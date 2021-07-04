package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceHTTPProxy() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": Required(String()),
			"port": Required(Int()),
		},
	}
}

func ExpandResourceHTTPProxy(in map[string]interface{}) kops.HTTPProxy {
	if in == nil {
		panic("expand HTTPProxy failure, in is nil")
	}
	out := kops.HTTPProxy{}
	if in, ok := in["host"]; ok && in != nil {
		out.Host = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["port"]; ok && in != nil {
		out.Port = func(in interface{}) int { return int(in.(int)) }(in)
	}
	return out
}

func FlattenResourceHTTPProxyInto(in kops.HTTPProxy, out map[string]interface{}) {
	out["host"] = func(in string) interface{} { return string(in) }(in.Host)
	out["port"] = func(in int) interface{} { return int(in) }(in.Port)
}

func FlattenResourceHTTPProxy(in kops.HTTPProxy) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceHTTPProxyInto(in, out)
	return out
}
