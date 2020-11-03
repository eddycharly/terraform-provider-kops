package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandUserData(in map[string]interface{}) kops.UserData {
	if in == nil {
		panic("expand UserData failure, in is nil")
	}
	return kops.UserData{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		Type: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["type"]),
		Content: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["content"]),
	}
}

func FlattenUserData(in kops.UserData) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Name),
		"type": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Type),
		"content": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Content),
	}
}
