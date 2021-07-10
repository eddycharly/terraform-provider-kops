package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceNodeTerminationHandlerConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.NodeTerminationHandlerConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceNodeTerminationHandlerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceNodeTerminationHandlerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceNodeTerminationHandlerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":                           nil,
		"enable_spot_interruption_draining": nil,
		"enable_scheduled_event_draining":   nil,
		"enable_prometheus_metrics":         nil,
		"enable_sqs_termination_draining":   nil,
		"managed_asg_tag":                   nil,
		"memory_request":                    nil,
		"cpu_request":                       nil,
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
			name: "CPURequest - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.CPURequest = nil
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
		"enable_prometheus_metrics":         nil,
		"enable_sqs_termination_draining":   nil,
		"managed_asg_tag":                   nil,
		"memory_request":                    nil,
		"cpu_request":                       nil,
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
			name: "CPURequest - default",
			args: args{
				in: func() kops.NodeTerminationHandlerConfig {
					subject := kops.NodeTerminationHandlerConfig{}
					subject.CPURequest = nil
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
