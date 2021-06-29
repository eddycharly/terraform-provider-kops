package schemas

import (
	"reflect"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ConfigKlog() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"v": OptionalInt(),
		},
	}
}

func ExpandConfigKlog(in map[string]interface{}) config.Klog {
	if in == nil {
		panic("expand Klog failure, in is nil")
	}
	return config.Klog{
		V: func(in interface{}) *int {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int) *int {
					return &in
				}(int(ExpandInt(in)))
			}(in)
		}(in["v"]),
	}
}

func FlattenConfigKlogInto(in config.Klog, out map[string]interface{}) {
	out["v"] = func(in *int) interface{} {
		return func(in *int) interface{} {
			if in == nil {
				return nil
			}
			return func(in int) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.V)
}

func FlattenConfigKlog(in config.Klog) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenConfigKlogInto(in, out)
	return out
}
