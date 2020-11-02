package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandKubenetNetworkingSpec(in map[string]interface{}) kops.KubenetNetworkingSpec {
	if in == nil {
		panic("expand KubenetNetworkingSpec failure, in is nil")
	}
	return kops.KubenetNetworkingSpec{}
}

func FlattenKubenetNetworkingSpec(in kops.KubenetNetworkingSpec) map[string]interface{} {
	return map[string]interface{}{}
}
