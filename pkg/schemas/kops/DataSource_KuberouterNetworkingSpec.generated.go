package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceKuberouterNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceKuberouterNetworkingSpec(in map[string]interface{}) kops.KuberouterNetworkingSpec {
	if in == nil {
		panic("expand KuberouterNetworkingSpec failure, in is nil")
	}
	return kops.KuberouterNetworkingSpec{}
}

func FlattenDataSourceKuberouterNetworkingSpecInto(in kops.KuberouterNetworkingSpec, out map[string]interface{}) {
}

func FlattenDataSourceKuberouterNetworkingSpec(in kops.KuberouterNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKuberouterNetworkingSpecInto(in, out)
	return out
}
