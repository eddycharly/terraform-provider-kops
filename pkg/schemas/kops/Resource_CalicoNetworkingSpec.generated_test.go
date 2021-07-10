package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceCalicoNetworkingSpec(t *testing.T) {
	_default := kops.CalicoNetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.CalicoNetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"registry":                  "",
					"version":                   "",
					"aws_src_dst_check":         "",
					"bpf_enabled":               false,
					"bpf_external_service_mode": "",
					"bpf_kube_proxy_iptables_cleanup_enabled": false,
					"bpf_log_level":                      "",
					"chain_insert_mode":                  "",
					"cpu_request":                        nil,
					"cross_subnet":                       false,
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
					"wireguard_enabled":                  false,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceCalicoNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceCalicoNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCalicoNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"registry":                  "",
		"version":                   "",
		"aws_src_dst_check":         "",
		"bpf_enabled":               false,
		"bpf_external_service_mode": "",
		"bpf_kube_proxy_iptables_cleanup_enabled": false,
		"bpf_log_level":                      "",
		"chain_insert_mode":                  "",
		"cpu_request":                        nil,
		"cross_subnet":                       false,
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
		"wireguard_enabled":                  false,
	}
	type args struct {
		in kops.CalicoNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CalicoNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "Registry - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.Registry = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.Version = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AwsSrcDstCheck - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.AWSSrcDstCheck = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.BPFEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFExternalServiceMode - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.BPFExternalServiceMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFKubeProxyIptablesCleanupEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.BPFKubeProxyIptablesCleanupEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLogLevel - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.BPFLogLevel = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ChainInsertMode - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.ChainInsertMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CPURequest - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CrossSubnet - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.CrossSubnet = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EncapsulationMode - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.EncapsulationMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpIpMode - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.IPIPMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv4AutoDetectionMethod - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.IPv4AutoDetectionMethod = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6AutoDetectionMethod - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.IPv6AutoDetectionMethod = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IptablesBackend - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.IptablesBackend = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogSeverityScreen - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.LogSeverityScreen = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MTU - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.MTU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusMetricsEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.PrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusMetricsPort - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.PrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusGoMetricsEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.PrometheusGoMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusProcessMetricsEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.PrometheusProcessMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MajorVersion - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.MajorVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaPrometheusMetricsEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.TyphaPrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaPrometheusMetricsPort - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.TyphaPrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaReplicas - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.TyphaReplicas = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WireguardEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
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
			FlattenResourceCalicoNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCalicoNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCalicoNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"registry":                  "",
		"version":                   "",
		"aws_src_dst_check":         "",
		"bpf_enabled":               false,
		"bpf_external_service_mode": "",
		"bpf_kube_proxy_iptables_cleanup_enabled": false,
		"bpf_log_level":                      "",
		"chain_insert_mode":                  "",
		"cpu_request":                        nil,
		"cross_subnet":                       false,
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
		"wireguard_enabled":                  false,
	}
	type args struct {
		in kops.CalicoNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CalicoNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "Registry - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.Registry = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.Version = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AwsSrcDstCheck - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.AWSSrcDstCheck = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.BPFEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFExternalServiceMode - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.BPFExternalServiceMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFKubeProxyIptablesCleanupEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.BPFKubeProxyIptablesCleanupEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLogLevel - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.BPFLogLevel = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ChainInsertMode - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.ChainInsertMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CPURequest - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CrossSubnet - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.CrossSubnet = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EncapsulationMode - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.EncapsulationMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpIpMode - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.IPIPMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv4AutoDetectionMethod - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.IPv4AutoDetectionMethod = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6AutoDetectionMethod - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.IPv6AutoDetectionMethod = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IptablesBackend - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.IptablesBackend = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogSeverityScreen - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.LogSeverityScreen = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MTU - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.MTU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusMetricsEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.PrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusMetricsPort - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.PrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusGoMetricsEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.PrometheusGoMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusProcessMetricsEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.PrometheusProcessMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MajorVersion - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.MajorVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaPrometheusMetricsEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.TyphaPrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaPrometheusMetricsPort - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.TyphaPrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaReplicas - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.TyphaReplicas = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WireguardEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.WireguardEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceCalicoNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCalicoNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
