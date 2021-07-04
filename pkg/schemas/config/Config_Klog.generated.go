package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ConfigKlog() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"v": Optional(Nullable(Int())),
		},
	}
}

func ExpandConfigKlog(in map[string]interface{}) config.Klog {
	if in == nil {
		panic("expand Klog failure, in is nil")
	}
	out := config.Klog{}
	if in, ok := in["v"]; ok && in != nil {
		out.V = func(in interface{}) *int {
			if in == nil {
				return nil
			}
			return func(in int) *int { return &in }(func(in interface{}) int { return int(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}

func FlattenConfigKlogInto(in config.Klog, out map[string]interface{}) {
	out["v"] = func(in *int) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int) interface{} { return int(in) }(*in)}
	}(in.V)
}

func FlattenConfigKlog(in config.Klog) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenConfigKlogInto(in, out)
	return out
}
