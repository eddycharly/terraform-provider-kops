package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CalicoNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu_request":                        OptionalQuantity(),
			"cross_subnet":                       OptionalBool(),
			"log_severity_screen":                OptionalString(),
			"mtu":                                OptionalInt(),
			"prometheus_metrics_enabled":         OptionalBool(),
			"prometheus_metrics_port":            OptionalInt(),
			"prometheus_go_metrics_enabled":      OptionalBool(),
			"prometheus_process_metrics_enabled": OptionalBool(),
			"major_version":                      OptionalString(),
			"iptables_backend":                   OptionalString(),
			"ip_ip_mode":                         OptionalString(),
			"typha_prometheus_metrics_enabled":   OptionalBool(),
			"typha_prometheus_metrics_port":      OptionalInt(),
			"typha_replicas":                     OptionalInt(),
			"wireguard_enabled":                  OptionalBool(),
		},
	}
}
