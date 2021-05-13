package utils

import (
	"context"
	"time"

	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/pkg/commands"
	"k8s.io/kops/pkg/kubeconfig"
)

func GetKubeConfigBuilder(clientset simple.Clientset, clusterName string, admin *time.Duration, internal bool) (*kubeconfig.KubeconfigBuilder, error) {
	cluster, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return nil, err
	}
	keyStore, err := clientset.KeyStore(cluster)
	if err != nil {
		return nil, err
	}
	secretStore, err := clientset.SecretStore(cluster)
	if err != nil {
		return nil, err
	}
	duration := kubeconfig.DefaultKubecfgAdminLifetime
	if admin != nil {
		duration = *admin
	}
	conf, err := kubeconfig.BuildKubecfg(cluster, keyStore, secretStore, &commands.CloudDiscoveryStatusStore{}, duration, "", internal, "", false)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
