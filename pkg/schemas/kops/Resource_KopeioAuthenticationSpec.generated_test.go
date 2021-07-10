package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceKopeioAuthenticationSpec(t *testing.T) {
	_default := kops.KopeioAuthenticationSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KopeioAuthenticationSpec
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
			got := ExpandResourceKopeioAuthenticationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceKopeioAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceKopeioAuthenticationSpecInto(t *testing.T) {
	_default := map[string]interface{}{}
	type args struct {
		in kops.KopeioAuthenticationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KopeioAuthenticationSpec{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceKopeioAuthenticationSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceKopeioAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceKopeioAuthenticationSpec(t *testing.T) {
	_default := map[string]interface{}{}
	type args struct {
		in kops.KopeioAuthenticationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KopeioAuthenticationSpec{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceKopeioAuthenticationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceKopeioAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
