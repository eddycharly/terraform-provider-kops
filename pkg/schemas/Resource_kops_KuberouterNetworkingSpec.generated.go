package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsKuberouterNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandResourceKopsKuberouterNetworkingSpec(in map[string]interface{}) kops.KuberouterNetworkingSpec {
	if in == nil {
		panic("expand KuberouterNetworkingSpec failure, in is nil")
	}
	return kops.KuberouterNetworkingSpec{}
}

func FlattenResourceKopsKuberouterNetworkingSpecInto(in kops.KuberouterNetworkingSpec, out map[string]interface{}) {
}

func FlattenResourceKopsKuberouterNetworkingSpec(in kops.KuberouterNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsKuberouterNetworkingSpecInto(in, out)
	return out
}
