package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceClusterSubnetSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":        Required(String()),
			"cidr":        Optional(Computed(String())),
			"zone":        Required(String()),
			"region":      Optional(String()),
			"provider_id": Required(String()),
			"egress":      Optional(String()),
			"type":        Required(String()),
			"public_ip":   Optional(String()),
		},
	}
}

func ExpandResourceClusterSubnetSpec(in map[string]interface{}) kops.ClusterSubnetSpec {
	if in == nil {
		panic("expand ClusterSubnetSpec failure, in is nil")
	}
	out := kops.ClusterSubnetSpec{}
	if in, ok := in["name"]; ok && in != nil {
		out.Name = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cidr"]; ok && in != nil {
		out.CIDR = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["zone"]; ok && in != nil {
		out.Zone = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["region"]; ok && in != nil {
		out.Region = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["provider_id"]; ok && in != nil {
		out.ProviderID = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["egress"]; ok && in != nil {
		out.Egress = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["type"]; ok && in != nil {
		out.Type = func(in interface{}) kops.SubnetType { return kops.SubnetType(in.(string)) }(in)
	}
	if in, ok := in["public_ip"]; ok && in != nil {
		out.PublicIP = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}
