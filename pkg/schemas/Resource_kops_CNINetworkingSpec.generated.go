package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsCNINetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"uses_secondary_ip": OptionalBool(),
		},
	}
}

func ExpandResourceKopsCNINetworkingSpec(in map[string]interface{}) kops.CNINetworkingSpec {
	if in == nil {
		panic("expand CNINetworkingSpec failure, in is nil")
	}
	return kops.CNINetworkingSpec{
		UsesSecondaryIP: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["uses_secondary_ip"]),
	}
}

func FlattenResourceKopsCNINetworkingSpecInto(in kops.CNINetworkingSpec, out map[string]interface{}) {
	out["uses_secondary_ip"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.UsesSecondaryIP)
}

func FlattenResourceKopsCNINetworkingSpec(in kops.CNINetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsCNINetworkingSpecInto(in, out)
	return out
}
