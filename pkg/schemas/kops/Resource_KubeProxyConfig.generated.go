package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceKubeProxyConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image":                  Optional(String()),
			"cpu_request":            Optional(String()),
			"cpu_limit":              Optional(String()),
			"memory_request":         Optional(String()),
			"memory_limit":           Optional(String()),
			"log_level":              Optional(Int()),
			"cluster_cidr":           Optional(String()),
			"hostname_override":      Optional(String()),
			"bind_address":           Optional(String()),
			"master":                 Optional(String()),
			"metrics_bind_address":   Optional(Ptr(String())),
			"enabled":                Optional(Ptr(Bool())),
			"proxy_mode":             Optional(String()),
			"ip_vs_exclude_cidr_s":   Optional(List(String())),
			"ip_vs_min_sync_period":  Optional(Ptr(Duration())),
			"ip_vs_scheduler":        Optional(Ptr(String())),
			"ip_vs_sync_period":      Optional(Ptr(Duration())),
			"feature_gates":          Optional(Map(String())),
			"conntrack_max_per_core": Optional(Ptr(Int())),
			"conntrack_min":          Optional(Ptr(Int())),
		},
	}
}

func ExpandResourceKubeProxyConfig(in map[string]interface{}) kops.KubeProxyConfig {
	if in == nil {
		panic("expand KubeProxyConfig failure, in is nil")
	}
	out := kops.KubeProxyConfig{}
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cpu_request"]; ok && in != nil {
		out.CPURequest = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cpu_limit"]; ok && in != nil {
		out.CPULimit = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["memory_request"]; ok && in != nil {
		out.MemoryRequest = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["memory_limit"]; ok && in != nil {
		out.MemoryLimit = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["log_level"]; ok && in != nil {
		out.LogLevel = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	if in, ok := in["cluster_cidr"]; ok && in != nil {
		out.ClusterCIDR = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["hostname_override"]; ok && in != nil {
		out.HostnameOverride = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["bind_address"]; ok && in != nil {
		out.BindAddress = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["master"]; ok && in != nil {
		out.Master = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["metrics_bind_address"]; ok && in != nil {
		out.MetricsBindAddress = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["enabled"]; ok && in != nil {
		out.Enabled = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["proxy_mode"]; ok && in != nil {
		out.ProxyMode = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["ip_vs_exclude_cidr_s"]; ok && in != nil {
		out.IPVSExcludeCIDRS = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["ip_vs_min_sync_period"]; ok && in != nil {
		out.IPVSMinSyncPeriod = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["ip_vs_scheduler"]; ok && in != nil {
		out.IPVSScheduler = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["ip_vs_sync_period"]; ok && in != nil {
		out.IPVSSyncPeriod = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["feature_gates"]; ok && in != nil {
		out.FeatureGates = func(in interface{}) map[string]string {
			if in == nil {
				return nil
			}
			out := map[string]string{}
			for key, in := range in.(map[string]interface{}) {
				out[key] = func(in interface{}) string { return string(in.(string)) }(in)
			}
			return out
		}(in)
	}
	if in, ok := in["conntrack_max_per_core"]; ok && in != nil {
		out.ConntrackMaxPerCore = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["conntrack_min"]; ok && in != nil {
		out.ConntrackMin = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
