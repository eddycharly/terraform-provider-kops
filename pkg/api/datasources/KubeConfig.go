package datasources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube"
	"k8s.io/kops/pkg/client/simple"
)

// KubeConfig represents the kubernets configuration needed to access a cluster
type KubeConfig struct {
	// The cluster name
	ClusterName string
	kube.Config
}

func (s *KubeConfig) GetKubeConfig(clientset simple.Clientset) error {
	if err := s.Config.GetConfig(clientset, s.ClusterName); err != nil {
		return err
	}
	return nil
}
