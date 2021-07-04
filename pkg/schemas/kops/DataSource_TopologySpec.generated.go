package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceTopologySpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"masters": Computed(String()),
			"nodes":   Computed(String()),
			"bastion": Computed(Struct(DataSourceBastionSpec())),
			"dns":     Computed(Struct(DataSourceDNSSpec())),
		},
	}
}

func ExpandDataSourceTopologySpec(in map[string]interface{}) kops.TopologySpec {
	if in == nil {
		panic("expand TopologySpec failure, in is nil")
	}
	out := kops.TopologySpec{}
	if in, ok := in["masters"]; ok && in != nil {
		out.Masters = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["nodes"]; ok && in != nil {
		out.Nodes = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["bastion"]; ok && in != nil {
		out.Bastion = func(in interface{}) *kops.BastionSpec {
			if in == nil {
				return nil
			}
			return func(in kops.BastionSpec) *kops.BastionSpec { return &in }(func(in interface{}) kops.BastionSpec {
				if in == nil {
					return kops.BastionSpec{}
				}
				return ExpandDataSourceBastionSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["dns"]; ok && in != nil {
		out.DNS = func(in interface{}) *kops.DNSSpec {
			if in == nil {
				return nil
			}
			return func(in kops.DNSSpec) *kops.DNSSpec { return &in }(func(in interface{}) kops.DNSSpec {
				if in == nil {
					return kops.DNSSpec{}
				}
				return ExpandDataSourceDNSSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	return out
}

func FlattenDataSourceTopologySpecInto(in kops.TopologySpec, out map[string]interface{}) {
	out["masters"] = func(in string) interface{} { return string(in) }(in.Masters)
	out["nodes"] = func(in string) interface{} { return string(in) }(in.Nodes)
	out["bastion"] = func(in *kops.BastionSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.BastionSpec) interface{} { return FlattenDataSourceBastionSpec(in) }(*in)
	}(in.Bastion)
	out["dns"] = func(in *kops.DNSSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.DNSSpec) interface{} { return FlattenDataSourceDNSSpec(in) }(*in)
	}(in.DNS)
}

func FlattenDataSourceTopologySpec(in kops.TopologySpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceTopologySpecInto(in, out)
	return out
}
