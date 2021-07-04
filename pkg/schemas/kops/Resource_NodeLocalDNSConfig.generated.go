package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceNodeLocalDNSConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":             Optional(Nullable(Bool())),
			"local_ip":            Optional(String()),
			"forward_to_kube_dns": Optional(Nullable(Bool())),
			"memory_request":      Optional(Nullable(Quantity())),
			"cpu_request":         Optional(Nullable(Quantity())),
		},
	}
}

func ExpandResourceNodeLocalDNSConfig(in map[string]interface{}) kops.NodeLocalDNSConfig {
	if in == nil {
		panic("expand NodeLocalDNSConfig failure, in is nil")
	}
	out := kops.NodeLocalDNSConfig{}
	if in, ok := in["enabled"]; ok && in != nil {
		out.Enabled = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["memory_request"]; ok && in != nil {
		out.MemoryRequest = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["cpu_request"]; ok && in != nil {
		out.CPURequest = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenResourceNodeLocalDNSConfigInto(in kops.NodeLocalDNSConfig, out map[string]interface{}) {
	out["enabled"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.Enabled)
	out["local_ip"] = func(in string) interface{} { return string(in) }(in.LocalIP)
	out["forward_to_kube_dns"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.ForwardToKubeDNS)
	out["memory_request"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.MemoryRequest)
	out["cpu_request"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.CPURequest)
}

func FlattenResourceNodeLocalDNSConfig(in kops.NodeLocalDNSConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceNodeLocalDNSConfigInto(in, out)
	return out
}
