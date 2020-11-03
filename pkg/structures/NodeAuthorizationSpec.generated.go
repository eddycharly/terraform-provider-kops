package structures

import (
	"reflect"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandNodeAuthorizationSpec(in map[string]interface{}) kops.NodeAuthorizationSpec {
	if in == nil {
		panic("expand NodeAuthorizationSpec failure, in is nil")
	}
	return kops.NodeAuthorizationSpec{
		NodeAuthorizer: func(in interface{}) *kops.NodeAuthorizerSpec {
			value := func(in interface{}) *kops.NodeAuthorizerSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.NodeAuthorizerSpec) *kops.NodeAuthorizerSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.NodeAuthorizerSpec {
					if in.([]interface{})[0] == nil {
						return kops.NodeAuthorizerSpec{}
					}
					return (ExpandNodeAuthorizerSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["node_authorizer"]),
	}
}

func FlattenNodeAuthorizationSpec(in kops.NodeAuthorizationSpec) map[string]interface{} {
	return map[string]interface{}{
		"node_authorizer": func(in *kops.NodeAuthorizerSpec) interface{} {
			value := func(in *kops.NodeAuthorizerSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.NodeAuthorizerSpec) interface{} {
					return func(in kops.NodeAuthorizerSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenNodeAuthorizerSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.NodeAuthorizer),
	}
}
