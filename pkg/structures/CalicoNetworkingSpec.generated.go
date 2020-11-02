package structures

import (
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandCalicoNetworkingSpec(in map[string]interface{}) kops.CalicoNetworkingSpec {
	if in == nil {
		panic("expand CalicoNetworkingSpec failure, in is nil")
	}
	return kops.CalicoNetworkingSpec{
		CPURequest: func(in interface{}) *resource.Quantity {
			value := func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in resource.Quantity) *resource.Quantity {
					// if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
					// 	return nil
					// }
					return &in
				}(ExpandQuantity(in))
				return tmp
			}(in)
			return value
		}(in["cpu_request"]),
		CrossSubnet: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["cross_subnet"]),
		LogSeverityScreen: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["log_severity_screen"]),
		MTU: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					// if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
					// 	return nil
					// }
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["mtu"]),
		PrometheusMetricsEnabled: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["prometheus_metrics_enabled"]),
		PrometheusMetricsPort: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			return value
		}(in["prometheus_metrics_port"]),
		PrometheusGoMetricsEnabled: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["prometheus_go_metrics_enabled"]),
		PrometheusProcessMetricsEnabled: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["prometheus_process_metrics_enabled"]),
		MajorVersion: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["major_version"]),
		IptablesBackend: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["iptables_backend"]),
		IPIPMode: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["ip_ip_mode"]),
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
		WireguardEnabled: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["wireguard_enabled"]),
	}
}

func FlattenCalicoNetworkingSpec(in kops.CalicoNetworkingSpec) map[string]interface{} {
	return map[string]interface{}{
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
		"cross_subnet": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.CrossSubnet),
		"log_severity_screen": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.LogSeverityScreen),
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
		"prometheus_metrics_enabled": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.PrometheusMetricsEnabled),
		"prometheus_metrics_port": func(in int32) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.PrometheusMetricsPort),
		"prometheus_go_metrics_enabled": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.PrometheusGoMetricsEnabled),
		"prometheus_process_metrics_enabled": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.PrometheusProcessMetricsEnabled),
		"major_version": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.MajorVersion),
		"iptables_backend": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.IptablesBackend),
		"ip_ip_mode": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.IPIPMode),
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
		"wireguard_enabled": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.WireguardEnabled),
	}
}
