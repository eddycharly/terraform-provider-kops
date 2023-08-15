package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceAddonSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"manifest": RequiredString(),
		},
	}

	return res
}

func ExpandResourceAddonSpec(in map[string]interface{}) kopsv1alpha2.AddonSpec {
	if in == nil {
		panic("expand AddonSpec failure, in is nil")
	}
	return kopsv1alpha2.AddonSpec{
		Manifest: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["manifest"]),
	}
}

func FlattenResourceAddonSpecInto(in kopsv1alpha2.AddonSpec, out map[string]interface{}) {
	out["manifest"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Manifest)
}

func FlattenResourceAddonSpec(in kopsv1alpha2.AddonSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAddonSpecInto(in, out)
	return out
}
