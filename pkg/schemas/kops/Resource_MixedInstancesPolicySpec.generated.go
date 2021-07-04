package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceMixedInstancesPolicySpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"instances":                     Optional(List(String())),
			"on_demand_allocation_strategy": Optional(Nullable(String())),
			"on_demand_base":                Required(Nullable(Int())),
			"on_demand_above_base":          Required(Nullable(Int())),
			"spot_allocation_strategy":      Optional(Nullable(String())),
			"spot_instance_pools":           Optional(Nullable(Int())),
		},
	}
}

func ExpandResourceMixedInstancesPolicySpec(in map[string]interface{}) kops.MixedInstancesPolicySpec {
	if in == nil {
		panic("expand MixedInstancesPolicySpec failure, in is nil")
	}
	out := kops.MixedInstancesPolicySpec{}
	if in, ok := in["instances"]; ok && in != nil {
		out.Instances = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["on_demand_allocation_strategy"]; ok && in != nil {
		out.OnDemandAllocationStrategy = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["on_demand_base"]; ok && in != nil {
		out.OnDemandBase = func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int64) *int64 { return &in }(func(in interface{}) int64 { return int64(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["on_demand_above_base"]; ok && in != nil {
		out.OnDemandAboveBase = func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int64) *int64 { return &in }(func(in interface{}) int64 { return int64(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["spot_allocation_strategy"]; ok && in != nil {
		out.SpotAllocationStrategy = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["spot_instance_pools"]; ok && in != nil {
		out.SpotInstancePools = func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int64) *int64 { return &in }(func(in interface{}) int64 { return int64(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenResourceMixedInstancesPolicySpecInto(in kops.MixedInstancesPolicySpec, out map[string]interface{}) {
	out["instances"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.Instances)
	out["on_demand_allocation_strategy"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.OnDemandAllocationStrategy)
	out["on_demand_base"] = func(in *int64) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int64) interface{} { return int(in) }(*in)}
	}(in.OnDemandBase)
	out["on_demand_above_base"] = func(in *int64) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int64) interface{} { return int(in) }(*in)}
	}(in.OnDemandAboveBase)
	out["spot_allocation_strategy"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.SpotAllocationStrategy)
	out["spot_instance_pools"] = func(in *int64) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int64) interface{} { return int(in) }(*in)}
	}(in.SpotInstancePools)
}

func FlattenResourceMixedInstancesPolicySpec(in kops.MixedInstancesPolicySpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceMixedInstancesPolicySpecInto(in, out)
	return out
}
