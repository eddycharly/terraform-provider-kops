package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsIAMSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"legacy":                   ComputedBool(),
			"allow_container_registry": ComputedBool(),
		},
	}
}

func ExpandDataSourceKopsIAMSpec(in map[string]interface{}) kops.IAMSpec {
	if in == nil {
		panic("expand IAMSpec failure, in is nil")
	}
	return kops.IAMSpec{
		Legacy: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["legacy"]),
		AllowContainerRegistry: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["allow_container_registry"]),
	}
}

func FlattenDataSourceKopsIAMSpecInto(in kops.IAMSpec, out map[string]interface{}) {
	out["legacy"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Legacy)
	out["allow_container_registry"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.AllowContainerRegistry)
}

func FlattenDataSourceKopsIAMSpec(in kops.IAMSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsIAMSpecInto(in, out)
	return out
}
