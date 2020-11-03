package structures

import (
	"reflect"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandOpenstackMonitor(in map[string]interface{}) kops.OpenstackMonitor {
	if in == nil {
		panic("expand OpenstackMonitor failure, in is nil")
	}
	return kops.OpenstackMonitor{
		Delay: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			return value
		}(in["delay"]),
		Timeout: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			return value
		}(in["timeout"]),
		MaxRetries: func(in interface{}) *int {
			value := func(in interface{}) *int {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int) *int {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["max_retries"]),
	}
}

func FlattenOpenstackMonitor(in kops.OpenstackMonitor) map[string]interface{} {
	return map[string]interface{}{
		"delay": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.Delay),
		"timeout": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.Timeout),
		"max_retries": func(in *int) interface{} {
			value := func(in *int) interface{} {
				if in == nil {
					return nil
				}
				return func(in int) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.MaxRetries),
	}
}
