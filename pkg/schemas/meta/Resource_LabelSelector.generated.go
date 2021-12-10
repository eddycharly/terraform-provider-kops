package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

var _ = Schema

func ResourceLabelSelector() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_labels":      OptionalMap(String()),
			"match_expressions": OptionalList(ResourceLabelSelectorRequirement()),
		},
	}

	return res
}

func ExpandResourceLabelSelector(in map[string]interface{}) v1.LabelSelector {
	if in == nil {
		panic("expand LabelSelector failure, in is nil")
	}
	return v1.LabelSelector{
		MatchLabels: func(in interface{}) map[string]string /**/ {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string]string{}
						for key, in := range in {
							out[key] = string(ExpandString(in))
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["match_labels"]),
		MatchExpressions: func(in interface{}) []v1.LabelSelectorRequirement /**/ {
			return func(in interface{}) []v1.LabelSelectorRequirement {
				if in == nil {
					return nil
				}
				var out []v1.LabelSelectorRequirement
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) v1.LabelSelectorRequirement {
						if in == nil {
							return v1.LabelSelectorRequirement{}
						}
						return (ExpandResourceLabelSelectorRequirement(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["match_expressions"]),
	}
}

func FlattenResourceLabelSelectorInto(in v1.LabelSelector, out map[string]interface{}) {
	out["match_labels"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.MatchLabels)
	out["match_expressions"] = func(in []v1.LabelSelectorRequirement) interface{} {
		return func(in []v1.LabelSelectorRequirement) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in v1.LabelSelectorRequirement) interface{} {
					return FlattenResourceLabelSelectorRequirement(in)
				}(in))
			}
			return out
		}(in)
	}(in.MatchExpressions)
}

func FlattenResourceLabelSelector(in v1.LabelSelector) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceLabelSelectorInto(in, out)
	return out
}
