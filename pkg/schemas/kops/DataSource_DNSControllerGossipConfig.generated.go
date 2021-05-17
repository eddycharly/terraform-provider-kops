package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceDNSControllerGossipConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"protocol":  ComputedString(),
			"listen":    ComputedString(),
			"secret":    ComputedString(),
			"secondary": ComputedStruct(DataSourceDNSControllerGossipConfigSecondary()),
			"seed":      ComputedString(),
		},
	}
}

func ExpandDataSourceDNSControllerGossipConfig(in map[string]interface{}) kops.DNSControllerGossipConfig {
	if in == nil {
		panic("expand DNSControllerGossipConfig failure, in is nil")
	}
	return kops.DNSControllerGossipConfig{
		Protocol: func(in interface{}) *string {
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
		Secondary: func(in interface{}) *kops.DNSControllerGossipConfigSecondary {
			return func(in interface{}) *kops.DNSControllerGossipConfigSecondary {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.DNSControllerGossipConfigSecondary) *kops.DNSControllerGossipConfigSecondary {
					return &in
				}(func(in interface{}) kops.DNSControllerGossipConfigSecondary {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.DNSControllerGossipConfigSecondary{}
					}
					return (ExpandDataSourceDNSControllerGossipConfigSecondary(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["secondary"]),
		Seed: func(in interface{}) *string {
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

func FlattenDataSourceDNSControllerGossipConfigInto(in kops.DNSControllerGossipConfig, out map[string]interface{}) {
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
	out["secondary"] = func(in *kops.DNSControllerGossipConfigSecondary) interface{} {
		return func(in *kops.DNSControllerGossipConfigSecondary) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.DNSControllerGossipConfigSecondary) interface{} {
				return func(in kops.DNSControllerGossipConfigSecondary) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceDNSControllerGossipConfigSecondary(in)}
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

func FlattenDataSourceDNSControllerGossipConfig(in kops.DNSControllerGossipConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceDNSControllerGossipConfigInto(in, out)
	return out
}
