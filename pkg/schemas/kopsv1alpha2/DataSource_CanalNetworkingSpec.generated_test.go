package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceCanalNetworkingSpec(t *testing.T) {
	_default := kopsv1alpha2.CanalNetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.CanalNetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"chain_insert_mode":                  "",
					"cpu_request":                        nil,
					"default_endpoint_to_host_action":    "",
					"flanneld_iptables_forward_rules":    nil,
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
			got := ExpandDataSourceCanalNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceCanalNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceCanalNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"chain_insert_mode":                  "",
		"cpu_request":                        nil,
		"default_endpoint_to_host_action":    "",
		"flanneld_iptables_forward_rules":    nil,
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
		in kopsv1alpha2.CanalNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.CanalNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "ChainInsertMode - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.ChainInsertMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DefaultEndpointToHostAction - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.DefaultEndpointToHostAction = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FlanneldIptablesForwardRules - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.FlanneldIptablesForwardRules = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableTxChecksumOffloading - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.DisableTxChecksumOffloading = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IptablesBackend - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.IptablesBackend = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogSeveritySys - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.LogSeveritySys = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MTU - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.MTU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusGoMetricsEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.PrometheusGoMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusMetricsEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.PrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusMetricsPort - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.PrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusProcessMetricsEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.PrometheusProcessMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaPrometheusMetricsEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.TyphaPrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaPrometheusMetricsPort - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.TyphaPrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaReplicas - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
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
			FlattenDataSourceCanalNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceCanalNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceCanalNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"chain_insert_mode":                  "",
		"cpu_request":                        nil,
		"default_endpoint_to_host_action":    "",
		"flanneld_iptables_forward_rules":    nil,
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
		in kopsv1alpha2.CanalNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.CanalNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "ChainInsertMode - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.ChainInsertMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DefaultEndpointToHostAction - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.DefaultEndpointToHostAction = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FlanneldIptablesForwardRules - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.FlanneldIptablesForwardRules = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableTxChecksumOffloading - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.DisableTxChecksumOffloading = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IptablesBackend - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.IptablesBackend = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogSeveritySys - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.LogSeveritySys = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MTU - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.MTU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusGoMetricsEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.PrometheusGoMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusMetricsEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.PrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusMetricsPort - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.PrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusProcessMetricsEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.PrometheusProcessMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaPrometheusMetricsEnabled - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.TyphaPrometheusMetricsEnabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaPrometheusMetricsPort - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.TyphaPrometheusMetricsPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TyphaReplicas - default",
			args: args{
				in: func() kopsv1alpha2.CanalNetworkingSpec {
					subject := kopsv1alpha2.CanalNetworkingSpec{}
					subject.TyphaReplicas = 0
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceCanalNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceCanalNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
