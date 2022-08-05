package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceEgressProxySpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_proxy":     RequiredStruct(ResourceHTTPProxy()),
			"proxy_excludes": OptionalString(),
		},
	}

	return res
}

func ExpandResourceEgressProxySpec(in map[string]interface{}) kops.EgressProxySpec {
	if in == nil {
		panic("expand EgressProxySpec failure, in is nil")
	}
	return kops.EgressProxySpec{
		HTTPProxy: func(in interface{}) kops.HTTPProxy {
			return func(in interface{}) kops.HTTPProxy {
				if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
					return ExpandResourceHTTPProxy(in[0].(map[string]interface{}))
				}
				return kops.HTTPProxy{}
			}(in)
		}(in["http_proxy"]),
		ProxyExcludes: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["proxy_excludes"]),
	}
}

func FlattenResourceEgressProxySpecInto(in kops.EgressProxySpec, out map[string]interface{}) {
	out["http_proxy"] = func(in kops.HTTPProxy) interface{} {
		return func(in kops.HTTPProxy) []interface{} {
			return []interface{}{FlattenResourceHTTPProxy(in)}
		}(in)
	}(in.HTTPProxy)
	out["proxy_excludes"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ProxyExcludes)
}

func FlattenResourceEgressProxySpec(in kops.EgressProxySpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceEgressProxySpecInto(in, out)
	return out
}
