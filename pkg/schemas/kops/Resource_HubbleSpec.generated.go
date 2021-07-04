package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceHubbleSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": Optional(Nullable(Bool())),
			"metrics": Optional(List(String())),
		},
	}
}

func ExpandResourceHubbleSpec(in map[string]interface{}) kops.HubbleSpec {
	if in == nil {
		panic("expand HubbleSpec failure, in is nil")
	}
	out := kops.HubbleSpec{}
	if in, ok := in["enabled"]; ok && in != nil {
		out.Enabled = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["metrics"]; ok && in != nil {
		out.Metrics = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	return out
}

func FlattenResourceHubbleSpecInto(in kops.HubbleSpec, out map[string]interface{}) {
	out["enabled"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.Enabled)
	out["metrics"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.Metrics)
}

func FlattenResourceHubbleSpec(in kops.HubbleSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceHubbleSpecInto(in, out)
	return out
}
