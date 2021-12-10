package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func ResourceNodeAffinity() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"required_during_scheduling_ignored_during_execution":  OptionalStruct(ResourceNodeSelector()),
			"preferred_during_scheduling_ignored_during_execution": OptionalList(ResourcePreferredSchedulingTerm()),
		},
	}

	return res
}

func ExpandResourceNodeAffinity(in map[string]interface{}) core.NodeAffinity {
	if in == nil {
		panic("expand NodeAffinity failure, in is nil")
	}
	return core.NodeAffinity{
		RequiredDuringSchedulingIgnoredDuringExecution: func(in interface{}) *core.NodeSelector {
			return func(in interface{}) *core.NodeSelector {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in core.NodeSelector) *core.NodeSelector {
					return &in
				}(func(in interface{}) core.NodeSelector {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return core.NodeSelector{}
					}
					return (ExpandResourceNodeSelector(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["required_during_scheduling_ignored_during_execution"]),
		PreferredDuringSchedulingIgnoredDuringExecution: func(in interface{}) []core.PreferredSchedulingTerm {
			return func(in interface{}) []core.PreferredSchedulingTerm {
				if in == nil {
					return nil
				}
				var out []core.PreferredSchedulingTerm
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) core.PreferredSchedulingTerm {
						if in == nil {
							return core.PreferredSchedulingTerm{}
						}
						return (ExpandResourcePreferredSchedulingTerm(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["preferred_during_scheduling_ignored_during_execution"]),
	}
}

func FlattenResourceNodeAffinityInto(in core.NodeAffinity, out map[string]interface{}) {
	out["required_during_scheduling_ignored_during_execution"] = func(in *core.NodeSelector) interface{} {
		return func(in *core.NodeSelector) interface{} {
			if in == nil {
				return nil
			}
			return func(in core.NodeSelector) interface{} {
				return func(in core.NodeSelector) []interface{} {
					return []interface{}{FlattenResourceNodeSelector(in)}
				}(in)
			}(*in)
		}(in)
	}(in.RequiredDuringSchedulingIgnoredDuringExecution)
	out["preferred_during_scheduling_ignored_during_execution"] = func(in []core.PreferredSchedulingTerm) interface{} {
		return func(in []core.PreferredSchedulingTerm) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in core.PreferredSchedulingTerm) interface{} {
					return FlattenResourcePreferredSchedulingTerm(in)
				}(in))
			}
			return out
		}(in)
	}(in.PreferredDuringSchedulingIgnoredDuringExecution)
}

func FlattenResourceNodeAffinity(in core.NodeAffinity) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceNodeAffinityInto(in, out)
	return out
}
