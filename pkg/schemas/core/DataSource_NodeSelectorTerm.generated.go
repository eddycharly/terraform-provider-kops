package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/api/core/v1"
)

var _ = Schema

func DataSourceNodeSelectorTerm() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_expressions": ComputedList(DataSourceNodeSelectorRequirement()),
			"match_fields":      ComputedList(DataSourceNodeSelectorRequirement()),
		},
	}

	return res
}

func ExpandDataSourceNodeSelectorTerm(in map[string]interface{}) v1.NodeSelectorTerm {
	if in == nil {
		panic("expand NodeSelectorTerm failure, in is nil")
	}
	return v1.NodeSelectorTerm{
		MatchExpressions: func(in interface{}) []v1.NodeSelectorRequirement /**/ {
			return func(in interface{}) []v1.NodeSelectorRequirement {
				if in == nil {
					return nil
				}
				var out []v1.NodeSelectorRequirement
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) v1.NodeSelectorRequirement {
						if in == nil {
							return v1.NodeSelectorRequirement{}
						}
						return (ExpandDataSourceNodeSelectorRequirement(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["match_expressions"]),
		MatchFields: func(in interface{}) []v1.NodeSelectorRequirement /**/ {
			return func(in interface{}) []v1.NodeSelectorRequirement {
				if in == nil {
					return nil
				}
				var out []v1.NodeSelectorRequirement
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) v1.NodeSelectorRequirement {
						if in == nil {
							return v1.NodeSelectorRequirement{}
						}
						return (ExpandDataSourceNodeSelectorRequirement(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["match_fields"]),
	}
}

func FlattenDataSourceNodeSelectorTermInto(in v1.NodeSelectorTerm, out map[string]interface{}) {
	out["match_expressions"] = func(in []v1.NodeSelectorRequirement) interface{} {
		return func(in []v1.NodeSelectorRequirement) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in v1.NodeSelectorRequirement) interface{} {
					return FlattenDataSourceNodeSelectorRequirement(in)
				}(in))
			}
			return out
		}(in)
	}(in.MatchExpressions)
	out["match_fields"] = func(in []v1.NodeSelectorRequirement) interface{} {
		return func(in []v1.NodeSelectorRequirement) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in v1.NodeSelectorRequirement) interface{} {
					return FlattenDataSourceNodeSelectorRequirement(in)
				}(in))
			}
			return out
		}(in)
	}(in.MatchFields)
}

func FlattenDataSourceNodeSelectorTerm(in v1.NodeSelectorTerm) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceNodeSelectorTermInto(in, out)
	return out
}
