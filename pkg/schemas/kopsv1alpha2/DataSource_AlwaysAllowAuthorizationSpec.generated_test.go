package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceAlwaysAllowAuthorizationSpec(t *testing.T) {
	_default := kopsv1alpha2.AlwaysAllowAuthorizationSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.AlwaysAllowAuthorizationSpec
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
			got := ExpandDataSourceAlwaysAllowAuthorizationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAlwaysAllowAuthorizationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAlwaysAllowAuthorizationSpecInto(t *testing.T) {
	_default := map[string]interface{}{}
	type args struct {
		in kopsv1alpha2.AlwaysAllowAuthorizationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AlwaysAllowAuthorizationSpec{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceAlwaysAllowAuthorizationSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAlwaysAllowAuthorizationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAlwaysAllowAuthorizationSpec(t *testing.T) {
	_default := map[string]interface{}{}
	type args struct {
		in kopsv1alpha2.AlwaysAllowAuthorizationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AlwaysAllowAuthorizationSpec{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceAlwaysAllowAuthorizationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAlwaysAllowAuthorizationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
