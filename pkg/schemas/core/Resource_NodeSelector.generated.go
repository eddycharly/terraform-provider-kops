package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/api/core/v1"
)

var _ = Schema

func ResourceNodeSelector() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_selector_terms": OptionalList(ResourceNodeSelectorTerm()),
		},
	}

	return res
}

func ExpandResourceNodeSelector(in map[string]interface{}) v1.NodeSelector {
	if in == nil {
		panic("expand NodeSelector failure, in is nil")
	}
	return v1.NodeSelector{
		NodeSelectorTerms: func(in interface{}) []v1.NodeSelectorTerm /**/ {
			return func(in interface{}) []v1.NodeSelectorTerm {
				if in == nil {
					return nil
				}
				var out []v1.NodeSelectorTerm
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) v1.NodeSelectorTerm {
						if in == nil {
							return v1.NodeSelectorTerm{}
						}
						return (ExpandResourceNodeSelectorTerm(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["node_selector_terms"]),
	}
}

func FlattenResourceNodeSelectorInto(in v1.NodeSelector, out map[string]interface{}) {
	out["node_selector_terms"] = func(in []v1.NodeSelectorTerm) interface{} {
		return func(in []v1.NodeSelectorTerm) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in v1.NodeSelectorTerm) interface{} {
					return FlattenResourceNodeSelectorTerm(in)
				}(in))
			}
			return out
		}(in)
	}(in.NodeSelectorTerms)
}

func FlattenResourceNodeSelector(in v1.NodeSelector) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceNodeSelectorInto(in, out)
	return out
}
