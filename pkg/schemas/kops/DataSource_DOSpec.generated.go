package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceDOSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceDOSpec(in map[string]interface{}) kops.DOSpec {
	if in == nil {
		panic("expand DOSpec failure, in is nil")
	}
	return kops.DOSpec{}
}

func FlattenDataSourceDOSpecInto(in kops.DOSpec, out map[string]interface{}) {
}

func FlattenDataSourceDOSpec(in kops.DOSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceDOSpecInto(in, out)
	return out
}
