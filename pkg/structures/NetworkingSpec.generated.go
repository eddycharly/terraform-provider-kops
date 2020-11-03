package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandNetworkingSpec(in map[string]interface{}) kops.NetworkingSpec {
	if in == nil {
		panic("expand NetworkingSpec failure, in is nil")
	}
	return kops.NetworkingSpec{
		Classic: func(in interface{}) *kops.ClassicNetworkingSpec {
			return func(in interface{}) *kops.ClassicNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.ClassicNetworkingSpec) *kops.ClassicNetworkingSpec {
					return &in
				}(func(in interface{}) kops.ClassicNetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.ClassicNetworkingSpec{}
					}
					return (ExpandClassicNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["classic"]),
		Kubenet: func(in interface{}) *kops.KubenetNetworkingSpec {
			return func(in interface{}) *kops.KubenetNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubenetNetworkingSpec) *kops.KubenetNetworkingSpec {
					return &in
				}(func(in interface{}) kops.KubenetNetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.KubenetNetworkingSpec{}
					}
					return (ExpandKubenetNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kubenet"]),
		External: func(in interface{}) *kops.ExternalNetworkingSpec {
			return func(in interface{}) *kops.ExternalNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.ExternalNetworkingSpec) *kops.ExternalNetworkingSpec {
					return &in
				}(func(in interface{}) kops.ExternalNetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.ExternalNetworkingSpec{}
					}
					return (ExpandExternalNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["external"]),
		CNI: func(in interface{}) *kops.CNINetworkingSpec {
			return func(in interface{}) *kops.CNINetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.CNINetworkingSpec) *kops.CNINetworkingSpec {
					return &in
				}(func(in interface{}) kops.CNINetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.CNINetworkingSpec{}
					}
					return (ExpandCNINetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["cni"]),
		Kopeio: func(in interface{}) *kops.KopeioNetworkingSpec {
			return func(in interface{}) *kops.KopeioNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KopeioNetworkingSpec) *kops.KopeioNetworkingSpec {
					return &in
				}(func(in interface{}) kops.KopeioNetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.KopeioNetworkingSpec{}
					}
					return (ExpandKopeioNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kopeio"]),
		Weave: func(in interface{}) *kops.WeaveNetworkingSpec {
			return func(in interface{}) *kops.WeaveNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.WeaveNetworkingSpec) *kops.WeaveNetworkingSpec {
					return &in
				}(func(in interface{}) kops.WeaveNetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.WeaveNetworkingSpec{}
					}
					return (ExpandWeaveNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["weave"]),
		Flannel: func(in interface{}) *kops.FlannelNetworkingSpec {
			return func(in interface{}) *kops.FlannelNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.FlannelNetworkingSpec) *kops.FlannelNetworkingSpec {
					return &in
				}(func(in interface{}) kops.FlannelNetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.FlannelNetworkingSpec{}
					}
					return (ExpandFlannelNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["flannel"]),
		Calico: func(in interface{}) *kops.CalicoNetworkingSpec {
			return func(in interface{}) *kops.CalicoNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.CalicoNetworkingSpec) *kops.CalicoNetworkingSpec {
					return &in
				}(func(in interface{}) kops.CalicoNetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.CalicoNetworkingSpec{}
					}
					return (ExpandCalicoNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["calico"]),
		Canal: func(in interface{}) *kops.CanalNetworkingSpec {
			return func(in interface{}) *kops.CanalNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.CanalNetworkingSpec) *kops.CanalNetworkingSpec {
					return &in
				}(func(in interface{}) kops.CanalNetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.CanalNetworkingSpec{}
					}
					return (ExpandCanalNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["canal"]),
		Kuberouter: func(in interface{}) *kops.KuberouterNetworkingSpec {
			return func(in interface{}) *kops.KuberouterNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KuberouterNetworkingSpec) *kops.KuberouterNetworkingSpec {
					return &in
				}(func(in interface{}) kops.KuberouterNetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.KuberouterNetworkingSpec{}
					}
					return (ExpandKuberouterNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kuberouter"]),
		Romana: func(in interface{}) *kops.RomanaNetworkingSpec {
			return func(in interface{}) *kops.RomanaNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.RomanaNetworkingSpec) *kops.RomanaNetworkingSpec {
					return &in
				}(func(in interface{}) kops.RomanaNetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.RomanaNetworkingSpec{}
					}
					return (ExpandRomanaNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["romana"]),
		AmazonVPC: func(in interface{}) *kops.AmazonVPCNetworkingSpec {
			return func(in interface{}) *kops.AmazonVPCNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AmazonVPCNetworkingSpec) *kops.AmazonVPCNetworkingSpec {
					return &in
				}(func(in interface{}) kops.AmazonVPCNetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.AmazonVPCNetworkingSpec{}
					}
					return (ExpandAmazonVPCNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["amazon_vpc"]),
		Cilium: func(in interface{}) *kops.CiliumNetworkingSpec {
			return func(in interface{}) *kops.CiliumNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.CiliumNetworkingSpec) *kops.CiliumNetworkingSpec {
					return &in
				}(func(in interface{}) kops.CiliumNetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.CiliumNetworkingSpec{}
					}
					return (ExpandCiliumNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["cilium"]),
		LyftVPC: func(in interface{}) *kops.LyftVPCNetworkingSpec {
			return func(in interface{}) *kops.LyftVPCNetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.LyftVPCNetworkingSpec) *kops.LyftVPCNetworkingSpec {
					return &in
				}(func(in interface{}) kops.LyftVPCNetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.LyftVPCNetworkingSpec{}
					}
					return (ExpandLyftVPCNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["lyft_vpc"]),
		GCE: func(in interface{}) *kops.GCENetworkingSpec {
			return func(in interface{}) *kops.GCENetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.GCENetworkingSpec) *kops.GCENetworkingSpec {
					return &in
				}(func(in interface{}) kops.GCENetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.GCENetworkingSpec{}
					}
					return (ExpandGCENetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["gce"]),
	}
}

func FlattenNetworkingSpec(in kops.NetworkingSpec) map[string]interface{} {
	return map[string]interface{}{
		"classic": func(in *kops.ClassicNetworkingSpec) interface{} {
			return func(in *kops.ClassicNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.ClassicNetworkingSpec) interface{} {
					return func(in kops.ClassicNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenClassicNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Classic),
		"kubenet": func(in *kops.KubenetNetworkingSpec) interface{} {
			return func(in *kops.KubenetNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KubenetNetworkingSpec) interface{} {
					return func(in kops.KubenetNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenKubenetNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Kubenet),
		"external": func(in *kops.ExternalNetworkingSpec) interface{} {
			return func(in *kops.ExternalNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.ExternalNetworkingSpec) interface{} {
					return func(in kops.ExternalNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenExternalNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.External),
		"cni": func(in *kops.CNINetworkingSpec) interface{} {
			return func(in *kops.CNINetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.CNINetworkingSpec) interface{} {
					return func(in kops.CNINetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenCNINetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.CNI),
		"kopeio": func(in *kops.KopeioNetworkingSpec) interface{} {
			return func(in *kops.KopeioNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KopeioNetworkingSpec) interface{} {
					return func(in kops.KopeioNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenKopeioNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Kopeio),
		"weave": func(in *kops.WeaveNetworkingSpec) interface{} {
			return func(in *kops.WeaveNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.WeaveNetworkingSpec) interface{} {
					return func(in kops.WeaveNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenWeaveNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Weave),
		"flannel": func(in *kops.FlannelNetworkingSpec) interface{} {
			return func(in *kops.FlannelNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.FlannelNetworkingSpec) interface{} {
					return func(in kops.FlannelNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenFlannelNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Flannel),
		"calico": func(in *kops.CalicoNetworkingSpec) interface{} {
			return func(in *kops.CalicoNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.CalicoNetworkingSpec) interface{} {
					return func(in kops.CalicoNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenCalicoNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Calico),
		"canal": func(in *kops.CanalNetworkingSpec) interface{} {
			return func(in *kops.CanalNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.CanalNetworkingSpec) interface{} {
					return func(in kops.CanalNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenCanalNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Canal),
		"kuberouter": func(in *kops.KuberouterNetworkingSpec) interface{} {
			return func(in *kops.KuberouterNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KuberouterNetworkingSpec) interface{} {
					return func(in kops.KuberouterNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenKuberouterNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Kuberouter),
		"romana": func(in *kops.RomanaNetworkingSpec) interface{} {
			return func(in *kops.RomanaNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.RomanaNetworkingSpec) interface{} {
					return func(in kops.RomanaNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenRomanaNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Romana),
		"amazon_vpc": func(in *kops.AmazonVPCNetworkingSpec) interface{} {
			return func(in *kops.AmazonVPCNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.AmazonVPCNetworkingSpec) interface{} {
					return func(in kops.AmazonVPCNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenAmazonVPCNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.AmazonVPC),
		"cilium": func(in *kops.CiliumNetworkingSpec) interface{} {
			return func(in *kops.CiliumNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.CiliumNetworkingSpec) interface{} {
					return func(in kops.CiliumNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenCiliumNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Cilium),
		"lyft_vpc": func(in *kops.LyftVPCNetworkingSpec) interface{} {
			return func(in *kops.LyftVPCNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.LyftVPCNetworkingSpec) interface{} {
					return func(in kops.LyftVPCNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenLyftVPCNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.LyftVPC),
		"gce": func(in *kops.GCENetworkingSpec) interface{} {
			return func(in *kops.GCENetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.GCENetworkingSpec) interface{} {
					return func(in kops.GCENetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenGCENetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.GCE),
	}
}
