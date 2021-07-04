package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceMetricsServerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":  Optional(Nullable(Bool())),
			"image":    Optional(Nullable(String())),
			"insecure": Optional(Nullable(Bool())),
		},
	}
}

func ExpandResourceMetricsServerConfig(in map[string]interface{}) kops.MetricsServerConfig {
	if in == nil {
		panic("expand MetricsServerConfig failure, in is nil")
	}
	out := kops.MetricsServerConfig{}
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
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["insecure"]; ok && in != nil {
		out.Insecure = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenResourceMetricsServerConfigInto(in kops.MetricsServerConfig, out map[string]interface{}) {
	out["enabled"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.Enabled)
	out["image"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Image)
	out["insecure"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.Insecure)
}

func FlattenResourceMetricsServerConfig(in kops.MetricsServerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceMetricsServerConfigInto(in, out)
	return out
}
