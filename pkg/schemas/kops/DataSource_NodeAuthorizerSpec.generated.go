package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceNodeAuthorizerSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authorizer": Computed(String()),
			"features":   Computed(List(String())),
			"image":      Computed(String()),
			"node_url":   Computed(String()),
			"port":       Computed(Int()),
			"interval":   Computed(Nullable(Duration())),
			"timeout":    Computed(Nullable(Duration())),
			"token_ttl":  Computed(Nullable(Duration())),
		},
	}
}

func ExpandDataSourceNodeAuthorizerSpec(in map[string]interface{}) kops.NodeAuthorizerSpec {
	if in == nil {
		panic("expand NodeAuthorizerSpec failure, in is nil")
	}
	out := kops.NodeAuthorizerSpec{}
	if in, ok := in["authorizer"]; ok && in != nil {
		out.Authorizer = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["features"]; ok && in != nil {
		out.Features = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["node_url"]; ok && in != nil {
		out.NodeURL = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["port"]; ok && in != nil {
		out.Port = func(in interface{}) int { return int(in.(int)) }(in)
	}
	if in, ok := in["interval"]; ok && in != nil {
		out.Interval = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["timeout"]; ok && in != nil {
		out.Timeout = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["token_ttl"]; ok && in != nil {
		out.TokenTTL = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenDataSourceNodeAuthorizerSpecInto(in kops.NodeAuthorizerSpec, out map[string]interface{}) {
	out["authorizer"] = func(in string) interface{} { return string(in) }(in.Authorizer)
	out["features"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.Features)
	out["image"] = func(in string) interface{} { return string(in) }(in.Image)
	out["node_url"] = func(in string) interface{} { return string(in) }(in.NodeURL)
	out["port"] = func(in int) interface{} { return int(in) }(in.Port)
	out["interval"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.Interval)
	out["timeout"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.Timeout)
	out["token_ttl"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.TokenTTL)
}

func FlattenDataSourceNodeAuthorizerSpec(in kops.NodeAuthorizerSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceNodeAuthorizerSpecInto(in, out)
	return out
}
