package kube

import "k8s.io/kops/pkg/kubeconfig"

type Config struct {
	Server          string
	Context         string
	Namespace       string
	KubeBearerToken string
	KubeUser        string
	KubePassword    string
	CaCert          string
	ClientCert      string
	ClientKey       string
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
