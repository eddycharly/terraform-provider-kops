package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsOpenstackRouter() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external_network": ComputedString(),
			"dns_servers":      ComputedString(),
			"external_subnet":  ComputedString(),
		},
	}
}

func ExpandDataSourceKopsOpenstackRouter(in map[string]interface{}) kops.OpenstackRouter {
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
	}
}

func FlattenDataSourceKopsOpenstackRouterInto(in kops.OpenstackRouter, out map[string]interface{}) {
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
}

func FlattenDataSourceKopsOpenstackRouter(in kops.OpenstackRouter) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsOpenstackRouterInto(in, out)
	return out
}
