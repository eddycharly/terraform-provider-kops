package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/api/core/v1"
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

func ExpandResourceAffinity(in map[string]interface{}) v1.Affinity {
	if in == nil {
		panic("expand Affinity failure, in is nil")
	}
	return v1.Affinity{
		NodeAffinity: func(in interface{}) *v1.NodeAffinity /*k8s.io/api/core/v1*/ {
			return func(in interface{}) *v1.NodeAffinity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.NodeAffinity) *v1.NodeAffinity {
					return &in
				}(func(in interface{}) v1.NodeAffinity {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return v1.NodeAffinity{}
					}
					return (ExpandResourceNodeAffinity(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["node_affinity"]),
		PodAffinity: func(in interface{}) *v1.PodAffinity /*k8s.io/api/core/v1*/ {
			return func(in interface{}) *v1.PodAffinity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.PodAffinity) *v1.PodAffinity {
					return &in
				}(func(in interface{}) v1.PodAffinity {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return v1.PodAffinity{}
					}
					return (ExpandResourcePodAffinity(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["pod_affinity"]),
		PodAntiAffinity: func(in interface{}) *v1.PodAntiAffinity /*k8s.io/api/core/v1*/ {
			return func(in interface{}) *v1.PodAntiAffinity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.PodAntiAffinity) *v1.PodAntiAffinity {
					return &in
				}(func(in interface{}) v1.PodAntiAffinity {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return v1.PodAntiAffinity{}
					}
					return (ExpandResourcePodAntiAffinity(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["pod_anti_affinity"]),
	}
}

func FlattenResourceAffinityInto(in v1.Affinity, out map[string]interface{}) {
	out["node_affinity"] = func(in *v1.NodeAffinity) interface{} {
		return func(in *v1.NodeAffinity) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.NodeAffinity) interface{} {
				return func(in v1.NodeAffinity) []interface{} {
					return []interface{}{FlattenResourceNodeAffinity(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NodeAffinity)
	out["pod_affinity"] = func(in *v1.PodAffinity) interface{} {
		return func(in *v1.PodAffinity) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.PodAffinity) interface{} {
				return func(in v1.PodAffinity) []interface{} {
					return []interface{}{FlattenResourcePodAffinity(in)}
				}(in)
			}(*in)
		}(in)
	}(in.PodAffinity)
	out["pod_anti_affinity"] = func(in *v1.PodAntiAffinity) interface{} {
		return func(in *v1.PodAntiAffinity) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.PodAntiAffinity) interface{} {
				return func(in v1.PodAntiAffinity) []interface{} {
					return []interface{}{FlattenResourcePodAntiAffinity(in)}
				}(in)
			}(*in)
		}(in)
	}(in.PodAntiAffinity)
}

func FlattenResourceAffinity(in v1.Affinity) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAffinityInto(in, out)
	return out
}
