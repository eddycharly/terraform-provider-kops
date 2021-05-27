package resources

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/pkg/resources"
	"k8s.io/kops/pkg/resources/ops"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/cloudup"
)

// Cluster defines the configuration for a cluster
type Cluster struct {
	// Revision is incremented every time the resource changes, this is useful for triggering cluster updater
	Revision int
	// Name defines the cluster name
	Name string
	// AdminSshKey defines the cluster admin ssh key
	AdminSshKey string
	kops.ClusterSpec
}

func makeCluster(adminSshKey string, cluster *kops.Cluster) *Cluster {
	return &Cluster{
		Name:        cluster.ObjectMeta.Name,
		AdminSshKey: adminSshKey,
		ClusterSpec: cluster.Spec,
	}
}

func makeKopsCluster(name string, spec kops.ClusterSpec) *kops.Cluster {
	return &kops.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: spec,
	}
}

func GetCluster(name string, clientset simple.Clientset) (*Cluster, error) {
	kc, err := clientset.GetCluster(context.Background(), name)
	if err != nil {
		return nil, err
	}
	sshCredentialStore, err := clientset.SSHCredentialStore(kc)
	if err != nil {
		return nil, err
	}
	pubKeys, err := sshCredentialStore.FindSSHPublicKeys(fi.SecretNameSSHPrimary)
	if err != nil {
		return nil, err
	}
	cluster := makeCluster(pubKeys[0].Spec.PublicKey, kc)
	return cluster, nil
}

func CreateCluster(name, adminSshKey string, spec kops.ClusterSpec, clientset simple.Clientset) (*Cluster, error) {
	kc := makeKopsCluster(name, spec)
	cloud, err := cloudup.BuildCloud(kc)
	if err != nil {
		return nil, err
	}
	if err := cloudup.PerformAssignments(kc, cloud); err != nil {
		return nil, err
	}
	kc, err = clientset.CreateCluster(context.Background(), kc)
	if err != nil {
		return nil, err
	}
	kc, err = clientset.GetCluster(context.Background(), name)
	if err != nil {
		return nil, err
	}
	sshCredentialStore, err := clientset.SSHCredentialStore(kc)
	if err != nil {
		return nil, err
	}
	pubKey := []byte(adminSshKey)
	err = sshCredentialStore.AddSSHPublicKey(fi.SecretNameSSHPrimary, pubKey)
	if err != nil {
		return nil, fmt.Errorf("error adding SSH public key: %v", err)
	}
	kc, err = clientset.GetCluster(context.Background(), name)
	if err != nil {
		return nil, err
	}
	return makeCluster(adminSshKey, kc), nil
}

func UpdateCluster(name, adminSshKey string, spec kops.ClusterSpec, clientset simple.Clientset) (*Cluster, error) {
	kc := makeKopsCluster(name, spec)
	cloud, err := cloudup.BuildCloud(kc)
	if err != nil {
		return nil, err
	}
	if err := cloudup.PerformAssignments(kc, cloud); err != nil {
		return nil, err
	}
	kc, err = clientset.UpdateCluster(context.Background(), kc, nil)
	if err != nil {
		return nil, err
	}
	sshCredentialStore, err := clientset.SSHCredentialStore(kc)
	if err != nil {
		return nil, err
	}
	pubKey := []byte(adminSshKey)
	err = sshCredentialStore.AddSSHPublicKey(fi.SecretNameSSHPrimary, pubKey)
	if err != nil {
		return nil, fmt.Errorf("error adding SSH public key: %v", err)
	}
	kc, err = clientset.GetCluster(context.Background(), name)
	if err != nil {
		return nil, err
	}
	return makeCluster(adminSshKey, kc), nil
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
	allResources, err := ops.ListResources(cloud, kc, "")
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
