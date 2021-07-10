package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceTopologySpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"masters": ComputedString(),
			"nodes":   ComputedString(),
			"bastion": ComputedStruct(DataSourceBastionSpec()),
			"dns":     ComputedStruct(DataSourceDNSSpec()),
		},
	}

	return res
}

func ExpandDataSourceTopologySpec(in map[string]interface{}) kops.TopologySpec {
	if in == nil {
		panic("expand TopologySpec failure, in is nil")
	}
	return kops.TopologySpec{
		Masters: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["masters"]),
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.BastionSpec{}
					}
					return (ExpandDataSourceBastionSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["bastion"]),
		DNS: func(in interface{}) *kops.DNSSpec {
			return func(in interface{}) *kops.DNSSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.DNSSpec) *kops.DNSSpec {
					return &in
				}(func(in interface{}) kops.DNSSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.DNSSpec{}
					}
					return (ExpandDataSourceDNSSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["dns"]),
	}
}

func FlattenDataSourceTopologySpecInto(in kops.TopologySpec, out map[string]interface{}) {
	out["masters"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Masters)
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
					return []interface{}{FlattenDataSourceBastionSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Bastion)
	out["dns"] = func(in *kops.DNSSpec) interface{} {
		return func(in *kops.DNSSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.DNSSpec) interface{} {
				return func(in kops.DNSSpec) []interface{} {
					return []interface{}{FlattenDataSourceDNSSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.DNS)
}

func FlattenDataSourceTopologySpec(in kops.TopologySpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceTopologySpecInto(in, out)
	return out
}
