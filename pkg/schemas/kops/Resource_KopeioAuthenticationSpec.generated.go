package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceKopeioAuthenticationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandResourceKopeioAuthenticationSpec(in map[string]interface{}) kops.KopeioAuthenticationSpec {
	if in == nil {
		panic("expand KopeioAuthenticationSpec failure, in is nil")
	}
	return kops.KopeioAuthenticationSpec{}
}

func FlattenResourceKopeioAuthenticationSpecInto(in kops.KopeioAuthenticationSpec, out map[string]interface{}) {
}

func FlattenResourceKopeioAuthenticationSpec(in kops.KopeioAuthenticationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopeioAuthenticationSpecInto(in, out)
	return out
}
