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
			"always_allow": Optional(Struct(ResourceAlwaysAllowAuthorizationSpec())),
			"rbac":         Optional(Struct(ResourceRBACAuthorizationSpec())),
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.AlwaysAllowAuthorizationSpec) *kops.AlwaysAllowAuthorizationSpec { return &in }(func(in interface{}) kops.AlwaysAllowAuthorizationSpec {
					return ExpandResourceAlwaysAllowAuthorizationSpec(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	if in, ok := in["rbac"]; ok && in != nil {
		out.RBAC = func(in interface{}) *kops.RBACAuthorizationSpec {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.RBACAuthorizationSpec) *kops.RBACAuthorizationSpec { return &in }(func(in interface{}) kops.RBACAuthorizationSpec {
					return ExpandResourceRBACAuthorizationSpec(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenResourceAuthorizationSpecInto(in kops.AuthorizationSpec, out map[string]interface{}) {
	out["always_allow"] = func(in *kops.AlwaysAllowAuthorizationSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.AlwaysAllowAuthorizationSpec) interface{} {
			return FlattenResourceAlwaysAllowAuthorizationSpec(in)
		}(*in)
	}(in.AlwaysAllow)
	out["rbac"] = func(in *kops.RBACAuthorizationSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.RBACAuthorizationSpec) interface{} { return FlattenResourceRBACAuthorizationSpec(in) }(*in)
	}(in.RBAC)
}

func FlattenResourceAuthorizationSpec(in kops.AuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAuthorizationSpecInto(in, out)
	return out
}
