package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceRollingUpdate() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"drain_and_terminate": Optional(Nullable(Bool())),
			"max_unavailable":     Optional(Nullable(IntOrString())),
			"max_surge":           Optional(Nullable(IntOrString())),
		},
	}
}

func ExpandResourceRollingUpdate(in map[string]interface{}) kops.RollingUpdate {
	if in == nil {
		panic("expand RollingUpdate failure, in is nil")
	}
	out := kops.RollingUpdate{}
	if in, ok := in["drain_and_terminate"]; ok && in != nil {
		out.DrainAndTerminate = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["max_unavailable"]; ok && in != nil {
		out.MaxUnavailable = func(in interface{}) *intstr.IntOrString {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in intstr.IntOrString) *intstr.IntOrString { return &in }(func(in interface{}) intstr.IntOrString { return ExpandIntOrString(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["max_surge"]; ok && in != nil {
		out.MaxSurge = func(in interface{}) *intstr.IntOrString {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in intstr.IntOrString) *intstr.IntOrString { return &in }(func(in interface{}) intstr.IntOrString { return ExpandIntOrString(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenResourceRollingUpdateInto(in kops.RollingUpdate, out map[string]interface{}) {
	out["drain_and_terminate"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.DrainAndTerminate)
	out["max_unavailable"] = func(in *intstr.IntOrString) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in intstr.IntOrString) interface{} { return FlattenIntOrString(in) }(*in)}
	}(in.MaxUnavailable)
	out["max_surge"] = func(in *intstr.IntOrString) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in intstr.IntOrString) interface{} { return FlattenIntOrString(in) }(*in)}
	}(in.MaxSurge)
}

func FlattenResourceRollingUpdate(in kops.RollingUpdate) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceRollingUpdateInto(in, out)
	return out
}
