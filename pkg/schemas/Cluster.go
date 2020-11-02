package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Cluster() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                  StringRequired(),
			"network_id":            StringRequired(),
			"cloud_provider":        StringEnumRequired("aws"), // TODO add other providers
			"kubernetes_version":    StringRequired(),
			"project":               StringOptional(),
			"master_public_name":    StringRequired(),
			"master_internal_name":  StringRequired(),
			"ssh_key_name":          StringOptionalComputed(),
			"channel":               StringOptionalComputed(),
			"config_base":           StringOptionalComputed(),
			"config_store":          StringOptionalComputed(),
			"secret_store":          StringOptionalComputed(),
			"key_store":             StringOptionalComputed(),
			"dns_zone":              StringOptionalComputed(),
			"ssh_access":            ListOptional(String()),
			"kubernetes_api_access": ListOptional(String()),
			"api":                   StructOptionalComputed(AccessSpec()),
			"networking":            StructOptionalComputed(NetworkingSpec()),
			"topology":              StructRequired(TopologySpec()),
			"subnet":                ListRequired(ClusterSubnetSpec()),
			"etcd_cluster":          ListRequired(EtcdClusterSpec()),
			"instance_group":        ListRequired(InstanceGroup()),
		},
	}
}
