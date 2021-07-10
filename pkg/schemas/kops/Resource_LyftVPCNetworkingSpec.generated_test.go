package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceLyftVPCNetworkingSpec(t *testing.T) {
	_default := kops.LyftVPCNetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.LyftVPCNetworkingSpec
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
			got := ExpandResourceLyftVPCNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceLyftVPCNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceLyftVPCNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"subnet_tags": func() map[string]interface{} { return nil }(),
	}
	type args struct {
		in kops.LyftVPCNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.LyftVPCNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "SubnetTags - default",
			args: args{
				in: func() kops.LyftVPCNetworkingSpec {
					subject := kops.LyftVPCNetworkingSpec{}
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
			FlattenResourceLyftVPCNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceLyftVPCNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceLyftVPCNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"subnet_tags": func() map[string]interface{} { return nil }(),
	}
	type args struct {
		in kops.LyftVPCNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.LyftVPCNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "SubnetTags - default",
			args: args{
				in: func() kops.LyftVPCNetworkingSpec {
					subject := kops.LyftVPCNetworkingSpec{}
					subject.SubnetTags = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceLyftVPCNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceLyftVPCNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
