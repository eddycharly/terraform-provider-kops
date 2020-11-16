package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsRBACAuthorizationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandDataSourceKopsRBACAuthorizationSpec(in map[string]interface{}) kops.RBACAuthorizationSpec {
	if in == nil {
		panic("expand RBACAuthorizationSpec failure, in is nil")
	}
	return kops.RBACAuthorizationSpec{}
}

func FlattenDataSourceKopsRBACAuthorizationSpecInto(in kops.RBACAuthorizationSpec, out map[string]interface{}) {
}

func FlattenDataSourceKopsRBACAuthorizationSpec(in kops.RBACAuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsRBACAuthorizationSpecInto(in, out)
	return out
}
