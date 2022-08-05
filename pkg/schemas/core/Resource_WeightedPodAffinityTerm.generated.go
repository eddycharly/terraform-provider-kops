package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func ResourceWeightedPodAffinityTerm() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"weight":            OptionalInt(),
			"pod_affinity_term": OptionalStruct(ResourcePodAffinityTerm()),
		},
	}

	return res
}

func ExpandResourceWeightedPodAffinityTerm(in map[string]interface{}) core.WeightedPodAffinityTerm {
	if in == nil {
		panic("expand WeightedPodAffinityTerm failure, in is nil")
	}
	return core.WeightedPodAffinityTerm{
		Weight: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["weight"]),
		PodAffinityTerm: func(in interface{}) core.PodAffinityTerm {
			return func(in interface{}) core.PodAffinityTerm {
				if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
					return ExpandResourcePodAffinityTerm(in[0].(map[string]interface{}))
				}
				return core.PodAffinityTerm{}
			}(in)
		}(in["pod_affinity_term"]),
	}
}

func FlattenResourceWeightedPodAffinityTermInto(in core.WeightedPodAffinityTerm, out map[string]interface{}) {
	out["weight"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.Weight)
	out["pod_affinity_term"] = func(in core.PodAffinityTerm) interface{} {
		return func(in core.PodAffinityTerm) []interface{} {
			return []interface{}{FlattenResourcePodAffinityTerm(in)}
		}(in)
	}(in.PodAffinityTerm)
}

func FlattenResourceWeightedPodAffinityTerm(in core.WeightedPodAffinityTerm) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceWeightedPodAffinityTermInto(in, out)
	return out
}
