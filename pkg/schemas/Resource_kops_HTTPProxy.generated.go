package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsHTTPProxy() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": RequiredString(),
			"port": RequiredInt(),
		},
	}
}

func ExpandResourceKopsHTTPProxy(in map[string]interface{}) kops.HTTPProxy {
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

func FlattenResourceKopsHTTPProxyInto(in kops.HTTPProxy, out map[string]interface{}) {
	out["host"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Host)
	out["port"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.Port)
}

func FlattenResourceKopsHTTPProxy(in kops.HTTPProxy) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsHTTPProxyInto(in, out)
	return out
}
