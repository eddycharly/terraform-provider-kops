package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandEnvVar(in map[string]interface{}) kops.EnvVar {
	if in == nil {
		panic("expand EnvVar failure, in is nil")
	}
	return kops.EnvVar{
		Name: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["name"]),
		Value: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["value"]),
	}
}

func FlattenEnvVar(in kops.EnvVar) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Name),
		"value": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Value),
	}
}
