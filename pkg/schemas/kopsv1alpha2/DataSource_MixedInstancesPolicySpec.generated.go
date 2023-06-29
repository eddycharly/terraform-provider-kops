package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceMixedInstancesPolicySpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"instances":                     ComputedList(String()),
			"instance_requirements":         ComputedStruct(DataSourceInstanceRequirementsSpec()),
			"on_demand_allocation_strategy": ComputedString(),
			"on_demand_base":                Nullable(ComputedInt()),
			"on_demand_above_base":          Nullable(ComputedInt()),
			"spot_allocation_strategy":      ComputedString(),
			"spot_instance_pools":           ComputedInt(),
		},
	}

	return res
}

func ExpandDataSourceMixedInstancesPolicySpec(in map[string]interface{}) kopsv1alpha2.MixedInstancesPolicySpec {
	if in == nil {
		panic("expand MixedInstancesPolicySpec failure, in is nil")
	}
	return kopsv1alpha2.MixedInstancesPolicySpec{
		Instances: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["instances"]),
		InstanceRequirements: func(in interface{}) *kopsv1alpha2.InstanceRequirementsSpec {
			return func(in interface{}) *kopsv1alpha2.InstanceRequirementsSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.InstanceRequirementsSpec) *kopsv1alpha2.InstanceRequirementsSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.InstanceRequirementsSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceInstanceRequirementsSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.InstanceRequirementsSpec{}
				}(in))
			}(in)
		}(in["instance_requirements"]),
		OnDemandAllocationStrategy: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["on_demand_allocation_strategy"]),
		OnDemandBase: func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && len(in) == 1 {
				return func(in interface{}) *int64 {
					return func(in interface{}) *int64 {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
							return nil
						}
						return func(in int64) *int64 {
							return &in
						}(int64(ExpandInt(in)))
					}(in)
				}(in[0].(map[string]interface{})["value"])
			}
			return nil
		}(in["on_demand_base"]),
		OnDemandAboveBase: func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && len(in) == 1 {
				return func(in interface{}) *int64 {
					return func(in interface{}) *int64 {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
							return nil
						}
						return func(in int64) *int64 {
							return &in
						}(int64(ExpandInt(in)))
					}(in)
				}(in[0].(map[string]interface{})["value"])
			}
			return nil
		}(in["on_demand_above_base"]),
		SpotAllocationStrategy: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["spot_allocation_strategy"]),
		SpotInstancePools: func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					return &in
				}(int64(ExpandInt(in)))
			}(in)
		}(in["spot_instance_pools"]),
	}
}

func FlattenDataSourceMixedInstancesPolicySpecInto(in kopsv1alpha2.MixedInstancesPolicySpec, out map[string]interface{}) {
	out["instances"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Instances)
	out["instance_requirements"] = func(in *kopsv1alpha2.InstanceRequirementsSpec) interface{} {
		return func(in *kopsv1alpha2.InstanceRequirementsSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.InstanceRequirementsSpec) interface{} {
				return func(in kopsv1alpha2.InstanceRequirementsSpec) []interface{} {
					return []interface{}{FlattenDataSourceInstanceRequirementsSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.InstanceRequirements)
	out["on_demand_allocation_strategy"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.OnDemandAllocationStrategy)
	out["on_demand_base"] = func(in *int64) interface{} {
		if in == nil {
			return nil
		}
		return []interface{}{map[string]interface{}{"value": func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)}}
	}(in.OnDemandBase)
	out["on_demand_above_base"] = func(in *int64) interface{} {
		if in == nil {
			return nil
		}
		return []interface{}{map[string]interface{}{"value": func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)}}
	}(in.OnDemandAboveBase)
	out["spot_allocation_strategy"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.SpotAllocationStrategy)
	out["spot_instance_pools"] = func(in *int64) interface{} {
		return func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.SpotInstancePools)
}

func FlattenDataSourceMixedInstancesPolicySpec(in kopsv1alpha2.MixedInstancesPolicySpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceMixedInstancesPolicySpecInto(in, out)
	return out
}
