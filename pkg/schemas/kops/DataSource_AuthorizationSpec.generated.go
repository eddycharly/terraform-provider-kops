package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAuthorizationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"always_allow": Computed(Struct(DataSourceAlwaysAllowAuthorizationSpec())),
			"rbac":         Computed(Struct(DataSourceRBACAuthorizationSpec())),
		},
	}
}

func ExpandDataSourceAuthorizationSpec(in map[string]interface{}) kops.AuthorizationSpec {
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
				return ExpandDataSourceAlwaysAllowAuthorizationSpec(in.(map[string]interface{}))
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
				return ExpandDataSourceRBACAuthorizationSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	return out
}

func FlattenDataSourceAuthorizationSpecInto(in kops.AuthorizationSpec, out map[string]interface{}) {
	out["always_allow"] = func(in *kops.AlwaysAllowAuthorizationSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.AlwaysAllowAuthorizationSpec) interface{} {
			return FlattenDataSourceAlwaysAllowAuthorizationSpec(in)
		}(*in)
	}(in.AlwaysAllow)
	out["rbac"] = func(in *kops.RBACAuthorizationSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.RBACAuthorizationSpec) interface{} { return FlattenDataSourceRBACAuthorizationSpec(in) }(*in)
	}(in.RBAC)
}

func FlattenDataSourceAuthorizationSpec(in kops.AuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAuthorizationSpecInto(in, out)
	return out
}
