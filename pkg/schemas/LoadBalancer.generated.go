package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func LoadBalancer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"load_balancer_name": OptionalString(),
			"target_group_arn":   OptionalString(),
		},
	}
}
