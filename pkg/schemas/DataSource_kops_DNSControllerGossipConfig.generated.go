package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsDNSControllerGossipConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"protocol":  ComputedString(),
			"listen":    ComputedString(),
			"secret":    ComputedString(),
			"secondary": ComputedStruct(DataSourceKopsDNSControllerGossipConfig()),
			"seed":      ComputedString(),
		},
	}
}

func ExpandDataSourceKopsDNSControllerGossipConfig(in map[string]interface{}) kops.DNSControllerGossipConfig {
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
		Secondary: func(in interface{}) *kops.DNSControllerGossipConfig {
			return func(in interface{}) *kops.DNSControllerGossipConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.DNSControllerGossipConfig) *kops.DNSControllerGossipConfig {
					return &in
				}(func(in interface{}) kops.DNSControllerGossipConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.DNSControllerGossipConfig{}
					}
					return (ExpandDataSourceKopsDNSControllerGossipConfig(in.([]interface{})[0].(map[string]interface{})))
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

func FlattenDataSourceKopsDNSControllerGossipConfigInto(in kops.DNSControllerGossipConfig, out map[string]interface{}) {
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
	out["secondary"] = func(in *kops.DNSControllerGossipConfig) interface{} {
		return func(in *kops.DNSControllerGossipConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.DNSControllerGossipConfig) interface{} {
				return func(in kops.DNSControllerGossipConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopsDNSControllerGossipConfig(in)}
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

func FlattenDataSourceKopsDNSControllerGossipConfig(in kops.DNSControllerGossipConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsDNSControllerGossipConfigInto(in, out)
	return out
}
