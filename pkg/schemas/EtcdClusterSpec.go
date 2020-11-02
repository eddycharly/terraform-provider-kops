package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func EtcdClusterSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                    StringEnumRequired("main", "events"),
			"provider":                StringEnumOptionalComputed("Manager", "Legacy"),
			"members":                 ListRequired(EtcdMemberSpec()),
			"enable_etcd_tls":         BoolOptionalComputed(),
			"enable_tls_auth":         BoolOptionalComputed(),
			"version":                 StringOptionalComputed(),
			"leader_election_timeout": DurationOptionalComputed(),
			"heartbeat_interval":      DurationOptionalComputed(),
			"image":                   StringOptionalComputed(),
			"backups":                 StructOptionalComputed(EtcdBackupSpec()),
			"manager":                 StructOptionalComputed(EtcdManagerSpec()),
			"memory_request":          QuantityOptionalComputed(),
			"cpu_request":             QuantityOptionalComputed(),
		},
	}
}
