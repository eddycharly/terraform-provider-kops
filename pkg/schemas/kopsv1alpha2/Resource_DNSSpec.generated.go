package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops/v1alpha2"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceDNSSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": RequiredString(),
		},
	}

	return res
}

func ExpandResourceDNSSpec(in map[string]interface{}) kopsv1alpha2.DNSSpec {
	if in == nil {
		panic("expand DNSSpec failure, in is nil")
	}
	return kopsv1alpha2.DNSSpec{
		Type: func(in interface{}) kopsv1alpha2.DNSType {
			return v1alpha2.DNSType(ExpandString(in))
		}(in["type"]),
	}
}

func FlattenResourceDNSSpecInto(in kopsv1alpha2.DNSSpec, out map[string]interface{}) {
	out["type"] = func(in kopsv1alpha2.DNSType) interface{} {
		return FlattenString(string(in))
	}(in.Type)
}

func FlattenResourceDNSSpec(in kopsv1alpha2.DNSSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceDNSSpecInto(in, out)
	return out
}
