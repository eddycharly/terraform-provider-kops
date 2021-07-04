package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceNodeAuthorizationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_authorizer": Computed(Struct(DataSourceNodeAuthorizerSpec())),
		},
	}
}

func ExpandDataSourceNodeAuthorizationSpec(in map[string]interface{}) kops.NodeAuthorizationSpec {
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
				return ExpandDataSourceNodeAuthorizerSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	return out
}

func FlattenDataSourceNodeAuthorizationSpecInto(in kops.NodeAuthorizationSpec, out map[string]interface{}) {
	out["node_authorizer"] = func(in *kops.NodeAuthorizerSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.NodeAuthorizerSpec) interface{} { return FlattenDataSourceNodeAuthorizerSpec(in) }(*in)
	}(in.NodeAuthorizer)
}

func FlattenDataSourceNodeAuthorizationSpec(in kops.NodeAuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceNodeAuthorizationSpecInto(in, out)
	return out
}
