package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceClusterSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"channel":                           Computed(String()),
			"addons":                            Computed(List(Struct(DataSourceAddonSpec()))),
			"config_base":                       Computed(String()),
			"cloud_provider":                    Computed(String()),
			"container_runtime":                 Computed(String()),
			"kubernetes_version":                Computed(String()),
			"subnet":                            Computed(List(Struct(DataSourceClusterSubnetSpec()))),
			"project":                           Computed(String()),
			"master_public_name":                Computed(String()),
			"master_internal_name":              Computed(String()),
			"network_cidr":                      Computed(String()),
			"additional_network_cidrs":          Computed(List(String())),
			"network_id":                        Computed(String()),
			"topology":                          Computed(Ptr(Struct(DataSourceTopologySpec()))),
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
			"egress_proxy":                      Computed(Ptr(Struct(DataSourceEgressProxySpec()))),
			"ssh_key_name":                      Computed(Ptr(String())),
			"kubernetes_api_access":             Computed(List(String())),
			"isolate_masters":                   Computed(Ptr(Bool())),
			"update_policy":                     Computed(Ptr(String())),
			"external_policies":                 Computed(Ptr(Map(List(String())))),
			"additional_policies":               Computed(Ptr(Map(String()))),
			"file_assets":                       Computed(List(Struct(DataSourceFileAssetSpec()))),
			"etcd_cluster":                      Computed(List(Struct(DataSourceEtcdClusterSpec()))),
			"containerd":                        Computed(Ptr(Struct(DataSourceContainerdConfig()))),
			"docker":                            Computed(Ptr(Struct(DataSourceDockerConfig()))),
			"kube_dns":                          Computed(Ptr(Struct(DataSourceKubeDNSConfig()))),
			"kube_api_server":                   Computed(Ptr(Struct(DataSourceKubeAPIServerConfig()))),
			"kube_controller_manager":           Computed(Ptr(Struct(DataSourceKubeControllerManagerConfig()))),
			"external_cloud_controller_manager": Computed(Ptr(Struct(DataSourceCloudControllerManagerConfig()))),
			"kube_scheduler":                    Computed(Ptr(Struct(DataSourceKubeSchedulerConfig()))),
			"kube_proxy":                        Computed(Ptr(Struct(DataSourceKubeProxyConfig()))),
			"kubelet":                           Computed(Ptr(Struct(DataSourceKubeletConfigSpec()))),
			"master_kubelet":                    Computed(Ptr(Struct(DataSourceKubeletConfigSpec()))),
			"cloud_config":                      Computed(Ptr(Struct(DataSourceCloudConfiguration()))),
			"external_dns":                      Computed(Ptr(Struct(DataSourceExternalDNSConfig()))),
			"ntp":                               Computed(Ptr(Struct(DataSourceNTPConfig()))),
			"node_termination_handler":          Computed(Ptr(Struct(DataSourceNodeTerminationHandlerConfig()))),
			"metrics_server":                    Computed(Ptr(Struct(DataSourceMetricsServerConfig()))),
			"cert_manager":                      Computed(Ptr(Struct(DataSourceCertManagerConfig()))),
			"aws_load_balancer_controller":      Computed(Ptr(Struct(DataSourceAWSLoadBalancerControllerConfig()))),
			"networking":                        Computed(Ptr(Struct(DataSourceNetworkingSpec()))),
			"api":                               Computed(Ptr(Struct(DataSourceAccessSpec()))),
			"authentication":                    Computed(Ptr(Struct(DataSourceAuthenticationSpec()))),
			"authorization":                     Computed(Ptr(Struct(DataSourceAuthorizationSpec()))),
			"node_authorization":                Computed(Ptr(Struct(DataSourceNodeAuthorizationSpec()))),
			"cloud_labels":                      Computed(Map(String())),
			"hooks":                             Computed(List(Struct(DataSourceHookSpec()))),
			"assets":                            Computed(Ptr(Struct(DataSourceAssets()))),
			"iam":                               Computed(Ptr(Struct(DataSourceIAMSpec()))),
			"encryption_config":                 Computed(Ptr(Bool())),
			"disable_subnet_tags":               Computed(Bool()),
			"use_host_certificates":             Computed(Ptr(Bool())),
			"sysctl_parameters":                 Computed(List(String())),
			"rolling_update":                    Computed(Ptr(Struct(DataSourceRollingUpdate()))),
			"cluster_autoscaler":                Computed(Ptr(Struct(DataSourceClusterAutoscalerConfig()))),
		},
	}
}

func ExpandDataSourceClusterSpec(in map[string]interface{}) kops.ClusterSpec {
	if in == nil {
		panic("expand ClusterSpec failure, in is nil")
	}
	out := kops.ClusterSpec{}
	if in, ok := in["channel"]; ok && in != nil {
		out.Channel = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["addons"]; ok && in != nil {
		out.Addons = func(in interface{}) []kops.AddonSpec {
			var out []kops.AddonSpec
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.AddonSpec {
					if in == nil {
						return kops.AddonSpec{}
					}
					return ExpandDataSourceAddonSpec(in.(map[string]interface{}))
				}(in))
			}
			return out
		}(in)
	}
	if in, ok := in["config_base"]; ok && in != nil {
		out.ConfigBase = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cloud_provider"]; ok && in != nil {
		out.CloudProvider = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["container_runtime"]; ok && in != nil {
		out.ContainerRuntime = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["kubernetes_version"]; ok && in != nil {
		out.KubernetesVersion = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["subnet"]; ok && in != nil {
		out.Subnets = func(in interface{}) []kops.ClusterSubnetSpec {
			var out []kops.ClusterSubnetSpec
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.ClusterSubnetSpec {
					if in == nil {
						return kops.ClusterSubnetSpec{}
					}
					return ExpandDataSourceClusterSubnetSpec(in.(map[string]interface{}))
				}(in))
			}
			return out
		}(in)
	}
	if in, ok := in["project"]; ok && in != nil {
		out.Project = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["master_public_name"]; ok && in != nil {
		out.MasterPublicName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["master_internal_name"]; ok && in != nil {
		out.MasterInternalName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["network_cidr"]; ok && in != nil {
		out.NetworkCIDR = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["additional_network_cidrs"]; ok && in != nil {
		out.AdditionalNetworkCIDRs = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["network_id"]; ok && in != nil {
		out.NetworkID = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["topology"]; ok && in != nil {
		out.Topology = func(in interface{}) *kops.TopologySpec {
			if in == nil {
				return nil
			}
			return func(in kops.TopologySpec) *kops.TopologySpec { return &in }(func(in interface{}) kops.TopologySpec {
				if in == nil {
					return kops.TopologySpec{}
				}
				return ExpandDataSourceTopologySpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["secret_store"]; ok && in != nil {
		out.SecretStore = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["key_store"]; ok && in != nil {
		out.KeyStore = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["config_store"]; ok && in != nil {
		out.ConfigStore = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["dns_zone"]; ok && in != nil {
		out.DNSZone = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["additional_sans"]; ok && in != nil {
		out.AdditionalSANs = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["cluster_dns_domain"]; ok && in != nil {
		out.ClusterDNSDomain = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["service_cluster_ip_range"]; ok && in != nil {
		out.ServiceClusterIPRange = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["pod_cidr"]; ok && in != nil {
		out.PodCIDR = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["non_masquerade_cidr"]; ok && in != nil {
		out.NonMasqueradeCIDR = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["ssh_access"]; ok && in != nil {
		out.SSHAccess = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["node_port_access"]; ok && in != nil {
		out.NodePortAccess = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["egress_proxy"]; ok && in != nil {
		out.EgressProxy = func(in interface{}) *kops.EgressProxySpec {
			if in == nil {
				return nil
			}
			return func(in kops.EgressProxySpec) *kops.EgressProxySpec { return &in }(func(in interface{}) kops.EgressProxySpec {
				if in == nil {
					return kops.EgressProxySpec{}
				}
				return ExpandDataSourceEgressProxySpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["ssh_key_name"]; ok && in != nil {
		out.SSHKeyName = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["kubernetes_api_access"]; ok && in != nil {
		out.KubernetesAPIAccess = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["isolate_masters"]; ok && in != nil {
		out.IsolateMasters = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["update_policy"]; ok && in != nil {
		out.UpdatePolicy = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["external_policies"]; ok && in != nil {
		out.ExternalPolicies = func(in interface{}) *map[string][]string {
			if in == nil {
				return nil
			}
			return func(in map[string][]string) *map[string][]string { return &in }(func(in interface{}) map[string][]string {
				if in == nil {
					return nil
				}
				out := map[string][]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = func(in interface{}) []string {
						var out []string
						for _, in := range in.([]interface{}) {
							out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
						}
						return out
					}(in)
				}
				return out
			}(in))
		}(in)
	}
	if in, ok := in["additional_policies"]; ok && in != nil {
		out.AdditionalPolicies = func(in interface{}) *map[string]string {
			if in == nil {
				return nil
			}
			return func(in map[string]string) *map[string]string { return &in }(func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = func(in interface{}) string { return string(in.(string)) }(in)
				}
				return out
			}(in))
		}(in)
	}
	if in, ok := in["file_assets"]; ok && in != nil {
		out.FileAssets = func(in interface{}) []kops.FileAssetSpec {
			var out []kops.FileAssetSpec
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.FileAssetSpec {
					if in == nil {
						return kops.FileAssetSpec{}
					}
					return ExpandDataSourceFileAssetSpec(in.(map[string]interface{}))
				}(in))
			}
			return out
		}(in)
	}
	if in, ok := in["etcd_cluster"]; ok && in != nil {
		out.EtcdClusters = func(in interface{}) []kops.EtcdClusterSpec {
			var out []kops.EtcdClusterSpec
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.EtcdClusterSpec {
					if in == nil {
						return kops.EtcdClusterSpec{}
					}
					return ExpandDataSourceEtcdClusterSpec(in.(map[string]interface{}))
				}(in))
			}
			return out
		}(in)
	}
	if in, ok := in["containerd"]; ok && in != nil {
		out.Containerd = func(in interface{}) *kops.ContainerdConfig {
			if in == nil {
				return nil
			}
			return func(in kops.ContainerdConfig) *kops.ContainerdConfig { return &in }(func(in interface{}) kops.ContainerdConfig {
				if in == nil {
					return kops.ContainerdConfig{}
				}
				return ExpandDataSourceContainerdConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["docker"]; ok && in != nil {
		out.Docker = func(in interface{}) *kops.DockerConfig {
			if in == nil {
				return nil
			}
			return func(in kops.DockerConfig) *kops.DockerConfig { return &in }(func(in interface{}) kops.DockerConfig {
				if in == nil {
					return kops.DockerConfig{}
				}
				return ExpandDataSourceDockerConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["kube_dns"]; ok && in != nil {
		out.KubeDNS = func(in interface{}) *kops.KubeDNSConfig {
			if in == nil {
				return nil
			}
			return func(in kops.KubeDNSConfig) *kops.KubeDNSConfig { return &in }(func(in interface{}) kops.KubeDNSConfig {
				if in == nil {
					return kops.KubeDNSConfig{}
				}
				return ExpandDataSourceKubeDNSConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["kube_api_server"]; ok && in != nil {
		out.KubeAPIServer = func(in interface{}) *kops.KubeAPIServerConfig {
			if in == nil {
				return nil
			}
			return func(in kops.KubeAPIServerConfig) *kops.KubeAPIServerConfig { return &in }(func(in interface{}) kops.KubeAPIServerConfig {
				if in == nil {
					return kops.KubeAPIServerConfig{}
				}
				return ExpandDataSourceKubeAPIServerConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["kube_controller_manager"]; ok && in != nil {
		out.KubeControllerManager = func(in interface{}) *kops.KubeControllerManagerConfig {
			if in == nil {
				return nil
			}
			return func(in kops.KubeControllerManagerConfig) *kops.KubeControllerManagerConfig { return &in }(func(in interface{}) kops.KubeControllerManagerConfig {
				if in == nil {
					return kops.KubeControllerManagerConfig{}
				}
				return ExpandDataSourceKubeControllerManagerConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["external_cloud_controller_manager"]; ok && in != nil {
		out.ExternalCloudControllerManager = func(in interface{}) *kops.CloudControllerManagerConfig {
			if in == nil {
				return nil
			}
			return func(in kops.CloudControllerManagerConfig) *kops.CloudControllerManagerConfig { return &in }(func(in interface{}) kops.CloudControllerManagerConfig {
				if in == nil {
					return kops.CloudControllerManagerConfig{}
				}
				return ExpandDataSourceCloudControllerManagerConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["kube_scheduler"]; ok && in != nil {
		out.KubeScheduler = func(in interface{}) *kops.KubeSchedulerConfig {
			if in == nil {
				return nil
			}
			return func(in kops.KubeSchedulerConfig) *kops.KubeSchedulerConfig { return &in }(func(in interface{}) kops.KubeSchedulerConfig {
				if in == nil {
					return kops.KubeSchedulerConfig{}
				}
				return ExpandDataSourceKubeSchedulerConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["kube_proxy"]; ok && in != nil {
		out.KubeProxy = func(in interface{}) *kops.KubeProxyConfig {
			if in == nil {
				return nil
			}
			return func(in kops.KubeProxyConfig) *kops.KubeProxyConfig { return &in }(func(in interface{}) kops.KubeProxyConfig {
				if in == nil {
					return kops.KubeProxyConfig{}
				}
				return ExpandDataSourceKubeProxyConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["kubelet"]; ok && in != nil {
		out.Kubelet = func(in interface{}) *kops.KubeletConfigSpec {
			if in == nil {
				return nil
			}
			return func(in kops.KubeletConfigSpec) *kops.KubeletConfigSpec { return &in }(func(in interface{}) kops.KubeletConfigSpec {
				if in == nil {
					return kops.KubeletConfigSpec{}
				}
				return ExpandDataSourceKubeletConfigSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["master_kubelet"]; ok && in != nil {
		out.MasterKubelet = func(in interface{}) *kops.KubeletConfigSpec {
			if in == nil {
				return nil
			}
			return func(in kops.KubeletConfigSpec) *kops.KubeletConfigSpec { return &in }(func(in interface{}) kops.KubeletConfigSpec {
				if in == nil {
					return kops.KubeletConfigSpec{}
				}
				return ExpandDataSourceKubeletConfigSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["cloud_config"]; ok && in != nil {
		out.CloudConfig = func(in interface{}) *kops.CloudConfiguration {
			if in == nil {
				return nil
			}
			return func(in kops.CloudConfiguration) *kops.CloudConfiguration { return &in }(func(in interface{}) kops.CloudConfiguration {
				if in == nil {
					return kops.CloudConfiguration{}
				}
				return ExpandDataSourceCloudConfiguration(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["external_dns"]; ok && in != nil {
		out.ExternalDNS = func(in interface{}) *kops.ExternalDNSConfig {
			if in == nil {
				return nil
			}
			return func(in kops.ExternalDNSConfig) *kops.ExternalDNSConfig { return &in }(func(in interface{}) kops.ExternalDNSConfig {
				if in == nil {
					return kops.ExternalDNSConfig{}
				}
				return ExpandDataSourceExternalDNSConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["ntp"]; ok && in != nil {
		out.NTP = func(in interface{}) *kops.NTPConfig {
			if in == nil {
				return nil
			}
			return func(in kops.NTPConfig) *kops.NTPConfig { return &in }(func(in interface{}) kops.NTPConfig {
				if in == nil {
					return kops.NTPConfig{}
				}
				return ExpandDataSourceNTPConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["node_termination_handler"]; ok && in != nil {
		out.NodeTerminationHandler = func(in interface{}) *kops.NodeTerminationHandlerConfig {
			if in == nil {
				return nil
			}
			return func(in kops.NodeTerminationHandlerConfig) *kops.NodeTerminationHandlerConfig { return &in }(func(in interface{}) kops.NodeTerminationHandlerConfig {
				if in == nil {
					return kops.NodeTerminationHandlerConfig{}
				}
				return ExpandDataSourceNodeTerminationHandlerConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["metrics_server"]; ok && in != nil {
		out.MetricsServer = func(in interface{}) *kops.MetricsServerConfig {
			if in == nil {
				return nil
			}
			return func(in kops.MetricsServerConfig) *kops.MetricsServerConfig { return &in }(func(in interface{}) kops.MetricsServerConfig {
				if in == nil {
					return kops.MetricsServerConfig{}
				}
				return ExpandDataSourceMetricsServerConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["cert_manager"]; ok && in != nil {
		out.CertManager = func(in interface{}) *kops.CertManagerConfig {
			if in == nil {
				return nil
			}
			return func(in kops.CertManagerConfig) *kops.CertManagerConfig { return &in }(func(in interface{}) kops.CertManagerConfig {
				if in == nil {
					return kops.CertManagerConfig{}
				}
				return ExpandDataSourceCertManagerConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["aws_load_balancer_controller"]; ok && in != nil {
		out.AWSLoadBalancerController = func(in interface{}) *kops.AWSLoadBalancerControllerConfig {
			if in == nil {
				return nil
			}
			return func(in kops.AWSLoadBalancerControllerConfig) *kops.AWSLoadBalancerControllerConfig { return &in }(func(in interface{}) kops.AWSLoadBalancerControllerConfig {
				if in == nil {
					return kops.AWSLoadBalancerControllerConfig{}
				}
				return ExpandDataSourceAWSLoadBalancerControllerConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["networking"]; ok && in != nil {
		out.Networking = func(in interface{}) *kops.NetworkingSpec {
			if in == nil {
				return nil
			}
			return func(in kops.NetworkingSpec) *kops.NetworkingSpec { return &in }(func(in interface{}) kops.NetworkingSpec {
				if in == nil {
					return kops.NetworkingSpec{}
				}
				return ExpandDataSourceNetworkingSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["api"]; ok && in != nil {
		out.API = func(in interface{}) *kops.AccessSpec {
			if in == nil {
				return nil
			}
			return func(in kops.AccessSpec) *kops.AccessSpec { return &in }(func(in interface{}) kops.AccessSpec {
				if in == nil {
					return kops.AccessSpec{}
				}
				return ExpandDataSourceAccessSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["authentication"]; ok && in != nil {
		out.Authentication = func(in interface{}) *kops.AuthenticationSpec {
			if in == nil {
				return nil
			}
			return func(in kops.AuthenticationSpec) *kops.AuthenticationSpec { return &in }(func(in interface{}) kops.AuthenticationSpec {
				if in == nil {
					return kops.AuthenticationSpec{}
				}
				return ExpandDataSourceAuthenticationSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["authorization"]; ok && in != nil {
		out.Authorization = func(in interface{}) *kops.AuthorizationSpec {
			if in == nil {
				return nil
			}
			return func(in kops.AuthorizationSpec) *kops.AuthorizationSpec { return &in }(func(in interface{}) kops.AuthorizationSpec {
				if in == nil {
					return kops.AuthorizationSpec{}
				}
				return ExpandDataSourceAuthorizationSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["node_authorization"]; ok && in != nil {
		out.NodeAuthorization = func(in interface{}) *kops.NodeAuthorizationSpec {
			if in == nil {
				return nil
			}
			return func(in kops.NodeAuthorizationSpec) *kops.NodeAuthorizationSpec { return &in }(func(in interface{}) kops.NodeAuthorizationSpec {
				if in == nil {
					return kops.NodeAuthorizationSpec{}
				}
				return ExpandDataSourceNodeAuthorizationSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["cloud_labels"]; ok && in != nil {
		out.CloudLabels = func(in interface{}) map[string]string {
			if in == nil {
				return nil
			}
			out := map[string]string{}
			for key, in := range in.(map[string]interface{}) {
				out[key] = func(in interface{}) string { return string(in.(string)) }(in)
			}
			return out
		}(in)
	}
	if in, ok := in["hooks"]; ok && in != nil {
		out.Hooks = func(in interface{}) []kops.HookSpec {
			var out []kops.HookSpec
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.HookSpec {
					if in == nil {
						return kops.HookSpec{}
					}
					return ExpandDataSourceHookSpec(in.(map[string]interface{}))
				}(in))
			}
			return out
		}(in)
	}
	if in, ok := in["assets"]; ok && in != nil {
		out.Assets = func(in interface{}) *kops.Assets {
			if in == nil {
				return nil
			}
			return func(in kops.Assets) *kops.Assets { return &in }(func(in interface{}) kops.Assets {
				if in == nil {
					return kops.Assets{}
				}
				return ExpandDataSourceAssets(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["iam"]; ok && in != nil {
		out.IAM = func(in interface{}) *kops.IAMSpec {
			if in == nil {
				return nil
			}
			return func(in kops.IAMSpec) *kops.IAMSpec { return &in }(func(in interface{}) kops.IAMSpec {
				if in == nil {
					return kops.IAMSpec{}
				}
				return ExpandDataSourceIAMSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["encryption_config"]; ok && in != nil {
		out.EncryptionConfig = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["disable_subnet_tags"]; ok && in != nil {
		out.DisableSubnetTags = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["use_host_certificates"]; ok && in != nil {
		out.UseHostCertificates = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["sysctl_parameters"]; ok && in != nil {
		out.SysctlParameters = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["rolling_update"]; ok && in != nil {
		out.RollingUpdate = func(in interface{}) *kops.RollingUpdate {
			if in == nil {
				return nil
			}
			return func(in kops.RollingUpdate) *kops.RollingUpdate { return &in }(func(in interface{}) kops.RollingUpdate {
				if in == nil {
					return kops.RollingUpdate{}
				}
				return ExpandDataSourceRollingUpdate(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["cluster_autoscaler"]; ok && in != nil {
		out.ClusterAutoscaler = func(in interface{}) *kops.ClusterAutoscalerConfig {
			if in == nil {
				return nil
			}
			return func(in kops.ClusterAutoscalerConfig) *kops.ClusterAutoscalerConfig { return &in }(func(in interface{}) kops.ClusterAutoscalerConfig {
				if in == nil {
					return kops.ClusterAutoscalerConfig{}
				}
				return ExpandDataSourceClusterAutoscalerConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	return out
}
