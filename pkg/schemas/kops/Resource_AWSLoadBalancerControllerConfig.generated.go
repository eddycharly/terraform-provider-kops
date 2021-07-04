package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceAWSLoadBalancerControllerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": Optional(Nullable(Bool())),
			"version": Optional(Nullable(String())),
		},
	}
}

func ExpandResourceAWSLoadBalancerControllerConfig(in map[string]interface{}) kops.AWSLoadBalancerControllerConfig {
	if in == nil {
		panic("expand AWSLoadBalancerControllerConfig failure, in is nil")
	}
	out := kops.AWSLoadBalancerControllerConfig{}
	if in, ok := in["enabled"]; ok && in != nil {
		out.Enabled = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["version"]; ok && in != nil {
		out.Version = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}

func FlattenResourceAWSLoadBalancerControllerConfigInto(in kops.AWSLoadBalancerControllerConfig, out map[string]interface{}) {
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

func FlattenResourceAWSLoadBalancerControllerConfig(in kops.AWSLoadBalancerControllerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAWSLoadBalancerControllerConfigInto(in, out)
	return out
}
