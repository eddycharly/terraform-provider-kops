package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceGCENetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandResourceGCENetworkingSpec(in map[string]interface{}) kopsv1alpha2.GCENetworkingSpec {
	if in == nil {
		panic("expand GCENetworkingSpec failure, in is nil")
	}
	return kopsv1alpha2.GCENetworkingSpec{}
}

func FlattenResourceGCENetworkingSpecInto(in kopsv1alpha2.GCENetworkingSpec, out map[string]interface{}) {
}

func FlattenResourceGCENetworkingSpec(in kopsv1alpha2.GCENetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceGCENetworkingSpecInto(in, out)
	return out
}
