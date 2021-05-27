package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ExpandResourceClusterSpec(in map[string]interface{}) kops.ClusterSpec {
	if in == nil {
		panic("expand ClusterSpec failure, in is nil")
	}
	return kops.ClusterSpec{
		Channel: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["channel"]),
		Addons: func(in interface{}) []kops.AddonSpec {
			return func(in interface{}) []kops.AddonSpec {
				var out []kops.AddonSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.AddonSpec {
						if in == nil {
							return kops.AddonSpec{}
						}
						return (ExpandResourceAddonSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["addons"]),
		ConfigBase: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["config_base"]),
		CloudProvider: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cloud_provider"]),
		ContainerRuntime: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["container_runtime"]),
		KubernetesVersion: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kubernetes_version"]),
		Subnets: func(in interface{}) []kops.ClusterSubnetSpec {
			return func(in interface{}) []kops.ClusterSubnetSpec {
				var out []kops.ClusterSubnetSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.ClusterSubnetSpec {
						if in == nil {
							return kops.ClusterSubnetSpec{}
						}
						return (ExpandResourceClusterSubnetSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["subnet"]),
		Project: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["project"]),
		MasterPublicName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["master_public_name"]),
		MasterInternalName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["master_internal_name"]),
		NetworkCIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["network_cidr"]),
		AdditionalNetworkCIDRs: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["additional_network_cidrs"]),
		NetworkID: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["network_id"]),
		Topology: func(in interface{}) *kops.TopologySpec {
			return func(in interface{}) *kops.TopologySpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.TopologySpec) *kops.TopologySpec {
					return &in
				}(func(in interface{}) kops.TopologySpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.TopologySpec{}
					}
					return (ExpandResourceTopologySpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["topology"]),
		SecretStore: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["secret_store"]),
		KeyStore: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["key_store"]),
		ConfigStore: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["config_store"]),
		DNSZone: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["dns_zone"]),
		AdditionalSANs: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["additional_sans"]),
		ClusterDNSDomain: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_dns_domain"]),
		ServiceClusterIPRange: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["service_cluster_ip_range"]),
		PodCIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["pod_cidr"]),
		NonMasqueradeCIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["non_masquerade_cidr"]),
		SSHAccess: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["ssh_access"]),
		NodePortAccess: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["node_port_access"]),
		EgressProxy: func(in interface{}) *kops.EgressProxySpec {
			return func(in interface{}) *kops.EgressProxySpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.EgressProxySpec) *kops.EgressProxySpec {
					return &in
				}(func(in interface{}) kops.EgressProxySpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.EgressProxySpec{}
					}
					return (ExpandResourceEgressProxySpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["egress_proxy"]),
		SSHKeyName: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["ssh_key_name"]),
		KubernetesAPIAccess: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["kubernetes_api_access"]),
		IsolateMasters: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["isolate_masters"]),
		UpdatePolicy: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["update_policy"]),
		ExternalPolicies: func(in interface{}) *map[string][]string {
			return func(in interface{}) *map[string][]string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in map[string][]string) *map[string][]string {
					return &in
				}(func(in interface{}) map[string][]string {
					if in == nil {
						return nil
					}
					out := map[string][]string{}
					for key, in := range in.(map[string]interface{}) {
						out[key] = func(in interface{}) []string {
							var out []string
							for _, in := range in.([]interface{}) {
								out = append(out, string(ExpandString(in)))
							}
							return out
						}(in)
					}
					return out
				}(in))
			}(in)
		}(in["external_policies"]),
		AdditionalPolicies: func(in interface{}) *map[string]string {
			return func(in interface{}) *map[string]string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in map[string]string) *map[string]string {
					return &in
				}(func(in interface{}) map[string]string {
					if in == nil {
						return nil
					}
					out := map[string]string{}
					for key, in := range in.(map[string]interface{}) {
						out[key] = string(ExpandString(in))
					}
					return out
				}(in))
			}(in)
		}(in["additional_policies"]),
		FileAssets: func(in interface{}) []kops.FileAssetSpec {
			return func(in interface{}) []kops.FileAssetSpec {
				var out []kops.FileAssetSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.FileAssetSpec {
						if in == nil {
							return kops.FileAssetSpec{}
						}
						return (ExpandResourceFileAssetSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["file_assets"]),
		EtcdClusters: func(in interface{}) []kops.EtcdClusterSpec {
			return func(in interface{}) []kops.EtcdClusterSpec {
				var out []kops.EtcdClusterSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.EtcdClusterSpec {
						if in == nil {
							return kops.EtcdClusterSpec{}
						}
						return (ExpandResourceEtcdClusterSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["etcd_cluster"]),
		Containerd: func(in interface{}) *kops.ContainerdConfig {
			return func(in interface{}) *kops.ContainerdConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.ContainerdConfig) *kops.ContainerdConfig {
					return &in
				}(func(in interface{}) kops.ContainerdConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.ContainerdConfig{}
					}
					return (ExpandResourceContainerdConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["containerd"]),
		Docker: func(in interface{}) *kops.DockerConfig {
			return func(in interface{}) *kops.DockerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.DockerConfig) *kops.DockerConfig {
					return &in
				}(func(in interface{}) kops.DockerConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.DockerConfig{}
					}
					return (ExpandResourceDockerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["docker"]),
		KubeDNS: func(in interface{}) *kops.KubeDNSConfig {
			return func(in interface{}) *kops.KubeDNSConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeDNSConfig) *kops.KubeDNSConfig {
					return &in
				}(func(in interface{}) kops.KubeDNSConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.KubeDNSConfig{}
					}
					return (ExpandResourceKubeDNSConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kube_dns"]),
		KubeAPIServer: func(in interface{}) *kops.KubeAPIServerConfig {
			return func(in interface{}) *kops.KubeAPIServerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeAPIServerConfig) *kops.KubeAPIServerConfig {
					return &in
				}(func(in interface{}) kops.KubeAPIServerConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.KubeAPIServerConfig{}
					}
					return (ExpandResourceKubeAPIServerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kube_api_server"]),
		KubeControllerManager: func(in interface{}) *kops.KubeControllerManagerConfig {
			return func(in interface{}) *kops.KubeControllerManagerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeControllerManagerConfig) *kops.KubeControllerManagerConfig {
					return &in
				}(func(in interface{}) kops.KubeControllerManagerConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.KubeControllerManagerConfig{}
					}
					return (ExpandResourceKubeControllerManagerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kube_controller_manager"]),
		ExternalCloudControllerManager: func(in interface{}) *kops.CloudControllerManagerConfig {
			return func(in interface{}) *kops.CloudControllerManagerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.CloudControllerManagerConfig) *kops.CloudControllerManagerConfig {
					return &in
				}(func(in interface{}) kops.CloudControllerManagerConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.CloudControllerManagerConfig{}
					}
					return (ExpandResourceCloudControllerManagerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["external_cloud_controller_manager"]),
		KubeScheduler: func(in interface{}) *kops.KubeSchedulerConfig {
			return func(in interface{}) *kops.KubeSchedulerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeSchedulerConfig) *kops.KubeSchedulerConfig {
					return &in
				}(func(in interface{}) kops.KubeSchedulerConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.KubeSchedulerConfig{}
					}
					return (ExpandResourceKubeSchedulerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kube_scheduler"]),
		KubeProxy: func(in interface{}) *kops.KubeProxyConfig {
			return func(in interface{}) *kops.KubeProxyConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeProxyConfig) *kops.KubeProxyConfig {
					return &in
				}(func(in interface{}) kops.KubeProxyConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.KubeProxyConfig{}
					}
					return (ExpandResourceKubeProxyConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kube_proxy"]),
		Kubelet: func(in interface{}) *kops.KubeletConfigSpec {
			return func(in interface{}) *kops.KubeletConfigSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeletConfigSpec) *kops.KubeletConfigSpec {
					return &in
				}(func(in interface{}) kops.KubeletConfigSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.KubeletConfigSpec{}
					}
					return (ExpandResourceKubeletConfigSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kubelet"]),
		MasterKubelet: func(in interface{}) *kops.KubeletConfigSpec {
			return func(in interface{}) *kops.KubeletConfigSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeletConfigSpec) *kops.KubeletConfigSpec {
					return &in
				}(func(in interface{}) kops.KubeletConfigSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.KubeletConfigSpec{}
					}
					return (ExpandResourceKubeletConfigSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["master_kubelet"]),
		CloudConfig: func(in interface{}) *kops.CloudConfiguration {
			return func(in interface{}) *kops.CloudConfiguration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.CloudConfiguration) *kops.CloudConfiguration {
					return &in
				}(func(in interface{}) kops.CloudConfiguration {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.CloudConfiguration{}
					}
					return (ExpandResourceCloudConfiguration(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["cloud_config"]),
		ExternalDNS: func(in interface{}) *kops.ExternalDNSConfig {
			return func(in interface{}) *kops.ExternalDNSConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.ExternalDNSConfig) *kops.ExternalDNSConfig {
					return &in
				}(func(in interface{}) kops.ExternalDNSConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.ExternalDNSConfig{}
					}
					return (ExpandResourceExternalDNSConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["external_dns"]),
		NTP: func(in interface{}) *kops.NTPConfig {
			return func(in interface{}) *kops.NTPConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.NTPConfig) *kops.NTPConfig {
					return &in
				}(func(in interface{}) kops.NTPConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.NTPConfig{}
					}
					return (ExpandResourceNTPConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["ntp"]),
		NodeTerminationHandler: func(in interface{}) *kops.NodeTerminationHandlerConfig {
			return func(in interface{}) *kops.NodeTerminationHandlerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.NodeTerminationHandlerConfig) *kops.NodeTerminationHandlerConfig {
					return &in
				}(func(in interface{}) kops.NodeTerminationHandlerConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.NodeTerminationHandlerConfig{}
					}
					return (ExpandResourceNodeTerminationHandlerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["node_termination_handler"]),
		MetricsServer: func(in interface{}) *kops.MetricsServerConfig {
			return func(in interface{}) *kops.MetricsServerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.MetricsServerConfig) *kops.MetricsServerConfig {
					return &in
				}(func(in interface{}) kops.MetricsServerConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.MetricsServerConfig{}
					}
					return (ExpandResourceMetricsServerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["metrics_server"]),
		CertManager: func(in interface{}) *kops.CertManagerConfig {
			return func(in interface{}) *kops.CertManagerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.CertManagerConfig) *kops.CertManagerConfig {
					return &in
				}(func(in interface{}) kops.CertManagerConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.CertManagerConfig{}
					}
					return (ExpandResourceCertManagerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["cert_manager"]),
		AWSLoadBalancerController: func(in interface{}) *kops.AWSLoadBalancerControllerConfig {
			return func(in interface{}) *kops.AWSLoadBalancerControllerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AWSLoadBalancerControllerConfig) *kops.AWSLoadBalancerControllerConfig {
					return &in
				}(func(in interface{}) kops.AWSLoadBalancerControllerConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.AWSLoadBalancerControllerConfig{}
					}
					return (ExpandResourceAWSLoadBalancerControllerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["aws_load_balancer_controller"]),
		Networking: func(in interface{}) *kops.NetworkingSpec {
			return func(in interface{}) *kops.NetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.NetworkingSpec) *kops.NetworkingSpec {
					return &in
				}(func(in interface{}) kops.NetworkingSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.NetworkingSpec{}
					}
					return (ExpandResourceNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["networking"]),
		API: func(in interface{}) *kops.AccessSpec {
			return func(in interface{}) *kops.AccessSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AccessSpec) *kops.AccessSpec {
					return &in
				}(func(in interface{}) kops.AccessSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.AccessSpec{}
					}
					return (ExpandResourceAccessSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["api"]),
		Authentication: func(in interface{}) *kops.AuthenticationSpec {
			return func(in interface{}) *kops.AuthenticationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AuthenticationSpec) *kops.AuthenticationSpec {
					return &in
				}(func(in interface{}) kops.AuthenticationSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.AuthenticationSpec{}
					}
					return (ExpandResourceAuthenticationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["authentication"]),
		Authorization: func(in interface{}) *kops.AuthorizationSpec {
			return func(in interface{}) *kops.AuthorizationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AuthorizationSpec) *kops.AuthorizationSpec {
					return &in
				}(func(in interface{}) kops.AuthorizationSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.AuthorizationSpec{}
					}
					return (ExpandResourceAuthorizationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["authorization"]),
		NodeAuthorization: func(in interface{}) *kops.NodeAuthorizationSpec {
			return func(in interface{}) *kops.NodeAuthorizationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.NodeAuthorizationSpec) *kops.NodeAuthorizationSpec {
					return &in
				}(func(in interface{}) kops.NodeAuthorizationSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.NodeAuthorizationSpec{}
					}
					return (ExpandResourceNodeAuthorizationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["node_authorization"]),
		CloudLabels: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
		}(in["cloud_labels"]),
		Hooks: func(in interface{}) []kops.HookSpec {
			return func(in interface{}) []kops.HookSpec {
				var out []kops.HookSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.HookSpec {
						if in == nil {
							return kops.HookSpec{}
						}
						return (ExpandResourceHookSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["hooks"]),
		Assets: func(in interface{}) *kops.Assets {
			return func(in interface{}) *kops.Assets {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.Assets) *kops.Assets {
					return &in
				}(func(in interface{}) kops.Assets {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.Assets{}
					}
					return (ExpandResourceAssets(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["assets"]),
		IAM: func(in interface{}) *kops.IAMSpec {
			return func(in interface{}) *kops.IAMSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.IAMSpec) *kops.IAMSpec {
					return &in
				}(func(in interface{}) kops.IAMSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.IAMSpec{}
					}
					return (ExpandResourceIAMSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["iam"]),
		EncryptionConfig: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["encryption_config"]),
		DisableSubnetTags: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disable_subnet_tags"]),
		UseHostCertificates: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["use_host_certificates"]),
		SysctlParameters: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["sysctl_parameters"]),
		RollingUpdate: func(in interface{}) *kops.RollingUpdate {
			return func(in interface{}) *kops.RollingUpdate {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.RollingUpdate) *kops.RollingUpdate {
					return &in
				}(func(in interface{}) kops.RollingUpdate {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.RollingUpdate{}
					}
					return (ExpandResourceRollingUpdate(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["rolling_update"]),
		ClusterAutoscaler: func(in interface{}) *kops.ClusterAutoscalerConfig {
			return func(in interface{}) *kops.ClusterAutoscalerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.ClusterAutoscalerConfig) *kops.ClusterAutoscalerConfig {
					return &in
				}(func(in interface{}) kops.ClusterAutoscalerConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.ClusterAutoscalerConfig{}
					}
					return (ExpandResourceClusterAutoscalerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["cluster_autoscaler"]),
	}
}

func FlattenResourceClusterSpecInto(in kops.ClusterSpec, out map[string]interface{}) {
	out["channel"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Channel)
	out["addons"] = func(in []kops.AddonSpec) interface{} {
		return func(in []kops.AddonSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.AddonSpec) interface{} {
					return FlattenResourceAddonSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.Addons)
	out["config_base"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ConfigBase)
	out["cloud_provider"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CloudProvider)
	out["container_runtime"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ContainerRuntime)
	out["kubernetes_version"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.KubernetesVersion)
	out["subnet"] = func(in []kops.ClusterSubnetSpec) interface{} {
		return func(in []kops.ClusterSubnetSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.ClusterSubnetSpec) interface{} {
					return FlattenResourceClusterSubnetSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.Subnets)
	out["project"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Project)
	out["master_public_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.MasterPublicName)
	out["master_internal_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.MasterInternalName)
	out["network_cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.NetworkCIDR)
	out["additional_network_cidrs"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AdditionalNetworkCIDRs)
	out["network_id"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.NetworkID)
	out["topology"] = func(in *kops.TopologySpec) interface{} {
		return func(in *kops.TopologySpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.TopologySpec) interface{} {
				return func(in kops.TopologySpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceTopologySpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Topology)
	out["secret_store"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.SecretStore)
	out["key_store"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.KeyStore)
	out["config_store"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ConfigStore)
	out["dns_zone"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.DNSZone)
	out["additional_sans"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AdditionalSANs)
	out["cluster_dns_domain"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterDNSDomain)
	out["service_cluster_ip_range"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ServiceClusterIPRange)
	out["pod_cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.PodCIDR)
	out["non_masquerade_cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.NonMasqueradeCIDR)
	out["ssh_access"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.SSHAccess)
	out["node_port_access"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.NodePortAccess)
	out["egress_proxy"] = func(in *kops.EgressProxySpec) interface{} {
		return func(in *kops.EgressProxySpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.EgressProxySpec) interface{} {
				return func(in kops.EgressProxySpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceEgressProxySpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.EgressProxy)
	out["ssh_key_name"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.SSHKeyName)
	out["kubernetes_api_access"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.KubernetesAPIAccess)
	out["isolate_masters"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.IsolateMasters)
	out["update_policy"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.UpdatePolicy)
	out["external_policies"] = func(in *map[string][]string) interface{} {
		return func(in *map[string][]string) interface{} {
			if in == nil {
				return nil
			}
			return func(in map[string][]string) interface{} {
				return func(in map[string][]string) map[string]interface{} {
					if in == nil {
						return nil
					}
					out := map[string]interface{}{}
					for key, in := range in {
						out[key] = func(in []string) []interface{} {
							var out []interface{}
							for _, in := range in {
								out = append(out, FlattenString(string(in)))
							}
							return out
						}(in)
					}
					return out
				}(in)
			}(*in)
		}(in)
	}(in.ExternalPolicies)
	out["additional_policies"] = func(in *map[string]string) interface{} {
		return func(in *map[string]string) interface{} {
			if in == nil {
				return nil
			}
			return func(in map[string]string) interface{} {
				return func(in map[string]string) map[string]interface{} {
					if in == nil {
						return nil
					}
					out := map[string]interface{}{}
					for key, in := range in {
						out[key] = FlattenString(string(in))
					}
					return out
				}(in)
			}(*in)
		}(in)
	}(in.AdditionalPolicies)
	out["file_assets"] = func(in []kops.FileAssetSpec) interface{} {
		return func(in []kops.FileAssetSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.FileAssetSpec) interface{} {
					return FlattenResourceFileAssetSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.FileAssets)
	out["etcd_cluster"] = func(in []kops.EtcdClusterSpec) interface{} {
		return func(in []kops.EtcdClusterSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.EtcdClusterSpec) interface{} {
					return FlattenResourceEtcdClusterSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.EtcdClusters)
	out["containerd"] = func(in *kops.ContainerdConfig) interface{} {
		return func(in *kops.ContainerdConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.ContainerdConfig) interface{} {
				return func(in kops.ContainerdConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceContainerdConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Containerd)
	out["docker"] = func(in *kops.DockerConfig) interface{} {
		return func(in *kops.DockerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.DockerConfig) interface{} {
				return func(in kops.DockerConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceDockerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Docker)
	out["kube_dns"] = func(in *kops.KubeDNSConfig) interface{} {
		return func(in *kops.KubeDNSConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KubeDNSConfig) interface{} {
				return func(in kops.KubeDNSConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceKubeDNSConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeDNS)
	out["kube_api_server"] = func(in *kops.KubeAPIServerConfig) interface{} {
		return func(in *kops.KubeAPIServerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KubeAPIServerConfig) interface{} {
				return func(in kops.KubeAPIServerConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceKubeAPIServerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeAPIServer)
	out["kube_controller_manager"] = func(in *kops.KubeControllerManagerConfig) interface{} {
		return func(in *kops.KubeControllerManagerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KubeControllerManagerConfig) interface{} {
				return func(in kops.KubeControllerManagerConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceKubeControllerManagerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeControllerManager)
	out["external_cloud_controller_manager"] = func(in *kops.CloudControllerManagerConfig) interface{} {
		return func(in *kops.CloudControllerManagerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.CloudControllerManagerConfig) interface{} {
				return func(in kops.CloudControllerManagerConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceCloudControllerManagerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ExternalCloudControllerManager)
	out["kube_scheduler"] = func(in *kops.KubeSchedulerConfig) interface{} {
		return func(in *kops.KubeSchedulerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KubeSchedulerConfig) interface{} {
				return func(in kops.KubeSchedulerConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceKubeSchedulerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeScheduler)
	out["kube_proxy"] = func(in *kops.KubeProxyConfig) interface{} {
		return func(in *kops.KubeProxyConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KubeProxyConfig) interface{} {
				return func(in kops.KubeProxyConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceKubeProxyConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeProxy)
	out["kubelet"] = func(in *kops.KubeletConfigSpec) interface{} {
		return func(in *kops.KubeletConfigSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KubeletConfigSpec) interface{} {
				return func(in kops.KubeletConfigSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceKubeletConfigSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Kubelet)
	out["master_kubelet"] = func(in *kops.KubeletConfigSpec) interface{} {
		return func(in *kops.KubeletConfigSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KubeletConfigSpec) interface{} {
				return func(in kops.KubeletConfigSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceKubeletConfigSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.MasterKubelet)
	out["cloud_config"] = func(in *kops.CloudConfiguration) interface{} {
		return func(in *kops.CloudConfiguration) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.CloudConfiguration) interface{} {
				return func(in kops.CloudConfiguration) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceCloudConfiguration(in)}
				}(in)
			}(*in)
		}(in)
	}(in.CloudConfig)
	out["external_dns"] = func(in *kops.ExternalDNSConfig) interface{} {
		return func(in *kops.ExternalDNSConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.ExternalDNSConfig) interface{} {
				return func(in kops.ExternalDNSConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceExternalDNSConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ExternalDNS)
	out["ntp"] = func(in *kops.NTPConfig) interface{} {
		return func(in *kops.NTPConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.NTPConfig) interface{} {
				return func(in kops.NTPConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceNTPConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NTP)
	out["node_termination_handler"] = func(in *kops.NodeTerminationHandlerConfig) interface{} {
		return func(in *kops.NodeTerminationHandlerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.NodeTerminationHandlerConfig) interface{} {
				return func(in kops.NodeTerminationHandlerConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceNodeTerminationHandlerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NodeTerminationHandler)
	out["metrics_server"] = func(in *kops.MetricsServerConfig) interface{} {
		return func(in *kops.MetricsServerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.MetricsServerConfig) interface{} {
				return func(in kops.MetricsServerConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceMetricsServerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.MetricsServer)
	out["cert_manager"] = func(in *kops.CertManagerConfig) interface{} {
		return func(in *kops.CertManagerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.CertManagerConfig) interface{} {
				return func(in kops.CertManagerConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceCertManagerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.CertManager)
	out["aws_load_balancer_controller"] = func(in *kops.AWSLoadBalancerControllerConfig) interface{} {
		return func(in *kops.AWSLoadBalancerControllerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AWSLoadBalancerControllerConfig) interface{} {
				return func(in kops.AWSLoadBalancerControllerConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceAWSLoadBalancerControllerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AWSLoadBalancerController)
	out["networking"] = func(in *kops.NetworkingSpec) interface{} {
		return func(in *kops.NetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.NetworkingSpec) interface{} {
				return func(in kops.NetworkingSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Networking)
	out["api"] = func(in *kops.AccessSpec) interface{} {
		return func(in *kops.AccessSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AccessSpec) interface{} {
				return func(in kops.AccessSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceAccessSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.API)
	out["authentication"] = func(in *kops.AuthenticationSpec) interface{} {
		return func(in *kops.AuthenticationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AuthenticationSpec) interface{} {
				return func(in kops.AuthenticationSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceAuthenticationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Authentication)
	out["authorization"] = func(in *kops.AuthorizationSpec) interface{} {
		return func(in *kops.AuthorizationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AuthorizationSpec) interface{} {
				return func(in kops.AuthorizationSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceAuthorizationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Authorization)
	out["node_authorization"] = func(in *kops.NodeAuthorizationSpec) interface{} {
		return func(in *kops.NodeAuthorizationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.NodeAuthorizationSpec) interface{} {
				return func(in kops.NodeAuthorizationSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceNodeAuthorizationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NodeAuthorization)
	out["cloud_labels"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.CloudLabels)
	out["hooks"] = func(in []kops.HookSpec) interface{} {
		return func(in []kops.HookSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.HookSpec) interface{} {
					return FlattenResourceHookSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.Hooks)
	out["assets"] = func(in *kops.Assets) interface{} {
		return func(in *kops.Assets) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.Assets) interface{} {
				return func(in kops.Assets) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceAssets(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Assets)
	out["iam"] = func(in *kops.IAMSpec) interface{} {
		return func(in *kops.IAMSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.IAMSpec) interface{} {
				return func(in kops.IAMSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceIAMSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.IAM)
	out["encryption_config"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EncryptionConfig)
	out["disable_subnet_tags"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.DisableSubnetTags)
	out["use_host_certificates"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.UseHostCertificates)
	out["sysctl_parameters"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.SysctlParameters)
	out["rolling_update"] = func(in *kops.RollingUpdate) interface{} {
		return func(in *kops.RollingUpdate) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.RollingUpdate) interface{} {
				return func(in kops.RollingUpdate) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceRollingUpdate(in)}
				}(in)
			}(*in)
		}(in)
	}(in.RollingUpdate)
	out["cluster_autoscaler"] = func(in *kops.ClusterAutoscalerConfig) interface{} {
		return func(in *kops.ClusterAutoscalerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.ClusterAutoscalerConfig) interface{} {
				return func(in kops.ClusterAutoscalerConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceClusterAutoscalerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ClusterAutoscaler)
}

func FlattenResourceClusterSpec(in kops.ClusterSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceClusterSpecInto(in, out)
	return out
}
