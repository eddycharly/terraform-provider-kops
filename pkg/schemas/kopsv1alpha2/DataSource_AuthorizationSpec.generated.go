package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceAuthorizationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"always_allow": ComputedStruct(DataSourceAlwaysAllowAuthorizationSpec()),
			"rbac":         ComputedStruct(DataSourceRBACAuthorizationSpec()),
		},
	}

	return res
}

func ExpandDataSourceAuthorizationSpec(in map[string]interface{}) kopsv1alpha2.AuthorizationSpec {
	if in == nil {
		panic("expand AuthorizationSpec failure, in is nil")
	}
	return kopsv1alpha2.AuthorizationSpec{
		AlwaysAllow: func(in interface{}) *kopsv1alpha2.AlwaysAllowAuthorizationSpec {
			return func(in interface{}) *kopsv1alpha2.AlwaysAllowAuthorizationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.AlwaysAllowAuthorizationSpec) *kopsv1alpha2.AlwaysAllowAuthorizationSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.AlwaysAllowAuthorizationSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceAlwaysAllowAuthorizationSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.AlwaysAllowAuthorizationSpec{}
				}(in))
			}(in)
		}(in["always_allow"]),
		RBAC: func(in interface{}) *kopsv1alpha2.RBACAuthorizationSpec {
			return func(in interface{}) *kopsv1alpha2.RBACAuthorizationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.RBACAuthorizationSpec) *kopsv1alpha2.RBACAuthorizationSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.RBACAuthorizationSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceRBACAuthorizationSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.RBACAuthorizationSpec{}
				}(in))
			}(in)
		}(in["rbac"]),
	}
}

func FlattenDataSourceAuthorizationSpecInto(in kopsv1alpha2.AuthorizationSpec, out map[string]interface{}) {
	out["always_allow"] = func(in *kopsv1alpha2.AlwaysAllowAuthorizationSpec) interface{} {
		return func(in *kopsv1alpha2.AlwaysAllowAuthorizationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.AlwaysAllowAuthorizationSpec) interface{} {
				return func(in kopsv1alpha2.AlwaysAllowAuthorizationSpec) []interface{} {
					return []interface{}{FlattenDataSourceAlwaysAllowAuthorizationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AlwaysAllow)
	out["rbac"] = func(in *kopsv1alpha2.RBACAuthorizationSpec) interface{} {
		return func(in *kopsv1alpha2.RBACAuthorizationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.RBACAuthorizationSpec) interface{} {
				return func(in kopsv1alpha2.RBACAuthorizationSpec) []interface{} {
					return []interface{}{FlattenDataSourceRBACAuthorizationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.RBAC)
}

func FlattenDataSourceAuthorizationSpec(in kopsv1alpha2.AuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAuthorizationSpecInto(in, out)
	return out
}
