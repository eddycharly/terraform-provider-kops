package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceLoadBalancerSubnetSpec(t *testing.T) {
	_default := kops.LoadBalancerSubnetSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.LoadBalancerSubnetSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"name":                 "",
					"private_ipv4_address": nil,
					"allocation_id":        nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceLoadBalancerSubnetSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceLoadBalancerSubnetSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLoadBalancerSubnetSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"name":                 "",
		"private_ipv4_address": nil,
		"allocation_id":        nil,
	}
	type args struct {
		in kops.LoadBalancerSubnetSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.LoadBalancerSubnetSpec{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() kops.LoadBalancerSubnetSpec {
					subject := kops.LoadBalancerSubnetSpec{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrivateIpv4Address - default",
			args: args{
				in: func() kops.LoadBalancerSubnetSpec {
					subject := kops.LoadBalancerSubnetSpec{}
					subject.PrivateIPv4Address = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllocationID - default",
			args: args{
				in: func() kops.LoadBalancerSubnetSpec {
					subject := kops.LoadBalancerSubnetSpec{}
					subject.AllocationID = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceLoadBalancerSubnetSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLoadBalancerSubnetSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLoadBalancerSubnetSpec(t *testing.T) {
	_default := map[string]interface{}{
		"name":                 "",
		"private_ipv4_address": nil,
		"allocation_id":        nil,
	}
	type args struct {
		in kops.LoadBalancerSubnetSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.LoadBalancerSubnetSpec{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() kops.LoadBalancerSubnetSpec {
					subject := kops.LoadBalancerSubnetSpec{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrivateIpv4Address - default",
			args: args{
				in: func() kops.LoadBalancerSubnetSpec {
					subject := kops.LoadBalancerSubnetSpec{}
					subject.PrivateIPv4Address = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllocationID - default",
			args: args{
				in: func() kops.LoadBalancerSubnetSpec {
					subject := kops.LoadBalancerSubnetSpec{}
					subject.AllocationID = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceLoadBalancerSubnetSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLoadBalancerSubnetSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
