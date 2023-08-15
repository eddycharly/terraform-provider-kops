package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceBastionLoadBalancerSpec(t *testing.T) {
	_default := kopsv1alpha2.BastionLoadBalancerSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.BastionLoadBalancerSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"additional_security_groups": func() []interface{} { return nil }(),
					"type":                       "",
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
		"type":                       "",
	}
	type args struct {
		in kopsv1alpha2.BastionLoadBalancerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.BastionLoadBalancerSpec{},
			},
			want: _default,
		},
		{
			name: "AdditionalSecurityGroups - default",
			args: args{
				in: func() kopsv1alpha2.BastionLoadBalancerSpec {
					subject := kopsv1alpha2.BastionLoadBalancerSpec{}
					subject.AdditionalSecurityGroups = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kopsv1alpha2.BastionLoadBalancerSpec {
					subject := kopsv1alpha2.BastionLoadBalancerSpec{}
					subject.Type = ""
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
		"type":                       "",
	}
	type args struct {
		in kopsv1alpha2.BastionLoadBalancerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.BastionLoadBalancerSpec{},
			},
			want: _default,
		},
		{
			name: "AdditionalSecurityGroups - default",
			args: args{
				in: func() kopsv1alpha2.BastionLoadBalancerSpec {
					subject := kopsv1alpha2.BastionLoadBalancerSpec{}
					subject.AdditionalSecurityGroups = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kopsv1alpha2.BastionLoadBalancerSpec {
					subject := kopsv1alpha2.BastionLoadBalancerSpec{}
					subject.Type = ""
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
