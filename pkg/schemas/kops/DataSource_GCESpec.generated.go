package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceGCESpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceGCESpec(in map[string]interface{}) kops.GCESpec {
	if in == nil {
		panic("expand GCESpec failure, in is nil")
	}
	return kops.GCESpec{}
}

func FlattenDataSourceGCESpecInto(in kops.GCESpec, out map[string]interface{}) {
}

func FlattenDataSourceGCESpec(in kops.GCESpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceGCESpecInto(in, out)
	return out
}
