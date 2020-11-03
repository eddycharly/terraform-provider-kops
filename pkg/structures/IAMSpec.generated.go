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

func FlattenIAMSpec(in kops.IAMSpec) map[string]interface{} {
	return map[string]interface{}{
		"legacy": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.Legacy),
		"allow_container_registry": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.AllowContainerRegistry),
	}
}
