package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func EtcdClusterSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                    RequiredString(),
			"provider":                OptionalString(),
			"members":                 RequiredList(EtcdMemberSpec()),
			"enable_etcd_tls":         OptionalBool(),
			"enable_tls_auth":         OptionalBool(),
			"version":                 OptionalString(),
			"leader_election_timeout": OptionalDuration(),
			"heartbeat_interval":      OptionalDuration(),
			"image":                   OptionalString(),
			"backups":                 OptionalStruct(EtcdBackupSpec()),
			"manager":                 OptionalStruct(EtcdManagerSpec()),
			"memory_request":          OptionalQuantity(),
			"cpu_request":             OptionalQuantity(),
		},
	}
}
