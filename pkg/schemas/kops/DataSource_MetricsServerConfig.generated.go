package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceMetricsServerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":  Computed(Ptr(Bool())),
			"image":    Computed(Ptr(String())),
			"insecure": Computed(Ptr(Bool())),
		},
	}
}

func ExpandDataSourceMetricsServerConfig(in map[string]interface{}) kops.MetricsServerConfig {
	if in == nil {
		panic("expand MetricsServerConfig failure, in is nil")
	}
	out := kops.MetricsServerConfig{}
	if in, ok := in["enabled"]; ok && in != nil {
		out.Enabled = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["insecure"]; ok && in != nil {
		out.Insecure = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
