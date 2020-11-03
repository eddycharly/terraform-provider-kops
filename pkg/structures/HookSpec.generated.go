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
			return string(ExpandString(in))
		}(in["name"]),
		Disabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disabled"]),
		Roles: func(in interface{}) []kops.InstanceGroupRole {
			return func(in interface{}) []kops.InstanceGroupRole {
				var out []kops.InstanceGroupRole
				for _, in := range in.([]interface{}) {
					out = append(out, kops.InstanceGroupRole(ExpandString(in)))
				}
				return out
			}(in)
		}(in["roles"]),
		Requires: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["requires"]),
		Before: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["before"]),
		ExecContainer: func(in interface{}) *kops.ExecContainerAction {
			return func(in interface{}) *kops.ExecContainerAction {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
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
		}(in["exec_container"]),
		Manifest: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["manifest"]),
		UseRawManifest: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["use_raw_manifest"]),
	}
}

func FlattenHookSpec(in kops.HookSpec) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Name),
		"disabled": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.Disabled),
		"roles": func(in []kops.InstanceGroupRole) interface{} {
			return func(in []kops.InstanceGroupRole) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.Roles),
		"requires": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.Requires),
		"before": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.Before),
		"exec_container": func(in *kops.ExecContainerAction) interface{} {
			return func(in *kops.ExecContainerAction) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.ExecContainerAction) interface{} {
					return func(in kops.ExecContainerAction) []map[string]interface{} {
						return []map[string]interface{}{FlattenExecContainerAction(in)}
					}(in)
				}(*in)
			}(in)
		}(in.ExecContainer),
		"manifest": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Manifest),
		"use_raw_manifest": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.UseRawManifest),
	}
}
