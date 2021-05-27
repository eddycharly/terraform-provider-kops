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
			"availability_zone_hints": OptionalList(String()),
		},
	}
}

func ExpandResourceOpenstackNetwork(in map[string]interface{}) kops.OpenstackNetwork {
	if in == nil {
		panic("expand OpenstackNetwork failure, in is nil")
	}
	return kops.OpenstackNetwork{
		AvailabilityZoneHints: func(in interface{}) []*string {
			return func(in interface{}) []*string {
				var out []*string
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *string {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
							return nil
						}
						return func(in string) *string {
							return &in
						}(string(ExpandString(in)))
					}(in))
				}
				return out
			}(in)
		}(in["availability_zone_hints"]),
	}
}

func FlattenResourceOpenstackNetworkInto(in kops.OpenstackNetwork, out map[string]interface{}) {
	out["availability_zone_hints"] = func(in []*string) interface{} {
		return func(in []*string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in *string) interface{} {
					if in == nil {
						return nil
					}
					return func(in string) interface{} {
						return FlattenString(string(in))
					}(*in)
				}(in))
			}
			return out
		}(in)
	}(in.AvailabilityZoneHints)
}

func FlattenResourceOpenstackNetwork(in kops.OpenstackNetwork) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceOpenstackNetworkInto(in, out)
	return out
}
