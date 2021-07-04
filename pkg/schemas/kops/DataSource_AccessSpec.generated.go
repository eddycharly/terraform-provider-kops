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
			"dns":           Computed(Struct(DataSourceDNSAccessSpec())),
			"load_balancer": Computed(Struct(DataSourceLoadBalancerAccessSpec())),
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.DNSAccessSpec) *kops.DNSAccessSpec { return &in }(func(in interface{}) kops.DNSAccessSpec {
					return ExpandDataSourceDNSAccessSpec(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	if in, ok := in["load_balancer"]; ok && in != nil {
		out.LoadBalancer = func(in interface{}) *kops.LoadBalancerAccessSpec {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.LoadBalancerAccessSpec) *kops.LoadBalancerAccessSpec { return &in }(func(in interface{}) kops.LoadBalancerAccessSpec {
					return ExpandDataSourceLoadBalancerAccessSpec(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenDataSourceAccessSpecInto(in kops.AccessSpec, out map[string]interface{}) {
	out["dns"] = func(in *kops.DNSAccessSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.DNSAccessSpec) interface{} { return FlattenDataSourceDNSAccessSpec(in) }(*in)
	}(in.DNS)
	out["load_balancer"] = func(in *kops.LoadBalancerAccessSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.LoadBalancerAccessSpec) interface{} { return FlattenDataSourceLoadBalancerAccessSpec(in) }(*in)
	}(in.LoadBalancer)
}

func FlattenDataSourceAccessSpec(in kops.AccessSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAccessSpecInto(in, out)
	return out
}
