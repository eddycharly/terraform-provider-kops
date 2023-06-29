package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceIAMProfileSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"profile": ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceIAMProfileSpec(in map[string]interface{}) kopsv1alpha2.IAMProfileSpec {
	if in == nil {
		panic("expand IAMProfileSpec failure, in is nil")
	}
	return kopsv1alpha2.IAMProfileSpec{
		Profile: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
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

func FlattenDataSourceIAMProfileSpecInto(in kopsv1alpha2.IAMProfileSpec, out map[string]interface{}) {
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

func FlattenDataSourceIAMProfileSpec(in kopsv1alpha2.IAMProfileSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceIAMProfileSpecInto(in, out)
	return out
}
