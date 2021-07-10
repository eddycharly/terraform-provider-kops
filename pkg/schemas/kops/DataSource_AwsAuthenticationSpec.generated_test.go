package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAwsAuthenticationSpec(t *testing.T) {
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
			if got := ExpandDataSourceAwsAuthenticationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceAwsAuthenticationSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceAwsAuthenticationSpecInto(t *testing.T) {
	type args struct {
		in  kops.AwsAuthenticationSpec
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
			FlattenDataSourceAwsAuthenticationSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceAwsAuthenticationSpec(t *testing.T) {
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
			want: map[string]interface{}{
				"image":          "",
				"backend_mode":   "",
				"cluster_id":     "",
				"memory_request": nil,
				"cpu_request":    nil,
				"memory_limit":   nil,
				"cpu_limit":      nil,
			},
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
			want: map[string]interface{}{
				"image":          "",
				"backend_mode":   "",
				"cluster_id":     "",
				"memory_request": nil,
				"cpu_request":    nil,
				"memory_limit":   nil,
				"cpu_limit":      nil,
			},
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
			want: map[string]interface{}{
				"image":          "",
				"backend_mode":   "",
				"cluster_id":     "",
				"memory_request": nil,
				"cpu_request":    nil,
				"memory_limit":   nil,
				"cpu_limit":      nil,
			},
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
			want: map[string]interface{}{
				"image":          "",
				"backend_mode":   "",
				"cluster_id":     "",
				"memory_request": nil,
				"cpu_request":    nil,
				"memory_limit":   nil,
				"cpu_limit":      nil,
			},
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
			want: map[string]interface{}{
				"image":          "",
				"backend_mode":   "",
				"cluster_id":     "",
				"memory_request": nil,
				"cpu_request":    nil,
				"memory_limit":   nil,
				"cpu_limit":      nil,
			},
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
			want: map[string]interface{}{
				"image":          "",
				"backend_mode":   "",
				"cluster_id":     "",
				"memory_request": nil,
				"cpu_request":    nil,
				"memory_limit":   nil,
				"cpu_limit":      nil,
			},
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
			want: map[string]interface{}{
				"image":          "",
				"backend_mode":   "",
				"cluster_id":     "",
				"memory_request": nil,
				"cpu_request":    nil,
				"memory_limit":   nil,
				"cpu_limit":      nil,
			},
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
			want: map[string]interface{}{
				"image":          "",
				"backend_mode":   "",
				"cluster_id":     "",
				"memory_request": nil,
				"cpu_request":    nil,
				"memory_limit":   nil,
				"cpu_limit":      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceAwsAuthenticationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceAwsAuthenticationSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
