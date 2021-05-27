package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceOpenstackRouter() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external_network":        ComputedString(),
			"dns_servers":             ComputedString(),
			"external_subnet":         ComputedString(),
			"availability_zone_hints": ComputedList(String()),
		},
	}
}

func ExpandDataSourceOpenstackRouter(in map[string]interface{}) kops.OpenstackRouter {
	if in == nil {
		panic("expand OpenstackRouter failure, in is nil")
	}
	return kops.OpenstackRouter{
		ExternalNetwork: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["external_network"]),
		DNSServers: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["dns_servers"]),
		ExternalSubnet: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["external_subnet"]),
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

func FlattenDataSourceOpenstackRouterInto(in kops.OpenstackRouter, out map[string]interface{}) {
	out["external_network"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.ExternalNetwork)
	out["dns_servers"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.DNSServers)
	out["external_subnet"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.ExternalSubnet)
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

func FlattenDataSourceOpenstackRouter(in kops.OpenstackRouter) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceOpenstackRouterInto(in, out)
	return out
}
