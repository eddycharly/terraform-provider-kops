package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceHTTPProxy() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": ComputedString(),
			"port": ComputedInt(),
		},
	}
}

func ExpandDataSourceHTTPProxy(in map[string]interface{}) kops.HTTPProxy {
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

func FlattenDataSourceHTTPProxyInto(in kops.HTTPProxy, out map[string]interface{}) {
	out["host"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Host)
	out["port"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.Port)
}

func FlattenDataSourceHTTPProxy(in kops.HTTPProxy) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceHTTPProxyInto(in, out)
	return out
}
