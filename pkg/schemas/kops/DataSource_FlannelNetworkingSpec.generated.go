package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceFlannelNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backend":                        Computed(String()),
			"disable_tx_checksum_offloading": Computed(Bool()),
			"iptables_resync_seconds":        Computed(Nullable(Int())),
		},
	}
}

func ExpandDataSourceFlannelNetworkingSpec(in map[string]interface{}) kops.FlannelNetworkingSpec {
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

func FlattenDataSourceFlannelNetworkingSpecInto(in kops.FlannelNetworkingSpec, out map[string]interface{}) {
	out["backend"] = func(in string) interface{} { return string(in) }(in.Backend)
	out["disable_tx_checksum_offloading"] = func(in bool) interface{} { return in }(in.DisableTxChecksumOffloading)
	out["iptables_resync_seconds"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.IptablesResyncSeconds)
}

func FlattenDataSourceFlannelNetworkingSpec(in kops.FlannelNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceFlannelNetworkingSpecInto(in, out)
	return out
}
