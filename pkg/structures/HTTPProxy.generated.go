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
			value := string(ExpandString(in))
			return value
		}(in["host"]),
		Port: func(in interface{}) int {
			value := int(ExpandInt(in))
			return value
		}(in["port"]),
	}
}

func FlattenHTTPProxy(in kops.HTTPProxy) map[string]interface{} {
	return map[string]interface{}{
		"host": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Host),
		"port": func(in int) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.Port),
	}
}
