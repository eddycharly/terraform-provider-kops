package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsKubenetNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandResourceKopsKubenetNetworkingSpec(in map[string]interface{}) kops.KubenetNetworkingSpec {
	if in == nil {
		panic("expand KubenetNetworkingSpec failure, in is nil")
	}
	return kops.KubenetNetworkingSpec{}
}

func FlattenResourceKopsKubenetNetworkingSpecInto(in kops.KubenetNetworkingSpec, out map[string]interface{}) {
}

func FlattenResourceKopsKubenetNetworkingSpec(in kops.KubenetNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsKubenetNetworkingSpecInto(in, out)
	return out
}
