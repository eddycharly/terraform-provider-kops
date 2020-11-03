package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DockerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authorization_plugins": OptionalList(String()),
			"bridge":                OptionalString(),
			"bridge_ip":             OptionalString(),
			"data_root":             OptionalString(),
			"default_ulimit":        OptionalList(String()),
			"exec_opt":              OptionalList(String()),
			"exec_root":             OptionalString(),
			"experimental":          OptionalBool(),
			"health_check":          OptionalBool(),
			"hosts":                 OptionalList(String()),
			"ip_masq":               OptionalBool(),
			"ip_tables":             OptionalBool(),
			"insecure_registry":     OptionalString(),
			"insecure_registries":   OptionalList(String()),
			"live_restore":          OptionalBool(),
			"log_driver":            OptionalString(),
			"log_level":             OptionalString(),
			"log_opt":               OptionalList(String()),
			"metrics_address":       OptionalString(),
			"mtu":                   OptionalInt(),
			"registry_mirrors":      OptionalList(String()),
			"skip_install":          OptionalBool(),
			"storage":               OptionalString(),
			"storage_opts":          OptionalList(String()),
			"user_namespace_remap":  OptionalString(),
			"version":               OptionalString(),
		},
	}
}
