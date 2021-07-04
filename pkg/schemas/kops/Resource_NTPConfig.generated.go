package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceNTPConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"managed": Optional(Nullable(Bool())),
		},
	}
}

func ExpandResourceNTPConfig(in map[string]interface{}) kops.NTPConfig {
	if in == nil {
		panic("expand NTPConfig failure, in is nil")
	}
	out := kops.NTPConfig{}
	if in, ok := in["managed"]; ok && in != nil {
		out.Managed = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}

func FlattenResourceNTPConfigInto(in kops.NTPConfig, out map[string]interface{}) {
	out["managed"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.Managed)
}

func FlattenResourceNTPConfig(in kops.NTPConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceNTPConfigInto(in, out)
	return out
}
