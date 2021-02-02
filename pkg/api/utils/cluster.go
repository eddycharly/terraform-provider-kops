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
	"k8s.io/kops/pkg/instancegroups"
	kopsValidation "k8s.io/kops/pkg/validation"
	"k8s.io/kops/upup/pkg/fi/cloudup"
)

func makeValidator(name string, clientset simple.Clientset) (kopsValidation.ClusterValidator, error) {
	kc, err := clientset.GetCluster(context.Background(), name)
	if err != nil {
		return nil, err
	}
	cloud, err := cloudup.BuildCloud(kc)
	if err != nil {
		return nil, err
	}
	list, err := clientset.InstanceGroupsFor(kc).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("cannot get InstanceGroups for %q: %v", kc.ObjectMeta.Name, err)
	}
	var instanceGroups []kops.InstanceGroup
	for _, ig := range list.Items {
		instanceGroups = append(instanceGroups, ig)
	}
	if len(instanceGroups) == 0 {
		return nil, fmt.Errorf("no InstanceGroup objects found")
	}
	configBuilder, err := GetKubeConfigBuilder(clientset, name)
	if err != nil {
		return nil, err
	}
	config, err := configBuilder.BuildRestConfig()
	if err != nil {
		return nil, err
	}
	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("cannot build kubernetes api client for %q: %v", kc.Name, err)
	}
	validator, err := validation.NewClusterValidator(kc, cloud, list, config, k8sClient)
	if err != nil {
		return nil, fmt.Errorf("unexpected error creating validatior: %v", err)
	}
	return validator, nil
}

func IsClusterValid(name string, clientset simple.Clientset) (bool, error) {
	if validator, err := makeValidator(name, clientset); err != nil {
		return false, err
	} else {
		result, err := validator.Validate()
		if err != nil {
			return false, err
		}
		if len(result.Failures) != 0 {
			return false, fmt.Errorf("Validation failures")
		}
		return true, nil
	}
}

func NeedsUpdate(name string, clientset simple.Clientset) (bool, error) {
	kc, err := clientset.GetCluster(context.Background(), name)
	if err != nil {
		return false, err
	}
	var k8sClient kubernetes.Interface
	var nodes []v1.Node
	configBuilder, err := GetKubeConfigBuilder(clientset, name)
	if err != nil {
		return false, err
	}
	config, err := configBuilder.BuildRestConfig()
	if err != nil {
		return false, err
	}
	k8sClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		return false, fmt.Errorf("cannot build kube client for %q: %v", kc.Name, err)
	}
	nodeList, err := k8sClient.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return false, err
	}
	if nodeList != nil {
		nodes = nodeList.Items
	}
	list, err := clientset.InstanceGroupsFor(kc).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return false, err
	}
	var instanceGroups []*kops.InstanceGroup
	for i := range list.Items {
		instanceGroups = append(instanceGroups, &list.Items[i])
	}
	cloud, err := cloudup.BuildCloud(kc)
	if err != nil {
		return false, err
	}
	groups, err := cloud.GetCloudGroups(kc, instanceGroups, false, nodes)
	if err != nil {
		return false, err
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
	err = d.AdjustNeedUpdate(groups, kc, list)
	if err != nil {
		return false, err
	}
	needUpdate := false
	for _, group := range groups {
		if len(group.NeedUpdate) != 0 {
			needUpdate = true
		}
	}
	return needUpdate, nil
}
