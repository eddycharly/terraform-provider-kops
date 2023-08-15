package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceAPISpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns":             ComputedStruct(DataSourceDNSAccessSpec()),
			"load_balancer":   ComputedStruct(DataSourceLoadBalancerAccessSpec()),
			"public_name":     ComputedString(),
			"additional_sans": ComputedList(String()),
			"access":          ComputedList(String()),
		},
	}

	return res
}

func ExpandDataSourceAPISpec(in map[string]interface{}) kopsv1alpha2.APISpec {
	if in == nil {
		panic("expand APISpec failure, in is nil")
	}
	return kopsv1alpha2.APISpec{
		DNS: func(in interface{}) *kopsv1alpha2.DNSAccessSpec {
			return func(in interface{}) *kopsv1alpha2.DNSAccessSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.DNSAccessSpec) *kopsv1alpha2.DNSAccessSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.DNSAccessSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceDNSAccessSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.DNSAccessSpec{}
				}(in))
			}(in)
		}(in["dns"]),
		LoadBalancer: func(in interface{}) *kopsv1alpha2.LoadBalancerAccessSpec {
			return func(in interface{}) *kopsv1alpha2.LoadBalancerAccessSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.LoadBalancerAccessSpec) *kopsv1alpha2.LoadBalancerAccessSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.LoadBalancerAccessSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceLoadBalancerAccessSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.LoadBalancerAccessSpec{}
				}(in))
			}(in)
		}(in["load_balancer"]),
		PublicName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["public_name"]),
		AdditionalSANs: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["additional_sans"]),
		Access: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["access"]),
	}
}

func FlattenDataSourceAPISpecInto(in kopsv1alpha2.APISpec, out map[string]interface{}) {
	out["dns"] = func(in *kopsv1alpha2.DNSAccessSpec) interface{} {
		return func(in *kopsv1alpha2.DNSAccessSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.DNSAccessSpec) interface{} {
				return func(in kopsv1alpha2.DNSAccessSpec) []interface{} {
					return []interface{}{FlattenDataSourceDNSAccessSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.DNS)
	out["load_balancer"] = func(in *kopsv1alpha2.LoadBalancerAccessSpec) interface{} {
		return func(in *kopsv1alpha2.LoadBalancerAccessSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.LoadBalancerAccessSpec) interface{} {
				return func(in kopsv1alpha2.LoadBalancerAccessSpec) []interface{} {
					return []interface{}{FlattenDataSourceLoadBalancerAccessSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.LoadBalancer)
	out["public_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.PublicName)
	out["additional_sans"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AdditionalSANs)
	out["access"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Access)
}

func FlattenDataSourceAPISpec(in kopsv1alpha2.APISpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAPISpecInto(in, out)
	return out
}
