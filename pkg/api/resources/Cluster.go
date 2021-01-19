package resources

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/upup/pkg/fi"
)

// Cluster defines the configuration for a cluster
// It includes cluster instance groups.
type Cluster struct {
	// The cluster name
	Name string
	// The cluster admin ssh key
	AdminSshKey string
	kops.ClusterSpec
}

func fromKopsCluster(cluster *kops.Cluster) *Cluster {
	return &Cluster{
		Name:        cluster.ObjectMeta.Name,
		ClusterSpec: cluster.Spec,
	}
}

func toKopsCluster(cluster *Cluster) (*kops.Cluster, []*kops.InstanceGroup) {
	c := kops.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: cluster.Name,
		},
		Spec: cluster.ClusterSpec,
	}
	// var ig []*kops.InstanceGroup
	// for _, instanceGroup := range cluster.InstanceGroup {
	// 	if instanceGroup != nil && instanceGroup.Name != "" {
	// 		ig = append(ig, toKopsInstanceGroup(instanceGroup))
	// 	}
	// }
	// return &c, ig
	return &c, nil
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
	cluster := fromKopsCluster(kc)
	if len(pubKeys) > 0 {
		cluster.AdminSshKey = pubKeys[0].Spec.PublicKey
	}
	return cluster, nil
}
