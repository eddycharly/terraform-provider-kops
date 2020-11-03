package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ContainerdConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"address":         OptionalString(),
			"config_override": OptionalString(),
			"log_level":       OptionalString(),
			"root":            OptionalString(),
			"skip_install":    OptionalBool(),
			"state":           OptionalString(),
			"version":         OptionalString(),
		},
	}
}
