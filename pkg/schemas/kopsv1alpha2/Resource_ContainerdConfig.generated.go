package schemas

import (
	"reflect"
	"sort"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceContainerdConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"address":          OptionalString(),
			"config_override":  OptionalString(),
			"log_level":        OptionalString(),
			"packages":         OptionalStruct(ResourcePackagesConfig()),
			"registry_mirrors": OptionalComplexMap(List(String())),
			"root":             OptionalString(),
			"skip_install":     OptionalBool(),
			"state":            OptionalString(),
			"version":          OptionalString(),
			"nvidia_gpu":       OptionalStruct(ResourceNvidiaGPUConfig()),
			"runc":             OptionalStruct(ResourceRunc()),
		},
	}

	return res
}

func ExpandResourceContainerdConfig(in map[string]interface{}) kopsv1alpha2.ContainerdConfig {
	if in == nil {
		panic("expand ContainerdConfig failure, in is nil")
	}
	return kopsv1alpha2.ContainerdConfig{
		Address: func(in interface{}) *string {
			if in == nil {
				return nil
			}
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
		}(in["address"]),
		ConfigOverride: func(in interface{}) *string {
			if in == nil {
				return nil
			}
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
		}(in["config_override"]),
		LogLevel: func(in interface{}) *string {
			if in == nil {
				return nil
			}
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
		}(in["log_level"]),
		Packages: func(in interface{}) *kopsv1alpha2.PackagesConfig {
			return func(in interface{}) *kopsv1alpha2.PackagesConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.PackagesConfig) *kopsv1alpha2.PackagesConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.PackagesConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourcePackagesConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.PackagesConfig{}
				}(in))
			}(in)
		}(in["packages"]),
		RegistryMirrors: func(in interface{}) map[string][]string {
			return func(in interface{}) map[string][]string {
				if in == nil {
					return nil
				}
				if in, ok := in.([]interface{}); ok {
					if len(in) > 0 {
						out := map[string][]string{}
						for _, in := range in {
							if in, ok := in.(map[string]interface{}); ok {
								key := ExpandString(in["key"])
								value := func(in interface{}) []string {
									return func(in interface{}) []string {
										if in == nil {
											return nil
										}
										var out []string
										for _, in := range in.([]interface{}) {
											out = append(out, string(ExpandString(in)))
										}
										return out
									}(in)
								}(in["value"])
								out[key] = value
							}
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["registry_mirrors"]),
		Root: func(in interface{}) *string {
			if in == nil {
				return nil
			}
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
		}(in["root"]),
		SkipInstall: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["skip_install"]),
		State: func(in interface{}) *string {
			if in == nil {
				return nil
			}
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
		}(in["state"]),
		Version: func(in interface{}) *string {
			if in == nil {
				return nil
			}
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
		}(in["version"]),
		NvidiaGPU: func(in interface{}) *kopsv1alpha2.NvidiaGPUConfig {
			return func(in interface{}) *kopsv1alpha2.NvidiaGPUConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.NvidiaGPUConfig) *kopsv1alpha2.NvidiaGPUConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.NvidiaGPUConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceNvidiaGPUConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.NvidiaGPUConfig{}
				}(in))
			}(in)
		}(in["nvidia_gpu"]),
		Runc: func(in interface{}) *kopsv1alpha2.Runc {
			return func(in interface{}) *kopsv1alpha2.Runc {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.Runc) *kopsv1alpha2.Runc {
					return &in
				}(func(in interface{}) kopsv1alpha2.Runc {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceRunc(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.Runc{}
				}(in))
			}(in)
		}(in["runc"]),
	}
}

func FlattenResourceContainerdConfigInto(in kopsv1alpha2.ContainerdConfig, out map[string]interface{}) {
	out["address"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Address)
	out["config_override"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.ConfigOverride)
	out["log_level"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.LogLevel)
	out["packages"] = func(in *kopsv1alpha2.PackagesConfig) interface{} {
		return func(in *kopsv1alpha2.PackagesConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.PackagesConfig) interface{} {
				return func(in kopsv1alpha2.PackagesConfig) []interface{} {
					return []interface{}{FlattenResourcePackagesConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Packages)
	out["registry_mirrors"] = func(in map[string][]string) interface{} {
		return func(in map[string][]string) []interface{} {
			if in == nil {
				return nil
			}
			keys := make([]string, 0, len(in))
			for key := range in {
				keys = append(keys, key)
			}
			sort.SliceStable(keys, func(i, j int) bool {
				return keys[i] < keys[j]
			})
			var out []interface{}
			for _, key := range keys {
				in := in[key]
				out = append(out, map[string]interface{}{
					"key": key,
					"value": func(in []string) []interface{} {
						var out []interface{}
						for _, in := range in {
							out = append(out, FlattenString(string(in)))
						}
						return out
					}(in),
				})
			}
			return out
		}(in)
	}(in.RegistryMirrors)
	out["root"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Root)
	out["skip_install"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.SkipInstall)
	out["state"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.State)
	out["version"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Version)
	out["nvidia_gpu"] = func(in *kopsv1alpha2.NvidiaGPUConfig) interface{} {
		return func(in *kopsv1alpha2.NvidiaGPUConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.NvidiaGPUConfig) interface{} {
				return func(in kopsv1alpha2.NvidiaGPUConfig) []interface{} {
					return []interface{}{FlattenResourceNvidiaGPUConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NvidiaGPU)
	out["runc"] = func(in *kopsv1alpha2.Runc) interface{} {
		return func(in *kopsv1alpha2.Runc) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.Runc) interface{} {
				return func(in kopsv1alpha2.Runc) []interface{} {
					return []interface{}{FlattenResourceRunc(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Runc)
}

func FlattenResourceContainerdConfig(in kopsv1alpha2.ContainerdConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceContainerdConfigInto(in, out)
	return out
}
