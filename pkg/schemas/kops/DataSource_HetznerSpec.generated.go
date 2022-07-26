package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceHetznerSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceHetznerSpec(in map[string]interface{}) kops.HetznerSpec {
	if in == nil {
		panic("expand HetznerSpec failure, in is nil")
	}
	return kops.HetznerSpec{}
}

func FlattenDataSourceHetznerSpecInto(in kops.HetznerSpec, out map[string]interface{}) {
}

func FlattenDataSourceHetznerSpec(in kops.HetznerSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceHetznerSpecInto(in, out)
	return out
}
