package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceAuthorizationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"always_allow": Optional(Ptr(Struct(ResourceAlwaysAllowAuthorizationSpec()))),
			"rbac":         Optional(Ptr(Struct(ResourceRBACAuthorizationSpec()))),
		},
	}
}

func ExpandResourceAuthorizationSpec(in map[string]interface{}) kops.AuthorizationSpec {
	if in == nil {
		panic("expand AuthorizationSpec failure, in is nil")
	}
	out := kops.AuthorizationSpec{}
	if in, ok := in["always_allow"]; ok && in != nil {
		out.AlwaysAllow = func(in interface{}) *kops.AlwaysAllowAuthorizationSpec {
			if in == nil {
				return nil
			}
			return func(in kops.AlwaysAllowAuthorizationSpec) *kops.AlwaysAllowAuthorizationSpec { return &in }(func(in interface{}) kops.AlwaysAllowAuthorizationSpec {
				if in == nil {
					return kops.AlwaysAllowAuthorizationSpec{}
				}
				return ExpandResourceAlwaysAllowAuthorizationSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["rbac"]; ok && in != nil {
		out.RBAC = func(in interface{}) *kops.RBACAuthorizationSpec {
			if in == nil {
				return nil
			}
			return func(in kops.RBACAuthorizationSpec) *kops.RBACAuthorizationSpec { return &in }(func(in interface{}) kops.RBACAuthorizationSpec {
				if in == nil {
					return kops.RBACAuthorizationSpec{}
				}
				return ExpandResourceRBACAuthorizationSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	return out
}
