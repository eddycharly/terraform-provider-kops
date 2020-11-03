package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandRBACAuthorizationSpec(in map[string]interface{}) kops.RBACAuthorizationSpec {
	if in == nil {
		panic("expand RBACAuthorizationSpec failure, in is nil")
	}
	return kops.RBACAuthorizationSpec{}
}

func FlattenRBACAuthorizationSpec(in kops.RBACAuthorizationSpec) map[string]interface{} {
	return map[string]interface{}{}
}
