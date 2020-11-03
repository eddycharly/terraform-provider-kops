package structures

import (
	"reflect"

	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandNodeLocalDNSConfig(in map[string]interface{}) kops.NodeLocalDNSConfig {
	if in == nil {
		panic("expand NodeLocalDNSConfig failure, in is nil")
	}
	return kops.NodeLocalDNSConfig{
		Enabled: func(in interface{}) *bool {
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
			return value
		}(in["enabled"]),
		LocalIP: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["local_ip"]),
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
	}
}

func FlattenNodeLocalDNSConfig(in kops.NodeLocalDNSConfig) map[string]interface{} {
	return map[string]interface{}{
		"enabled": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.Enabled),
		"local_ip": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.LocalIP),
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
	}
}
