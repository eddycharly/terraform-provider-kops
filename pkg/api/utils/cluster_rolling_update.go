package utils

import (
	"context"
	"fmt"
	"time"

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

type RollingUpdateOptions struct {
	// MasterInterval is the amount of time to wait after stopping a master instance
	MasterInterval *metav1.Duration
	// NodeInterval is the amount of time to wait after stopping a non-master instance
	NodeInterval *metav1.Duration
	// BastionInterval is the amount of time to wait after stopping a bastion instance
	BastionInterval *metav1.Duration
	// FailOnDrainError will fail when a drain error occurs
	FailOnDrainError bool
	// FailOnValidate will fail when a validation error occurs
	FailOnValidate bool
	// PostDrainDelay is the duration we wait after draining each node
	PostDrainDelay *metav1.Duration
	// ValidationTimeout is the maximum time to wait for the cluster to validate, once we start validation
	ValidationTimeout *metav1.Duration
	// ValidateCount is the amount of time that a cluster needs to be validated after single node update
	ValidateCount *int
	// CloudOnly perform rolling update without confirming progress with k8s
	CloudOnly bool
}

func ClusterInstanceGroupsNeedingUpdate(clientset simple.Clientset, clusterName string) ([]string, error) {
	kc, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return nil, err
	}
	var k8sClient kubernetes.Interface
	var nodes []v1.Node
	configBuilder, err := GetKubeConfigBuilder(clientset, clusterName, nil, false)
	if err != nil {
		return nil, err
	}
	config, err := configBuilder.BuildRestConfig()
	if err != nil {
		return nil, err
	}
	k8sClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("cannot build kube client for %q: %v", kc.Name, err)
	}
	nodeList, err := k8sClient.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	if nodeList != nil {
		nodes = nodeList.Items
	}
	list, err := clientset.InstanceGroupsFor(kc).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var instanceGroups []*kops.InstanceGroup
	for i := range list.Items {
		instanceGroups = append(instanceGroups, &list.Items[i])
	}
	cloud, err := cloudup.BuildCloud(kc)
	if err != nil {
		return nil, err
	}
	groups, err := cloud.GetCloudGroups(kc, instanceGroups, false, nodes)
	if err != nil {
		return nil, err
	}
	d := &instancegroups.RollingUpdateCluster{
		Interactive:             false,
		Force:                   false,
		Cloud:                   cloud,
		K8sClient:               k8sClient,
		CloudOnly:               false,
		ClusterName:             kc.Name,
		ValidateTickDuration:    30 * time.Second,
		ValidateSuccessDuration: 10 * time.Second,
	}
	err = d.AdjustNeedUpdate(groups)
	if err != nil {
		return nil, err
	}
	var needUpdate []string
	for _, group := range groups {
		if len(group.NeedUpdate) != 0 {
			needUpdate = append(needUpdate, group.InstanceGroup.Name)
		}
	}
	return needUpdate, nil
}

func ClusterRollingUpdate(clientset simple.Clientset, clusterName string, options RollingUpdateOptions) error {
	kc, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return err
	}
	var k8sClient kubernetes.Interface
	var nodes []v1.Node
	configBuilder, err := GetKubeConfigBuilder(clientset, clusterName, nil, false)
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
		Ctx:                     context.TODO(),
		Cluster:                 kc,
		Clientset:               clientset,
		MasterInterval:          MasterInterval,
		NodeInterval:            NodeInterval,
		BastionInterval:         BastionInterval,
		Interactive:             false,
		Force:                   false,
		Cloud:                   cloud,
		K8sClient:               k8sClient,
		FailOnDrainError:        options.FailOnDrainError,
		FailOnValidate:          options.FailOnValidate,
		CloudOnly:               options.CloudOnly,
		ClusterName:             kc.Name,
		PostDrainDelay:          PostDrainDelay,
		ValidationTimeout:       ValidationTimeout,
		ValidateCount:           ValidateCount,
		ValidateTickDuration:    30 * time.Second,
		ValidateSuccessDuration: 10 * time.Second,
	}
	err = d.AdjustNeedUpdate(groups)
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
	return d.RollingUpdate(groups, list)
}
