package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceOpenstackLoadbalancerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"method":              Computed(Nullable(String())),
			"provider":            Computed(Nullable(String())),
			"use_octavia":         Computed(Nullable(Bool())),
			"floating_network":    Computed(Nullable(String())),
			"floating_network_id": Computed(Nullable(String())),
			"floating_subnet":     Computed(Nullable(String())),
			"subnet_id":           Computed(Nullable(String())),
			"manage_sec_groups":   Computed(Nullable(Bool())),
		},
	}
}

func ExpandDataSourceOpenstackLoadbalancerConfig(in map[string]interface{}) kops.OpenstackLoadbalancerConfig {
	if in == nil {
		panic("expand OpenstackLoadbalancerConfig failure, in is nil")
	}
	out := kops.OpenstackLoadbalancerConfig{}
	if in, ok := in["method"]; ok && in != nil {
		out.Method = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["provider"]; ok && in != nil {
		out.Provider = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["use_octavia"]; ok && in != nil {
		out.UseOctavia = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["floating_network"]; ok && in != nil {
		out.FloatingNetwork = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["floating_network_id"]; ok && in != nil {
		out.FloatingNetworkID = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["floating_subnet"]; ok && in != nil {
		out.FloatingSubnet = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["subnet_id"]; ok && in != nil {
		out.SubnetID = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["manage_sec_groups"]; ok && in != nil {
		out.ManageSecGroups = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}

func FlattenDataSourceOpenstackLoadbalancerConfigInto(in kops.OpenstackLoadbalancerConfig, out map[string]interface{}) {
	out["method"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Method)
	out["provider"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Provider)
	out["use_octavia"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.UseOctavia)
	out["floating_network"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.FloatingNetwork)
	out["floating_network_id"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.FloatingNetworkID)
	out["floating_subnet"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.FloatingSubnet)
	out["subnet_id"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.SubnetID)
	out["manage_sec_groups"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.ManageSecGroups)
}

func FlattenDataSourceOpenstackLoadbalancerConfig(in kops.OpenstackLoadbalancerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceOpenstackLoadbalancerConfigInto(in, out)
	return out
}
