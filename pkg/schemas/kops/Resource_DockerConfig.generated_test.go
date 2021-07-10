package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceDockerConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.DockerConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceDockerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceDockerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceDockerConfigInto(t *testing.T) {
	type args struct {
		in  kops.DockerConfig
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
			FlattenResourceDockerConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceDockerConfig(t *testing.T) {
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
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
			want: map[string]interface{}{
				"authorization_plugins": func() []interface{} { return nil }(),
				"bridge":                nil,
				"bridge_ip":             nil,
				"data_root":             nil,
				"default_ulimit":        func() []interface{} { return nil }(),
				"default_runtime":       nil,
				"exec_opt":              func() []interface{} { return nil }(),
				"exec_root":             nil,
				"experimental":          nil,
				"health_check":          false,
				"hosts":                 func() []interface{} { return nil }(),
				"ip_masq":               nil,
				"ip_tables":             nil,
				"insecure_registry":     nil,
				"insecure_registries":   func() []interface{} { return nil }(),
				"live_restore":          nil,
				"log_driver":            nil,
				"log_level":             nil,
				"log_opt":               func() []interface{} { return nil }(),
				"metrics_address":       nil,
				"mtu":                   nil,
				"packages":              nil,
				"registry_mirrors":      func() []interface{} { return nil }(),
				"runtimes":              func() []interface{} { return nil }(),
				"selinux_enabled":       nil,
				"skip_install":          false,
				"storage":               nil,
				"storage_opts":          func() []interface{} { return nil }(),
				"user_namespace_remap":  "",
				"version":               nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceDockerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceDockerConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
