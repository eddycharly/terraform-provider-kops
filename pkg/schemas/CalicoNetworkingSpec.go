package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CalicoNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu_request":                        QuantityOptionalComputed(),
			"cross_subnet":                       BoolOptionalComputed(),
			"log_severity_screen":                StringOptionalComputed(),
			"mtu":                                IntOptionalComputed(),
			"prometheus_metrics_enabled":         BoolOptionalComputed(),
			"prometheus_metrics_port":            IntOptionalComputed(),
			"prometheus_go_metrics_enabled":      BoolOptionalComputed(),
			"prometheus_process_metrics_enabled": BoolOptionalComputed(),
			"major_version":                      StringOptionalComputed(),
			"iptables_backend":                   StringOptionalComputed(),
			"ip_ip_mode":                         StringOptional(),
			"typha_prometheus_metrics_enabled":   BoolOptionalComputed(),
			"typha_prometheus_metrics_port":      IntOptionalComputed(),
			"typha_replicas":                     IntOptionalComputed(),
			"wireguard_enabled":                  BoolOptionalComputed(),
		},
	}
}
