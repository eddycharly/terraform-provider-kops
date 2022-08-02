package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceRouteSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cidr":   OptionalString(),
			"target": OptionalString(),
		},
	}

	return res
}

func ExpandResourceRouteSpec(in map[string]interface{}) kops.RouteSpec {
	if in == nil {
		panic("expand RouteSpec failure, in is nil")
	}
	return kops.RouteSpec{
		CIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cidr"]),
		Target: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["target"]),
	}
}

func FlattenResourceRouteSpecInto(in kops.RouteSpec, out map[string]interface{}) {
	out["cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CIDR)
	out["target"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Target)
}

func FlattenResourceRouteSpec(in kops.RouteSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceRouteSpecInto(in, out)
	return out
}
