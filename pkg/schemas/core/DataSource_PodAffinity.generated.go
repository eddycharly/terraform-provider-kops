package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	core "k8s.io/api/core/v1"
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

func ExpandDataSourcePodAffinity(in map[string]interface{}) core.PodAffinity {
	if in == nil {
		panic("expand PodAffinity failure, in is nil")
	}
	return core.PodAffinity{
		RequiredDuringSchedulingIgnoredDuringExecution: func(in interface{}) []core.PodAffinityTerm {
			return func(in interface{}) []core.PodAffinityTerm {
				if in == nil {
					return nil
				}
				var out []core.PodAffinityTerm
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) core.PodAffinityTerm {
						if in == nil {
							return core.PodAffinityTerm{}
						}
						return ExpandDataSourcePodAffinityTerm(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["required_during_scheduling_ignored_during_execution"]),
		PreferredDuringSchedulingIgnoredDuringExecution: func(in interface{}) []core.WeightedPodAffinityTerm {
			return func(in interface{}) []core.WeightedPodAffinityTerm {
				if in == nil {
					return nil
				}
				var out []core.WeightedPodAffinityTerm
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) core.WeightedPodAffinityTerm {
						if in == nil {
							return core.WeightedPodAffinityTerm{}
						}
						return ExpandDataSourceWeightedPodAffinityTerm(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["preferred_during_scheduling_ignored_during_execution"]),
	}
}

func FlattenDataSourcePodAffinityInto(in core.PodAffinity, out map[string]interface{}) {
	out["required_during_scheduling_ignored_during_execution"] = func(in []core.PodAffinityTerm) interface{} {
		return func(in []core.PodAffinityTerm) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in core.PodAffinityTerm) interface{} {
					return FlattenDataSourcePodAffinityTerm(in)
				}(in))
			}
			return out
		}(in)
	}(in.RequiredDuringSchedulingIgnoredDuringExecution)
	out["preferred_during_scheduling_ignored_during_execution"] = func(in []core.WeightedPodAffinityTerm) interface{} {
		return func(in []core.WeightedPodAffinityTerm) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in core.WeightedPodAffinityTerm) interface{} {
					return FlattenDataSourceWeightedPodAffinityTerm(in)
				}(in))
			}
			return out
		}(in)
	}(in.PreferredDuringSchedulingIgnoredDuringExecution)
}

func FlattenDataSourcePodAffinity(in core.PodAffinity) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourcePodAffinityInto(in, out)
	return out
}
