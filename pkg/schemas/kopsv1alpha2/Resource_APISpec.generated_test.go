package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceAPISpec(t *testing.T) {
	_default := kopsv1alpha2.APISpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.APISpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"dns":             nil,
					"load_balancer":   nil,
					"public_name":     "",
					"additional_sans": func() []interface{} { return nil }(),
					"access":          func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceAPISpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceAPISpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAPISpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"dns":             nil,
		"load_balancer":   nil,
		"public_name":     "",
		"additional_sans": func() []interface{} { return nil }(),
		"access":          func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.APISpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.APISpec{},
			},
			want: _default,
		},
		{
			name: "Dns - default",
			args: args{
				in: func() kopsv1alpha2.APISpec {
					subject := kopsv1alpha2.APISpec{}
					subject.DNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LoadBalancer - default",
			args: args{
				in: func() kopsv1alpha2.APISpec {
					subject := kopsv1alpha2.APISpec{}
					subject.LoadBalancer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PublicName - default",
			args: args{
				in: func() kopsv1alpha2.APISpec {
					subject := kopsv1alpha2.APISpec{}
					subject.PublicName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalSans - default",
			args: args{
				in: func() kopsv1alpha2.APISpec {
					subject := kopsv1alpha2.APISpec{}
					subject.AdditionalSANs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Access - default",
			args: args{
				in: func() kopsv1alpha2.APISpec {
					subject := kopsv1alpha2.APISpec{}
					subject.Access = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceAPISpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAPISpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAPISpec(t *testing.T) {
	_default := map[string]interface{}{
		"dns":             nil,
		"load_balancer":   nil,
		"public_name":     "",
		"additional_sans": func() []interface{} { return nil }(),
		"access":          func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.APISpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.APISpec{},
			},
			want: _default,
		},
		{
			name: "Dns - default",
			args: args{
				in: func() kopsv1alpha2.APISpec {
					subject := kopsv1alpha2.APISpec{}
					subject.DNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LoadBalancer - default",
			args: args{
				in: func() kopsv1alpha2.APISpec {
					subject := kopsv1alpha2.APISpec{}
					subject.LoadBalancer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PublicName - default",
			args: args{
				in: func() kopsv1alpha2.APISpec {
					subject := kopsv1alpha2.APISpec{}
					subject.PublicName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalSans - default",
			args: args{
				in: func() kopsv1alpha2.APISpec {
					subject := kopsv1alpha2.APISpec{}
					subject.AdditionalSANs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Access - default",
			args: args{
				in: func() kopsv1alpha2.APISpec {
					subject := kopsv1alpha2.APISpec{}
					subject.Access = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceAPISpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAPISpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
