package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func LoadBalancerAccessSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type":                       RequiredString(),
			"idle_timeout_seconds":       OptionalInt(),
			"security_group_override":    OptionalString(),
			"additional_security_groups": OptionalList(String()),
			"use_for_internal_api":       OptionalBool(),
			"ssl_certificate":            OptionalString(),
			"cross_zone_load_balancing":  OptionalBool(),
		},
	}
}
