package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAmazonVPCNetworkingSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AmazonVPCNetworkingSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceAmazonVPCNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceAmazonVPCNetworkingSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceAmazonVPCNetworkingSpecInto(t *testing.T) {
	type args struct {
		in  kops.AmazonVPCNetworkingSpec
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
			FlattenDataSourceAmazonVPCNetworkingSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceAmazonVPCNetworkingSpec(t *testing.T) {
	type args struct {
		in kops.AmazonVPCNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AmazonVPCNetworkingSpec{},
			},
			want: map[string]interface{}{
				"image_name":      "",
				"init_image_name": "",
				"env":             func() []interface{} { return nil }(),
			},
		},
		{
			name: "ImageName - default",
			args: args{
				in: func() kops.AmazonVPCNetworkingSpec {
					subject := kops.AmazonVPCNetworkingSpec{}
					subject.ImageName = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image_name":      "",
				"init_image_name": "",
				"env":             func() []interface{} { return nil }(),
			},
		},
		{
			name: "InitImageName - default",
			args: args{
				in: func() kops.AmazonVPCNetworkingSpec {
					subject := kops.AmazonVPCNetworkingSpec{}
					subject.InitImageName = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image_name":      "",
				"init_image_name": "",
				"env":             func() []interface{} { return nil }(),
			},
		},
		{
			name: "Env - default",
			args: args{
				in: func() kops.AmazonVPCNetworkingSpec {
					subject := kops.AmazonVPCNetworkingSpec{}
					subject.Env = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image_name":      "",
				"init_image_name": "",
				"env":             func() []interface{} { return nil }(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceAmazonVPCNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceAmazonVPCNetworkingSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
