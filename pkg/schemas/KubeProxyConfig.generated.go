package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func KubeProxyConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image":                  OptionalString(),
			"cpu_request":            OptionalString(),
			"cpu_limit":              OptionalString(),
			"memory_request":         OptionalString(),
			"memory_limit":           OptionalString(),
			"log_level":              OptionalInt(),
			"cluster_cidr":           OptionalString(),
			"hostname_override":      OptionalString(),
			"bind_address":           OptionalString(),
			"master":                 OptionalString(),
			"metrics_bind_address":   OptionalString(),
			"enabled":                OptionalBool(),
			"proxy_mode":             OptionalString(),
			"ip_vs_exclude_cidr_s":   OptionalList(String()),
			"ip_vs_min_sync_period":  OptionalDuration(),
			"ip_vs_scheduler":        OptionalString(),
			"ip_vs_sync_period":      OptionalDuration(),
			"feature_gates":          OptionalMap(String()),
			"conntrack_max_per_core": OptionalInt(),
			"conntrack_min":          OptionalInt(),
		},
	}
}
