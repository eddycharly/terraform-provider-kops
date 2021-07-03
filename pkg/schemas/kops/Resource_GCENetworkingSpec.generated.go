package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceGCENetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandResourceGCENetworkingSpec(in map[string]interface{}) kops.GCENetworkingSpec {
	if in == nil {
		panic("expand GCENetworkingSpec failure, in is nil")
	}
	out := kops.GCENetworkingSpec{}
	return out
}
