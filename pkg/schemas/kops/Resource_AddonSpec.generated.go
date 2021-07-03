package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceAddonSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"manifest": Required(String()),
		},
	}
}

func ExpandResourceAddonSpec(in map[string]interface{}) kops.AddonSpec {
	if in == nil {
		panic("expand AddonSpec failure, in is nil")
	}
	out := kops.AddonSpec{}
	if in, ok := in["manifest"]; ok && in != nil {
		out.Manifest = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}
