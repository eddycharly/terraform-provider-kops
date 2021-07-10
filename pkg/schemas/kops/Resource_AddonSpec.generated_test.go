package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceAddonSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AddonSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceAddonSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceAddonSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceAddonSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"manifest": "",
	}
	type args struct {
		in kops.AddonSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AddonSpec{},
			},
			want: _default,
		},
		{
			name: "Manifest - default",
			args: args{
				in: func() kops.AddonSpec {
					subject := kops.AddonSpec{}
					subject.Manifest = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceAddonSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAddonSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAddonSpec(t *testing.T) {
	_default := map[string]interface{}{
		"manifest": "",
	}
	type args struct {
		in kops.AddonSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AddonSpec{},
			},
			want: _default,
		},
		{
			name: "Manifest - default",
			args: args{
				in: func() kops.AddonSpec {
					subject := kops.AddonSpec{}
					subject.Manifest = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceAddonSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAddonSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
