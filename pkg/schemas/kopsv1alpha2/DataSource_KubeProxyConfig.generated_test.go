package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceKubeProxyConfig(t *testing.T) {
	_default := kopsv1alpha2.KubeProxyConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.KubeProxyConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"image":                  "",
					"cpu_request":            nil,
					"cpu_limit":              nil,
					"memory_request":         nil,
					"memory_limit":           nil,
					"log_level":              0,
					"cluster_cidr":           nil,
					"hostname_override":      "",
					"bind_address":           "",
					"master":                 "",
					"metrics_bind_address":   nil,
					"enabled":                nil,
					"proxy_mode":             "",
					"ip_vs_exclude_cidrs":    func() []interface{} { return nil }(),
					"ip_vs_min_sync_period":  nil,
					"ip_vs_scheduler":        nil,
					"ip_vs_sync_period":      nil,
					"feature_gates":          func() map[string]interface{} { return nil }(),
					"conntrack_max_per_core": nil,
					"conntrack_min":          nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceKubeProxyConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceKubeProxyConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKubeProxyConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"image":                  "",
		"cpu_request":            nil,
		"cpu_limit":              nil,
		"memory_request":         nil,
		"memory_limit":           nil,
		"log_level":              0,
		"cluster_cidr":           nil,
		"hostname_override":      "",
		"bind_address":           "",
		"master":                 "",
		"metrics_bind_address":   nil,
		"enabled":                nil,
		"proxy_mode":             "",
		"ip_vs_exclude_cidrs":    func() []interface{} { return nil }(),
		"ip_vs_min_sync_period":  nil,
		"ip_vs_scheduler":        nil,
		"ip_vs_sync_period":      nil,
		"feature_gates":          func() map[string]interface{} { return nil }(),
		"conntrack_max_per_core": nil,
		"conntrack_min":          nil,
	}
	type args struct {
		in kopsv1alpha2.KubeProxyConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.KubeProxyConfig{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuLimit - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.CPULimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.LogLevel = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterCidr - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.ClusterCIDR = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HostnameOverride - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.HostnameOverride = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BindAddress - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.BindAddress = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Master - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.Master = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MetricsBindAddress - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.MetricsBindAddress = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ProxyMode - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.ProxyMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpVSExcludeCidrs - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.IPVSExcludeCIDRs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpVSMinSyncPeriod - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.IPVSMinSyncPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpVSScheduler - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.IPVSScheduler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpVSSyncPeriod - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.IPVSSyncPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FeatureGates - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.FeatureGates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConntrackMaxPerCore - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.ConntrackMaxPerCore = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConntrackMin - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.ConntrackMin = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceKubeProxyConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKubeProxyConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKubeProxyConfig(t *testing.T) {
	_default := map[string]interface{}{
		"image":                  "",
		"cpu_request":            nil,
		"cpu_limit":              nil,
		"memory_request":         nil,
		"memory_limit":           nil,
		"log_level":              0,
		"cluster_cidr":           nil,
		"hostname_override":      "",
		"bind_address":           "",
		"master":                 "",
		"metrics_bind_address":   nil,
		"enabled":                nil,
		"proxy_mode":             "",
		"ip_vs_exclude_cidrs":    func() []interface{} { return nil }(),
		"ip_vs_min_sync_period":  nil,
		"ip_vs_scheduler":        nil,
		"ip_vs_sync_period":      nil,
		"feature_gates":          func() map[string]interface{} { return nil }(),
		"conntrack_max_per_core": nil,
		"conntrack_min":          nil,
	}
	type args struct {
		in kopsv1alpha2.KubeProxyConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.KubeProxyConfig{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuLimit - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.CPULimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.LogLevel = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterCidr - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.ClusterCIDR = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HostnameOverride - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.HostnameOverride = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BindAddress - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.BindAddress = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Master - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.Master = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MetricsBindAddress - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.MetricsBindAddress = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ProxyMode - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.ProxyMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpVSExcludeCidrs - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.IPVSExcludeCIDRs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpVSMinSyncPeriod - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.IPVSMinSyncPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpVSScheduler - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.IPVSScheduler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpVSSyncPeriod - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.IPVSSyncPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FeatureGates - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.FeatureGates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConntrackMaxPerCore - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.ConntrackMaxPerCore = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConntrackMin - default",
			args: args{
				in: func() kopsv1alpha2.KubeProxyConfig {
					subject := kopsv1alpha2.KubeProxyConfig{}
					subject.ConntrackMin = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceKubeProxyConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKubeProxyConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
