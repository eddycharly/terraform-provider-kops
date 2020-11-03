package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func OpenstackMonitor() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"delay":       OptionalString(),
			"timeout":     OptionalString(),
			"max_retries": OptionalInt(),
		},
	}
}
