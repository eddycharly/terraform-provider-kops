package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsKopeioAuthenticationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandResourceKopsKopeioAuthenticationSpec(in map[string]interface{}) kops.KopeioAuthenticationSpec {
	if in == nil {
		panic("expand KopeioAuthenticationSpec failure, in is nil")
	}
	return kops.KopeioAuthenticationSpec{}
}

func FlattenResourceKopsKopeioAuthenticationSpecInto(in kops.KopeioAuthenticationSpec, out map[string]interface{}) {
}

func FlattenResourceKopsKopeioAuthenticationSpec(in kops.KopeioAuthenticationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsKopeioAuthenticationSpecInto(in, out)
	return out
}
