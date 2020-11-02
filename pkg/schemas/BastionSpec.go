package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func BastionSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bastion_public_name":  StringRequired(),
			"idle_timeout_seconds": IntOptionalComputed(),
			"load_balancer":        StructOptionalComputed(BastionLoadBalancerSpec()),
		},
	}
}
