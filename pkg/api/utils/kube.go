package utils

import (
	"context"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/pkg/commands"
	"k8s.io/kops/pkg/kubeconfig"
)

func GetKubeConfigBuilder(clientset simple.Clientset, clusterName string) (*kubeconfig.KubeconfigBuilder, error) {
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
	conf, err := kubeconfig.BuildKubecfg(cluster, keyStore, secretStore, &commands.CloudDiscoveryStatusStore{}, clientcmd.NewDefaultPathOptions())
	if err != nil {
		return nil, err
	}
	return conf, nil
}
