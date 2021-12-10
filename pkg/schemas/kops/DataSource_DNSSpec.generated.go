package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceDNSSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceDNSSpec(in map[string]interface{}) kops.DNSSpec {
	if in == nil {
		panic("expand DNSSpec failure, in is nil")
	}
	return kops.DNSSpec{
		Type: func(in interface{}) kops.DNSType {
			return kops.DNSType(ExpandString(in))
		}(in["type"]),
	}
}

func FlattenDataSourceDNSSpecInto(in kops.DNSSpec, out map[string]interface{}) {
	out["type"] = func(in kops.DNSType) interface{} {
		return FlattenString(string(in))
	}(in.Type)
}

func FlattenDataSourceDNSSpec(in kops.DNSSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceDNSSpecInto(in, out)
	return out
}
