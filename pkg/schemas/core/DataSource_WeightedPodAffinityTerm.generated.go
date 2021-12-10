package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func DataSourceWeightedPodAffinityTerm() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"weight":            ComputedInt(),
			"pod_affinity_term": ComputedStruct(DataSourcePodAffinityTerm()),
		},
	}

	return res
}

func ExpandDataSourceWeightedPodAffinityTerm(in map[string]interface{}) core.WeightedPodAffinityTerm {
	if in == nil {
		panic("expand WeightedPodAffinityTerm failure, in is nil")
	}
	return core.WeightedPodAffinityTerm{
		Weight: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["weight"]),
		PodAffinityTerm: func(in interface{}) core.PodAffinityTerm {
			return func(in interface{}) core.PodAffinityTerm {
				if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
					return core.PodAffinityTerm{}
				}
				return (ExpandDataSourcePodAffinityTerm(in.([]interface{})[0].(map[string]interface{})))
			}(in)
		}(in["pod_affinity_term"]),
	}
}

func FlattenDataSourceWeightedPodAffinityTermInto(in core.WeightedPodAffinityTerm, out map[string]interface{}) {
	out["weight"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.Weight)
	out["pod_affinity_term"] = func(in core.PodAffinityTerm) interface{} {
		return func(in core.PodAffinityTerm) []interface{} {
			return []interface{}{FlattenDataSourcePodAffinityTerm(in)}
		}(in)
	}(in.PodAffinityTerm)
}

func FlattenDataSourceWeightedPodAffinityTerm(in core.WeightedPodAffinityTerm) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceWeightedPodAffinityTermInto(in, out)
	return out
}
