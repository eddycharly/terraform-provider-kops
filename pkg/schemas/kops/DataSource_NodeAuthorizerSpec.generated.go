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
			"interval":   Computed(Ptr(Duration())),
			"timeout":    Computed(Ptr(Duration())),
			"token_ttl":  Computed(Ptr(Duration())),
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
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["timeout"]; ok && in != nil {
		out.Timeout = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["token_ttl"]; ok && in != nil {
		out.TokenTTL = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
