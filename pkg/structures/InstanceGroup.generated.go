package structures

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandInstanceGroup(in map[string]interface{}) api.InstanceGroup {
	if in == nil {
		panic("expand InstanceGroup failure, in is nil")
	}
	return api.InstanceGroup{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		InstanceGroupSpec: func(in interface{}) kops.InstanceGroupSpec {
			return func(in interface{}) kops.InstanceGroupSpec {
				if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
					return kops.InstanceGroupSpec{}
				}
				return (ExpandInstanceGroupSpec(in.([]interface{})[0].(map[string]interface{})))
			}(in)
		}(in),
	}
}

func FlattenInstanceGroupInto(in api.InstanceGroup, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	FlattenInstanceGroupSpecInto(in.InstanceGroupSpec, out)
}

func FlattenInstanceGroup(in api.InstanceGroup) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenInstanceGroupInto(in, out)
	return out
}
