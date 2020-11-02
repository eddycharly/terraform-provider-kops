package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func VolumeMountSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"device":         StringOptionalComputed(),
			"filesystem":     StringOptionalComputed(),
			"format_options": ListOptional(String()),
			"mount_options":  ListOptional(String()),
			"path":           StringRequired(),
		},
	}
}
