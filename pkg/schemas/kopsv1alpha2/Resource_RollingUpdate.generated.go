package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceRollingUpdate() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"drain_and_terminate": OptionalBool(),
			"max_unavailable":     OptionalIntOrString(),
			"max_surge":           OptionalIntOrString(),
		},
	}

	return res
}

func ExpandResourceRollingUpdate(in map[string]interface{}) kopsv1alpha2.RollingUpdate {
	if in == nil {
		panic("expand RollingUpdate failure, in is nil")
	}
	return kopsv1alpha2.RollingUpdate{
		DrainAndTerminate: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
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
		}(in["drain_and_terminate"]),
		MaxUnavailable: func(in interface{}) *intstr.IntOrString {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *intstr.IntOrString {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in intstr.IntOrString) *intstr.IntOrString {
					return &in
				}(ExpandIntOrString(in))
			}(in)
		}(in["max_unavailable"]),
		MaxSurge: func(in interface{}) *intstr.IntOrString {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *intstr.IntOrString {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in intstr.IntOrString) *intstr.IntOrString {
					return &in
				}(ExpandIntOrString(in))
			}(in)
		}(in["max_surge"]),
	}
}

func FlattenResourceRollingUpdateInto(in kopsv1alpha2.RollingUpdate, out map[string]interface{}) {
	out["drain_and_terminate"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.DrainAndTerminate)
	out["max_unavailable"] = func(in *intstr.IntOrString) interface{} {
		return func(in *intstr.IntOrString) interface{} {
			if in == nil {
				return nil
			}
			return func(in intstr.IntOrString) interface{} {
				return FlattenIntOrString(in)
			}(*in)
		}(in)
	}(in.MaxUnavailable)
	out["max_surge"] = func(in *intstr.IntOrString) interface{} {
		return func(in *intstr.IntOrString) interface{} {
			if in == nil {
				return nil
			}
			return func(in intstr.IntOrString) interface{} {
				return FlattenIntOrString(in)
			}(*in)
		}(in)
	}(in.MaxSurge)
}

func FlattenResourceRollingUpdate(in kopsv1alpha2.RollingUpdate) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceRollingUpdateInto(in, out)
	return out
}
