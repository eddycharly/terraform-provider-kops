package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"classic":    OptionalStruct(ClassicNetworkingSpec()),
			"kubenet":    OptionalStruct(KubenetNetworkingSpec()),
			"external":   OptionalStruct(ExternalNetworkingSpec()),
			"cni":        OptionalStruct(CNINetworkingSpec()),
			"kopeio":     OptionalStruct(KopeioNetworkingSpec()),
			"weave":      OptionalStruct(WeaveNetworkingSpec()),
			"flannel":    OptionalStruct(FlannelNetworkingSpec()),
			"calico":     OptionalStruct(CalicoNetworkingSpec()),
			"canal":      OptionalStruct(CanalNetworkingSpec()),
			"kuberouter": OptionalStruct(KuberouterNetworkingSpec()),
			"romana":     OptionalStruct(RomanaNetworkingSpec()),
			"amazon_vpc": OptionalStruct(AmazonVPCNetworkingSpec()),
			"cilium":     OptionalStruct(CiliumNetworkingSpec()),
			"lyft_vpc":   OptionalStruct(LyftVPCNetworkingSpec()),
			"gce":        OptionalStruct(GCENetworkingSpec()),
		},
	}
}
