package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
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

func ExpandResourceEgressProxySpec(in map[string]interface{}) kopsv1alpha2.EgressProxySpec {
	if in == nil {
		panic("expand EgressProxySpec failure, in is nil")
	}
	return kopsv1alpha2.EgressProxySpec{
		HTTPProxy: func(in interface{}) kopsv1alpha2.HTTPProxy {
			return func(in interface{}) kopsv1alpha2.HTTPProxy {
				if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
					return ExpandResourceHTTPProxy(in[0].(map[string]interface{}))
				}
				return kopsv1alpha2.HTTPProxy{}
			}(in)
		}(in["http_proxy"]),
		ProxyExcludes: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["proxy_excludes"]),
	}
}

func FlattenResourceEgressProxySpecInto(in kopsv1alpha2.EgressProxySpec, out map[string]interface{}) {
	out["http_proxy"] = func(in kopsv1alpha2.HTTPProxy) interface{} {
		return func(in kopsv1alpha2.HTTPProxy) []interface{} {
			return []interface{}{FlattenResourceHTTPProxy(in)}
		}(in)
	}(in.HTTPProxy)
	out["proxy_excludes"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ProxyExcludes)
}

func FlattenResourceEgressProxySpec(in kopsv1alpha2.EgressProxySpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceEgressProxySpecInto(in, out)
	return out
}
