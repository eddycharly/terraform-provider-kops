package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandKopeioNetworkingSpec(in map[string]interface{}) kops.KopeioNetworkingSpec {
	if in == nil {
		panic("expand KopeioNetworkingSpec failure, in is nil")
	}
	return kops.KopeioNetworkingSpec{}
}

func FlattenKopeioNetworkingSpec(in kops.KopeioNetworkingSpec) map[string]interface{} {
	return map[string]interface{}{}
}
