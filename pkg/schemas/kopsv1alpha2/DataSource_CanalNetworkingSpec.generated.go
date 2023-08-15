package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceCanalNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"chain_insert_mode":                  ComputedString(),
			"cpu_request":                        ComputedQuantity(),
			"default_endpoint_to_host_action":    ComputedString(),
			"flanneld_iptables_forward_rules":    ComputedBool(),
			"disable_tx_checksum_offloading":     ComputedBool(),
			"iptables_backend":                   ComputedString(),
			"log_severity_sys":                   ComputedString(),
			"mtu":                                ComputedInt(),
			"prometheus_go_metrics_enabled":      ComputedBool(),
			"prometheus_metrics_enabled":         ComputedBool(),
			"prometheus_metrics_port":            ComputedInt(),
			"prometheus_process_metrics_enabled": ComputedBool(),
			"typha_prometheus_metrics_enabled":   ComputedBool(),
			"typha_prometheus_metrics_port":      ComputedInt(),
			"typha_replicas":                     ComputedInt(),
		},
	}

	return res
}

func ExpandDataSourceCanalNetworkingSpec(in map[string]interface{}) kopsv1alpha2.CanalNetworkingSpec {
	if in == nil {
		panic("expand CanalNetworkingSpec failure, in is nil")
	}
	return kopsv1alpha2.CanalNetworkingSpec{
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
		FlanneldIptablesForwardRules: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["flanneld_iptables_forward_rules"]),
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

func FlattenDataSourceCanalNetworkingSpecInto(in kopsv1alpha2.CanalNetworkingSpec, out map[string]interface{}) {
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
	out["flanneld_iptables_forward_rules"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.FlanneldIptablesForwardRules)
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

func FlattenDataSourceCanalNetworkingSpec(in kopsv1alpha2.CanalNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceCanalNetworkingSpecInto(in, out)
	return out
}
