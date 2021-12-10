package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceCanalNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"chain_insert_mode":                  OptionalString(),
			"cpu_request":                        OptionalQuantity(),
			"default_endpoint_to_host_action":    OptionalString(),
			"disable_flannel_forward_rules":      OptionalBool(),
			"disable_tx_checksum_offloading":     OptionalBool(),
			"iptables_backend":                   OptionalString(),
			"log_severity_sys":                   OptionalString(),
			"mtu":                                OptionalInt(),
			"prometheus_go_metrics_enabled":      OptionalBool(),
			"prometheus_metrics_enabled":         OptionalBool(),
			"prometheus_metrics_port":            OptionalInt(),
			"prometheus_process_metrics_enabled": OptionalBool(),
			"typha_prometheus_metrics_enabled":   OptionalBool(),
			"typha_prometheus_metrics_port":      OptionalInt(),
			"typha_replicas":                     OptionalInt(),
		},
	}

	return res
}

func ExpandResourceCanalNetworkingSpec(in map[string]interface{}) kops.CanalNetworkingSpec {
	if in == nil {
		panic("expand CanalNetworkingSpec failure, in is nil")
	}
	return kops.CanalNetworkingSpec{
		ChainInsertMode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["chain_insert_mode"]),
		CPURequest: func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					return &in
				}(ExpandQuantity(in))
			}(in)
		}(in["cpu_request"]),
		DefaultEndpointToHostAction: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["default_endpoint_to_host_action"]),
		DisableFlannelForwardRules: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disable_flannel_forward_rules"]),
		DisableTxChecksumOffloading: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disable_tx_checksum_offloading"]),
		IptablesBackend: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["iptables_backend"]),
		LogSeveritySys: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["log_severity_sys"]),
		MTU: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["mtu"]),
		PrometheusGoMetricsEnabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["prometheus_go_metrics_enabled"]),
		PrometheusMetricsEnabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["prometheus_metrics_enabled"]),
		PrometheusMetricsPort: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["prometheus_metrics_port"]),
		PrometheusProcessMetricsEnabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["prometheus_process_metrics_enabled"]),
		TyphaPrometheusMetricsEnabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["typha_prometheus_metrics_enabled"]),
		TyphaPrometheusMetricsPort: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["typha_prometheus_metrics_port"]),
		TyphaReplicas: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["typha_replicas"]),
	}
}

func FlattenResourceCanalNetworkingSpecInto(in kops.CanalNetworkingSpec, out map[string]interface{}) {
	out["chain_insert_mode"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ChainInsertMode)
	out["cpu_request"] = func(in *resource.Quantity) interface{} {
		return func(in *resource.Quantity) interface{} {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) interface{} {
				return FlattenQuantity(in)
			}(*in)
		}(in)
	}(in.CPURequest)
	out["default_endpoint_to_host_action"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.DefaultEndpointToHostAction)
	out["disable_flannel_forward_rules"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.DisableFlannelForwardRules)
	out["disable_tx_checksum_offloading"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.DisableTxChecksumOffloading)
	out["iptables_backend"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.IptablesBackend)
	out["log_severity_sys"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.LogSeveritySys)
	out["mtu"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.MTU)
	out["prometheus_go_metrics_enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.PrometheusGoMetricsEnabled)
	out["prometheus_metrics_enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.PrometheusMetricsEnabled)
	out["prometheus_metrics_port"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.PrometheusMetricsPort)
	out["prometheus_process_metrics_enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.PrometheusProcessMetricsEnabled)
	out["typha_prometheus_metrics_enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.TyphaPrometheusMetricsEnabled)
	out["typha_prometheus_metrics_port"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.TyphaPrometheusMetricsPort)
	out["typha_replicas"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.TyphaReplicas)
}

func FlattenResourceCanalNetworkingSpec(in kops.CanalNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceCanalNetworkingSpecInto(in, out)
	return out
}
