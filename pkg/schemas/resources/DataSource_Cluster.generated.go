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
			"addons":                            Computed(List(Struct(kopsschemas.DataSourceAddonSpec()))),
			"config_base":                       Computed(String()),
			"cloud_provider":                    Computed(String()),
			"container_runtime":                 Computed(String()),
			"kubernetes_version":                Computed(String()),
			"subnet":                            Computed(List(Struct(kopsschemas.DataSourceClusterSubnetSpec()))),
			"project":                           Computed(String()),
			"master_public_name":                Computed(String()),
			"master_internal_name":              Computed(String()),
			"network_cidr":                      Computed(String()),
			"additional_network_cidrs":          Computed(List(String())),
			"network_id":                        Computed(String()),
			"topology":                          Computed(Ptr(Struct(kopsschemas.DataSourceTopologySpec()))),
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
			"egress_proxy":                      Computed(Ptr(Struct(kopsschemas.DataSourceEgressProxySpec()))),
			"ssh_key_name":                      Computed(Ptr(String())),
			"kubernetes_api_access":             Computed(List(String())),
			"isolate_masters":                   Computed(Ptr(Bool())),
			"update_policy":                     Computed(Ptr(String())),
			"external_policies":                 Computed(Ptr(Map(List(String())))),
			"additional_policies":               Computed(Ptr(Map(String()))),
			"file_assets":                       Computed(List(Struct(kopsschemas.DataSourceFileAssetSpec()))),
			"etcd_cluster":                      Computed(List(Struct(kopsschemas.DataSourceEtcdClusterSpec()))),
			"containerd":                        Computed(Ptr(Struct(kopsschemas.DataSourceContainerdConfig()))),
			"docker":                            Computed(Ptr(Struct(kopsschemas.DataSourceDockerConfig()))),
			"kube_dns":                          Computed(Ptr(Struct(kopsschemas.DataSourceKubeDNSConfig()))),
			"kube_api_server":                   Computed(Ptr(Struct(kopsschemas.DataSourceKubeAPIServerConfig()))),
			"kube_controller_manager":           Computed(Ptr(Struct(kopsschemas.DataSourceKubeControllerManagerConfig()))),
			"external_cloud_controller_manager": Computed(Ptr(Struct(kopsschemas.DataSourceCloudControllerManagerConfig()))),
			"kube_scheduler":                    Computed(Ptr(Struct(kopsschemas.DataSourceKubeSchedulerConfig()))),
			"kube_proxy":                        Computed(Ptr(Struct(kopsschemas.DataSourceKubeProxyConfig()))),
			"kubelet":                           Computed(Ptr(Struct(kopsschemas.DataSourceKubeletConfigSpec()))),
			"master_kubelet":                    Computed(Ptr(Struct(kopsschemas.DataSourceKubeletConfigSpec()))),
			"cloud_config":                      Computed(Ptr(Struct(kopsschemas.DataSourceCloudConfiguration()))),
			"external_dns":                      Computed(Ptr(Struct(kopsschemas.DataSourceExternalDNSConfig()))),
			"ntp":                               Computed(Ptr(Struct(kopsschemas.DataSourceNTPConfig()))),
			"node_termination_handler":          Computed(Ptr(Struct(kopsschemas.DataSourceNodeTerminationHandlerConfig()))),
			"metrics_server":                    Computed(Ptr(Struct(kopsschemas.DataSourceMetricsServerConfig()))),
			"cert_manager":                      Computed(Ptr(Struct(kopsschemas.DataSourceCertManagerConfig()))),
			"aws_load_balancer_controller":      Computed(Ptr(Struct(kopsschemas.DataSourceAWSLoadBalancerControllerConfig()))),
			"networking":                        Computed(Ptr(Struct(kopsschemas.DataSourceNetworkingSpec()))),
			"api":                               Computed(Ptr(Struct(kopsschemas.DataSourceAccessSpec()))),
			"authentication":                    Computed(Ptr(Struct(kopsschemas.DataSourceAuthenticationSpec()))),
			"authorization":                     Computed(Ptr(Struct(kopsschemas.DataSourceAuthorizationSpec()))),
			"node_authorization":                Computed(Ptr(Struct(kopsschemas.DataSourceNodeAuthorizationSpec()))),
			"cloud_labels":                      Computed(Map(String())),
			"hooks":                             Computed(List(Struct(kopsschemas.DataSourceHookSpec()))),
			"assets":                            Computed(Ptr(Struct(kopsschemas.DataSourceAssets()))),
			"iam":                               Computed(Ptr(Struct(kopsschemas.DataSourceIAMSpec()))),
			"encryption_config":                 Computed(Ptr(Bool())),
			"disable_subnet_tags":               Computed(Bool()),
			"use_host_certificates":             Computed(Ptr(Bool())),
			"sysctl_parameters":                 Computed(List(String())),
			"rolling_update":                    Computed(Ptr(Struct(kopsschemas.DataSourceRollingUpdate()))),
			"cluster_autoscaler":                Computed(Ptr(Struct(kopsschemas.DataSourceClusterAutoscalerConfig()))),
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
