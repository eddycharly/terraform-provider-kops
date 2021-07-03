package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceEgressProxySpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_proxy":     Required(Struct(ResourceHTTPProxy())),
			"proxy_excludes": Optional(String()),
		},
	}
}

func ExpandResourceEgressProxySpec(in map[string]interface{}) kops.EgressProxySpec {
	if in == nil {
		panic("expand EgressProxySpec failure, in is nil")
	}
	out := kops.EgressProxySpec{}
	if in, ok := in["http_proxy"]; ok && in != nil {
		out.HTTPProxy = func(in interface{}) kops.HTTPProxy {
			if in == nil {
				return kops.HTTPProxy{}
			}
			return ExpandResourceHTTPProxy(in.(map[string]interface{}))
		}(in)
	}
	if in, ok := in["proxy_excludes"]; ok && in != nil {
		out.ProxyExcludes = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}
