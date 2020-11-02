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

func FlattenGCENetworkingSpec(in kops.GCENetworkingSpec) map[string]interface{} {
	return map[string]interface{}{}
}
