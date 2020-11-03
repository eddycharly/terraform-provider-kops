package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandEgressProxySpec(in map[string]interface{}) kops.EgressProxySpec {
	if in == nil {
		panic("expand EgressProxySpec failure, in is nil")
	}
	return kops.EgressProxySpec{
		HTTPProxy: func(in interface{}) kops.HTTPProxy {
			value := func(in interface{}) kops.HTTPProxy {
				if in.([]interface{})[0] == nil {
					return kops.HTTPProxy{}
				}
				return (ExpandHTTPProxy(in.([]interface{})[0].(map[string]interface{})))
			}(in)
			return value
		}(in["http_proxy"]),
		ProxyExcludes: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["proxy_excludes"]),
	}
}

func FlattenEgressProxySpec(in kops.EgressProxySpec) map[string]interface{} {
	return map[string]interface{}{
		"http_proxy": func(in kops.HTTPProxy) interface{} {
			value := func(in kops.HTTPProxy) []map[string]interface{} {
				return []map[string]interface{}{FlattenHTTPProxy(in)}
			}(in)
			return value
		}(in.HTTPProxy),
		"proxy_excludes": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ProxyExcludes),
	}
}
