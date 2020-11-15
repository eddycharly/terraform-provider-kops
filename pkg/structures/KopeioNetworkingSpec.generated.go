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

func FlattenKopeioNetworkingSpecInto(in kops.KopeioNetworkingSpec, out map[string]interface{}) {
}

func FlattenKopeioNetworkingSpec(in kops.KopeioNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenKopeioNetworkingSpecInto(in, out)
	return out
}
