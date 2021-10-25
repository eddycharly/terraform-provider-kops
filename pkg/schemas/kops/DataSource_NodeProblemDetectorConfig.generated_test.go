package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceNodeProblemDetectorConfig(t *testing.T) {
	_default := kops.NodeProblemDetectorConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.NodeProblemDetectorConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":        nil,
					"image":          nil,
					"memory_request": nil,
					"cpu_request":    nil,
					"memory_limit":   nil,
					"cpu_limit":      nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceNodeProblemDetectorConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceNodeProblemDetectorConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNodeProblemDetectorConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":        nil,
		"image":          nil,
		"memory_request": nil,
		"cpu_request":    nil,
		"memory_limit":   nil,
		"cpu_limit":      nil,
	}
	type args struct {
		in kops.NodeProblemDetectorConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.NodeProblemDetectorConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.NodeProblemDetectorConfig {
					subject := kops.NodeProblemDetectorConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.NodeProblemDetectorConfig {
					subject := kops.NodeProblemDetectorConfig{}
					subject.Image = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.NodeProblemDetectorConfig {
					subject := kops.NodeProblemDetectorConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.NodeProblemDetectorConfig {
					subject := kops.NodeProblemDetectorConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.NodeProblemDetectorConfig {
					subject := kops.NodeProblemDetectorConfig{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuLimit - default",
			args: args{
				in: func() kops.NodeProblemDetectorConfig {
					subject := kops.NodeProblemDetectorConfig{}
					subject.CPULimit = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceNodeProblemDetectorConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNodeProblemDetectorConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNodeProblemDetectorConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":        nil,
		"image":          nil,
		"memory_request": nil,
		"cpu_request":    nil,
		"memory_limit":   nil,
		"cpu_limit":      nil,
	}
	type args struct {
		in kops.NodeProblemDetectorConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.NodeProblemDetectorConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.NodeProblemDetectorConfig {
					subject := kops.NodeProblemDetectorConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.NodeProblemDetectorConfig {
					subject := kops.NodeProblemDetectorConfig{}
					subject.Image = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.NodeProblemDetectorConfig {
					subject := kops.NodeProblemDetectorConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.NodeProblemDetectorConfig {
					subject := kops.NodeProblemDetectorConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.NodeProblemDetectorConfig {
					subject := kops.NodeProblemDetectorConfig{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuLimit - default",
			args: args{
				in: func() kops.NodeProblemDetectorConfig {
					subject := kops.NodeProblemDetectorConfig{}
					subject.CPULimit = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceNodeProblemDetectorConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNodeProblemDetectorConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
