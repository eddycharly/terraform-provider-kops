package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceClusterAutoscalerConfig(t *testing.T) {
	_default := kopsv1alpha2.ClusterAutoscalerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.ClusterAutoscalerConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":                          nil,
					"expander":                         "",
					"balance_similar_node_groups":      nil,
					"aws_use_static_instance_list":     nil,
					"scale_down_utilization_threshold": nil,
					"skip_nodes_with_system_pods":      nil,
					"skip_nodes_with_local_storage":    nil,
					"new_pod_scale_up_delay":           nil,
					"scale_down_delay_after_add":       nil,
					"scale_down_unneeded_time":         nil,
					"scale_down_unready_time":          nil,
					"cordon_node_before_terminating":   nil,
					"image":                            nil,
					"memory_request":                   nil,
					"cpu_request":                      nil,
					"max_node_provision_time":          "",
					"pod_annotations":                  func() map[string]interface{} { return nil }(),
					"create_priority_expender_config":  nil,
					"custom_priority_expander_config":  func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceClusterAutoscalerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceClusterAutoscalerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceClusterAutoscalerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":                          nil,
		"expander":                         "",
		"balance_similar_node_groups":      nil,
		"aws_use_static_instance_list":     nil,
		"scale_down_utilization_threshold": nil,
		"skip_nodes_with_system_pods":      nil,
		"skip_nodes_with_local_storage":    nil,
		"new_pod_scale_up_delay":           nil,
		"scale_down_delay_after_add":       nil,
		"scale_down_unneeded_time":         nil,
		"scale_down_unready_time":          nil,
		"cordon_node_before_terminating":   nil,
		"image":                            nil,
		"memory_request":                   nil,
		"cpu_request":                      nil,
		"max_node_provision_time":          "",
		"pod_annotations":                  func() map[string]interface{} { return nil }(),
		"create_priority_expender_config":  nil,
		"custom_priority_expander_config":  func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.ClusterAutoscalerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.ClusterAutoscalerConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Expander - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.Expander = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BalanceSimilarNodeGroups - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.BalanceSimilarNodeGroups = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AwsUseStaticInstanceList - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.AWSUseStaticInstanceList = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ScaleDownUtilizationThreshold - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.ScaleDownUtilizationThreshold = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SkipNodesWithSystemPods - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.SkipNodesWithSystemPods = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SkipNodesWithLocalStorage - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.SkipNodesWithLocalStorage = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NewPodScaleUpDelay - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.NewPodScaleUpDelay = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ScaleDownDelayAfterAdd - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.ScaleDownDelayAfterAdd = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ScaleDownUnneededTime - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.ScaleDownUnneededTime = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ScaleDownUnreadyTime - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.ScaleDownUnreadyTime = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CordonNodeBeforeTerminating - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.CordonNodeBeforeTerminating = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.Image = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxNodeProvisionTime - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.MaxNodeProvisionTime = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodAnnotations - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.PodAnnotations = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CreatePriorityExpenderConfig - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.CreatePriorityExpenderConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CustomPriorityExpanderConfig - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.CustomPriorityExpanderConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceClusterAutoscalerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceClusterAutoscalerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceClusterAutoscalerConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":                          nil,
		"expander":                         "",
		"balance_similar_node_groups":      nil,
		"aws_use_static_instance_list":     nil,
		"scale_down_utilization_threshold": nil,
		"skip_nodes_with_system_pods":      nil,
		"skip_nodes_with_local_storage":    nil,
		"new_pod_scale_up_delay":           nil,
		"scale_down_delay_after_add":       nil,
		"scale_down_unneeded_time":         nil,
		"scale_down_unready_time":          nil,
		"cordon_node_before_terminating":   nil,
		"image":                            nil,
		"memory_request":                   nil,
		"cpu_request":                      nil,
		"max_node_provision_time":          "",
		"pod_annotations":                  func() map[string]interface{} { return nil }(),
		"create_priority_expender_config":  nil,
		"custom_priority_expander_config":  func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.ClusterAutoscalerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.ClusterAutoscalerConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Expander - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.Expander = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BalanceSimilarNodeGroups - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.BalanceSimilarNodeGroups = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AwsUseStaticInstanceList - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.AWSUseStaticInstanceList = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ScaleDownUtilizationThreshold - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.ScaleDownUtilizationThreshold = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SkipNodesWithSystemPods - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.SkipNodesWithSystemPods = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SkipNodesWithLocalStorage - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.SkipNodesWithLocalStorage = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NewPodScaleUpDelay - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.NewPodScaleUpDelay = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ScaleDownDelayAfterAdd - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.ScaleDownDelayAfterAdd = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ScaleDownUnneededTime - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.ScaleDownUnneededTime = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ScaleDownUnreadyTime - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.ScaleDownUnreadyTime = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CordonNodeBeforeTerminating - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.CordonNodeBeforeTerminating = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.Image = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxNodeProvisionTime - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.MaxNodeProvisionTime = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodAnnotations - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.PodAnnotations = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CreatePriorityExpenderConfig - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.CreatePriorityExpenderConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CustomPriorityExpanderConfig - default",
			args: args{
				in: func() kopsv1alpha2.ClusterAutoscalerConfig {
					subject := kopsv1alpha2.ClusterAutoscalerConfig{}
					subject.CustomPriorityExpanderConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceClusterAutoscalerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceClusterAutoscalerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
