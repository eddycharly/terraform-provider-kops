package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsIAMSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"legacy":                   OptionalBool(),
			"allow_container_registry": OptionalBool(),
		},
	}
}

func ExpandResourceKopsIAMSpec(in map[string]interface{}) kops.IAMSpec {
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

func FlattenResourceKopsIAMSpecInto(in kops.IAMSpec, out map[string]interface{}) {
	out["legacy"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Legacy)
	out["allow_container_registry"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.AllowContainerRegistry)
}

func FlattenResourceKopsIAMSpec(in kops.IAMSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsIAMSpecInto(in, out)
	return out
}
