package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandAuthenticationSpec(in map[string]interface{}) kops.AuthenticationSpec {
	if in == nil {
		panic("expand AuthenticationSpec failure, in is nil")
	}
	return kops.AuthenticationSpec{
		Kopeio: func(in interface{}) *kops.KopeioAuthenticationSpec {
			return func(in interface{}) *kops.KopeioAuthenticationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KopeioAuthenticationSpec) *kops.KopeioAuthenticationSpec {
					return &in
				}(func(in interface{}) kops.KopeioAuthenticationSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.KopeioAuthenticationSpec{}
					}
					return (ExpandKopeioAuthenticationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kopeio"]),
		Aws: func(in interface{}) *kops.AwsAuthenticationSpec {
			return func(in interface{}) *kops.AwsAuthenticationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AwsAuthenticationSpec) *kops.AwsAuthenticationSpec {
					return &in
				}(func(in interface{}) kops.AwsAuthenticationSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.AwsAuthenticationSpec{}
					}
					return (ExpandAwsAuthenticationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["aws"]),
	}
}

func FlattenAuthenticationSpec(in kops.AuthenticationSpec) map[string]interface{} {
	return map[string]interface{}{
		"kopeio": func(in *kops.KopeioAuthenticationSpec) interface{} {
			return func(in *kops.KopeioAuthenticationSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KopeioAuthenticationSpec) interface{} {
					return func(in kops.KopeioAuthenticationSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenKopeioAuthenticationSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Kopeio),
		"aws": func(in *kops.AwsAuthenticationSpec) interface{} {
			return func(in *kops.AwsAuthenticationSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.AwsAuthenticationSpec) interface{} {
					return func(in kops.AwsAuthenticationSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenAwsAuthenticationSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Aws),
	}
}
