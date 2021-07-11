package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceCanalNetworkingSpec(t *testing.T) {
	_default := kops.CanalNetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.CanalNetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"chain_insert_mode":                  "",
					"cpu_request":                        nil,
					"default_endpoint_to_host_action":    "",
					"disable_flannel_forward_rules":      false,
					"disable_tx_checksum_offloading":     false,
					"iptables_backend":                   "",
					"log_severity_sys":                   "",
					"mtu":                                nil,
					"prometheus_go_metrics_enabled":      false,
					"prometheus_metrics_enabled":         false,
					"prometheus_metrics_port":            0,
					"prometheus_process_metrics_enabled": false,
					"typha_prometheus_metrics_enabled":   false,
					"typha_prometheus_metrics_port":      0,
					"typha_replicas":                     0,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceCanalNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceCanalNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCanalNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"chain_insert_mode":                  "",
		"cpu_request":                        nil,
		"default_endpoint_to_host_action":    "",
		"disable_flannel_forward_rules":      false,
		"disable_tx_checksum_offloading":     false,
		"iptables_backend":                   "",
		"log_severity_sys":                   "",
		"mtu":                                nil,
		"prometheus_go_metrics_enabled":      false,
		"prometheus_metrics_enabled":         false,
		"prometheus_metrics_port":            0,
		"prometheus_process_metrics_enabled": false,
		"typha_prometheus_metrics_enabled":   false,
		"typha_prometheus_metrics_port":      0,
		"typha_replicas":                     0,
	}
	type args struct {
		in kops.CanalNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CanalNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "ChainInsertMode - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.ChainInsertMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CPURequest - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DefaultEndpointToHostAction - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.DefaultEndpointToHostAction = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableFlannelForwardRules - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.DisableFlannelForwardRules = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableTxChecksumOffloading - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.DisableTxChecksumOffloading = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IptablesBackend - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.IptablesBackend = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogSeveritySys - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.LogSeveritySys = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MTU - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.MTU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusGoMetricsEnabled - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.PrometheusGoMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusMetricsEnabled - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.PrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusMetricsPort - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.PrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusProcessMetricsEnabled - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.PrometheusProcessMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaPrometheusMetricsEnabled - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.TyphaPrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaPrometheusMetricsPort - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.TyphaPrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaReplicas - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.TyphaReplicas = 0
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceCanalNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCanalNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCanalNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"chain_insert_mode":                  "",
		"cpu_request":                        nil,
		"default_endpoint_to_host_action":    "",
		"disable_flannel_forward_rules":      false,
		"disable_tx_checksum_offloading":     false,
		"iptables_backend":                   "",
		"log_severity_sys":                   "",
		"mtu":                                nil,
		"prometheus_go_metrics_enabled":      false,
		"prometheus_metrics_enabled":         false,
		"prometheus_metrics_port":            0,
		"prometheus_process_metrics_enabled": false,
		"typha_prometheus_metrics_enabled":   false,
		"typha_prometheus_metrics_port":      0,
		"typha_replicas":                     0,
	}
	type args struct {
		in kops.CanalNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CanalNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "ChainInsertMode - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.ChainInsertMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CPURequest - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DefaultEndpointToHostAction - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.DefaultEndpointToHostAction = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableFlannelForwardRules - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.DisableFlannelForwardRules = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableTxChecksumOffloading - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.DisableTxChecksumOffloading = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IptablesBackend - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.IptablesBackend = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogSeveritySys - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.LogSeveritySys = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MTU - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.MTU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusGoMetricsEnabled - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.PrometheusGoMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusMetricsEnabled - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.PrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusMetricsPort - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.PrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusProcessMetricsEnabled - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.PrometheusProcessMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaPrometheusMetricsEnabled - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.TyphaPrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaPrometheusMetricsPort - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.TyphaPrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaReplicas - default",
			args: args{
				in: func() kops.CanalNetworkingSpec {
					subject := kops.CanalNetworkingSpec{}
					subject.TyphaReplicas = 0
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceCanalNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCanalNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
