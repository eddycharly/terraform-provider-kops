package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandRomanaNetworkingSpec(in map[string]interface{}) kops.RomanaNetworkingSpec {
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

func FlattenRomanaNetworkingSpec(in kops.RomanaNetworkingSpec) map[string]interface{} {
	return map[string]interface{}{
		"daemon_service_ip": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.DaemonServiceIP),
		"etcd_service_ip": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.EtcdServiceIP),
	}
}
