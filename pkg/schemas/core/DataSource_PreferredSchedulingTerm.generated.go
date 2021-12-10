package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func DataSourcePreferredSchedulingTerm() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"weight":     ComputedInt(),
			"preference": ComputedStruct(DataSourceNodeSelectorTerm()),
		},
	}

	return res
}

func ExpandDataSourcePreferredSchedulingTerm(in map[string]interface{}) core.PreferredSchedulingTerm {
	if in == nil {
		panic("expand PreferredSchedulingTerm failure, in is nil")
	}
	return core.PreferredSchedulingTerm{
		Weight: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["weight"]),
		Preference: func(in interface{}) core.NodeSelectorTerm {
			return func(in interface{}) core.NodeSelectorTerm {
				if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
					return core.NodeSelectorTerm{}
				}
				return (ExpandDataSourceNodeSelectorTerm(in.([]interface{})[0].(map[string]interface{})))
			}(in)
		}(in["preference"]),
	}
}

func FlattenDataSourcePreferredSchedulingTermInto(in core.PreferredSchedulingTerm, out map[string]interface{}) {
	out["weight"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.Weight)
	out["preference"] = func(in core.NodeSelectorTerm) interface{} {
		return func(in core.NodeSelectorTerm) []interface{} {
			return []interface{}{FlattenDataSourceNodeSelectorTerm(in)}
		}(in)
	}(in.Preference)
}

func FlattenDataSourcePreferredSchedulingTerm(in core.PreferredSchedulingTerm) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourcePreferredSchedulingTermInto(in, out)
	return out
}
