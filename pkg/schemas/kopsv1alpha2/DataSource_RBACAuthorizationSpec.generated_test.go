package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceRBACAuthorizationSpec(t *testing.T) {
	_default := kopsv1alpha2.RBACAuthorizationSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.RBACAuthorizationSpec
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
			got := ExpandDataSourceRBACAuthorizationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceRBACAuthorizationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceRBACAuthorizationSpecInto(t *testing.T) {
	_default := map[string]interface{}{}
	type args struct {
		in kopsv1alpha2.RBACAuthorizationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.RBACAuthorizationSpec{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceRBACAuthorizationSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceRBACAuthorizationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceRBACAuthorizationSpec(t *testing.T) {
	_default := map[string]interface{}{}
	type args struct {
		in kopsv1alpha2.RBACAuthorizationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.RBACAuthorizationSpec{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceRBACAuthorizationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceRBACAuthorizationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
