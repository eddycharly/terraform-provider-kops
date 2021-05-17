package kube

import (
	"time"

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

func (s *Config) GetConfig(clientset simple.Clientset, clusterName string, admin *time.Duration, internal bool) error {
	conf, err := utils.GetKubeConfigBuilder(clientset, clusterName, admin, internal)
	if err != nil {
		return err
	}
	s.Server = conf.Server
	s.Context = conf.Context
	s.Namespace = conf.Namespace
	s.KubeUser = conf.KubeUser
	s.KubePassword = conf.KubePassword
	s.CaCert = string(conf.CACert)
	s.ClientCert = string(conf.ClientCert)
	s.ClientKey = string(conf.ClientKey)
	return nil
}
