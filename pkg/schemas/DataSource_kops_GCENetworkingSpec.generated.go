package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsGCENetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandDataSourceKopsGCENetworkingSpec(in map[string]interface{}) kops.GCENetworkingSpec {
	if in == nil {
		panic("expand GCENetworkingSpec failure, in is nil")
	}
	return kops.GCENetworkingSpec{}
}

func FlattenDataSourceKopsGCENetworkingSpecInto(in kops.GCENetworkingSpec, out map[string]interface{}) {
}

func FlattenDataSourceKopsGCENetworkingSpec(in kops.GCENetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsGCENetworkingSpecInto(in, out)
	return out
}
