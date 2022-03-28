package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceBastionSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"public_name":          RequiredString(),
			"idle_timeout_seconds": OptionalInt(),
			"load_balancer":        OptionalStruct(ResourceBastionLoadBalancerSpec()),
		},
	}

	return res
}

func ExpandResourceBastionSpec(in map[string]interface{}) kops.BastionSpec {
	if in == nil {
		panic("expand BastionSpec failure, in is nil")
	}
	return kops.BastionSpec{
		PublicName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["public_name"]),
		IdleTimeoutSeconds: func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
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
					return (ExpandResourceBastionLoadBalancerSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["load_balancer"]),
	}
}

func FlattenResourceBastionSpecInto(in kops.BastionSpec, out map[string]interface{}) {
	out["public_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.PublicName)
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
				return func(in kops.BastionLoadBalancerSpec) []interface{} {
					return []interface{}{FlattenResourceBastionLoadBalancerSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.LoadBalancer)
}

func FlattenResourceBastionSpec(in kops.BastionSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceBastionSpecInto(in, out)
	return out
}
