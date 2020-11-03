package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NodeLocalDNSConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":        OptionalBool(),
			"local_ip":       OptionalString(),
			"memory_request": OptionalQuantity(),
			"cpu_request":    OptionalQuantity(),
		},
	}
}
