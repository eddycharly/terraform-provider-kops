package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceClusterSubnetSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":              RequiredString(),
			"cidr":              OptionalComputedString(),
			"ipv6_cidr":         OptionalString(),
			"zone":              RequiredString(),
			"region":            OptionalString(),
			"provider_id":       RequiredString(),
			"egress":            OptionalString(),
			"type":              RequiredString(),
			"public_ip":         OptionalString(),
			"additional_routes": OptionalList(ResourceRouteSpec()),
		},
	}

	return res
}

func ExpandResourceClusterSubnetSpec(in map[string]interface{}) kops.ClusterSubnetSpec {
	if in == nil {
		panic("expand ClusterSubnetSpec failure, in is nil")
	}
	return kops.ClusterSubnetSpec{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		CIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cidr"]),
		IPv6CIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv6_cidr"]),
		Zone: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["zone"]),
		Region: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["region"]),
		ProviderID: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["provider_id"]),
		Egress: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["egress"]),
		Type: func(in interface{}) kops.SubnetType {
			return kops.SubnetType(ExpandString(in))
		}(in["type"]),
		PublicIP: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["public_ip"]),
		AdditionalRoutes: func(in interface{}) []kops.RouteSpec {
			return func(in interface{}) []kops.RouteSpec {
				if in == nil {
					return nil
				}
				var out []kops.RouteSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.RouteSpec {
						if in == nil {
							return kops.RouteSpec{}
						}
						return (ExpandResourceRouteSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["additional_routes"]),
	}
}

func FlattenResourceClusterSubnetSpecInto(in kops.ClusterSubnetSpec, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CIDR)
	out["ipv6_cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.IPv6CIDR)
	out["zone"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Zone)
	out["region"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Region)
	out["provider_id"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ProviderID)
	out["egress"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Egress)
	out["type"] = func(in kops.SubnetType) interface{} {
		return FlattenString(string(in))
	}(in.Type)
	out["public_ip"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.PublicIP)
	out["additional_routes"] = func(in []kops.RouteSpec) interface{} {
		return func(in []kops.RouteSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.RouteSpec) interface{} {
					return FlattenResourceRouteSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.AdditionalRoutes)
}

func FlattenResourceClusterSubnetSpec(in kops.ClusterSubnetSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceClusterSubnetSpecInto(in, out)
	return out
}
