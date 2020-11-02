package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandClassicNetworkingSpec(in map[string]interface{}) kops.ClassicNetworkingSpec {
	if in == nil {
		panic("expand ClassicNetworkingSpec failure, in is nil")
	}
	return kops.ClassicNetworkingSpec{}
}

func FlattenClassicNetworkingSpec(in kops.ClassicNetworkingSpec) map[string]interface{} {
	return map[string]interface{}{}
}
