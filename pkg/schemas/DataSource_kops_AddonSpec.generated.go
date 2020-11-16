package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsAddonSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"manifest": ComputedString(),
		},
	}
}

func ExpandDataSourceKopsAddonSpec(in map[string]interface{}) kops.AddonSpec {
	if in == nil {
		panic("expand AddonSpec failure, in is nil")
	}
	return kops.AddonSpec{
		Manifest: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["manifest"]),
	}
}

func FlattenDataSourceKopsAddonSpecInto(in kops.AddonSpec, out map[string]interface{}) {
	out["manifest"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Manifest)
}

func FlattenDataSourceKopsAddonSpec(in kops.AddonSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsAddonSpecInto(in, out)
	return out
}
