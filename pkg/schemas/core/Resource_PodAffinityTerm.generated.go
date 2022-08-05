package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	metaschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/meta"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Schema

func ResourcePodAffinityTerm() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"label_selector":     OptionalStruct(metaschemas.ResourceLabelSelector()),
			"namespaces":         OptionalList(String()),
			"topology_key":       OptionalString(),
			"namespace_selector": OptionalStruct(metaschemas.ResourceLabelSelector()),
		},
	}

	return res
}

func ExpandResourcePodAffinityTerm(in map[string]interface{}) core.PodAffinityTerm {
	if in == nil {
		panic("expand PodAffinityTerm failure, in is nil")
	}
	return core.PodAffinityTerm{
		LabelSelector: func(in interface{}) *meta.LabelSelector {
			return func(in interface{}) *meta.LabelSelector {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in meta.LabelSelector) *meta.LabelSelector {
					return &in
				}(func(in interface{}) meta.LabelSelector {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return metaschemas.ExpandResourceLabelSelector(in[0].(map[string]interface{}))
					}
					return meta.LabelSelector{}
				}(in))
			}(in)
		}(in["label_selector"]),
		Namespaces: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["namespaces"]),
		TopologyKey: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["topology_key"]),
		NamespaceSelector: func(in interface{}) *meta.LabelSelector {
			return func(in interface{}) *meta.LabelSelector {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in meta.LabelSelector) *meta.LabelSelector {
					return &in
				}(func(in interface{}) meta.LabelSelector {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return metaschemas.ExpandResourceLabelSelector(in[0].(map[string]interface{}))
					}
					return meta.LabelSelector{}
				}(in))
			}(in)
		}(in["namespace_selector"]),
	}
}

func FlattenResourcePodAffinityTermInto(in core.PodAffinityTerm, out map[string]interface{}) {
	out["label_selector"] = func(in *meta.LabelSelector) interface{} {
		return func(in *meta.LabelSelector) interface{} {
			if in == nil {
				return nil
			}
			return func(in meta.LabelSelector) interface{} {
				return func(in meta.LabelSelector) []interface{} {
					return []interface{}{metaschemas.FlattenResourceLabelSelector(in)}
				}(in)
			}(*in)
		}(in)
	}(in.LabelSelector)
	out["namespaces"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Namespaces)
	out["topology_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.TopologyKey)
	out["namespace_selector"] = func(in *meta.LabelSelector) interface{} {
		return func(in *meta.LabelSelector) interface{} {
			if in == nil {
				return nil
			}
			return func(in meta.LabelSelector) interface{} {
				return func(in meta.LabelSelector) []interface{} {
					return []interface{}{metaschemas.FlattenResourceLabelSelector(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NamespaceSelector)
}

func FlattenResourcePodAffinityTerm(in core.PodAffinityTerm) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourcePodAffinityTermInto(in, out)
	return out
}
