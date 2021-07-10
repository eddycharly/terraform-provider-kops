package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceKubeProxyConfig(t *testing.T) {
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
			if got := ExpandResourceKubeProxyConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceKubeProxyConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceKubeProxyConfigInto(t *testing.T) {
	_default := map[string]interface{}{
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
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceKubeProxyConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceKubeProxyConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceKubeProxyConfig(t *testing.T) {
	_default := map[string]interface{}{
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
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceKubeProxyConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceKubeProxyConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
