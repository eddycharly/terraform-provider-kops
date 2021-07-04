package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kopsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kops"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func DataSourceCluster() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                              Required(String()),
			"admin_ssh_key":                     Computed(String()),
			"channel":                           Computed(String()),
			"addons":                            Computed(List(kopsschemas.DataSourceAddonSpec())),
			"config_base":                       Computed(String()),
			"cloud_provider":                    Computed(String()),
			"container_runtime":                 Computed(String()),
			"kubernetes_version":                Computed(String()),
			"subnet":                            Computed(List(kopsschemas.DataSourceClusterSubnetSpec())),
			"project":                           Computed(String()),
			"master_public_name":                Computed(String()),
			"master_internal_name":              Computed(String()),
			"network_cidr":                      Computed(String()),
			"additional_network_cidrs":          Computed(List(String())),
			"network_id":                        Computed(String()),
			"topology":                          Computed(Struct(kopsschemas.DataSourceTopologySpec())),
			"secret_store":                      Computed(String()),
			"key_store":                         Computed(String()),
			"config_store":                      Computed(String()),
			"dns_zone":                          Computed(String()),
			"additional_sans":                   Computed(List(String())),
			"cluster_dns_domain":                Computed(String()),
			"service_cluster_ip_range":          Computed(String()),
			"pod_cidr":                          Computed(String()),
			"non_masquerade_cidr":               Computed(String()),
			"ssh_access":                        Computed(List(String())),
			"node_port_access":                  Computed(List(String())),
			"egress_proxy":                      Computed(Struct(kopsschemas.DataSourceEgressProxySpec())),
			"ssh_key_name":                      Computed(Nullable(String())),
			"kubernetes_api_access":             Computed(List(String())),
			"isolate_masters":                   Computed(Nullable(Bool())),
			"update_policy":                     Computed(Nullable(String())),
			"external_policies":                 Computed(Map(List(String()))),
			"additional_policies":               Computed(Map(String())),
			"file_assets":                       Computed(List(kopsschemas.DataSourceFileAssetSpec())),
			"etcd_cluster":                      Computed(List(kopsschemas.DataSourceEtcdClusterSpec())),
			"containerd":                        Computed(Struct(kopsschemas.DataSourceContainerdConfig())),
			"docker":                            Computed(Struct(kopsschemas.DataSourceDockerConfig())),
			"kube_dns":                          Computed(Struct(kopsschemas.DataSourceKubeDNSConfig())),
			"kube_api_server":                   Computed(Struct(kopsschemas.DataSourceKubeAPIServerConfig())),
			"kube_controller_manager":           Computed(Struct(kopsschemas.DataSourceKubeControllerManagerConfig())),
			"external_cloud_controller_manager": Computed(Struct(kopsschemas.DataSourceCloudControllerManagerConfig())),
			"kube_scheduler":                    Computed(Struct(kopsschemas.DataSourceKubeSchedulerConfig())),
			"kube_proxy":                        Computed(Struct(kopsschemas.DataSourceKubeProxyConfig())),
			"kubelet":                           Computed(Struct(kopsschemas.DataSourceKubeletConfigSpec())),
			"master_kubelet":                    Computed(Struct(kopsschemas.DataSourceKubeletConfigSpec())),
			"cloud_config":                      Computed(Struct(kopsschemas.DataSourceCloudConfiguration())),
			"external_dns":                      Computed(Struct(kopsschemas.DataSourceExternalDNSConfig())),
			"ntp":                               Computed(Struct(kopsschemas.DataSourceNTPConfig())),
			"node_termination_handler":          Computed(Struct(kopsschemas.DataSourceNodeTerminationHandlerConfig())),
			"metrics_server":                    Computed(Struct(kopsschemas.DataSourceMetricsServerConfig())),
			"cert_manager":                      Computed(Struct(kopsschemas.DataSourceCertManagerConfig())),
			"aws_load_balancer_controller":      Computed(Struct(kopsschemas.DataSourceAWSLoadBalancerControllerConfig())),
			"networking":                        Computed(Struct(kopsschemas.DataSourceNetworkingSpec())),
			"api":                               Computed(Struct(kopsschemas.DataSourceAccessSpec())),
			"authentication":                    Computed(Struct(kopsschemas.DataSourceAuthenticationSpec())),
			"authorization":                     Computed(Struct(kopsschemas.DataSourceAuthorizationSpec())),
			"node_authorization":                Computed(Struct(kopsschemas.DataSourceNodeAuthorizationSpec())),
			"cloud_labels":                      Computed(Map(String())),
			"hooks":                             Computed(List(kopsschemas.DataSourceHookSpec())),
			"assets":                            Computed(Struct(kopsschemas.DataSourceAssets())),
			"iam":                               Computed(Struct(kopsschemas.DataSourceIAMSpec())),
			"encryption_config":                 Computed(Nullable(Bool())),
			"disable_subnet_tags":               Computed(Bool()),
			"use_host_certificates":             Computed(Nullable(Bool())),
			"sysctl_parameters":                 Computed(List(String())),
			"rolling_update":                    Computed(Struct(kopsschemas.DataSourceRollingUpdate())),
			"cluster_autoscaler":                Computed(Struct(kopsschemas.DataSourceClusterAutoscalerConfig())),
		},
	}
}

func ExpandDataSourceCluster(in map[string]interface{}) resources.Cluster {
	if in == nil {
		panic("expand Cluster failure, in is nil")
	}
	out := resources.Cluster{}
	if in, ok := in["name"]; ok && in != nil {
		out.Name = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["admin_ssh_key"]; ok && in != nil {
		out.AdminSshKey = func(in interface{}) string { return string(in.(string)) }(in)
	}
	out.ClusterSpec = kopsschemas.ExpandDataSourceClusterSpec(in)
	return out
}

func FlattenDataSourceClusterInto(in resources.Cluster, out map[string]interface{}) {
	out["name"] = func(in string) interface{} { return string(in) }(in.Name)
	out["admin_ssh_key"] = func(in string) interface{} { return string(in) }(in.AdminSshKey)
	kopsschemas.FlattenDataSourceClusterSpecInto(in.ClusterSpec, out)
}

func FlattenDataSourceCluster(in resources.Cluster) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceClusterInto(in, out)
	return out
}
