package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceOpenstackLoadbalancerConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.OpenstackLoadbalancerConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceOpenstackLoadbalancerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceOpenstackLoadbalancerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackLoadbalancerConfigInto(t *testing.T) {
	type args struct {
		in  kops.OpenstackLoadbalancerConfig
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
			FlattenDataSourceOpenstackLoadbalancerConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceOpenstackLoadbalancerConfig(t *testing.T) {
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
			want: map[string]interface{}{
				"method":              nil,
				"provider":            nil,
				"use_octavia":         nil,
				"floating_network":    nil,
				"floating_network_id": nil,
				"floating_subnet":     nil,
				"subnet_id":           nil,
				"manage_sec_groups":   nil,
			},
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
			want: map[string]interface{}{
				"method":              nil,
				"provider":            nil,
				"use_octavia":         nil,
				"floating_network":    nil,
				"floating_network_id": nil,
				"floating_subnet":     nil,
				"subnet_id":           nil,
				"manage_sec_groups":   nil,
			},
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
			want: map[string]interface{}{
				"method":              nil,
				"provider":            nil,
				"use_octavia":         nil,
				"floating_network":    nil,
				"floating_network_id": nil,
				"floating_subnet":     nil,
				"subnet_id":           nil,
				"manage_sec_groups":   nil,
			},
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
			want: map[string]interface{}{
				"method":              nil,
				"provider":            nil,
				"use_octavia":         nil,
				"floating_network":    nil,
				"floating_network_id": nil,
				"floating_subnet":     nil,
				"subnet_id":           nil,
				"manage_sec_groups":   nil,
			},
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
			want: map[string]interface{}{
				"method":              nil,
				"provider":            nil,
				"use_octavia":         nil,
				"floating_network":    nil,
				"floating_network_id": nil,
				"floating_subnet":     nil,
				"subnet_id":           nil,
				"manage_sec_groups":   nil,
			},
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
			want: map[string]interface{}{
				"method":              nil,
				"provider":            nil,
				"use_octavia":         nil,
				"floating_network":    nil,
				"floating_network_id": nil,
				"floating_subnet":     nil,
				"subnet_id":           nil,
				"manage_sec_groups":   nil,
			},
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
			want: map[string]interface{}{
				"method":              nil,
				"provider":            nil,
				"use_octavia":         nil,
				"floating_network":    nil,
				"floating_network_id": nil,
				"floating_subnet":     nil,
				"subnet_id":           nil,
				"manage_sec_groups":   nil,
			},
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
			want: map[string]interface{}{
				"method":              nil,
				"provider":            nil,
				"use_octavia":         nil,
				"floating_network":    nil,
				"floating_network_id": nil,
				"floating_subnet":     nil,
				"subnet_id":           nil,
				"manage_sec_groups":   nil,
			},
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
			want: map[string]interface{}{
				"method":              nil,
				"provider":            nil,
				"use_octavia":         nil,
				"floating_network":    nil,
				"floating_network_id": nil,
				"floating_subnet":     nil,
				"subnet_id":           nil,
				"manage_sec_groups":   nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceOpenstackLoadbalancerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceOpenstackLoadbalancerConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
