package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func AlwaysAllowAuthorizationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandAlwaysAllowAuthorizationSpec(in map[string]interface{}) kops.AlwaysAllowAuthorizationSpec {
	if in == nil {
		panic("expand AlwaysAllowAuthorizationSpec failure, in is nil")
	}
	return kops.AlwaysAllowAuthorizationSpec{}
}

func FlattenAlwaysAllowAuthorizationSpecInto(in kops.AlwaysAllowAuthorizationSpec, out map[string]interface{}) {
}

func FlattenAlwaysAllowAuthorizationSpec(in kops.AlwaysAllowAuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenAlwaysAllowAuthorizationSpecInto(in, out)
	return out
}
