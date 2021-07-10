package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceIAMSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.IAMSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceIAMSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceIAMSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceIAMSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"legacy":                               false,
		"allow_container_registry":             false,
		"permissions_boundary":                 nil,
		"service_account_external_permissions": func() []interface{} { return nil }(),
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
			FlattenDataSourceIAMSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceIAMSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceIAMSpec(t *testing.T) {
	_default := map[string]interface{}{
		"legacy":                               false,
		"allow_container_registry":             false,
		"permissions_boundary":                 nil,
		"service_account_external_permissions": func() []interface{} { return nil }(),
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
			got := FlattenDataSourceIAMSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceIAMSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
