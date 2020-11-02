package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func MixedInstancesPolicySpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"instances":                     ListOptional(String()),
			"on_demand_allocation_strategy": StringOptionalComputed(),
			"on_demand_base":                IntOptionalComputed(),
			"on_demand_above_base":          IntOptionalComputed(),
			"spot_allocation_strategy":      StringOptionalComputed(),
			"spot_instance_pools":           IntOptionalComputed(),
		},
	}
}
