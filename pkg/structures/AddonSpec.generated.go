package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandAddonSpec(in map[string]interface{}) kops.AddonSpec {
	if in == nil {
		panic("expand AddonSpec failure, in is nil")
	}
	return kops.AddonSpec{
		Manifest: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["manifest"]),
	}
}

func FlattenAddonSpec(in kops.AddonSpec) map[string]interface{} {
	return map[string]interface{}{
		"manifest": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Manifest),
	}
}
