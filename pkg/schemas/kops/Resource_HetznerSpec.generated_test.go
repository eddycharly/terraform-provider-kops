package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceHetznerSpec(t *testing.T) {
	_default := kops.HetznerSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.HetznerSpec
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
			got := ExpandResourceHetznerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceHetznerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceHetznerSpecInto(t *testing.T) {
	_default := map[string]interface{}{}
	type args struct {
		in kops.HetznerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.HetznerSpec{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceHetznerSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceHetznerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceHetznerSpec(t *testing.T) {
	_default := map[string]interface{}{}
	type args struct {
		in kops.HetznerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.HetznerSpec{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceHetznerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceHetznerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
