package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceEgressProxySpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_proxy":     ComputedStruct(DataSourceHTTPProxy()),
			"proxy_excludes": ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceEgressProxySpec(in map[string]interface{}) kops.EgressProxySpec {
	if in == nil {
		panic("expand EgressProxySpec failure, in is nil")
	}
	return kops.EgressProxySpec{
		HTTPProxy: func(in interface{}) kops.HTTPProxy {
			return func(in interface{}) kops.HTTPProxy {
				if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
					return kops.HTTPProxy{}
				}
				return (ExpandDataSourceHTTPProxy(in.([]interface{})[0].(map[string]interface{})))
			}(in)
		}(in["http_proxy"]),
		ProxyExcludes: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["proxy_excludes"]),
	}
}

func FlattenDataSourceEgressProxySpecInto(in kops.EgressProxySpec, out map[string]interface{}) {
	out["http_proxy"] = func(in kops.HTTPProxy) interface{} {
		return func(in kops.HTTPProxy) []interface{} {
			return []interface{}{FlattenDataSourceHTTPProxy(in)}
		}(in)
	}(in.HTTPProxy)
	out["proxy_excludes"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ProxyExcludes)
}

func FlattenDataSourceEgressProxySpec(in kops.EgressProxySpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceEgressProxySpecInto(in, out)
	return out
}
