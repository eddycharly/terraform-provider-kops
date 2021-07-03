package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceIAMSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"legacy":                   Computed(Bool()),
			"allow_container_registry": Computed(Bool()),
			"permissions_boundary":     Computed(Ptr(String())),
		},
	}
}

func ExpandDataSourceIAMSpec(in map[string]interface{}) kops.IAMSpec {
	if in == nil {
		panic("expand IAMSpec failure, in is nil")
	}
	out := kops.IAMSpec{}
	if in, ok := in["legacy"]; ok && in != nil {
		out.Legacy = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["allow_container_registry"]; ok && in != nil {
		out.AllowContainerRegistry = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["permissions_boundary"]; ok && in != nil {
		out.PermissionsBoundary = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
