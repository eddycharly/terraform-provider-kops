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
			"addons":                            Optional(List(Struct(kopsschemas.ResourceAddonSpec()))),
			"config_base":                       Optional(Computed(String())),
			"cloud_provider":                    Required(String()),
			"container_runtime":                 Optional(String()),
			"kubernetes_version":                Optional(String()),
			"subnet":                            Required(List(Struct(kopsschemas.ResourceClusterSubnetSpec()))),
			"project":                           Optional(String()),
			"master_public_name":                Optional(Computed(String())),
			"master_internal_name":              Optional(Computed(String())),
			"network_cidr":                      Optional(Computed(String())),
			"additional_network_cidrs":          Optional(List(String())),
			"network_id":                        Required(String()),
			"topology":                          Required(Ptr(Struct(kopsschemas.ResourceTopologySpec()))),
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
			"egress_proxy":                      Optional(Ptr(Struct(kopsschemas.ResourceEgressProxySpec()))),
			"ssh_key_name":                      Optional(Ptr(String())),
			"kubernetes_api_access":             Optional(List(String())),
			"isolate_masters":                   Optional(Ptr(Bool())),
			"update_policy":                     Optional(Ptr(String())),
			"external_policies":                 Optional(Ptr(Map(List(String())))),
			"additional_policies":               Optional(Ptr(Map(String()))),
			"file_assets":                       Optional(List(Struct(kopsschemas.ResourceFileAssetSpec()))),
			"etcd_cluster":                      Required(List(Struct(kopsschemas.ResourceEtcdClusterSpec()))),
			"containerd":                        Optional(Ptr(Struct(kopsschemas.ResourceContainerdConfig()))),
			"docker":                            Optional(Ptr(Struct(kopsschemas.ResourceDockerConfig()))),
			"kube_dns":                          Optional(Ptr(Struct(kopsschemas.ResourceKubeDNSConfig()))),
			"kube_api_server":                   Optional(Ptr(Struct(kopsschemas.ResourceKubeAPIServerConfig()))),
			"kube_controller_manager":           Optional(Ptr(Struct(kopsschemas.ResourceKubeControllerManagerConfig()))),
			"external_cloud_controller_manager": Optional(Ptr(Struct(kopsschemas.ResourceCloudControllerManagerConfig()))),
			"kube_scheduler":                    Optional(Ptr(Struct(kopsschemas.ResourceKubeSchedulerConfig()))),
			"kube_proxy":                        Optional(Ptr(Struct(kopsschemas.ResourceKubeProxyConfig()))),
			"kubelet":                           Optional(Ptr(Struct(kopsschemas.ResourceKubeletConfigSpec()))),
			"master_kubelet":                    Optional(Ptr(Struct(kopsschemas.ResourceKubeletConfigSpec()))),
			"cloud_config":                      Optional(Ptr(Struct(kopsschemas.ResourceCloudConfiguration()))),
			"external_dns":                      Optional(Ptr(Struct(kopsschemas.ResourceExternalDNSConfig()))),
			"ntp":                               Optional(Ptr(Struct(kopsschemas.ResourceNTPConfig()))),
			"node_termination_handler":          Optional(Ptr(Struct(kopsschemas.ResourceNodeTerminationHandlerConfig()))),
			"metrics_server":                    Optional(Ptr(Struct(kopsschemas.ResourceMetricsServerConfig()))),
			"cert_manager":                      Optional(Ptr(Struct(kopsschemas.ResourceCertManagerConfig()))),
			"aws_load_balancer_controller":      Optional(Ptr(Struct(kopsschemas.ResourceAWSLoadBalancerControllerConfig()))),
			"networking":                        Required(Ptr(Struct(kopsschemas.ResourceNetworkingSpec()))),
			"api":                               Optional(Ptr(Struct(kopsschemas.ResourceAccessSpec()))),
			"authentication":                    Optional(Ptr(Struct(kopsschemas.ResourceAuthenticationSpec()))),
			"authorization":                     Optional(Ptr(Struct(kopsschemas.ResourceAuthorizationSpec()))),
			"node_authorization":                Optional(Ptr(Struct(kopsschemas.ResourceNodeAuthorizationSpec()))),
			"cloud_labels":                      Optional(Map(String())),
			"hooks":                             Optional(List(Struct(kopsschemas.ResourceHookSpec()))),
			"assets":                            Optional(Ptr(Struct(kopsschemas.ResourceAssets()))),
			"iam":                               Optional(Computed(Ptr(Struct(kopsschemas.ResourceIAMSpec())))),
			"encryption_config":                 Optional(Ptr(Bool())),
			"disable_subnet_tags":               Optional(Bool()),
			"use_host_certificates":             Optional(Ptr(Bool())),
			"sysctl_parameters":                 Optional(List(String())),
			"rolling_update":                    Optional(Ptr(Struct(kopsschemas.ResourceRollingUpdate()))),
			"cluster_autoscaler":                Optional(Ptr(Struct(kopsschemas.ResourceClusterAutoscalerConfig()))),
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
