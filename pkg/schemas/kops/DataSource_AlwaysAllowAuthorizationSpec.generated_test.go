package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAlwaysAllowAuthorizationSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AlwaysAllowAuthorizationSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceAlwaysAllowAuthorizationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceAlwaysAllowAuthorizationSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceAlwaysAllowAuthorizationSpecInto(t *testing.T) {
	type args struct {
		in  kops.AlwaysAllowAuthorizationSpec
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
			FlattenDataSourceAlwaysAllowAuthorizationSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceAlwaysAllowAuthorizationSpec(t *testing.T) {
	type args struct {
		in kops.AlwaysAllowAuthorizationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AlwaysAllowAuthorizationSpec{},
			},
			want: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceAlwaysAllowAuthorizationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceAlwaysAllowAuthorizationSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
