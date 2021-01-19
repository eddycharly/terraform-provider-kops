package resources

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/kubeconfig"
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

func fromKopsCluster(cluster *kops.Cluster, kubeConfig *kubeconfig.KubeconfigBuilder, instanceGroups ...*kops.InstanceGroup) *Cluster {
	return &Cluster{
		Name:        cluster.ObjectMeta.Name,
		ClusterSpec: cluster.Spec,
		// KubeConfig:  kube.FromKubeconfigBuilder(kubeConfig),
		// InstanceGroup: func(in ...*kops.InstanceGroup) []*InstanceGroup {
		// 	var out []*InstanceGroup
		// 	for _, in := range in {
		// 		out = append(out, fromKopsInstanceGroup(in))
		// 	}
		// 	return out
		// }(instanceGroups...),
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
