package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceOpenstackRouter() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external_network":        Optional(Nullable(String())),
			"dns_servers":             Optional(Nullable(String())),
			"external_subnet":         Optional(Nullable(String())),
			"availability_zone_hints": Optional(List(String())),
		},
	}
}

func ExpandResourceOpenstackRouter(in map[string]interface{}) kops.OpenstackRouter {
	if in == nil {
		panic("expand OpenstackRouter failure, in is nil")
	}
	out := kops.OpenstackRouter{}
	if in, ok := in["external_network"]; ok && in != nil {
		out.ExternalNetwork = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["dns_servers"]; ok && in != nil {
		out.DNSServers = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["external_subnet"]; ok && in != nil {
		out.ExternalSubnet = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
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
					return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in))
				}(in))
			}
			return out
		}(in)
	}
	return out
}

func FlattenResourceOpenstackRouterInto(in kops.OpenstackRouter, out map[string]interface{}) {
	out["external_network"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.ExternalNetwork)
	out["dns_servers"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.DNSServers)
	out["external_subnet"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.ExternalSubnet)
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

func FlattenResourceOpenstackRouter(in kops.OpenstackRouter) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceOpenstackRouterInto(in, out)
	return out
}
