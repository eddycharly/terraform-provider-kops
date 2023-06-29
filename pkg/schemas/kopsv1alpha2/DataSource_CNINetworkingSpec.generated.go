package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceCNINetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"uses_secondary_ip": ComputedBool(),
		},
	}

	return res
}

func ExpandDataSourceCNINetworkingSpec(in map[string]interface{}) kopsv1alpha2.CNINetworkingSpec {
	if in == nil {
		panic("expand CNINetworkingSpec failure, in is nil")
	}
	return kopsv1alpha2.CNINetworkingSpec{
		UsesSecondaryIP: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["uses_secondary_ip"]),
	}
}

func FlattenDataSourceCNINetworkingSpecInto(in kopsv1alpha2.CNINetworkingSpec, out map[string]interface{}) {
	out["uses_secondary_ip"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.UsesSecondaryIP)
}

func FlattenDataSourceCNINetworkingSpec(in kopsv1alpha2.CNINetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceCNINetworkingSpecInto(in, out)
	return out
}
