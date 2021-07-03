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
			"classic":    Optional(Ptr(Struct(ResourceClassicNetworkingSpec()))),
			"kubenet":    Optional(Ptr(Struct(ResourceKubenetNetworkingSpec()))),
			"external":   Optional(Ptr(Struct(ResourceExternalNetworkingSpec()))),
			"cni":        Optional(Ptr(Struct(ResourceCNINetworkingSpec()))),
			"kopeio":     Optional(Ptr(Struct(ResourceKopeioNetworkingSpec()))),
			"weave":      Optional(Ptr(Struct(ResourceWeaveNetworkingSpec()))),
			"flannel":    Optional(Ptr(Struct(ResourceFlannelNetworkingSpec()))),
			"calico":     Optional(Ptr(Struct(ResourceCalicoNetworkingSpec()))),
			"canal":      Optional(Ptr(Struct(ResourceCanalNetworkingSpec()))),
			"kuberouter": Optional(Ptr(Struct(ResourceKuberouterNetworkingSpec()))),
			"romana":     Optional(Ptr(Struct(ResourceRomanaNetworkingSpec()))),
			"amazon_vpc": Optional(Ptr(Struct(ResourceAmazonVPCNetworkingSpec()))),
			"cilium":     Optional(Ptr(Struct(ResourceCiliumNetworkingSpec()))),
			"lyft_vpc":   Optional(Ptr(Struct(ResourceLyftVPCNetworkingSpec()))),
			"gce":        Optional(Ptr(Struct(ResourceGCENetworkingSpec()))),
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
