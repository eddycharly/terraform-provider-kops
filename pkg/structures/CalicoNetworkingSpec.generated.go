package structures

import (
	"log"
	"reflect"

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
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandQuantity(in))
				return tmp
			}(in)
			log.Printf("%s - %#v", "cpu_request", value)
			return value
		}(in["cpu_request"]),
		CrossSubnet: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			log.Printf("%s - %#v", "cross_subnet", value)
			return value
		}(in["cross_subnet"]),
		LogSeverityScreen: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "log_severity_screen", value)
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
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "mtu", value)
			return value
		}(in["mtu"]),
		PrometheusMetricsEnabled: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			log.Printf("%s - %#v", "prometheus_metrics_enabled", value)
			return value
		}(in["prometheus_metrics_enabled"]),
		PrometheusMetricsPort: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			log.Printf("%s - %#v", "prometheus_metrics_port", value)
			return value
		}(in["prometheus_metrics_port"]),
		PrometheusGoMetricsEnabled: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			log.Printf("%s - %#v", "prometheus_go_metrics_enabled", value)
			return value
		}(in["prometheus_go_metrics_enabled"]),
		PrometheusProcessMetricsEnabled: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			log.Printf("%s - %#v", "prometheus_process_metrics_enabled", value)
			return value
		}(in["prometheus_process_metrics_enabled"]),
		MajorVersion: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "major_version", value)
			return value
		}(in["major_version"]),
		IptablesBackend: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "iptables_backend", value)
			return value
		}(in["iptables_backend"]),
		IPIPMode: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "ipip_mode", value)
			return value
		}(in["ipip_mode"]),
		TyphaPrometheusMetricsEnabled: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			log.Printf("%s - %#v", "typha_prometheus_metrics_enabled", value)
			return value
		}(in["typha_prometheus_metrics_enabled"]),
		TyphaPrometheusMetricsPort: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			log.Printf("%s - %#v", "typha_prometheus_metrics_port", value)
			return value
		}(in["typha_prometheus_metrics_port"]),
		TyphaReplicas: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			log.Printf("%s - %#v", "typha_replicas", value)
			return value
		}(in["typha_replicas"]),
		WireguardEnabled: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			log.Printf("%s - %#v", "wireguard_enabled", value)
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
			log.Printf("%s - %v", "cpu_request", value)
			return value
		}(in.CPURequest),
		"cross_subnet": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			log.Printf("%s - %v", "cross_subnet", value)
			return value
		}(in.CrossSubnet),
		"log_severity_screen": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "log_severity_screen", value)
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
			log.Printf("%s - %v", "mtu", value)
			return value
		}(in.MTU),
		"prometheus_metrics_enabled": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			log.Printf("%s - %v", "prometheus_metrics_enabled", value)
			return value
		}(in.PrometheusMetricsEnabled),
		"prometheus_metrics_port": func(in int32) interface{} {
			value := FlattenInt(int(in))
			log.Printf("%s - %v", "prometheus_metrics_port", value)
			return value
		}(in.PrometheusMetricsPort),
		"prometheus_go_metrics_enabled": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			log.Printf("%s - %v", "prometheus_go_metrics_enabled", value)
			return value
		}(in.PrometheusGoMetricsEnabled),
		"prometheus_process_metrics_enabled": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			log.Printf("%s - %v", "prometheus_process_metrics_enabled", value)
			return value
		}(in.PrometheusProcessMetricsEnabled),
		"major_version": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "major_version", value)
			return value
		}(in.MajorVersion),
		"iptables_backend": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "iptables_backend", value)
			return value
		}(in.IptablesBackend),
		"ipip_mode": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "ipip_mode", value)
			return value
		}(in.IPIPMode),
		"typha_prometheus_metrics_enabled": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			log.Printf("%s - %v", "typha_prometheus_metrics_enabled", value)
			return value
		}(in.TyphaPrometheusMetricsEnabled),
		"typha_prometheus_metrics_port": func(in int32) interface{} {
			value := FlattenInt(int(in))
			log.Printf("%s - %v", "typha_prometheus_metrics_port", value)
			return value
		}(in.TyphaPrometheusMetricsPort),
		"typha_replicas": func(in int32) interface{} {
			value := FlattenInt(int(in))
			log.Printf("%s - %v", "typha_replicas", value)
			return value
		}(in.TyphaReplicas),
		"wireguard_enabled": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			log.Printf("%s - %v", "wireguard_enabled", value)
			return value
		}(in.WireguardEnabled),
	}
}
