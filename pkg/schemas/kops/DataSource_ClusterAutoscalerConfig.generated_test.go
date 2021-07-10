package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceClusterAutoscalerConfig(t *testing.T) {
	_default := kops.ClusterAutoscalerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ClusterAutoscalerConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceClusterAutoscalerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceClusterAutoscalerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceClusterAutoscalerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
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
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceClusterAutoscalerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceClusterAutoscalerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceClusterAutoscalerConfig(t *testing.T) {
	_default := map[string]interface{}{
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
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceClusterAutoscalerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceClusterAutoscalerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
