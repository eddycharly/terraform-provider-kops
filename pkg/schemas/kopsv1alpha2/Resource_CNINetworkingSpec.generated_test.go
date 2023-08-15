package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceCNINetworkingSpec(t *testing.T) {
	_default := kopsv1alpha2.CNINetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.CNINetworkingSpec
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
		in kopsv1alpha2.CNINetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.CNINetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "UsesSecondaryIp - default",
			args: args{
				in: func() kopsv1alpha2.CNINetworkingSpec {
					subject := kopsv1alpha2.CNINetworkingSpec{}
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
		in kopsv1alpha2.CNINetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.CNINetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "UsesSecondaryIp - default",
			args: args{
				in: func() kopsv1alpha2.CNINetworkingSpec {
					subject := kopsv1alpha2.CNINetworkingSpec{}
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
