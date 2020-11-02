package structures

import (
	"log"
	"reflect"

	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandRollingUpdate(in map[string]interface{}) kops.RollingUpdate {
	if in == nil {
		panic("expand RollingUpdate failure, in is nil")
	}
	return kops.RollingUpdate{
		DrainAndTerminate: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "drain_and_terminate", value)
			return value
		}(in["drain_and_terminate"]),
		MaxUnavailable: func(in interface{}) *intstr.IntOrString {
			value := func(in interface{}) *intstr.IntOrString {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in intstr.IntOrString) *intstr.IntOrString {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandIntOrString(in))
				return tmp
			}(in)
			log.Printf("%s - %#v", "max_unavailable", value)
			return value
		}(in["max_unavailable"]),
		MaxSurge: func(in interface{}) *intstr.IntOrString {
			value := func(in interface{}) *intstr.IntOrString {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in intstr.IntOrString) *intstr.IntOrString {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandIntOrString(in))
				return tmp
			}(in)
			log.Printf("%s - %#v", "max_surge", value)
			return value
		}(in["max_surge"]),
	}
}

func FlattenRollingUpdate(in kops.RollingUpdate) map[string]interface{} {
	return map[string]interface{}{
		"drain_and_terminate": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			log.Printf("%s - %v", "drain_and_terminate", value)
			return value
		}(in.DrainAndTerminate),
		"max_unavailable": func(in *intstr.IntOrString) interface{} {
			value := func(in *intstr.IntOrString) interface{} {
				if in == nil {
					return nil
				}
				return func(in intstr.IntOrString) interface{} {
					return FlattenIntOrString(in)
				}(*in)
			}(in)
			log.Printf("%s - %v", "max_unavailable", value)
			return value
		}(in.MaxUnavailable),
		"max_surge": func(in *intstr.IntOrString) interface{} {
			value := func(in *intstr.IntOrString) interface{} {
				if in == nil {
					return nil
				}
				return func(in intstr.IntOrString) interface{} {
					return FlattenIntOrString(in)
				}(*in)
			}(in)
			log.Printf("%s - %v", "max_surge", value)
			return value
		}(in.MaxSurge),
	}
}
