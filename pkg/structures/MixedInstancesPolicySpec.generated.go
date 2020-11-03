package structures

import (
	"reflect"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandMixedInstancesPolicySpec(in map[string]interface{}) kops.MixedInstancesPolicySpec {
	if in == nil {
		panic("expand MixedInstancesPolicySpec failure, in is nil")
	}
	return kops.MixedInstancesPolicySpec{
		Instances: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["instances"]),
		OnDemandAllocationStrategy: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["on_demand_allocation_strategy"]),
		OnDemandBase: func(in interface{}) *int64 {
			value := func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int64(ExpandInt(in)))
			}(in)
			return value
		}(in["on_demand_base"]),
		OnDemandAboveBase: func(in interface{}) *int64 {
			value := func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int64(ExpandInt(in)))
			}(in)
			return value
		}(in["on_demand_above_base"]),
		SpotAllocationStrategy: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["spot_allocation_strategy"]),
		SpotInstancePools: func(in interface{}) *int64 {
			value := func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int64(ExpandInt(in)))
			}(in)
			return value
		}(in["spot_instance_pools"]),
	}
}

func FlattenMixedInstancesPolicySpec(in kops.MixedInstancesPolicySpec) map[string]interface{} {
	return map[string]interface{}{
		"instances": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.Instances),
		"on_demand_allocation_strategy": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.OnDemandAllocationStrategy),
		"on_demand_base": func(in *int64) interface{} {
			value := func(in *int64) interface{} {
				if in == nil {
					return nil
				}
				return func(in int64) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.OnDemandBase),
		"on_demand_above_base": func(in *int64) interface{} {
			value := func(in *int64) interface{} {
				if in == nil {
					return nil
				}
				return func(in int64) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.OnDemandAboveBase),
		"spot_allocation_strategy": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.SpotAllocationStrategy),
		"spot_instance_pools": func(in *int64) interface{} {
			value := func(in *int64) interface{} {
				if in == nil {
					return nil
				}
				return func(in int64) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.SpotInstancePools),
	}
}
