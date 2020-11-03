package structures

import (
	"reflect"

	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandAwsAuthenticationSpec(in map[string]interface{}) kops.AwsAuthenticationSpec {
	if in == nil {
		panic("expand AwsAuthenticationSpec failure, in is nil")
	}
	return kops.AwsAuthenticationSpec{
		Image: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["image"]),
		MemoryRequest: func(in interface{}) *resource.Quantity {
			value := func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in resource.Quantity) *resource.Quantity {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandQuantity(in))
				return tmp
			}(in)
			return value
		}(in["memory_request"]),
		CPURequest: func(in interface{}) *resource.Quantity {
			value := func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in resource.Quantity) *resource.Quantity {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandQuantity(in))
				return tmp
			}(in)
			return value
		}(in["cpu_request"]),
		MemoryLimit: func(in interface{}) *resource.Quantity {
			value := func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in resource.Quantity) *resource.Quantity {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandQuantity(in))
				return tmp
			}(in)
			return value
		}(in["memory_limit"]),
		CPULimit: func(in interface{}) *resource.Quantity {
			value := func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in resource.Quantity) *resource.Quantity {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandQuantity(in))
				return tmp
			}(in)
			return value
		}(in["cpu_limit"]),
	}
}

func FlattenAwsAuthenticationSpec(in kops.AwsAuthenticationSpec) map[string]interface{} {
	return map[string]interface{}{
		"image": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Image),
		"memory_request": func(in *resource.Quantity) interface{} {
			value := func(in *resource.Quantity) interface{} {
				if in == nil {
					return nil
				}
				return func(in resource.Quantity) interface{} {
					return FlattenQuantity(in)
				}(*in)
			}(in)
			return value
		}(in.MemoryRequest),
		"cpu_request": func(in *resource.Quantity) interface{} {
			value := func(in *resource.Quantity) interface{} {
				if in == nil {
					return nil
				}
				return func(in resource.Quantity) interface{} {
					return FlattenQuantity(in)
				}(*in)
			}(in)
			return value
		}(in.CPURequest),
		"memory_limit": func(in *resource.Quantity) interface{} {
			value := func(in *resource.Quantity) interface{} {
				if in == nil {
					return nil
				}
				return func(in resource.Quantity) interface{} {
					return FlattenQuantity(in)
				}(*in)
			}(in)
			return value
		}(in.MemoryLimit),
		"cpu_limit": func(in *resource.Quantity) interface{} {
			value := func(in *resource.Quantity) interface{} {
				if in == nil {
					return nil
				}
				return func(in resource.Quantity) interface{} {
					return FlattenQuantity(in)
				}(*in)
			}(in)
			return value
		}(in.CPULimit),
	}
}
