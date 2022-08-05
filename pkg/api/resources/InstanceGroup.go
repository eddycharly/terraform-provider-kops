package resources

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/upup/pkg/fi/cloudup"
)

// InstanceGroup represents a group of instances (either bastions, nodes or masters) with the same configuration
type InstanceGroup struct {
	kops.InstanceGroupSpec
	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and services.
	Labels map[string]string
	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata. They are not
	// queryable and should be preserved when modifying objects.
	Annotations map[string]string
	// ClusterName defines the cluster name the instance group belongs to
	ClusterName string
	// Name defines the instance group name
	Name string
	// Revision is incremented every time the resource changes, this is useful for triggering cluster updater
	Revision int
}

func makeInstanceGroup(clusterName string, instanceGroup *kops.InstanceGroup) *InstanceGroup {
	labels := instanceGroup.Labels
	if labels != nil {
		delete(labels, kops.LabelClusterName)
	}
	return &InstanceGroup{
		InstanceGroupSpec: instanceGroup.Spec,
		Labels:            labels,
		Annotations:       instanceGroup.Annotations,
		ClusterName:       clusterName,
		Name:              instanceGroup.ObjectMeta.Name,
	}
}

func makeKopsInstanceGroup(name string, labels map[string]string, annotations map[string]string, spec kops.InstanceGroupSpec) *kops.InstanceGroup {
	return &kops.InstanceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Labels:      labels,
			Annotations: annotations,
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

func CreateInstanceGroup(clusterName, name string, labels map[string]string, annotations map[string]string, spec kops.InstanceGroupSpec, clientset simple.Clientset) (*InstanceGroup, error) {
	cluster, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return nil, err
	}
	ig := makeKopsInstanceGroup(name, labels, annotations, spec)
	channel, err := cloudup.ChannelForCluster(cluster)
	if err != nil {
		return nil, err
	}
	cloud, err := cloudup.BuildCloud(cluster)
	if err != nil {
		return nil, err
	}
	ig, err = cloudup.PopulateInstanceGroupSpec(cluster, ig, cloud, channel)
	if err != nil {
		return nil, err
	}
	ig, err = clientset.InstanceGroupsFor(cluster).Create(context.Background(), ig, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return makeInstanceGroup(clusterName, ig), nil
}

func UpdateInstanceGroup(clusterName, name string, labels map[string]string, annotations map[string]string, spec kops.InstanceGroupSpec, clientset simple.Clientset) (*InstanceGroup, error) {
	cluster, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return nil, err
	}
	ig := makeKopsInstanceGroup(name, labels, annotations, spec)
	channel, err := cloudup.ChannelForCluster(cluster)
	if err != nil {
		return nil, err
	}
	cloud, err := cloudup.BuildCloud(cluster)
	if err != nil {
		return nil, err
	}
	ig, err = cloudup.PopulateInstanceGroupSpec(cluster, ig, cloud, channel)
	if err != nil {
		return nil, err
	}
	ig, err = clientset.InstanceGroupsFor(cluster).Update(context.Background(), ig, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return makeInstanceGroup(clusterName, ig), nil
}
