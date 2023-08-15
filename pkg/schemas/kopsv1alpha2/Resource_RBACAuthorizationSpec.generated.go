package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceRBACAuthorizationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandResourceRBACAuthorizationSpec(in map[string]interface{}) kopsv1alpha2.RBACAuthorizationSpec {
	if in == nil {
		panic("expand RBACAuthorizationSpec failure, in is nil")
	}
	return kopsv1alpha2.RBACAuthorizationSpec{}
}

func FlattenResourceRBACAuthorizationSpecInto(in kopsv1alpha2.RBACAuthorizationSpec, out map[string]interface{}) {
}

func FlattenResourceRBACAuthorizationSpec(in kopsv1alpha2.RBACAuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceRBACAuthorizationSpecInto(in, out)
	return out
}
