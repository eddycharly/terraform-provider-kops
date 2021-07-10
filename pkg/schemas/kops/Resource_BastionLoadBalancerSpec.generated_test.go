package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceBastionLoadBalancerSpec(t *testing.T) {
	_default := kops.BastionLoadBalancerSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.BastionLoadBalancerSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"additional_security_groups": func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceBastionLoadBalancerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceBastionLoadBalancerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceBastionLoadBalancerSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"additional_security_groups": func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.BastionLoadBalancerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.BastionLoadBalancerSpec{},
			},
			want: _default,
		},
		{
			name: "AdditionalSecurityGroups - default",
			args: args{
				in: func() kops.BastionLoadBalancerSpec {
					subject := kops.BastionLoadBalancerSpec{}
					subject.AdditionalSecurityGroups = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceBastionLoadBalancerSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceBastionLoadBalancerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceBastionLoadBalancerSpec(t *testing.T) {
	_default := map[string]interface{}{
		"additional_security_groups": func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.BastionLoadBalancerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.BastionLoadBalancerSpec{},
			},
			want: _default,
		},
		{
			name: "AdditionalSecurityGroups - default",
			args: args{
				in: func() kops.BastionLoadBalancerSpec {
					subject := kops.BastionLoadBalancerSpec{}
					subject.AdditionalSecurityGroups = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceBastionLoadBalancerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceBastionLoadBalancerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
