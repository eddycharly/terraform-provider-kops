package resources

import "k8s.io/kops/pkg/client/simple"

// ClusterUpdater takes care of rolling update the cluster when needed
type ClusterUpdater struct {
	// The target cluster name
	ClusterName string
}

func (u *ClusterUpdater) Apply(clientset simple.Clientset) error {
	return nil
}
