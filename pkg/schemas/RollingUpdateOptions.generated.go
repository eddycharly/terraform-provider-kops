package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func RollingUpdateOptions() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"skip":                OptionalBool(),
			"master_interval":     OptionalDuration(),
			"node_interval":       OptionalDuration(),
			"bastion_interval":    OptionalDuration(),
			"fail_on_drain_error": OptionalBool(),
			"fail_on_validate":    OptionalBool(),
			"post_drain_delay":    OptionalDuration(),
			"validation_timeout":  OptionalDuration(),
			"validate_count":      OptionalInt(),
		},
	}
}
