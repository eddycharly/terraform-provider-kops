package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceAmazonVPCNetworkingSpec(t *testing.T) {
	_default := kopsv1alpha2.AmazonVPCNetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.AmazonVPCNetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"image":      "",
					"init_image": "",
					"env":        func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceAmazonVPCNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAmazonVPCNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAmazonVPCNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"image":      "",
		"init_image": "",
		"env":        func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.AmazonVPCNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AmazonVPCNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.AmazonVPCNetworkingSpec {
					subject := kopsv1alpha2.AmazonVPCNetworkingSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InitImage - default",
			args: args{
				in: func() kopsv1alpha2.AmazonVPCNetworkingSpec {
					subject := kopsv1alpha2.AmazonVPCNetworkingSpec{}
					subject.InitImage = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Env - default",
			args: args{
				in: func() kopsv1alpha2.AmazonVPCNetworkingSpec {
					subject := kopsv1alpha2.AmazonVPCNetworkingSpec{}
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
			FlattenDataSourceAmazonVPCNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAmazonVPCNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAmazonVPCNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"image":      "",
		"init_image": "",
		"env":        func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.AmazonVPCNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AmazonVPCNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.AmazonVPCNetworkingSpec {
					subject := kopsv1alpha2.AmazonVPCNetworkingSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InitImage - default",
			args: args{
				in: func() kopsv1alpha2.AmazonVPCNetworkingSpec {
					subject := kopsv1alpha2.AmazonVPCNetworkingSpec{}
					subject.InitImage = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Env - default",
			args: args{
				in: func() kopsv1alpha2.AmazonVPCNetworkingSpec {
					subject := kopsv1alpha2.AmazonVPCNetworkingSpec{}
					subject.Env = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceAmazonVPCNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAmazonVPCNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
