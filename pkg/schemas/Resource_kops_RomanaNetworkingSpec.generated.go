package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsRomanaNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"daemon_service_ip": OptionalString(),
			"etcd_service_ip":   OptionalString(),
		},
	}
}

func ExpandResourceKopsRomanaNetworkingSpec(in map[string]interface{}) kops.RomanaNetworkingSpec {
	if in == nil {
		panic("expand RomanaNetworkingSpec failure, in is nil")
	}
	return kops.RomanaNetworkingSpec{
		DaemonServiceIP: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["daemon_service_ip"]),
		EtcdServiceIP: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["etcd_service_ip"]),
	}
}

func FlattenResourceKopsRomanaNetworkingSpecInto(in kops.RomanaNetworkingSpec, out map[string]interface{}) {
	out["daemon_service_ip"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.DaemonServiceIP)
	out["etcd_service_ip"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.EtcdServiceIP)
}

func FlattenResourceKopsRomanaNetworkingSpec(in kops.RomanaNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsRomanaNetworkingSpecInto(in, out)
	return out
}
