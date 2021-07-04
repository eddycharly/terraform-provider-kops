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
			"node_authorizer": Optional(Struct(ResourceNodeAuthorizerSpec())),
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.NodeAuthorizerSpec) *kops.NodeAuthorizerSpec { return &in }(func(in interface{}) kops.NodeAuthorizerSpec {
					return ExpandResourceNodeAuthorizerSpec(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenResourceNodeAuthorizationSpecInto(in kops.NodeAuthorizationSpec, out map[string]interface{}) {
	out["node_authorizer"] = func(in *kops.NodeAuthorizerSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.NodeAuthorizerSpec) interface{} { return FlattenResourceNodeAuthorizerSpec(in) }(*in)
	}(in.NodeAuthorizer)
}

func FlattenResourceNodeAuthorizationSpec(in kops.NodeAuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceNodeAuthorizationSpecInto(in, out)
	return out
}
