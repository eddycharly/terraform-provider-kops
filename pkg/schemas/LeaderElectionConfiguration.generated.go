package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func LeaderElectionConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"leader_elect":                         OptionalBool(),
			"leader_elect_lease_duration":          OptionalDuration(),
			"leader_elect_renew_deadline_duration": OptionalDuration(),
			"leader_elect_resource_lock":           OptionalString(),
			"leader_elect_resource_name":           OptionalString(),
			"leader_elect_resource_namespace":      OptionalString(),
			"leader_elect_retry_period":            OptionalDuration(),
		},
	}
}
