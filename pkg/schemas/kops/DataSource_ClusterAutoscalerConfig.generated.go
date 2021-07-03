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
			"enabled":                          Computed(Ptr(Bool())),
			"expander":                         Computed(Ptr(String())),
			"balance_similar_node_groups":      Computed(Ptr(Bool())),
			"scale_down_utilization_threshold": Computed(Ptr(String())),
			"skip_nodes_with_system_pods":      Computed(Ptr(Bool())),
			"skip_nodes_with_local_storage":    Computed(Ptr(Bool())),
			"new_pod_scale_up_delay":           Computed(Ptr(String())),
			"scale_down_delay_after_add":       Computed(Ptr(String())),
			"image":                            Computed(Ptr(String())),
			"memory_request":                   Computed(Ptr(Quantity())),
			"cpu_request":                      Computed(Ptr(Quantity())),
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
