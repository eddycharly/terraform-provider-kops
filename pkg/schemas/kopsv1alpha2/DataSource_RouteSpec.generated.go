package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceRouteSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cidr":   ComputedString(),
			"target": ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceRouteSpec(in map[string]interface{}) kopsv1alpha2.RouteSpec {
	if in == nil {
		panic("expand RouteSpec failure, in is nil")
	}
	return kopsv1alpha2.RouteSpec{
		CIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cidr"]),
		Target: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["target"]),
	}
}

func FlattenDataSourceRouteSpecInto(in kopsv1alpha2.RouteSpec, out map[string]interface{}) {
	out["cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CIDR)
	out["target"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Target)
}

func FlattenDataSourceRouteSpec(in kopsv1alpha2.RouteSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceRouteSpecInto(in, out)
	return out
}
