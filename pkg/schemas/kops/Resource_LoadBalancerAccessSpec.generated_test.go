package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceLoadBalancerAccessSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.LoadBalancerAccessSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceLoadBalancerAccessSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceLoadBalancerAccessSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceLoadBalancerAccessSpecInto(t *testing.T) {
	type args struct {
		in  kops.LoadBalancerAccessSpec
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
			FlattenResourceLoadBalancerAccessSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceLoadBalancerAccessSpec(t *testing.T) {
	type args struct {
		in kops.LoadBalancerAccessSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.LoadBalancerAccessSpec{},
			},
			want: map[string]interface{}{
				"class":                      "",
				"type":                       "",
				"idle_timeout_seconds":       nil,
				"security_group_override":    nil,
				"additional_security_groups": func() []interface{} { return nil }(),
				"use_for_internal_api":       false,
				"ssl_certificate":            "",
				"ssl_policy":                 nil,
				"cross_zone_load_balancing":  nil,
				"subnets":                    func() []interface{} { return nil }(),
			},
		},
		{
			name: "Class - default",
			args: args{
				in: func() kops.LoadBalancerAccessSpec {
					subject := kops.LoadBalancerAccessSpec{}
					subject.Class = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"class":                      "",
				"type":                       "",
				"idle_timeout_seconds":       nil,
				"security_group_override":    nil,
				"additional_security_groups": func() []interface{} { return nil }(),
				"use_for_internal_api":       false,
				"ssl_certificate":            "",
				"ssl_policy":                 nil,
				"cross_zone_load_balancing":  nil,
				"subnets":                    func() []interface{} { return nil }(),
			},
		},
		{
			name: "Type - default",
			args: args{
				in: func() kops.LoadBalancerAccessSpec {
					subject := kops.LoadBalancerAccessSpec{}
					subject.Type = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"class":                      "",
				"type":                       "",
				"idle_timeout_seconds":       nil,
				"security_group_override":    nil,
				"additional_security_groups": func() []interface{} { return nil }(),
				"use_for_internal_api":       false,
				"ssl_certificate":            "",
				"ssl_policy":                 nil,
				"cross_zone_load_balancing":  nil,
				"subnets":                    func() []interface{} { return nil }(),
			},
		},
		{
			name: "IdleTimeoutSeconds - default",
			args: args{
				in: func() kops.LoadBalancerAccessSpec {
					subject := kops.LoadBalancerAccessSpec{}
					subject.IdleTimeoutSeconds = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"class":                      "",
				"type":                       "",
				"idle_timeout_seconds":       nil,
				"security_group_override":    nil,
				"additional_security_groups": func() []interface{} { return nil }(),
				"use_for_internal_api":       false,
				"ssl_certificate":            "",
				"ssl_policy":                 nil,
				"cross_zone_load_balancing":  nil,
				"subnets":                    func() []interface{} { return nil }(),
			},
		},
		{
			name: "SecurityGroupOverride - default",
			args: args{
				in: func() kops.LoadBalancerAccessSpec {
					subject := kops.LoadBalancerAccessSpec{}
					subject.SecurityGroupOverride = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"class":                      "",
				"type":                       "",
				"idle_timeout_seconds":       nil,
				"security_group_override":    nil,
				"additional_security_groups": func() []interface{} { return nil }(),
				"use_for_internal_api":       false,
				"ssl_certificate":            "",
				"ssl_policy":                 nil,
				"cross_zone_load_balancing":  nil,
				"subnets":                    func() []interface{} { return nil }(),
			},
		},
		{
			name: "AdditionalSecurityGroups - default",
			args: args{
				in: func() kops.LoadBalancerAccessSpec {
					subject := kops.LoadBalancerAccessSpec{}
					subject.AdditionalSecurityGroups = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"class":                      "",
				"type":                       "",
				"idle_timeout_seconds":       nil,
				"security_group_override":    nil,
				"additional_security_groups": func() []interface{} { return nil }(),
				"use_for_internal_api":       false,
				"ssl_certificate":            "",
				"ssl_policy":                 nil,
				"cross_zone_load_balancing":  nil,
				"subnets":                    func() []interface{} { return nil }(),
			},
		},
		{
			name: "UseForInternalApi - default",
			args: args{
				in: func() kops.LoadBalancerAccessSpec {
					subject := kops.LoadBalancerAccessSpec{}
					subject.UseForInternalApi = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"class":                      "",
				"type":                       "",
				"idle_timeout_seconds":       nil,
				"security_group_override":    nil,
				"additional_security_groups": func() []interface{} { return nil }(),
				"use_for_internal_api":       false,
				"ssl_certificate":            "",
				"ssl_policy":                 nil,
				"cross_zone_load_balancing":  nil,
				"subnets":                    func() []interface{} { return nil }(),
			},
		},
		{
			name: "SSLCertificate - default",
			args: args{
				in: func() kops.LoadBalancerAccessSpec {
					subject := kops.LoadBalancerAccessSpec{}
					subject.SSLCertificate = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"class":                      "",
				"type":                       "",
				"idle_timeout_seconds":       nil,
				"security_group_override":    nil,
				"additional_security_groups": func() []interface{} { return nil }(),
				"use_for_internal_api":       false,
				"ssl_certificate":            "",
				"ssl_policy":                 nil,
				"cross_zone_load_balancing":  nil,
				"subnets":                    func() []interface{} { return nil }(),
			},
		},
		{
			name: "SSLPolicy - default",
			args: args{
				in: func() kops.LoadBalancerAccessSpec {
					subject := kops.LoadBalancerAccessSpec{}
					subject.SSLPolicy = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"class":                      "",
				"type":                       "",
				"idle_timeout_seconds":       nil,
				"security_group_override":    nil,
				"additional_security_groups": func() []interface{} { return nil }(),
				"use_for_internal_api":       false,
				"ssl_certificate":            "",
				"ssl_policy":                 nil,
				"cross_zone_load_balancing":  nil,
				"subnets":                    func() []interface{} { return nil }(),
			},
		},
		{
			name: "CrossZoneLoadBalancing - default",
			args: args{
				in: func() kops.LoadBalancerAccessSpec {
					subject := kops.LoadBalancerAccessSpec{}
					subject.CrossZoneLoadBalancing = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"class":                      "",
				"type":                       "",
				"idle_timeout_seconds":       nil,
				"security_group_override":    nil,
				"additional_security_groups": func() []interface{} { return nil }(),
				"use_for_internal_api":       false,
				"ssl_certificate":            "",
				"ssl_policy":                 nil,
				"cross_zone_load_balancing":  nil,
				"subnets":                    func() []interface{} { return nil }(),
			},
		},
		{
			name: "Subnets - default",
			args: args{
				in: func() kops.LoadBalancerAccessSpec {
					subject := kops.LoadBalancerAccessSpec{}
					subject.Subnets = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"class":                      "",
				"type":                       "",
				"idle_timeout_seconds":       nil,
				"security_group_override":    nil,
				"additional_security_groups": func() []interface{} { return nil }(),
				"use_for_internal_api":       false,
				"ssl_certificate":            "",
				"ssl_policy":                 nil,
				"cross_zone_load_balancing":  nil,
				"subnets":                    func() []interface{} { return nil }(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceLoadBalancerAccessSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceLoadBalancerAccessSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
