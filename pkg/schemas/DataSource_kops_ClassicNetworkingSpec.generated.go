package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsClassicNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandDataSourceKopsClassicNetworkingSpec(in map[string]interface{}) kops.ClassicNetworkingSpec {
	if in == nil {
		panic("expand ClassicNetworkingSpec failure, in is nil")
	}
	return kops.ClassicNetworkingSpec{}
}

func FlattenDataSourceKopsClassicNetworkingSpecInto(in kops.ClassicNetworkingSpec, out map[string]interface{}) {
}

func FlattenDataSourceKopsClassicNetworkingSpec(in kops.ClassicNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsClassicNetworkingSpecInto(in, out)
	return out
}
