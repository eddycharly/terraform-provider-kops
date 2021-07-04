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
			"profile": Computed(Nullable(String())),
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenDataSourceIAMProfileSpecInto(in kops.IAMProfileSpec, out map[string]interface{}) {
	out["profile"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Profile)
}

func FlattenDataSourceIAMProfileSpec(in kops.IAMProfileSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceIAMProfileSpecInto(in, out)
	return out
}
