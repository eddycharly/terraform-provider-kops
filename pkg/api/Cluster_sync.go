package api

import (
	"context"
	"fmt"
	"io/ioutil"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/pkg/resources"
	"k8s.io/kops/pkg/resources/ops"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/cloudup"
	"k8s.io/kops/upup/pkg/fi/utils"
)

func getClusterAndInstanceGroups(name string, clientset simple.Clientset) (*kops.Cluster, []kops.InstanceGroup, error) {
	cluster, err := clientset.GetCluster(context.Background(), name)
	if err != nil {
		return nil, nil, err
	}
	ig := clientset.InstanceGroupsFor(cluster)
	instanceGroups, err := ig.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}
	return cluster, instanceGroups.Items, err
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
	c := FromKopsCluster(*kc, kig...)
	return &c, nil
}

func SyncCluster(cluster Cluster, clientset simple.Clientset) (*Cluster, error) {
	exists, err := ClusterExists(cluster.Name, clientset)
	if err != nil {
		return nil, err
	}
	kc, kig := ToKopsCluster(cluster)
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
	for _, ig := range kig {
		_, err = clientset.InstanceGroupsFor(kc).Get(context.Background(), ig.Name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				_, err = clientset.InstanceGroupsFor(kc).Create(context.Background(), ig, metav1.CreateOptions{})
				if err != nil {
					return nil, err
				}
			}
			return nil, err
		} else {
			_, err = clientset.InstanceGroupsFor(kc).Update(context.Background(), ig, metav1.UpdateOptions{})
			if err != nil {
				return nil, err
			}
		}
	}
	apply := &cloudup.ApplyClusterCmd{
		Cluster:    kc,
		Clientset:  clientset,
		TargetName: cloudup.TargetDirect,
	}
	err = apply.Run(context.Background())
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
