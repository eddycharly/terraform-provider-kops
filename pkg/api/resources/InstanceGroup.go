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

func fromKopsInstanceGroup(clusterName string, instanceGroup *kops.InstanceGroup) *InstanceGroup {
	return &InstanceGroup{
		ClusterName:       clusterName,
		Name:              instanceGroup.ObjectMeta.Name,
		InstanceGroupSpec: instanceGroup.Spec,
	}
}

// func toKopsInstanceGroup(instanceGroup *InstanceGroup) *kops.InstanceGroup {
// 	return &kops.InstanceGroup{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: instanceGroup.Name,
// 		},
// 		Spec: instanceGroup.InstanceGroupSpec,
// 	}
// }

// func deleteInstanceGroup(cluster *kops.Cluster, instanceGroup *kops.InstanceGroup, clientset simple.Clientset) error {
// 	cloud, err := cloudup.BuildCloud(cluster)
// 	if err != nil {
// 		return err
// 	}
// 	d := &instancegroups.DeleteInstanceGroup{
// 		Cluster:   cluster,
// 		Cloud:     cloud,
// 		Clientset: clientset,
// 	}
// 	err = d.DeleteInstanceGroup(instanceGroup)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func syncInstanceGroups(cluster *Cluster, clientset simple.Clientset) error {
// 	kc, kig := toKopsCluster(cluster)
// 	synced := sets.NewString()
// 	for _, ig := range kig {
// 		_, err := clientset.InstanceGroupsFor(kc).Get(context.Background(), ig.Name, metav1.GetOptions{})
// 		if err != nil {
// 			if errors.IsNotFound(err) {
// 				_, err := clientset.InstanceGroupsFor(kc).Create(context.Background(), ig, metav1.CreateOptions{})
// 				if err != nil {
// 					return err
// 				}
// 			} else {
// 				return err
// 			}
// 		} else {
// 			_, err := clientset.InstanceGroupsFor(kc).Update(context.Background(), ig, metav1.UpdateOptions{})
// 			if err != nil {
// 				return err
// 			}
// 		}
// 		synced.Insert(ig.Name)
// 	}
// 	kc, kig, err := getClusterAndInstanceGroups(cluster.Name, clientset)
// 	if err != nil {
// 		return err
// 	}
// 	for _, ig := range kig {
// 		if !synced.Has(ig.Name) {
// 			err := deleteInstanceGroup(kc, ig, clientset)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

func GetInstanceGroup(clusterName, name string, clientset simple.Clientset) (*InstanceGroup, error) {
	cluster, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return nil, err
	}
	instanceGroup, err := clientset.InstanceGroupsFor(cluster).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return fromKopsInstanceGroup(clusterName, instanceGroup), nil
}
