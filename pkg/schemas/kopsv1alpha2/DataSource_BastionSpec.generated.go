package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceBastionSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bastion_public_name":  ComputedString(),
			"idle_timeout_seconds": ComputedInt(),
			"load_balancer":        ComputedStruct(DataSourceBastionLoadBalancerSpec()),
		},
	}

	return res
}

func ExpandDataSourceBastionSpec(in map[string]interface{}) kopsv1alpha2.BastionSpec {
	if in == nil {
		panic("expand BastionSpec failure, in is nil")
	}
	return kopsv1alpha2.BastionSpec{
		PublicName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["bastion_public_name"]),
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
		LoadBalancer: func(in interface{}) *kopsv1alpha2.BastionLoadBalancerSpec {
			return func(in interface{}) *kopsv1alpha2.BastionLoadBalancerSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.BastionLoadBalancerSpec) *kopsv1alpha2.BastionLoadBalancerSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.BastionLoadBalancerSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceBastionLoadBalancerSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.BastionLoadBalancerSpec{}
				}(in))
			}(in)
		}(in["load_balancer"]),
	}
}

func FlattenDataSourceBastionSpecInto(in kopsv1alpha2.BastionSpec, out map[string]interface{}) {
	out["bastion_public_name"] = func(in string) interface{} {
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
	out["load_balancer"] = func(in *kopsv1alpha2.BastionLoadBalancerSpec) interface{} {
		return func(in *kopsv1alpha2.BastionLoadBalancerSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.BastionLoadBalancerSpec) interface{} {
				return func(in kopsv1alpha2.BastionLoadBalancerSpec) []interface{} {
					return []interface{}{FlattenDataSourceBastionLoadBalancerSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.LoadBalancer)
}

func FlattenDataSourceBastionSpec(in kopsv1alpha2.BastionSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceBastionSpecInto(in, out)
	return out
}
