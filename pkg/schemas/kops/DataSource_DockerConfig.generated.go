package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceDockerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authorization_plugins": Computed(List(String())),
			"bridge":                Computed(Nullable(String())),
			"bridge_ip":             Computed(Nullable(String())),
			"data_root":             Computed(Nullable(String())),
			"default_ulimit":        Computed(List(String())),
			"default_runtime":       Computed(Nullable(String())),
			"exec_opt":              Computed(List(String())),
			"exec_root":             Computed(Nullable(String())),
			"experimental":          Computed(Nullable(Bool())),
			"health_check":          Computed(Bool()),
			"hosts":                 Computed(List(String())),
			"ip_masq":               Computed(Nullable(Bool())),
			"ip_tables":             Computed(Nullable(Bool())),
			"insecure_registry":     Computed(Nullable(String())),
			"insecure_registries":   Computed(List(String())),
			"live_restore":          Computed(Nullable(Bool())),
			"log_driver":            Computed(Nullable(String())),
			"log_level":             Computed(Nullable(String())),
			"log_opt":               Computed(List(String())),
			"metrics_address":       Computed(Nullable(String())),
			"mtu":                   Computed(Nullable(Int())),
			"packages":              Computed(Struct(DataSourcePackagesConfig())),
			"registry_mirrors":      Computed(List(String())),
			"runtimes":              Computed(List(String())),
			"selinux_enabled":       Computed(Nullable(Bool())),
			"skip_install":          Computed(Bool()),
			"storage":               Computed(Nullable(String())),
			"storage_opts":          Computed(List(String())),
			"user_namespace_remap":  Computed(String()),
			"version":               Computed(Nullable(String())),
		},
	}
}

func ExpandDataSourceDockerConfig(in map[string]interface{}) kops.DockerConfig {
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["bridge_ip"]; ok && in != nil {
		out.BridgeIP = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["data_root"]; ok && in != nil {
		out.DataRoot = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["experimental"]; ok && in != nil {
		out.Experimental = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["ip_tables"]; ok && in != nil {
		out.IPTables = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["insecure_registry"]; ok && in != nil {
		out.InsecureRegistry = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["log_driver"]; ok && in != nil {
		out.LogDriver = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["log_level"]; ok && in != nil {
		out.LogLevel = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["mtu"]; ok && in != nil {
		out.MTU = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["packages"]; ok && in != nil {
		out.Packages = func(in interface{}) *kops.PackagesConfig {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.PackagesConfig) *kops.PackagesConfig { return &in }(func(in interface{}) kops.PackagesConfig {
					return ExpandDataSourcePackagesConfig(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenDataSourceDockerConfigInto(in kops.DockerConfig, out map[string]interface{}) {
	out["authorization_plugins"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.AuthorizationPlugins)
	out["bridge"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Bridge)
	out["bridge_ip"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.BridgeIP)
	out["data_root"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.DataRoot)
	out["default_ulimit"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.DefaultUlimit)
	out["default_runtime"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.DefaultRuntime)
	out["exec_opt"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.ExecOpt)
	out["exec_root"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.ExecRoot)
	out["experimental"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.Experimental)
	out["health_check"] = func(in bool) interface{} { return in }(in.HealthCheck)
	out["hosts"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.Hosts)
	out["ip_masq"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.IPMasq)
	out["ip_tables"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.IPTables)
	out["insecure_registry"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.InsecureRegistry)
	out["insecure_registries"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.InsecureRegistries)
	out["live_restore"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.LiveRestore)
	out["log_driver"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.LogDriver)
	out["log_level"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.LogLevel)
	out["log_opt"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.LogOpt)
	out["metrics_address"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.MetricsAddress)
	out["mtu"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.MTU)
	out["packages"] = func(in *kops.PackagesConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.PackagesConfig) interface{} { return FlattenDataSourcePackagesConfig(in) }(*in)
	}(in.Packages)
	out["registry_mirrors"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.RegistryMirrors)
	out["runtimes"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.Runtimes)
	out["selinux_enabled"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.SelinuxEnabled)
	out["skip_install"] = func(in bool) interface{} { return in }(in.SkipInstall)
	out["storage"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Storage)
	out["storage_opts"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.StorageOpts)
	out["user_namespace_remap"] = func(in string) interface{} { return string(in) }(in.UserNamespaceRemap)
	out["version"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Version)
}

func FlattenDataSourceDockerConfig(in kops.DockerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceDockerConfigInto(in, out)
	return out
}
