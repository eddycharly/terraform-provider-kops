package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsAmazonVPCNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image_name": OptionalString(),
			"env":        OptionalList(ResourceKopsEnvVar()),
		},
	}
}

func ExpandResourceKopsAmazonVPCNetworkingSpec(in map[string]interface{}) kops.AmazonVPCNetworkingSpec {
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
						return (ExpandResourceKopsEnvVar(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["env"]),
	}
}

func FlattenResourceKopsAmazonVPCNetworkingSpecInto(in kops.AmazonVPCNetworkingSpec, out map[string]interface{}) {
	out["image_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ImageName)
	out["env"] = func(in []kops.EnvVar) interface{} {
		return func(in []kops.EnvVar) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.EnvVar) interface{} {
					return FlattenResourceKopsEnvVar(in)
				}(in))
			}
			return out
		}(in)
	}(in.Env)
}

func FlattenResourceKopsAmazonVPCNetworkingSpec(in kops.AmazonVPCNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsAmazonVPCNetworkingSpecInto(in, out)
	return out
}
