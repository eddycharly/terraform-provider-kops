package datasources

import (
	"time"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube"
	"k8s.io/kops/pkg/client/simple"
)

// KubeConfig represents the kubernets configuration needed to access a cluster
type KubeConfig struct {
	// The cluster name
	ClusterName string
	// Admin is the cluster admin user credential lifetime
	Admin *time.Duration
	// Internal use the cluster's internal DNS name
	Internal bool
	kube.Config
}

func (s *KubeConfig) GetKubeConfig(clientset simple.Clientset) error {
	if err := s.Config.GetConfig(clientset, s.ClusterName, s.Admin, s.Internal); err != nil {
		return err
	}
	return nil
}
