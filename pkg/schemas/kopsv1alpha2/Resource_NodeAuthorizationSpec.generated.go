package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceNodeAuthorizationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_authorizer": OptionalStruct(ResourceNodeAuthorizerSpec()),
		},
	}

	return res
}

func ExpandResourceNodeAuthorizationSpec(in map[string]interface{}) kopsv1alpha2.NodeAuthorizationSpec {
	if in == nil {
		panic("expand NodeAuthorizationSpec failure, in is nil")
	}
	return kopsv1alpha2.NodeAuthorizationSpec{
		NodeAuthorizer: func(in interface{}) *kopsv1alpha2.NodeAuthorizerSpec {
			return func(in interface{}) *kopsv1alpha2.NodeAuthorizerSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.NodeAuthorizerSpec) *kopsv1alpha2.NodeAuthorizerSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.NodeAuthorizerSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceNodeAuthorizerSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.NodeAuthorizerSpec{}
				}(in))
			}(in)
		}(in["node_authorizer"]),
	}
}

func FlattenResourceNodeAuthorizationSpecInto(in kopsv1alpha2.NodeAuthorizationSpec, out map[string]interface{}) {
	out["node_authorizer"] = func(in *kopsv1alpha2.NodeAuthorizerSpec) interface{} {
		return func(in *kopsv1alpha2.NodeAuthorizerSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.NodeAuthorizerSpec) interface{} {
				return func(in kopsv1alpha2.NodeAuthorizerSpec) []interface{} {
					return []interface{}{FlattenResourceNodeAuthorizerSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NodeAuthorizer)
}

func FlattenResourceNodeAuthorizationSpec(in kopsv1alpha2.NodeAuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceNodeAuthorizationSpecInto(in, out)
	return out
}
