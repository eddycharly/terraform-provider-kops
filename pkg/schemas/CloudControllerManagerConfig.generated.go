package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CloudControllerManagerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"master":                          OptionalString(),
			"log_level":                       OptionalInt(),
			"image":                           OptionalString(),
			"cloud_provider":                  OptionalString(),
			"cluster_name":                    OptionalString(),
			"cluster_cidr":                    OptionalString(),
			"allocate_node_cidrs":             OptionalBool(),
			"configure_cloud_routes":          OptionalBool(),
			"cidr_allocator_type":             OptionalString(),
			"leader_election":                 OptionalStruct(LeaderElectionConfiguration()),
			"use_service_account_credentials": OptionalBool(),
		},
	}
}
