package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func OpenstackLoadbalancerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"method":              OptionalString(),
			"provider":            OptionalString(),
			"use_octavia":         OptionalBool(),
			"floating_network":    OptionalString(),
			"floating_network_id": OptionalString(),
			"floating_subnet":     OptionalString(),
			"subnet_id":           OptionalString(),
			"manage_sec_groups":   OptionalBool(),
		},
	}
}
