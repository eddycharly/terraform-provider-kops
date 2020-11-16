package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func KubeConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"server":            ComputedString(),
			"context":           ComputedString(),
			"namespace":         ComputedString(),
			"kube_bearer_token": Sensitive(ComputedString()),
			"kube_user":         Sensitive(ComputedString()),
			"kube_password":     Sensitive(ComputedString()),
			"ca_cert":           Sensitive(ComputedString()),
			"client_cert":       Sensitive(ComputedString()),
			"client_key":        Sensitive(ComputedString()),
		},
	}
}

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

func FlattenKubeConfigInto(in api.KubeConfig, out map[string]interface{}) {
	out["server"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Server)
	out["context"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Context)
	out["namespace"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Namespace)
	out["kube_bearer_token"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.KubeBearerToken)
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

func FlattenKubeConfig(in api.KubeConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenKubeConfigInto(in, out)
	return out
}
