package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func VolumeSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"delete_on_termination": BoolOptionalComputed(),
			"device":                StringOptionalComputed(),
			"encrypted":             BoolOptionalComputed(),
			"iops":                  IntOptionalComputed(),
			"size":                  IntOptionalComputed(),
			"type":                  StringOptionalComputed(),
		},
	}
}
