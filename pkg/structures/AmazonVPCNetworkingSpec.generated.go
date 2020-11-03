package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandAmazonVPCNetworkingSpec(in map[string]interface{}) kops.AmazonVPCNetworkingSpec {
	if in == nil {
		panic("expand AmazonVPCNetworkingSpec failure, in is nil")
	}
	return kops.AmazonVPCNetworkingSpec{
		ImageName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image_name"]),
		Env: func(in interface{}) []kops.EnvVar {
			return func(in interface{}) []kops.EnvVar {
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
		}(in["env"]),
	}
}

func FlattenAmazonVPCNetworkingSpec(in kops.AmazonVPCNetworkingSpec) map[string]interface{} {
	return map[string]interface{}{
		"image_name": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ImageName),
		"env": func(in []kops.EnvVar) interface{} {
			return func(in []kops.EnvVar) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.EnvVar) interface{} {
						return FlattenEnvVar(in)
					}(in))
				}
				return out
			}(in)
		}(in.Env),
	}
}
