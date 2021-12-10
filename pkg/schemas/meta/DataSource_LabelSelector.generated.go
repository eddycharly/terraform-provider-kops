package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Schema

func DataSourceLabelSelector() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_labels":      ComputedMap(String()),
			"match_expressions": ComputedList(DataSourceLabelSelectorRequirement()),
		},
	}

	return res
}

func ExpandDataSourceLabelSelector(in map[string]interface{}) meta.LabelSelector {
	if in == nil {
		panic("expand LabelSelector failure, in is nil")
	}
	return meta.LabelSelector{
		MatchLabels: func(in interface{}) map[string]string {
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
		MatchExpressions: func(in interface{}) []meta.LabelSelectorRequirement {
			return func(in interface{}) []meta.LabelSelectorRequirement {
				if in == nil {
					return nil
				}
				var out []meta.LabelSelectorRequirement
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) meta.LabelSelectorRequirement {
						if in == nil {
							return meta.LabelSelectorRequirement{}
						}
						return (ExpandDataSourceLabelSelectorRequirement(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["match_expressions"]),
	}
}

func FlattenDataSourceLabelSelectorInto(in meta.LabelSelector, out map[string]interface{}) {
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
	out["match_expressions"] = func(in []meta.LabelSelectorRequirement) interface{} {
		return func(in []meta.LabelSelectorRequirement) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in meta.LabelSelectorRequirement) interface{} {
					return FlattenDataSourceLabelSelectorRequirement(in)
				}(in))
			}
			return out
		}(in)
	}(in.MatchExpressions)
}

func FlattenDataSourceLabelSelector(in meta.LabelSelector) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceLabelSelectorInto(in, out)
	return out
}
