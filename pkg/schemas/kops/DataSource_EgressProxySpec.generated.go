package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceEgressProxySpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_proxy":     Computed(Struct(DataSourceHTTPProxy())),
			"proxy_excludes": Computed(String()),
		},
	}
}

func ExpandDataSourceEgressProxySpec(in map[string]interface{}) kops.EgressProxySpec {
	if in == nil {
		panic("expand EgressProxySpec failure, in is nil")
	}
	out := kops.EgressProxySpec{}
	if in, ok := in["http_proxy"]; ok && in != nil {
		out.HTTPProxy = func(in interface{}) kops.HTTPProxy {
			if in == nil {
				return kops.HTTPProxy{}
			}
			return ExpandDataSourceHTTPProxy(in.(map[string]interface{}))
		}(in)
	}
	if in, ok := in["proxy_excludes"]; ok && in != nil {
		out.ProxyExcludes = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}
