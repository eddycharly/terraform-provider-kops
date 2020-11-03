package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func HTTPProxy() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": RequiredString(),
			"port": RequiredInt(),
		},
	}
}
