package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops/v1alpha2"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceExternalDNSConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable":         ComputedBool(),
			"watch_ingress":   ComputedBool(),
			"watch_namespace": ComputedString(),
			"provider":        ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceExternalDNSConfig(in map[string]interface{}) kopsv1alpha2.ExternalDNSConfig {
	if in == nil {
		panic("expand ExternalDNSConfig failure, in is nil")
	}
	return kopsv1alpha2.ExternalDNSConfig{
		Disable: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disable"]),
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
		Provider: func(in interface{}) kopsv1alpha2.ExternalDNSProvider {
			return v1alpha2.ExternalDNSProvider(ExpandString(in))
		}(in["provider"]),
	}
}

func FlattenDataSourceExternalDNSConfigInto(in kopsv1alpha2.ExternalDNSConfig, out map[string]interface{}) {
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
	out["provider"] = func(in kopsv1alpha2.ExternalDNSProvider) interface{} {
		return FlattenString(string(in))
	}(in.Provider)
}

func FlattenDataSourceExternalDNSConfig(in kopsv1alpha2.ExternalDNSConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceExternalDNSConfigInto(in, out)
	return out
}
