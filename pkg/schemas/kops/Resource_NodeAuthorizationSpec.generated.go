package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceNodeAuthorizationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_authorizer": Optional(Ptr(Struct(ResourceNodeAuthorizerSpec()))),
		},
	}
}

func ExpandResourceNodeAuthorizationSpec(in map[string]interface{}) kops.NodeAuthorizationSpec {
	if in == nil {
		panic("expand NodeAuthorizationSpec failure, in is nil")
	}
	out := kops.NodeAuthorizationSpec{}
	if in, ok := in["node_authorizer"]; ok && in != nil {
		out.NodeAuthorizer = func(in interface{}) *kops.NodeAuthorizerSpec {
			if in == nil {
				return nil
			}
			return func(in kops.NodeAuthorizerSpec) *kops.NodeAuthorizerSpec { return &in }(func(in interface{}) kops.NodeAuthorizerSpec {
				if in == nil {
					return kops.NodeAuthorizerSpec{}
				}
				return ExpandResourceNodeAuthorizerSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	return out
}
