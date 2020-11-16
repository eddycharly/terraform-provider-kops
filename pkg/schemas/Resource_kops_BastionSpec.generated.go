package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsBastionSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bastion_public_name":  RequiredString(),
			"idle_timeout_seconds": OptionalInt(),
			"load_balancer":        OptionalStruct(ResourceKopsBastionLoadBalancerSpec()),
		},
	}
}

func ExpandResourceKopsBastionSpec(in map[string]interface{}) kops.BastionSpec {
	if in == nil {
		panic("expand BastionSpec failure, in is nil")
	}
	return kops.BastionSpec{
		BastionPublicName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["bastion_public_name"]),
		IdleTimeoutSeconds: func(in interface{}) *int64 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					return &in
				}(int64(ExpandInt(in)))
			}(in)
		}(in["idle_timeout_seconds"]),
		LoadBalancer: func(in interface{}) *kops.BastionLoadBalancerSpec {
			return func(in interface{}) *kops.BastionLoadBalancerSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.BastionLoadBalancerSpec) *kops.BastionLoadBalancerSpec {
					return &in
				}(func(in interface{}) kops.BastionLoadBalancerSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.BastionLoadBalancerSpec{}
					}
					return (ExpandResourceKopsBastionLoadBalancerSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["load_balancer"]),
	}
}

func FlattenResourceKopsBastionSpecInto(in kops.BastionSpec, out map[string]interface{}) {
	out["bastion_public_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.BastionPublicName)
	out["idle_timeout_seconds"] = func(in *int64) interface{} {
		return func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.IdleTimeoutSeconds)
	out["load_balancer"] = func(in *kops.BastionLoadBalancerSpec) interface{} {
		return func(in *kops.BastionLoadBalancerSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.BastionLoadBalancerSpec) interface{} {
				return func(in kops.BastionLoadBalancerSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceKopsBastionLoadBalancerSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.LoadBalancer)
}

func FlattenResourceKopsBastionSpec(in kops.BastionSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsBastionSpecInto(in, out)
	return out
}
