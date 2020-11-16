package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsKubenetNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandDataSourceKopsKubenetNetworkingSpec(in map[string]interface{}) kops.KubenetNetworkingSpec {
	if in == nil {
		panic("expand KubenetNetworkingSpec failure, in is nil")
	}
	return kops.KubenetNetworkingSpec{}
}

func FlattenDataSourceKopsKubenetNetworkingSpecInto(in kops.KubenetNetworkingSpec, out map[string]interface{}) {
}

func FlattenDataSourceKopsKubenetNetworkingSpec(in kops.KubenetNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsKubenetNetworkingSpecInto(in, out)
	return out
}
