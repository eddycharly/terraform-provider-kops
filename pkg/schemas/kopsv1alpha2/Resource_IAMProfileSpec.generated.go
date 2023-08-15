package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceIAMProfileSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"profile": RequiredString(),
		},
	}

	return res
}

func ExpandResourceIAMProfileSpec(in map[string]interface{}) kopsv1alpha2.IAMProfileSpec {
	if in == nil {
		panic("expand IAMProfileSpec failure, in is nil")
	}
	return kopsv1alpha2.IAMProfileSpec{
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

func FlattenResourceIAMProfileSpecInto(in kopsv1alpha2.IAMProfileSpec, out map[string]interface{}) {
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

func FlattenResourceIAMProfileSpec(in kopsv1alpha2.IAMProfileSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceIAMProfileSpecInto(in, out)
	return out
}
