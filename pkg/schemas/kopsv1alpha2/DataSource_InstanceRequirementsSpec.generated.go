package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceInstanceRequirementsSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu":    ComputedStruct(DataSourceMinMaxSpec()),
			"memory": ComputedStruct(DataSourceMinMaxSpec()),
		},
	}

	return res
}

func ExpandDataSourceInstanceRequirementsSpec(in map[string]interface{}) kopsv1alpha2.InstanceRequirementsSpec {
	if in == nil {
		panic("expand InstanceRequirementsSpec failure, in is nil")
	}
	return kopsv1alpha2.InstanceRequirementsSpec{
		CPU: func(in interface{}) *kopsv1alpha2.MinMaxSpec {
			return func(in interface{}) *kopsv1alpha2.MinMaxSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.MinMaxSpec) *kopsv1alpha2.MinMaxSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.MinMaxSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceMinMaxSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.MinMaxSpec{}
				}(in))
			}(in)
		}(in["cpu"]),
		Memory: func(in interface{}) *kopsv1alpha2.MinMaxSpec {
			return func(in interface{}) *kopsv1alpha2.MinMaxSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.MinMaxSpec) *kopsv1alpha2.MinMaxSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.MinMaxSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceMinMaxSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.MinMaxSpec{}
				}(in))
			}(in)
		}(in["memory"]),
	}
}

func FlattenDataSourceInstanceRequirementsSpecInto(in kopsv1alpha2.InstanceRequirementsSpec, out map[string]interface{}) {
	out["cpu"] = func(in *kopsv1alpha2.MinMaxSpec) interface{} {
		return func(in *kopsv1alpha2.MinMaxSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.MinMaxSpec) interface{} {
				return func(in kopsv1alpha2.MinMaxSpec) []interface{} {
					return []interface{}{FlattenDataSourceMinMaxSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.CPU)
	out["memory"] = func(in *kopsv1alpha2.MinMaxSpec) interface{} {
		return func(in *kopsv1alpha2.MinMaxSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.MinMaxSpec) interface{} {
				return func(in kopsv1alpha2.MinMaxSpec) []interface{} {
					return []interface{}{FlattenDataSourceMinMaxSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Memory)
}

func FlattenDataSourceInstanceRequirementsSpec(in kopsv1alpha2.InstanceRequirementsSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceInstanceRequirementsSpecInto(in, out)
	return out
}
