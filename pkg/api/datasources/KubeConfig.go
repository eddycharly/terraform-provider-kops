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

func GetConfig(name string, clientset simple.Clientset) (*KubeConfig, error) {
	config, err := kube.GetConfig(name, clientset)
	if err != nil {
		return nil, err
	}
	return &KubeConfig{
		ClusterName: name,
		Config:      *config,
	}, nil
}
