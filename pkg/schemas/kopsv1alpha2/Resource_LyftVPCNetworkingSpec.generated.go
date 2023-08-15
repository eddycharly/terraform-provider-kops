package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceLyftVPCNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subnet_tags": OptionalMap(String()),
		},
	}

	return res
}

func ExpandResourceLyftVPCNetworkingSpec(in map[string]interface{}) kopsv1alpha2.LyftVPCNetworkingSpec {
	if in == nil {
		panic("expand LyftVPCNetworkingSpec failure, in is nil")
	}
	return kopsv1alpha2.LyftVPCNetworkingSpec{
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

func FlattenResourceLyftVPCNetworkingSpecInto(in kopsv1alpha2.LyftVPCNetworkingSpec, out map[string]interface{}) {
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

func FlattenResourceLyftVPCNetworkingSpec(in kopsv1alpha2.LyftVPCNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceLyftVPCNetworkingSpecInto(in, out)
	return out
}
