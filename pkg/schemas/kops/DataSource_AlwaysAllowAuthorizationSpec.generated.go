package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAlwaysAllowAuthorizationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceAlwaysAllowAuthorizationSpec(in map[string]interface{}) kops.AlwaysAllowAuthorizationSpec {
	if in == nil {
		panic("expand AlwaysAllowAuthorizationSpec failure, in is nil")
	}
	return kops.AlwaysAllowAuthorizationSpec{}
}

func FlattenDataSourceAlwaysAllowAuthorizationSpecInto(in kops.AlwaysAllowAuthorizationSpec, out map[string]interface{}) {
}

func FlattenDataSourceAlwaysAllowAuthorizationSpec(in kops.AlwaysAllowAuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAlwaysAllowAuthorizationSpecInto(in, out)
	return out
}
