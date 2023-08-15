package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops/v1alpha2"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceTopologySpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"control_plane": RequiredString(),
			"nodes":         RequiredString(),
			"bastion":       OptionalStruct(ResourceBastionSpec()),
			"dns":           RequiredString(),
			"legacy_dns":    OptionalStruct(ResourceDNSSpec()),
		},
	}

	return res
}

func ExpandResourceTopologySpec(in map[string]interface{}) kopsv1alpha2.TopologySpec {
	if in == nil {
		panic("expand TopologySpec failure, in is nil")
	}
	return kopsv1alpha2.TopologySpec{
		ControlPlane: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["control_plane"]),
		Nodes: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["nodes"]),
		Bastion: func(in interface{}) *kopsv1alpha2.BastionSpec {
			return func(in interface{}) *kopsv1alpha2.BastionSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.BastionSpec) *kopsv1alpha2.BastionSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.BastionSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceBastionSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.BastionSpec{}
				}(in))
			}(in)
		}(in["bastion"]),
		DNS: func(in interface{}) kopsv1alpha2.DNSType {
			return v1alpha2.DNSType(ExpandString(in))
		}(in["dns"]),
		LegacyDNS: func(in interface{}) *kopsv1alpha2.DNSSpec {
			return func(in interface{}) *kopsv1alpha2.DNSSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.DNSSpec) *kopsv1alpha2.DNSSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.DNSSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceDNSSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.DNSSpec{}
				}(in))
			}(in)
		}(in["legacy_dns"]),
	}
}

func FlattenResourceTopologySpecInto(in kopsv1alpha2.TopologySpec, out map[string]interface{}) {
	out["control_plane"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ControlPlane)
	out["nodes"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Nodes)
	out["bastion"] = func(in *kopsv1alpha2.BastionSpec) interface{} {
		return func(in *kopsv1alpha2.BastionSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.BastionSpec) interface{} {
				return func(in kopsv1alpha2.BastionSpec) []interface{} {
					return []interface{}{FlattenResourceBastionSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Bastion)
	out["dns"] = func(in kopsv1alpha2.DNSType) interface{} {
		return FlattenString(string(in))
	}(in.DNS)
	out["legacy_dns"] = func(in *kopsv1alpha2.DNSSpec) interface{} {
		return func(in *kopsv1alpha2.DNSSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.DNSSpec) interface{} {
				return func(in kopsv1alpha2.DNSSpec) []interface{} {
					return []interface{}{FlattenResourceDNSSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.LegacyDNS)
}

func FlattenResourceTopologySpec(in kopsv1alpha2.TopologySpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceTopologySpecInto(in, out)
	return out
}
