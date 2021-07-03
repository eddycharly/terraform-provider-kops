package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceLoadBalancerAccessSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"class":                      Computed(String()),
			"type":                       Computed(String()),
			"idle_timeout_seconds":       Computed(Ptr(Int())),
			"security_group_override":    Computed(Ptr(String())),
			"additional_security_groups": Computed(List(String())),
			"use_for_internal_api":       Computed(Bool()),
			"ssl_certificate":            Computed(String()),
			"ssl_policy":                 Computed(Ptr(String())),
			"cross_zone_load_balancing":  Computed(Ptr(Bool())),
			"subnets":                    Computed(List(Struct(DataSourceLoadBalancerSubnetSpec()))),
		},
	}
}

func ExpandDataSourceLoadBalancerAccessSpec(in map[string]interface{}) kops.LoadBalancerAccessSpec {
	if in == nil {
		panic("expand LoadBalancerAccessSpec failure, in is nil")
	}
	out := kops.LoadBalancerAccessSpec{}
	if in, ok := in["class"]; ok && in != nil {
		out.Class = func(in interface{}) kops.LoadBalancerClass { return kops.LoadBalancerClass(in.(string)) }(in)
	}
	if in, ok := in["type"]; ok && in != nil {
		out.Type = func(in interface{}) kops.LoadBalancerType { return kops.LoadBalancerType(in.(string)) }(in)
	}
	if in, ok := in["idle_timeout_seconds"]; ok && in != nil {
		out.IdleTimeoutSeconds = func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			return func(in int64) *int64 { return &in }(func(in interface{}) int64 { return int64(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["security_group_override"]; ok && in != nil {
		out.SecurityGroupOverride = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["additional_security_groups"]; ok && in != nil {
		out.AdditionalSecurityGroups = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["use_for_internal_api"]; ok && in != nil {
		out.UseForInternalApi = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["ssl_certificate"]; ok && in != nil {
		out.SSLCertificate = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["ssl_policy"]; ok && in != nil {
		out.SSLPolicy = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["cross_zone_load_balancing"]; ok && in != nil {
		out.CrossZoneLoadBalancing = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["subnets"]; ok && in != nil {
		out.Subnets = func(in interface{}) []kops.LoadBalancerSubnetSpec {
			var out []kops.LoadBalancerSubnetSpec
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.LoadBalancerSubnetSpec {
					if in == nil {
						return kops.LoadBalancerSubnetSpec{}
					}
					return ExpandDataSourceLoadBalancerSubnetSpec(in.(map[string]interface{}))
				}(in))
			}
			return out
		}(in)
	}
	return out
}
