package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceClusterAutoscalerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":                          Computed(Nullable(Bool())),
			"expander":                         Computed(Nullable(String())),
			"balance_similar_node_groups":      Computed(Nullable(Bool())),
			"scale_down_utilization_threshold": Computed(Nullable(String())),
			"skip_nodes_with_system_pods":      Computed(Nullable(Bool())),
			"skip_nodes_with_local_storage":    Computed(Nullable(Bool())),
			"new_pod_scale_up_delay":           Computed(Nullable(String())),
			"scale_down_delay_after_add":       Computed(Nullable(String())),
			"image":                            Computed(Nullable(String())),
			"memory_request":                   Computed(Nullable(Quantity())),
			"cpu_request":                      Computed(Nullable(Quantity())),
		},
	}
}

func ExpandDataSourceClusterAutoscalerConfig(in map[string]interface{}) kops.ClusterAutoscalerConfig {
	if in == nil {
		panic("expand ClusterAutoscalerConfig failure, in is nil")
	}
	out := kops.ClusterAutoscalerConfig{}
	if in, ok := in["enabled"]; ok && in != nil {
		out.Enabled = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["expander"]; ok && in != nil {
		out.Expander = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["balance_similar_node_groups"]; ok && in != nil {
		out.BalanceSimilarNodeGroups = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["scale_down_utilization_threshold"]; ok && in != nil {
		out.ScaleDownUtilizationThreshold = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["skip_nodes_with_system_pods"]; ok && in != nil {
		out.SkipNodesWithSystemPods = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["skip_nodes_with_local_storage"]; ok && in != nil {
		out.SkipNodesWithLocalStorage = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["new_pod_scale_up_delay"]; ok && in != nil {
		out.NewPodScaleUpDelay = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["scale_down_delay_after_add"]; ok && in != nil {
		out.ScaleDownDelayAfterAdd = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["memory_request"]; ok && in != nil {
		out.MemoryRequest = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["cpu_request"]; ok && in != nil {
		out.CPURequest = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}

func FlattenDataSourceClusterAutoscalerConfigInto(in kops.ClusterAutoscalerConfig, out map[string]interface{}) {
	out["enabled"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.Enabled)
	out["expander"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Expander)
	out["balance_similar_node_groups"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.BalanceSimilarNodeGroups)
	out["scale_down_utilization_threshold"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.ScaleDownUtilizationThreshold)
	out["skip_nodes_with_system_pods"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.SkipNodesWithSystemPods)
	out["skip_nodes_with_local_storage"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.SkipNodesWithLocalStorage)
	out["new_pod_scale_up_delay"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.NewPodScaleUpDelay)
	out["scale_down_delay_after_add"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.ScaleDownDelayAfterAdd)
	out["image"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Image)
	out["memory_request"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.MemoryRequest)
	out["cpu_request"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.CPURequest)
}

func FlattenDataSourceClusterAutoscalerConfig(in kops.ClusterAutoscalerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceClusterAutoscalerConfigInto(in, out)
	return out
}
