package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceKopeioNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandResourceKopeioNetworkingSpec(in map[string]interface{}) kops.KopeioNetworkingSpec {
	if in == nil {
		panic("expand KopeioNetworkingSpec failure, in is nil")
	}
	out := kops.KopeioNetworkingSpec{}
	return out
}

func FlattenResourceKopeioNetworkingSpecInto(in kops.KopeioNetworkingSpec, out map[string]interface{}) {
}

func FlattenResourceKopeioNetworkingSpec(in kops.KopeioNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopeioNetworkingSpecInto(in, out)
	return out
}
