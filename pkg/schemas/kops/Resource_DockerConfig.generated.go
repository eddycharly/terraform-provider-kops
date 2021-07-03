package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceDockerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authorization_plugins": Optional(List(String())),
			"bridge":                Optional(Ptr(String())),
			"bridge_ip":             Optional(Ptr(String())),
			"data_root":             Optional(Ptr(String())),
			"default_ulimit":        Optional(List(String())),
			"default_runtime":       Optional(Ptr(String())),
			"exec_opt":              Optional(List(String())),
			"exec_root":             Optional(Ptr(String())),
			"experimental":          Optional(Ptr(Bool())),
			"health_check":          Optional(Bool()),
			"hosts":                 Optional(List(String())),
			"ip_masq":               Optional(Ptr(Bool())),
			"ip_tables":             Optional(Ptr(Bool())),
			"insecure_registry":     Optional(Ptr(String())),
			"insecure_registries":   Optional(List(String())),
			"live_restore":          Optional(Ptr(Bool())),
			"log_driver":            Optional(Ptr(String())),
			"log_level":             Optional(Ptr(String())),
			"log_opt":               Optional(List(String())),
			"metrics_address":       Optional(Ptr(String())),
			"mtu":                   Optional(Ptr(Int())),
			"packages":              Optional(Ptr(Struct(ResourcePackagesConfig()))),
			"registry_mirrors":      Optional(List(String())),
			"runtimes":              Optional(List(String())),
			"selinux_enabled":       Optional(Ptr(Bool())),
			"skip_install":          Optional(Bool()),
			"storage":               Optional(Ptr(String())),
			"storage_opts":          Optional(List(String())),
			"user_namespace_remap":  Optional(String()),
			"version":               Optional(Ptr(String())),
		},
	}
}

func ExpandResourceDockerConfig(in map[string]interface{}) kops.DockerConfig {
	if in == nil {
		panic("expand DockerConfig failure, in is nil")
	}
	out := kops.DockerConfig{}
	if in, ok := in["authorization_plugins"]; ok && in != nil {
		out.AuthorizationPlugins = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["bridge"]; ok && in != nil {
		out.Bridge = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["bridge_ip"]; ok && in != nil {
		out.BridgeIP = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["data_root"]; ok && in != nil {
		out.DataRoot = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["default_ulimit"]; ok && in != nil {
		out.DefaultUlimit = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["default_runtime"]; ok && in != nil {
		out.DefaultRuntime = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["exec_opt"]; ok && in != nil {
		out.ExecOpt = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["exec_root"]; ok && in != nil {
		out.ExecRoot = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["experimental"]; ok && in != nil {
		out.Experimental = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["health_check"]; ok && in != nil {
		out.HealthCheck = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["hosts"]; ok && in != nil {
		out.Hosts = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["ip_masq"]; ok && in != nil {
		out.IPMasq = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["ip_tables"]; ok && in != nil {
		out.IPTables = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["insecure_registry"]; ok && in != nil {
		out.InsecureRegistry = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["insecure_registries"]; ok && in != nil {
		out.InsecureRegistries = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["live_restore"]; ok && in != nil {
		out.LiveRestore = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["log_driver"]; ok && in != nil {
		out.LogDriver = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["log_level"]; ok && in != nil {
		out.LogLevel = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["log_opt"]; ok && in != nil {
		out.LogOpt = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["metrics_address"]; ok && in != nil {
		out.MetricsAddress = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["mtu"]; ok && in != nil {
		out.MTU = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["packages"]; ok && in != nil {
		out.Packages = func(in interface{}) *kops.PackagesConfig {
			if in == nil {
				return nil
			}
			return func(in kops.PackagesConfig) *kops.PackagesConfig { return &in }(func(in interface{}) kops.PackagesConfig {
				if in == nil {
					return kops.PackagesConfig{}
				}
				return ExpandResourcePackagesConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["registry_mirrors"]; ok && in != nil {
		out.RegistryMirrors = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["runtimes"]; ok && in != nil {
		out.Runtimes = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["selinux_enabled"]; ok && in != nil {
		out.SelinuxEnabled = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["skip_install"]; ok && in != nil {
		out.SkipInstall = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["storage"]; ok && in != nil {
		out.Storage = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["storage_opts"]; ok && in != nil {
		out.StorageOpts = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["user_namespace_remap"]; ok && in != nil {
		out.UserNamespaceRemap = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["version"]; ok && in != nil {
		out.Version = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
