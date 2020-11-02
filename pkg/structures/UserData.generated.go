package structures

import (
	"log"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandUserData(in map[string]interface{}) kops.UserData {
	if in == nil {
		panic("expand UserData failure, in is nil")
	}
	return kops.UserData{
		Name: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "name", value)
			return value
		}(in["name"]),
		Type: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "type", value)
			return value
		}(in["type"]),
		Content: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "content", value)
			return value
		}(in["content"]),
	}
}

func FlattenUserData(in kops.UserData) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "name", value)
			return value
		}(in.Name),
		"type": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "type", value)
			return value
		}(in.Type),
		"content": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "content", value)
			return value
		}(in.Content),
	}
}
