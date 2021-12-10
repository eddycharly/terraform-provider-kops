package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/api/core/v1"
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

func ExpandDataSourceWeightedPodAffinityTerm(in map[string]interface{}) v1.WeightedPodAffinityTerm {
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
				return (ExpandDataSourcePodAffinityTerm(in.([]interface{})[0].(map[string]interface{})))
			}(in)
		}(in["pod_affinity_term"]),
	}
}

func FlattenDataSourceWeightedPodAffinityTermInto(in v1.WeightedPodAffinityTerm, out map[string]interface{}) {
	out["weight"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.Weight)
	out["pod_affinity_term"] = func(in v1.PodAffinityTerm) interface{} {
		return func(in v1.PodAffinityTerm) []interface{} {
			return []interface{}{FlattenDataSourcePodAffinityTerm(in)}
		}(in)
	}(in.PodAffinityTerm)
}

func FlattenDataSourceWeightedPodAffinityTerm(in v1.WeightedPodAffinityTerm) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceWeightedPodAffinityTermInto(in, out)
	return out
}
