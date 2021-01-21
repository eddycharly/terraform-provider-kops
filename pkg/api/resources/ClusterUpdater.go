package resources

import (
	"context"
	"fmt"
	"time"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/utils"
	"github.com/eddycharly/terraform-provider-kops/pkg/validation"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/pkg/cloudinstances"
	"k8s.io/kops/pkg/instancegroups"
	"k8s.io/kops/upup/pkg/fi/cloudup"
)

// ClusterUpdater takes care of rolling update the cluster when needed
type ClusterUpdater struct {
	// The target cluster name
	ClusterName string
}

func (u *ClusterUpdater) Apply(clientset simple.Clientset, options config.RollingUpdate) error {
	if options.Skip {
		return nil
	}
	kc, err := clientset.GetCluster(context.Background(), u.ClusterName)
	if err != nil {
		return err
	}
	var k8sClient kubernetes.Interface
	var nodes []v1.Node
	configBuilder, err := utils.GetKubeConfigBuilder(u.ClusterName, clientset)
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
	PostDrainDelay := 5 * time.Second
	ValidationTimeout := 15 * time.Minute
	ValidateCount := 2
	if options.MasterInterval != nil {
		MasterInterval = options.MasterInterval.Duration
	}
	if options.NodeInterval != nil {
		NodeInterval = options.NodeInterval.Duration
	}
	if options.BastionInterval != nil {
		BastionInterval = options.BastionInterval.Duration
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
	d := &instancegroups.RollingUpdateCluster{
		MasterInterval:          MasterInterval,
		NodeInterval:            NodeInterval,
		BastionInterval:         BastionInterval,
		Interactive:             false,
		Force:                   false,
		Cloud:                   cloud,
		K8sClient:               k8sClient,
		FailOnDrainError:        options.FailOnDrainError,
		FailOnValidate:          options.FailOnValidate,
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
	clusterValidator, err := validation.NewClusterValidator(kc, cloud, list, config, k8sClient)
	if err != nil {
		return fmt.Errorf("cannot create cluster validator: %v", err)
	}
	d.ClusterValidator = clusterValidator
	return d.RollingUpdate(context.Background(), groups, kc, list)
}
