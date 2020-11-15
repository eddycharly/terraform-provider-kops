package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandGCENetworkingSpec(in map[string]interface{}) kops.GCENetworkingSpec {
	if in == nil {
		panic("expand GCENetworkingSpec failure, in is nil")
	}
	return kops.GCENetworkingSpec{}
}

func FlattenGCENetworkingSpecInto(in kops.GCENetworkingSpec, out map[string]interface{}) {
}

func FlattenGCENetworkingSpec(in kops.GCENetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenGCENetworkingSpecInto(in, out)
	return out
}
