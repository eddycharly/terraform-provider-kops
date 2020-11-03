package structures

import (
	"reflect"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandAuthenticationSpec(in map[string]interface{}) kops.AuthenticationSpec {
	if in == nil {
		panic("expand AuthenticationSpec failure, in is nil")
	}
	return kops.AuthenticationSpec{
		Kopeio: func(in interface{}) *kops.KopeioAuthenticationSpec {
			value := func(in interface{}) *kops.KopeioAuthenticationSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.KopeioAuthenticationSpec) *kops.KopeioAuthenticationSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.KopeioAuthenticationSpec {
					if in.([]interface{})[0] == nil {
						return kops.KopeioAuthenticationSpec{}
					}
					return (ExpandKopeioAuthenticationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["kopeio"]),
		Aws: func(in interface{}) *kops.AwsAuthenticationSpec {
			value := func(in interface{}) *kops.AwsAuthenticationSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.AwsAuthenticationSpec) *kops.AwsAuthenticationSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.AwsAuthenticationSpec {
					if in.([]interface{})[0] == nil {
						return kops.AwsAuthenticationSpec{}
					}
					return (ExpandAwsAuthenticationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["aws"]),
	}
}

func FlattenAuthenticationSpec(in kops.AuthenticationSpec) map[string]interface{} {
	return map[string]interface{}{
		"kopeio": func(in *kops.KopeioAuthenticationSpec) interface{} {
			value := func(in *kops.KopeioAuthenticationSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KopeioAuthenticationSpec) interface{} {
					return func(in kops.KopeioAuthenticationSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenKopeioAuthenticationSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Kopeio),
		"aws": func(in *kops.AwsAuthenticationSpec) interface{} {
			value := func(in *kops.AwsAuthenticationSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.AwsAuthenticationSpec) interface{} {
					return func(in kops.AwsAuthenticationSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenAwsAuthenticationSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Aws),
	}
}
