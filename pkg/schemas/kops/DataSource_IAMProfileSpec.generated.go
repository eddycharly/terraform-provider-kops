package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
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

func ExpandDataSourceIAMProfileSpec(in map[string]interface{}) kops.IAMProfileSpec {
	if in == nil {
		panic("expand IAMProfileSpec failure, in is nil")
	}
	return kops.IAMProfileSpec{
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

func FlattenDataSourceIAMProfileSpecInto(in kops.IAMProfileSpec, out map[string]interface{}) {
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

func FlattenDataSourceIAMProfileSpec(in kops.IAMProfileSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceIAMProfileSpecInto(in, out)
	return out
}
