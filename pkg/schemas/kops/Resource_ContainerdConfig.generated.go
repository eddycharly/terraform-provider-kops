package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceContainerdConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"address":          Optional(Nullable(String())),
			"config_override":  Optional(Nullable(String())),
			"log_level":        Optional(Nullable(String())),
			"packages":         Optional(Struct(ResourcePackagesConfig())),
			"registry_mirrors": Optional(Map(List(String()))),
			"root":             Optional(Nullable(String())),
			"skip_install":     Optional(Bool()),
			"state":            Optional(Nullable(String())),
			"version":          Optional(Nullable(String())),
		},
	}
}

func ExpandResourceContainerdConfig(in map[string]interface{}) kops.ContainerdConfig {
	if in == nil {
		panic("expand ContainerdConfig failure, in is nil")
	}
	out := kops.ContainerdConfig{}
	if in, ok := in["address"]; ok && in != nil {
		out.Address = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["config_override"]; ok && in != nil {
		out.ConfigOverride = func(in interface{}) *string {
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
		out.RegistryMirrors = func(in interface{}) map[string][]string {
			if in == nil {
				return nil
			}
			out := map[string][]string{}
			for key, in := range in.(map[string]interface{}) {
				out[key] = func(in interface{}) []string {
					var out []string
					for _, in := range in.([]interface{}) {
						out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
					}
					return out
				}(in)
			}
			return out
		}(in)
	}
	if in, ok := in["root"]; ok && in != nil {
		out.Root = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["skip_install"]; ok && in != nil {
		out.SkipInstall = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["state"]; ok && in != nil {
		out.State = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
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

func FlattenResourceContainerdConfigInto(in kops.ContainerdConfig, out map[string]interface{}) {
	out["address"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Address)
	out["config_override"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.ConfigOverride)
	out["log_level"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.LogLevel)
	out["packages"] = func(in *kops.PackagesConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.PackagesConfig) interface{} { return FlattenResourcePackagesConfig(in) }(*in)
	}(in.Packages)
	out["registry_mirrors"] = func(in map[string][]string) interface{} {
		if in == nil {
			return nil
		}
		out := map[string]interface{}{}
		for key, in := range in {
			out[key] = func(in []string) interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in string) interface{} { return string(in) }(in))
				}
				return out
			}(in)
		}
		return out
	}(in.RegistryMirrors)
	out["root"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Root)
	out["skip_install"] = func(in bool) interface{} { return in }(in.SkipInstall)
	out["state"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.State)
	out["version"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Version)
}

func FlattenResourceContainerdConfig(in kops.ContainerdConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceContainerdConfigInto(in, out)
	return out
}
