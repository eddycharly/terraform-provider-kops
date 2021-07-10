package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceKubeProxyConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KubeProxyConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceKubeProxyConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceKubeProxyConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceKubeProxyConfigInto(t *testing.T) {
	type args struct {
		in  kops.KubeProxyConfig
		out map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FlattenDataSourceKubeProxyConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceKubeProxyConfig(t *testing.T) {
	type args struct {
		in kops.KubeProxyConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KubeProxyConfig{},
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "CPURequest - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.CPURequest = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "CPULimit - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.CPULimit = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.MemoryRequest = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.MemoryLimit = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.LogLevel = 0
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "ClusterCidr - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.ClusterCIDR = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "HostnameOverride - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.HostnameOverride = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "BindAddress - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.BindAddress = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "Master - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.Master = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "MetricsBindAddress - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.MetricsBindAddress = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "ProxyMode - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.ProxyMode = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "IpVSExcludeCidrS - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.IPVSExcludeCIDRS = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "IpVSMinSyncPeriod - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.IPVSMinSyncPeriod = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "IpVSScheduler - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.IPVSScheduler = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "IpVSSyncPeriod - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.IPVSSyncPeriod = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "FeatureGates - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.FeatureGates = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "ConntrackMaxPerCore - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.ConntrackMaxPerCore = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
		{
			name: "ConntrackMin - default",
			args: args{
				in: func() kops.KubeProxyConfig {
					subject := kops.KubeProxyConfig{}
					subject.ConntrackMin = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                  "",
				"cpu_request":            "",
				"cpu_limit":              "",
				"memory_request":         "",
				"memory_limit":           "",
				"log_level":              0,
				"cluster_cidr":           "",
				"hostname_override":      "",
				"bind_address":           "",
				"master":                 "",
				"metrics_bind_address":   nil,
				"enabled":                nil,
				"proxy_mode":             "",
				"ip_vs_exclude_cidr_s":   func() []interface{} { return nil }(),
				"ip_vs_min_sync_period":  nil,
				"ip_vs_scheduler":        nil,
				"ip_vs_sync_period":      nil,
				"feature_gates":          func() map[string]interface{} { return nil }(),
				"conntrack_max_per_core": nil,
				"conntrack_min":          nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceKubeProxyConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceKubeProxyConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
