package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsRBACAuthorizationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandResourceKopsRBACAuthorizationSpec(in map[string]interface{}) kops.RBACAuthorizationSpec {
	if in == nil {
		panic("expand RBACAuthorizationSpec failure, in is nil")
	}
	return kops.RBACAuthorizationSpec{}
}

func FlattenResourceKopsRBACAuthorizationSpecInto(in kops.RBACAuthorizationSpec, out map[string]interface{}) {
}

func FlattenResourceKopsRBACAuthorizationSpec(in kops.RBACAuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsRBACAuthorizationSpecInto(in, out)
	return out
}
