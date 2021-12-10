package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/api/core/v1"
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

func ExpandResourcePreferredSchedulingTerm(in map[string]interface{}) v1.PreferredSchedulingTerm {
	if in == nil {
		panic("expand PreferredSchedulingTerm failure, in is nil")
	}
	return v1.PreferredSchedulingTerm{
		Weight: func(in interface{}) int32 /**/ {
			return int32(ExpandInt(in))
		}(in["weight"]),
		Preference: func(in interface{}) v1.NodeSelectorTerm /*k8s.io/api/core/v1*/ {
			return func(in interface{}) v1.NodeSelectorTerm {
				if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
					return v1.NodeSelectorTerm{}
				}
				return (ExpandResourceNodeSelectorTerm(in.([]interface{})[0].(map[string]interface{})))
			}(in)
		}(in["preference"]),
	}
}

func FlattenResourcePreferredSchedulingTermInto(in v1.PreferredSchedulingTerm, out map[string]interface{}) {
	out["weight"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.Weight)
	out["preference"] = func(in v1.NodeSelectorTerm) interface{} {
		return func(in v1.NodeSelectorTerm) []interface{} {
			return []interface{}{FlattenResourceNodeSelectorTerm(in)}
		}(in)
	}(in.Preference)
}

func FlattenResourcePreferredSchedulingTerm(in v1.PreferredSchedulingTerm) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourcePreferredSchedulingTermInto(in, out)
	return out
}
