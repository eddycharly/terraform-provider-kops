package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandAlwaysAllowAuthorizationSpec(in map[string]interface{}) kops.AlwaysAllowAuthorizationSpec {
	if in == nil {
		panic("expand AlwaysAllowAuthorizationSpec failure, in is nil")
	}
	return kops.AlwaysAllowAuthorizationSpec{}
}

func FlattenAlwaysAllowAuthorizationSpec(in kops.AlwaysAllowAuthorizationSpec) map[string]interface{} {
	return map[string]interface{}{}
}
