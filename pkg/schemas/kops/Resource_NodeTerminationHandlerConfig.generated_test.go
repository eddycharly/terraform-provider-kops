package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceNodeTerminationHandlerConfig(t *testing.T) {
	_default := kops.NodeTerminationHandlerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.NodeTerminationHandlerConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":                           nil,
					"enable_spot_interruption_draining": nil,
					"enable_scheduled_event_draining":   nil,
					"enable_rebalance_monitoring":       nil,
					"enable_rebalance_draining":         nil,
					"enable_prometheus_metrics":         nil,
					"enable_sqs_termination_draining":   nil,
					"exclude_from_load_balancers":       nil,
					"managed_asg_tag":                   nil,
					"memory_request":                    nil,
					"cpu_request":                       nil,
					"version":                           nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceNodeTerminationHandlerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceNodeTerminationHandlerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceNodeTerminationHandlerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":                           nil,
		"enable_spot_interruption_draining": nil,
		"enable_scheduled_event_draining":   nil,
		"enable_rebalance_monitoring":       nil,
		"enable_rebalance_draining":         nil,
		"enable_prometheus_metrics":         nil,
		"enable_sqs_termination_draining":   nil,
		"exclude_from_load_balancers":       nil,
		"managed_asg_tag":                   nil,
		"memory_request":                    nil,
		"cpu_request":                       nil,
		"version":                           nil,
	}
	type args struct {
		in kops.NodeTerminationHandlerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.NodeTerminationHandlerConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableSpotInterruptionDraining - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.EnableSpotInterruptionDraining = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableScheduledEventDraining - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.EnableScheduledEventDraining = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableRebalanceMonitoring - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.EnableRebalanceMonitoring = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableRebalanceDraining - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.EnableRebalanceDraining = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnablePrometheusMetrics - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.EnablePrometheusMetrics = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableSqsTerminationDraining - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.EnableSQSTerminationDraining = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExcludeFromLoadBalancers - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.ExcludeFromLoadBalancers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ManagedASGTag - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.ManagedASGTag = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceNodeTerminationHandlerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceNodeTerminationHandlerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceNodeTerminationHandlerConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":                           nil,
		"enable_spot_interruption_draining": nil,
		"enable_scheduled_event_draining":   nil,
		"enable_rebalance_monitoring":       nil,
		"enable_rebalance_draining":         nil,
		"enable_prometheus_metrics":         nil,
		"enable_sqs_termination_draining":   nil,
		"exclude_from_load_balancers":       nil,
		"managed_asg_tag":                   nil,
		"memory_request":                    nil,
		"cpu_request":                       nil,
		"version":                           nil,
	}
	type args struct {
		in kops.NodeTerminationHandlerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.NodeTerminationHandlerConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableSpotInterruptionDraining - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.EnableSpotInterruptionDraining = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableScheduledEventDraining - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.EnableScheduledEventDraining = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableRebalanceMonitoring - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.EnableRebalanceMonitoring = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableRebalanceDraining - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.EnableRebalanceDraining = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnablePrometheusMetrics - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.EnablePrometheusMetrics = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableSqsTerminationDraining - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.EnableSQSTerminationDraining = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExcludeFromLoadBalancers - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.ExcludeFromLoadBalancers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ManagedASGTag - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.ManagedASGTag = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceNodeTerminationHandlerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceNodeTerminationHandlerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
