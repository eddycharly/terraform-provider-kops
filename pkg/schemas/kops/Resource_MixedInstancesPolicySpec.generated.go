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
			"on_demand_allocation_strategy": Optional(Ptr(String())),
			"on_demand_base":                Required(Ptr(Int())),
			"on_demand_above_base":          Required(Ptr(Int())),
			"spot_allocation_strategy":      Optional(Ptr(String())),
			"spot_instance_pools":           Optional(Ptr(Int())),
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
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["on_demand_base"]; ok && in != nil {
		out.OnDemandBase = func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			return func(in int64) *int64 { return &in }(func(in interface{}) int64 { return int64(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["on_demand_above_base"]; ok && in != nil {
		out.OnDemandAboveBase = func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			return func(in int64) *int64 { return &in }(func(in interface{}) int64 { return int64(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["spot_allocation_strategy"]; ok && in != nil {
		out.SpotAllocationStrategy = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["spot_instance_pools"]; ok && in != nil {
		out.SpotInstancePools = func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			return func(in int64) *int64 { return &in }(func(in interface{}) int64 { return int64(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
