package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAccessSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns":           ComputedStruct(DataSourceDNSAccessSpec()),
			"load_balancer": ComputedStruct(DataSourceLoadBalancerAccessSpec()),
		},
	}

	return res
}

func ExpandDataSourceAccessSpec(in map[string]interface{}) kops.AccessSpec {
	if in == nil {
		panic("expand AccessSpec failure, in is nil")
	}
	return kops.AccessSpec{
		DNS: func(in interface{}) *kops.DNSAccessSpec {
			return func(in interface{}) *kops.DNSAccessSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.DNSAccessSpec) *kops.DNSAccessSpec {
					return &in
				}(func(in interface{}) kops.DNSAccessSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.DNSAccessSpec{}
					}
					return (ExpandDataSourceDNSAccessSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["dns"]),
		LoadBalancer: func(in interface{}) *kops.LoadBalancerAccessSpec {
			return func(in interface{}) *kops.LoadBalancerAccessSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.LoadBalancerAccessSpec) *kops.LoadBalancerAccessSpec {
					return &in
				}(func(in interface{}) kops.LoadBalancerAccessSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.LoadBalancerAccessSpec{}
					}
					return (ExpandDataSourceLoadBalancerAccessSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["load_balancer"]),
	}
}

func FlattenDataSourceAccessSpecInto(in kops.AccessSpec, out map[string]interface{}) {
	out["dns"] = func(in *kops.DNSAccessSpec) interface{} {
		return func(in *kops.DNSAccessSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.DNSAccessSpec) interface{} {
				return func(in kops.DNSAccessSpec) []interface{} {
					return []interface{}{FlattenDataSourceDNSAccessSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.DNS)
	out["load_balancer"] = func(in *kops.LoadBalancerAccessSpec) interface{} {
		return func(in *kops.LoadBalancerAccessSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.LoadBalancerAccessSpec) interface{} {
				return func(in kops.LoadBalancerAccessSpec) []interface{} {
					return []interface{}{FlattenDataSourceLoadBalancerAccessSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.LoadBalancer)
}

func FlattenDataSourceAccessSpec(in kops.AccessSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAccessSpecInto(in, out)
	return out
}
