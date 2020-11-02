package structures

import (
	"log"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandAmazonVPCNetworkingSpec(in map[string]interface{}) kops.AmazonVPCNetworkingSpec {
	if in == nil {
		panic("expand AmazonVPCNetworkingSpec failure, in is nil")
	}
	return kops.AmazonVPCNetworkingSpec{
		ImageName: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "image_name", value)
			return value
		}(in["image_name"]),
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

func FlattenAmazonVPCNetworkingSpec(in kops.AmazonVPCNetworkingSpec) map[string]interface{} {
	return map[string]interface{}{
		"image_name": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "image_name", value)
			return value
		}(in.ImageName),
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
