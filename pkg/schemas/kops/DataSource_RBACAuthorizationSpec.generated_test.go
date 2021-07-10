package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceRBACAuthorizationSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.RBACAuthorizationSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceRBACAuthorizationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceRBACAuthorizationSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceRBACAuthorizationSpecInto(t *testing.T) {
	type args struct {
		in  kops.RBACAuthorizationSpec
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
			FlattenDataSourceRBACAuthorizationSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceRBACAuthorizationSpec(t *testing.T) {
	type args struct {
		in kops.RBACAuthorizationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.RBACAuthorizationSpec{},
			},
			want: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceRBACAuthorizationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceRBACAuthorizationSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
