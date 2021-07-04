package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceOpenstackNetwork() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone_hints": Computed(List(String())),
		},
	}
}

func ExpandDataSourceOpenstackNetwork(in map[string]interface{}) kops.OpenstackNetwork {
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

func FlattenDataSourceOpenstackNetworkInto(in kops.OpenstackNetwork, out map[string]interface{}) {
	out["availability_zone_hints"] = func(in []*string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
			}(in))
		}
		return out
	}(in.AvailabilityZoneHints)
}

func FlattenDataSourceOpenstackNetwork(in kops.OpenstackNetwork) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceOpenstackNetworkInto(in, out)
	return out
}
