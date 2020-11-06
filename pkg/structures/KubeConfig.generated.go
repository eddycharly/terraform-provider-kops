package structures

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api"
)

func ExpandKubeConfig(in map[string]interface{}) api.KubeConfig {
	if in == nil {
		panic("expand KubeConfig failure, in is nil")
	}
	return api.KubeConfig{
		Server: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["server"]),
		Context: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["context"]),
		Namespace: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["namespace"]),
		KubeBearerToken: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kube_bearer_token"]),
		KubeUser: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kube_user"]),
		KubePassword: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kube_password"]),
		CaCert: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ca_cert"]),
		ClientCert: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["client_cert"]),
		ClientKey: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["client_key"]),
	}
}

func FlattenKubeConfig(in api.KubeConfig) map[string]interface{} {
	return map[string]interface{}{
		"server": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Server),
		"context": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Context),
		"namespace": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Namespace),
		"kube_bearer_token": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.KubeBearerToken),
		"kube_user": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.KubeUser),
		"kube_password": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.KubePassword),
		"ca_cert": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.CaCert),
		"client_cert": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ClientCert),
		"client_key": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ClientKey),
	}
}
