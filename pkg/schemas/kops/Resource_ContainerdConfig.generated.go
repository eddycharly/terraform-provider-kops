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
			"address":          Optional(Ptr(String())),
			"config_override":  Optional(Ptr(String())),
			"log_level":        Optional(Ptr(String())),
			"packages":         Optional(Ptr(Struct(ResourcePackagesConfig()))),
			"registry_mirrors": Optional(Map(List(String()))),
			"root":             Optional(Ptr(String())),
			"skip_install":     Optional(Bool()),
			"state":            Optional(Ptr(String())),
			"version":          Optional(Ptr(String())),
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
