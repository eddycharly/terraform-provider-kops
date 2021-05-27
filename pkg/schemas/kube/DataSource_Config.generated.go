package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
)

var _ = Schema

func ExpandDataSourceConfig(in map[string]interface{}) kube.Config {
	if in == nil {
		panic("expand Config failure, in is nil")
	}
	return kube.Config{
		Server: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["server"]),
		Context: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["context"]),
		Namespace: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["namespace"]),
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

func FlattenDataSourceConfigInto(in kube.Config, out map[string]interface{}) {
	out["server"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Server)
	out["context"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Context)
	out["namespace"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Namespace)
	out["kube_user"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.KubeUser)
	out["kube_password"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.KubePassword)
	out["ca_cert"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CaCert)
	out["client_cert"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClientCert)
	out["client_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClientKey)
}

func FlattenDataSourceConfig(in kube.Config) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceConfigInto(in, out)
	return out
}
