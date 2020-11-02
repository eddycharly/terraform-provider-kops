package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TopologySpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"masters": StringEnumRequired("public", "private"),
			"nodes":   StringEnumRequired("public", "private"),
			"bastion": StructOptionalComputed(BastionSpec()),
			"dns":     StructOptionalComputed(DNSSpec()),
		},
	}
}
