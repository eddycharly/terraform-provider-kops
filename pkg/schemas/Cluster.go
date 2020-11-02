package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Cluster() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                     StringRequired(),
			"channel":                  StringOptionalComputed(),
			"config_base":              StringOptionalComputed(),
			"cloud_provider":           StringEnumRequired("aws"), // TODO add other providers
			"container_runtime":        StringOptionalComputed(),
			"kubernetes_version":       StringRequired(),
			"subnet":                   ListRequired(ClusterSubnetSpec()),
			"project":                  StringOptional(),
			"master_public_name":       StringRequired(),
			"master_internal_name":     StringRequired(),
			"network_cidr":             StringOptionalComputed(),
			"additional_network_cidrs": ListOptional(String()),
			"network_id":               StringRequired(),
			"topology":                 StructRequired(TopologySpec()),
			"secret_store":             StringOptionalComputed(),
			"key_store":                StringOptionalComputed(),
			"config_store":             StringOptionalComputed(),
			"dns_zone":                 StringOptionalComputed(),
			"cluster_dns_domain":       StringOptionalComputed(),
			"service_cluster_ip_range": StringOptionalComputed(),
			"pod_cidr":                 StringOptionalComputed(),
			"non_masquerade_cidr":      StringOptionalComputed(),
			"ssh_access":               ListOptional(String()),
			"node_port_access":         ListOptional(String()),
			"ssh_key_name":             StringOptionalComputed(),
			"kubernetes_api_access":    ListOptional(String()),
			"isolate_masters":          BoolOptionalComputed(),
			"update_policy":            StringOptionalComputed(),
			"external_policies":        MapOptionalComputed(ListOptional(String())),
			"additional_policies":      MapOptionalComputed(String()),
			"etcd_cluster":             ListRequired(EtcdClusterSpec()),
			"networking":               StructOptionalComputed(NetworkingSpec()),
			"api":                      StructOptionalComputed(AccessSpec()),
			"cloud_labels":             MapOptionalComputed(String()),
			"encryption_config":        BoolOptionalComputed(),
			"disable_subnet_tags":      BoolOptionalComputed(),
			"use_host_certificates":    BoolOptionalComputed(),
			"sysctl_parameters":        ListOptional(String()),
			"rolling_update":           StructOptionalComputed(RollingUpdate()),
			"instance_group":           ListRequired(InstanceGroup()),
		},
	}
}
