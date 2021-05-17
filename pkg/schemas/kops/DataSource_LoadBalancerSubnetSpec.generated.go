package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceLoadBalancerSubnetSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                 ComputedString(),
			"private_ipv4_address": ComputedString(),
			"allocation_id":        ComputedString(),
		},
	}
}

func ExpandDataSourceLoadBalancerSubnetSpec(in map[string]interface{}) kops.LoadBalancerSubnetSpec {
	if in == nil {
		panic("expand LoadBalancerSubnetSpec failure, in is nil")
	}
	return kops.LoadBalancerSubnetSpec{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		PrivateIPv4Address: func(in interface{}) *string {
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
		}(in["private_ipv4_address"]),
		AllocationID: func(in interface{}) *string {
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
		}(in["allocation_id"]),
	}
}

func FlattenDataSourceLoadBalancerSubnetSpecInto(in kops.LoadBalancerSubnetSpec, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["private_ipv4_address"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.PrivateIPv4Address)
	out["allocation_id"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.AllocationID)
}

func FlattenDataSourceLoadBalancerSubnetSpec(in kops.LoadBalancerSubnetSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceLoadBalancerSubnetSpecInto(in, out)
	return out
}
