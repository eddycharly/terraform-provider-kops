package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceBastionSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bastion_public_name":  Computed(String()),
			"idle_timeout_seconds": Computed(Nullable(Int())),
			"load_balancer":        Computed(Struct(DataSourceBastionLoadBalancerSpec())),
		},
	}
}

func ExpandDataSourceBastionSpec(in map[string]interface{}) kops.BastionSpec {
	if in == nil {
		panic("expand BastionSpec failure, in is nil")
	}
	out := kops.BastionSpec{}
	if in, ok := in["bastion_public_name"]; ok && in != nil {
		out.BastionPublicName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["idle_timeout_seconds"]; ok && in != nil {
		out.IdleTimeoutSeconds = func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			return func(in int64) *int64 { return &in }(func(in interface{}) int64 { return int64(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["load_balancer"]; ok && in != nil {
		out.LoadBalancer = func(in interface{}) *kops.BastionLoadBalancerSpec {
			if in == nil {
				return nil
			}
			return func(in kops.BastionLoadBalancerSpec) *kops.BastionLoadBalancerSpec { return &in }(func(in interface{}) kops.BastionLoadBalancerSpec {
				if in == nil {
					return kops.BastionLoadBalancerSpec{}
				}
				return ExpandDataSourceBastionLoadBalancerSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	return out
}

func FlattenDataSourceBastionSpecInto(in kops.BastionSpec, out map[string]interface{}) {
	out["bastion_public_name"] = func(in string) interface{} { return string(in) }(in.BastionPublicName)
	out["idle_timeout_seconds"] = func(in *int64) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int64) interface{} { return int(in) }(*in)}
	}(in.IdleTimeoutSeconds)
	out["load_balancer"] = func(in *kops.BastionLoadBalancerSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.BastionLoadBalancerSpec) interface{} { return FlattenDataSourceBastionLoadBalancerSpec(in) }(*in)
	}(in.LoadBalancer)
}

func FlattenDataSourceBastionSpec(in kops.BastionSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceBastionSpecInto(in, out)
	return out
}
