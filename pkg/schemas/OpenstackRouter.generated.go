package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func OpenstackRouter() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external_network": OptionalString(),
			"dns_servers":      OptionalString(),
			"external_subnet":  OptionalString(),
		},
	}
}
