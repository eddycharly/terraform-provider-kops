package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceAlwaysAllowAuthorizationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceAlwaysAllowAuthorizationSpec(in map[string]interface{}) kopsv1alpha2.AlwaysAllowAuthorizationSpec {
	if in == nil {
		panic("expand AlwaysAllowAuthorizationSpec failure, in is nil")
	}
	return kopsv1alpha2.AlwaysAllowAuthorizationSpec{}
}

func FlattenDataSourceAlwaysAllowAuthorizationSpecInto(in kopsv1alpha2.AlwaysAllowAuthorizationSpec, out map[string]interface{}) {
}

func FlattenDataSourceAlwaysAllowAuthorizationSpec(in kopsv1alpha2.AlwaysAllowAuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAlwaysAllowAuthorizationSpecInto(in, out)
	return out
}
