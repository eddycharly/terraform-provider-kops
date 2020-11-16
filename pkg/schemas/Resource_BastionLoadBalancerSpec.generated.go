package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceBastionLoadBalancerSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_security_groups": RequiredList(String()),
		},
	}
}

func ExpandResourceBastionLoadBalancerSpec(in map[string]interface{}) kops.BastionLoadBalancerSpec {
	if in == nil {
		panic("expand BastionLoadBalancerSpec failure, in is nil")
	}
	return kops.BastionLoadBalancerSpec{
		AdditionalSecurityGroups: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["additional_security_groups"]),
	}
}

func FlattenResourceBastionLoadBalancerSpecInto(in kops.BastionLoadBalancerSpec, out map[string]interface{}) {
	out["additional_security_groups"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AdditionalSecurityGroups)
}

func FlattenResourceBastionLoadBalancerSpec(in kops.BastionLoadBalancerSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceBastionLoadBalancerSpecInto(in, out)
	return out
}
