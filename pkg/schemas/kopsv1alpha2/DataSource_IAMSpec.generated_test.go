package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceIAMSpec(t *testing.T) {
	_default := kopsv1alpha2.IAMSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.IAMSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"legacy":                   false,
					"allow_container_registry": false,
					"permissions_boundary":     nil,
					"use_service_account_external_permissions": nil,
					"service_account_external_permissions":     func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceIAMSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceIAMSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceIAMSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"legacy":                   false,
		"allow_container_registry": false,
		"permissions_boundary":     nil,
		"use_service_account_external_permissions": nil,
		"service_account_external_permissions":     func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.IAMSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.IAMSpec{},
			},
			want: _default,
		},
		{
			name: "Legacy - default",
			args: args{
				in: func() kopsv1alpha2.IAMSpec {
					subject := kopsv1alpha2.IAMSpec{}
					subject.Legacy = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowContainerRegistry - default",
			args: args{
				in: func() kopsv1alpha2.IAMSpec {
					subject := kopsv1alpha2.IAMSpec{}
					subject.AllowContainerRegistry = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PermissionsBoundary - default",
			args: args{
				in: func() kopsv1alpha2.IAMSpec {
					subject := kopsv1alpha2.IAMSpec{}
					subject.PermissionsBoundary = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseServiceAccountExternalPermissions - default",
			args: args{
				in: func() kopsv1alpha2.IAMSpec {
					subject := kopsv1alpha2.IAMSpec{}
					subject.UseServiceAccountExternalPermissions = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountExternalPermissions - default",
			args: args{
				in: func() kopsv1alpha2.IAMSpec {
					subject := kopsv1alpha2.IAMSpec{}
					subject.ServiceAccountExternalPermissions = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceIAMSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceIAMSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceIAMSpec(t *testing.T) {
	_default := map[string]interface{}{
		"legacy":                   false,
		"allow_container_registry": false,
		"permissions_boundary":     nil,
		"use_service_account_external_permissions": nil,
		"service_account_external_permissions":     func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.IAMSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.IAMSpec{},
			},
			want: _default,
		},
		{
			name: "Legacy - default",
			args: args{
				in: func() kopsv1alpha2.IAMSpec {
					subject := kopsv1alpha2.IAMSpec{}
					subject.Legacy = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowContainerRegistry - default",
			args: args{
				in: func() kopsv1alpha2.IAMSpec {
					subject := kopsv1alpha2.IAMSpec{}
					subject.AllowContainerRegistry = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PermissionsBoundary - default",
			args: args{
				in: func() kopsv1alpha2.IAMSpec {
					subject := kopsv1alpha2.IAMSpec{}
					subject.PermissionsBoundary = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseServiceAccountExternalPermissions - default",
			args: args{
				in: func() kopsv1alpha2.IAMSpec {
					subject := kopsv1alpha2.IAMSpec{}
					subject.UseServiceAccountExternalPermissions = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountExternalPermissions - default",
			args: args{
				in: func() kopsv1alpha2.IAMSpec {
					subject := kopsv1alpha2.IAMSpec{}
					subject.ServiceAccountExternalPermissions = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceIAMSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceIAMSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
