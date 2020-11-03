package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandKopeioAuthenticationSpec(in map[string]interface{}) kops.KopeioAuthenticationSpec {
	if in == nil {
		panic("expand KopeioAuthenticationSpec failure, in is nil")
	}
	return kops.KopeioAuthenticationSpec{}
}

func FlattenKopeioAuthenticationSpec(in kops.KopeioAuthenticationSpec) map[string]interface{} {
	return map[string]interface{}{}
}
