package resources

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/client/simple"
)

// InstanceGroup represents a group of instances (either bastions, nodes or masters) with the same configuration
type InstanceGroup struct {
	kops.InstanceGroupSpec
	// Revision is incremented every time the resource changes, this is useful for triggering cluster updater
	Revision int
	// ClusterName defines the cluster name the instance group belongs to
	ClusterName string
	// Name defines the instance group name
	Name string
}

func makeInstanceGroup(clusterName string, instanceGroup *kops.InstanceGroup) *InstanceGroup {
	return &InstanceGroup{
		ClusterName:       clusterName,
		Name:              instanceGroup.ObjectMeta.Name,
		InstanceGroupSpec: instanceGroup.Spec,
	}
}

func makeKopsInstanceGroup(name string, spec kops.InstanceGroupSpec) *kops.InstanceGroup {
	return &kops.InstanceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: spec,
	}
}

func GetInstanceGroup(clusterName, name string, clientset simple.Clientset) (*InstanceGroup, error) {
	cluster, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return nil, err
	}
	instanceGroup, err := clientset.InstanceGroupsFor(cluster).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return makeInstanceGroup(clusterName, instanceGroup), nil
}

func CreateInstanceGroup(clusterName, name string, spec kops.InstanceGroupSpec, clientset simple.Clientset) (*InstanceGroup, error) {
	cluster, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return nil, err
	}
	instanceGroup, err := clientset.InstanceGroupsFor(cluster).Create(context.Background(), makeKopsInstanceGroup(name, spec), metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return makeInstanceGroup(clusterName, instanceGroup), nil
}

func UpdateInstanceGroup(clusterName, name string, spec kops.InstanceGroupSpec, clientset simple.Clientset) (*InstanceGroup, error) {
	cluster, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return nil, err
	}
	instanceGroup, err := clientset.InstanceGroupsFor(cluster).Update(context.Background(), makeKopsInstanceGroup(name, spec), metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return makeInstanceGroup(clusterName, instanceGroup), nil
}
