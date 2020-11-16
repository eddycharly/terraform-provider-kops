package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsKopeioAuthenticationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandDataSourceKopsKopeioAuthenticationSpec(in map[string]interface{}) kops.KopeioAuthenticationSpec {
	if in == nil {
		panic("expand KopeioAuthenticationSpec failure, in is nil")
	}
	return kops.KopeioAuthenticationSpec{}
}

func FlattenDataSourceKopsKopeioAuthenticationSpecInto(in kops.KopeioAuthenticationSpec, out map[string]interface{}) {
}

func FlattenDataSourceKopsKopeioAuthenticationSpec(in kops.KopeioAuthenticationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsKopeioAuthenticationSpecInto(in, out)
	return out
}
