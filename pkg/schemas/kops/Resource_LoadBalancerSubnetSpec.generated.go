package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceLoadBalancerSubnetSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                 Optional(String()),
			"private_ipv4_address": Optional(Nullable(String())),
			"allocation_id":        Optional(Nullable(String())),
		},
	}
}

func ExpandResourceLoadBalancerSubnetSpec(in map[string]interface{}) kops.LoadBalancerSubnetSpec {
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
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["allocation_id"]; ok && in != nil {
		out.AllocationID = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}

func FlattenResourceLoadBalancerSubnetSpecInto(in kops.LoadBalancerSubnetSpec, out map[string]interface{}) {
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

func FlattenResourceLoadBalancerSubnetSpec(in kops.LoadBalancerSubnetSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceLoadBalancerSubnetSpecInto(in, out)
	return out
}
