package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandExternalNetworkingSpec(in map[string]interface{}) kops.ExternalNetworkingSpec {
	if in == nil {
		panic("expand ExternalNetworkingSpec failure, in is nil")
	}
	return kops.ExternalNetworkingSpec{}
}

func FlattenExternalNetworkingSpec(in kops.ExternalNetworkingSpec) map[string]interface{} {
	return map[string]interface{}{}
}
