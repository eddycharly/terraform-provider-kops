package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceTopologySpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"masters": Required(String()),
			"nodes":   Required(String()),
			"bastion": Optional(Ptr(Struct(ResourceBastionSpec()))),
			"dns":     Required(Ptr(Struct(ResourceDNSSpec()))),
		},
	}
}

func ExpandResourceTopologySpec(in map[string]interface{}) kops.TopologySpec {
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
				return ExpandResourceBastionSpec(in.(map[string]interface{}))
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
				return ExpandResourceDNSSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	return out
}
