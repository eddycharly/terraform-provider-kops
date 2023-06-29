package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceKopeioAuthenticationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceKopeioAuthenticationSpec(in map[string]interface{}) kopsv1alpha2.KopeioAuthenticationSpec {
	if in == nil {
		panic("expand KopeioAuthenticationSpec failure, in is nil")
	}
	return kopsv1alpha2.KopeioAuthenticationSpec{}
}

func FlattenDataSourceKopeioAuthenticationSpecInto(in kopsv1alpha2.KopeioAuthenticationSpec, out map[string]interface{}) {
}

func FlattenDataSourceKopeioAuthenticationSpec(in kopsv1alpha2.KopeioAuthenticationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopeioAuthenticationSpecInto(in, out)
	return out
}
