package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceHookSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":             RequiredString(),
			"disabled":         OptionalBool(),
			"roles":            OptionalList(String()),
			"requires":         OptionalList(String()),
			"before":           OptionalList(String()),
			"exec_container":   OptionalStruct(ResourceExecContainerAction()),
			"manifest":         OptionalString(),
			"use_raw_manifest": OptionalBool(),
		},
	}
}

func ExpandResourceHookSpec(in map[string]interface{}) kops.HookSpec {
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.ExecContainerAction{}
					}
					return (ExpandResourceExecContainerAction(in.([]interface{})[0].(map[string]interface{})))
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

func FlattenResourceHookSpecInto(in kops.HookSpec, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["disabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Disabled)
	out["roles"] = func(in []kops.InstanceGroupRole) interface{} {
		return func(in []kops.InstanceGroupRole) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Roles)
	out["requires"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Requires)
	out["before"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Before)
	out["exec_container"] = func(in *kops.ExecContainerAction) interface{} {
		return func(in *kops.ExecContainerAction) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.ExecContainerAction) interface{} {
				return func(in kops.ExecContainerAction) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceExecContainerAction(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ExecContainer)
	out["manifest"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Manifest)
	out["use_raw_manifest"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.UseRawManifest)
}

func FlattenResourceHookSpec(in kops.HookSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceHookSpecInto(in, out)
	return out
}
