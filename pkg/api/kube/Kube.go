package kube

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/utils"
	"k8s.io/kops/pkg/client/simple"
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

func GetConfig(clientset simple.Clientset, clusterName string) (*Config, error) {
	conf, err := utils.GetKubeConfigBuilder(clientset, clusterName)
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
