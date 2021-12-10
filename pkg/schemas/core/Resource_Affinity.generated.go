package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func ResourceAffinity() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_affinity":     OptionalStruct(ResourceNodeAffinity()),
			"pod_affinity":      OptionalStruct(ResourcePodAffinity()),
			"pod_anti_affinity": OptionalStruct(ResourcePodAntiAffinity()),
		},
	}

	return res
}

func ExpandResourceAffinity(in map[string]interface{}) core.Affinity {
	if in == nil {
		panic("expand Affinity failure, in is nil")
	}
	return core.Affinity{
		NodeAffinity: func(in interface{}) *core.NodeAffinity {
			return func(in interface{}) *core.NodeAffinity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in core.NodeAffinity) *core.NodeAffinity {
					return &in
				}(func(in interface{}) core.NodeAffinity {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return core.NodeAffinity{}
					}
					return (ExpandResourceNodeAffinity(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["node_affinity"]),
		PodAffinity: func(in interface{}) *core.PodAffinity {
			return func(in interface{}) *core.PodAffinity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in core.PodAffinity) *core.PodAffinity {
					return &in
				}(func(in interface{}) core.PodAffinity {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return core.PodAffinity{}
					}
					return (ExpandResourcePodAffinity(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["pod_affinity"]),
		PodAntiAffinity: func(in interface{}) *core.PodAntiAffinity {
			return func(in interface{}) *core.PodAntiAffinity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in core.PodAntiAffinity) *core.PodAntiAffinity {
					return &in
				}(func(in interface{}) core.PodAntiAffinity {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return core.PodAntiAffinity{}
					}
					return (ExpandResourcePodAntiAffinity(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["pod_anti_affinity"]),
	}
}

func FlattenResourceAffinityInto(in core.Affinity, out map[string]interface{}) {
	out["node_affinity"] = func(in *core.NodeAffinity) interface{} {
		return func(in *core.NodeAffinity) interface{} {
			if in == nil {
				return nil
			}
			return func(in core.NodeAffinity) interface{} {
				return func(in core.NodeAffinity) []interface{} {
					return []interface{}{FlattenResourceNodeAffinity(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NodeAffinity)
	out["pod_affinity"] = func(in *core.PodAffinity) interface{} {
		return func(in *core.PodAffinity) interface{} {
			if in == nil {
				return nil
			}
			return func(in core.PodAffinity) interface{} {
				return func(in core.PodAffinity) []interface{} {
					return []interface{}{FlattenResourcePodAffinity(in)}
				}(in)
			}(*in)
		}(in)
	}(in.PodAffinity)
	out["pod_anti_affinity"] = func(in *core.PodAntiAffinity) interface{} {
		return func(in *core.PodAntiAffinity) interface{} {
			if in == nil {
				return nil
			}
			return func(in core.PodAntiAffinity) interface{} {
				return func(in core.PodAntiAffinity) []interface{} {
					return []interface{}{FlattenResourcePodAntiAffinity(in)}
				}(in)
			}(*in)
		}(in)
	}(in.PodAntiAffinity)
}

func FlattenResourceAffinity(in core.Affinity) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAffinityInto(in, out)
	return out
}
