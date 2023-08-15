package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceKuberouterNetworkingSpec(t *testing.T) {
	_default := kopsv1alpha2.KuberouterNetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.KuberouterNetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceKuberouterNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceKuberouterNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKuberouterNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{}
	type args struct {
		in kopsv1alpha2.KuberouterNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.KuberouterNetworkingSpec{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceKuberouterNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKuberouterNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKuberouterNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{}
	type args struct {
		in kopsv1alpha2.KuberouterNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.KuberouterNetworkingSpec{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceKuberouterNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKuberouterNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
