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
	_default := map[string]interface{}{
		"instances":                     func() []interface{} { return nil }(),
		"on_demand_allocation_strategy": nil,
		"on_demand_base":                nil,
		"on_demand_above_base":          nil,
		"spot_allocation_strategy":      nil,
		"spot_instance_pools":           nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceMixedInstancesPolicySpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceMixedInstancesPolicySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceMixedInstancesPolicySpec(t *testing.T) {
	_default := map[string]interface{}{
		"instances":                     func() []interface{} { return nil }(),
		"on_demand_allocation_strategy": nil,
		"on_demand_base":                nil,
		"on_demand_above_base":          nil,
		"spot_allocation_strategy":      nil,
		"spot_instance_pools":           nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceMixedInstancesPolicySpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceMixedInstancesPolicySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
