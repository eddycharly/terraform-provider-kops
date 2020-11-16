package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsKuberouterNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandDataSourceKopsKuberouterNetworkingSpec(in map[string]interface{}) kops.KuberouterNetworkingSpec {
	if in == nil {
		panic("expand KuberouterNetworkingSpec failure, in is nil")
	}
	return kops.KuberouterNetworkingSpec{}
}

func FlattenDataSourceKopsKuberouterNetworkingSpecInto(in kops.KuberouterNetworkingSpec, out map[string]interface{}) {
}

func FlattenDataSourceKopsKuberouterNetworkingSpec(in kops.KuberouterNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsKuberouterNetworkingSpecInto(in, out)
	return out
}
