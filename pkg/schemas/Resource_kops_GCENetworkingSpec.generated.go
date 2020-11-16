package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsGCENetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandResourceKopsGCENetworkingSpec(in map[string]interface{}) kops.GCENetworkingSpec {
	if in == nil {
		panic("expand GCENetworkingSpec failure, in is nil")
	}
	return kops.GCENetworkingSpec{}
}

func FlattenResourceKopsGCENetworkingSpecInto(in kops.GCENetworkingSpec, out map[string]interface{}) {
}

func FlattenResourceKopsGCENetworkingSpec(in kops.GCENetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsGCENetworkingSpecInto(in, out)
	return out
}
