package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ValidateOptions() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"skip":          OptionalBool(),
			"timeout":       OptionalDuration(),
			"poll_interval": OptionalDuration(),
		},
	}
}
