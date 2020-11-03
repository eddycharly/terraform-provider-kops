package structures

import (
	"reflect"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandNodeAuthorizerSpec(in map[string]interface{}) kops.NodeAuthorizerSpec {
	if in == nil {
		panic("expand NodeAuthorizerSpec failure, in is nil")
	}
	return kops.NodeAuthorizerSpec{
		Authorizer: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["authorizer"]),
		Features: func(in interface{}) *[]string {
			value := func(in interface{}) *[]string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in []string) *[]string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) []string {
					var out []string
					for _, in := range in.([]interface{}) {
						out = append(out, string(ExpandString(in)))
					}
					return out
				}(in))
				return tmp
			}(in)
			return value
		}(in["features"]),
		Image: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["image"]),
		NodeURL: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["node_url"]),
		Port: func(in interface{}) int {
			value := int(ExpandInt(in))
			return value
		}(in["port"]),
		Interval: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["interval"]),
		Timeout: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["timeout"]),
		TokenTTL: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["token_ttl"]),
	}
}

func FlattenNodeAuthorizerSpec(in kops.NodeAuthorizerSpec) map[string]interface{} {
	return map[string]interface{}{
		"authorizer": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Authorizer),
		"features": func(in *[]string) interface{} {
			value := func(in *[]string) interface{} {
				if in == nil {
					return nil
				}
				return func(in []string) interface{} {
					return func(in []string) []interface{} {
						var out []interface{}
						for _, in := range in {
							out = append(out, FlattenString(string(in)))
						}
						return out
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Features),
		"image": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Image),
		"node_url": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.NodeURL),
		"port": func(in int) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.Port),
		"interval": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.Interval),
		"timeout": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.Timeout),
		"token_ttl": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.TokenTTL),
	}
}
