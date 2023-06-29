package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceExternalNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandResourceExternalNetworkingSpec(in map[string]interface{}) kopsv1alpha2.ExternalNetworkingSpec {
	if in == nil {
		panic("expand ExternalNetworkingSpec failure, in is nil")
	}
	return kopsv1alpha2.ExternalNetworkingSpec{}
}

func FlattenResourceExternalNetworkingSpecInto(in kopsv1alpha2.ExternalNetworkingSpec, out map[string]interface{}) {
}

func FlattenResourceExternalNetworkingSpec(in kopsv1alpha2.ExternalNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceExternalNetworkingSpecInto(in, out)
	return out
}
