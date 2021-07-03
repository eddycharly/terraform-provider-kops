package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceIAMProfileSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"profile": Computed(Ptr(String())),
		},
	}
}

func ExpandDataSourceIAMProfileSpec(in map[string]interface{}) kops.IAMProfileSpec {
	if in == nil {
		panic("expand IAMProfileSpec failure, in is nil")
	}
	out := kops.IAMProfileSpec{}
	if in, ok := in["profile"]; ok && in != nil {
		out.Profile = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
