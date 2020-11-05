package api

import (
	"context"
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/pkg/cloudinstances"
	"k8s.io/kops/pkg/commands"
	"k8s.io/kops/pkg/instancegroups"
	"k8s.io/kops/pkg/kubeconfig"
	"k8s.io/kops/pkg/resources"
	"k8s.io/kops/pkg/resources/ops"
	"k8s.io/kops/pkg/validation"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/cloudup"
)

func getClusterAndInstanceGroups(name string, clientset simple.Clientset) (*kops.Cluster, []*kops.InstanceGroup, error) {
	cluster, err := clientset.GetCluster(context.Background(), name)
	if err != nil {
		return nil, nil, err
	}
	igs := clientset.InstanceGroupsFor(cluster)
	instanceGroups, err := igs.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}
	var kigs []*kops.InstanceGroup
	for _, ig := range instanceGroups.Items {
		x := ig
		kigs = append(kigs, &x)
	}
	return cluster, kigs, err
}

func ClusterExists(name string, clientset simple.Clientset) (bool, error) {
	_, err := clientset.GetCluster(context.Background(), name)
	if err != nil {
		if errors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func GetCluster(name string, clientset simple.Clientset) (*Cluster, error) {
	kc, kig, err := getClusterAndInstanceGroups(name, clientset)
	if err != nil {
		return nil, err
	}
	kube, err := getKubeConfig(name, clientset)
	if err != nil {
		return nil, err
	}
	return fromKopsCluster(kc, kube, kig...), nil
}

func SyncCluster(cluster *Cluster, clientset simple.Clientset) (*Cluster, error) {
	exists, err := ClusterExists(cluster.Name, clientset)
	if err != nil {
		return nil, err
	}
	doRollingUpdate := false
	doValidateCluster := true

	kc, _ := toKopsCluster(cluster)
	if err := cloudup.PerformAssignments(kc); err != nil {
		return nil, err
	}
	if exists {
		kc, err = clientset.UpdateCluster(context.Background(), kc, nil)
		if err != nil {
			return nil, err
		}
		doRollingUpdate = true
	} else {
		kc, err = clientset.CreateCluster(context.Background(), kc)
		if err != nil {
			return nil, err
		}
		kc, err = clientset.GetCluster(context.Background(), cluster.Name)
		if err != nil {
			return nil, err
		}
	}
	sshCredentialStore, err := clientset.SSHCredentialStore(kc)
	if err != nil {
		return nil, err
	}
	pubKey := []byte(cluster.AdminSshKey)
	err = sshCredentialStore.AddSSHPublicKey(fi.SecretNameSSHPrimary, pubKey)
	if err != nil {
		return nil, fmt.Errorf("error adding SSH public key: %v", err)
	}
	kc, err = clientset.GetCluster(context.Background(), cluster.Name)
	if err != nil {
		return nil, err
	}
	SyncInstanceGroups(cluster, clientset)
	apply := &cloudup.ApplyClusterCmd{
		Cluster:    kc,
		Clientset:  clientset,
		TargetName: cloudup.TargetDirect,
	}
	err = apply.Run(context.Background())
	if err != nil {
		return nil, err
	}
	if doRollingUpdate {
		err = rollingUpdate(cluster.Name, clientset, cluster.RollingUpdateOptions)
		if err != nil {
			return nil, err
		}
	}
	if doValidateCluster {
		err = validateCluster(cluster.Name, clientset)
		if err != nil {
			return nil, err
		}
	}
	return GetCluster(cluster.Name, clientset)
}

func DeleteCluster(name string, clientset simple.Clientset) error {
	kc, err := clientset.GetCluster(context.Background(), name)
	if err != nil {
		return err
	}
	cloud, err := cloudup.BuildCloud(kc)
	if err != nil {
		return err
	}
	allResources, err := ops.ListResources(cloud, kc.Name, "")
	if err != nil {
		return err
	}
	clusterResources := make(map[string]*resources.Resource)
	for k, resource := range allResources {
		if resource.Shared {
			continue
		}
		clusterResources[k] = resource
	}
	if len(clusterResources) != 0 {
		var l []*resources.Resource
		for _, v := range clusterResources {
			l = append(l, v)
		}
		err = ops.DeleteResources(cloud, clusterResources)
		if err != nil {
			return err
		}
	}
	err = clientset.DeleteCluster(context.Background(), kc)
	if err != nil {
		return err
	}
	return nil
}

func rollingUpdate(name string, clientset simple.Clientset, options *RollingUpdateOptions) error {
	kc, err := clientset.GetCluster(context.Background(), name)
	if err != nil {
		return err
	}
	var k8sClient kubernetes.Interface
	var nodes []v1.Node
	configBuilder, err := getKubeConfig(name, clientset)
	if err != nil {
		return err
	}
	config, err := configBuilder.BuildRestConfig()
	if err != nil {
		return err
	}
	k8sClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("cannot build kube client for %q: %v", kc.Name, err)
	}
	nodeList, err := k8sClient.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	if nodeList != nil {
		nodes = nodeList.Items
	}
	list, err := clientset.InstanceGroupsFor(kc).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	var instanceGroups []*kops.InstanceGroup
	for i := range list.Items {
		instanceGroups = append(instanceGroups, &list.Items[i])
	}
	cloud, err := cloudup.BuildCloud(kc)
	if err != nil {
		return err
	}
	groups, err := cloud.GetCloudGroups(kc, instanceGroups, false, nodes)
	if err != nil {
		return err
	}
	MasterInterval := 15 * time.Second
	NodeInterval := 15 * time.Second
	BastionInterval := 15 * time.Second
	FailOnDrainError := false
	FailOnValidate := false
	PostDrainDelay := 5 * time.Second
	ValidationTimeout := 15 * time.Minute
	ValidateCount := 2
	if options != nil {
		if options.MasterInterval != nil {
			MasterInterval = options.MasterInterval.Duration
		}
		if options.NodeInterval != nil {
			NodeInterval = options.NodeInterval.Duration
		}
		if options.BastionInterval != nil {
			BastionInterval = options.BastionInterval.Duration
		}
		if options.FailOnDrainError != nil {
			FailOnDrainError = *options.FailOnDrainError
		}
		if options.FailOnValidate != nil {
			FailOnValidate = *options.FailOnValidate
		}
		if options.PostDrainDelay != nil {
			PostDrainDelay = options.PostDrainDelay.Duration
		}
		if options.ValidationTimeout != nil {
			ValidationTimeout = options.ValidationTimeout.Duration
		}
		if options.ValidateCount != nil {
			ValidateCount = *options.ValidateCount
		}
	}
	d := &instancegroups.RollingUpdateCluster{
		MasterInterval:          MasterInterval,
		NodeInterval:            NodeInterval,
		BastionInterval:         BastionInterval,
		Interactive:             false,
		Force:                   false,
		Cloud:                   cloud,
		K8sClient:               k8sClient,
		FailOnDrainError:        FailOnDrainError,
		FailOnValidate:          FailOnValidate,
		CloudOnly:               false,
		ClusterName:             kc.Name,
		PostDrainDelay:          PostDrainDelay,
		ValidationTimeout:       ValidationTimeout,
		ValidateCount:           ValidateCount,
		ValidateTickDuration:    30 * time.Second,
		ValidateSuccessDuration: 10 * time.Second,
	}
	err = d.AdjustNeedUpdate(groups, kc, list)
	if err != nil {
		return err
	}
	var l []*cloudinstances.CloudInstanceGroup
	for _, v := range groups {
		l = append(l, v)
	}
	needUpdate := false
	for _, group := range groups {
		if len(group.NeedUpdate) != 0 {
			needUpdate = true
		}
	}
	if !needUpdate {
		return nil
	}
	clusterValidator, err := validation.NewClusterValidator(kc, cloud, list, k8sClient)
	if err != nil {
		return fmt.Errorf("cannot create cluster validator: %v", err)
	}
	d.ClusterValidator = clusterValidator
	return d.RollingUpdate(context.Background(), groups, kc, list)
}

func validateCluster(name string, clientSet simple.Clientset) error {
	kc, err := clientSet.GetCluster(context.Background(), name)
	if err != nil {
		return err
	}
	cloud, err := cloudup.BuildCloud(kc)
	if err != nil {
		return err
	}

	list, err := clientSet.InstanceGroupsFor(kc).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("cannot get InstanceGroups for %q: %v", kc.ObjectMeta.Name, err)
	}

	var instanceGroups []kops.InstanceGroup
	for _, ig := range list.Items {
		instanceGroups = append(instanceGroups, ig)
	}

	if len(instanceGroups) == 0 {
		return fmt.Errorf("no InstanceGroup objects found")
	}

	configBuilder, err := getKubeConfig(name, clientSet)
	if err != nil {
		return err
	}
	config, err := configBuilder.BuildRestConfig()
	if err != nil {
		return err
	}
	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("cannot build kubernetes api client for %q: %v", kc.Name, err)
	}

	timeout := time.Now().Add(15 * time.Minute)
	pollInterval := 10 * time.Second

	validator, err := validation.NewClusterValidator(kc, cloud, list, k8sClient)
	if err != nil {
		return fmt.Errorf("unexpected error creating validatior: %v", err)
	}

	consecutive := 0
	for {
		if time.Now().After(timeout) {
			return fmt.Errorf("wait time exceeded during validation")
		}

		result, err := validator.Validate()
		if err != nil {
			consecutive = 0
			klog.Warningf("(will retry): unexpected error during validation: %v", err)
			time.Sleep(pollInterval)
			continue
		}

		if len(result.Failures) == 0 {
			consecutive++
			if consecutive < 0 {
				klog.Infof("(will retry): cluster passed validation %d consecutive times", consecutive)
				time.Sleep(pollInterval)
				continue
			} else {
				return nil
			}
		} else {
			if consecutive == 0 {
				klog.Warningf("(will retry): cluster not yet healthy")
				time.Sleep(pollInterval)
				continue
			}
		}
	}
}

func getKubeConfig(name string, clientset simple.Clientset) (*kubeconfig.KubeconfigBuilder, error) {
	cluster, err := clientset.GetCluster(context.Background(), name)
	if err != nil {
		return nil, err
	}
	keyStore, err := clientset.KeyStore(cluster)
	if err != nil {
		return nil, err
	}
	secretStore, err := clientset.SecretStore(cluster)
	if err != nil {
		return nil, err
	}
	conf, err := kubeconfig.BuildKubecfg(cluster, keyStore, secretStore, &commands.CloudDiscoveryStatusStore{}, clientcmd.NewDefaultPathOptions())
	if err != nil {
		return nil, err
	}

	return conf, nil
}
