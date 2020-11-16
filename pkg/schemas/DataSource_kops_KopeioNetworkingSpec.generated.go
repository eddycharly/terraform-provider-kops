package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsKopeioNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandDataSourceKopsKopeioNetworkingSpec(in map[string]interface{}) kops.KopeioNetworkingSpec {
	if in == nil {
		panic("expand KopeioNetworkingSpec failure, in is nil")
	}
	return kops.KopeioNetworkingSpec{}
}

func FlattenDataSourceKopsKopeioNetworkingSpecInto(in kops.KopeioNetworkingSpec, out map[string]interface{}) {
}

func FlattenDataSourceKopsKopeioNetworkingSpec(in kops.KopeioNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsKopeioNetworkingSpecInto(in, out)
	return out
}
