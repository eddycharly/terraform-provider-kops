package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceClassicNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandResourceClassicNetworkingSpec(in map[string]interface{}) kopsv1alpha2.ClassicNetworkingSpec {
	if in == nil {
		panic("expand ClassicNetworkingSpec failure, in is nil")
	}
	return kopsv1alpha2.ClassicNetworkingSpec{}
}

func FlattenResourceClassicNetworkingSpecInto(in kopsv1alpha2.ClassicNetworkingSpec, out map[string]interface{}) {
}

func FlattenResourceClassicNetworkingSpec(in kopsv1alpha2.ClassicNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceClassicNetworkingSpecInto(in, out)
	return out
}
