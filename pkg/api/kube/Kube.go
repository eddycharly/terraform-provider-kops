package kube

import "k8s.io/kops/pkg/kubeconfig"

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

func FromKubeconfigBuilder(kubeConfig *kubeconfig.KubeconfigBuilder) *Config {
	return &Config{
		Server:          kubeConfig.Server,
		Context:         kubeConfig.Context,
		Namespace:       kubeConfig.Namespace,
		KubeBearerToken: kubeConfig.KubeBearerToken,
		KubeUser:        kubeConfig.KubeUser,
		KubePassword:    kubeConfig.KubePassword,
		CaCert:          string(kubeConfig.CACert),
		ClientCert:      string(kubeConfig.ClientCert),
		ClientKey:       string(kubeConfig.ClientKey),
	}
}
