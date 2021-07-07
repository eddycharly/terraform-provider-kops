package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ConfigKlog() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"verbosity": Nullable(OptionalInt()),
		},
	}

	return res
}

func ExpandConfigKlog(in map[string]interface{}) config.Klog {
	if in == nil {
		panic("expand Klog failure, in is nil")
	}
	return config.Klog{
		Verbosity: func(in interface{}) *int {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && len(in) == 1 {
				return func(in interface{}) *int {
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
				}(in[0].(map[string]interface{})["value"])
			}
			return nil
		}(in["verbosity"]),
	}
}

func FlattenConfigKlogInto(in config.Klog, out map[string]interface{}) {
	out["verbosity"] = func(in *int) interface{} {
		if in == nil {
			return nil
		}
		return []interface{}{map[string]interface{}{"value": func(in *int) interface{} {
			if in == nil {
				return nil
			}
			return func(in int) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)}}
	}(in.Verbosity)
}

func FlattenConfigKlog(in config.Klog) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenConfigKlogInto(in, out)
	return out
}
