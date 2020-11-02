package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ExecContainerAction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image":       StringRequired(),
			"command":     ListRequired(String()),
			"environment": MapOptionalComputed(String()),
		},
	}
}
