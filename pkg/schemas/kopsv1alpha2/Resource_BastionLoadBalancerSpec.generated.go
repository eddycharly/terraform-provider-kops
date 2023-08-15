package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops/v1alpha2"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceBastionLoadBalancerSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_security_groups": OptionalList(String()),
			"type":                       RequiredString(),
		},
	}

	return res
}

func ExpandResourceBastionLoadBalancerSpec(in map[string]interface{}) kopsv1alpha2.BastionLoadBalancerSpec {
	if in == nil {
		panic("expand BastionLoadBalancerSpec failure, in is nil")
	}
	return kopsv1alpha2.BastionLoadBalancerSpec{
		AdditionalSecurityGroups: func(in interface{}) []string {
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
		}(in["additional_security_groups"]),
		Type: func(in interface{}) kopsv1alpha2.LoadBalancerType {
			return v1alpha2.LoadBalancerType(ExpandString(in))
		}(in["type"]),
	}
}

func FlattenResourceBastionLoadBalancerSpecInto(in kopsv1alpha2.BastionLoadBalancerSpec, out map[string]interface{}) {
	out["additional_security_groups"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AdditionalSecurityGroups)
	out["type"] = func(in kopsv1alpha2.LoadBalancerType) interface{} {
		return FlattenString(string(in))
	}(in.Type)
}

func FlattenResourceBastionLoadBalancerSpec(in kopsv1alpha2.BastionLoadBalancerSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceBastionLoadBalancerSpecInto(in, out)
	return out
}
