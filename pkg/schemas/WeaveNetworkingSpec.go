package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func WeaveNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mtu":                IntOptionalComputed(),
			"conn_limit":         IntOptionalComputed(),
			"no_masq_local":      IntOptionalComputed(),
			"memory_request":     QuantityOptionalComputed(),
			"cpu_request":        QuantityOptionalComputed(),
			"memory_limit":       QuantityOptionalComputed(),
			"cpu_limit":          QuantityOptionalComputed(),
			"net_extra_args":     StringOptionalComputed(),
			"npc_memory_request": QuantityOptionalComputed(),
			"npccpu_request":     QuantityOptionalComputed(),
			"npc_memory_limit":   QuantityOptionalComputed(),
			"npccpu_limit":       QuantityOptionalComputed(),
			"npc_extra_args":     StringOptionalComputed(),
		},
	}
}
