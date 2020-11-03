package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func VolumeMountSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"device":         RequiredString(),
			"filesystem":     RequiredString(),
			"format_options": OptionalList(String()),
			"mount_options":  OptionalList(String()),
			"path":           RequiredString(),
		},
	}
}
