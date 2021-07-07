package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceKopeioNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceKopeioNetworkingSpec(in map[string]interface{}) kops.KopeioNetworkingSpec {
	if in == nil {
		panic("expand KopeioNetworkingSpec failure, in is nil")
	}
	return kops.KopeioNetworkingSpec{}
}

func FlattenDataSourceKopeioNetworkingSpecInto(in kops.KopeioNetworkingSpec, out map[string]interface{}) {
}

func FlattenDataSourceKopeioNetworkingSpec(in kops.KopeioNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopeioNetworkingSpecInto(in, out)
	return out
}
