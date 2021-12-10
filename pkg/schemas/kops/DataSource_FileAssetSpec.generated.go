package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceFileAssetSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":      ComputedString(),
			"path":      ComputedString(),
			"roles":     ComputedList(String()),
			"content":   ComputedString(),
			"is_base64": ComputedBool(),
		},
	}

	return res
}

func ExpandDataSourceFileAssetSpec(in map[string]interface{}) kops.FileAssetSpec {
	if in == nil {
		panic("expand FileAssetSpec failure, in is nil")
	}
	return kops.FileAssetSpec{
		Name: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["name"]),
		Path: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["path"]),
		Roles: func(in interface{}) []kops.InstanceGroupRole /**/ {
			return func(in interface{}) []kops.InstanceGroupRole {
				if in == nil {
					return nil
				}
				var out []kops.InstanceGroupRole
				for _, in := range in.([]interface{}) {
					out = append(out, kops.InstanceGroupRole(ExpandString(in)))
				}
				return out
			}(in)
		}(in["roles"]),
		Content: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["content"]),
		IsBase64: func(in interface{}) bool /**/ {
			return bool(ExpandBool(in))
		}(in["is_base64"]),
	}
}

func FlattenDataSourceFileAssetSpecInto(in kops.FileAssetSpec, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["path"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Path)
	out["roles"] = func(in []kops.InstanceGroupRole) interface{} {
		return func(in []kops.InstanceGroupRole) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Roles)
	out["content"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Content)
	out["is_base64"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.IsBase64)
}

func FlattenDataSourceFileAssetSpec(in kops.FileAssetSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceFileAssetSpecInto(in, out)
	return out
}
