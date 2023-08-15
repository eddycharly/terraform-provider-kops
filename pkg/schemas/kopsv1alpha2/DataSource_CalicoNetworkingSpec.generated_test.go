package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceCalicoNetworkingSpec(t *testing.T) {
	_default := kopsv1alpha2.CalicoNetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.CalicoNetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"registry":                  "",
					"version":                   "",
					"allow_ip_forwarding":       false,
					"aws_src_dst_check":         "",
					"bpf_enabled":               false,
					"bpf_external_service_mode": "",
					"bpf_kube_proxy_iptables_cleanup_enabled": false,
					"bpf_log_level":                      "",
					"chain_insert_mode":                  "",
					"cpu_request":                        nil,
					"cross_subnet":                       nil,
					"encapsulation_mode":                 "",
					"ip_ip_mode":                         "",
					"ipv4_auto_detection_method":         "",
					"ipv6_auto_detection_method":         "",
					"iptables_backend":                   "",
					"log_severity_screen":                "",
					"mtu":                                nil,
					"prometheus_metrics_enabled":         false,
					"prometheus_metrics_port":            0,
					"prometheus_go_metrics_enabled":      false,
					"prometheus_process_metrics_enabled": false,
					"major_version":                      "",
					"typha_prometheus_metrics_enabled":   false,
					"typha_prometheus_metrics_port":      0,
					"typha_replicas":                     0,
					"vxlan_mode":                         "",
					"wireguard_enabled":                  false,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceCalicoNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceCalicoNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceCalicoNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"registry":                  "",
		"version":                   "",
		"allow_ip_forwarding":       false,
		"aws_src_dst_check":         "",
		"bpf_enabled":               false,
		"bpf_external_service_mode": "",
		"bpf_kube_proxy_iptables_cleanup_enabled": false,
		"bpf_log_level":                      "",
		"chain_insert_mode":                  "",
		"cpu_request":                        nil,
		"cross_subnet":                       nil,
		"encapsulation_mode":                 "",
		"ip_ip_mode":                         "",
		"ipv4_auto_detection_method":         "",
		"ipv6_auto_detection_method":         "",
		"iptables_backend":                   "",
		"log_severity_screen":                "",
		"mtu":                                nil,
		"prometheus_metrics_enabled":         false,
		"prometheus_metrics_port":            0,
		"prometheus_go_metrics_enabled":      false,
		"prometheus_process_metrics_enabled": false,
		"major_version":                      "",
		"typha_prometheus_metrics_enabled":   false,
		"typha_prometheus_metrics_port":      0,
		"typha_replicas":                     0,
		"vxlan_mode":                         "",
		"wireguard_enabled":                  false,
	}
	type args struct {
		in kopsv1alpha2.CalicoNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.CalicoNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "Registry - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.Registry = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.Version = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowIpForwarding - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.AllowIPForwarding = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AwsSrcDstCheck - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.AWSSrcDstCheck = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.BPFEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFExternalServiceMode - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.BPFExternalServiceMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFKubeProxyIptablesCleanupEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.BPFKubeProxyIptablesCleanupEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLogLevel - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.BPFLogLevel = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ChainInsertMode - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.ChainInsertMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CrossSubnet - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.CrossSubnet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EncapsulationMode - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.EncapsulationMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpIpMode - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.IPIPMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv4AutoDetectionMethod - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.IPv4AutoDetectionMethod = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6AutoDetectionMethod - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.IPv6AutoDetectionMethod = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IptablesBackend - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.IptablesBackend = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogSeverityScreen - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.LogSeverityScreen = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MTU - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.MTU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusMetricsEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.PrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusMetricsPort - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.PrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusGoMetricsEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.PrometheusGoMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusProcessMetricsEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.PrometheusProcessMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MajorVersion - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.MajorVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaPrometheusMetricsEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.TyphaPrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaPrometheusMetricsPort - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.TyphaPrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaReplicas - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.TyphaReplicas = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VXLANMode - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.VXLANMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WireguardEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.WireguardEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceCalicoNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceCalicoNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceCalicoNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"registry":                  "",
		"version":                   "",
		"allow_ip_forwarding":       false,
		"aws_src_dst_check":         "",
		"bpf_enabled":               false,
		"bpf_external_service_mode": "",
		"bpf_kube_proxy_iptables_cleanup_enabled": false,
		"bpf_log_level":                      "",
		"chain_insert_mode":                  "",
		"cpu_request":                        nil,
		"cross_subnet":                       nil,
		"encapsulation_mode":                 "",
		"ip_ip_mode":                         "",
		"ipv4_auto_detection_method":         "",
		"ipv6_auto_detection_method":         "",
		"iptables_backend":                   "",
		"log_severity_screen":                "",
		"mtu":                                nil,
		"prometheus_metrics_enabled":         false,
		"prometheus_metrics_port":            0,
		"prometheus_go_metrics_enabled":      false,
		"prometheus_process_metrics_enabled": false,
		"major_version":                      "",
		"typha_prometheus_metrics_enabled":   false,
		"typha_prometheus_metrics_port":      0,
		"typha_replicas":                     0,
		"vxlan_mode":                         "",
		"wireguard_enabled":                  false,
	}
	type args struct {
		in kopsv1alpha2.CalicoNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.CalicoNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "Registry - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.Registry = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.Version = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowIpForwarding - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.AllowIPForwarding = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AwsSrcDstCheck - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.AWSSrcDstCheck = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.BPFEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFExternalServiceMode - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.BPFExternalServiceMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFKubeProxyIptablesCleanupEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.BPFKubeProxyIptablesCleanupEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLogLevel - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.BPFLogLevel = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ChainInsertMode - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.ChainInsertMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CrossSubnet - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.CrossSubnet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EncapsulationMode - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.EncapsulationMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpIpMode - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.IPIPMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv4AutoDetectionMethod - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.IPv4AutoDetectionMethod = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6AutoDetectionMethod - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.IPv6AutoDetectionMethod = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IptablesBackend - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.IptablesBackend = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogSeverityScreen - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.LogSeverityScreen = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MTU - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.MTU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusMetricsEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.PrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusMetricsPort - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.PrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusGoMetricsEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.PrometheusGoMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusProcessMetricsEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.PrometheusProcessMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MajorVersion - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.MajorVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaPrometheusMetricsEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.TyphaPrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaPrometheusMetricsPort - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.TyphaPrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaReplicas - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.TyphaReplicas = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VXLANMode - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.VXLANMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WireguardEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CalicoNetworkingSpec {
					subject := kopsv1alpha2.CalicoNetworkingSpec{}
					subject.WireguardEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceCalicoNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceCalicoNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
