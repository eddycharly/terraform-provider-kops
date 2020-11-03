package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func MixedInstancesPolicySpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"instances":                     OptionalList(String()),
			"on_demand_allocation_strategy": OptionalString(),
			"on_demand_base":                OptionalInt(),
			"on_demand_above_base":          OptionalInt(),
			"spot_allocation_strategy":      OptionalString(),
			"spot_instance_pools":           OptionalInt(),
		},
	}
}
