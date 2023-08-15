package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceKuberouterNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceKuberouterNetworkingSpec(in map[string]interface{}) kopsv1alpha2.KuberouterNetworkingSpec {
	if in == nil {
		panic("expand KuberouterNetworkingSpec failure, in is nil")
	}
	return kopsv1alpha2.KuberouterNetworkingSpec{}
}

func FlattenDataSourceKuberouterNetworkingSpecInto(in kopsv1alpha2.KuberouterNetworkingSpec, out map[string]interface{}) {
}

func FlattenDataSourceKuberouterNetworkingSpec(in kopsv1alpha2.KuberouterNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKuberouterNetworkingSpecInto(in, out)
	return out
}
