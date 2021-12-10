package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/api/core/v1"
)

var _ = Schema

func DataSourcePodAffinity() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"required_during_scheduling_ignored_during_execution":  ComputedList(DataSourcePodAffinityTerm()),
			"preferred_during_scheduling_ignored_during_execution": ComputedList(DataSourceWeightedPodAffinityTerm()),
		},
	}

	return res
}

func ExpandDataSourcePodAffinity(in map[string]interface{}) v1.PodAffinity {
	if in == nil {
		panic("expand PodAffinity failure, in is nil")
	}
	return v1.PodAffinity{
		RequiredDuringSchedulingIgnoredDuringExecution: func(in interface{}) []v1.PodAffinityTerm /**/ {
			return func(in interface{}) []v1.PodAffinityTerm {
				if in == nil {
					return nil
				}
				var out []v1.PodAffinityTerm
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) v1.PodAffinityTerm {
						if in == nil {
							return v1.PodAffinityTerm{}
						}
						return (ExpandDataSourcePodAffinityTerm(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["required_during_scheduling_ignored_during_execution"]),
		PreferredDuringSchedulingIgnoredDuringExecution: func(in interface{}) []v1.WeightedPodAffinityTerm /**/ {
			return func(in interface{}) []v1.WeightedPodAffinityTerm {
				if in == nil {
					return nil
				}
				var out []v1.WeightedPodAffinityTerm
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) v1.WeightedPodAffinityTerm {
						if in == nil {
							return v1.WeightedPodAffinityTerm{}
						}
						return (ExpandDataSourceWeightedPodAffinityTerm(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["preferred_during_scheduling_ignored_during_execution"]),
	}
}

func FlattenDataSourcePodAffinityInto(in v1.PodAffinity, out map[string]interface{}) {
	out["required_during_scheduling_ignored_during_execution"] = func(in []v1.PodAffinityTerm) interface{} {
		return func(in []v1.PodAffinityTerm) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in v1.PodAffinityTerm) interface{} {
					return FlattenDataSourcePodAffinityTerm(in)
				}(in))
			}
			return out
		}(in)
	}(in.RequiredDuringSchedulingIgnoredDuringExecution)
	out["preferred_during_scheduling_ignored_during_execution"] = func(in []v1.WeightedPodAffinityTerm) interface{} {
		return func(in []v1.WeightedPodAffinityTerm) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in v1.WeightedPodAffinityTerm) interface{} {
					return FlattenDataSourceWeightedPodAffinityTerm(in)
				}(in))
			}
			return out
		}(in)
	}(in.PreferredDuringSchedulingIgnoredDuringExecution)
}

func FlattenDataSourcePodAffinity(in v1.PodAffinity) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourcePodAffinityInto(in, out)
	return out
}
