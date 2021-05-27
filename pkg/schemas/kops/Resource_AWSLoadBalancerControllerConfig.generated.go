package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceAWSLoadBalancerControllerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": OptionalBool(),
			"version": OptionalString(),
		},
	}
}

func ExpandResourceAWSLoadBalancerControllerConfig(in map[string]interface{}) kops.AWSLoadBalancerControllerConfig {
	if in == nil {
		panic("expand AWSLoadBalancerControllerConfig failure, in is nil")
	}
	return kops.AWSLoadBalancerControllerConfig{
		Enabled: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["enabled"]),
		Version: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["version"]),
	}
}

func FlattenResourceAWSLoadBalancerControllerConfigInto(in kops.AWSLoadBalancerControllerConfig, out map[string]interface{}) {
	out["enabled"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Enabled)
	out["version"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Version)
}

func FlattenResourceAWSLoadBalancerControllerConfig(in kops.AWSLoadBalancerControllerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAWSLoadBalancerControllerConfigInto(in, out)
	return out
}
