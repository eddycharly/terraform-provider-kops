package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceDOSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandResourceDOSpec(in map[string]interface{}) kops.DOSpec {
	if in == nil {
		panic("expand DOSpec failure, in is nil")
	}
	return kops.DOSpec{}
}

func FlattenResourceDOSpecInto(in kops.DOSpec, out map[string]interface{}) {
}

func FlattenResourceDOSpec(in kops.DOSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceDOSpecInto(in, out)
	return out
}
