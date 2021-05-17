package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceGossipConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"protocol":  ComputedString(),
			"listen":    ComputedString(),
			"secret":    ComputedString(),
			"secondary": ComputedStruct(DataSourceGossipConfigSecondary()),
		},
	}
}

func ExpandDataSourceGossipConfig(in map[string]interface{}) kops.GossipConfig {
	if in == nil {
		panic("expand GossipConfig failure, in is nil")
	}
	return kops.GossipConfig{
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
		Secondary: func(in interface{}) *kops.GossipConfigSecondary {
			return func(in interface{}) *kops.GossipConfigSecondary {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.GossipConfigSecondary) *kops.GossipConfigSecondary {
					return &in
				}(func(in interface{}) kops.GossipConfigSecondary {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.GossipConfigSecondary{}
					}
					return (ExpandDataSourceGossipConfigSecondary(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["secondary"]),
	}
}

func FlattenDataSourceGossipConfigInto(in kops.GossipConfig, out map[string]interface{}) {
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
	out["secondary"] = func(in *kops.GossipConfigSecondary) interface{} {
		return func(in *kops.GossipConfigSecondary) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.GossipConfigSecondary) interface{} {
				return func(in kops.GossipConfigSecondary) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceGossipConfigSecondary(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Secondary)
}

func FlattenDataSourceGossipConfig(in kops.GossipConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceGossipConfigInto(in, out)
	return out
}
