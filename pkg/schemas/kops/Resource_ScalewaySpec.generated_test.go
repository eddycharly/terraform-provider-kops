package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceScalewaySpec(t *testing.T) {
	_default := kops.ScalewaySpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ScalewaySpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceScalewaySpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceScalewaySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceScalewaySpecInto(t *testing.T) {
	_default := map[string]interface{}{}
	type args struct {
		in kops.ScalewaySpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ScalewaySpec{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceScalewaySpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceScalewaySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceScalewaySpec(t *testing.T) {
	_default := map[string]interface{}{}
	type args struct {
		in kops.ScalewaySpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ScalewaySpec{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceScalewaySpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceScalewaySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
