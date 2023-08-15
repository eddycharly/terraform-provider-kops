package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops/v1alpha2"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceClusterSubnetSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":              RequiredString(),
			"zone":              RequiredString(),
			"region":            OptionalString(),
			"cidr":              OptionalComputedString(),
			"ipv6_cidr":         OptionalString(),
			"id":                OptionalString(),
			"egress":            OptionalString(),
			"type":              RequiredString(),
			"public_ip":         OptionalString(),
			"additional_routes": OptionalList(ResourceRouteSpec()),
		},
	}

	return res
}

func ExpandResourceClusterSubnetSpec(in map[string]interface{}) kopsv1alpha2.ClusterSubnetSpec {
	if in == nil {
		panic("expand ClusterSubnetSpec failure, in is nil")
	}
	return kopsv1alpha2.ClusterSubnetSpec{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		Zone: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["zone"]),
		Region: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["region"]),
		CIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cidr"]),
		IPv6CIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv6_cidr"]),
		ID: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["id"]),
		Egress: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["egress"]),
		Type: func(in interface{}) kopsv1alpha2.SubnetType {
			return v1alpha2.SubnetType(ExpandString(in))
		}(in["type"]),
		PublicIP: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["public_ip"]),
		AdditionalRoutes: func(in interface{}) []kopsv1alpha2.RouteSpec {
			return func(in interface{}) []kopsv1alpha2.RouteSpec {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.RouteSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.RouteSpec {
						if in == nil {
							return kopsv1alpha2.RouteSpec{}
						}
						return ExpandResourceRouteSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["additional_routes"]),
	}
}

func FlattenResourceClusterSubnetSpecInto(in kopsv1alpha2.ClusterSubnetSpec, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["zone"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Zone)
	out["region"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Region)
	out["cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CIDR)
	out["ipv6_cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.IPv6CIDR)
	out["id"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ID)
	out["egress"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Egress)
	out["type"] = func(in kopsv1alpha2.SubnetType) interface{} {
		return FlattenString(string(in))
	}(in.Type)
	out["public_ip"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.PublicIP)
	out["additional_routes"] = func(in []kopsv1alpha2.RouteSpec) interface{} {
		return func(in []kopsv1alpha2.RouteSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.RouteSpec) interface{} {
					return FlattenResourceRouteSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.AdditionalRoutes)
}

func FlattenResourceClusterSubnetSpec(in kopsv1alpha2.ClusterSubnetSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceClusterSubnetSpecInto(in, out)
	return out
}
