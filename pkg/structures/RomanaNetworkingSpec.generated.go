package structures

import (
	"log"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandRomanaNetworkingSpec(in map[string]interface{}) kops.RomanaNetworkingSpec {
	if in == nil {
		panic("expand RomanaNetworkingSpec failure, in is nil")
	}
	return kops.RomanaNetworkingSpec{
		DaemonServiceIP: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "daemon_service_ip", value)
			return value
		}(in["daemon_service_ip"]),
		EtcdServiceIP: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "etcd_service_ip", value)
			return value
		}(in["etcd_service_ip"]),
	}
}

func FlattenRomanaNetworkingSpec(in kops.RomanaNetworkingSpec) map[string]interface{} {
	return map[string]interface{}{
		"daemon_service_ip": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "daemon_service_ip", value)
			return value
		}(in.DaemonServiceIP),
		"etcd_service_ip": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "etcd_service_ip", value)
			return value
		}(in.EtcdServiceIP),
	}
}
