package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceAlwaysAllowAuthorizationSpec(t *testing.T) {
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
			if got := ExpandResourceAlwaysAllowAuthorizationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceAlwaysAllowAuthorizationSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceAlwaysAllowAuthorizationSpecInto(t *testing.T) {
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
			FlattenResourceAlwaysAllowAuthorizationSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceAlwaysAllowAuthorizationSpec(t *testing.T) {
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
			if got := FlattenResourceAlwaysAllowAuthorizationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceAlwaysAllowAuthorizationSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
