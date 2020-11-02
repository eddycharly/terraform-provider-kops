package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandLoadBalancerAccessSpec(in map[string]interface{}) kops.LoadBalancerAccessSpec {
	if in == nil {
		panic("expand LoadBalancerAccessSpec failure, in is nil")
	}
	return kops.LoadBalancerAccessSpec{
		Type: func(in interface{}) kops.LoadBalancerType {
			value := kops.LoadBalancerType(ExpandString(in))
			return value
		}(in["type"]),
		IdleTimeoutSeconds: func(in interface{}) *int64 {
			value := func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int64) *int64 {
					// if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
					// 	return nil
					// }
					return &in
				}(int64(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["idle_timeout_seconds"]),
		SecurityGroupOverride: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					// if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
					// 	return nil
					// }
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			return value
		}(in["security_group_override"]),
		AdditionalSecurityGroups: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["additional_security_groups"]),
		UseForInternalApi: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["use_for_internal_api"]),
		SSLCertificate: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["ssl_certificate"]),
		CrossZoneLoadBalancing: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					// if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
					// 	return nil
					// }
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			return value
		}(in["cross_zone_load_balancing"]),
	}
}

func FlattenLoadBalancerAccessSpec(in kops.LoadBalancerAccessSpec) map[string]interface{} {
	return map[string]interface{}{
		"type": func(in kops.LoadBalancerType) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Type),
		"idle_timeout_seconds": func(in *int64) interface{} {
			value := func(in *int64) interface{} {
				if in == nil {
					return nil
				}
				return func(in int64) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.IdleTimeoutSeconds),
		"security_group_override": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.SecurityGroupOverride),
		"additional_security_groups": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.AdditionalSecurityGroups),
		"use_for_internal_api": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.UseForInternalApi),
		"ssl_certificate": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.SSLCertificate),
		"cross_zone_load_balancing": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.CrossZoneLoadBalancing),
	}
}
