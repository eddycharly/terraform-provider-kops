package datasources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube"
)

// KubeConfig represents the kubernets configuration needed to access a cluster
type KubeConfig struct {
	// The cluster name
	ClusterName string
	kube.Config
}
