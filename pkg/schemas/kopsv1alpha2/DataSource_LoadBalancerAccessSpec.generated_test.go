package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceLoadBalancerAccessSpec(t *testing.T) {
	_default := kopsv1alpha2.LoadBalancerAccessSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.LoadBalancerAccessSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
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
					"access_log":                 nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceLoadBalancerAccessSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceLoadBalancerAccessSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLoadBalancerAccessSpecInto(t *testing.T) {
	_default := map[string]interface{}{
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
		"access_log":                 nil,
	}
	type args struct {
		in kopsv1alpha2.LoadBalancerAccessSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.LoadBalancerAccessSpec{},
			},
			want: _default,
		},
		{
			name: "Class - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.Class = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.Type = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdleTimeoutSeconds - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.IdleTimeoutSeconds = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecurityGroupOverride - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.SecurityGroupOverride = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalSecurityGroups - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.AdditionalSecurityGroups = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseForInternalApi - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.UseForInternalAPI = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SSLCertificate - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.SSLCertificate = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SSLPolicy - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.SSLPolicy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CrossZoneLoadBalancing - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.CrossZoneLoadBalancing = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Subnets - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.Subnets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AccessLog - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.AccessLog = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceLoadBalancerAccessSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLoadBalancerAccessSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLoadBalancerAccessSpec(t *testing.T) {
	_default := map[string]interface{}{
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
		"access_log":                 nil,
	}
	type args struct {
		in kopsv1alpha2.LoadBalancerAccessSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.LoadBalancerAccessSpec{},
			},
			want: _default,
		},
		{
			name: "Class - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.Class = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.Type = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdleTimeoutSeconds - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.IdleTimeoutSeconds = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecurityGroupOverride - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.SecurityGroupOverride = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalSecurityGroups - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.AdditionalSecurityGroups = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseForInternalApi - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.UseForInternalAPI = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SSLCertificate - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.SSLCertificate = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SSLPolicy - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.SSLPolicy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CrossZoneLoadBalancing - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.CrossZoneLoadBalancing = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Subnets - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.Subnets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AccessLog - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerAccessSpec {
					subject := kopsv1alpha2.LoadBalancerAccessSpec{}
					subject.AccessLog = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceLoadBalancerAccessSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLoadBalancerAccessSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
