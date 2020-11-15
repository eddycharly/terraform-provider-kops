package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandIAMSpec(in map[string]interface{}) kops.IAMSpec {
	if in == nil {
		panic("expand IAMSpec failure, in is nil")
	}
	return kops.IAMSpec{
		Legacy: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["legacy"]),
		AllowContainerRegistry: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["allow_container_registry"]),
	}
}

func FlattenIAMSpecInto(in kops.IAMSpec, out map[string]interface{}) {
	out["legacy"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Legacy)
	out["allow_container_registry"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.AllowContainerRegistry)
}

func FlattenIAMSpec(in kops.IAMSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenIAMSpecInto(in, out)
	return out
}
