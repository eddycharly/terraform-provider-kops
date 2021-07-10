package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceNodeTerminationHandlerConfig(t *testing.T) {
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
			if got := ExpandDataSourceNodeTerminationHandlerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceNodeTerminationHandlerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceNodeTerminationHandlerConfigInto(t *testing.T) {
	type args struct {
		in  kops.NodeTerminationHandlerConfig
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
			FlattenDataSourceNodeTerminationHandlerConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceNodeTerminationHandlerConfig(t *testing.T) {
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
			want: map[string]interface{}{
				"enabled":                           nil,
				"enable_spot_interruption_draining": nil,
				"enable_scheduled_event_draining":   nil,
				"enable_prometheus_metrics":         nil,
				"enable_sqs_termination_draining":   nil,
				"managed_asg_tag":                   nil,
				"memory_request":                    nil,
				"cpu_request":                       nil,
			},
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
			want: map[string]interface{}{
				"enabled":                           nil,
				"enable_spot_interruption_draining": nil,
				"enable_scheduled_event_draining":   nil,
				"enable_prometheus_metrics":         nil,
				"enable_sqs_termination_draining":   nil,
				"managed_asg_tag":                   nil,
				"memory_request":                    nil,
				"cpu_request":                       nil,
			},
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
			want: map[string]interface{}{
				"enabled":                           nil,
				"enable_spot_interruption_draining": nil,
				"enable_scheduled_event_draining":   nil,
				"enable_prometheus_metrics":         nil,
				"enable_sqs_termination_draining":   nil,
				"managed_asg_tag":                   nil,
				"memory_request":                    nil,
				"cpu_request":                       nil,
			},
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
			want: map[string]interface{}{
				"enabled":                           nil,
				"enable_spot_interruption_draining": nil,
				"enable_scheduled_event_draining":   nil,
				"enable_prometheus_metrics":         nil,
				"enable_sqs_termination_draining":   nil,
				"managed_asg_tag":                   nil,
				"memory_request":                    nil,
				"cpu_request":                       nil,
			},
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
			want: map[string]interface{}{
				"enabled":                           nil,
				"enable_spot_interruption_draining": nil,
				"enable_scheduled_event_draining":   nil,
				"enable_prometheus_metrics":         nil,
				"enable_sqs_termination_draining":   nil,
				"managed_asg_tag":                   nil,
				"memory_request":                    nil,
				"cpu_request":                       nil,
			},
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
			want: map[string]interface{}{
				"enabled":                           nil,
				"enable_spot_interruption_draining": nil,
				"enable_scheduled_event_draining":   nil,
				"enable_prometheus_metrics":         nil,
				"enable_sqs_termination_draining":   nil,
				"managed_asg_tag":                   nil,
				"memory_request":                    nil,
				"cpu_request":                       nil,
			},
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
			want: map[string]interface{}{
				"enabled":                           nil,
				"enable_spot_interruption_draining": nil,
				"enable_scheduled_event_draining":   nil,
				"enable_prometheus_metrics":         nil,
				"enable_sqs_termination_draining":   nil,
				"managed_asg_tag":                   nil,
				"memory_request":                    nil,
				"cpu_request":                       nil,
			},
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
			want: map[string]interface{}{
				"enabled":                           nil,
				"enable_spot_interruption_draining": nil,
				"enable_scheduled_event_draining":   nil,
				"enable_prometheus_metrics":         nil,
				"enable_sqs_termination_draining":   nil,
				"managed_asg_tag":                   nil,
				"memory_request":                    nil,
				"cpu_request":                       nil,
			},
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
			want: map[string]interface{}{
				"enabled":                           nil,
				"enable_spot_interruption_draining": nil,
				"enable_scheduled_event_draining":   nil,
				"enable_prometheus_metrics":         nil,
				"enable_sqs_termination_draining":   nil,
				"managed_asg_tag":                   nil,
				"memory_request":                    nil,
				"cpu_request":                       nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceNodeTerminationHandlerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceNodeTerminationHandlerConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
