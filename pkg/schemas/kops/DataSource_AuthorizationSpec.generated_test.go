package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAuthorizationSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AuthorizationSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceAuthorizationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceAuthorizationSpec() = %v, want %v", got, tt.want)
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
		in kops.AuthorizationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AuthorizationSpec{},
			},
			want: _default,
		},
		{
			name: "AlwaysAllow - default",
			args: args{
				in: func() kops.AuthorizationSpec {
					subject := kops.AuthorizationSpec{}
					subject.AlwaysAllow = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RBAC - default",
			args: args{
				in: func() kops.AuthorizationSpec {
					subject := kops.AuthorizationSpec{}
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
		in kops.AuthorizationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AuthorizationSpec{},
			},
			want: _default,
		},
		{
			name: "AlwaysAllow - default",
			args: args{
				in: func() kops.AuthorizationSpec {
					subject := kops.AuthorizationSpec{}
					subject.AlwaysAllow = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RBAC - default",
			args: args{
				in: func() kops.AuthorizationSpec {
					subject := kops.AuthorizationSpec{}
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
