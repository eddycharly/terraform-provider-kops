package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceNodeLocalDNSConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":             Computed(Ptr(Bool())),
			"local_ip":            Computed(String()),
			"forward_to_kube_dns": Computed(Ptr(Bool())),
			"memory_request":      Computed(Ptr(Quantity())),
			"cpu_request":         Computed(Ptr(Quantity())),
		},
	}
}

func ExpandDataSourceNodeLocalDNSConfig(in map[string]interface{}) kops.NodeLocalDNSConfig {
	if in == nil {
		panic("expand NodeLocalDNSConfig failure, in is nil")
	}
	out := kops.NodeLocalDNSConfig{}
	if in, ok := in["enabled"]; ok && in != nil {
		out.Enabled = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["local_ip"]; ok && in != nil {
		out.LocalIP = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["forward_to_kube_dns"]; ok && in != nil {
		out.ForwardToKubeDNS = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["memory_request"]; ok && in != nil {
		out.MemoryRequest = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["cpu_request"]; ok && in != nil {
		out.CPURequest = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
