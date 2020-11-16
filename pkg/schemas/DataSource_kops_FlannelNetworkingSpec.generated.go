package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsFlannelNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backend":                        ComputedString(),
			"disable_tx_checksum_offloading": ComputedBool(),
			"iptables_resync_seconds":        ComputedInt(),
		},
	}
}

func ExpandDataSourceKopsFlannelNetworkingSpec(in map[string]interface{}) kops.FlannelNetworkingSpec {
	if in == nil {
		panic("expand FlannelNetworkingSpec failure, in is nil")
	}
	return kops.FlannelNetworkingSpec{
		Backend: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["backend"]),
		DisableTxChecksumOffloading: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disable_tx_checksum_offloading"]),
		IptablesResyncSeconds: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["iptables_resync_seconds"]),
	}
}

func FlattenDataSourceKopsFlannelNetworkingSpecInto(in kops.FlannelNetworkingSpec, out map[string]interface{}) {
	out["backend"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Backend)
	out["disable_tx_checksum_offloading"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.DisableTxChecksumOffloading)
	out["iptables_resync_seconds"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.IptablesResyncSeconds)
}

func FlattenDataSourceKopsFlannelNetworkingSpec(in kops.FlannelNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsFlannelNetworkingSpecInto(in, out)
	return out
}
