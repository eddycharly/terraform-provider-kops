package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsAlwaysAllowAuthorizationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandDataSourceKopsAlwaysAllowAuthorizationSpec(in map[string]interface{}) kops.AlwaysAllowAuthorizationSpec {
	if in == nil {
		panic("expand AlwaysAllowAuthorizationSpec failure, in is nil")
	}
	return kops.AlwaysAllowAuthorizationSpec{}
}

func FlattenDataSourceKopsAlwaysAllowAuthorizationSpecInto(in kops.AlwaysAllowAuthorizationSpec, out map[string]interface{}) {
}

func FlattenDataSourceKopsAlwaysAllowAuthorizationSpec(in kops.AlwaysAllowAuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsAlwaysAllowAuthorizationSpecInto(in, out)
	return out
}
