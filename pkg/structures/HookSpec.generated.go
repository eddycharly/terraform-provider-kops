package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandHookSpec(in map[string]interface{}) kops.HookSpec {
	if in == nil {
		panic("expand HookSpec failure, in is nil")
	}
	return kops.HookSpec{
		Name: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["name"]),
		Disabled: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["disabled"]),
		Roles: func(in interface{}) []kops.InstanceGroupRole {
			value := func(in interface{}) []kops.InstanceGroupRole {
				var out []kops.InstanceGroupRole
				for _, in := range in.([]interface{}) {
					out = append(out, kops.InstanceGroupRole(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["roles"]),
		Requires: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["requires"]),
		Before: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["before"]),
		ExecContainer: func(in interface{}) *kops.ExecContainerAction {
			value := func(in interface{}) *kops.ExecContainerAction {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in kops.ExecContainerAction) *kops.ExecContainerAction {
					return &in
				}(func(in interface{}) kops.ExecContainerAction {
					if in.([]interface{})[0] == nil {
						return kops.ExecContainerAction{}
					}
					return (ExpandExecContainerAction(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
			return value
		}(in["exec_container"]),
		Manifest: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["manifest"]),
		UseRawManifest: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["use_raw_manifest"]),
	}
}

func FlattenHookSpec(in kops.HookSpec) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Name),
		"disabled": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.Disabled),
		"roles": func(in []kops.InstanceGroupRole) interface{} {
			value := func(in []kops.InstanceGroupRole) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.Roles),
		"requires": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.Requires),
		"before": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.Before),
		"exec_container": func(in *kops.ExecContainerAction) interface{} {
			value := func(in *kops.ExecContainerAction) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.ExecContainerAction) interface{} {
					return func(in kops.ExecContainerAction) []map[string]interface{} {
						return []map[string]interface{}{FlattenExecContainerAction(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.ExecContainer),
		"manifest": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Manifest),
		"use_raw_manifest": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.UseRawManifest),
	}
}
