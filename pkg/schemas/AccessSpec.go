package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AccessSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns":           StructOptionalComputed(DNSAccessSpec()),
			"load_balancer": StructOptionalComputed(LoadBalancerAccessSpec()),
		},
	}
}
