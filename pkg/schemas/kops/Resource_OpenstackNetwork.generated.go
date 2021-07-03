package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceOpenstackNetwork() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone_hints": Optional(List(Ptr(String()))),
		},
	}
}

func ExpandResourceOpenstackNetwork(in map[string]interface{}) kops.OpenstackNetwork {
	if in == nil {
		panic("expand OpenstackNetwork failure, in is nil")
	}
	out := kops.OpenstackNetwork{}
	if in, ok := in["availability_zone_hints"]; ok && in != nil {
		out.AvailabilityZoneHints = func(in interface{}) []*string {
			var out []*string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) *string {
					if in == nil {
						return nil
					}
					return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
				}(in))
			}
			return out
		}(in)
	}
	return out
}
