package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandKuberouterNetworkingSpec(in map[string]interface{}) kops.KuberouterNetworkingSpec {
	if in == nil {
		panic("expand KuberouterNetworkingSpec failure, in is nil")
	}
	return kops.KuberouterNetworkingSpec{}
}

func FlattenKuberouterNetworkingSpecInto(in kops.KuberouterNetworkingSpec, out map[string]interface{}) {
}

func FlattenKuberouterNetworkingSpec(in kops.KuberouterNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenKuberouterNetworkingSpecInto(in, out)
	return out
}
