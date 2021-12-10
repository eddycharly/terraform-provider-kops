package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/api/core/v1"
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

func ExpandResourceWeightedPodAffinityTerm(in map[string]interface{}) v1.WeightedPodAffinityTerm {
	if in == nil {
		panic("expand WeightedPodAffinityTerm failure, in is nil")
	}
	return v1.WeightedPodAffinityTerm{
		Weight: func(in interface{}) int32 /**/ {
			return int32(ExpandInt(in))
		}(in["weight"]),
		PodAffinityTerm: func(in interface{}) v1.PodAffinityTerm /*k8s.io/api/core/v1*/ {
			return func(in interface{}) v1.PodAffinityTerm {
				if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
					return v1.PodAffinityTerm{}
				}
				return (ExpandResourcePodAffinityTerm(in.([]interface{})[0].(map[string]interface{})))
			}(in)
		}(in["pod_affinity_term"]),
	}
}

func FlattenResourceWeightedPodAffinityTermInto(in v1.WeightedPodAffinityTerm, out map[string]interface{}) {
	out["weight"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.Weight)
	out["pod_affinity_term"] = func(in v1.PodAffinityTerm) interface{} {
		return func(in v1.PodAffinityTerm) []interface{} {
			return []interface{}{FlattenResourcePodAffinityTerm(in)}
		}(in)
	}(in.PodAffinityTerm)
}

func FlattenResourceWeightedPodAffinityTerm(in v1.WeightedPodAffinityTerm) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceWeightedPodAffinityTermInto(in, out)
	return out
}
