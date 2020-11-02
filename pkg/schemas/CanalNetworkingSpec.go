package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CanalNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"chain_insert_mode":                  StringOptionalComputed(),
			"cpu_request":                        QuantityOptionalComputed(),
			"default_endpoint_to_host_action":    StringOptionalComputed(),
			"disable_flannel_forward_rules":      BoolOptionalComputed(),
			"disable_tx_checksum_offloading":     BoolOptionalComputed(),
			"iptables_backend":                   StringOptionalComputed(),
			"log_severity_sys":                   StringOptionalComputed(),
			"mtu":                                IntOptionalComputed(),
			"prometheus_go_metrics_enabled":      BoolOptionalComputed(),
			"prometheus_metrics_enabled":         BoolOptionalComputed(),
			"prometheus_metrics_port":            IntOptionalComputed(),
			"prometheus_process_metrics_enabled": BoolOptionalComputed(),
			"typha_prometheus_metrics_enabled":   BoolOptionalComputed(),
			"typha_prometheus_metrics_port":      IntOptionalComputed(),
			"typha_replicas":                     IntOptionalComputed(),
		},
	}
}
