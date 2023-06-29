package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops/v1alpha2"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceHookSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":             RequiredString(),
			"enabled":          OptionalBool(),
			"roles":            OptionalList(String()),
			"requires":         OptionalList(String()),
			"before":           OptionalList(String()),
			"exec_container":   OptionalStruct(ResourceExecContainerAction()),
			"manifest":         OptionalString(),
			"use_raw_manifest": OptionalBool(),
		},
	}

	return res
}

func ExpandResourceHookSpec(in map[string]interface{}) kopsv1alpha2.HookSpec {
	if in == nil {
		panic("expand HookSpec failure, in is nil")
	}
	return kopsv1alpha2.HookSpec{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		Enabled: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["enabled"]),
		Roles: func(in interface{}) []kopsv1alpha2.InstanceGroupRole {
			return func(in interface{}) []kopsv1alpha2.InstanceGroupRole {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.InstanceGroupRole
				for _, in := range in.([]interface{}) {
					out = append(out, v1alpha2.InstanceGroupRole(ExpandString(in)))
				}
				return out
			}(in)
		}(in["roles"]),
		Requires: func(in interface{}) []string {
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
		}(in["requires"]),
		Before: func(in interface{}) []string {
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
		}(in["before"]),
		ExecContainer: func(in interface{}) *kopsv1alpha2.ExecContainerAction {
			return func(in interface{}) *kopsv1alpha2.ExecContainerAction {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.ExecContainerAction) *kopsv1alpha2.ExecContainerAction {
					return &in
				}(func(in interface{}) kopsv1alpha2.ExecContainerAction {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceExecContainerAction(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.ExecContainerAction{}
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

func FlattenResourceHookSpecInto(in kopsv1alpha2.HookSpec, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["enabled"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Enabled)
	out["roles"] = func(in []kopsv1alpha2.InstanceGroupRole) interface{} {
		return func(in []kopsv1alpha2.InstanceGroupRole) []interface{} {
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
	out["exec_container"] = func(in *kopsv1alpha2.ExecContainerAction) interface{} {
		return func(in *kopsv1alpha2.ExecContainerAction) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.ExecContainerAction) interface{} {
				return func(in kopsv1alpha2.ExecContainerAction) []interface{} {
					return []interface{}{FlattenResourceExecContainerAction(in)}
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

func FlattenResourceHookSpec(in kopsv1alpha2.HookSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceHookSpecInto(in, out)
	return out
}
