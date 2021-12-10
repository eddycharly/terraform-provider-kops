package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	core "k8s.io/api/core/v1"
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

func ExpandDataSourceNodeSelectorTerm(in map[string]interface{}) core.NodeSelectorTerm {
	if in == nil {
		panic("expand NodeSelectorTerm failure, in is nil")
	}
	return core.NodeSelectorTerm{
		MatchExpressions: func(in interface{}) []core.NodeSelectorRequirement {
			return func(in interface{}) []core.NodeSelectorRequirement {
				if in == nil {
					return nil
				}
				var out []core.NodeSelectorRequirement
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) core.NodeSelectorRequirement {
						if in == nil {
							return core.NodeSelectorRequirement{}
						}
						return (ExpandDataSourceNodeSelectorRequirement(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["match_expressions"]),
		MatchFields: func(in interface{}) []core.NodeSelectorRequirement {
			return func(in interface{}) []core.NodeSelectorRequirement {
				if in == nil {
					return nil
				}
				var out []core.NodeSelectorRequirement
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) core.NodeSelectorRequirement {
						if in == nil {
							return core.NodeSelectorRequirement{}
						}
						return (ExpandDataSourceNodeSelectorRequirement(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["match_fields"]),
	}
}

func FlattenDataSourceNodeSelectorTermInto(in core.NodeSelectorTerm, out map[string]interface{}) {
	out["match_expressions"] = func(in []core.NodeSelectorRequirement) interface{} {
		return func(in []core.NodeSelectorRequirement) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in core.NodeSelectorRequirement) interface{} {
					return FlattenDataSourceNodeSelectorRequirement(in)
				}(in))
			}
			return out
		}(in)
	}(in.MatchExpressions)
	out["match_fields"] = func(in []core.NodeSelectorRequirement) interface{} {
		return func(in []core.NodeSelectorRequirement) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in core.NodeSelectorRequirement) interface{} {
					return FlattenDataSourceNodeSelectorRequirement(in)
				}(in))
			}
			return out
		}(in)
	}(in.MatchFields)
}

func FlattenDataSourceNodeSelectorTerm(in core.NodeSelectorTerm) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceNodeSelectorTermInto(in, out)
	return out
}
