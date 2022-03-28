package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceOpenstackLoadbalancerConfig(t *testing.T) {
	_default := kops.OpenstackLoadbalancerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.OpenstackLoadbalancerConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"method":                  nil,
					"provider":                nil,
					"use_octavia":             nil,
					"floating_network":        nil,
					"floating_network_id":     nil,
					"floating_subnet":         nil,
					"subnet_id":               nil,
					"manage_sec_groups":       nil,
					"enable_ingress_hostname": nil,
					"ingress_hostname_suffix": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceOpenstackLoadbalancerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceOpenstackLoadbalancerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceOpenstackLoadbalancerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"method":                  nil,
		"provider":                nil,
		"use_octavia":             nil,
		"floating_network":        nil,
		"floating_network_id":     nil,
		"floating_subnet":         nil,
		"subnet_id":               nil,
		"manage_sec_groups":       nil,
		"enable_ingress_hostname": nil,
		"ingress_hostname_suffix": nil,
	}
	type args struct {
		in kops.OpenstackLoadbalancerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackLoadbalancerConfig{},
			},
			want: _default,
		},
		{
			name: "Method - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.Method = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Provider - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.Provider = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseOctavia - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.UseOctavia = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FloatingNetwork - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.FloatingNetwork = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FloatingNetworkID - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.FloatingNetworkID = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FloatingSubnet - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.FloatingSubnet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SubnetID - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.SubnetID = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ManageSecGroups - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.ManageSecGroups = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableIngressHostname - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.EnableIngressHostname = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IngressHostnameSuffix - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.IngressHostnameSuffix = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceOpenstackLoadbalancerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceOpenstackLoadbalancerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceOpenstackLoadbalancerConfig(t *testing.T) {
	_default := map[string]interface{}{
		"method":                  nil,
		"provider":                nil,
		"use_octavia":             nil,
		"floating_network":        nil,
		"floating_network_id":     nil,
		"floating_subnet":         nil,
		"subnet_id":               nil,
		"manage_sec_groups":       nil,
		"enable_ingress_hostname": nil,
		"ingress_hostname_suffix": nil,
	}
	type args struct {
		in kops.OpenstackLoadbalancerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackLoadbalancerConfig{},
			},
			want: _default,
		},
		{
			name: "Method - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.Method = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Provider - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.Provider = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseOctavia - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.UseOctavia = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FloatingNetwork - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.FloatingNetwork = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FloatingNetworkID - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.FloatingNetworkID = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FloatingSubnet - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.FloatingSubnet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SubnetID - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.SubnetID = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ManageSecGroups - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.ManageSecGroups = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableIngressHostname - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.EnableIngressHostname = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IngressHostnameSuffix - default",
			args: args{
				in: func() kops.OpenstackLoadbalancerConfig {
					subject := kops.OpenstackLoadbalancerConfig{}
					subject.IngressHostnameSuffix = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceOpenstackLoadbalancerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceOpenstackLoadbalancerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
