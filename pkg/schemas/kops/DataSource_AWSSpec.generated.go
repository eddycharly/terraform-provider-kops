package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAWSSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceAWSSpec(in map[string]interface{}) kops.AWSSpec {
	if in == nil {
		panic("expand AWSSpec failure, in is nil")
	}
	return kops.AWSSpec{}
}

func FlattenDataSourceAWSSpecInto(in kops.AWSSpec, out map[string]interface{}) {
}

func FlattenDataSourceAWSSpec(in kops.AWSSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAWSSpecInto(in, out)
	return out
}
