package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func OpenstackConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"loadbalancer":         OptionalStruct(OpenstackLoadbalancerConfig()),
			"monitor":              OptionalStruct(OpenstackMonitor()),
			"router":               OptionalStruct(OpenstackRouter()),
			"block_storage":        OptionalStruct(OpenstackBlockStorageConfig()),
			"insecure_skip_verify": OptionalBool(),
		},
	}
}
