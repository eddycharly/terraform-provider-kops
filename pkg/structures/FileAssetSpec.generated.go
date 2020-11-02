package structures

import (
	"log"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandFileAssetSpec(in map[string]interface{}) kops.FileAssetSpec {
	if in == nil {
		panic("expand FileAssetSpec failure, in is nil")
	}
	return kops.FileAssetSpec{
		Name: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "name", value)
			return value
		}(in["name"]),
		Path: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "path", value)
			return value
		}(in["path"]),
		Roles: func(in interface{}) []kops.InstanceGroupRole {
			value := func(in interface{}) []kops.InstanceGroupRole {
				var out []kops.InstanceGroupRole
				for _, in := range in.([]interface{}) {
					out = append(out, kops.InstanceGroupRole(ExpandString(in)))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "roles", value)
			return value
		}(in["roles"]),
		Content: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "content", value)
			return value
		}(in["content"]),
		IsBase64: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			log.Printf("%s - %#v", "is_base_64", value)
			return value
		}(in["is_base_64"]),
	}
}

func FlattenFileAssetSpec(in kops.FileAssetSpec) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "name", value)
			return value
		}(in.Name),
		"path": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "path", value)
			return value
		}(in.Path),
		"roles": func(in []kops.InstanceGroupRole) interface{} {
			value := func(in []kops.InstanceGroupRole) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			log.Printf("%s - %v", "roles", value)
			return value
		}(in.Roles),
		"content": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "content", value)
			return value
		}(in.Content),
		"is_base_64": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			log.Printf("%s - %v", "is_base_64", value)
			return value
		}(in.IsBase64),
	}
}
