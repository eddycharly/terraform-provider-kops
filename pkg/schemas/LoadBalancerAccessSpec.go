package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func LoadBalancerAccessSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type":                       StringOptionalComputed(),
			"idle_timeout_seconds":       IntOptionalComputed(),
			"security_group_override":    StringOptionalComputed(),
			"additional_security_groups": ListOptional(String()),
			"use_for_internal_api":       BoolOptionalComputed(),
			"ssl_certificate":            StringOptionalComputed(),
			"cross_zone_load_balancing":  BoolOptionalComputed(),
		},
	}
}
