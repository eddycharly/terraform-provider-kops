package structures

import (
	"reflect"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandExternalDNSConfig(in map[string]interface{}) kops.ExternalDNSConfig {
	if in == nil {
		panic("expand ExternalDNSConfig failure, in is nil")
	}
	return kops.ExternalDNSConfig{
		Disable: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["disable"]),
		WatchIngress: func(in interface{}) *bool {
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
		}(in["watch_ingress"]),
		WatchNamespace: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["watch_namespace"]),
	}
}

func FlattenExternalDNSConfig(in kops.ExternalDNSConfig) map[string]interface{} {
	return map[string]interface{}{
		"disable": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.Disable),
		"watch_ingress": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.WatchIngress),
		"watch_namespace": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.WatchNamespace),
	}
}
