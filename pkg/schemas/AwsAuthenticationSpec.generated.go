package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AwsAuthenticationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image":          OptionalString(),
			"memory_request": OptionalQuantity(),
			"cpu_request":    OptionalQuantity(),
			"memory_limit":   OptionalQuantity(),
			"cpu_limit":      OptionalQuantity(),
		},
	}
}
