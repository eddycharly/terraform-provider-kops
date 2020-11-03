package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func VolumeSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"delete_on_termination": OptionalBool(),
			"device":                RequiredString(),
			"encrypted":             OptionalBool(),
			"iops":                  OptionalInt(),
			"size":                  OptionalInt(),
			"type":                  OptionalString(),
		},
	}
}
