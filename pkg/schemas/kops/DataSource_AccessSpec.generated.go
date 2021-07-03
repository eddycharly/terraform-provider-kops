package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAccessSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns":           Computed(Ptr(Struct(DataSourceDNSAccessSpec()))),
			"load_balancer": Computed(Ptr(Struct(DataSourceLoadBalancerAccessSpec()))),
		},
	}
}

func ExpandDataSourceAccessSpec(in map[string]interface{}) kops.AccessSpec {
	if in == nil {
		panic("expand AccessSpec failure, in is nil")
	}
	out := kops.AccessSpec{}
	if in, ok := in["dns"]; ok && in != nil {
		out.DNS = func(in interface{}) *kops.DNSAccessSpec {
			if in == nil {
				return nil
			}
			return func(in kops.DNSAccessSpec) *kops.DNSAccessSpec { return &in }(func(in interface{}) kops.DNSAccessSpec {
				if in == nil {
					return kops.DNSAccessSpec{}
				}
				return ExpandDataSourceDNSAccessSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["load_balancer"]; ok && in != nil {
		out.LoadBalancer = func(in interface{}) *kops.LoadBalancerAccessSpec {
			if in == nil {
				return nil
			}
			return func(in kops.LoadBalancerAccessSpec) *kops.LoadBalancerAccessSpec { return &in }(func(in interface{}) kops.LoadBalancerAccessSpec {
				if in == nil {
					return kops.LoadBalancerAccessSpec{}
				}
				return ExpandDataSourceLoadBalancerAccessSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	return out
}
