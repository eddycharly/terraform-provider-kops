package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceIAMProfileSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.IAMProfileSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceIAMProfileSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceIAMProfileSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceIAMProfileSpecInto(t *testing.T) {
	type args struct {
		in  kops.IAMProfileSpec
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
			FlattenDataSourceIAMProfileSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceIAMProfileSpec(t *testing.T) {
	type args struct {
		in kops.IAMProfileSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.IAMProfileSpec{},
			},
			want: map[string]interface{}{
				"profile": nil,
			},
		},
		{
			name: "Profile - default",
			args: args{
				in: func() kops.IAMProfileSpec {
					subject := kops.IAMProfileSpec{}
					subject.Profile = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"profile": nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceIAMProfileSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceIAMProfileSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
