package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceCalicoNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"registry":                  Optional(String()),
			"version":                   Optional(String()),
			"aws_src_dst_check":         Optional(String()),
			"bpf_enabled":               Optional(Bool()),
			"bpf_external_service_mode": Optional(String()),
			"bpf_kube_proxy_iptables_cleanup_enabled": Optional(Bool()),
			"bpf_log_level":                      Optional(String()),
			"chain_insert_mode":                  Optional(String()),
			"cpu_request":                        Optional(Ptr(Quantity())),
			"cross_subnet":                       Optional(Bool()),
			"encapsulation_mode":                 Optional(String()),
			"ip_ip_mode":                         Optional(String()),
			"ipv4_auto_detection_method":         Optional(String()),
			"ipv6_auto_detection_method":         Optional(String()),
			"iptables_backend":                   Optional(String()),
			"log_severity_screen":                Optional(String()),
			"mtu":                                Optional(Ptr(Int())),
			"prometheus_metrics_enabled":         Optional(Bool()),
			"prometheus_metrics_port":            Optional(Int()),
			"prometheus_go_metrics_enabled":      Optional(Bool()),
			"prometheus_process_metrics_enabled": Optional(Bool()),
			"major_version":                      Optional(String()),
			"typha_prometheus_metrics_enabled":   Optional(Bool()),
			"typha_prometheus_metrics_port":      Optional(Int()),
			"typha_replicas":                     Optional(Int()),
			"wireguard_enabled":                  Optional(Bool()),
		},
	}
}

func ExpandResourceCalicoNetworkingSpec(in map[string]interface{}) kops.CalicoNetworkingSpec {
	if in == nil {
		panic("expand CalicoNetworkingSpec failure, in is nil")
	}
	out := kops.CalicoNetworkingSpec{}
	if in, ok := in["registry"]; ok && in != nil {
		out.Registry = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["version"]; ok && in != nil {
		out.Version = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["aws_src_dst_check"]; ok && in != nil {
		out.AWSSrcDstCheck = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["bpf_enabled"]; ok && in != nil {
		out.BPFEnabled = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["bpf_external_service_mode"]; ok && in != nil {
		out.BPFExternalServiceMode = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["bpf_kube_proxy_iptables_cleanup_enabled"]; ok && in != nil {
		out.BPFKubeProxyIptablesCleanupEnabled = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["bpf_log_level"]; ok && in != nil {
		out.BPFLogLevel = func(in interface{}) string { return string(in.(string)) }(in)
	}
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
	if in, ok := in["cross_subnet"]; ok && in != nil {
		out.CrossSubnet = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["encapsulation_mode"]; ok && in != nil {
		out.EncapsulationMode = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["ip_ip_mode"]; ok && in != nil {
		out.IPIPMode = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["ipv4_auto_detection_method"]; ok && in != nil {
		out.IPv4AutoDetectionMethod = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["ipv6_auto_detection_method"]; ok && in != nil {
		out.IPv6AutoDetectionMethod = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["iptables_backend"]; ok && in != nil {
		out.IptablesBackend = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["log_severity_screen"]; ok && in != nil {
		out.LogSeverityScreen = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["mtu"]; ok && in != nil {
		out.MTU = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["prometheus_metrics_enabled"]; ok && in != nil {
		out.PrometheusMetricsEnabled = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["prometheus_metrics_port"]; ok && in != nil {
		out.PrometheusMetricsPort = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	if in, ok := in["prometheus_go_metrics_enabled"]; ok && in != nil {
		out.PrometheusGoMetricsEnabled = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["prometheus_process_metrics_enabled"]; ok && in != nil {
		out.PrometheusProcessMetricsEnabled = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["major_version"]; ok && in != nil {
		out.MajorVersion = func(in interface{}) string { return string(in.(string)) }(in)
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
	if in, ok := in["wireguard_enabled"]; ok && in != nil {
		out.WireguardEnabled = func(in interface{}) bool { return in.(bool) }(in)
	}
	return out
}
