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

func GetConfig(clientset simple.Clientset, clusterName string) (*KubeConfig, error) {
	config, err := kube.GetConfig(clientset, clusterName)
	if err != nil {
		return nil, err
	}
	return &KubeConfig{
		ClusterName: clusterName,
		Config:      *config,
	}, nil
}
