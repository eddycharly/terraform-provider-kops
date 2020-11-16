package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsKopeioNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandResourceKopsKopeioNetworkingSpec(in map[string]interface{}) kops.KopeioNetworkingSpec {
	if in == nil {
		panic("expand KopeioNetworkingSpec failure, in is nil")
	}
	return kops.KopeioNetworkingSpec{}
}

func FlattenResourceKopsKopeioNetworkingSpecInto(in kops.KopeioNetworkingSpec, out map[string]interface{}) {
}

func FlattenResourceKopsKopeioNetworkingSpec(in kops.KopeioNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsKopeioNetworkingSpecInto(in, out)
	return out
}
