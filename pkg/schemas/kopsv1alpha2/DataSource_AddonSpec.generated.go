package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceAddonSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"manifest": ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceAddonSpec(in map[string]interface{}) kopsv1alpha2.AddonSpec {
	if in == nil {
		panic("expand AddonSpec failure, in is nil")
	}
	return kopsv1alpha2.AddonSpec{
		Manifest: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["manifest"]),
	}
}

func FlattenDataSourceAddonSpecInto(in kopsv1alpha2.AddonSpec, out map[string]interface{}) {
	out["manifest"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Manifest)
}

func FlattenDataSourceAddonSpec(in kopsv1alpha2.AddonSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAddonSpecInto(in, out)
	return out
}
