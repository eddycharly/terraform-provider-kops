package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsExternalNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandResourceKopsExternalNetworkingSpec(in map[string]interface{}) kops.ExternalNetworkingSpec {
	if in == nil {
		panic("expand ExternalNetworkingSpec failure, in is nil")
	}
	return kops.ExternalNetworkingSpec{}
}

func FlattenResourceKopsExternalNetworkingSpecInto(in kops.ExternalNetworkingSpec, out map[string]interface{}) {
}

func FlattenResourceKopsExternalNetworkingSpec(in kops.ExternalNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsExternalNetworkingSpecInto(in, out)
	return out
}
