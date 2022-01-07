package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceDockerConfig(t *testing.T) {
	_default := kops.DockerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.DockerConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"authorization_plugins":    func() []interface{} { return nil }(),
					"bridge":                   nil,
					"bridge_ip":                nil,
					"data_root":                nil,
					"default_ulimit":           func() []interface{} { return nil }(),
					"default_runtime":          nil,
					"dns":                      func() []interface{} { return nil }(),
					"exec_opt":                 func() []interface{} { return nil }(),
					"exec_root":                nil,
					"experimental":             nil,
					"health_check":             false,
					"hosts":                    func() []interface{} { return nil }(),
					"ip_masq":                  nil,
					"ip_tables":                nil,
					"insecure_registry":        nil,
					"insecure_registries":      func() []interface{} { return nil }(),
					"live_restore":             nil,
					"log_driver":               nil,
					"log_level":                nil,
					"log_opt":                  func() []interface{} { return nil }(),
					"max_concurrent_downloads": nil,
					"max_concurrent_uploads":   nil,
					"max_download_attempts":    nil,
					"metrics_address":          nil,
					"mtu":                      nil,
					"packages":                 nil,
					"registry_mirrors":         func() []interface{} { return nil }(),
					"runtimes":                 func() []interface{} { return nil }(),
					"selinux_enabled":          nil,
					"skip_install":             false,
					"storage":                  nil,
					"storage_opts":             func() []interface{} { return nil }(),
					"user_namespace_remap":     "",
					"version":                  nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceDockerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceDockerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceDockerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"authorization_plugins":    func() []interface{} { return nil }(),
		"bridge":                   nil,
		"bridge_ip":                nil,
		"data_root":                nil,
		"default_ulimit":           func() []interface{} { return nil }(),
		"default_runtime":          nil,
		"dns":                      func() []interface{} { return nil }(),
		"exec_opt":                 func() []interface{} { return nil }(),
		"exec_root":                nil,
		"experimental":             nil,
		"health_check":             false,
		"hosts":                    func() []interface{} { return nil }(),
		"ip_masq":                  nil,
		"ip_tables":                nil,
		"insecure_registry":        nil,
		"insecure_registries":      func() []interface{} { return nil }(),
		"live_restore":             nil,
		"log_driver":               nil,
		"log_level":                nil,
		"log_opt":                  func() []interface{} { return nil }(),
		"max_concurrent_downloads": nil,
		"max_concurrent_uploads":   nil,
		"max_download_attempts":    nil,
		"metrics_address":          nil,
		"mtu":                      nil,
		"packages":                 nil,
		"registry_mirrors":         func() []interface{} { return nil }(),
		"runtimes":                 func() []interface{} { return nil }(),
		"selinux_enabled":          nil,
		"skip_install":             false,
		"storage":                  nil,
		"storage_opts":             func() []interface{} { return nil }(),
		"user_namespace_remap":     "",
		"version":                  nil,
	}
	type args struct {
		in kops.DockerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.DockerConfig{},
			},
			want: _default,
		},
		{
			name: "AuthorizationPlugins - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.AuthorizationPlugins = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Bridge - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.Bridge = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BridgeIp - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.BridgeIP = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DataRoot - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.DataRoot = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DefaultUlimit - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.DefaultUlimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DefaultRuntime - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.DefaultRuntime = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Dns - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.DNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExecOpt - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.ExecOpt = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExecRoot - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.ExecRoot = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Experimental - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.Experimental = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HealthCheck - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.HealthCheck = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Hosts - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.Hosts = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpMasq - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.IPMasq = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpTables - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.IPTables = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InsecureRegistry - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.InsecureRegistry = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InsecureRegistries - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.InsecureRegistries = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LiveRestore - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.LiveRestore = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogDriver - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.LogDriver = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.LogLevel = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogOpt - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.LogOpt = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxConcurrentDownloads - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.MaxConcurrentDownloads = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxConcurrentUploads - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.MaxConcurrentUploads = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxDownloadAttempts - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.MaxDownloadAttempts = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MetricsAddress - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.MetricsAddress = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MTU - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.MTU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Packages - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.Packages = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RegistryMirrors - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.RegistryMirrors = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Runtimes - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.Runtimes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SelinuxEnabled - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.SelinuxEnabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SkipInstall - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.SkipInstall = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Storage - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.Storage = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "StorageOpts - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.StorageOpts = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UserNamespaceRemap - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.UserNamespaceRemap = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceDockerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceDockerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceDockerConfig(t *testing.T) {
	_default := map[string]interface{}{
		"authorization_plugins":    func() []interface{} { return nil }(),
		"bridge":                   nil,
		"bridge_ip":                nil,
		"data_root":                nil,
		"default_ulimit":           func() []interface{} { return nil }(),
		"default_runtime":          nil,
		"dns":                      func() []interface{} { return nil }(),
		"exec_opt":                 func() []interface{} { return nil }(),
		"exec_root":                nil,
		"experimental":             nil,
		"health_check":             false,
		"hosts":                    func() []interface{} { return nil }(),
		"ip_masq":                  nil,
		"ip_tables":                nil,
		"insecure_registry":        nil,
		"insecure_registries":      func() []interface{} { return nil }(),
		"live_restore":             nil,
		"log_driver":               nil,
		"log_level":                nil,
		"log_opt":                  func() []interface{} { return nil }(),
		"max_concurrent_downloads": nil,
		"max_concurrent_uploads":   nil,
		"max_download_attempts":    nil,
		"metrics_address":          nil,
		"mtu":                      nil,
		"packages":                 nil,
		"registry_mirrors":         func() []interface{} { return nil }(),
		"runtimes":                 func() []interface{} { return nil }(),
		"selinux_enabled":          nil,
		"skip_install":             false,
		"storage":                  nil,
		"storage_opts":             func() []interface{} { return nil }(),
		"user_namespace_remap":     "",
		"version":                  nil,
	}
	type args struct {
		in kops.DockerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.DockerConfig{},
			},
			want: _default,
		},
		{
			name: "AuthorizationPlugins - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.AuthorizationPlugins = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Bridge - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.Bridge = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BridgeIp - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.BridgeIP = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DataRoot - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.DataRoot = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DefaultUlimit - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.DefaultUlimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DefaultRuntime - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.DefaultRuntime = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Dns - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.DNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExecOpt - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.ExecOpt = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExecRoot - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.ExecRoot = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Experimental - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.Experimental = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HealthCheck - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.HealthCheck = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Hosts - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.Hosts = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpMasq - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.IPMasq = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpTables - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.IPTables = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InsecureRegistry - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.InsecureRegistry = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InsecureRegistries - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.InsecureRegistries = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LiveRestore - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.LiveRestore = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogDriver - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.LogDriver = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.LogLevel = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogOpt - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.LogOpt = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxConcurrentDownloads - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.MaxConcurrentDownloads = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxConcurrentUploads - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.MaxConcurrentUploads = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxDownloadAttempts - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.MaxDownloadAttempts = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MetricsAddress - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.MetricsAddress = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MTU - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.MTU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Packages - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.Packages = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RegistryMirrors - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.RegistryMirrors = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Runtimes - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.Runtimes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SelinuxEnabled - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.SelinuxEnabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SkipInstall - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.SkipInstall = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Storage - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.Storage = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "StorageOpts - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.StorageOpts = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UserNamespaceRemap - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.UserNamespaceRemap = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.DockerConfig {
					subject := kops.DockerConfig{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceDockerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceDockerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
