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
			return string(ExpandString(in))
		}(in["name"]),
		CIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cidr"]),
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
	}
}

func FlattenClusterSubnetSpecInto(in kops.ClusterSubnetSpec, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CIDR)
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
}

func FlattenClusterSubnetSpec(in kops.ClusterSubnetSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenClusterSubnetSpecInto(in, out)
	return out
}
