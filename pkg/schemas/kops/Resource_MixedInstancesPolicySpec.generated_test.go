package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceMixedInstancesPolicySpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.MixedInstancesPolicySpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceMixedInstancesPolicySpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceMixedInstancesPolicySpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceMixedInstancesPolicySpecInto(t *testing.T) {
	type args struct {
		in  kops.MixedInstancesPolicySpec
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
			FlattenResourceMixedInstancesPolicySpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceMixedInstancesPolicySpec(t *testing.T) {
	type args struct {
		in kops.MixedInstancesPolicySpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.MixedInstancesPolicySpec{},
			},
			want: map[string]interface{}{
				"instances":                     func() []interface{} { return nil }(),
				"on_demand_allocation_strategy": nil,
				"on_demand_base":                nil,
				"on_demand_above_base":          nil,
				"spot_allocation_strategy":      nil,
				"spot_instance_pools":           nil,
			},
		},
		{
			name: "Instances - default",
			args: args{
				in: func() kops.MixedInstancesPolicySpec {
					subject := kops.MixedInstancesPolicySpec{}
					subject.Instances = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"instances":                     func() []interface{} { return nil }(),
				"on_demand_allocation_strategy": nil,
				"on_demand_base":                nil,
				"on_demand_above_base":          nil,
				"spot_allocation_strategy":      nil,
				"spot_instance_pools":           nil,
			},
		},
		{
			name: "OnDemandAllocationStrategy - default",
			args: args{
				in: func() kops.MixedInstancesPolicySpec {
					subject := kops.MixedInstancesPolicySpec{}
					subject.OnDemandAllocationStrategy = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"instances":                     func() []interface{} { return nil }(),
				"on_demand_allocation_strategy": nil,
				"on_demand_base":                nil,
				"on_demand_above_base":          nil,
				"spot_allocation_strategy":      nil,
				"spot_instance_pools":           nil,
			},
		},
		{
			name: "OnDemandBase - default",
			args: args{
				in: func() kops.MixedInstancesPolicySpec {
					subject := kops.MixedInstancesPolicySpec{}
					subject.OnDemandBase = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"instances":                     func() []interface{} { return nil }(),
				"on_demand_allocation_strategy": nil,
				"on_demand_base":                nil,
				"on_demand_above_base":          nil,
				"spot_allocation_strategy":      nil,
				"spot_instance_pools":           nil,
			},
		},
		{
			name: "OnDemandAboveBase - default",
			args: args{
				in: func() kops.MixedInstancesPolicySpec {
					subject := kops.MixedInstancesPolicySpec{}
					subject.OnDemandAboveBase = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"instances":                     func() []interface{} { return nil }(),
				"on_demand_allocation_strategy": nil,
				"on_demand_base":                nil,
				"on_demand_above_base":          nil,
				"spot_allocation_strategy":      nil,
				"spot_instance_pools":           nil,
			},
		},
		{
			name: "SpotAllocationStrategy - default",
			args: args{
				in: func() kops.MixedInstancesPolicySpec {
					subject := kops.MixedInstancesPolicySpec{}
					subject.SpotAllocationStrategy = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"instances":                     func() []interface{} { return nil }(),
				"on_demand_allocation_strategy": nil,
				"on_demand_base":                nil,
				"on_demand_above_base":          nil,
				"spot_allocation_strategy":      nil,
				"spot_instance_pools":           nil,
			},
		},
		{
			name: "SpotInstancePools - default",
			args: args{
				in: func() kops.MixedInstancesPolicySpec {
					subject := kops.MixedInstancesPolicySpec{}
					subject.SpotInstancePools = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"instances":                     func() []interface{} { return nil }(),
				"on_demand_allocation_strategy": nil,
				"on_demand_base":                nil,
				"on_demand_above_base":          nil,
				"spot_allocation_strategy":      nil,
				"spot_instance_pools":           nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceMixedInstancesPolicySpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceMixedInstancesPolicySpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
