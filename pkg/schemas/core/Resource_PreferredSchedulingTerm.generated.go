package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func ResourcePreferredSchedulingTerm() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"weight":     OptionalInt(),
			"preference": OptionalStruct(ResourceNodeSelectorTerm()),
		},
	}

	return res
}

func ExpandResourcePreferredSchedulingTerm(in map[string]interface{}) core.PreferredSchedulingTerm {
	if in == nil {
		panic("expand PreferredSchedulingTerm failure, in is nil")
	}
	return core.PreferredSchedulingTerm{
		Weight: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["weight"]),
		Preference: func(in interface{}) core.NodeSelectorTerm {
			return func(in interface{}) core.NodeSelectorTerm {
				if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
					return ExpandResourceNodeSelectorTerm(in[0].(map[string]interface{}))
				}
				return core.NodeSelectorTerm{}
			}(in)
		}(in["preference"]),
	}
}

func FlattenResourcePreferredSchedulingTermInto(in core.PreferredSchedulingTerm, out map[string]interface{}) {
	out["weight"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.Weight)
	out["preference"] = func(in core.NodeSelectorTerm) interface{} {
		return func(in core.NodeSelectorTerm) []interface{} {
			return []interface{}{FlattenResourceNodeSelectorTerm(in)}
		}(in)
	}(in.Preference)
}

func FlattenResourcePreferredSchedulingTerm(in core.PreferredSchedulingTerm) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourcePreferredSchedulingTermInto(in, out)
	return out
}
