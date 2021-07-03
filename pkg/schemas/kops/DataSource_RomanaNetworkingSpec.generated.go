package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceRomanaNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"daemon_service_ip": Computed(String()),
			"etcd_service_ip":   Computed(String()),
		},
	}
}

func ExpandDataSourceRomanaNetworkingSpec(in map[string]interface{}) kops.RomanaNetworkingSpec {
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
