package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"classic":    ComputedStruct(DataSourceKopsClassicNetworkingSpec()),
			"kubenet":    ComputedStruct(DataSourceKopsKubenetNetworkingSpec()),
			"external":   ComputedStruct(DataSourceKopsExternalNetworkingSpec()),
			"cni":        ComputedStruct(DataSourceKopsCNINetworkingSpec()),
			"kopeio":     ComputedStruct(DataSourceKopsKopeioNetworkingSpec()),
			"weave":      ComputedStruct(DataSourceKopsWeaveNetworkingSpec()),
			"flannel":    ComputedStruct(DataSourceKopsFlannelNetworkingSpec()),
			"calico":     ComputedStruct(DataSourceKopsCalicoNetworkingSpec()),
			"canal":      ComputedStruct(DataSourceKopsCanalNetworkingSpec()),
			"kuberouter": ComputedStruct(DataSourceKopsKuberouterNetworkingSpec()),
			"romana":     ComputedStruct(DataSourceKopsRomanaNetworkingSpec()),
			"amazon_vpc": ComputedStruct(DataSourceKopsAmazonVPCNetworkingSpec()),
			"cilium":     ComputedStruct(DataSourceKopsCiliumNetworkingSpec()),
			"lyft_vpc":   ComputedStruct(DataSourceKopsLyftVPCNetworkingSpec()),
			"gce":        ComputedStruct(DataSourceKopsGCENetworkingSpec()),
		},
	}
}

func ExpandDataSourceKopsNetworkingSpec(in map[string]interface{}) kops.NetworkingSpec {
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.ClassicNetworkingSpec{}
					}
					return (ExpandDataSourceKopsClassicNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.KubenetNetworkingSpec{}
					}
					return (ExpandDataSourceKopsKubenetNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.ExternalNetworkingSpec{}
					}
					return (ExpandDataSourceKopsExternalNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.CNINetworkingSpec{}
					}
					return (ExpandDataSourceKopsCNINetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.KopeioNetworkingSpec{}
					}
					return (ExpandDataSourceKopsKopeioNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.WeaveNetworkingSpec{}
					}
					return (ExpandDataSourceKopsWeaveNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.FlannelNetworkingSpec{}
					}
					return (ExpandDataSourceKopsFlannelNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.CalicoNetworkingSpec{}
					}
					return (ExpandDataSourceKopsCalicoNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.CanalNetworkingSpec{}
					}
					return (ExpandDataSourceKopsCanalNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.KuberouterNetworkingSpec{}
					}
					return (ExpandDataSourceKopsKuberouterNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.RomanaNetworkingSpec{}
					}
					return (ExpandDataSourceKopsRomanaNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.AmazonVPCNetworkingSpec{}
					}
					return (ExpandDataSourceKopsAmazonVPCNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.CiliumNetworkingSpec{}
					}
					return (ExpandDataSourceKopsCiliumNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.LyftVPCNetworkingSpec{}
					}
					return (ExpandDataSourceKopsLyftVPCNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.GCENetworkingSpec{}
					}
					return (ExpandDataSourceKopsGCENetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["gce"]),
	}
}

func FlattenDataSourceKopsNetworkingSpecInto(in kops.NetworkingSpec, out map[string]interface{}) {
	out["classic"] = func(in *kops.ClassicNetworkingSpec) interface{} {
		return func(in *kops.ClassicNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.ClassicNetworkingSpec) interface{} {
				return func(in kops.ClassicNetworkingSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopsClassicNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Classic)
	out["kubenet"] = func(in *kops.KubenetNetworkingSpec) interface{} {
		return func(in *kops.KubenetNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KubenetNetworkingSpec) interface{} {
				return func(in kops.KubenetNetworkingSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopsKubenetNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Kubenet)
	out["external"] = func(in *kops.ExternalNetworkingSpec) interface{} {
		return func(in *kops.ExternalNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.ExternalNetworkingSpec) interface{} {
				return func(in kops.ExternalNetworkingSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopsExternalNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.External)
	out["cni"] = func(in *kops.CNINetworkingSpec) interface{} {
		return func(in *kops.CNINetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.CNINetworkingSpec) interface{} {
				return func(in kops.CNINetworkingSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopsCNINetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.CNI)
	out["kopeio"] = func(in *kops.KopeioNetworkingSpec) interface{} {
		return func(in *kops.KopeioNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KopeioNetworkingSpec) interface{} {
				return func(in kops.KopeioNetworkingSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopsKopeioNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Kopeio)
	out["weave"] = func(in *kops.WeaveNetworkingSpec) interface{} {
		return func(in *kops.WeaveNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.WeaveNetworkingSpec) interface{} {
				return func(in kops.WeaveNetworkingSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopsWeaveNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Weave)
	out["flannel"] = func(in *kops.FlannelNetworkingSpec) interface{} {
		return func(in *kops.FlannelNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.FlannelNetworkingSpec) interface{} {
				return func(in kops.FlannelNetworkingSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopsFlannelNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Flannel)
	out["calico"] = func(in *kops.CalicoNetworkingSpec) interface{} {
		return func(in *kops.CalicoNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.CalicoNetworkingSpec) interface{} {
				return func(in kops.CalicoNetworkingSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopsCalicoNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Calico)
	out["canal"] = func(in *kops.CanalNetworkingSpec) interface{} {
		return func(in *kops.CanalNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.CanalNetworkingSpec) interface{} {
				return func(in kops.CanalNetworkingSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopsCanalNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Canal)
	out["kuberouter"] = func(in *kops.KuberouterNetworkingSpec) interface{} {
		return func(in *kops.KuberouterNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KuberouterNetworkingSpec) interface{} {
				return func(in kops.KuberouterNetworkingSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopsKuberouterNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Kuberouter)
	out["romana"] = func(in *kops.RomanaNetworkingSpec) interface{} {
		return func(in *kops.RomanaNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.RomanaNetworkingSpec) interface{} {
				return func(in kops.RomanaNetworkingSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopsRomanaNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Romana)
	out["amazon_vpc"] = func(in *kops.AmazonVPCNetworkingSpec) interface{} {
		return func(in *kops.AmazonVPCNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AmazonVPCNetworkingSpec) interface{} {
				return func(in kops.AmazonVPCNetworkingSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopsAmazonVPCNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AmazonVPC)
	out["cilium"] = func(in *kops.CiliumNetworkingSpec) interface{} {
		return func(in *kops.CiliumNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.CiliumNetworkingSpec) interface{} {
				return func(in kops.CiliumNetworkingSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopsCiliumNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Cilium)
	out["lyft_vpc"] = func(in *kops.LyftVPCNetworkingSpec) interface{} {
		return func(in *kops.LyftVPCNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.LyftVPCNetworkingSpec) interface{} {
				return func(in kops.LyftVPCNetworkingSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopsLyftVPCNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.LyftVPC)
	out["gce"] = func(in *kops.GCENetworkingSpec) interface{} {
		return func(in *kops.GCENetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.GCENetworkingSpec) interface{} {
				return func(in kops.GCENetworkingSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopsGCENetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.GCE)
}

func FlattenDataSourceKopsNetworkingSpec(in kops.NetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsNetworkingSpecInto(in, out)
	return out
}
