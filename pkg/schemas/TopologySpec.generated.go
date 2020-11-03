package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TopologySpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"masters": RequiredString(),
			"nodes":   RequiredString(),
			"bastion": OptionalStruct(BastionSpec()),
			"dns":     RequiredStruct(DNSSpec()),
		},
	}
}
