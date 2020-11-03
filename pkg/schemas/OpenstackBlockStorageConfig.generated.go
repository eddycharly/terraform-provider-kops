package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func OpenstackBlockStorageConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"version":     OptionalString(),
			"ignore_az":   OptionalBool(),
			"override_az": OptionalString(),
		},
	}
}
