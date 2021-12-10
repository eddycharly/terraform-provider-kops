package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"classic":    ComputedStruct(DataSourceClassicNetworkingSpec()),
			"kubenet":    ComputedStruct(DataSourceKubenetNetworkingSpec()),
			"external":   ComputedStruct(DataSourceExternalNetworkingSpec()),
			"cni":        ComputedStruct(DataSourceCNINetworkingSpec()),
			"kopeio":     ComputedStruct(DataSourceKopeioNetworkingSpec()),
			"weave":      ComputedStruct(DataSourceWeaveNetworkingSpec()),
			"flannel":    ComputedStruct(DataSourceFlannelNetworkingSpec()),
			"calico":     ComputedStruct(DataSourceCalicoNetworkingSpec()),
			"canal":      ComputedStruct(DataSourceCanalNetworkingSpec()),
			"kuberouter": ComputedStruct(DataSourceKuberouterNetworkingSpec()),
			"romana":     ComputedStruct(DataSourceRomanaNetworkingSpec()),
			"amazon_vpc": ComputedStruct(DataSourceAmazonVPCNetworkingSpec()),
			"cilium":     ComputedStruct(DataSourceCiliumNetworkingSpec()),
			"lyft_vpc":   ComputedStruct(DataSourceLyftVPCNetworkingSpec()),
			"gce":        ComputedStruct(DataSourceGCENetworkingSpec()),
		},
	}

	return res
}

func ExpandDataSourceNetworkingSpec(in map[string]interface{}) kops.NetworkingSpec {
	if in == nil {
		panic("expand NetworkingSpec failure, in is nil")
	}
	return kops.NetworkingSpec{
		Classic: func(in interface{}) *kops.ClassicNetworkingSpec /*k8s.io/kops/pkg/apis/kops*/ {
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
					return (ExpandDataSourceClassicNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["classic"]),
		Kubenet: func(in interface{}) *kops.KubenetNetworkingSpec /*k8s.io/kops/pkg/apis/kops*/ {
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
					return (ExpandDataSourceKubenetNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kubenet"]),
		External: func(in interface{}) *kops.ExternalNetworkingSpec /*k8s.io/kops/pkg/apis/kops*/ {
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
					return (ExpandDataSourceExternalNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["external"]),
		CNI: func(in interface{}) *kops.CNINetworkingSpec /*k8s.io/kops/pkg/apis/kops*/ {
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
					return (ExpandDataSourceCNINetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["cni"]),
		Kopeio: func(in interface{}) *kops.KopeioNetworkingSpec /*k8s.io/kops/pkg/apis/kops*/ {
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
					return (ExpandDataSourceKopeioNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kopeio"]),
		Weave: func(in interface{}) *kops.WeaveNetworkingSpec /*k8s.io/kops/pkg/apis/kops*/ {
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
					return (ExpandDataSourceWeaveNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["weave"]),
		Flannel: func(in interface{}) *kops.FlannelNetworkingSpec /*k8s.io/kops/pkg/apis/kops*/ {
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
					return (ExpandDataSourceFlannelNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["flannel"]),
		Calico: func(in interface{}) *kops.CalicoNetworkingSpec /*k8s.io/kops/pkg/apis/kops*/ {
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
					return (ExpandDataSourceCalicoNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["calico"]),
		Canal: func(in interface{}) *kops.CanalNetworkingSpec /*k8s.io/kops/pkg/apis/kops*/ {
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
					return (ExpandDataSourceCanalNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["canal"]),
		Kuberouter: func(in interface{}) *kops.KuberouterNetworkingSpec /*k8s.io/kops/pkg/apis/kops*/ {
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
					return (ExpandDataSourceKuberouterNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kuberouter"]),
		Romana: func(in interface{}) *kops.RomanaNetworkingSpec /*k8s.io/kops/pkg/apis/kops*/ {
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
					return (ExpandDataSourceRomanaNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["romana"]),
		AmazonVPC: func(in interface{}) *kops.AmazonVPCNetworkingSpec /*k8s.io/kops/pkg/apis/kops*/ {
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
					return (ExpandDataSourceAmazonVPCNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["amazon_vpc"]),
		Cilium: func(in interface{}) *kops.CiliumNetworkingSpec /*k8s.io/kops/pkg/apis/kops*/ {
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
					return (ExpandDataSourceCiliumNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["cilium"]),
		LyftVPC: func(in interface{}) *kops.LyftVPCNetworkingSpec /*k8s.io/kops/pkg/apis/kops*/ {
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
					return (ExpandDataSourceLyftVPCNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["lyft_vpc"]),
		GCE: func(in interface{}) *kops.GCENetworkingSpec /*k8s.io/kops/pkg/apis/kops*/ {
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
					return (ExpandDataSourceGCENetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["gce"]),
	}
}

func FlattenDataSourceNetworkingSpecInto(in kops.NetworkingSpec, out map[string]interface{}) {
	out["classic"] = func(in *kops.ClassicNetworkingSpec) interface{} {
		return func(in *kops.ClassicNetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.ClassicNetworkingSpec) interface{} {
				return func(in kops.ClassicNetworkingSpec) []interface{} {
					return []interface{}{FlattenDataSourceClassicNetworkingSpec(in)}
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
				return func(in kops.KubenetNetworkingSpec) []interface{} {
					return []interface{}{FlattenDataSourceKubenetNetworkingSpec(in)}
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
				return func(in kops.ExternalNetworkingSpec) []interface{} {
					return []interface{}{FlattenDataSourceExternalNetworkingSpec(in)}
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
				return func(in kops.CNINetworkingSpec) []interface{} {
					return []interface{}{FlattenDataSourceCNINetworkingSpec(in)}
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
				return func(in kops.KopeioNetworkingSpec) []interface{} {
					return []interface{}{FlattenDataSourceKopeioNetworkingSpec(in)}
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
				return func(in kops.WeaveNetworkingSpec) []interface{} {
					return []interface{}{FlattenDataSourceWeaveNetworkingSpec(in)}
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
				return func(in kops.FlannelNetworkingSpec) []interface{} {
					return []interface{}{FlattenDataSourceFlannelNetworkingSpec(in)}
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
				return func(in kops.CalicoNetworkingSpec) []interface{} {
					return []interface{}{FlattenDataSourceCalicoNetworkingSpec(in)}
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
				return func(in kops.CanalNetworkingSpec) []interface{} {
					return []interface{}{FlattenDataSourceCanalNetworkingSpec(in)}
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
				return func(in kops.KuberouterNetworkingSpec) []interface{} {
					return []interface{}{FlattenDataSourceKuberouterNetworkingSpec(in)}
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
				return func(in kops.RomanaNetworkingSpec) []interface{} {
					return []interface{}{FlattenDataSourceRomanaNetworkingSpec(in)}
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
				return func(in kops.AmazonVPCNetworkingSpec) []interface{} {
					return []interface{}{FlattenDataSourceAmazonVPCNetworkingSpec(in)}
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
				return func(in kops.CiliumNetworkingSpec) []interface{} {
					return []interface{}{FlattenDataSourceCiliumNetworkingSpec(in)}
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
				return func(in kops.LyftVPCNetworkingSpec) []interface{} {
					return []interface{}{FlattenDataSourceLyftVPCNetworkingSpec(in)}
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
				return func(in kops.GCENetworkingSpec) []interface{} {
					return []interface{}{FlattenDataSourceGCENetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.GCE)
}

func FlattenDataSourceNetworkingSpec(in kops.NetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceNetworkingSpecInto(in, out)
	return out
}
