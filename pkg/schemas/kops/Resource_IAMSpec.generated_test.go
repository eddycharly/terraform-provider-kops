package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceIAMSpec(t *testing.T) {
	_default := kops.IAMSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.IAMSpec
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
			got := ExpandResourceIAMSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceIAMSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceIAMSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"legacy":                   false,
		"allow_container_registry": false,
		"permissions_boundary":     nil,
		"use_service_account_external_permissions": nil,
		"service_account_external_permissions":     func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.IAMSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.IAMSpec{},
			},
			want: _default,
		},
		{
			name: "Legacy - default",
			args: args{
				in: func() kops.IAMSpec {
					subject := kops.IAMSpec{}
					subject.Legacy = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowContainerRegistry - default",
			args: args{
				in: func() kops.IAMSpec {
					subject := kops.IAMSpec{}
					subject.AllowContainerRegistry = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PermissionsBoundary - default",
			args: args{
				in: func() kops.IAMSpec {
					subject := kops.IAMSpec{}
					subject.PermissionsBoundary = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseServiceAccountExternalPermissions - default",
			args: args{
				in: func() kops.IAMSpec {
					subject := kops.IAMSpec{}
					subject.UseServiceAccountExternalPermissions = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountExternalPermissions - default",
			args: args{
				in: func() kops.IAMSpec {
					subject := kops.IAMSpec{}
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
			FlattenResourceIAMSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceIAMSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceIAMSpec(t *testing.T) {
	_default := map[string]interface{}{
		"legacy":                   false,
		"allow_container_registry": false,
		"permissions_boundary":     nil,
		"use_service_account_external_permissions": nil,
		"service_account_external_permissions":     func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.IAMSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.IAMSpec{},
			},
			want: _default,
		},
		{
			name: "Legacy - default",
			args: args{
				in: func() kops.IAMSpec {
					subject := kops.IAMSpec{}
					subject.Legacy = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowContainerRegistry - default",
			args: args{
				in: func() kops.IAMSpec {
					subject := kops.IAMSpec{}
					subject.AllowContainerRegistry = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PermissionsBoundary - default",
			args: args{
				in: func() kops.IAMSpec {
					subject := kops.IAMSpec{}
					subject.PermissionsBoundary = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseServiceAccountExternalPermissions - default",
			args: args{
				in: func() kops.IAMSpec {
					subject := kops.IAMSpec{}
					subject.UseServiceAccountExternalPermissions = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountExternalPermissions - default",
			args: args{
				in: func() kops.IAMSpec {
					subject := kops.IAMSpec{}
					subject.ServiceAccountExternalPermissions = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceIAMSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceIAMSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
