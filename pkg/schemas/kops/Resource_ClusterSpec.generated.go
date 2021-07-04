package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ExpandResourceClusterSpec(in map[string]interface{}) kops.ClusterSpec {
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
					return ExpandResourceAddonSpec(in.(map[string]interface{}))
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
					return ExpandResourceClusterSubnetSpec(in.(map[string]interface{}))
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
				return ExpandResourceTopologySpec(in.(map[string]interface{}))
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
				return ExpandResourceEgressProxySpec(in.(map[string]interface{}))
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
					return ExpandResourceFileAssetSpec(in.(map[string]interface{}))
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
					return ExpandResourceEtcdClusterSpec(in.(map[string]interface{}))
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
				return ExpandResourceContainerdConfig(in.(map[string]interface{}))
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
				return ExpandResourceDockerConfig(in.(map[string]interface{}))
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
				return ExpandResourceKubeDNSConfig(in.(map[string]interface{}))
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
				return ExpandResourceKubeAPIServerConfig(in.(map[string]interface{}))
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
				return ExpandResourceKubeControllerManagerConfig(in.(map[string]interface{}))
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
				return ExpandResourceCloudControllerManagerConfig(in.(map[string]interface{}))
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
				return ExpandResourceKubeSchedulerConfig(in.(map[string]interface{}))
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
				return ExpandResourceKubeProxyConfig(in.(map[string]interface{}))
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
				return ExpandResourceKubeletConfigSpec(in.(map[string]interface{}))
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
				return ExpandResourceKubeletConfigSpec(in.(map[string]interface{}))
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
				return ExpandResourceCloudConfiguration(in.(map[string]interface{}))
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
				return ExpandResourceExternalDNSConfig(in.(map[string]interface{}))
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
				return ExpandResourceNTPConfig(in.(map[string]interface{}))
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
				return ExpandResourceNodeTerminationHandlerConfig(in.(map[string]interface{}))
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
				return ExpandResourceMetricsServerConfig(in.(map[string]interface{}))
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
				return ExpandResourceCertManagerConfig(in.(map[string]interface{}))
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
				return ExpandResourceAWSLoadBalancerControllerConfig(in.(map[string]interface{}))
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
				return ExpandResourceNetworkingSpec(in.(map[string]interface{}))
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
				return ExpandResourceAccessSpec(in.(map[string]interface{}))
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
				return ExpandResourceAuthenticationSpec(in.(map[string]interface{}))
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
				return ExpandResourceAuthorizationSpec(in.(map[string]interface{}))
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
				return ExpandResourceNodeAuthorizationSpec(in.(map[string]interface{}))
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
					return ExpandResourceHookSpec(in.(map[string]interface{}))
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
				return ExpandResourceAssets(in.(map[string]interface{}))
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
				return ExpandResourceIAMSpec(in.(map[string]interface{}))
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
				return ExpandResourceRollingUpdate(in.(map[string]interface{}))
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
				return ExpandResourceClusterAutoscalerConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	return out
}

func FlattenResourceClusterSpecInto(in kops.ClusterSpec, out map[string]interface{}) {
	out["channel"] = func(in string) interface{} { return string(in) }(in.Channel)
	out["addons"] = func(in []kops.AddonSpec) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in kops.AddonSpec) interface{} { return FlattenResourceAddonSpec(in) }(in))
		}
		return out
	}(in.Addons)
	out["config_base"] = func(in string) interface{} { return string(in) }(in.ConfigBase)
	out["cloud_provider"] = func(in string) interface{} { return string(in) }(in.CloudProvider)
	out["container_runtime"] = func(in string) interface{} { return string(in) }(in.ContainerRuntime)
	out["kubernetes_version"] = func(in string) interface{} { return string(in) }(in.KubernetesVersion)
	out["subnet"] = func(in []kops.ClusterSubnetSpec) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in kops.ClusterSubnetSpec) interface{} { return FlattenResourceClusterSubnetSpec(in) }(in))
		}
		return out
	}(in.Subnets)
	out["project"] = func(in string) interface{} { return string(in) }(in.Project)
	out["master_public_name"] = func(in string) interface{} { return string(in) }(in.MasterPublicName)
	out["master_internal_name"] = func(in string) interface{} { return string(in) }(in.MasterInternalName)
	out["network_cidr"] = func(in string) interface{} { return string(in) }(in.NetworkCIDR)
	out["additional_network_cidrs"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.AdditionalNetworkCIDRs)
	out["network_id"] = func(in string) interface{} { return string(in) }(in.NetworkID)
	out["topology"] = func(in *kops.TopologySpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.TopologySpec) interface{} { return FlattenResourceTopologySpec(in) }(*in)
	}(in.Topology)
	out["secret_store"] = func(in string) interface{} { return string(in) }(in.SecretStore)
	out["key_store"] = func(in string) interface{} { return string(in) }(in.KeyStore)
	out["config_store"] = func(in string) interface{} { return string(in) }(in.ConfigStore)
	out["dns_zone"] = func(in string) interface{} { return string(in) }(in.DNSZone)
	out["additional_sans"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.AdditionalSANs)
	out["cluster_dns_domain"] = func(in string) interface{} { return string(in) }(in.ClusterDNSDomain)
	out["service_cluster_ip_range"] = func(in string) interface{} { return string(in) }(in.ServiceClusterIPRange)
	out["pod_cidr"] = func(in string) interface{} { return string(in) }(in.PodCIDR)
	out["non_masquerade_cidr"] = func(in string) interface{} { return string(in) }(in.NonMasqueradeCIDR)
	out["ssh_access"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.SSHAccess)
	out["node_port_access"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.NodePortAccess)
	out["egress_proxy"] = func(in *kops.EgressProxySpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.EgressProxySpec) interface{} { return FlattenResourceEgressProxySpec(in) }(*in)
	}(in.EgressProxy)
	out["ssh_key_name"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.SSHKeyName)
	out["kubernetes_api_access"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.KubernetesAPIAccess)
	out["isolate_masters"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.IsolateMasters)
	out["update_policy"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.UpdatePolicy)
	out["external_policies"] = func(in *map[string][]string) interface{} {
		if in == nil {
			return nil
		}
		return func(in map[string][]string) interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = func(in []string) interface{} {
					var out []interface{}
					for _, in := range in {
						out = append(out, func(in string) interface{} { return string(in) }(in))
					}
					return out
				}(in)
			}
			return out
		}(*in)
	}(in.ExternalPolicies)
	out["additional_policies"] = func(in *map[string]string) interface{} {
		if in == nil {
			return nil
		}
		return func(in map[string]string) interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = func(in string) interface{} { return string(in) }(in)
			}
			return out
		}(*in)
	}(in.AdditionalPolicies)
	out["file_assets"] = func(in []kops.FileAssetSpec) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in kops.FileAssetSpec) interface{} { return FlattenResourceFileAssetSpec(in) }(in))
		}
		return out
	}(in.FileAssets)
	out["etcd_cluster"] = func(in []kops.EtcdClusterSpec) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in kops.EtcdClusterSpec) interface{} { return FlattenResourceEtcdClusterSpec(in) }(in))
		}
		return out
	}(in.EtcdClusters)
	out["containerd"] = func(in *kops.ContainerdConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.ContainerdConfig) interface{} { return FlattenResourceContainerdConfig(in) }(*in)
	}(in.Containerd)
	out["docker"] = func(in *kops.DockerConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.DockerConfig) interface{} { return FlattenResourceDockerConfig(in) }(*in)
	}(in.Docker)
	out["kube_dns"] = func(in *kops.KubeDNSConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.KubeDNSConfig) interface{} { return FlattenResourceKubeDNSConfig(in) }(*in)
	}(in.KubeDNS)
	out["kube_api_server"] = func(in *kops.KubeAPIServerConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.KubeAPIServerConfig) interface{} { return FlattenResourceKubeAPIServerConfig(in) }(*in)
	}(in.KubeAPIServer)
	out["kube_controller_manager"] = func(in *kops.KubeControllerManagerConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.KubeControllerManagerConfig) interface{} {
			return FlattenResourceKubeControllerManagerConfig(in)
		}(*in)
	}(in.KubeControllerManager)
	out["external_cloud_controller_manager"] = func(in *kops.CloudControllerManagerConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.CloudControllerManagerConfig) interface{} {
			return FlattenResourceCloudControllerManagerConfig(in)
		}(*in)
	}(in.ExternalCloudControllerManager)
	out["kube_scheduler"] = func(in *kops.KubeSchedulerConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.KubeSchedulerConfig) interface{} { return FlattenResourceKubeSchedulerConfig(in) }(*in)
	}(in.KubeScheduler)
	out["kube_proxy"] = func(in *kops.KubeProxyConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.KubeProxyConfig) interface{} { return FlattenResourceKubeProxyConfig(in) }(*in)
	}(in.KubeProxy)
	out["kubelet"] = func(in *kops.KubeletConfigSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.KubeletConfigSpec) interface{} { return FlattenResourceKubeletConfigSpec(in) }(*in)
	}(in.Kubelet)
	out["master_kubelet"] = func(in *kops.KubeletConfigSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.KubeletConfigSpec) interface{} { return FlattenResourceKubeletConfigSpec(in) }(*in)
	}(in.MasterKubelet)
	out["cloud_config"] = func(in *kops.CloudConfiguration) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.CloudConfiguration) interface{} { return FlattenResourceCloudConfiguration(in) }(*in)
	}(in.CloudConfig)
	out["external_dns"] = func(in *kops.ExternalDNSConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.ExternalDNSConfig) interface{} { return FlattenResourceExternalDNSConfig(in) }(*in)
	}(in.ExternalDNS)
	out["ntp"] = func(in *kops.NTPConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.NTPConfig) interface{} { return FlattenResourceNTPConfig(in) }(*in)
	}(in.NTP)
	out["node_termination_handler"] = func(in *kops.NodeTerminationHandlerConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.NodeTerminationHandlerConfig) interface{} {
			return FlattenResourceNodeTerminationHandlerConfig(in)
		}(*in)
	}(in.NodeTerminationHandler)
	out["metrics_server"] = func(in *kops.MetricsServerConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.MetricsServerConfig) interface{} { return FlattenResourceMetricsServerConfig(in) }(*in)
	}(in.MetricsServer)
	out["cert_manager"] = func(in *kops.CertManagerConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.CertManagerConfig) interface{} { return FlattenResourceCertManagerConfig(in) }(*in)
	}(in.CertManager)
	out["aws_load_balancer_controller"] = func(in *kops.AWSLoadBalancerControllerConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.AWSLoadBalancerControllerConfig) interface{} {
			return FlattenResourceAWSLoadBalancerControllerConfig(in)
		}(*in)
	}(in.AWSLoadBalancerController)
	out["networking"] = func(in *kops.NetworkingSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.NetworkingSpec) interface{} { return FlattenResourceNetworkingSpec(in) }(*in)
	}(in.Networking)
	out["api"] = func(in *kops.AccessSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.AccessSpec) interface{} { return FlattenResourceAccessSpec(in) }(*in)
	}(in.API)
	out["authentication"] = func(in *kops.AuthenticationSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.AuthenticationSpec) interface{} { return FlattenResourceAuthenticationSpec(in) }(*in)
	}(in.Authentication)
	out["authorization"] = func(in *kops.AuthorizationSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.AuthorizationSpec) interface{} { return FlattenResourceAuthorizationSpec(in) }(*in)
	}(in.Authorization)
	out["node_authorization"] = func(in *kops.NodeAuthorizationSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.NodeAuthorizationSpec) interface{} { return FlattenResourceNodeAuthorizationSpec(in) }(*in)
	}(in.NodeAuthorization)
	out["cloud_labels"] = func(in map[string]string) interface{} {
		if in == nil {
			return nil
		}
		out := map[string]interface{}{}
		for key, in := range in {
			out[key] = func(in string) interface{} { return string(in) }(in)
		}
		return out
	}(in.CloudLabels)
	out["hooks"] = func(in []kops.HookSpec) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in kops.HookSpec) interface{} { return FlattenResourceHookSpec(in) }(in))
		}
		return out
	}(in.Hooks)
	out["assets"] = func(in *kops.Assets) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.Assets) interface{} { return FlattenResourceAssets(in) }(*in)
	}(in.Assets)
	out["iam"] = func(in *kops.IAMSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.IAMSpec) interface{} { return FlattenResourceIAMSpec(in) }(*in)
	}(in.IAM)
	out["encryption_config"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.EncryptionConfig)
	out["disable_subnet_tags"] = func(in bool) interface{} { return in }(in.DisableSubnetTags)
	out["use_host_certificates"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.UseHostCertificates)
	out["sysctl_parameters"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.SysctlParameters)
	out["rolling_update"] = func(in *kops.RollingUpdate) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.RollingUpdate) interface{} { return FlattenResourceRollingUpdate(in) }(*in)
	}(in.RollingUpdate)
	out["cluster_autoscaler"] = func(in *kops.ClusterAutoscalerConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.ClusterAutoscalerConfig) interface{} { return FlattenResourceClusterAutoscalerConfig(in) }(*in)
	}(in.ClusterAutoscaler)
}

func FlattenResourceClusterSpec(in kops.ClusterSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceClusterSpecInto(in, out)
	return out
}
