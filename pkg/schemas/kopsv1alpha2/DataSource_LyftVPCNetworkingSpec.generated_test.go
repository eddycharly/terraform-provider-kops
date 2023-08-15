package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceLyftVPCNetworkingSpec(t *testing.T) {
	_default := kopsv1alpha2.LyftVPCNetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.LyftVPCNetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"subnet_tags": func() map[string]interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceLyftVPCNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceLyftVPCNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLyftVPCNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"subnet_tags": func() map[string]interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.LyftVPCNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.LyftVPCNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "SubnetTags - default",
			args: args{
				in: func() kopsv1alpha2.LyftVPCNetworkingSpec {
					subject := kopsv1alpha2.LyftVPCNetworkingSpec{}
					subject.SubnetTags = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceLyftVPCNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLyftVPCNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLyftVPCNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"subnet_tags": func() map[string]interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.LyftVPCNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.LyftVPCNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "SubnetTags - default",
			args: args{
				in: func() kopsv1alpha2.LyftVPCNetworkingSpec {
					subject := kopsv1alpha2.LyftVPCNetworkingSpec{}
					subject.SubnetTags = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceLyftVPCNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLyftVPCNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
