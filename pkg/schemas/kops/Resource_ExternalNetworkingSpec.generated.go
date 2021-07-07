package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceExternalNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandResourceExternalNetworkingSpec(in map[string]interface{}) kops.ExternalNetworkingSpec {
	if in == nil {
		panic("expand ExternalNetworkingSpec failure, in is nil")
	}
	return kops.ExternalNetworkingSpec{}
}

func FlattenResourceExternalNetworkingSpecInto(in kops.ExternalNetworkingSpec, out map[string]interface{}) {
}

func FlattenResourceExternalNetworkingSpec(in kops.ExternalNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceExternalNetworkingSpecInto(in, out)
	return out
}
