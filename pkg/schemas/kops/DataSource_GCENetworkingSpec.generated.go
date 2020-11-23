package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceGCENetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandDataSourceGCENetworkingSpec(in map[string]interface{}) kops.GCENetworkingSpec {
	if in == nil {
		panic("expand GCENetworkingSpec failure, in is nil")
	}
	return kops.GCENetworkingSpec{}
}

func FlattenDataSourceGCENetworkingSpecInto(in kops.GCENetworkingSpec, out map[string]interface{}) {
}

func FlattenDataSourceGCENetworkingSpec(in kops.GCENetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceGCENetworkingSpecInto(in, out)
	return out
}
