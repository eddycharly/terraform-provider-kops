package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Assets() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"container_registry": OptionalString(),
			"file_repository":    OptionalString(),
			"container_proxy":    OptionalString(),
		},
	}
}
