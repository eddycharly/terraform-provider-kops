package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func RomanaNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"daemon_service_ip": OptionalString(),
			"etcd_service_ip":   OptionalString(),
		},
	}
}
