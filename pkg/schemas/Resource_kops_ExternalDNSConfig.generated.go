package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsExternalDNSConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable":         OptionalBool(),
			"watch_ingress":   OptionalBool(),
			"watch_namespace": OptionalString(),
		},
	}
}

func ExpandResourceKopsExternalDNSConfig(in map[string]interface{}) kops.ExternalDNSConfig {
	if in == nil {
		panic("expand ExternalDNSConfig failure, in is nil")
	}
	return kops.ExternalDNSConfig{
		Disable: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disable"]),
		WatchIngress: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["watch_ingress"]),
		WatchNamespace: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["watch_namespace"]),
	}
}

func FlattenResourceKopsExternalDNSConfigInto(in kops.ExternalDNSConfig, out map[string]interface{}) {
	out["disable"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Disable)
	out["watch_ingress"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.WatchIngress)
	out["watch_namespace"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.WatchNamespace)
}

func FlattenResourceKopsExternalDNSConfig(in kops.ExternalDNSConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsExternalDNSConfigInto(in, out)
	return out
}
