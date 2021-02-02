package utils

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/pkg/instancegroups"
	"k8s.io/kops/upup/pkg/fi/cloudup"
)

func InstanceGroupDelete(clientset simple.Clientset, clusterName string, instanceGroupName string) error {
	cluster, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return err
	}
	instanceGroup, err := clientset.InstanceGroupsFor(cluster).Get(context.Background(), instanceGroupName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	cloud, err := cloudup.BuildCloud(cluster)
	if err != nil {
		return err
	}
	d := &instancegroups.DeleteInstanceGroup{
		Cluster:   cluster,
		Cloud:     cloud,
		Clientset: clientset,
	}
	err = d.DeleteInstanceGroup(instanceGroup)
	if err != nil {
		return err
	}
	return nil
}
