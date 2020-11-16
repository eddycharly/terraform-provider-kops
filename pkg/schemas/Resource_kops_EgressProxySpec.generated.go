package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsEgressProxySpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_proxy":     RequiredStruct(ResourceKopsHTTPProxy()),
			"proxy_excludes": OptionalString(),
		},
	}
}

func ExpandResourceKopsEgressProxySpec(in map[string]interface{}) kops.EgressProxySpec {
	if in == nil {
		panic("expand EgressProxySpec failure, in is nil")
	}
	return kops.EgressProxySpec{
		HTTPProxy: func(in interface{}) kops.HTTPProxy {
			return func(in interface{}) kops.HTTPProxy {
				if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
					return kops.HTTPProxy{}
				}
				return (ExpandResourceKopsHTTPProxy(in.([]interface{})[0].(map[string]interface{})))
			}(in)
		}(in["http_proxy"]),
		ProxyExcludes: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["proxy_excludes"]),
	}
}

func FlattenResourceKopsEgressProxySpecInto(in kops.EgressProxySpec, out map[string]interface{}) {
	out["http_proxy"] = func(in kops.HTTPProxy) interface{} {
		return func(in kops.HTTPProxy) []map[string]interface{} {
			return []map[string]interface{}{FlattenResourceKopsHTTPProxy(in)}
		}(in)
	}(in.HTTPProxy)
	out["proxy_excludes"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ProxyExcludes)
}

func FlattenResourceKopsEgressProxySpec(in kops.EgressProxySpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsEgressProxySpecInto(in, out)
	return out
}
