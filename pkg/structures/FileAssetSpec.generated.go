package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandFileAssetSpec(in map[string]interface{}) kops.FileAssetSpec {
	if in == nil {
		panic("expand FileAssetSpec failure, in is nil")
	}
	return kops.FileAssetSpec{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		Path: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["path"]),
		Roles: func(in interface{}) []kops.InstanceGroupRole {
			return func(in interface{}) []kops.InstanceGroupRole {
				var out []kops.InstanceGroupRole
				for _, in := range in.([]interface{}) {
					out = append(out, kops.InstanceGroupRole(ExpandString(in)))
				}
				return out
			}(in)
		}(in["roles"]),
		Content: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["content"]),
		IsBase64: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["is_base_64"]),
	}
}

func FlattenFileAssetSpec(in kops.FileAssetSpec) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Name),
		"path": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Path),
		"roles": func(in []kops.InstanceGroupRole) interface{} {
			return func(in []kops.InstanceGroupRole) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.Roles),
		"content": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Content),
		"is_base_64": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.IsBase64),
	}
}
