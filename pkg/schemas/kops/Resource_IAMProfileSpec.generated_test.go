package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceIAMProfileSpec(t *testing.T) {
	_default := kops.IAMProfileSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.IAMProfileSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"profile": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceIAMProfileSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceIAMProfileSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceIAMProfileSpecInto(t *testing.T) {
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
			FlattenResourceIAMProfileSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceIAMProfileSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceIAMProfileSpec(t *testing.T) {
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
			got := FlattenResourceIAMProfileSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceIAMProfileSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
