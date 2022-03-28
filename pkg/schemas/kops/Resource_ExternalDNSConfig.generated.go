package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceExternalDNSConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"watch_ingress":   OptionalBool(),
			"watch_namespace": OptionalString(),
			"provider":        OptionalString(),
		},
	}

	return res
}

func ExpandResourceExternalDNSConfig(in map[string]interface{}) kops.ExternalDNSConfig {
	if in == nil {
		panic("expand ExternalDNSConfig failure, in is nil")
	}
	return kops.ExternalDNSConfig{
		WatchIngress: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
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
		Provider: func(in interface{}) kops.ExternalDNSProvider {
			return kops.ExternalDNSProvider(ExpandString(in))
		}(in["provider"]),
	}
}

func FlattenResourceExternalDNSConfigInto(in kops.ExternalDNSConfig, out map[string]interface{}) {
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
	out["provider"] = func(in kops.ExternalDNSProvider) interface{} {
		return FlattenString(string(in))
	}(in.Provider)
}

func FlattenResourceExternalDNSConfig(in kops.ExternalDNSConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceExternalDNSConfigInto(in, out)
	return out
}
