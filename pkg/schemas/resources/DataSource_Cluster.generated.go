package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kopsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kops"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceCluster() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                              RequiredString(),
			"admin_ssh_key":                     ComputedString(),
			"channel":                           ComputedString(),
			"addons":                            ComputedList(kopsschemas.DataSourceAddonSpec()),
			"config_base":                       ComputedString(),
			"cloud_provider":                    ComputedString(),
			"container_runtime":                 ComputedString(),
			"kubernetes_version":                ComputedString(),
			"subnet":                            ComputedList(kopsschemas.DataSourceClusterSubnetSpec()),
			"project":                           ComputedString(),
			"master_public_name":                ComputedString(),
			"master_internal_name":              ComputedString(),
			"network_cidr":                      ComputedString(),
			"additional_network_cidrs":          ComputedList(String()),
			"network_id":                        ComputedString(),
			"topology":                          ComputedStruct(kopsschemas.DataSourceTopologySpec()),
			"secret_store":                      ComputedString(),
			"key_store":                         ComputedString(),
			"config_store":                      ComputedString(),
			"dns_zone":                          ComputedString(),
			"additional_sans":                   ComputedList(String()),
			"cluster_dns_domain":                ComputedString(),
			"service_cluster_ip_range":          ComputedString(),
			"pod_cidr":                          ComputedString(),
			"non_masquerade_cidr":               ComputedString(),
			"ssh_access":                        ComputedList(String()),
			"node_port_access":                  ComputedList(String()),
			"egress_proxy":                      ComputedStruct(kopsschemas.DataSourceEgressProxySpec()),
			"ssh_key_name":                      ComputedString(),
			"kubernetes_api_access":             ComputedList(String()),
			"isolate_masters":                   ComputedBool(),
			"update_policy":                     ComputedString(),
			"external_policies":                 ComputedMap(List(String())),
			"additional_policies":               ComputedMap(String()),
			"file_assets":                       ComputedList(kopsschemas.DataSourceFileAssetSpec()),
			"etcd_cluster":                      ComputedList(kopsschemas.DataSourceEtcdClusterSpec()),
			"containerd":                        ComputedStruct(kopsschemas.DataSourceContainerdConfig()),
			"docker":                            ComputedStruct(kopsschemas.DataSourceDockerConfig()),
			"kube_dns":                          ComputedStruct(kopsschemas.DataSourceKubeDNSConfig()),
			"kube_api_server":                   ComputedStruct(kopsschemas.DataSourceKubeAPIServerConfig()),
			"kube_controller_manager":           ComputedStruct(kopsschemas.DataSourceKubeControllerManagerConfig()),
			"external_cloud_controller_manager": ComputedStruct(kopsschemas.DataSourceCloudControllerManagerConfig()),
			"kube_scheduler":                    ComputedStruct(kopsschemas.DataSourceKubeSchedulerConfig()),
			"kube_proxy":                        ComputedStruct(kopsschemas.DataSourceKubeProxyConfig()),
			"kubelet":                           ComputedStruct(kopsschemas.DataSourceKubeletConfigSpec()),
			"master_kubelet":                    ComputedStruct(kopsschemas.DataSourceKubeletConfigSpec()),
			"cloud_config":                      ComputedStruct(kopsschemas.DataSourceCloudConfiguration()),
			"external_dns":                      ComputedStruct(kopsschemas.DataSourceExternalDNSConfig()),
			"ntp":                               ComputedStruct(kopsschemas.DataSourceNTPConfig()),
			"node_termination_handler":          ComputedStruct(kopsschemas.DataSourceNodeTerminationHandlerConfig()),
			"metrics_server":                    ComputedStruct(kopsschemas.DataSourceMetricsServerConfig()),
			"cert_manager":                      ComputedStruct(kopsschemas.DataSourceCertManagerConfig()),
			"aws_load_balancer_controller":      ComputedStruct(kopsschemas.DataSourceAWSLoadBalancerControllerConfig()),
			"networking":                        ComputedStruct(kopsschemas.DataSourceNetworkingSpec()),
			"api":                               ComputedStruct(kopsschemas.DataSourceAccessSpec()),
			"authentication":                    ComputedStruct(kopsschemas.DataSourceAuthenticationSpec()),
			"authorization":                     ComputedStruct(kopsschemas.DataSourceAuthorizationSpec()),
			"node_authorization":                ComputedStruct(kopsschemas.DataSourceNodeAuthorizationSpec()),
			"cloud_labels":                      ComputedMap(String()),
			"hooks":                             ComputedList(kopsschemas.DataSourceHookSpec()),
			"assets":                            ComputedStruct(kopsschemas.DataSourceAssets()),
			"iam":                               ComputedStruct(kopsschemas.DataSourceIAMSpec()),
			"encryption_config":                 ComputedBool(),
			"disable_subnet_tags":               ComputedBool(),
			"use_host_certificates":             ComputedBool(),
			"sysctl_parameters":                 ComputedList(String()),
			"rolling_update":                    ComputedStruct(kopsschemas.DataSourceRollingUpdate()),
			"cluster_autoscaler":                ComputedStruct(kopsschemas.DataSourceClusterAutoscalerConfig()),
		},
	}
}

func ExpandDataSourceCluster(in map[string]interface{}) resources.Cluster {
	if in == nil {
		panic("expand Cluster failure, in is nil")
	}
	return resources.Cluster{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		AdminSshKey: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["admin_ssh_key"]),
		ClusterSpec: func(in interface{}) kops.ClusterSpec {
			return kopsschemas.ExpandDataSourceClusterSpec(in.(map[string]interface{}))
		}(in),
	}
}

func FlattenDataSourceClusterInto(in resources.Cluster, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["admin_ssh_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AdminSshKey)
	kopsschemas.FlattenDataSourceClusterSpecInto(in.ClusterSpec, out)
}

func FlattenDataSourceCluster(in resources.Cluster) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceClusterInto(in, out)
	return out
}
