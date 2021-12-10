package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	metaschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/meta"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Schema

func DataSourcePodAffinityTerm() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"label_selector":     ComputedStruct(metaschemas.DataSourceLabelSelector()),
			"namespaces":         ComputedList(String()),
			"topology_key":       ComputedString(),
			"namespace_selector": ComputedStruct(metaschemas.DataSourceLabelSelector()),
		},
	}

	return res
}

func ExpandDataSourcePodAffinityTerm(in map[string]interface{}) core.PodAffinityTerm {
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return meta.LabelSelector{}
					}
					return (metaschemas.ExpandDataSourceLabelSelector(in.([]interface{})[0].(map[string]interface{})))
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return meta.LabelSelector{}
					}
					return (metaschemas.ExpandDataSourceLabelSelector(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["namespace_selector"]),
	}
}

func FlattenDataSourcePodAffinityTermInto(in core.PodAffinityTerm, out map[string]interface{}) {
	out["label_selector"] = func(in *meta.LabelSelector) interface{} {
		return func(in *meta.LabelSelector) interface{} {
			if in == nil {
				return nil
			}
			return func(in meta.LabelSelector) interface{} {
				return func(in meta.LabelSelector) []interface{} {
					return []interface{}{metaschemas.FlattenDataSourceLabelSelector(in)}
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
					return []interface{}{metaschemas.FlattenDataSourceLabelSelector(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NamespaceSelector)
}

func FlattenDataSourcePodAffinityTerm(in core.PodAffinityTerm) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourcePodAffinityTermInto(in, out)
	return out
}
