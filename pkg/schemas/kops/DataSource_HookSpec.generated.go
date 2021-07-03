package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceHookSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":             Computed(String()),
			"disabled":         Computed(Bool()),
			"roles":            Computed(List(String())),
			"requires":         Computed(List(String())),
			"before":           Computed(List(String())),
			"exec_container":   Computed(Ptr(Struct(DataSourceExecContainerAction()))),
			"manifest":         Computed(String()),
			"use_raw_manifest": Computed(Bool()),
		},
	}
}

func ExpandDataSourceHookSpec(in map[string]interface{}) kops.HookSpec {
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
				return ExpandDataSourceExecContainerAction(in.(map[string]interface{}))
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
