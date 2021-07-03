package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceCalicoNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"registry":                  Computed(String()),
			"version":                   Computed(String()),
			"aws_src_dst_check":         Computed(String()),
			"bpf_enabled":               Computed(Bool()),
			"bpf_external_service_mode": Computed(String()),
			"bpf_kube_proxy_iptables_cleanup_enabled": Computed(Bool()),
			"bpf_log_level":                      Computed(String()),
			"chain_insert_mode":                  Computed(String()),
			"cpu_request":                        Computed(Ptr(Quantity())),
			"cross_subnet":                       Computed(Bool()),
			"encapsulation_mode":                 Computed(String()),
			"ip_ip_mode":                         Computed(String()),
			"ipv4_auto_detection_method":         Computed(String()),
			"ipv6_auto_detection_method":         Computed(String()),
			"iptables_backend":                   Computed(String()),
			"log_severity_screen":                Computed(String()),
			"mtu":                                Computed(Ptr(Int())),
			"prometheus_metrics_enabled":         Computed(Bool()),
			"prometheus_metrics_port":            Computed(Int()),
			"prometheus_go_metrics_enabled":      Computed(Bool()),
			"prometheus_process_metrics_enabled": Computed(Bool()),
			"major_version":                      Computed(String()),
			"typha_prometheus_metrics_enabled":   Computed(Bool()),
			"typha_prometheus_metrics_port":      Computed(Int()),
			"typha_replicas":                     Computed(Int()),
			"wireguard_enabled":                  Computed(Bool()),
		},
	}
}

func ExpandDataSourceCalicoNetworkingSpec(in map[string]interface{}) kops.CalicoNetworkingSpec {
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
