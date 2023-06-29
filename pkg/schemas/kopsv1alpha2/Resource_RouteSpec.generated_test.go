package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceRouteSpec(t *testing.T) {
	_default := kopsv1alpha2.RouteSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.RouteSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"cidr":   "",
					"target": "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceRouteSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceRouteSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceRouteSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"cidr":   "",
		"target": "",
	}
	type args struct {
		in kopsv1alpha2.RouteSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.RouteSpec{},
			},
			want: _default,
		},
		{
			name: "Cidr - default",
			args: args{
				in: func() kopsv1alpha2.RouteSpec {
					subject := kopsv1alpha2.RouteSpec{}
					subject.CIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Target - default",
			args: args{
				in: func() kopsv1alpha2.RouteSpec {
					subject := kopsv1alpha2.RouteSpec{}
					subject.Target = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceRouteSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceRouteSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceRouteSpec(t *testing.T) {
	_default := map[string]interface{}{
		"cidr":   "",
		"target": "",
	}
	type args struct {
		in kopsv1alpha2.RouteSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.RouteSpec{},
			},
			want: _default,
		},
		{
			name: "Cidr - default",
			args: args{
				in: func() kopsv1alpha2.RouteSpec {
					subject := kopsv1alpha2.RouteSpec{}
					subject.CIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Target - default",
			args: args{
				in: func() kopsv1alpha2.RouteSpec {
					subject := kopsv1alpha2.RouteSpec{}
					subject.Target = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceRouteSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceRouteSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
