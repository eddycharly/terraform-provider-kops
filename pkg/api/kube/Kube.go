package kube

import (
	"context"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/pkg/commands"
	"k8s.io/kops/pkg/kubeconfig"
)

type Config struct {
	// Kubernetes server url
	Server string
	// Kubernetes context
	Context string
	// Kubernetes namespace
	Namespace string
	// Kubernetes bearer token
	KubeBearerToken string
	// Kubernetes user
	KubeUser string
	// Kubernetes password
	KubePassword string
	// Kubernetes cluster certificate
	CaCert string
	// Kubernetes client certificate
	ClientCert string
	// Kubernetes client key
	ClientKey string
}

func GetConfig(name string, clientset simple.Clientset) (*Config, error) {
	cluster, err := clientset.GetCluster(context.Background(), name)
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
	return &Config{
		Server:          conf.Server,
		Context:         conf.Context,
		Namespace:       conf.Namespace,
		KubeBearerToken: conf.KubeBearerToken,
		KubeUser:        conf.KubeUser,
		KubePassword:    conf.KubePassword,
		CaCert:          string(conf.CACert),
		ClientCert:      string(conf.ClientCert),
		ClientKey:       string(conf.ClientKey),
	}, nil
}
