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
			"idle_timeout_seconds":       Computed(Nullable(Int())),
			"security_group_override":    Computed(Nullable(String())),
			"additional_security_groups": Computed(List(String())),
			"use_for_internal_api":       Computed(Bool()),
			"ssl_certificate":            Computed(String()),
			"ssl_policy":                 Computed(Nullable(String())),
			"cross_zone_load_balancing":  Computed(Nullable(Bool())),
			"subnets":                    Computed(List(DataSourceLoadBalancerSubnetSpec())),
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int64) *int64 { return &in }(func(in interface{}) int64 { return int64(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["security_group_override"]; ok && in != nil {
		out.SecurityGroupOverride = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["cross_zone_load_balancing"]; ok && in != nil {
		out.CrossZoneLoadBalancing = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["subnets"]; ok && in != nil {
		out.Subnets = func(in interface{}) []kops.LoadBalancerSubnetSpec {
			var out []kops.LoadBalancerSubnetSpec
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.LoadBalancerSubnetSpec {
					return ExpandDataSourceLoadBalancerSubnetSpec(in.(map[string]interface{}))
				}(in))
			}
			return out
		}(in)
	}
	return out
}

func FlattenDataSourceLoadBalancerAccessSpecInto(in kops.LoadBalancerAccessSpec, out map[string]interface{}) {
	out["class"] = func(in kops.LoadBalancerClass) interface{} { return string(in) }(in.Class)
	out["type"] = func(in kops.LoadBalancerType) interface{} { return string(in) }(in.Type)
	out["idle_timeout_seconds"] = func(in *int64) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int64) interface{} { return int(in) }(*in)}
	}(in.IdleTimeoutSeconds)
	out["security_group_override"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.SecurityGroupOverride)
	out["additional_security_groups"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.AdditionalSecurityGroups)
	out["use_for_internal_api"] = func(in bool) interface{} { return in }(in.UseForInternalApi)
	out["ssl_certificate"] = func(in string) interface{} { return string(in) }(in.SSLCertificate)
	out["ssl_policy"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.SSLPolicy)
	out["cross_zone_load_balancing"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.CrossZoneLoadBalancing)
	out["subnets"] = func(in []kops.LoadBalancerSubnetSpec) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in kops.LoadBalancerSubnetSpec) interface{} { return FlattenDataSourceLoadBalancerSubnetSpec(in) }(in))
		}
		return out
	}(in.Subnets)
}

func FlattenDataSourceLoadBalancerAccessSpec(in kops.LoadBalancerAccessSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceLoadBalancerAccessSpecInto(in, out)
	return out
}
