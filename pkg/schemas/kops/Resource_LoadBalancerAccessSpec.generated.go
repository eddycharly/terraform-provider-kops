package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceLoadBalancerAccessSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"class":                      OptionalString(),
			"type":                       RequiredString(),
			"idle_timeout_seconds":       OptionalInt(),
			"security_group_override":    OptionalString(),
			"additional_security_groups": OptionalList(String()),
			"use_for_internal_api":       OptionalBool(),
			"ssl_certificate":            OptionalString(),
			"ssl_policy":                 OptionalString(),
			"cross_zone_load_balancing":  OptionalBool(),
			"subnets":                    OptionalList(ResourceLoadBalancerSubnetSpec()),
		},
	}
}

func ExpandResourceLoadBalancerAccessSpec(in map[string]interface{}) kops.LoadBalancerAccessSpec {
	if in == nil {
		panic("expand LoadBalancerAccessSpec failure, in is nil")
	}
	return kops.LoadBalancerAccessSpec{
		Class: func(in interface{}) kops.LoadBalancerClass {
			return kops.LoadBalancerClass(ExpandString(in))
		}(in["class"]),
		Type: func(in interface{}) kops.LoadBalancerType {
			return kops.LoadBalancerType(ExpandString(in))
		}(in["type"]),
		IdleTimeoutSeconds: func(in interface{}) *int64 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					return &in
				}(int64(ExpandInt(in)))
			}(in)
		}(in["idle_timeout_seconds"]),
		SecurityGroupOverride: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["security_group_override"]),
		AdditionalSecurityGroups: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["additional_security_groups"]),
		UseForInternalApi: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["use_for_internal_api"]),
		SSLCertificate: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ssl_certificate"]),
		SSLPolicy: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["ssl_policy"]),
		CrossZoneLoadBalancing: func(in interface{}) *bool {
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
		}(in["cross_zone_load_balancing"]),
		Subnets: func(in interface{}) []kops.LoadBalancerSubnetSpec {
			return func(in interface{}) []kops.LoadBalancerSubnetSpec {
				var out []kops.LoadBalancerSubnetSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.LoadBalancerSubnetSpec {
						if in == nil {
							return kops.LoadBalancerSubnetSpec{}
						}
						return (ExpandResourceLoadBalancerSubnetSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["subnets"]),
	}
}

func FlattenResourceLoadBalancerAccessSpecInto(in kops.LoadBalancerAccessSpec, out map[string]interface{}) {
	out["class"] = func(in kops.LoadBalancerClass) interface{} {
		return FlattenString(string(in))
	}(in.Class)
	out["type"] = func(in kops.LoadBalancerType) interface{} {
		return FlattenString(string(in))
	}(in.Type)
	out["idle_timeout_seconds"] = func(in *int64) interface{} {
		return func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.IdleTimeoutSeconds)
	out["security_group_override"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.SecurityGroupOverride)
	out["additional_security_groups"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AdditionalSecurityGroups)
	out["use_for_internal_api"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.UseForInternalApi)
	out["ssl_certificate"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.SSLCertificate)
	out["ssl_policy"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.SSLPolicy)
	out["cross_zone_load_balancing"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.CrossZoneLoadBalancing)
	out["subnets"] = func(in []kops.LoadBalancerSubnetSpec) interface{} {
		return func(in []kops.LoadBalancerSubnetSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.LoadBalancerSubnetSpec) interface{} {
					return FlattenResourceLoadBalancerSubnetSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.Subnets)
}

func FlattenResourceLoadBalancerAccessSpec(in kops.LoadBalancerAccessSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceLoadBalancerAccessSpecInto(in, out)
	return out
}
