package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceFlannelNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backend":                        Optional(String()),
			"disable_tx_checksum_offloading": Optional(Bool()),
			"iptables_resync_seconds":        Optional(Nullable(Int())),
		},
	}
}

func ExpandResourceFlannelNetworkingSpec(in map[string]interface{}) kops.FlannelNetworkingSpec {
	if in == nil {
		panic("expand FlannelNetworkingSpec failure, in is nil")
	}
	out := kops.FlannelNetworkingSpec{}
	if in, ok := in["backend"]; ok && in != nil {
		out.Backend = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["disable_tx_checksum_offloading"]; ok && in != nil {
		out.DisableTxChecksumOffloading = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["iptables_resync_seconds"]; ok && in != nil {
		out.IptablesResyncSeconds = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenResourceFlannelNetworkingSpecInto(in kops.FlannelNetworkingSpec, out map[string]interface{}) {
	out["backend"] = func(in string) interface{} { return string(in) }(in.Backend)
	out["disable_tx_checksum_offloading"] = func(in bool) interface{} { return in }(in.DisableTxChecksumOffloading)
	out["iptables_resync_seconds"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.IptablesResyncSeconds)
}

func FlattenResourceFlannelNetworkingSpec(in kops.FlannelNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceFlannelNetworkingSpecInto(in, out)
	return out
}
