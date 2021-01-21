package resources

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/client/simple"
)

// InstanceGroup represents a group of instances (either nodes or masters) with the same configuration
type InstanceGroup struct {
	// The cluster name the instance group belongs to
	ClusterName string
	// Instance group name
	Name string
	kops.InstanceGroupSpec
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

func DeleteInstanceGroup(clusterName, name string, clientset simple.Clientset) error {
	cluster, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return err
	}
	if err := clientset.InstanceGroupsFor(cluster).Delete(context.Background(), name, metav1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}
