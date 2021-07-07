package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceAmazonVPCNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image_name":      OptionalString(),
			"init_image_name": OptionalString(),
			"env":             OptionalList(ResourceEnvVar()),
		},
	}

	return res
}

func ExpandResourceAmazonVPCNetworkingSpec(in map[string]interface{}) kops.AmazonVPCNetworkingSpec {
	if in == nil {
		panic("expand AmazonVPCNetworkingSpec failure, in is nil")
	}
	return kops.AmazonVPCNetworkingSpec{
		ImageName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image_name"]),
		InitImageName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["init_image_name"]),
		Env: func(in interface{}) []kops.EnvVar {
			return func(in interface{}) []kops.EnvVar {
				if in == nil {
					return nil
				}
				var out []kops.EnvVar
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.EnvVar {
						if in == nil {
							return kops.EnvVar{}
						}
						return (ExpandResourceEnvVar(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["env"]),
	}
}

func FlattenResourceAmazonVPCNetworkingSpecInto(in kops.AmazonVPCNetworkingSpec, out map[string]interface{}) {
	out["image_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ImageName)
	out["init_image_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.InitImageName)
	out["env"] = func(in []kops.EnvVar) interface{} {
		return func(in []kops.EnvVar) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.EnvVar) interface{} {
					return FlattenResourceEnvVar(in)
				}(in))
			}
			return out
		}(in)
	}(in.Env)
}

func FlattenResourceAmazonVPCNetworkingSpec(in kops.AmazonVPCNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAmazonVPCNetworkingSpecInto(in, out)
	return out
}
