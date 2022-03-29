package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAWSAuthenticationSpec(t *testing.T) {
	_default := kops.AWSAuthenticationSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AWSAuthenticationSpec
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
			got := ExpandDataSourceAWSAuthenticationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAWSAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAWSAuthenticationSpecInto(t *testing.T) {
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
		in kops.AWSAuthenticationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AWSAuthenticationSpec{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.AWSAuthenticationSpec {
					subject := kops.AWSAuthenticationSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BackendMode - default",
			args: args{
				in: func() kops.AWSAuthenticationSpec {
					subject := kops.AWSAuthenticationSpec{}
					subject.BackendMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterID - default",
			args: args{
				in: func() kops.AWSAuthenticationSpec {
					subject := kops.AWSAuthenticationSpec{}
					subject.ClusterID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.AWSAuthenticationSpec {
					subject := kops.AWSAuthenticationSpec{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.AWSAuthenticationSpec {
					subject := kops.AWSAuthenticationSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.AWSAuthenticationSpec {
					subject := kops.AWSAuthenticationSpec{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuLimit - default",
			args: args{
				in: func() kops.AWSAuthenticationSpec {
					subject := kops.AWSAuthenticationSpec{}
					subject.CPULimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdentityMappings - default",
			args: args{
				in: func() kops.AWSAuthenticationSpec {
					subject := kops.AWSAuthenticationSpec{}
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
			FlattenDataSourceAWSAuthenticationSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAWSAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAWSAuthenticationSpec(t *testing.T) {
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
		in kops.AWSAuthenticationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AWSAuthenticationSpec{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.AWSAuthenticationSpec {
					subject := kops.AWSAuthenticationSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BackendMode - default",
			args: args{
				in: func() kops.AWSAuthenticationSpec {
					subject := kops.AWSAuthenticationSpec{}
					subject.BackendMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterID - default",
			args: args{
				in: func() kops.AWSAuthenticationSpec {
					subject := kops.AWSAuthenticationSpec{}
					subject.ClusterID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.AWSAuthenticationSpec {
					subject := kops.AWSAuthenticationSpec{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.AWSAuthenticationSpec {
					subject := kops.AWSAuthenticationSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.AWSAuthenticationSpec {
					subject := kops.AWSAuthenticationSpec{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuLimit - default",
			args: args{
				in: func() kops.AWSAuthenticationSpec {
					subject := kops.AWSAuthenticationSpec{}
					subject.CPULimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdentityMappings - default",
			args: args{
				in: func() kops.AWSAuthenticationSpec {
					subject := kops.AWSAuthenticationSpec{}
					subject.IdentityMappings = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceAWSAuthenticationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAWSAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
