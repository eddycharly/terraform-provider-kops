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
	_default := map[string]interface{}{
		"profile": nil,
	}
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceIAMProfileSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceIAMProfileSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceIAMProfileSpec(t *testing.T) {
	_default := map[string]interface{}{
		"profile": nil,
	}
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceIAMProfileSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceIAMProfileSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
