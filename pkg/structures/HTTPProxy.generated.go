package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandHTTPProxy(in map[string]interface{}) kops.HTTPProxy {
	if in == nil {
		panic("expand HTTPProxy failure, in is nil")
	}
	return kops.HTTPProxy{
		Host: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["host"]),
		Port: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["port"]),
	}
}

func FlattenHTTPProxy(in kops.HTTPProxy) map[string]interface{} {
	return map[string]interface{}{
		"host": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Host),
		"port": func(in int) interface{} {
			return FlattenInt(int(in))
		}(in.Port),
	}
}
