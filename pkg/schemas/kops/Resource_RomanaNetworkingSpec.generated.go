package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceRomanaNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"daemon_service_ip": Optional(String()),
			"etcd_service_ip":   Optional(String()),
		},
	}
}

func ExpandResourceRomanaNetworkingSpec(in map[string]interface{}) kops.RomanaNetworkingSpec {
	if in == nil {
		panic("expand RomanaNetworkingSpec failure, in is nil")
	}
	out := kops.RomanaNetworkingSpec{}
	if in, ok := in["daemon_service_ip"]; ok && in != nil {
		out.DaemonServiceIP = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["etcd_service_ip"]; ok && in != nil {
		out.EtcdServiceIP = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}

func FlattenResourceRomanaNetworkingSpecInto(in kops.RomanaNetworkingSpec, out map[string]interface{}) {
	out["daemon_service_ip"] = func(in string) interface{} { return string(in) }(in.DaemonServiceIP)
	out["etcd_service_ip"] = func(in string) interface{} { return string(in) }(in.EtcdServiceIP)
}

func FlattenResourceRomanaNetworkingSpec(in kops.RomanaNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceRomanaNetworkingSpecInto(in, out)
	return out
}
