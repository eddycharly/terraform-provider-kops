package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsDNSSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": RequiredString(),
		},
	}
}

func ExpandResourceKopsDNSSpec(in map[string]interface{}) kops.DNSSpec {
	if in == nil {
		panic("expand DNSSpec failure, in is nil")
	}
	return kops.DNSSpec{
		Type: func(in interface{}) kops.DNSType {
			return kops.DNSType(ExpandString(in))
		}(in["type"]),
	}
}

func FlattenResourceKopsDNSSpecInto(in kops.DNSSpec, out map[string]interface{}) {
	out["type"] = func(in kops.DNSType) interface{} {
		return FlattenString(string(in))
	}(in.Type)
}

func FlattenResourceKopsDNSSpec(in kops.DNSSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsDNSSpecInto(in, out)
	return out
}
