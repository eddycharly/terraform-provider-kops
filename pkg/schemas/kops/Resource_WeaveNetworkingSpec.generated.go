package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceWeaveNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mtu":                Optional(Nullable(Int())),
			"conn_limit":         Optional(Nullable(Int())),
			"no_masq_local":      Optional(Nullable(Int())),
			"memory_request":     Optional(Nullable(Quantity())),
			"cpu_request":        Optional(Nullable(Quantity())),
			"memory_limit":       Optional(Nullable(Quantity())),
			"cpu_limit":          Optional(Nullable(Quantity())),
			"net_extra_args":     Optional(String()),
			"npc_memory_request": Optional(Nullable(Quantity())),
			"npccpu_request":     Optional(Nullable(Quantity())),
			"npc_memory_limit":   Optional(Nullable(Quantity())),
			"npccpu_limit":       Optional(Nullable(Quantity())),
			"npc_extra_args":     Optional(String()),
			"version":            Optional(String()),
		},
	}
}

func ExpandResourceWeaveNetworkingSpec(in map[string]interface{}) kops.WeaveNetworkingSpec {
	if in == nil {
		panic("expand WeaveNetworkingSpec failure, in is nil")
	}
	out := kops.WeaveNetworkingSpec{}
	if in, ok := in["mtu"]; ok && in != nil {
		out.MTU = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["conn_limit"]; ok && in != nil {
		out.ConnLimit = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["no_masq_local"]; ok && in != nil {
		out.NoMasqLocal = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
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
	if in, ok := in["memory_limit"]; ok && in != nil {
		out.MemoryLimit = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["cpu_limit"]; ok && in != nil {
		out.CPULimit = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["net_extra_args"]; ok && in != nil {
		out.NetExtraArgs = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["npc_memory_request"]; ok && in != nil {
		out.NPCMemoryRequest = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["npccpu_request"]; ok && in != nil {
		out.NPCCPURequest = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["npc_memory_limit"]; ok && in != nil {
		out.NPCMemoryLimit = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["npccpu_limit"]; ok && in != nil {
		out.NPCCPULimit = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["npc_extra_args"]; ok && in != nil {
		out.NPCExtraArgs = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["version"]; ok && in != nil {
		out.Version = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}

func FlattenResourceWeaveNetworkingSpecInto(in kops.WeaveNetworkingSpec, out map[string]interface{}) {
	out["mtu"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.MTU)
	out["conn_limit"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.ConnLimit)
	out["no_masq_local"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.NoMasqLocal)
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
	out["memory_limit"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.MemoryLimit)
	out["cpu_limit"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.CPULimit)
	out["net_extra_args"] = func(in string) interface{} { return string(in) }(in.NetExtraArgs)
	out["npc_memory_request"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.NPCMemoryRequest)
	out["npccpu_request"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.NPCCPURequest)
	out["npc_memory_limit"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.NPCMemoryLimit)
	out["npccpu_limit"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.NPCCPULimit)
	out["npc_extra_args"] = func(in string) interface{} { return string(in) }(in.NPCExtraArgs)
	out["version"] = func(in string) interface{} { return string(in) }(in.Version)
}

func FlattenResourceWeaveNetworkingSpec(in kops.WeaveNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceWeaveNetworkingSpecInto(in, out)
	return out
}
