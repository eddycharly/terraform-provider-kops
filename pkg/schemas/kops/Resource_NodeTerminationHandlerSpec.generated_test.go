package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceNodeTerminationHandlerSpec(t *testing.T) {
	_default := kops.NodeTerminationHandlerSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.NodeTerminationHandlerSpec
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
			got := ExpandResourceNodeTerminationHandlerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceNodeTerminationHandlerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceNodeTerminationHandlerSpecInto(t *testing.T) {
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
		in kops.NodeTerminationHandlerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.NodeTerminationHandlerSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableSpotInterruptionDraining - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.EnableSpotInterruptionDraining = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableScheduledEventDraining - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.EnableScheduledEventDraining = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableRebalanceMonitoring - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.EnableRebalanceMonitoring = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableRebalanceDraining - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.EnableRebalanceDraining = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnablePrometheusMetrics - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.EnablePrometheusMetrics = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableSqsTerminationDraining - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.EnableSQSTerminationDraining = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExcludeFromLoadBalancers - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.ExcludeFromLoadBalancers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ManagedASGTag - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.ManagedASGTag = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
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
			FlattenResourceNodeTerminationHandlerSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceNodeTerminationHandlerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceNodeTerminationHandlerSpec(t *testing.T) {
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
		in kops.NodeTerminationHandlerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.NodeTerminationHandlerSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableSpotInterruptionDraining - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.EnableSpotInterruptionDraining = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableScheduledEventDraining - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.EnableScheduledEventDraining = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableRebalanceMonitoring - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.EnableRebalanceMonitoring = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableRebalanceDraining - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.EnableRebalanceDraining = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnablePrometheusMetrics - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.EnablePrometheusMetrics = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableSqsTerminationDraining - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.EnableSQSTerminationDraining = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExcludeFromLoadBalancers - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.ExcludeFromLoadBalancers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ManagedASGTag - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.ManagedASGTag = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.NodeTerminationHandlerSpec {
					subject := kops.NodeTerminationHandlerSpec{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceNodeTerminationHandlerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceNodeTerminationHandlerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
