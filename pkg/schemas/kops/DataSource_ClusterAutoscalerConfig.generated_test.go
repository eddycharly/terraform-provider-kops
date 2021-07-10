package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceClusterAutoscalerConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ClusterAutoscalerConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceClusterAutoscalerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceClusterAutoscalerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceClusterAutoscalerConfigInto(t *testing.T) {
	type args struct {
		in  kops.ClusterAutoscalerConfig
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
			FlattenDataSourceClusterAutoscalerConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceClusterAutoscalerConfig(t *testing.T) {
	type args struct {
		in kops.ClusterAutoscalerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ClusterAutoscalerConfig{},
			},
			want: map[string]interface{}{
				"enabled":                          nil,
				"expander":                         nil,
				"balance_similar_node_groups":      nil,
				"scale_down_utilization_threshold": nil,
				"skip_nodes_with_system_pods":      nil,
				"skip_nodes_with_local_storage":    nil,
				"new_pod_scale_up_delay":           nil,
				"scale_down_delay_after_add":       nil,
				"image":                            nil,
				"memory_request":                   nil,
				"cpu_request":                      nil,
			},
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.ClusterAutoscalerConfig {
					subject := kops.ClusterAutoscalerConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":                          nil,
				"expander":                         nil,
				"balance_similar_node_groups":      nil,
				"scale_down_utilization_threshold": nil,
				"skip_nodes_with_system_pods":      nil,
				"skip_nodes_with_local_storage":    nil,
				"new_pod_scale_up_delay":           nil,
				"scale_down_delay_after_add":       nil,
				"image":                            nil,
				"memory_request":                   nil,
				"cpu_request":                      nil,
			},
		},
		{
			name: "Expander - default",
			args: args{
				in: func() kops.ClusterAutoscalerConfig {
					subject := kops.ClusterAutoscalerConfig{}
					subject.Expander = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":                          nil,
				"expander":                         nil,
				"balance_similar_node_groups":      nil,
				"scale_down_utilization_threshold": nil,
				"skip_nodes_with_system_pods":      nil,
				"skip_nodes_with_local_storage":    nil,
				"new_pod_scale_up_delay":           nil,
				"scale_down_delay_after_add":       nil,
				"image":                            nil,
				"memory_request":                   nil,
				"cpu_request":                      nil,
			},
		},
		{
			name: "BalanceSimilarNodeGroups - default",
			args: args{
				in: func() kops.ClusterAutoscalerConfig {
					subject := kops.ClusterAutoscalerConfig{}
					subject.BalanceSimilarNodeGroups = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":                          nil,
				"expander":                         nil,
				"balance_similar_node_groups":      nil,
				"scale_down_utilization_threshold": nil,
				"skip_nodes_with_system_pods":      nil,
				"skip_nodes_with_local_storage":    nil,
				"new_pod_scale_up_delay":           nil,
				"scale_down_delay_after_add":       nil,
				"image":                            nil,
				"memory_request":                   nil,
				"cpu_request":                      nil,
			},
		},
		{
			name: "ScaleDownUtilizationThreshold - default",
			args: args{
				in: func() kops.ClusterAutoscalerConfig {
					subject := kops.ClusterAutoscalerConfig{}
					subject.ScaleDownUtilizationThreshold = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":                          nil,
				"expander":                         nil,
				"balance_similar_node_groups":      nil,
				"scale_down_utilization_threshold": nil,
				"skip_nodes_with_system_pods":      nil,
				"skip_nodes_with_local_storage":    nil,
				"new_pod_scale_up_delay":           nil,
				"scale_down_delay_after_add":       nil,
				"image":                            nil,
				"memory_request":                   nil,
				"cpu_request":                      nil,
			},
		},
		{
			name: "SkipNodesWithSystemPods - default",
			args: args{
				in: func() kops.ClusterAutoscalerConfig {
					subject := kops.ClusterAutoscalerConfig{}
					subject.SkipNodesWithSystemPods = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":                          nil,
				"expander":                         nil,
				"balance_similar_node_groups":      nil,
				"scale_down_utilization_threshold": nil,
				"skip_nodes_with_system_pods":      nil,
				"skip_nodes_with_local_storage":    nil,
				"new_pod_scale_up_delay":           nil,
				"scale_down_delay_after_add":       nil,
				"image":                            nil,
				"memory_request":                   nil,
				"cpu_request":                      nil,
			},
		},
		{
			name: "SkipNodesWithLocalStorage - default",
			args: args{
				in: func() kops.ClusterAutoscalerConfig {
					subject := kops.ClusterAutoscalerConfig{}
					subject.SkipNodesWithLocalStorage = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":                          nil,
				"expander":                         nil,
				"balance_similar_node_groups":      nil,
				"scale_down_utilization_threshold": nil,
				"skip_nodes_with_system_pods":      nil,
				"skip_nodes_with_local_storage":    nil,
				"new_pod_scale_up_delay":           nil,
				"scale_down_delay_after_add":       nil,
				"image":                            nil,
				"memory_request":                   nil,
				"cpu_request":                      nil,
			},
		},
		{
			name: "NewPodScaleUpDelay - default",
			args: args{
				in: func() kops.ClusterAutoscalerConfig {
					subject := kops.ClusterAutoscalerConfig{}
					subject.NewPodScaleUpDelay = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":                          nil,
				"expander":                         nil,
				"balance_similar_node_groups":      nil,
				"scale_down_utilization_threshold": nil,
				"skip_nodes_with_system_pods":      nil,
				"skip_nodes_with_local_storage":    nil,
				"new_pod_scale_up_delay":           nil,
				"scale_down_delay_after_add":       nil,
				"image":                            nil,
				"memory_request":                   nil,
				"cpu_request":                      nil,
			},
		},
		{
			name: "ScaleDownDelayAfterAdd - default",
			args: args{
				in: func() kops.ClusterAutoscalerConfig {
					subject := kops.ClusterAutoscalerConfig{}
					subject.ScaleDownDelayAfterAdd = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":                          nil,
				"expander":                         nil,
				"balance_similar_node_groups":      nil,
				"scale_down_utilization_threshold": nil,
				"skip_nodes_with_system_pods":      nil,
				"skip_nodes_with_local_storage":    nil,
				"new_pod_scale_up_delay":           nil,
				"scale_down_delay_after_add":       nil,
				"image":                            nil,
				"memory_request":                   nil,
				"cpu_request":                      nil,
			},
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.ClusterAutoscalerConfig {
					subject := kops.ClusterAutoscalerConfig{}
					subject.Image = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":                          nil,
				"expander":                         nil,
				"balance_similar_node_groups":      nil,
				"scale_down_utilization_threshold": nil,
				"skip_nodes_with_system_pods":      nil,
				"skip_nodes_with_local_storage":    nil,
				"new_pod_scale_up_delay":           nil,
				"scale_down_delay_after_add":       nil,
				"image":                            nil,
				"memory_request":                   nil,
				"cpu_request":                      nil,
			},
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.ClusterAutoscalerConfig {
					subject := kops.ClusterAutoscalerConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":                          nil,
				"expander":                         nil,
				"balance_similar_node_groups":      nil,
				"scale_down_utilization_threshold": nil,
				"skip_nodes_with_system_pods":      nil,
				"skip_nodes_with_local_storage":    nil,
				"new_pod_scale_up_delay":           nil,
				"scale_down_delay_after_add":       nil,
				"image":                            nil,
				"memory_request":                   nil,
				"cpu_request":                      nil,
			},
		},
		{
			name: "CPURequest - default",
			args: args{
				in: func() kops.ClusterAutoscalerConfig {
					subject := kops.ClusterAutoscalerConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":                          nil,
				"expander":                         nil,
				"balance_similar_node_groups":      nil,
				"scale_down_utilization_threshold": nil,
				"skip_nodes_with_system_pods":      nil,
				"skip_nodes_with_local_storage":    nil,
				"new_pod_scale_up_delay":           nil,
				"scale_down_delay_after_add":       nil,
				"image":                            nil,
				"memory_request":                   nil,
				"cpu_request":                      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceClusterAutoscalerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceClusterAutoscalerConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
