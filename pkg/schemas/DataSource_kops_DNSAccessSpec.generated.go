package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsDNSAccessSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandDataSourceKopsDNSAccessSpec(in map[string]interface{}) kops.DNSAccessSpec {
	if in == nil {
		panic("expand DNSAccessSpec failure, in is nil")
	}
	return kops.DNSAccessSpec{}
}

func FlattenDataSourceKopsDNSAccessSpecInto(in kops.DNSAccessSpec, out map[string]interface{}) {
}

func FlattenDataSourceKopsDNSAccessSpec(in kops.DNSAccessSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsDNSAccessSpecInto(in, out)
	return out
}
