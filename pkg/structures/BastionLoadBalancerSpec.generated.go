package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandBastionLoadBalancerSpec(in map[string]interface{}) kops.BastionLoadBalancerSpec {
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

func FlattenBastionLoadBalancerSpec(in kops.BastionLoadBalancerSpec) map[string]interface{} {
	return map[string]interface{}{
		"additional_security_groups": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.AdditionalSecurityGroups),
	}
}
