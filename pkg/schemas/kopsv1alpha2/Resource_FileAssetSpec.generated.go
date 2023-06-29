package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops/v1alpha2"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceFileAssetSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":      RequiredString(),
			"path":      RequiredString(),
			"roles":     OptionalList(String()),
			"content":   RequiredString(),
			"is_base64": OptionalBool(),
			"mode":      OptionalString(),
		},
	}

	return res
}

func ExpandResourceFileAssetSpec(in map[string]interface{}) kopsv1alpha2.FileAssetSpec {
	if in == nil {
		panic("expand FileAssetSpec failure, in is nil")
	}
	return kopsv1alpha2.FileAssetSpec{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		Path: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["path"]),
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
		Content: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["content"]),
		IsBase64: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["is_base64"]),
		Mode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["mode"]),
	}
}

func FlattenResourceFileAssetSpecInto(in kopsv1alpha2.FileAssetSpec, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["path"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Path)
	out["roles"] = func(in []kopsv1alpha2.InstanceGroupRole) interface{} {
		return func(in []kopsv1alpha2.InstanceGroupRole) []interface{} {
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
	out["mode"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Mode)
}

func FlattenResourceFileAssetSpec(in kopsv1alpha2.FileAssetSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceFileAssetSpecInto(in, out)
	return out
}
