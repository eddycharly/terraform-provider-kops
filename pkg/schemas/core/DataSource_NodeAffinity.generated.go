package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/api/core/v1"
)

var _ = Schema

func DataSourceNodeAffinity() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"required_during_scheduling_ignored_during_execution":  ComputedStruct(DataSourceNodeSelector()),
			"preferred_during_scheduling_ignored_during_execution": ComputedList(DataSourcePreferredSchedulingTerm()),
		},
	}

	return res
}

func ExpandDataSourceNodeAffinity(in map[string]interface{}) v1.NodeAffinity {
	if in == nil {
		panic("expand NodeAffinity failure, in is nil")
	}
	return v1.NodeAffinity{
		RequiredDuringSchedulingIgnoredDuringExecution: func(in interface{}) *v1.NodeSelector /*k8s.io/api/core/v1*/ {
			return func(in interface{}) *v1.NodeSelector {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.NodeSelector) *v1.NodeSelector {
					return &in
				}(func(in interface{}) v1.NodeSelector {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return v1.NodeSelector{}
					}
					return (ExpandDataSourceNodeSelector(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["required_during_scheduling_ignored_during_execution"]),
		PreferredDuringSchedulingIgnoredDuringExecution: func(in interface{}) []v1.PreferredSchedulingTerm /**/ {
			return func(in interface{}) []v1.PreferredSchedulingTerm {
				if in == nil {
					return nil
				}
				var out []v1.PreferredSchedulingTerm
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) v1.PreferredSchedulingTerm {
						if in == nil {
							return v1.PreferredSchedulingTerm{}
						}
						return (ExpandDataSourcePreferredSchedulingTerm(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["preferred_during_scheduling_ignored_during_execution"]),
	}
}

func FlattenDataSourceNodeAffinityInto(in v1.NodeAffinity, out map[string]interface{}) {
	out["required_during_scheduling_ignored_during_execution"] = func(in *v1.NodeSelector) interface{} {
		return func(in *v1.NodeSelector) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.NodeSelector) interface{} {
				return func(in v1.NodeSelector) []interface{} {
					return []interface{}{FlattenDataSourceNodeSelector(in)}
				}(in)
			}(*in)
		}(in)
	}(in.RequiredDuringSchedulingIgnoredDuringExecution)
	out["preferred_during_scheduling_ignored_during_execution"] = func(in []v1.PreferredSchedulingTerm) interface{} {
		return func(in []v1.PreferredSchedulingTerm) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in v1.PreferredSchedulingTerm) interface{} {
					return FlattenDataSourcePreferredSchedulingTerm(in)
				}(in))
			}
			return out
		}(in)
	}(in.PreferredDuringSchedulingIgnoredDuringExecution)
}

func FlattenDataSourceNodeAffinity(in v1.NodeAffinity) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceNodeAffinityInto(in, out)
	return out
}
