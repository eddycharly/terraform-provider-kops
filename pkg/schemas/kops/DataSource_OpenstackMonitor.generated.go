package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceOpenstackMonitor() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"delay":       Computed(Nullable(String())),
			"timeout":     Computed(Nullable(String())),
			"max_retries": Computed(Nullable(Int())),
		},
	}
}

func ExpandDataSourceOpenstackMonitor(in map[string]interface{}) kops.OpenstackMonitor {
	if in == nil {
		panic("expand OpenstackMonitor failure, in is nil")
	}
	out := kops.OpenstackMonitor{}
	if in, ok := in["delay"]; ok && in != nil {
		out.Delay = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["timeout"]; ok && in != nil {
		out.Timeout = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["max_retries"]; ok && in != nil {
		out.MaxRetries = func(in interface{}) *int {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int) *int { return &in }(func(in interface{}) int { return int(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenDataSourceOpenstackMonitorInto(in kops.OpenstackMonitor, out map[string]interface{}) {
	out["delay"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Delay)
	out["timeout"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Timeout)
	out["max_retries"] = func(in *int) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int) interface{} { return int(in) }(*in)}
	}(in.MaxRetries)
}

func FlattenDataSourceOpenstackMonitor(in kops.OpenstackMonitor) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceOpenstackMonitorInto(in, out)
	return out
}
