package structures

import (
	"reflect"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandKubeProxyConfig(in map[string]interface{}) kops.KubeProxyConfig {
	if in == nil {
		panic("expand KubeProxyConfig failure, in is nil")
	}
	return kops.KubeProxyConfig{
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		CPURequest: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cpu_request"]),
		CPULimit: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cpu_limit"]),
		MemoryRequest: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["memory_request"]),
		MemoryLimit: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["memory_limit"]),
		LogLevel: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["log_level"]),
		ClusterCIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_cidr"]),
		HostnameOverride: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["hostname_override"]),
		BindAddress: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["bind_address"]),
		Master: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["master"]),
		MetricsBindAddress: func(in interface{}) *string {
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
		}(in["metrics_bind_address"]),
		Enabled: func(in interface{}) *bool {
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
		}(in["enabled"]),
		ProxyMode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["proxy_mode"]),
		IPVSExcludeCIDRS: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["ip_vs_exclude_cidr_s"]),
		IPVSMinSyncPeriod: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["ip_vs_min_sync_period"]),
		IPVSScheduler: func(in interface{}) *string {
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
		}(in["ip_vs_scheduler"]),
		IPVSSyncPeriod: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["ip_vs_sync_period"]),
		FeatureGates: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
		}(in["feature_gates"]),
		ConntrackMaxPerCore: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["conntrack_max_per_core"]),
		ConntrackMin: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["conntrack_min"]),
	}
}

func FlattenKubeProxyConfig(in kops.KubeProxyConfig) map[string]interface{} {
	return map[string]interface{}{
		"image": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Image),
		"cpu_request": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.CPURequest),
		"cpu_limit": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.CPULimit),
		"memory_request": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.MemoryRequest),
		"memory_limit": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.MemoryLimit),
		"log_level": func(in int32) interface{} {
			return FlattenInt(int(in))
		}(in.LogLevel),
		"cluster_cidr": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ClusterCIDR),
		"hostname_override": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.HostnameOverride),
		"bind_address": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.BindAddress),
		"master": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Master),
		"metrics_bind_address": func(in *string) interface{} {
			return func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
		}(in.MetricsBindAddress),
		"enabled": func(in *bool) interface{} {
			return func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
		}(in.Enabled),
		"proxy_mode": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ProxyMode),
		"ip_vs_exclude_cidr_s": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.IPVSExcludeCIDRS),
		"ip_vs_min_sync_period": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.IPVSMinSyncPeriod),
		"ip_vs_scheduler": func(in *string) interface{} {
			return func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
		}(in.IPVSScheduler),
		"ip_vs_sync_period": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.IPVSSyncPeriod),
		"feature_gates": func(in map[string]string) interface{} {
			return func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				out := map[string]interface{}{}
				for key, in := range in {
					out[key] = FlattenString(string(in))
				}
				return out
			}(in)
		}(in.FeatureGates),
		"conntrack_max_per_core": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.ConntrackMaxPerCore),
		"conntrack_min": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.ConntrackMin),
	}
}
