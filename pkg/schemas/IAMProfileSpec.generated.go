package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func IAMProfileSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"profile": RequiredString(),
		},
	}
}

func ExpandIAMProfileSpec(in map[string]interface{}) kops.IAMProfileSpec {
	if in == nil {
		panic("expand IAMProfileSpec failure, in is nil")
	}
	return kops.IAMProfileSpec{
		Profile: func(in interface{}) *string {
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["profile"]),
	}
}

func FlattenIAMProfileSpecInto(in kops.IAMProfileSpec, out map[string]interface{}) {
	out["profile"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Profile)
}

func FlattenIAMProfileSpec(in kops.IAMProfileSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenIAMProfileSpecInto(in, out)
	return out
}
