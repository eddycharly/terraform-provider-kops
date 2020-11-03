package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ExternalDNSConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable":         OptionalBool(),
			"watch_ingress":   OptionalBool(),
			"watch_namespace": OptionalString(),
		},
	}
}
