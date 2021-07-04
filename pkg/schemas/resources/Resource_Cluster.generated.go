package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kopsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kops"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ResourceCluster() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"revision":                          Computed(Int()),
			"name":                              ForceNew(Required(String())),
			"admin_ssh_key":                     Sensitive(Required(String())),
			"channel":                           Optional(String()),
			"addons":                            Optional(List(kopsschemas.ResourceAddonSpec())),
			"config_base":                       Optional(Computed(String())),
			"cloud_provider":                    Required(String()),
			"container_runtime":                 Optional(String()),
			"kubernetes_version":                Optional(String()),
			"subnet":                            Required(List(kopsschemas.ResourceClusterSubnetSpec())),
			"project":                           Optional(String()),
			"master_public_name":                Optional(Computed(String())),
			"master_internal_name":              Optional(Computed(String())),
			"network_cidr":                      Optional(Computed(String())),
			"additional_network_cidrs":          Optional(List(String())),
			"network_id":                        Required(String()),
			"topology":                          Required(Struct(kopsschemas.ResourceTopologySpec())),
			"secret_store":                      Optional(String()),
			"key_store":                         Optional(String()),
			"config_store":                      Optional(String()),
			"dns_zone":                          Optional(String()),
			"additional_sans":                   Optional(List(String())),
			"cluster_dns_domain":                Optional(String()),
			"service_cluster_ip_range":          Optional(String()),
			"pod_cidr":                          Optional(String()),
			"non_masquerade_cidr":               Optional(Computed(String())),
			"ssh_access":                        Optional(List(String())),
			"node_port_access":                  Optional(List(String())),
			"egress_proxy":                      Optional(Struct(kopsschemas.ResourceEgressProxySpec())),
			"ssh_key_name":                      Optional(Nullable(String())),
			"kubernetes_api_access":             Optional(List(String())),
			"isolate_masters":                   Optional(Nullable(Bool())),
			"update_policy":                     Optional(Nullable(String())),
			"external_policies":                 Optional(Map(List(String()))),
			"additional_policies":               Optional(Map(String())),
			"file_assets":                       Optional(List(kopsschemas.ResourceFileAssetSpec())),
			"etcd_cluster":                      Required(List(kopsschemas.ResourceEtcdClusterSpec())),
			"containerd":                        Optional(Struct(kopsschemas.ResourceContainerdConfig())),
			"docker":                            Optional(Struct(kopsschemas.ResourceDockerConfig())),
			"kube_dns":                          Optional(Struct(kopsschemas.ResourceKubeDNSConfig())),
			"kube_api_server":                   Optional(Struct(kopsschemas.ResourceKubeAPIServerConfig())),
			"kube_controller_manager":           Optional(Struct(kopsschemas.ResourceKubeControllerManagerConfig())),
			"external_cloud_controller_manager": Optional(Struct(kopsschemas.ResourceCloudControllerManagerConfig())),
			"kube_scheduler":                    Optional(Struct(kopsschemas.ResourceKubeSchedulerConfig())),
			"kube_proxy":                        Optional(Struct(kopsschemas.ResourceKubeProxyConfig())),
			"kubelet":                           Optional(Struct(kopsschemas.ResourceKubeletConfigSpec())),
			"master_kubelet":                    Optional(Struct(kopsschemas.ResourceKubeletConfigSpec())),
			"cloud_config":                      Optional(Struct(kopsschemas.ResourceCloudConfiguration())),
			"external_dns":                      Optional(Struct(kopsschemas.ResourceExternalDNSConfig())),
			"ntp":                               Optional(Struct(kopsschemas.ResourceNTPConfig())),
			"node_termination_handler":          Optional(Struct(kopsschemas.ResourceNodeTerminationHandlerConfig())),
			"metrics_server":                    Optional(Struct(kopsschemas.ResourceMetricsServerConfig())),
			"cert_manager":                      Optional(Struct(kopsschemas.ResourceCertManagerConfig())),
			"aws_load_balancer_controller":      Optional(Struct(kopsschemas.ResourceAWSLoadBalancerControllerConfig())),
			"networking":                        Required(Struct(kopsschemas.ResourceNetworkingSpec())),
			"api":                               Optional(Struct(kopsschemas.ResourceAccessSpec())),
			"authentication":                    Optional(Struct(kopsschemas.ResourceAuthenticationSpec())),
			"authorization":                     Optional(Struct(kopsschemas.ResourceAuthorizationSpec())),
			"node_authorization":                Optional(Struct(kopsschemas.ResourceNodeAuthorizationSpec())),
			"cloud_labels":                      Optional(Map(String())),
			"hooks":                             Optional(List(kopsschemas.ResourceHookSpec())),
			"assets":                            Optional(Struct(kopsschemas.ResourceAssets())),
			"iam":                               Optional(Computed(Struct(kopsschemas.ResourceIAMSpec()))),
			"encryption_config":                 Optional(Nullable(Bool())),
			"disable_subnet_tags":               Optional(Bool()),
			"use_host_certificates":             Optional(Nullable(Bool())),
			"sysctl_parameters":                 Optional(List(String())),
			"rolling_update":                    Optional(Struct(kopsschemas.ResourceRollingUpdate())),
			"cluster_autoscaler":                Optional(Struct(kopsschemas.ResourceClusterAutoscalerConfig())),
		},
	}
}

func ExpandResourceCluster(in map[string]interface{}) resources.Cluster {
	if in == nil {
		panic("expand Cluster failure, in is nil")
	}
	out := resources.Cluster{}
	if in, ok := in["revision"]; ok && in != nil {
		out.Revision = func(in interface{}) int { return int(in.(int)) }(in)
	}
	if in, ok := in["name"]; ok && in != nil {
		out.Name = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["admin_ssh_key"]; ok && in != nil {
		out.AdminSshKey = func(in interface{}) string { return string(in.(string)) }(in)
	}
	out.ClusterSpec = kopsschemas.ExpandResourceClusterSpec(in)
	return out
}

func FlattenResourceClusterInto(in resources.Cluster, out map[string]interface{}) {
	out["revision"] = func(in int) interface{} { return int(in) }(in.Revision)
	out["name"] = func(in string) interface{} { return string(in) }(in.Name)
	out["admin_ssh_key"] = func(in string) interface{} { return string(in) }(in.AdminSshKey)
	kopsschemas.FlattenResourceClusterSpecInto(in.ClusterSpec, out)
}

func FlattenResourceCluster(in resources.Cluster) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceClusterInto(in, out)
	return out
}
