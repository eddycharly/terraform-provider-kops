package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceCNINetworkingSpec(t *testing.T) {
	_default := kops.CNINetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.CNINetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"uses_secondary_ip": false,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceCNINetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceCNINetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCNINetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"uses_secondary_ip": false,
	}
	type args struct {
		in kops.CNINetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CNINetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "UsesSecondaryIp - default",
			args: args{
				in: func() kops.CNINetworkingSpec {
					subject := kops.CNINetworkingSpec{}
					subject.UsesSecondaryIP = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceCNINetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCNINetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCNINetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"uses_secondary_ip": false,
	}
	type args struct {
		in kops.CNINetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CNINetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "UsesSecondaryIp - default",
			args: args{
				in: func() kops.CNINetworkingSpec {
					subject := kops.CNINetworkingSpec{}
					subject.UsesSecondaryIP = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceCNINetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCNINetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
