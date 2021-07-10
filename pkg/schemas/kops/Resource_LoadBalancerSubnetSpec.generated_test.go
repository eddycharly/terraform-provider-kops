package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceLoadBalancerSubnetSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.LoadBalancerSubnetSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceLoadBalancerSubnetSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceLoadBalancerSubnetSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceLoadBalancerSubnetSpecInto(t *testing.T) {
	type args struct {
		in  kops.LoadBalancerSubnetSpec
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
			FlattenResourceLoadBalancerSubnetSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceLoadBalancerSubnetSpec(t *testing.T) {
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
			want: map[string]interface{}{
				"name":                 "",
				"private_ipv4_address": nil,
				"allocation_id":        nil,
			},
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
			want: map[string]interface{}{
				"name":                 "",
				"private_ipv4_address": nil,
				"allocation_id":        nil,
			},
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
			want: map[string]interface{}{
				"name":                 "",
				"private_ipv4_address": nil,
				"allocation_id":        nil,
			},
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
			want: map[string]interface{}{
				"name":                 "",
				"private_ipv4_address": nil,
				"allocation_id":        nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceLoadBalancerSubnetSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceLoadBalancerSubnetSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
