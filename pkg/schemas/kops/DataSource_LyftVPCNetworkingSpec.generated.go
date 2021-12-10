package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceLyftVPCNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subnet_tags": ComputedMap(String()),
		},
	}

	return res
}

func ExpandDataSourceLyftVPCNetworkingSpec(in map[string]interface{}) kops.LyftVPCNetworkingSpec {
	if in == nil {
		panic("expand LyftVPCNetworkingSpec failure, in is nil")
	}
	return kops.LyftVPCNetworkingSpec{
		SubnetTags: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string]string{}
						for key, in := range in {
							out[key] = string(ExpandString(in))
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["subnet_tags"]),
	}
}

func FlattenDataSourceLyftVPCNetworkingSpecInto(in kops.LyftVPCNetworkingSpec, out map[string]interface{}) {
	out["subnet_tags"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.SubnetTags)
}

func FlattenDataSourceLyftVPCNetworkingSpec(in kops.LyftVPCNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceLyftVPCNetworkingSpecInto(in, out)
	return out
}
