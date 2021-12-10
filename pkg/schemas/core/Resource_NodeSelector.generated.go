package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	core "k8s.io/api/core/v1"
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

func ExpandResourceNodeSelector(in map[string]interface{}) core.NodeSelector {
	if in == nil {
		panic("expand NodeSelector failure, in is nil")
	}
	return core.NodeSelector{
		NodeSelectorTerms: func(in interface{}) []core.NodeSelectorTerm {
			return func(in interface{}) []core.NodeSelectorTerm {
				if in == nil {
					return nil
				}
				var out []core.NodeSelectorTerm
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) core.NodeSelectorTerm {
						if in == nil {
							return core.NodeSelectorTerm{}
						}
						return (ExpandResourceNodeSelectorTerm(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["node_selector_terms"]),
	}
}

func FlattenResourceNodeSelectorInto(in core.NodeSelector, out map[string]interface{}) {
	out["node_selector_terms"] = func(in []core.NodeSelectorTerm) interface{} {
		return func(in []core.NodeSelectorTerm) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in core.NodeSelectorTerm) interface{} {
					return FlattenResourceNodeSelectorTerm(in)
				}(in))
			}
			return out
		}(in)
	}(in.NodeSelectorTerms)
}

func FlattenResourceNodeSelector(in core.NodeSelector) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceNodeSelectorInto(in, out)
	return out
}
