package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func IAMSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"legacy":                   OptionalBool(),
			"allow_container_registry": OptionalBool(),
		},
	}
}
