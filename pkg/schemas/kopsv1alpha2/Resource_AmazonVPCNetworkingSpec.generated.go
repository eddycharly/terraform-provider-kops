package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceAmazonVPCNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image":      OptionalString(),
			"init_image": OptionalString(),
			"env":        OptionalList(ResourceEnvVar()),
		},
	}

	return res
}

func ExpandResourceAmazonVPCNetworkingSpec(in map[string]interface{}) kopsv1alpha2.AmazonVPCNetworkingSpec {
	if in == nil {
		panic("expand AmazonVPCNetworkingSpec failure, in is nil")
	}
	return kopsv1alpha2.AmazonVPCNetworkingSpec{
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		InitImage: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["init_image"]),
		Env: func(in interface{}) []kopsv1alpha2.EnvVar {
			return func(in interface{}) []kopsv1alpha2.EnvVar {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.EnvVar
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.EnvVar {
						if in == nil {
							return kopsv1alpha2.EnvVar{}
						}
						return ExpandResourceEnvVar(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["env"]),
	}
}

func FlattenResourceAmazonVPCNetworkingSpecInto(in kopsv1alpha2.AmazonVPCNetworkingSpec, out map[string]interface{}) {
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
	out["init_image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.InitImage)
	out["env"] = func(in []kopsv1alpha2.EnvVar) interface{} {
		return func(in []kopsv1alpha2.EnvVar) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.EnvVar) interface{} {
					return FlattenResourceEnvVar(in)
				}(in))
			}
			return out
		}(in)
	}(in.Env)
}

func FlattenResourceAmazonVPCNetworkingSpec(in kopsv1alpha2.AmazonVPCNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAmazonVPCNetworkingSpecInto(in, out)
	return out
}
