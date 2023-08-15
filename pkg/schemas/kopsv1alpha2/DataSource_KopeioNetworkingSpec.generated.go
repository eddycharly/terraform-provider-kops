package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceKopeioNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceKopeioNetworkingSpec(in map[string]interface{}) kopsv1alpha2.KopeioNetworkingSpec {
	if in == nil {
		panic("expand KopeioNetworkingSpec failure, in is nil")
	}
	return kopsv1alpha2.KopeioNetworkingSpec{}
}

func FlattenDataSourceKopeioNetworkingSpecInto(in kopsv1alpha2.KopeioNetworkingSpec, out map[string]interface{}) {
}

func FlattenDataSourceKopeioNetworkingSpec(in kopsv1alpha2.KopeioNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopeioNetworkingSpecInto(in, out)
	return out
}
