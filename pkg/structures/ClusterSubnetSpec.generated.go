package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandClusterSubnetSpec(in map[string]interface{}) kops.ClusterSubnetSpec {
	if in == nil {
		panic("expand ClusterSubnetSpec failure, in is nil")
	}
	return kops.ClusterSubnetSpec{
		Name: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["name"]),
		CIDR: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cidr"]),
		Zone: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["zone"]),
		Region: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["region"]),
		ProviderID: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["provider_id"]),
		Egress: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["egress"]),
		Type: func(in interface{}) kops.SubnetType {
			value := kops.SubnetType(ExpandString(in))
			return value
		}(in["type"]),
		PublicIP: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["public_ip"]),
	}
}

func FlattenClusterSubnetSpec(in kops.ClusterSubnetSpec) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Name),
		"cidr": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.CIDR),
		"zone": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Zone),
		"region": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Region),
		"provider_id": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ProviderID),
		"egress": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Egress),
		"type": func(in kops.SubnetType) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Type),
		"public_ip": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.PublicIP),
	}
}
