package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceBastionLoadBalancerSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.BastionLoadBalancerSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceBastionLoadBalancerSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceBastionLoadBalancerSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceBastionLoadBalancerSpecInto(t *testing.T) {
	type args struct {
		in  kops.BastionLoadBalancerSpec
		out map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FlattenResourceBastionLoadBalancerSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceBastionLoadBalancerSpec(t *testing.T) {
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
			want: map[string]interface{}{
				"additional_security_groups": func() []interface{} { return nil }(),
			},
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
			want: map[string]interface{}{
				"additional_security_groups": func() []interface{} { return nil }(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceBastionLoadBalancerSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceBastionLoadBalancerSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
