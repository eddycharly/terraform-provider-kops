package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func KopeioAuthenticationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}
