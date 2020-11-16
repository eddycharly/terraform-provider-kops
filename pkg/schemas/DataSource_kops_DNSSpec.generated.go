package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsDNSSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": ComputedString(),
		},
	}
}

func ExpandDataSourceKopsDNSSpec(in map[string]interface{}) kops.DNSSpec {
	if in == nil {
		panic("expand DNSSpec failure, in is nil")
	}
	return kops.DNSSpec{
		Type: func(in interface{}) kops.DNSType {
			return kops.DNSType(ExpandString(in))
		}(in["type"]),
	}
}

func FlattenDataSourceKopsDNSSpecInto(in kops.DNSSpec, out map[string]interface{}) {
	out["type"] = func(in kops.DNSType) interface{} {
		return FlattenString(string(in))
	}(in.Type)
}

func FlattenDataSourceKopsDNSSpec(in kops.DNSSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsDNSSpecInto(in, out)
	return out
}
