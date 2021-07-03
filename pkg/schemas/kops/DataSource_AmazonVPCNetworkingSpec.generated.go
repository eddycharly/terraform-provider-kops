package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAmazonVPCNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image_name":      Computed(String()),
			"init_image_name": Computed(String()),
			"env":             Computed(List(Struct(DataSourceEnvVar()))),
		},
	}
}

func ExpandDataSourceAmazonVPCNetworkingSpec(in map[string]interface{}) kops.AmazonVPCNetworkingSpec {
	if in == nil {
		panic("expand AmazonVPCNetworkingSpec failure, in is nil")
	}
	out := kops.AmazonVPCNetworkingSpec{}
	if in, ok := in["image_name"]; ok && in != nil {
		out.ImageName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["init_image_name"]; ok && in != nil {
		out.InitImageName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["env"]; ok && in != nil {
		out.Env = func(in interface{}) []kops.EnvVar {
			var out []kops.EnvVar
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.EnvVar {
					if in == nil {
						return kops.EnvVar{}
					}
					return ExpandDataSourceEnvVar(in.(map[string]interface{}))
				}(in))
			}
			return out
		}(in)
	}
	return out
}
