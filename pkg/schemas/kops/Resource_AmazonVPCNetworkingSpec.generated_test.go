package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceAmazonVPCNetworkingSpec(t *testing.T) {
	_default := kops.AmazonVPCNetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AmazonVPCNetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"image_name":      "",
					"init_image_name": "",
					"env":             func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceAmazonVPCNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceAmazonVPCNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAmazonVPCNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"image_name":      "",
		"init_image_name": "",
		"env":             func() []interface{} { return nil }(),
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceAmazonVPCNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAmazonVPCNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAmazonVPCNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"image_name":      "",
		"init_image_name": "",
		"env":             func() []interface{} { return nil }(),
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceAmazonVPCNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAmazonVPCNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
