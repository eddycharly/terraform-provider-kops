package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Cluster() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                     RequiredString(),
			"channel":                  OptionalString(),
			"config_base":              OptionalString(),
			"cloud_provider":           RequiredString(),
			"container_runtime":        OptionalString(),
			"kubernetes_version":       OptionalString(),
			"subnet":                   RequiredList(ClusterSubnetSpec()),
			"project":                  OptionalString(),
			"master_public_name":       OptionalString(),
			"master_internal_name":     OptionalString(),
			"network_cidr":             OptionalString(),
			"additional_network_cidrs": OptionalList(String()),
			"network_id":               RequiredString(),
			"topology":                 RequiredStruct(TopologySpec()),
			"secret_store":             OptionalString(),
			"key_store":                OptionalString(),
			"config_store":             OptionalString(),
			"dns_zone":                 OptionalString(),
			"cluster_dns_domain":       OptionalString(),
			"service_cluster_ip_range": OptionalString(),
			"pod_cidr":                 OptionalString(),
			"non_masquerade_cidr":      OptionalString(),
			"ssh_access":               OptionalList(String()),
			"node_port_access":         OptionalList(String()),
			"ssh_key_name":             OptionalString(),
			"kubernetes_api_access":    OptionalList(String()),
			"isolate_masters":          OptionalBool(),
			"update_policy":            OptionalString(),
			"external_policies":        OptionalMap(List(String())),
			"additional_policies":      OptionalMap(String()),
			"etcd_cluster":             RequiredList(EtcdClusterSpec()),
			"networking":               RequiredStruct(NetworkingSpec()),
			"api":                      OptionalStruct(AccessSpec()),
			"cloud_labels":             OptionalMap(String()),
			"encryption_config":        OptionalBool(),
			"disable_subnet_tags":      OptionalBool(),
			"use_host_certificates":    OptionalBool(),
			"sysctl_parameters":        OptionalList(String()),
			"rolling_update":           OptionalStruct(RollingUpdate()),
			"instance_group":           RequiredList(InstanceGroup()),
		},
	}
}
