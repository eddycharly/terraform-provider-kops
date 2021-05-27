package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceNTPConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"managed": ComputedBool(),
		},
	}
}

func ExpandDataSourceNTPConfig(in map[string]interface{}) kops.NTPConfig {
	if in == nil {
		panic("expand NTPConfig failure, in is nil")
	}
	return kops.NTPConfig{
		Managed: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["managed"]),
	}
}

func FlattenDataSourceNTPConfigInto(in kops.NTPConfig, out map[string]interface{}) {
	out["managed"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Managed)
}

func FlattenDataSourceNTPConfig(in kops.NTPConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceNTPConfigInto(in, out)
	return out
}
