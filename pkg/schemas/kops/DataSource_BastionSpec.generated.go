package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceBastionSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bastion_public_name": ComputedString(),
			"load_balancer":       ComputedStruct(DataSourceBastionLoadBalancerSpec()),
		},
	}

	return res
}

func ExpandDataSourceBastionSpec(in map[string]interface{}) kops.BastionSpec {
	if in == nil {
		panic("expand BastionSpec failure, in is nil")
	}
	return kops.BastionSpec{
		PublicName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["bastion_public_name"]),
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
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceBastionLoadBalancerSpec(in[0].(map[string]interface{}))
					}
					return kops.BastionLoadBalancerSpec{}
				}(in))
			}(in)
		}(in["load_balancer"]),
	}
}

func FlattenDataSourceBastionSpecInto(in kops.BastionSpec, out map[string]interface{}) {
	out["bastion_public_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.PublicName)
	out["load_balancer"] = func(in *kops.BastionLoadBalancerSpec) interface{} {
		return func(in *kops.BastionLoadBalancerSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.BastionLoadBalancerSpec) interface{} {
				return func(in kops.BastionLoadBalancerSpec) []interface{} {
					return []interface{}{FlattenDataSourceBastionLoadBalancerSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.LoadBalancer)
}

func FlattenDataSourceBastionSpec(in kops.BastionSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceBastionSpecInto(in, out)
	return out
}
