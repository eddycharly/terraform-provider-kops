package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func KubeSchedulerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"master":                 OptionalString(),
			"log_level":              OptionalInt(),
			"image":                  OptionalString(),
			"leader_election":        OptionalStruct(LeaderElectionConfiguration()),
			"use_policy_config_map":  OptionalBool(),
			"feature_gates":          OptionalMap(String()),
			"max_persistent_volumes": OptionalInt(),
			"qps":                    OptionalQuantity(),
			"burst":                  OptionalInt(),
			"enable_profiling":       OptionalBool(),
		},
	}
}
