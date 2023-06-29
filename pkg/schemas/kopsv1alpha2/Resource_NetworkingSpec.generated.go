package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_id":               RequiredString(),
			"network_cidr":             OptionalComputedString(),
			"additional_network_cidrs": OptionalList(String()),
			"subnet":                   RequiredList(ResourceClusterSubnetSpec()),
			"tag_subnets":              Nullable(OptionalBool()),
			"topology":                 RequiredStruct(ResourceTopologySpec()),
			"egress_proxy":             OptionalStruct(ResourceEgressProxySpec()),
			"non_masquerade_cidr":      OptionalComputedString(),
			"pod_cidr":                 OptionalString(),
			"service_cluster_ip_range": OptionalString(),
			"isolate_control_plane":    OptionalBool(),
			"classic":                  OptionalStruct(ResourceClassicNetworkingSpec()),
			"kubenet":                  OptionalStruct(ResourceKubenetNetworkingSpec()),
			"external":                 OptionalStruct(ResourceExternalNetworkingSpec()),
			"cni":                      OptionalStruct(ResourceCNINetworkingSpec()),
			"kopeio":                   OptionalStruct(ResourceKopeioNetworkingSpec()),
			"weave":                    OptionalStruct(ResourceWeaveNetworkingSpec()),
			"flannel":                  OptionalStruct(ResourceFlannelNetworkingSpec()),
			"calico":                   OptionalStruct(ResourceCalicoNetworkingSpec()),
			"canal":                    OptionalStruct(ResourceCanalNetworkingSpec()),
			"kube_router":              OptionalStruct(ResourceKuberouterNetworkingSpec()),
			"romana":                   OptionalStruct(ResourceRomanaNetworkingSpec()),
			"amazon_vpc":               OptionalStruct(ResourceAmazonVPCNetworkingSpec()),
			"cilium":                   OptionalStruct(ResourceCiliumNetworkingSpec()),
			"lyft_vpc":                 OptionalStruct(ResourceLyftVPCNetworkingSpec()),
			"gce":                      OptionalStruct(ResourceGCENetworkingSpec()),
		},
	}

	return res
}

func ExpandResourceNetworkingSpec(in map[string]interface{}) kopsv1alpha2.NetworkingSpec {
	if in == nil {
		panic("expand NetworkingSpec failure, in is nil")
	}
	return kopsv1alpha2.NetworkingSpec{
		NetworkID: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["network_id"]),
		NetworkCIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["network_cidr"]),
		AdditionalNetworkCIDRs: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["additional_network_cidrs"]),
		Subnets: func(in interface{}) []kopsv1alpha2.ClusterSubnetSpec {
			return func(in interface{}) []kopsv1alpha2.ClusterSubnetSpec {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.ClusterSubnetSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.ClusterSubnetSpec {
						if in == nil {
							return kopsv1alpha2.ClusterSubnetSpec{}
						}
						return ExpandResourceClusterSubnetSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["subnet"]),
		TagSubnets: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && len(in) == 1 {
				return func(in interface{}) *bool {
					return func(in interface{}) *bool {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
							return nil
						}
						return func(in bool) *bool {
							return &in
						}(bool(ExpandBool(in)))
					}(in)
				}(in[0].(map[string]interface{})["value"])
			}
			return nil
		}(in["tag_subnets"]),
		Topology: func(in interface{}) *kopsv1alpha2.TopologySpec {
			return func(in interface{}) *kopsv1alpha2.TopologySpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.TopologySpec) *kopsv1alpha2.TopologySpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.TopologySpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceTopologySpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.TopologySpec{}
				}(in))
			}(in)
		}(in["topology"]),
		EgressProxy: func(in interface{}) *kopsv1alpha2.EgressProxySpec {
			return func(in interface{}) *kopsv1alpha2.EgressProxySpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.EgressProxySpec) *kopsv1alpha2.EgressProxySpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.EgressProxySpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceEgressProxySpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.EgressProxySpec{}
				}(in))
			}(in)
		}(in["egress_proxy"]),
		NonMasqueradeCIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["non_masquerade_cidr"]),
		PodCIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["pod_cidr"]),
		ServiceClusterIPRange: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["service_cluster_ip_range"]),
		IsolateControlPlane: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["isolate_control_plane"]),
		Classic: func(in interface{}) *kopsv1alpha2.ClassicNetworkingSpec {
			return func(in interface{}) *kopsv1alpha2.ClassicNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.ClassicNetworkingSpec) *kopsv1alpha2.ClassicNetworkingSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.ClassicNetworkingSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceClassicNetworkingSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.ClassicNetworkingSpec{}
				}(in))
			}(in)
		}(in["classic"]),
		Kubenet: func(in interface{}) *kopsv1alpha2.KubenetNetworkingSpec {
			return func(in interface{}) *kopsv1alpha2.KubenetNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.KubenetNetworkingSpec) *kopsv1alpha2.KubenetNetworkingSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.KubenetNetworkingSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceKubenetNetworkingSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.KubenetNetworkingSpec{}
				}(in))
			}(in)
		}(in["kubenet"]),
		External: func(in interface{}) *kopsv1alpha2.ExternalNetworkingSpec {
			return func(in interface{}) *kopsv1alpha2.ExternalNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.ExternalNetworkingSpec) *kopsv1alpha2.ExternalNetworkingSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.ExternalNetworkingSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceExternalNetworkingSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.ExternalNetworkingSpec{}
				}(in))
			}(in)
		}(in["external"]),
		CNI: func(in interface{}) *kopsv1alpha2.CNINetworkingSpec {
			return func(in interface{}) *kopsv1alpha2.CNINetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.CNINetworkingSpec) *kopsv1alpha2.CNINetworkingSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.CNINetworkingSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceCNINetworkingSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.CNINetworkingSpec{}
				}(in))
			}(in)
		}(in["cni"]),
		Kopeio: func(in interface{}) *kopsv1alpha2.KopeioNetworkingSpec {
			return func(in interface{}) *kopsv1alpha2.KopeioNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.KopeioNetworkingSpec) *kopsv1alpha2.KopeioNetworkingSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.KopeioNetworkingSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceKopeioNetworkingSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.KopeioNetworkingSpec{}
				}(in))
			}(in)
		}(in["kopeio"]),
		Weave: func(in interface{}) *kopsv1alpha2.WeaveNetworkingSpec {
			return func(in interface{}) *kopsv1alpha2.WeaveNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.WeaveNetworkingSpec) *kopsv1alpha2.WeaveNetworkingSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.WeaveNetworkingSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceWeaveNetworkingSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.WeaveNetworkingSpec{}
				}(in))
			}(in)
		}(in["weave"]),
		Flannel: func(in interface{}) *kopsv1alpha2.FlannelNetworkingSpec {
			return func(in interface{}) *kopsv1alpha2.FlannelNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.FlannelNetworkingSpec) *kopsv1alpha2.FlannelNetworkingSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.FlannelNetworkingSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceFlannelNetworkingSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.FlannelNetworkingSpec{}
				}(in))
			}(in)
		}(in["flannel"]),
		Calico: func(in interface{}) *kopsv1alpha2.CalicoNetworkingSpec {
			return func(in interface{}) *kopsv1alpha2.CalicoNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.CalicoNetworkingSpec) *kopsv1alpha2.CalicoNetworkingSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.CalicoNetworkingSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceCalicoNetworkingSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.CalicoNetworkingSpec{}
				}(in))
			}(in)
		}(in["calico"]),
		Canal: func(in interface{}) *kopsv1alpha2.CanalNetworkingSpec {
			return func(in interface{}) *kopsv1alpha2.CanalNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.CanalNetworkingSpec) *kopsv1alpha2.CanalNetworkingSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.CanalNetworkingSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceCanalNetworkingSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.CanalNetworkingSpec{}
				}(in))
			}(in)
		}(in["canal"]),
		KubeRouter: func(in interface{}) *kopsv1alpha2.KuberouterNetworkingSpec {
			return func(in interface{}) *kopsv1alpha2.KuberouterNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.KuberouterNetworkingSpec) *kopsv1alpha2.KuberouterNetworkingSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.KuberouterNetworkingSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceKuberouterNetworkingSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.KuberouterNetworkingSpec{}
				}(in))
			}(in)
		}(in["kube_router"]),
		Romana: func(in interface{}) *kopsv1alpha2.RomanaNetworkingSpec {
			return func(in interface{}) *kopsv1alpha2.RomanaNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.RomanaNetworkingSpec) *kopsv1alpha2.RomanaNetworkingSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.RomanaNetworkingSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceRomanaNetworkingSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.RomanaNetworkingSpec{}
				}(in))
			}(in)
		}(in["romana"]),
		AmazonVPC: func(in interface{}) *kopsv1alpha2.AmazonVPCNetworkingSpec {
			return func(in interface{}) *kopsv1alpha2.AmazonVPCNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.AmazonVPCNetworkingSpec) *kopsv1alpha2.AmazonVPCNetworkingSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.AmazonVPCNetworkingSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceAmazonVPCNetworkingSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.AmazonVPCNetworkingSpec{}
				}(in))
			}(in)
		}(in["amazon_vpc"]),
		Cilium: func(in interface{}) *kopsv1alpha2.CiliumNetworkingSpec {
			return func(in interface{}) *kopsv1alpha2.CiliumNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.CiliumNetworkingSpec) *kopsv1alpha2.CiliumNetworkingSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.CiliumNetworkingSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceCiliumNetworkingSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.CiliumNetworkingSpec{}
				}(in))
			}(in)
		}(in["cilium"]),
		LyftVPC: func(in interface{}) *kopsv1alpha2.LyftVPCNetworkingSpec {
			return func(in interface{}) *kopsv1alpha2.LyftVPCNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.LyftVPCNetworkingSpec) *kopsv1alpha2.LyftVPCNetworkingSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.LyftVPCNetworkingSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceLyftVPCNetworkingSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.LyftVPCNetworkingSpec{}
				}(in))
			}(in)
		}(in["lyft_vpc"]),
		GCE: func(in interface{}) *kopsv1alpha2.GCENetworkingSpec {
			return func(in interface{}) *kopsv1alpha2.GCENetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.GCENetworkingSpec) *kopsv1alpha2.GCENetworkingSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.GCENetworkingSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceGCENetworkingSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.GCENetworkingSpec{}
				}(in))
			}(in)
		}(in["gce"]),
	}
}

func FlattenResourceNetworkingSpecInto(in kopsv1alpha2.NetworkingSpec, out map[string]interface{}) {
	out["network_id"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.NetworkID)
	out["network_cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.NetworkCIDR)
	out["additional_network_cidrs"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AdditionalNetworkCIDRs)
	out["subnet"] = func(in []kopsv1alpha2.ClusterSubnetSpec) interface{} {
		return func(in []kopsv1alpha2.ClusterSubnetSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.ClusterSubnetSpec) interface{} {
					return FlattenResourceClusterSubnetSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.Subnets)
	out["tag_subnets"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return []interface{}{map[string]interface{}{"value": func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)}}
	}(in.TagSubnets)
	out["topology"] = func(in *kopsv1alpha2.TopologySpec) interface{} {
		return func(in *kopsv1alpha2.TopologySpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.TopologySpec) interface{} {
				return func(in kopsv1alpha2.TopologySpec) []interface{} {
					return []interface{}{FlattenResourceTopologySpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Topology)
	out["egress_proxy"] = func(in *kopsv1alpha2.EgressProxySpec) interface{} {
		return func(in *kopsv1alpha2.EgressProxySpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.EgressProxySpec) interface{} {
				return func(in kopsv1alpha2.EgressProxySpec) []interface{} {
					return []interface{}{FlattenResourceEgressProxySpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.EgressProxy)
	out["non_masquerade_cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.NonMasqueradeCIDR)
	out["pod_cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.PodCIDR)
	out["service_cluster_ip_range"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ServiceClusterIPRange)
	out["isolate_control_plane"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.IsolateControlPlane)
	out["classic"] = func(in *kopsv1alpha2.ClassicNetworkingSpec) interface{} {
		return func(in *kopsv1alpha2.ClassicNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.ClassicNetworkingSpec) interface{} {
				return func(in kopsv1alpha2.ClassicNetworkingSpec) []interface{} {
					return []interface{}{FlattenResourceClassicNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Classic)
	out["kubenet"] = func(in *kopsv1alpha2.KubenetNetworkingSpec) interface{} {
		return func(in *kopsv1alpha2.KubenetNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.KubenetNetworkingSpec) interface{} {
				return func(in kopsv1alpha2.KubenetNetworkingSpec) []interface{} {
					return []interface{}{FlattenResourceKubenetNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Kubenet)
	out["external"] = func(in *kopsv1alpha2.ExternalNetworkingSpec) interface{} {
		return func(in *kopsv1alpha2.ExternalNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.ExternalNetworkingSpec) interface{} {
				return func(in kopsv1alpha2.ExternalNetworkingSpec) []interface{} {
					return []interface{}{FlattenResourceExternalNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.External)
	out["cni"] = func(in *kopsv1alpha2.CNINetworkingSpec) interface{} {
		return func(in *kopsv1alpha2.CNINetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.CNINetworkingSpec) interface{} {
				return func(in kopsv1alpha2.CNINetworkingSpec) []interface{} {
					return []interface{}{FlattenResourceCNINetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.CNI)
	out["kopeio"] = func(in *kopsv1alpha2.KopeioNetworkingSpec) interface{} {
		return func(in *kopsv1alpha2.KopeioNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.KopeioNetworkingSpec) interface{} {
				return func(in kopsv1alpha2.KopeioNetworkingSpec) []interface{} {
					return []interface{}{FlattenResourceKopeioNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Kopeio)
	out["weave"] = func(in *kopsv1alpha2.WeaveNetworkingSpec) interface{} {
		return func(in *kopsv1alpha2.WeaveNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.WeaveNetworkingSpec) interface{} {
				return func(in kopsv1alpha2.WeaveNetworkingSpec) []interface{} {
					return []interface{}{FlattenResourceWeaveNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Weave)
	out["flannel"] = func(in *kopsv1alpha2.FlannelNetworkingSpec) interface{} {
		return func(in *kopsv1alpha2.FlannelNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.FlannelNetworkingSpec) interface{} {
				return func(in kopsv1alpha2.FlannelNetworkingSpec) []interface{} {
					return []interface{}{FlattenResourceFlannelNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Flannel)
	out["calico"] = func(in *kopsv1alpha2.CalicoNetworkingSpec) interface{} {
		return func(in *kopsv1alpha2.CalicoNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.CalicoNetworkingSpec) interface{} {
				return func(in kopsv1alpha2.CalicoNetworkingSpec) []interface{} {
					return []interface{}{FlattenResourceCalicoNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Calico)
	out["canal"] = func(in *kopsv1alpha2.CanalNetworkingSpec) interface{} {
		return func(in *kopsv1alpha2.CanalNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.CanalNetworkingSpec) interface{} {
				return func(in kopsv1alpha2.CanalNetworkingSpec) []interface{} {
					return []interface{}{FlattenResourceCanalNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Canal)
	out["kube_router"] = func(in *kopsv1alpha2.KuberouterNetworkingSpec) interface{} {
		return func(in *kopsv1alpha2.KuberouterNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.KuberouterNetworkingSpec) interface{} {
				return func(in kopsv1alpha2.KuberouterNetworkingSpec) []interface{} {
					return []interface{}{FlattenResourceKuberouterNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeRouter)
	out["romana"] = func(in *kopsv1alpha2.RomanaNetworkingSpec) interface{} {
		return func(in *kopsv1alpha2.RomanaNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.RomanaNetworkingSpec) interface{} {
				return func(in kopsv1alpha2.RomanaNetworkingSpec) []interface{} {
					return []interface{}{FlattenResourceRomanaNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Romana)
	out["amazon_vpc"] = func(in *kopsv1alpha2.AmazonVPCNetworkingSpec) interface{} {
		return func(in *kopsv1alpha2.AmazonVPCNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.AmazonVPCNetworkingSpec) interface{} {
				return func(in kopsv1alpha2.AmazonVPCNetworkingSpec) []interface{} {
					return []interface{}{FlattenResourceAmazonVPCNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AmazonVPC)
	out["cilium"] = func(in *kopsv1alpha2.CiliumNetworkingSpec) interface{} {
		return func(in *kopsv1alpha2.CiliumNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.CiliumNetworkingSpec) interface{} {
				return func(in kopsv1alpha2.CiliumNetworkingSpec) []interface{} {
					return []interface{}{FlattenResourceCiliumNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Cilium)
	out["lyft_vpc"] = func(in *kopsv1alpha2.LyftVPCNetworkingSpec) interface{} {
		return func(in *kopsv1alpha2.LyftVPCNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.LyftVPCNetworkingSpec) interface{} {
				return func(in kopsv1alpha2.LyftVPCNetworkingSpec) []interface{} {
					return []interface{}{FlattenResourceLyftVPCNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.LyftVPC)
	out["gce"] = func(in *kopsv1alpha2.GCENetworkingSpec) interface{} {
		return func(in *kopsv1alpha2.GCENetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.GCENetworkingSpec) interface{} {
				return func(in kopsv1alpha2.GCENetworkingSpec) []interface{} {
					return []interface{}{FlattenResourceGCENetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.GCE)
}

func FlattenResourceNetworkingSpec(in kopsv1alpha2.NetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceNetworkingSpecInto(in, out)
	return out
}
