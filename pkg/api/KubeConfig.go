package api

import "k8s.io/kops/pkg/kubeconfig"

type KubeConfig struct {
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

func fromKubeconfigBuilder(kubeConfig *kubeconfig.KubeconfigBuilder) *KubeConfig {
	return &KubeConfig{
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
