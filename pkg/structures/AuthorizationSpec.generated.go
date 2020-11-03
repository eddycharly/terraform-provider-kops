package structures

import (
	"reflect"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandAuthorizationSpec(in map[string]interface{}) kops.AuthorizationSpec {
	if in == nil {
		panic("expand AuthorizationSpec failure, in is nil")
	}
	return kops.AuthorizationSpec{
		AlwaysAllow: func(in interface{}) *kops.AlwaysAllowAuthorizationSpec {
			value := func(in interface{}) *kops.AlwaysAllowAuthorizationSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.AlwaysAllowAuthorizationSpec) *kops.AlwaysAllowAuthorizationSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.AlwaysAllowAuthorizationSpec {
					if in.([]interface{})[0] == nil {
						return kops.AlwaysAllowAuthorizationSpec{}
					}
					return (ExpandAlwaysAllowAuthorizationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["always_allow"]),
		RBAC: func(in interface{}) *kops.RBACAuthorizationSpec {
			value := func(in interface{}) *kops.RBACAuthorizationSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.RBACAuthorizationSpec) *kops.RBACAuthorizationSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.RBACAuthorizationSpec {
					if in.([]interface{})[0] == nil {
						return kops.RBACAuthorizationSpec{}
					}
					return (ExpandRBACAuthorizationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["rbac"]),
	}
}

func FlattenAuthorizationSpec(in kops.AuthorizationSpec) map[string]interface{} {
	return map[string]interface{}{
		"always_allow": func(in *kops.AlwaysAllowAuthorizationSpec) interface{} {
			value := func(in *kops.AlwaysAllowAuthorizationSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.AlwaysAllowAuthorizationSpec) interface{} {
					return func(in kops.AlwaysAllowAuthorizationSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenAlwaysAllowAuthorizationSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.AlwaysAllow),
		"rbac": func(in *kops.RBACAuthorizationSpec) interface{} {
			value := func(in *kops.RBACAuthorizationSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.RBACAuthorizationSpec) interface{} {
					return func(in kops.RBACAuthorizationSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenRBACAuthorizationSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.RBAC),
	}
}
