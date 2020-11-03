package structures

import (
	"reflect"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandBastionSpec(in map[string]interface{}) kops.BastionSpec {
	if in == nil {
		panic("expand BastionSpec failure, in is nil")
	}
	return kops.BastionSpec{
		BastionPublicName: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["bastion_public_name"]),
		IdleTimeoutSeconds: func(in interface{}) *int64 {
			value := func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int64(ExpandInt(in)))
			}(in)
			return value
		}(in["idle_timeout_seconds"]),
		LoadBalancer: func(in interface{}) *kops.BastionLoadBalancerSpec {
			value := func(in interface{}) *kops.BastionLoadBalancerSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in kops.BastionLoadBalancerSpec) *kops.BastionLoadBalancerSpec {
					return &in
				}(func(in interface{}) kops.BastionLoadBalancerSpec {
					if in.([]interface{})[0] == nil {
						return kops.BastionLoadBalancerSpec{}
					}
					return (ExpandBastionLoadBalancerSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
			return value
		}(in["load_balancer"]),
	}
}

func FlattenBastionSpec(in kops.BastionSpec) map[string]interface{} {
	return map[string]interface{}{
		"bastion_public_name": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.BastionPublicName),
		"idle_timeout_seconds": func(in *int64) interface{} {
			value := func(in *int64) interface{} {
				if in == nil {
					return nil
				}
				return func(in int64) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.IdleTimeoutSeconds),
		"load_balancer": func(in *kops.BastionLoadBalancerSpec) interface{} {
			value := func(in *kops.BastionLoadBalancerSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.BastionLoadBalancerSpec) interface{} {
					return func(in kops.BastionLoadBalancerSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenBastionLoadBalancerSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.LoadBalancer),
	}
}
