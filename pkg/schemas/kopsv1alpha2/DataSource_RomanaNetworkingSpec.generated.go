package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceRomanaNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"daemon_service_ip": ComputedString(),
			"etcd_service_ip":   ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceRomanaNetworkingSpec(in map[string]interface{}) kopsv1alpha2.RomanaNetworkingSpec {
	if in == nil {
		panic("expand RomanaNetworkingSpec failure, in is nil")
	}
	return kopsv1alpha2.RomanaNetworkingSpec{
		DaemonServiceIP: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["daemon_service_ip"]),
		EtcdServiceIP: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["etcd_service_ip"]),
	}
}

func FlattenDataSourceRomanaNetworkingSpecInto(in kopsv1alpha2.RomanaNetworkingSpec, out map[string]interface{}) {
	out["daemon_service_ip"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.DaemonServiceIP)
	out["etcd_service_ip"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.EtcdServiceIP)
}

func FlattenDataSourceRomanaNetworkingSpec(in kopsv1alpha2.RomanaNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceRomanaNetworkingSpecInto(in, out)
	return out
}
