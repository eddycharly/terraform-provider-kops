package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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
	"k8s.io/kops/upup/pkg/fi/utils"
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
	kube, err := getKubeCertificateAndToken(name, clientset)
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
	kc, _ := toKopsCluster(cluster)
	if err := cloudup.PerformAssignments(kc); err != nil {
		return nil, err
	}
	if exists {
		kc, err = clientset.UpdateCluster(context.Background(), kc, nil)
		if err != nil {
			return nil, err
		}
	} else {
		kc, err = clientset.CreateCluster(context.Background(), kc)
		if err != nil {
			return nil, err
		}
		kc, err = clientset.GetCluster(context.Background(), cluster.Name)
		if err != nil {
			return nil, err
		}
		// TODO improve this part
		sshCredentialStore, err := clientset.SSHCredentialStore(kc)
		if err != nil {
			return nil, err
		}
		f := utils.ExpandPath(*cluster.SSHKeyName)
		pubKey, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, fmt.Errorf("error reading SSH key file %q: %v", f, err)
		}
		err = sshCredentialStore.AddSSHPublicKey(fi.SecretNameSSHPrimary, pubKey)
		if err != nil {
			return nil, fmt.Errorf("error adding SSH public key: %v", err)
		}
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
	err = rollingUpdate(cluster.Name, clientset)
	if err != nil {
		return nil, err
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
	// TODO shall we get region from cluster spec ?
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

func rollingUpdate(name string, clientset simple.Clientset) error {
	kc, err := clientset.GetCluster(context.Background(), name)
	if err != nil {
		return err
	}
	var k8sClient kubernetes.Interface
	var nodes []v1.Node

	// TODO cloud only
	// TODO validate

	// k8sClient, err = kubernetes.NewForConfig(config)
	// if err != nil {
	// 	return fmt.Errorf("cannot build kube client for %q: %v", kc.Name, err)
	// }
	// nodeList, err := k8sClient.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	// if err != nil {
	// 	return fmt.Errorf("error listing nodes in cluster: %v", err)
	// }
	// if nodeList != nil {
	// 	nodes = nodeList.Items
	// }
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
	// TODO make this configurable
	d := &instancegroups.RollingUpdateCluster{
		MasterInterval:          15 * time.Second,
		NodeInterval:            15 * time.Second,
		BastionInterval:         15 * time.Second,
		Interactive:             false,
		Force:                   false,
		Cloud:                   cloud,
		K8sClient:               k8sClient,
		FailOnDrainError:        false,
		FailOnValidate:          false,
		CloudOnly:               true,
		ClusterName:             kc.Name,
		PostDrainDelay:          5 * time.Second,
		ValidationTimeout:       15 * time.Minute,
		ValidateCount:           2,
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
	var clusterValidator validation.ClusterValidator
	if false /*!options.CloudOnly*/ {
		clusterValidator, err = validation.NewClusterValidator(kc, cloud, list, k8sClient)
		if err != nil {
			return fmt.Errorf("cannot create cluster validator: %v", err)
		}
	}
	d.ClusterValidator = clusterValidator
	return d.RollingUpdate(context.Background(), groups, kc, list)
}

func getKubeCertificateAndToken(name string, clientset simple.Clientset) (*kubeconfig.KubeconfigBuilder, error) {
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
