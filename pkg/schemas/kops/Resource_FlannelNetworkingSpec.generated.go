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
			"iptables_resync_seconds":        Optional(Ptr(Int())),
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
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
