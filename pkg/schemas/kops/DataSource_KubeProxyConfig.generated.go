package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceKubeProxyConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image":                  Computed(String()),
			"cpu_request":            Computed(String()),
			"cpu_limit":              Computed(String()),
			"memory_request":         Computed(String()),
			"memory_limit":           Computed(String()),
			"log_level":              Computed(Int()),
			"cluster_cidr":           Computed(String()),
			"hostname_override":      Computed(String()),
			"bind_address":           Computed(String()),
			"master":                 Computed(String()),
			"metrics_bind_address":   Computed(Ptr(String())),
			"enabled":                Computed(Ptr(Bool())),
			"proxy_mode":             Computed(String()),
			"ip_vs_exclude_cidr_s":   Computed(List(String())),
			"ip_vs_min_sync_period":  Computed(Ptr(Duration())),
			"ip_vs_scheduler":        Computed(Ptr(String())),
			"ip_vs_sync_period":      Computed(Ptr(Duration())),
			"feature_gates":          Computed(Map(String())),
			"conntrack_max_per_core": Computed(Ptr(Int())),
			"conntrack_min":          Computed(Ptr(Int())),
		},
	}
}

func ExpandDataSourceKubeProxyConfig(in map[string]interface{}) kops.KubeProxyConfig {
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
