package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"classic":    Optional(Struct(ResourceClassicNetworkingSpec())),
			"kubenet":    Optional(Struct(ResourceKubenetNetworkingSpec())),
			"external":   Optional(Struct(ResourceExternalNetworkingSpec())),
			"cni":        Optional(Struct(ResourceCNINetworkingSpec())),
			"kopeio":     Optional(Struct(ResourceKopeioNetworkingSpec())),
			"weave":      Optional(Struct(ResourceWeaveNetworkingSpec())),
			"flannel":    Optional(Struct(ResourceFlannelNetworkingSpec())),
			"calico":     Optional(Struct(ResourceCalicoNetworkingSpec())),
			"canal":      Optional(Struct(ResourceCanalNetworkingSpec())),
			"kuberouter": Optional(Struct(ResourceKuberouterNetworkingSpec())),
			"romana":     Optional(Struct(ResourceRomanaNetworkingSpec())),
			"amazon_vpc": Optional(Struct(ResourceAmazonVPCNetworkingSpec())),
			"cilium":     Optional(Struct(ResourceCiliumNetworkingSpec())),
			"lyft_vpc":   Optional(Struct(ResourceLyftVPCNetworkingSpec())),
			"gce":        Optional(Struct(ResourceGCENetworkingSpec())),
		},
	}
}

func ExpandResourceNetworkingSpec(in map[string]interface{}) kops.NetworkingSpec {
	if in == nil {
		panic("expand NetworkingSpec failure, in is nil")
	}
	out := kops.NetworkingSpec{}
	if in, ok := in["classic"]; ok && in != nil {
		out.Classic = func(in interface{}) *kops.ClassicNetworkingSpec {
			if in == nil {
				return nil
			}
			return func(in kops.ClassicNetworkingSpec) *kops.ClassicNetworkingSpec { return &in }(func(in interface{}) kops.ClassicNetworkingSpec {
				if in == nil {
					return kops.ClassicNetworkingSpec{}
				}
				return ExpandResourceClassicNetworkingSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["kubenet"]; ok && in != nil {
		out.Kubenet = func(in interface{}) *kops.KubenetNetworkingSpec {
			if in == nil {
				return nil
			}
			return func(in kops.KubenetNetworkingSpec) *kops.KubenetNetworkingSpec { return &in }(func(in interface{}) kops.KubenetNetworkingSpec {
				if in == nil {
					return kops.KubenetNetworkingSpec{}
				}
				return ExpandResourceKubenetNetworkingSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["external"]; ok && in != nil {
		out.External = func(in interface{}) *kops.ExternalNetworkingSpec {
			if in == nil {
				return nil
			}
			return func(in kops.ExternalNetworkingSpec) *kops.ExternalNetworkingSpec { return &in }(func(in interface{}) kops.ExternalNetworkingSpec {
				if in == nil {
					return kops.ExternalNetworkingSpec{}
				}
				return ExpandResourceExternalNetworkingSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["cni"]; ok && in != nil {
		out.CNI = func(in interface{}) *kops.CNINetworkingSpec {
			if in == nil {
				return nil
			}
			return func(in kops.CNINetworkingSpec) *kops.CNINetworkingSpec { return &in }(func(in interface{}) kops.CNINetworkingSpec {
				if in == nil {
					return kops.CNINetworkingSpec{}
				}
				return ExpandResourceCNINetworkingSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["kopeio"]; ok && in != nil {
		out.Kopeio = func(in interface{}) *kops.KopeioNetworkingSpec {
			if in == nil {
				return nil
			}
			return func(in kops.KopeioNetworkingSpec) *kops.KopeioNetworkingSpec { return &in }(func(in interface{}) kops.KopeioNetworkingSpec {
				if in == nil {
					return kops.KopeioNetworkingSpec{}
				}
				return ExpandResourceKopeioNetworkingSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["weave"]; ok && in != nil {
		out.Weave = func(in interface{}) *kops.WeaveNetworkingSpec {
			if in == nil {
				return nil
			}
			return func(in kops.WeaveNetworkingSpec) *kops.WeaveNetworkingSpec { return &in }(func(in interface{}) kops.WeaveNetworkingSpec {
				if in == nil {
					return kops.WeaveNetworkingSpec{}
				}
				return ExpandResourceWeaveNetworkingSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["flannel"]; ok && in != nil {
		out.Flannel = func(in interface{}) *kops.FlannelNetworkingSpec {
			if in == nil {
				return nil
			}
			return func(in kops.FlannelNetworkingSpec) *kops.FlannelNetworkingSpec { return &in }(func(in interface{}) kops.FlannelNetworkingSpec {
				if in == nil {
					return kops.FlannelNetworkingSpec{}
				}
				return ExpandResourceFlannelNetworkingSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["calico"]; ok && in != nil {
		out.Calico = func(in interface{}) *kops.CalicoNetworkingSpec {
			if in == nil {
				return nil
			}
			return func(in kops.CalicoNetworkingSpec) *kops.CalicoNetworkingSpec { return &in }(func(in interface{}) kops.CalicoNetworkingSpec {
				if in == nil {
					return kops.CalicoNetworkingSpec{}
				}
				return ExpandResourceCalicoNetworkingSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["canal"]; ok && in != nil {
		out.Canal = func(in interface{}) *kops.CanalNetworkingSpec {
			if in == nil {
				return nil
			}
			return func(in kops.CanalNetworkingSpec) *kops.CanalNetworkingSpec { return &in }(func(in interface{}) kops.CanalNetworkingSpec {
				if in == nil {
					return kops.CanalNetworkingSpec{}
				}
				return ExpandResourceCanalNetworkingSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["kuberouter"]; ok && in != nil {
		out.Kuberouter = func(in interface{}) *kops.KuberouterNetworkingSpec {
			if in == nil {
				return nil
			}
			return func(in kops.KuberouterNetworkingSpec) *kops.KuberouterNetworkingSpec { return &in }(func(in interface{}) kops.KuberouterNetworkingSpec {
				if in == nil {
					return kops.KuberouterNetworkingSpec{}
				}
				return ExpandResourceKuberouterNetworkingSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["romana"]; ok && in != nil {
		out.Romana = func(in interface{}) *kops.RomanaNetworkingSpec {
			if in == nil {
				return nil
			}
			return func(in kops.RomanaNetworkingSpec) *kops.RomanaNetworkingSpec { return &in }(func(in interface{}) kops.RomanaNetworkingSpec {
				if in == nil {
					return kops.RomanaNetworkingSpec{}
				}
				return ExpandResourceRomanaNetworkingSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["amazon_vpc"]; ok && in != nil {
		out.AmazonVPC = func(in interface{}) *kops.AmazonVPCNetworkingSpec {
			if in == nil {
				return nil
			}
			return func(in kops.AmazonVPCNetworkingSpec) *kops.AmazonVPCNetworkingSpec { return &in }(func(in interface{}) kops.AmazonVPCNetworkingSpec {
				if in == nil {
					return kops.AmazonVPCNetworkingSpec{}
				}
				return ExpandResourceAmazonVPCNetworkingSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["cilium"]; ok && in != nil {
		out.Cilium = func(in interface{}) *kops.CiliumNetworkingSpec {
			if in == nil {
				return nil
			}
			return func(in kops.CiliumNetworkingSpec) *kops.CiliumNetworkingSpec { return &in }(func(in interface{}) kops.CiliumNetworkingSpec {
				if in == nil {
					return kops.CiliumNetworkingSpec{}
				}
				return ExpandResourceCiliumNetworkingSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["lyft_vpc"]; ok && in != nil {
		out.LyftVPC = func(in interface{}) *kops.LyftVPCNetworkingSpec {
			if in == nil {
				return nil
			}
			return func(in kops.LyftVPCNetworkingSpec) *kops.LyftVPCNetworkingSpec { return &in }(func(in interface{}) kops.LyftVPCNetworkingSpec {
				if in == nil {
					return kops.LyftVPCNetworkingSpec{}
				}
				return ExpandResourceLyftVPCNetworkingSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["gce"]; ok && in != nil {
		out.GCE = func(in interface{}) *kops.GCENetworkingSpec {
			if in == nil {
				return nil
			}
			return func(in kops.GCENetworkingSpec) *kops.GCENetworkingSpec { return &in }(func(in interface{}) kops.GCENetworkingSpec {
				if in == nil {
					return kops.GCENetworkingSpec{}
				}
				return ExpandResourceGCENetworkingSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	return out
}

func FlattenResourceNetworkingSpecInto(in kops.NetworkingSpec, out map[string]interface{}) {
	out["classic"] = func(in *kops.ClassicNetworkingSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.ClassicNetworkingSpec) interface{} { return FlattenResourceClassicNetworkingSpec(in) }(*in)
	}(in.Classic)
	out["kubenet"] = func(in *kops.KubenetNetworkingSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.KubenetNetworkingSpec) interface{} { return FlattenResourceKubenetNetworkingSpec(in) }(*in)
	}(in.Kubenet)
	out["external"] = func(in *kops.ExternalNetworkingSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.ExternalNetworkingSpec) interface{} { return FlattenResourceExternalNetworkingSpec(in) }(*in)
	}(in.External)
	out["cni"] = func(in *kops.CNINetworkingSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.CNINetworkingSpec) interface{} { return FlattenResourceCNINetworkingSpec(in) }(*in)
	}(in.CNI)
	out["kopeio"] = func(in *kops.KopeioNetworkingSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.KopeioNetworkingSpec) interface{} { return FlattenResourceKopeioNetworkingSpec(in) }(*in)
	}(in.Kopeio)
	out["weave"] = func(in *kops.WeaveNetworkingSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.WeaveNetworkingSpec) interface{} { return FlattenResourceWeaveNetworkingSpec(in) }(*in)
	}(in.Weave)
	out["flannel"] = func(in *kops.FlannelNetworkingSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.FlannelNetworkingSpec) interface{} { return FlattenResourceFlannelNetworkingSpec(in) }(*in)
	}(in.Flannel)
	out["calico"] = func(in *kops.CalicoNetworkingSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.CalicoNetworkingSpec) interface{} { return FlattenResourceCalicoNetworkingSpec(in) }(*in)
	}(in.Calico)
	out["canal"] = func(in *kops.CanalNetworkingSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.CanalNetworkingSpec) interface{} { return FlattenResourceCanalNetworkingSpec(in) }(*in)
	}(in.Canal)
	out["kuberouter"] = func(in *kops.KuberouterNetworkingSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.KuberouterNetworkingSpec) interface{} { return FlattenResourceKuberouterNetworkingSpec(in) }(*in)
	}(in.Kuberouter)
	out["romana"] = func(in *kops.RomanaNetworkingSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.RomanaNetworkingSpec) interface{} { return FlattenResourceRomanaNetworkingSpec(in) }(*in)
	}(in.Romana)
	out["amazon_vpc"] = func(in *kops.AmazonVPCNetworkingSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.AmazonVPCNetworkingSpec) interface{} { return FlattenResourceAmazonVPCNetworkingSpec(in) }(*in)
	}(in.AmazonVPC)
	out["cilium"] = func(in *kops.CiliumNetworkingSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.CiliumNetworkingSpec) interface{} { return FlattenResourceCiliumNetworkingSpec(in) }(*in)
	}(in.Cilium)
	out["lyft_vpc"] = func(in *kops.LyftVPCNetworkingSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.LyftVPCNetworkingSpec) interface{} { return FlattenResourceLyftVPCNetworkingSpec(in) }(*in)
	}(in.LyftVPC)
	out["gce"] = func(in *kops.GCENetworkingSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.GCENetworkingSpec) interface{} { return FlattenResourceGCENetworkingSpec(in) }(*in)
	}(in.GCE)
}

func FlattenResourceNetworkingSpec(in kops.NetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceNetworkingSpecInto(in, out)
	return out
}
