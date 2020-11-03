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
			value := func(in interface{}) *kops.ClassicNetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
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
			return value
		}(in["classic"]),
		Kubenet: func(in interface{}) *kops.KubenetNetworkingSpec {
			value := func(in interface{}) *kops.KubenetNetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
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
			return value
		}(in["kubenet"]),
		External: func(in interface{}) *kops.ExternalNetworkingSpec {
			value := func(in interface{}) *kops.ExternalNetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
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
			return value
		}(in["external"]),
		CNI: func(in interface{}) *kops.CNINetworkingSpec {
			value := func(in interface{}) *kops.CNINetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
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
			return value
		}(in["cni"]),
		Kopeio: func(in interface{}) *kops.KopeioNetworkingSpec {
			value := func(in interface{}) *kops.KopeioNetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
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
			return value
		}(in["kopeio"]),
		Weave: func(in interface{}) *kops.WeaveNetworkingSpec {
			value := func(in interface{}) *kops.WeaveNetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
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
			return value
		}(in["weave"]),
		Flannel: func(in interface{}) *kops.FlannelNetworkingSpec {
			value := func(in interface{}) *kops.FlannelNetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
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
			return value
		}(in["flannel"]),
		Calico: func(in interface{}) *kops.CalicoNetworkingSpec {
			value := func(in interface{}) *kops.CalicoNetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
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
			return value
		}(in["calico"]),
		Canal: func(in interface{}) *kops.CanalNetworkingSpec {
			value := func(in interface{}) *kops.CanalNetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
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
			return value
		}(in["canal"]),
		Kuberouter: func(in interface{}) *kops.KuberouterNetworkingSpec {
			value := func(in interface{}) *kops.KuberouterNetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
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
			return value
		}(in["kuberouter"]),
		Romana: func(in interface{}) *kops.RomanaNetworkingSpec {
			value := func(in interface{}) *kops.RomanaNetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
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
			return value
		}(in["romana"]),
		AmazonVPC: func(in interface{}) *kops.AmazonVPCNetworkingSpec {
			value := func(in interface{}) *kops.AmazonVPCNetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
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
			return value
		}(in["amazon_vpc"]),
		Cilium: func(in interface{}) *kops.CiliumNetworkingSpec {
			value := func(in interface{}) *kops.CiliumNetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
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
			return value
		}(in["cilium"]),
		LyftVPC: func(in interface{}) *kops.LyftVPCNetworkingSpec {
			value := func(in interface{}) *kops.LyftVPCNetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
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
			return value
		}(in["lyft_vpc"]),
		GCE: func(in interface{}) *kops.GCENetworkingSpec {
			value := func(in interface{}) *kops.GCENetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
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
			return value
		}(in["gce"]),
	}
}

func FlattenNetworkingSpec(in kops.NetworkingSpec) map[string]interface{} {
	return map[string]interface{}{
		"classic": func(in *kops.ClassicNetworkingSpec) interface{} {
			value := func(in *kops.ClassicNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.ClassicNetworkingSpec) interface{} {
					return func(in kops.ClassicNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenClassicNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Classic),
		"kubenet": func(in *kops.KubenetNetworkingSpec) interface{} {
			value := func(in *kops.KubenetNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KubenetNetworkingSpec) interface{} {
					return func(in kops.KubenetNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenKubenetNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Kubenet),
		"external": func(in *kops.ExternalNetworkingSpec) interface{} {
			value := func(in *kops.ExternalNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.ExternalNetworkingSpec) interface{} {
					return func(in kops.ExternalNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenExternalNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.External),
		"cni": func(in *kops.CNINetworkingSpec) interface{} {
			value := func(in *kops.CNINetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.CNINetworkingSpec) interface{} {
					return func(in kops.CNINetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenCNINetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.CNI),
		"kopeio": func(in *kops.KopeioNetworkingSpec) interface{} {
			value := func(in *kops.KopeioNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KopeioNetworkingSpec) interface{} {
					return func(in kops.KopeioNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenKopeioNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Kopeio),
		"weave": func(in *kops.WeaveNetworkingSpec) interface{} {
			value := func(in *kops.WeaveNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.WeaveNetworkingSpec) interface{} {
					return func(in kops.WeaveNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenWeaveNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Weave),
		"flannel": func(in *kops.FlannelNetworkingSpec) interface{} {
			value := func(in *kops.FlannelNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.FlannelNetworkingSpec) interface{} {
					return func(in kops.FlannelNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenFlannelNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Flannel),
		"calico": func(in *kops.CalicoNetworkingSpec) interface{} {
			value := func(in *kops.CalicoNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.CalicoNetworkingSpec) interface{} {
					return func(in kops.CalicoNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenCalicoNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Calico),
		"canal": func(in *kops.CanalNetworkingSpec) interface{} {
			value := func(in *kops.CanalNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.CanalNetworkingSpec) interface{} {
					return func(in kops.CanalNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenCanalNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Canal),
		"kuberouter": func(in *kops.KuberouterNetworkingSpec) interface{} {
			value := func(in *kops.KuberouterNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KuberouterNetworkingSpec) interface{} {
					return func(in kops.KuberouterNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenKuberouterNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Kuberouter),
		"romana": func(in *kops.RomanaNetworkingSpec) interface{} {
			value := func(in *kops.RomanaNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.RomanaNetworkingSpec) interface{} {
					return func(in kops.RomanaNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenRomanaNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Romana),
		"amazon_vpc": func(in *kops.AmazonVPCNetworkingSpec) interface{} {
			value := func(in *kops.AmazonVPCNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.AmazonVPCNetworkingSpec) interface{} {
					return func(in kops.AmazonVPCNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenAmazonVPCNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.AmazonVPC),
		"cilium": func(in *kops.CiliumNetworkingSpec) interface{} {
			value := func(in *kops.CiliumNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.CiliumNetworkingSpec) interface{} {
					return func(in kops.CiliumNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenCiliumNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Cilium),
		"lyft_vpc": func(in *kops.LyftVPCNetworkingSpec) interface{} {
			value := func(in *kops.LyftVPCNetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.LyftVPCNetworkingSpec) interface{} {
					return func(in kops.LyftVPCNetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenLyftVPCNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.LyftVPC),
		"gce": func(in *kops.GCENetworkingSpec) interface{} {
			value := func(in *kops.GCENetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.GCENetworkingSpec) interface{} {
					return func(in kops.GCENetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenGCENetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.GCE),
	}
}
