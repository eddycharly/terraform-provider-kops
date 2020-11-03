package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandAccessSpec(in map[string]interface{}) kops.AccessSpec {
	if in == nil {
		panic("expand AccessSpec failure, in is nil")
	}
	return kops.AccessSpec{
		DNS: func(in interface{}) *kops.DNSAccessSpec {
			value := func(in interface{}) *kops.DNSAccessSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in kops.DNSAccessSpec) *kops.DNSAccessSpec {
					return &in
				}(func(in interface{}) kops.DNSAccessSpec {
					if in.([]interface{})[0] == nil {
						return kops.DNSAccessSpec{}
					}
					return (ExpandDNSAccessSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
			return value
		}(in["dns"]),
		LoadBalancer: func(in interface{}) *kops.LoadBalancerAccessSpec {
			value := func(in interface{}) *kops.LoadBalancerAccessSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in kops.LoadBalancerAccessSpec) *kops.LoadBalancerAccessSpec {
					return &in
				}(func(in interface{}) kops.LoadBalancerAccessSpec {
					if in.([]interface{})[0] == nil {
						return kops.LoadBalancerAccessSpec{}
					}
					return (ExpandLoadBalancerAccessSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
			return value
		}(in["load_balancer"]),
	}
}

func FlattenAccessSpec(in kops.AccessSpec) map[string]interface{} {
	return map[string]interface{}{
		"dns": func(in *kops.DNSAccessSpec) interface{} {
			value := func(in *kops.DNSAccessSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.DNSAccessSpec) interface{} {
					return func(in kops.DNSAccessSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenDNSAccessSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.DNS),
		"load_balancer": func(in *kops.LoadBalancerAccessSpec) interface{} {
			value := func(in *kops.LoadBalancerAccessSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.LoadBalancerAccessSpec) interface{} {
					return func(in kops.LoadBalancerAccessSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenLoadBalancerAccessSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.LoadBalancer),
	}
}
