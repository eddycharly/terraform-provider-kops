package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func LyftVPCNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subnet_tags": OptionalMap(String()),
		},
	}
}

func ExpandLyftVPCNetworkingSpec(in map[string]interface{}) kops.LyftVPCNetworkingSpec {
	if in == nil {
		panic("expand LyftVPCNetworkingSpec failure, in is nil")
	}
	return kops.LyftVPCNetworkingSpec{
		SubnetTags: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
		}(in["subnet_tags"]),
	}
}

func FlattenLyftVPCNetworkingSpecInto(in kops.LyftVPCNetworkingSpec, out map[string]interface{}) {
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

func FlattenLyftVPCNetworkingSpec(in kops.LyftVPCNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenLyftVPCNetworkingSpecInto(in, out)
	return out
}
