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
			"private_ipv4_address": Computed(Ptr(String())),
			"allocation_id":        Computed(Ptr(String())),
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
