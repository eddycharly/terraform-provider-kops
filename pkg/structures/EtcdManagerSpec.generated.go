package structures

import (
	"log"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandEtcdManagerSpec(in map[string]interface{}) kops.EtcdManagerSpec {
	if in == nil {
		panic("expand EtcdManagerSpec failure, in is nil")
	}
	return kops.EtcdManagerSpec{
		Image: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "image", value)
			return value
		}(in["image"]),
		Env: func(in interface{}) []kops.EnvVar {
			value := func(in interface{}) []kops.EnvVar {
				var out []kops.EnvVar
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.EnvVar {
						if in == nil {
							return kops.EnvVar{}
						}
						return (ExpandEnvVar(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "env", value)
			return value
		}(in["env"]),
	}
}

func FlattenEtcdManagerSpec(in kops.EtcdManagerSpec) map[string]interface{} {
	return map[string]interface{}{
		"image": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "image", value)
			return value
		}(in.Image),
		"env": func(in []kops.EnvVar) interface{} {
			value := func(in []kops.EnvVar) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.EnvVar) interface{} {
						return FlattenEnvVar(in)
					}(in))
				}
				return out
			}(in)
			log.Printf("%s - %v", "env", value)
			return value
		}(in.Env),
	}
}
