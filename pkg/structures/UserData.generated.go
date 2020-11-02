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
			value := string(ExpandString(in))
			return value
		}(in["name"]),
		Type: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["type"]),
		Content: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["content"]),
	}
}

func FlattenUserData(in kops.UserData) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Name),
		"type": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Type),
		"content": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Content),
	}
}
