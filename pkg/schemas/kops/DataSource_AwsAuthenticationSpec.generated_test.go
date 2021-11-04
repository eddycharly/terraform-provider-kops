package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAwsAuthenticationSpec(t *testing.T) {
	_default := kops.AwsAuthenticationSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AwsAuthenticationSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"image":          "",
					"backend_mode":   "",
					"cluster_id":     "",
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
			got := ExpandDataSourceAwsAuthenticationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAwsAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAwsAuthenticationSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"image":          "",
		"backend_mode":   "",
		"cluster_id":     "",
		"memory_request": nil,
		"cpu_request":    nil,
		"memory_limit":   nil,
		"cpu_limit":      nil,
	}
	type args struct {
		in kops.AwsAuthenticationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AwsAuthenticationSpec{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.AwsAuthenticationSpec {
					subject := kops.AwsAuthenticationSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BackendMode - default",
			args: args{
				in: func() kops.AwsAuthenticationSpec {
					subject := kops.AwsAuthenticationSpec{}
					subject.BackendMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterID - default",
			args: args{
				in: func() kops.AwsAuthenticationSpec {
					subject := kops.AwsAuthenticationSpec{}
					subject.ClusterID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.AwsAuthenticationSpec {
					subject := kops.AwsAuthenticationSpec{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.AwsAuthenticationSpec {
					subject := kops.AwsAuthenticationSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.AwsAuthenticationSpec {
					subject := kops.AwsAuthenticationSpec{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuLimit - default",
			args: args{
				in: func() kops.AwsAuthenticationSpec {
					subject := kops.AwsAuthenticationSpec{}
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
			FlattenDataSourceAwsAuthenticationSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAwsAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAwsAuthenticationSpec(t *testing.T) {
	_default := map[string]interface{}{
		"image":          "",
		"backend_mode":   "",
		"cluster_id":     "",
		"memory_request": nil,
		"cpu_request":    nil,
		"memory_limit":   nil,
		"cpu_limit":      nil,
	}
	type args struct {
		in kops.AwsAuthenticationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AwsAuthenticationSpec{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.AwsAuthenticationSpec {
					subject := kops.AwsAuthenticationSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BackendMode - default",
			args: args{
				in: func() kops.AwsAuthenticationSpec {
					subject := kops.AwsAuthenticationSpec{}
					subject.BackendMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterID - default",
			args: args{
				in: func() kops.AwsAuthenticationSpec {
					subject := kops.AwsAuthenticationSpec{}
					subject.ClusterID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.AwsAuthenticationSpec {
					subject := kops.AwsAuthenticationSpec{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.AwsAuthenticationSpec {
					subject := kops.AwsAuthenticationSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.AwsAuthenticationSpec {
					subject := kops.AwsAuthenticationSpec{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuLimit - default",
			args: args{
				in: func() kops.AwsAuthenticationSpec {
					subject := kops.AwsAuthenticationSpec{}
					subject.CPULimit = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceAwsAuthenticationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAwsAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
