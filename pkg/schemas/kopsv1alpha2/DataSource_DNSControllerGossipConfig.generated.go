package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceDNSControllerGossipConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"protocol":  ComputedString(),
			"listen":    ComputedString(),
			"secret":    ComputedString(),
			"secondary": ComputedStruct(DataSourceDNSControllerGossipConfigSecondary()),
			"seed":      ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceDNSControllerGossipConfig(in map[string]interface{}) kopsv1alpha2.DNSControllerGossipConfig {
	if in == nil {
		panic("expand DNSControllerGossipConfig failure, in is nil")
	}
	return kopsv1alpha2.DNSControllerGossipConfig{
		Protocol: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["protocol"]),
		Listen: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["listen"]),
		Secret: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["secret"]),
		Secondary: func(in interface{}) *kopsv1alpha2.DNSControllerGossipConfigSecondary {
			return func(in interface{}) *kopsv1alpha2.DNSControllerGossipConfigSecondary {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.DNSControllerGossipConfigSecondary) *kopsv1alpha2.DNSControllerGossipConfigSecondary {
					return &in
				}(func(in interface{}) kopsv1alpha2.DNSControllerGossipConfigSecondary {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceDNSControllerGossipConfigSecondary(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.DNSControllerGossipConfigSecondary{}
				}(in))
			}(in)
		}(in["secondary"]),
		Seed: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["seed"]),
	}
}

func FlattenDataSourceDNSControllerGossipConfigInto(in kopsv1alpha2.DNSControllerGossipConfig, out map[string]interface{}) {
	out["protocol"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Protocol)
	out["listen"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Listen)
	out["secret"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Secret)
	out["secondary"] = func(in *kopsv1alpha2.DNSControllerGossipConfigSecondary) interface{} {
		return func(in *kopsv1alpha2.DNSControllerGossipConfigSecondary) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.DNSControllerGossipConfigSecondary) interface{} {
				return func(in kopsv1alpha2.DNSControllerGossipConfigSecondary) []interface{} {
					return []interface{}{FlattenDataSourceDNSControllerGossipConfigSecondary(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Secondary)
	out["seed"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Seed)
}

func FlattenDataSourceDNSControllerGossipConfig(in kopsv1alpha2.DNSControllerGossipConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceDNSControllerGossipConfigInto(in, out)
	return out
}
