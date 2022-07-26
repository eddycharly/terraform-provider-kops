package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceInstanceRequirementsSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu":    OptionalStruct(ResourceMinMaxSpec()),
			"memory": OptionalStruct(ResourceMinMaxSpec()),
		},
	}

	return res
}

func ExpandResourceInstanceRequirementsSpec(in map[string]interface{}) kops.InstanceRequirementsSpec {
	if in == nil {
		panic("expand InstanceRequirementsSpec failure, in is nil")
	}
	return kops.InstanceRequirementsSpec{
		CPU: func(in interface{}) *kops.MinMaxSpec {
			return func(in interface{}) *kops.MinMaxSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.MinMaxSpec) *kops.MinMaxSpec {
					return &in
				}(func(in interface{}) kops.MinMaxSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.MinMaxSpec{}
					}
					return (ExpandResourceMinMaxSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["cpu"]),
		Memory: func(in interface{}) *kops.MinMaxSpec {
			return func(in interface{}) *kops.MinMaxSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.MinMaxSpec) *kops.MinMaxSpec {
					return &in
				}(func(in interface{}) kops.MinMaxSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.MinMaxSpec{}
					}
					return (ExpandResourceMinMaxSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["memory"]),
	}
}

func FlattenResourceInstanceRequirementsSpecInto(in kops.InstanceRequirementsSpec, out map[string]interface{}) {
	out["cpu"] = func(in *kops.MinMaxSpec) interface{} {
		return func(in *kops.MinMaxSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.MinMaxSpec) interface{} {
				return func(in kops.MinMaxSpec) []interface{} {
					return []interface{}{FlattenResourceMinMaxSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.CPU)
	out["memory"] = func(in *kops.MinMaxSpec) interface{} {
		return func(in *kops.MinMaxSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.MinMaxSpec) interface{} {
				return func(in kops.MinMaxSpec) []interface{} {
					return []interface{}{FlattenResourceMinMaxSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Memory)
}

func FlattenResourceInstanceRequirementsSpec(in kops.InstanceRequirementsSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceInstanceRequirementsSpecInto(in, out)
	return out
}
