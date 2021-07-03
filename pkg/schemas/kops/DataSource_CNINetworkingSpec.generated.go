package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceCNINetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"uses_secondary_ip": Computed(Bool()),
		},
	}
}

func ExpandDataSourceCNINetworkingSpec(in map[string]interface{}) kops.CNINetworkingSpec {
	if in == nil {
		panic("expand CNINetworkingSpec failure, in is nil")
	}
	out := kops.CNINetworkingSpec{}
	if in, ok := in["uses_secondary_ip"]; ok && in != nil {
		out.UsesSecondaryIP = func(in interface{}) bool { return in.(bool) }(in)
	}
	return out
}
