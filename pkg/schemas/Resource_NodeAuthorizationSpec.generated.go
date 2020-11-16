package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceNodeAuthorizationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_authorizer": OptionalStruct(ResourceNodeAuthorizerSpec()),
		},
	}
}

func ExpandResourceNodeAuthorizationSpec(in map[string]interface{}) kops.NodeAuthorizationSpec {
	if in == nil {
		panic("expand NodeAuthorizationSpec failure, in is nil")
	}
	return kops.NodeAuthorizationSpec{
		NodeAuthorizer: func(in interface{}) *kops.NodeAuthorizerSpec {
			return func(in interface{}) *kops.NodeAuthorizerSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.NodeAuthorizerSpec) *kops.NodeAuthorizerSpec {
					return &in
				}(func(in interface{}) kops.NodeAuthorizerSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.NodeAuthorizerSpec{}
					}
					return (ExpandResourceNodeAuthorizerSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["node_authorizer"]),
	}
}

func FlattenResourceNodeAuthorizationSpecInto(in kops.NodeAuthorizationSpec, out map[string]interface{}) {
	out["node_authorizer"] = func(in *kops.NodeAuthorizerSpec) interface{} {
		return func(in *kops.NodeAuthorizerSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.NodeAuthorizerSpec) interface{} {
				return func(in kops.NodeAuthorizerSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceNodeAuthorizerSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NodeAuthorizer)
}

func FlattenResourceNodeAuthorizationSpec(in kops.NodeAuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceNodeAuthorizationSpecInto(in, out)
	return out
}
