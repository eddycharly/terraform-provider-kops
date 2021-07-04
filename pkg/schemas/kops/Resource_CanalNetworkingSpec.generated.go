package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceCanalNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"chain_insert_mode":                  Optional(String()),
			"cpu_request":                        Optional(Nullable(Quantity())),
			"default_endpoint_to_host_action":    Optional(String()),
			"disable_flannel_forward_rules":      Optional(Bool()),
			"disable_tx_checksum_offloading":     Optional(Bool()),
			"iptables_backend":                   Optional(String()),
			"log_severity_sys":                   Optional(String()),
			"mtu":                                Optional(Nullable(Int())),
			"prometheus_go_metrics_enabled":      Optional(Bool()),
			"prometheus_metrics_enabled":         Optional(Bool()),
			"prometheus_metrics_port":            Optional(Int()),
			"prometheus_process_metrics_enabled": Optional(Bool()),
			"typha_prometheus_metrics_enabled":   Optional(Bool()),
			"typha_prometheus_metrics_port":      Optional(Int()),
			"typha_replicas":                     Optional(Int()),
		},
	}
}

func ExpandResourceCanalNetworkingSpec(in map[string]interface{}) kops.CanalNetworkingSpec {
	if in == nil {
		panic("expand CanalNetworkingSpec failure, in is nil")
	}
	out := kops.CanalNetworkingSpec{}
	if in, ok := in["chain_insert_mode"]; ok && in != nil {
		out.ChainInsertMode = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cpu_request"]; ok && in != nil {
		out.CPURequest = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["default_endpoint_to_host_action"]; ok && in != nil {
		out.DefaultEndpointToHostAction = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["disable_flannel_forward_rules"]; ok && in != nil {
		out.DisableFlannelForwardRules = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["disable_tx_checksum_offloading"]; ok && in != nil {
		out.DisableTxChecksumOffloading = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["iptables_backend"]; ok && in != nil {
		out.IptablesBackend = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["log_severity_sys"]; ok && in != nil {
		out.LogSeveritySys = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["mtu"]; ok && in != nil {
		out.MTU = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["prometheus_go_metrics_enabled"]; ok && in != nil {
		out.PrometheusGoMetricsEnabled = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["prometheus_metrics_enabled"]; ok && in != nil {
		out.PrometheusMetricsEnabled = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["prometheus_metrics_port"]; ok && in != nil {
		out.PrometheusMetricsPort = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	if in, ok := in["prometheus_process_metrics_enabled"]; ok && in != nil {
		out.PrometheusProcessMetricsEnabled = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["typha_prometheus_metrics_enabled"]; ok && in != nil {
		out.TyphaPrometheusMetricsEnabled = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["typha_prometheus_metrics_port"]; ok && in != nil {
		out.TyphaPrometheusMetricsPort = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	if in, ok := in["typha_replicas"]; ok && in != nil {
		out.TyphaReplicas = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	return out
}

func FlattenResourceCanalNetworkingSpecInto(in kops.CanalNetworkingSpec, out map[string]interface{}) {
	out["chain_insert_mode"] = func(in string) interface{} { return string(in) }(in.ChainInsertMode)
	out["cpu_request"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.CPURequest)
	out["default_endpoint_to_host_action"] = func(in string) interface{} { return string(in) }(in.DefaultEndpointToHostAction)
	out["disable_flannel_forward_rules"] = func(in bool) interface{} { return in }(in.DisableFlannelForwardRules)
	out["disable_tx_checksum_offloading"] = func(in bool) interface{} { return in }(in.DisableTxChecksumOffloading)
	out["iptables_backend"] = func(in string) interface{} { return string(in) }(in.IptablesBackend)
	out["log_severity_sys"] = func(in string) interface{} { return string(in) }(in.LogSeveritySys)
	out["mtu"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.MTU)
	out["prometheus_go_metrics_enabled"] = func(in bool) interface{} { return in }(in.PrometheusGoMetricsEnabled)
	out["prometheus_metrics_enabled"] = func(in bool) interface{} { return in }(in.PrometheusMetricsEnabled)
	out["prometheus_metrics_port"] = func(in int32) interface{} { return int(in) }(in.PrometheusMetricsPort)
	out["prometheus_process_metrics_enabled"] = func(in bool) interface{} { return in }(in.PrometheusProcessMetricsEnabled)
	out["typha_prometheus_metrics_enabled"] = func(in bool) interface{} { return in }(in.TyphaPrometheusMetricsEnabled)
	out["typha_prometheus_metrics_port"] = func(in int32) interface{} { return int(in) }(in.TyphaPrometheusMetricsPort)
	out["typha_replicas"] = func(in int32) interface{} { return int(in) }(in.TyphaReplicas)
}

func FlattenResourceCanalNetworkingSpec(in kops.CanalNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceCanalNetworkingSpecInto(in, out)
	return out
}
