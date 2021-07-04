package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceGossipConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"protocol":  Computed(Nullable(String())),
			"listen":    Computed(Nullable(String())),
			"secret":    Computed(Nullable(String())),
			"secondary": Computed(Struct(DataSourceGossipConfigSecondary())),
		},
	}
}

func ExpandDataSourceGossipConfig(in map[string]interface{}) kops.GossipConfig {
	if in == nil {
		panic("expand GossipConfig failure, in is nil")
	}
	out := kops.GossipConfig{}
	if in, ok := in["protocol"]; ok && in != nil {
		out.Protocol = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["listen"]; ok && in != nil {
		out.Listen = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["secret"]; ok && in != nil {
		out.Secret = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["secondary"]; ok && in != nil {
		out.Secondary = func(in interface{}) *kops.GossipConfigSecondary {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.GossipConfigSecondary) *kops.GossipConfigSecondary { return &in }(func(in interface{}) kops.GossipConfigSecondary {
					return ExpandDataSourceGossipConfigSecondary(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenDataSourceGossipConfigInto(in kops.GossipConfig, out map[string]interface{}) {
	out["protocol"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Protocol)
	out["listen"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Listen)
	out["secret"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Secret)
	out["secondary"] = func(in *kops.GossipConfigSecondary) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.GossipConfigSecondary) interface{} { return FlattenDataSourceGossipConfigSecondary(in) }(*in)
	}(in.Secondary)
}

func FlattenDataSourceGossipConfig(in kops.GossipConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceGossipConfigInto(in, out)
	return out
}
