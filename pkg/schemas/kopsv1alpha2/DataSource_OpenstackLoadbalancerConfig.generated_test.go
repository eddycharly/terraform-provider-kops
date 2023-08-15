package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceOpenstackLoadbalancerConfig(t *testing.T) {
	_default := kopsv1alpha2.OpenstackLoadbalancerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.OpenstackLoadbalancerConfig
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
			got := ExpandDataSourceOpenstackLoadbalancerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceOpenstackLoadbalancerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackLoadbalancerConfigInto(t *testing.T) {
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
		in kopsv1alpha2.OpenstackLoadbalancerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.OpenstackLoadbalancerConfig{},
			},
			want: _default,
		},
		{
			name: "Method - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.Method = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Provider - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.Provider = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseOctavia - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.UseOctavia = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FloatingNetwork - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.FloatingNetwork = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FloatingNetworkId - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.FloatingNetworkID = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FloatingSubnet - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.FloatingSubnet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SubnetId - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.SubnetID = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ManageSecGroups - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.ManageSecGroups = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableIngressHostname - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.EnableIngressHostname = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IngressHostnameSuffix - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
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
			FlattenDataSourceOpenstackLoadbalancerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceOpenstackLoadbalancerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackLoadbalancerConfig(t *testing.T) {
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
		in kopsv1alpha2.OpenstackLoadbalancerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.OpenstackLoadbalancerConfig{},
			},
			want: _default,
		},
		{
			name: "Method - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.Method = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Provider - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.Provider = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseOctavia - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.UseOctavia = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FloatingNetwork - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.FloatingNetwork = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FloatingNetworkId - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.FloatingNetworkID = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FloatingSubnet - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.FloatingSubnet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SubnetId - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.SubnetID = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ManageSecGroups - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.ManageSecGroups = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableIngressHostname - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.EnableIngressHostname = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IngressHostnameSuffix - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackLoadbalancerConfig {
					subject := kopsv1alpha2.OpenstackLoadbalancerConfig{}
					subject.IngressHostnameSuffix = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceOpenstackLoadbalancerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceOpenstackLoadbalancerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
