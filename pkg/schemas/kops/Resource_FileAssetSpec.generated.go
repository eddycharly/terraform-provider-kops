package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceFileAssetSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":      Required(String()),
			"path":      Required(String()),
			"roles":     Optional(List(String())),
			"content":   Required(String()),
			"is_base64": Optional(Bool()),
		},
	}
}

func ExpandResourceFileAssetSpec(in map[string]interface{}) kops.FileAssetSpec {
	if in == nil {
		panic("expand FileAssetSpec failure, in is nil")
	}
	out := kops.FileAssetSpec{}
	if in, ok := in["name"]; ok && in != nil {
		out.Name = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["path"]; ok && in != nil {
		out.Path = func(in interface{}) string { return string(in.(string)) }(in)
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
	if in, ok := in["content"]; ok && in != nil {
		out.Content = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["is_base64"]; ok && in != nil {
		out.IsBase64 = func(in interface{}) bool { return in.(bool) }(in)
	}
	return out
}

func FlattenResourceFileAssetSpecInto(in kops.FileAssetSpec, out map[string]interface{}) {
	out["name"] = func(in string) interface{} { return string(in) }(in.Name)
	out["path"] = func(in string) interface{} { return string(in) }(in.Path)
	out["roles"] = func(in []kops.InstanceGroupRole) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in kops.InstanceGroupRole) interface{} { return string(in) }(in))
		}
		return out
	}(in.Roles)
	out["content"] = func(in string) interface{} { return string(in) }(in.Content)
	out["is_base64"] = func(in bool) interface{} { return in }(in.IsBase64)
}

func FlattenResourceFileAssetSpec(in kops.FileAssetSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceFileAssetSpecInto(in, out)
	return out
}
