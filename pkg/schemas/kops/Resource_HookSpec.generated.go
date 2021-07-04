package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceHookSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":             Required(String()),
			"disabled":         Optional(Bool()),
			"roles":            Optional(List(String())),
			"requires":         Optional(List(String())),
			"before":           Optional(List(String())),
			"exec_container":   Optional(Struct(ResourceExecContainerAction())),
			"manifest":         Optional(String()),
			"use_raw_manifest": Optional(Bool()),
		},
	}
}

func ExpandResourceHookSpec(in map[string]interface{}) kops.HookSpec {
	if in == nil {
		panic("expand HookSpec failure, in is nil")
	}
	out := kops.HookSpec{}
	if in, ok := in["name"]; ok && in != nil {
		out.Name = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["disabled"]; ok && in != nil {
		out.Disabled = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["roles"]; ok && in != nil {
		out.Roles = func(in interface{}) []kops.InstanceGroupRole {
			var out []kops.InstanceGroupRole
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.InstanceGroupRole { return kops.InstanceGroupRole(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["requires"]; ok && in != nil {
		out.Requires = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["before"]; ok && in != nil {
		out.Before = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["exec_container"]; ok && in != nil {
		out.ExecContainer = func(in interface{}) *kops.ExecContainerAction {
			if in == nil {
				return nil
			}
			return func(in kops.ExecContainerAction) *kops.ExecContainerAction { return &in }(func(in interface{}) kops.ExecContainerAction {
				if in == nil {
					return kops.ExecContainerAction{}
				}
				return ExpandResourceExecContainerAction(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["manifest"]; ok && in != nil {
		out.Manifest = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["use_raw_manifest"]; ok && in != nil {
		out.UseRawManifest = func(in interface{}) bool { return in.(bool) }(in)
	}
	return out
}

func FlattenResourceHookSpecInto(in kops.HookSpec, out map[string]interface{}) {
	out["name"] = func(in string) interface{} { return string(in) }(in.Name)
	out["disabled"] = func(in bool) interface{} { return in }(in.Disabled)
	out["roles"] = func(in []kops.InstanceGroupRole) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in kops.InstanceGroupRole) interface{} { return string(in) }(in))
		}
		return out
	}(in.Roles)
	out["requires"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.Requires)
	out["before"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.Before)
	out["exec_container"] = func(in *kops.ExecContainerAction) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.ExecContainerAction) interface{} { return FlattenResourceExecContainerAction(in) }(*in)
	}(in.ExecContainer)
	out["manifest"] = func(in string) interface{} { return string(in) }(in.Manifest)
	out["use_raw_manifest"] = func(in bool) interface{} { return in }(in.UseRawManifest)
}

func FlattenResourceHookSpec(in kops.HookSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceHookSpecInto(in, out)
	return out
}
