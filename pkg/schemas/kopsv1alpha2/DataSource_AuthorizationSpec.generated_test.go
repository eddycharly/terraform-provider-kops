package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceAuthorizationSpec(t *testing.T) {
	_default := kopsv1alpha2.AuthorizationSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.AuthorizationSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"always_allow": nil,
					"rbac":         nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceAuthorizationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAuthorizationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAuthorizationSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"always_allow": nil,
		"rbac":         nil,
	}
	type args struct {
		in kopsv1alpha2.AuthorizationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AuthorizationSpec{},
			},
			want: _default,
		},
		{
			name: "AlwaysAllow - default",
			args: args{
				in: func() kopsv1alpha2.AuthorizationSpec {
					subject := kopsv1alpha2.AuthorizationSpec{}
					subject.AlwaysAllow = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Rbac - default",
			args: args{
				in: func() kopsv1alpha2.AuthorizationSpec {
					subject := kopsv1alpha2.AuthorizationSpec{}
					subject.RBAC = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceAuthorizationSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAuthorizationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAuthorizationSpec(t *testing.T) {
	_default := map[string]interface{}{
		"always_allow": nil,
		"rbac":         nil,
	}
	type args struct {
		in kopsv1alpha2.AuthorizationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AuthorizationSpec{},
			},
			want: _default,
		},
		{
			name: "AlwaysAllow - default",
			args: args{
				in: func() kopsv1alpha2.AuthorizationSpec {
					subject := kopsv1alpha2.AuthorizationSpec{}
					subject.AlwaysAllow = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Rbac - default",
			args: args{
				in: func() kopsv1alpha2.AuthorizationSpec {
					subject := kopsv1alpha2.AuthorizationSpec{}
					subject.RBAC = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceAuthorizationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAuthorizationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
