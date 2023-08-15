package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceKubenetNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandResourceKubenetNetworkingSpec(in map[string]interface{}) kopsv1alpha2.KubenetNetworkingSpec {
	if in == nil {
		panic("expand KubenetNetworkingSpec failure, in is nil")
	}
	return kopsv1alpha2.KubenetNetworkingSpec{}
}

func FlattenResourceKubenetNetworkingSpecInto(in kopsv1alpha2.KubenetNetworkingSpec, out map[string]interface{}) {
}

func FlattenResourceKubenetNetworkingSpec(in kopsv1alpha2.KubenetNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKubenetNetworkingSpecInto(in, out)
	return out
}
