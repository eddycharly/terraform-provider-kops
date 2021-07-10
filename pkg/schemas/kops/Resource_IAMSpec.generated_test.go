package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceIAMSpec(t *testing.T) {
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
			if got := ExpandResourceIAMSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceIAMSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceIAMSpecInto(t *testing.T) {
	type args struct {
		in  kops.IAMSpec
		out map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FlattenResourceIAMSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceIAMSpec(t *testing.T) {
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
			want: map[string]interface{}{
				"legacy":                               false,
				"allow_container_registry":             false,
				"permissions_boundary":                 nil,
				"service_account_external_permissions": func() []interface{} { return nil }(),
			},
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
			want: map[string]interface{}{
				"legacy":                               false,
				"allow_container_registry":             false,
				"permissions_boundary":                 nil,
				"service_account_external_permissions": func() []interface{} { return nil }(),
			},
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
			want: map[string]interface{}{
				"legacy":                               false,
				"allow_container_registry":             false,
				"permissions_boundary":                 nil,
				"service_account_external_permissions": func() []interface{} { return nil }(),
			},
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
			want: map[string]interface{}{
				"legacy":                               false,
				"allow_container_registry":             false,
				"permissions_boundary":                 nil,
				"service_account_external_permissions": func() []interface{} { return nil }(),
			},
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
			want: map[string]interface{}{
				"legacy":                               false,
				"allow_container_registry":             false,
				"permissions_boundary":                 nil,
				"service_account_external_permissions": func() []interface{} { return nil }(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceIAMSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceIAMSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
