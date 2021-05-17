package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kopsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kops"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceCluster() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"revision":                          ComputedInt(),
			"name":                              ForceNew(RequiredString()),
			"admin_ssh_key":                     Sensitive(RequiredString()),
			"channel":                           OptionalString(),
			"addons":                            OptionalList(kopsschemas.ResourceAddonSpec()),
			"config_base":                       OptionalComputedString(),
			"cloud_provider":                    RequiredString(),
			"container_runtime":                 OptionalString(),
			"kubernetes_version":                OptionalString(),
			"subnet":                            RequiredList(kopsschemas.ResourceClusterSubnetSpec()),
			"project":                           OptionalString(),
			"master_public_name":                OptionalComputedString(),
			"master_internal_name":              OptionalComputedString(),
			"network_cidr":                      OptionalComputedString(),
			"additional_network_cidrs":          OptionalList(String()),
			"network_id":                        RequiredString(),
			"topology":                          RequiredStruct(kopsschemas.ResourceTopologySpec()),
			"secret_store":                      OptionalString(),
			"key_store":                         OptionalString(),
			"config_store":                      OptionalString(),
			"dns_zone":                          OptionalString(),
			"additional_sans":                   OptionalList(String()),
			"cluster_dns_domain":                OptionalString(),
			"service_cluster_ip_range":          OptionalString(),
			"pod_cidr":                          OptionalString(),
			"non_masquerade_cidr":               OptionalComputedString(),
			"ssh_access":                        OptionalList(String()),
			"node_port_access":                  OptionalList(String()),
			"egress_proxy":                      OptionalStruct(kopsschemas.ResourceEgressProxySpec()),
			"ssh_key_name":                      OptionalString(),
			"kubernetes_api_access":             OptionalList(String()),
			"isolate_masters":                   OptionalBool(),
			"update_policy":                     OptionalString(),
			"external_policies":                 OptionalMap(List(String())),
			"additional_policies":               OptionalMap(String()),
			"file_assets":                       OptionalList(kopsschemas.ResourceFileAssetSpec()),
			"etcd_cluster":                      RequiredList(kopsschemas.ResourceEtcdClusterSpec()),
			"containerd":                        OptionalStruct(kopsschemas.ResourceContainerdConfig()),
			"docker":                            OptionalStruct(kopsschemas.ResourceDockerConfig()),
			"kube_dns":                          OptionalStruct(kopsschemas.ResourceKubeDNSConfig()),
			"kube_api_server":                   OptionalStruct(kopsschemas.ResourceKubeAPIServerConfig()),
			"kube_controller_manager":           OptionalStruct(kopsschemas.ResourceKubeControllerManagerConfig()),
			"external_cloud_controller_manager": OptionalStruct(kopsschemas.ResourceCloudControllerManagerConfig()),
			"kube_scheduler":                    OptionalStruct(kopsschemas.ResourceKubeSchedulerConfig()),
			"kube_proxy":                        OptionalStruct(kopsschemas.ResourceKubeProxyConfig()),
			"kubelet":                           OptionalStruct(kopsschemas.ResourceKubeletConfigSpec()),
			"master_kubelet":                    OptionalStruct(kopsschemas.ResourceKubeletConfigSpec()),
			"cloud_config":                      OptionalStruct(kopsschemas.ResourceCloudConfiguration()),
			"external_dns":                      OptionalStruct(kopsschemas.ResourceExternalDNSConfig()),
			"ntp":                               OptionalStruct(kopsschemas.ResourceNTPConfig()),
			"node_termination_handler":          OptionalStruct(kopsschemas.ResourceNodeTerminationHandlerConfig()),
			"metrics_server":                    OptionalStruct(kopsschemas.ResourceMetricsServerConfig()),
			"cert_manager":                      OptionalStruct(kopsschemas.ResourceCertManagerConfig()),
			"aws_load_balancer_controller":      OptionalStruct(kopsschemas.ResourceAWSLoadBalancerControllerConfig()),
			"networking":                        RequiredStruct(kopsschemas.ResourceNetworkingSpec()),
			"api":                               OptionalStruct(kopsschemas.ResourceAccessSpec()),
			"authentication":                    OptionalStruct(kopsschemas.ResourceAuthenticationSpec()),
			"authorization":                     OptionalStruct(kopsschemas.ResourceAuthorizationSpec()),
			"node_authorization":                OptionalStruct(kopsschemas.ResourceNodeAuthorizationSpec()),
			"cloud_labels":                      OptionalMap(String()),
			"hooks":                             OptionalList(kopsschemas.ResourceHookSpec()),
			"assets":                            OptionalStruct(kopsschemas.ResourceAssets()),
			"iam":                               OptionalComputedStruct(kopsschemas.ResourceIAMSpec()),
			"encryption_config":                 OptionalBool(),
			"disable_subnet_tags":               OptionalBool(),
			"use_host_certificates":             OptionalBool(),
			"sysctl_parameters":                 OptionalList(String()),
			"rolling_update":                    OptionalStruct(kopsschemas.ResourceRollingUpdate()),
			"cluster_autoscaler":                OptionalStruct(kopsschemas.ResourceClusterAutoscalerConfig()),
		},
	}
}

func ExpandResourceCluster(in map[string]interface{}) resources.Cluster {
	if in == nil {
		panic("expand Cluster failure, in is nil")
	}
	return resources.Cluster{
		Revision: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["revision"]),
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		AdminSshKey: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["admin_ssh_key"]),
		ClusterSpec: func(in interface{}) kops.ClusterSpec {
			return kopsschemas.ExpandResourceClusterSpec(in.(map[string]interface{}))
		}(in),
	}
}

func FlattenResourceClusterInto(in resources.Cluster, out map[string]interface{}) {
	out["revision"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.Revision)
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["admin_ssh_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AdminSshKey)
	kopsschemas.FlattenResourceClusterSpecInto(in.ClusterSpec, out)
}

func FlattenResourceCluster(in resources.Cluster) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceClusterInto(in, out)
	return out
}
