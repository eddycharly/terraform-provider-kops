package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceTopologySpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"control_plane": RequiredString(),
			"nodes":         RequiredString(),
			"bastion":       OptionalStruct(ResourceBastionSpec()),
			"dns":           RequiredString(),
		},
	}

	return res
}

func ExpandResourceTopologySpec(in map[string]interface{}) kops.TopologySpec {
	if in == nil {
		panic("expand TopologySpec failure, in is nil")
	}
	return kops.TopologySpec{
		ControlPlane: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["control_plane"]),
		Nodes: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["nodes"]),
		Bastion: func(in interface{}) *kops.BastionSpec {
			return func(in interface{}) *kops.BastionSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.BastionSpec) *kops.BastionSpec {
					return &in
				}(func(in interface{}) kops.BastionSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceBastionSpec(in[0].(map[string]interface{}))
					}
					return kops.BastionSpec{}
				}(in))
			}(in)
		}(in["bastion"]),
		DNS: func(in interface{}) kops.DNSType {
			return kops.DNSType(ExpandString(in))
		}(in["dns"]),
	}
}

func FlattenResourceTopologySpecInto(in kops.TopologySpec, out map[string]interface{}) {
	out["control_plane"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ControlPlane)
	out["nodes"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Nodes)
	out["bastion"] = func(in *kops.BastionSpec) interface{} {
		return func(in *kops.BastionSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.BastionSpec) interface{} {
				return func(in kops.BastionSpec) []interface{} {
					return []interface{}{FlattenResourceBastionSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Bastion)
	out["dns"] = func(in kops.DNSType) interface{} {
		return FlattenString(string(in))
	}(in.DNS)
}

func FlattenResourceTopologySpec(in kops.TopologySpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceTopologySpecInto(in, out)
	return out
}
