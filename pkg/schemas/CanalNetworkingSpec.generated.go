package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CanalNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"chain_insert_mode":                  OptionalString(),
			"cpu_request":                        OptionalQuantity(),
			"default_endpoint_to_host_action":    OptionalString(),
			"disable_flannel_forward_rules":      OptionalBool(),
			"disable_tx_checksum_offloading":     OptionalBool(),
			"iptables_backend":                   OptionalString(),
			"log_severity_sys":                   OptionalString(),
			"mtu":                                OptionalInt(),
			"prometheus_go_metrics_enabled":      OptionalBool(),
			"prometheus_metrics_enabled":         OptionalBool(),
			"prometheus_metrics_port":            OptionalInt(),
			"prometheus_process_metrics_enabled": OptionalBool(),
			"typha_prometheus_metrics_enabled":   OptionalBool(),
			"typha_prometheus_metrics_port":      OptionalInt(),
			"typha_replicas":                     OptionalInt(),
		},
	}
}
