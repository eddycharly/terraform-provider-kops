package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceGCESpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandResourceGCESpec(in map[string]interface{}) kops.GCESpec {
	if in == nil {
		panic("expand GCESpec failure, in is nil")
	}
	return kops.GCESpec{}
}

func FlattenResourceGCESpecInto(in kops.GCESpec, out map[string]interface{}) {
}

func FlattenResourceGCESpec(in kops.GCESpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceGCESpecInto(in, out)
	return out
}
