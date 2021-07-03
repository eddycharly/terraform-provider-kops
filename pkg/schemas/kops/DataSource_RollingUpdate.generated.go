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
			"drain_and_terminate": Computed(Ptr(Bool())),
			"max_unavailable":     Computed(Ptr(IntOrString())),
			"max_surge":           Computed(Ptr(IntOrString())),
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
