package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ClusterSubnetSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":        StringRequired(),
			"cidr":        StringRequired(),
			"zone":        StringOptionalComputed(),
			"region":      StringOptionalComputed(),
			"provider_id": StringOptionalComputed(),
			"egress":      StringOptionalComputed(),
			"type":        StringEnumRequired("Public", "Private", "Utility"),
			"public_ip":   StringOptionalComputed(),
		},
	}
}
