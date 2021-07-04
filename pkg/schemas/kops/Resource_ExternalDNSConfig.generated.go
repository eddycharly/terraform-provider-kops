package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceExternalDNSConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable":         Optional(Bool()),
			"watch_ingress":   Optional(Nullable(Bool())),
			"watch_namespace": Optional(String()),
		},
	}
}

func ExpandResourceExternalDNSConfig(in map[string]interface{}) kops.ExternalDNSConfig {
	if in == nil {
		panic("expand ExternalDNSConfig failure, in is nil")
	}
	out := kops.ExternalDNSConfig{}
	if in, ok := in["disable"]; ok && in != nil {
		out.Disable = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["watch_ingress"]; ok && in != nil {
		out.WatchIngress = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["watch_namespace"]; ok && in != nil {
		out.WatchNamespace = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}

func FlattenResourceExternalDNSConfigInto(in kops.ExternalDNSConfig, out map[string]interface{}) {
	out["disable"] = func(in bool) interface{} { return in }(in.Disable)
	out["watch_ingress"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.WatchIngress)
	out["watch_namespace"] = func(in string) interface{} { return string(in) }(in.WatchNamespace)
}

func FlattenResourceExternalDNSConfig(in kops.ExternalDNSConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceExternalDNSConfigInto(in, out)
	return out
}
