package structures

import (
	"reflect"

	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandCanalNetworkingSpec(in map[string]interface{}) kops.CanalNetworkingSpec {
	if in == nil {
		panic("expand CanalNetworkingSpec failure, in is nil")
	}
	return kops.CanalNetworkingSpec{
		ChainInsertMode: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["chain_insert_mode"]),
		CPURequest: func(in interface{}) *resource.Quantity {
			value := func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandQuantity(in))
			}(in)
			return value
		}(in["cpu_request"]),
		DefaultEndpointToHostAction: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["default_endpoint_to_host_action"]),
		DisableFlannelForwardRules: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["disable_flannel_forward_rules"]),
		DisableTxChecksumOffloading: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["disable_tx_checksum_offloading"]),
		IptablesBackend: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["iptables_backend"]),
		LogSeveritySys: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["log_severity_sys"]),
		MTU: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["mtu"]),
		PrometheusGoMetricsEnabled: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["prometheus_go_metrics_enabled"]),
		PrometheusMetricsEnabled: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["prometheus_metrics_enabled"]),
		PrometheusMetricsPort: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			return value
		}(in["prometheus_metrics_port"]),
		PrometheusProcessMetricsEnabled: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["prometheus_process_metrics_enabled"]),
		TyphaPrometheusMetricsEnabled: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["typha_prometheus_metrics_enabled"]),
		TyphaPrometheusMetricsPort: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			return value
		}(in["typha_prometheus_metrics_port"]),
		TyphaReplicas: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			return value
		}(in["typha_replicas"]),
	}
}

func FlattenCanalNetworkingSpec(in kops.CanalNetworkingSpec) map[string]interface{} {
	return map[string]interface{}{
		"chain_insert_mode": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ChainInsertMode),
		"cpu_request": func(in *resource.Quantity) interface{} {
			value := func(in *resource.Quantity) interface{} {
				if in == nil {
					return nil
				}
				return func(in resource.Quantity) interface{} {
					return FlattenQuantity(in)
				}(*in)
			}(in)
			return value
		}(in.CPURequest),
		"default_endpoint_to_host_action": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.DefaultEndpointToHostAction),
		"disable_flannel_forward_rules": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.DisableFlannelForwardRules),
		"disable_tx_checksum_offloading": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.DisableTxChecksumOffloading),
		"iptables_backend": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.IptablesBackend),
		"log_severity_sys": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.LogSeveritySys),
		"mtu": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.MTU),
		"prometheus_go_metrics_enabled": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.PrometheusGoMetricsEnabled),
		"prometheus_metrics_enabled": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.PrometheusMetricsEnabled),
		"prometheus_metrics_port": func(in int32) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.PrometheusMetricsPort),
		"prometheus_process_metrics_enabled": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.PrometheusProcessMetricsEnabled),
		"typha_prometheus_metrics_enabled": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.TyphaPrometheusMetricsEnabled),
		"typha_prometheus_metrics_port": func(in int32) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.TyphaPrometheusMetricsPort),
		"typha_replicas": func(in int32) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.TyphaReplicas),
	}
}
