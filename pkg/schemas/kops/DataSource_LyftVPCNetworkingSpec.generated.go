package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceLyftVPCNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subnet_tags": Computed(Map(String())),
		},
	}
}

func ExpandDataSourceLyftVPCNetworkingSpec(in map[string]interface{}) kops.LyftVPCNetworkingSpec {
	if in == nil {
		panic("expand LyftVPCNetworkingSpec failure, in is nil")
	}
	out := kops.LyftVPCNetworkingSpec{}
	if in, ok := in["subnet_tags"]; ok && in != nil {
		out.SubnetTags = func(in interface{}) map[string]string {
			if in == nil {
				return nil
			}
			out := map[string]string{}
			for key, in := range in.(map[string]interface{}) {
				out[key] = func(in interface{}) string { return string(in.(string)) }(in)
			}
			return out
		}(in)
	}
	return out
}

func FlattenDataSourceLyftVPCNetworkingSpecInto(in kops.LyftVPCNetworkingSpec, out map[string]interface{}) {
	out["subnet_tags"] = func(in map[string]string) interface{} {
		if in == nil {
			return nil
		}
		out := map[string]interface{}{}
		for key, in := range in {
			out[key] = func(in string) interface{} { return string(in) }(in)
		}
		return out
	}(in.SubnetTags)
}

func FlattenDataSourceLyftVPCNetworkingSpec(in kops.LyftVPCNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceLyftVPCNetworkingSpecInto(in, out)
	return out
}
