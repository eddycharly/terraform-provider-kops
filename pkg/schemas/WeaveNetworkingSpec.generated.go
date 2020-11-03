package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func WeaveNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mtu":                OptionalInt(),
			"conn_limit":         OptionalInt(),
			"no_masq_local":      OptionalInt(),
			"memory_request":     OptionalQuantity(),
			"cpu_request":        OptionalQuantity(),
			"memory_limit":       OptionalQuantity(),
			"cpu_limit":          OptionalQuantity(),
			"net_extra_args":     OptionalString(),
			"npc_memory_request": OptionalQuantity(),
			"npccpu_request":     OptionalQuantity(),
			"npc_memory_limit":   OptionalQuantity(),
			"npccpu_limit":       OptionalQuantity(),
			"npc_extra_args":     OptionalString(),
		},
	}
}
