package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceKopeioAuthenticationSpec(t *testing.T) {
	_default := kopsv1alpha2.KopeioAuthenticationSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.KopeioAuthenticationSpec
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
		in kopsv1alpha2.KopeioAuthenticationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.KopeioAuthenticationSpec{},
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
		in kopsv1alpha2.KopeioAuthenticationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.KopeioAuthenticationSpec{},
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
