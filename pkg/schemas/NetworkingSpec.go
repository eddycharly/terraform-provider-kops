package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"classic":    StructOptionalComputed(ClassicNetworkingSpec()),
			"kubenet":    StructOptionalComputed(KubenetNetworkingSpec()),
			"external":   StructOptionalComputed(ExternalNetworkingSpec()),
			"cni":        StructOptionalComputed(CNINetworkingSpec()),
			"kopeio":     StructOptionalComputed(KopeioNetworkingSpec()),
			"weave":      StructOptionalComputed(WeaveNetworkingSpec()),
			"flannel":    StructOptionalComputed(FlannelNetworkingSpec()),
			"calico":     StructOptionalComputed(CalicoNetworkingSpec()),
			"canal":      StructOptionalComputed(CanalNetworkingSpec()),
			"kuberouter": StructOptionalComputed(KuberouterNetworkingSpec()),
			"romana":     StructOptionalComputed(RomanaNetworkingSpec()),
			"amazon_vpc": StructOptionalComputed(AmazonVPCNetworkingSpec()),
			"cilium":     StructOptionalComputed(CiliumNetworkingSpec()),
			"lyft_vpc":   StructOptionalComputed(LyftVPCNetworkingSpec()),
			"gce":        StructOptionalComputed(GCENetworkingSpec()),
		},
	}
}
