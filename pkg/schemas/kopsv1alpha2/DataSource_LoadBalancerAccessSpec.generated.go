package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops/v1alpha2"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceLoadBalancerAccessSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"class":                      ComputedString(),
			"type":                       ComputedString(),
			"idle_timeout_seconds":       ComputedInt(),
			"security_group_override":    ComputedString(),
			"additional_security_groups": ComputedList(String()),
			"use_for_internal_api":       ComputedBool(),
			"ssl_certificate":            ComputedString(),
			"ssl_policy":                 ComputedString(),
			"cross_zone_load_balancing":  ComputedBool(),
			"subnets":                    ComputedList(DataSourceLoadBalancerSubnetSpec()),
			"access_log":                 ComputedStruct(DataSourceAccessLogSpec()),
		},
	}

	return res
}

func ExpandDataSourceLoadBalancerAccessSpec(in map[string]interface{}) kopsv1alpha2.LoadBalancerAccessSpec {
	if in == nil {
		panic("expand LoadBalancerAccessSpec failure, in is nil")
	}
	return kopsv1alpha2.LoadBalancerAccessSpec{
		Class: func(in interface{}) kopsv1alpha2.LoadBalancerClass {
			return v1alpha2.LoadBalancerClass(ExpandString(in))
		}(in["class"]),
		Type: func(in interface{}) kopsv1alpha2.LoadBalancerType {
			return v1alpha2.LoadBalancerType(ExpandString(in))
		}(in["type"]),
		IdleTimeoutSeconds: func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
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
			if in == nil {
				return nil
			}
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
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["additional_security_groups"]),
		UseForInternalAPI: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["use_for_internal_api"]),
		SSLCertificate: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ssl_certificate"]),
		SSLPolicy: func(in interface{}) *string {
			if in == nil {
				return nil
			}
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
			if in == nil {
				return nil
			}
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
		Subnets: func(in interface{}) []kopsv1alpha2.LoadBalancerSubnetSpec {
			return func(in interface{}) []kopsv1alpha2.LoadBalancerSubnetSpec {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.LoadBalancerSubnetSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.LoadBalancerSubnetSpec {
						if in == nil {
							return kopsv1alpha2.LoadBalancerSubnetSpec{}
						}
						return ExpandDataSourceLoadBalancerSubnetSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["subnets"]),
		AccessLog: func(in interface{}) *kopsv1alpha2.AccessLogSpec {
			return func(in interface{}) *kopsv1alpha2.AccessLogSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.AccessLogSpec) *kopsv1alpha2.AccessLogSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.AccessLogSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceAccessLogSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.AccessLogSpec{}
				}(in))
			}(in)
		}(in["access_log"]),
	}
}

func FlattenDataSourceLoadBalancerAccessSpecInto(in kopsv1alpha2.LoadBalancerAccessSpec, out map[string]interface{}) {
	out["class"] = func(in kopsv1alpha2.LoadBalancerClass) interface{} {
		return FlattenString(string(in))
	}(in.Class)
	out["type"] = func(in kopsv1alpha2.LoadBalancerType) interface{} {
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
	}(in.UseForInternalAPI)
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
	out["subnets"] = func(in []kopsv1alpha2.LoadBalancerSubnetSpec) interface{} {
		return func(in []kopsv1alpha2.LoadBalancerSubnetSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.LoadBalancerSubnetSpec) interface{} {
					return FlattenDataSourceLoadBalancerSubnetSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.Subnets)
	out["access_log"] = func(in *kopsv1alpha2.AccessLogSpec) interface{} {
		return func(in *kopsv1alpha2.AccessLogSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.AccessLogSpec) interface{} {
				return func(in kopsv1alpha2.AccessLogSpec) []interface{} {
					return []interface{}{FlattenDataSourceAccessLogSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AccessLog)
}

func FlattenDataSourceLoadBalancerAccessSpec(in kopsv1alpha2.LoadBalancerAccessSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceLoadBalancerAccessSpecInto(in, out)
	return out
}
