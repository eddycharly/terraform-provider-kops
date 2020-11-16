package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceAuthorizationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"always_allow": OptionalStruct(ResourceAlwaysAllowAuthorizationSpec()),
			"rbac":         OptionalStruct(ResourceRBACAuthorizationSpec()),
		},
	}
}

func ExpandResourceAuthorizationSpec(in map[string]interface{}) kops.AuthorizationSpec {
	if in == nil {
		panic("expand AuthorizationSpec failure, in is nil")
	}
	return kops.AuthorizationSpec{
		AlwaysAllow: func(in interface{}) *kops.AlwaysAllowAuthorizationSpec {
			return func(in interface{}) *kops.AlwaysAllowAuthorizationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AlwaysAllowAuthorizationSpec) *kops.AlwaysAllowAuthorizationSpec {
					return &in
				}(func(in interface{}) kops.AlwaysAllowAuthorizationSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.AlwaysAllowAuthorizationSpec{}
					}
					return (ExpandResourceAlwaysAllowAuthorizationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["always_allow"]),
		RBAC: func(in interface{}) *kops.RBACAuthorizationSpec {
			return func(in interface{}) *kops.RBACAuthorizationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.RBACAuthorizationSpec) *kops.RBACAuthorizationSpec {
					return &in
				}(func(in interface{}) kops.RBACAuthorizationSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.RBACAuthorizationSpec{}
					}
					return (ExpandResourceRBACAuthorizationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["rbac"]),
	}
}

func FlattenResourceAuthorizationSpecInto(in kops.AuthorizationSpec, out map[string]interface{}) {
	out["always_allow"] = func(in *kops.AlwaysAllowAuthorizationSpec) interface{} {
		return func(in *kops.AlwaysAllowAuthorizationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AlwaysAllowAuthorizationSpec) interface{} {
				return func(in kops.AlwaysAllowAuthorizationSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceAlwaysAllowAuthorizationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AlwaysAllow)
	out["rbac"] = func(in *kops.RBACAuthorizationSpec) interface{} {
		return func(in *kops.RBACAuthorizationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.RBACAuthorizationSpec) interface{} {
				return func(in kops.RBACAuthorizationSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceRBACAuthorizationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.RBAC)
}

func FlattenResourceAuthorizationSpec(in kops.AuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAuthorizationSpecInto(in, out)
	return out
}
