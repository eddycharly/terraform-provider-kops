package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceCalicoNetworkingSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.CalicoNetworkingSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceCalicoNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceCalicoNetworkingSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceCalicoNetworkingSpecInto(t *testing.T) {
	type args struct {
		in  kops.CalicoNetworkingSpec
		out map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FlattenResourceCalicoNetworkingSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceCalicoNetworkingSpec(t *testing.T) {
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
			want: map[string]interface{}{
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
		{
			name: "Registry - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.Registry = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "Version - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.Version = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "AwsSrcDstCheck - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.AWSSrcDstCheck = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "BPFEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.BPFEnabled = false
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "BPFExternalServiceMode - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.BPFExternalServiceMode = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "BPFKubeProxyIptablesCleanupEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.BPFKubeProxyIptablesCleanupEnabled = false
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "BPFLogLevel - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.BPFLogLevel = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "ChainInsertMode - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.ChainInsertMode = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "CPURequest - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "CrossSubnet - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.CrossSubnet = false
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "EncapsulationMode - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.EncapsulationMode = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "IpIpMode - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.IPIPMode = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "Ipv4AutoDetectionMethod - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.IPv4AutoDetectionMethod = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "Ipv6AutoDetectionMethod - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.IPv6AutoDetectionMethod = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "IptablesBackend - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.IptablesBackend = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "LogSeverityScreen - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.LogSeverityScreen = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "MTU - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.MTU = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "PrometheusMetricsEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.PrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "PrometheusMetricsPort - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.PrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "PrometheusGoMetricsEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.PrometheusGoMetricsEnabled = false
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "PrometheusProcessMetricsEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.PrometheusProcessMetricsEnabled = false
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "MajorVersion - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.MajorVersion = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "TyphaPrometheusMetricsEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.TyphaPrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "TyphaPrometheusMetricsPort - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.TyphaPrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "TyphaReplicas - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.TyphaReplicas = 0
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "WireguardEnabled - default",
			args: args{
				in: func() kops.CalicoNetworkingSpec {
					subject := kops.CalicoNetworkingSpec{}
					subject.WireguardEnabled = false
					return subject
				}(),
			},
			want: map[string]interface{}{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceCalicoNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceCalicoNetworkingSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
