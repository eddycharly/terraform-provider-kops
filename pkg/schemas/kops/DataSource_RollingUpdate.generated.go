package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceRollingUpdate() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"drain_and_terminate": Computed(Nullable(Bool())),
			"max_unavailable":     Computed(Nullable(IntOrString())),
			"max_surge":           Computed(Nullable(IntOrString())),
		},
	}
}

func ExpandDataSourceRollingUpdate(in map[string]interface{}) kops.RollingUpdate {
	if in == nil {
		panic("expand RollingUpdate failure, in is nil")
	}
	out := kops.RollingUpdate{}
	if in, ok := in["drain_and_terminate"]; ok && in != nil {
		out.DrainAndTerminate = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["max_unavailable"]; ok && in != nil {
		out.MaxUnavailable = func(in interface{}) *intstr.IntOrString {
			if in == nil {
				return nil
			}
			return func(in intstr.IntOrString) *intstr.IntOrString { return &in }(func(in interface{}) intstr.IntOrString { return ExpandIntOrString(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["max_surge"]; ok && in != nil {
		out.MaxSurge = func(in interface{}) *intstr.IntOrString {
			if in == nil {
				return nil
			}
			return func(in intstr.IntOrString) *intstr.IntOrString { return &in }(func(in interface{}) intstr.IntOrString { return ExpandIntOrString(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}

func FlattenDataSourceRollingUpdateInto(in kops.RollingUpdate, out map[string]interface{}) {
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

func FlattenDataSourceRollingUpdate(in kops.RollingUpdate) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceRollingUpdateInto(in, out)
	return out
}
