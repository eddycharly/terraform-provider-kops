package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAWSLoadBalancerControllerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": Computed(Nullable(Bool())),
			"version": Computed(Nullable(String())),
		},
	}
}

func ExpandDataSourceAWSLoadBalancerControllerConfig(in map[string]interface{}) kops.AWSLoadBalancerControllerConfig {
	if in == nil {
		panic("expand AWSLoadBalancerControllerConfig failure, in is nil")
	}
	out := kops.AWSLoadBalancerControllerConfig{}
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
	if in, ok := in["version"]; ok && in != nil {
		out.Version = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenDataSourceAWSLoadBalancerControllerConfigInto(in kops.AWSLoadBalancerControllerConfig, out map[string]interface{}) {
	out["enabled"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.Enabled)
	out["version"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Version)
}

func FlattenDataSourceAWSLoadBalancerControllerConfig(in kops.AWSLoadBalancerControllerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAWSLoadBalancerControllerConfigInto(in, out)
	return out
}
