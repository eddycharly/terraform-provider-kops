package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceLoadBalancerSubnetSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                 Computed(String()),
			"private_ipv4_address": Computed(Nullable(String())),
			"allocation_id":        Computed(Nullable(String())),
		},
	}
}

func ExpandDataSourceLoadBalancerSubnetSpec(in map[string]interface{}) kops.LoadBalancerSubnetSpec {
	if in == nil {
		panic("expand LoadBalancerSubnetSpec failure, in is nil")
	}
	out := kops.LoadBalancerSubnetSpec{}
	if in, ok := in["name"]; ok && in != nil {
		out.Name = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["private_ipv4_address"]; ok && in != nil {
		out.PrivateIPv4Address = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["allocation_id"]; ok && in != nil {
		out.AllocationID = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenDataSourceLoadBalancerSubnetSpecInto(in kops.LoadBalancerSubnetSpec, out map[string]interface{}) {
	out["name"] = func(in string) interface{} { return string(in) }(in.Name)
	out["private_ipv4_address"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.PrivateIPv4Address)
	out["allocation_id"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.AllocationID)
}

func FlattenDataSourceLoadBalancerSubnetSpec(in kops.LoadBalancerSubnetSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceLoadBalancerSubnetSpecInto(in, out)
	return out
}
