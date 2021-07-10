package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceAwsAuthenticationSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AwsAuthenticationSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceAwsAuthenticationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceAwsAuthenticationSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceAwsAuthenticationSpecInto(t *testing.T) {
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
			name: "CPURequest - default",
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
			name: "CPULimit - default",
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
			FlattenResourceAwsAuthenticationSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAwsAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAwsAuthenticationSpec(t *testing.T) {
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
			name: "CPURequest - default",
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
			name: "CPULimit - default",
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
			got := FlattenResourceAwsAuthenticationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAwsAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
