package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ClusterSubnetSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":        RequiredString(),
			"cidr":        OptionalComputedString(),
			"zone":        RequiredString(),
			"region":      OptionalString(),
			"provider_id": RequiredString(),
			"egress":      OptionalString(),
			"type":        RequiredString(),
			"public_ip":   OptionalString(),
		},
	}
}
