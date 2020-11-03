package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func FlannelNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backend":                        OptionalString(),
			"disable_tx_checksum_offloading": OptionalBool(),
			"iptables_resync_seconds":        OptionalInt(),
		},
	}
}
