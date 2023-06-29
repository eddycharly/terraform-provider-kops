package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceAWSAuthenticationSpec(t *testing.T) {
	_default := kopsv1alpha2.AWSAuthenticationSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.AWSAuthenticationSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"image":             "",
					"backend_mode":      "",
					"cluster_id":        "",
					"memory_request":    nil,
					"cpu_request":       nil,
					"memory_limit":      nil,
					"cpu_limit":         nil,
					"identity_mappings": func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceAWSAuthenticationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceAWSAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAWSAuthenticationSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"image":             "",
		"backend_mode":      "",
		"cluster_id":        "",
		"memory_request":    nil,
		"cpu_request":       nil,
		"memory_limit":      nil,
		"cpu_limit":         nil,
		"identity_mappings": func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.AWSAuthenticationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AWSAuthenticationSpec{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationSpec {
					subject := kopsv1alpha2.AWSAuthenticationSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BackendMode - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationSpec {
					subject := kopsv1alpha2.AWSAuthenticationSpec{}
					subject.BackendMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterId - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationSpec {
					subject := kopsv1alpha2.AWSAuthenticationSpec{}
					subject.ClusterID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationSpec {
					subject := kopsv1alpha2.AWSAuthenticationSpec{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationSpec {
					subject := kopsv1alpha2.AWSAuthenticationSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationSpec {
					subject := kopsv1alpha2.AWSAuthenticationSpec{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuLimit - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationSpec {
					subject := kopsv1alpha2.AWSAuthenticationSpec{}
					subject.CPULimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdentityMappings - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationSpec {
					subject := kopsv1alpha2.AWSAuthenticationSpec{}
					subject.IdentityMappings = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceAWSAuthenticationSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAWSAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAWSAuthenticationSpec(t *testing.T) {
	_default := map[string]interface{}{
		"image":             "",
		"backend_mode":      "",
		"cluster_id":        "",
		"memory_request":    nil,
		"cpu_request":       nil,
		"memory_limit":      nil,
		"cpu_limit":         nil,
		"identity_mappings": func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.AWSAuthenticationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AWSAuthenticationSpec{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationSpec {
					subject := kopsv1alpha2.AWSAuthenticationSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BackendMode - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationSpec {
					subject := kopsv1alpha2.AWSAuthenticationSpec{}
					subject.BackendMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterId - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationSpec {
					subject := kopsv1alpha2.AWSAuthenticationSpec{}
					subject.ClusterID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationSpec {
					subject := kopsv1alpha2.AWSAuthenticationSpec{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationSpec {
					subject := kopsv1alpha2.AWSAuthenticationSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationSpec {
					subject := kopsv1alpha2.AWSAuthenticationSpec{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuLimit - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationSpec {
					subject := kopsv1alpha2.AWSAuthenticationSpec{}
					subject.CPULimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdentityMappings - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationSpec {
					subject := kopsv1alpha2.AWSAuthenticationSpec{}
					subject.IdentityMappings = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceAWSAuthenticationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAWSAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
