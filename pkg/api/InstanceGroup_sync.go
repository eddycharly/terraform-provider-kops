package api

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/pkg/instancegroups"
	"k8s.io/kops/upup/pkg/fi/cloudup"
)

func deleteInstanceGroup(cluster *kops.Cluster, instanceGroup *kops.InstanceGroup, clientset simple.Clientset) error {
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

func SyncInstanceGroups(cluster *Cluster, clientset simple.Clientset) error {
	kc, kig := ToKopsCluster(cluster)
	synced := sets.NewString()
	for _, ig := range kig {
		_, err := clientset.InstanceGroupsFor(kc).Get(context.Background(), ig.Name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				_, err := clientset.InstanceGroupsFor(kc).Create(context.Background(), ig, metav1.CreateOptions{})
				if err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			_, err := clientset.InstanceGroupsFor(kc).Update(context.Background(), ig, metav1.UpdateOptions{})
			if err != nil {
				return err
			}
		}
		synced.Insert(ig.Name)
	}
	kc, kig, err := getClusterAndInstanceGroups(cluster.Name, clientset)
	if err != nil {
		return err
	}
	for _, ig := range kig {
		if !synced.Has(ig.Name) {
			err := deleteInstanceGroup(kc, ig, clientset)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
