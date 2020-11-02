package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func RollingUpdate() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"drain_and_terminate": BoolOptionalComputed(),
			"max_unavailable":     IntOrStringOptionalComputed(),
			"max_surge":           IntOrStringOptionalComputed(),
		},
	}
}
