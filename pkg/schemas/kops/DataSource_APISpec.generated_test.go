package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAPISpec(t *testing.T) {
	_default := kops.APISpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.APISpec
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
			got := ExpandDataSourceAPISpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAPISpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAPISpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"dns":             nil,
		"load_balancer":   nil,
		"public_name":     "",
		"additional_sans": func() []interface{} { return nil }(),
		"access":          func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.APISpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.APISpec{},
			},
			want: _default,
		},
		{
			name: "Dns - default",
			args: args{
				in: func() kops.APISpec {
					subject := kops.APISpec{}
					subject.DNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LoadBalancer - default",
			args: args{
				in: func() kops.APISpec {
					subject := kops.APISpec{}
					subject.LoadBalancer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PublicName - default",
			args: args{
				in: func() kops.APISpec {
					subject := kops.APISpec{}
					subject.PublicName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalSans - default",
			args: args{
				in: func() kops.APISpec {
					subject := kops.APISpec{}
					subject.AdditionalSANs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Access - default",
			args: args{
				in: func() kops.APISpec {
					subject := kops.APISpec{}
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
			FlattenDataSourceAPISpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAPISpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAPISpec(t *testing.T) {
	_default := map[string]interface{}{
		"dns":             nil,
		"load_balancer":   nil,
		"public_name":     "",
		"additional_sans": func() []interface{} { return nil }(),
		"access":          func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.APISpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.APISpec{},
			},
			want: _default,
		},
		{
			name: "Dns - default",
			args: args{
				in: func() kops.APISpec {
					subject := kops.APISpec{}
					subject.DNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LoadBalancer - default",
			args: args{
				in: func() kops.APISpec {
					subject := kops.APISpec{}
					subject.LoadBalancer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PublicName - default",
			args: args{
				in: func() kops.APISpec {
					subject := kops.APISpec{}
					subject.PublicName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalSans - default",
			args: args{
				in: func() kops.APISpec {
					subject := kops.APISpec{}
					subject.AdditionalSANs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Access - default",
			args: args{
				in: func() kops.APISpec {
					subject := kops.APISpec{}
					subject.Access = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceAPISpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAPISpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
