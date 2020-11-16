package datasources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube"
)

// Cluster defines the configuration for a cluster
// It includes cluster instance groups.
type KubeConfig struct {
	// The cluster name
	ClusterName string
	kube.Config
}
