package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceOpenstackRouter() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external_network":        Computed(Ptr(String())),
			"dns_servers":             Computed(Ptr(String())),
			"external_subnet":         Computed(Ptr(String())),
			"availability_zone_hints": Computed(List(Ptr(String()))),
		},
	}
}

func ExpandDataSourceOpenstackRouter(in map[string]interface{}) kops.OpenstackRouter {
	if in == nil {
		panic("expand OpenstackRouter failure, in is nil")
	}
	out := kops.OpenstackRouter{}
	if in, ok := in["external_network"]; ok && in != nil {
		out.ExternalNetwork = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["dns_servers"]; ok && in != nil {
		out.DNSServers = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["external_subnet"]; ok && in != nil {
		out.ExternalSubnet = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
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
