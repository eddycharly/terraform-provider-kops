package schemas

import (
	"reflect"
	"sort"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ExpandResourceClusterSpec(in map[string]interface{}) kopsv1alpha2.ClusterSpec {
	if in == nil {
		panic("expand ClusterSpec failure, in is nil")
	}
	return kopsv1alpha2.ClusterSpec{
		Channel: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["channel"]),
		Addons: func(in interface{}) []kopsv1alpha2.AddonSpec {
			return func(in interface{}) []kopsv1alpha2.AddonSpec {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.AddonSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.AddonSpec {
						if in == nil {
							return kopsv1alpha2.AddonSpec{}
						}
						return ExpandResourceAddonSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["addons"]),
		ConfigBase: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["config_base"]),
		LegacyCloudProvider: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["legacy_cloud_provider"]),
		ContainerRuntime: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["container_runtime"]),
		KubernetesVersion: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kubernetes_version"]),
		Subnets: func(in interface{}) []kopsv1alpha2.ClusterSubnetSpec {
			return func(in interface{}) []kopsv1alpha2.ClusterSubnetSpec {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.ClusterSubnetSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.ClusterSubnetSpec {
						if in == nil {
							return kopsv1alpha2.ClusterSubnetSpec{}
						}
						return ExpandResourceClusterSubnetSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["subnets"]),
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
				if in == nil {
					return nil
				}
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
		Topology: func(in interface{}) *kopsv1alpha2.TopologySpec {
			return func(in interface{}) *kopsv1alpha2.TopologySpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.TopologySpec) *kopsv1alpha2.TopologySpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.TopologySpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceTopologySpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.TopologySpec{}
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
				if in == nil {
					return nil
				}
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
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["ssh_access"]),
		NodePortAccess: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["node_port_access"]),
		EgressProxy: func(in interface{}) *kopsv1alpha2.EgressProxySpec {
			return func(in interface{}) *kopsv1alpha2.EgressProxySpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.EgressProxySpec) *kopsv1alpha2.EgressProxySpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.EgressProxySpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceEgressProxySpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.EgressProxySpec{}
				}(in))
			}(in)
		}(in["egress_proxy"]),
		SSHKeyName: func(in interface{}) *string {
			if in == nil {
				return nil
			}
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
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["kubernetes_api_access"]),
		IsolateMasters: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
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
			if in == nil {
				return nil
			}
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
					if in, ok := in.([]interface{}); ok {
						if len(in) > 0 {
							out := map[string][]string{}
							for _, in := range in {
								if in, ok := in.(map[string]interface{}); ok {
									key := ExpandString(in["key"])
									value := func(in interface{}) []string {
										return func(in interface{}) []string {
											if in == nil {
												return nil
											}
											var out []string
											for _, in := range in.([]interface{}) {
												out = append(out, string(ExpandString(in)))
											}
											return out
										}(in)
									}(in["value"])
									out[key] = value
								}
							}
							return out
						}
					}
					return nil
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
					if in, ok := in.(map[string]interface{}); ok {
						if len(in) > 0 {
							out := map[string]string{}
							for key, in := range in {
								out[key] = string(ExpandString(in))
							}
							return out
						}
					}
					return nil
				}(in))
			}(in)
		}(in["additional_policies"]),
		FileAssets: func(in interface{}) []kopsv1alpha2.FileAssetSpec {
			return func(in interface{}) []kopsv1alpha2.FileAssetSpec {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.FileAssetSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.FileAssetSpec {
						if in == nil {
							return kopsv1alpha2.FileAssetSpec{}
						}
						return ExpandResourceFileAssetSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["file_assets"]),
		EtcdClusters: func(in interface{}) []kopsv1alpha2.EtcdClusterSpec {
			return func(in interface{}) []kopsv1alpha2.EtcdClusterSpec {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.EtcdClusterSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.EtcdClusterSpec {
						if in == nil {
							return kopsv1alpha2.EtcdClusterSpec{}
						}
						return ExpandResourceEtcdClusterSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["etcd_cluster"]),
		Containerd: func(in interface{}) *kopsv1alpha2.ContainerdConfig {
			return func(in interface{}) *kopsv1alpha2.ContainerdConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.ContainerdConfig) *kopsv1alpha2.ContainerdConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.ContainerdConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceContainerdConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.ContainerdConfig{}
				}(in))
			}(in)
		}(in["containerd"]),
		Docker: func(in interface{}) *kopsv1alpha2.DockerConfig {
			return func(in interface{}) *kopsv1alpha2.DockerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.DockerConfig) *kopsv1alpha2.DockerConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.DockerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceDockerConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.DockerConfig{}
				}(in))
			}(in)
		}(in["docker"]),
		KubeDNS: func(in interface{}) *kopsv1alpha2.KubeDNSConfig {
			return func(in interface{}) *kopsv1alpha2.KubeDNSConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.KubeDNSConfig) *kopsv1alpha2.KubeDNSConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.KubeDNSConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceKubeDNSConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.KubeDNSConfig{}
				}(in))
			}(in)
		}(in["kube_dns"]),
		KubeAPIServer: func(in interface{}) *kopsv1alpha2.KubeAPIServerConfig {
			return func(in interface{}) *kopsv1alpha2.KubeAPIServerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.KubeAPIServerConfig) *kopsv1alpha2.KubeAPIServerConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.KubeAPIServerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceKubeAPIServerConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.KubeAPIServerConfig{}
				}(in))
			}(in)
		}(in["kube_api_server"]),
		KubeControllerManager: func(in interface{}) *kopsv1alpha2.KubeControllerManagerConfig {
			return func(in interface{}) *kopsv1alpha2.KubeControllerManagerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.KubeControllerManagerConfig) *kopsv1alpha2.KubeControllerManagerConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.KubeControllerManagerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceKubeControllerManagerConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.KubeControllerManagerConfig{}
				}(in))
			}(in)
		}(in["kube_controller_manager"]),
		ExternalCloudControllerManager: func(in interface{}) *kopsv1alpha2.CloudControllerManagerConfig {
			return func(in interface{}) *kopsv1alpha2.CloudControllerManagerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.CloudControllerManagerConfig) *kopsv1alpha2.CloudControllerManagerConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.CloudControllerManagerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceCloudControllerManagerConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.CloudControllerManagerConfig{}
				}(in))
			}(in)
		}(in["external_cloud_controller_manager"]),
		KubeScheduler: func(in interface{}) *kopsv1alpha2.KubeSchedulerConfig {
			return func(in interface{}) *kopsv1alpha2.KubeSchedulerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.KubeSchedulerConfig) *kopsv1alpha2.KubeSchedulerConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.KubeSchedulerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceKubeSchedulerConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.KubeSchedulerConfig{}
				}(in))
			}(in)
		}(in["kube_scheduler"]),
		KubeProxy: func(in interface{}) *kopsv1alpha2.KubeProxyConfig {
			return func(in interface{}) *kopsv1alpha2.KubeProxyConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.KubeProxyConfig) *kopsv1alpha2.KubeProxyConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.KubeProxyConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceKubeProxyConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.KubeProxyConfig{}
				}(in))
			}(in)
		}(in["kube_proxy"]),
		Kubelet: func(in interface{}) *kopsv1alpha2.KubeletConfigSpec {
			return func(in interface{}) *kopsv1alpha2.KubeletConfigSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.KubeletConfigSpec) *kopsv1alpha2.KubeletConfigSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.KubeletConfigSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceKubeletConfigSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.KubeletConfigSpec{}
				}(in))
			}(in)
		}(in["kubelet"]),
		ControlPlaneKubelet: func(in interface{}) *kopsv1alpha2.KubeletConfigSpec {
			return func(in interface{}) *kopsv1alpha2.KubeletConfigSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.KubeletConfigSpec) *kopsv1alpha2.KubeletConfigSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.KubeletConfigSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceKubeletConfigSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.KubeletConfigSpec{}
				}(in))
			}(in)
		}(in["control_plane_kubelet"]),
		CloudConfig: func(in interface{}) *kopsv1alpha2.CloudConfiguration {
			return func(in interface{}) *kopsv1alpha2.CloudConfiguration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.CloudConfiguration) *kopsv1alpha2.CloudConfiguration {
					return &in
				}(func(in interface{}) kopsv1alpha2.CloudConfiguration {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceCloudConfiguration(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.CloudConfiguration{}
				}(in))
			}(in)
		}(in["cloud_config"]),
		ExternalDNS: func(in interface{}) *kopsv1alpha2.ExternalDNSConfig {
			return func(in interface{}) *kopsv1alpha2.ExternalDNSConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.ExternalDNSConfig) *kopsv1alpha2.ExternalDNSConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.ExternalDNSConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceExternalDNSConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.ExternalDNSConfig{}
				}(in))
			}(in)
		}(in["external_dns"]),
		NTP: func(in interface{}) *kopsv1alpha2.NTPConfig {
			return func(in interface{}) *kopsv1alpha2.NTPConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.NTPConfig) *kopsv1alpha2.NTPConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.NTPConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceNTPConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.NTPConfig{}
				}(in))
			}(in)
		}(in["ntp"]),
		NodeTerminationHandler: func(in interface{}) *kopsv1alpha2.NodeTerminationHandlerSpec {
			return func(in interface{}) *kopsv1alpha2.NodeTerminationHandlerSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.NodeTerminationHandlerSpec) *kopsv1alpha2.NodeTerminationHandlerSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.NodeTerminationHandlerSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceNodeTerminationHandlerSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.NodeTerminationHandlerSpec{}
				}(in))
			}(in)
		}(in["node_termination_handler"]),
		NodeProblemDetector: func(in interface{}) *kopsv1alpha2.NodeProblemDetectorConfig {
			return func(in interface{}) *kopsv1alpha2.NodeProblemDetectorConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.NodeProblemDetectorConfig) *kopsv1alpha2.NodeProblemDetectorConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.NodeProblemDetectorConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceNodeProblemDetectorConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.NodeProblemDetectorConfig{}
				}(in))
			}(in)
		}(in["node_problem_detector"]),
		MetricsServer: func(in interface{}) *kopsv1alpha2.MetricsServerConfig {
			return func(in interface{}) *kopsv1alpha2.MetricsServerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.MetricsServerConfig) *kopsv1alpha2.MetricsServerConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.MetricsServerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceMetricsServerConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.MetricsServerConfig{}
				}(in))
			}(in)
		}(in["metrics_server"]),
		CertManager: func(in interface{}) *kopsv1alpha2.CertManagerConfig {
			return func(in interface{}) *kopsv1alpha2.CertManagerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.CertManagerConfig) *kopsv1alpha2.CertManagerConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.CertManagerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceCertManagerConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.CertManagerConfig{}
				}(in))
			}(in)
		}(in["cert_manager"]),
		AWSLoadBalancerController: func(in interface{}) *kopsv1alpha2.LoadBalancerControllerSpec {
			return func(in interface{}) *kopsv1alpha2.LoadBalancerControllerSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.LoadBalancerControllerSpec) *kopsv1alpha2.LoadBalancerControllerSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.LoadBalancerControllerSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceLoadBalancerControllerSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.LoadBalancerControllerSpec{}
				}(in))
			}(in)
		}(in["aws_load_balancer_controller"]),
		LegacyNetworking: func(in interface{}) *kopsv1alpha2.NetworkingSpec {
			return func(in interface{}) *kopsv1alpha2.NetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.NetworkingSpec) *kopsv1alpha2.NetworkingSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.NetworkingSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceNetworkingSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.NetworkingSpec{}
				}(in))
			}(in)
		}(in["legacy_networking"]),
		Networking: func(in interface{}) kopsv1alpha2.NetworkingSpec {
			return func(in interface{}) kopsv1alpha2.NetworkingSpec {
				if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
					return ExpandResourceNetworkingSpec(in[0].(map[string]interface{}))
				}
				return kopsv1alpha2.NetworkingSpec{}
			}(in)
		}(in["networking"]),
		LegacyAPI: func(in interface{}) *kopsv1alpha2.APISpec {
			return func(in interface{}) *kopsv1alpha2.APISpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.APISpec) *kopsv1alpha2.APISpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.APISpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceAPISpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.APISpec{}
				}(in))
			}(in)
		}(in["legacy_api"]),
		API: func(in interface{}) kopsv1alpha2.APISpec {
			return func(in interface{}) kopsv1alpha2.APISpec {
				if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
					return ExpandResourceAPISpec(in[0].(map[string]interface{}))
				}
				return kopsv1alpha2.APISpec{}
			}(in)
		}(in["api"]),
		Authentication: func(in interface{}) *kopsv1alpha2.AuthenticationSpec {
			return func(in interface{}) *kopsv1alpha2.AuthenticationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.AuthenticationSpec) *kopsv1alpha2.AuthenticationSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.AuthenticationSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceAuthenticationSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.AuthenticationSpec{}
				}(in))
			}(in)
		}(in["authentication"]),
		Authorization: func(in interface{}) *kopsv1alpha2.AuthorizationSpec {
			return func(in interface{}) *kopsv1alpha2.AuthorizationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.AuthorizationSpec) *kopsv1alpha2.AuthorizationSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.AuthorizationSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceAuthorizationSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.AuthorizationSpec{}
				}(in))
			}(in)
		}(in["authorization"]),
		NodeAuthorization: func(in interface{}) *kopsv1alpha2.NodeAuthorizationSpec {
			return func(in interface{}) *kopsv1alpha2.NodeAuthorizationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.NodeAuthorizationSpec) *kopsv1alpha2.NodeAuthorizationSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.NodeAuthorizationSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceNodeAuthorizationSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.NodeAuthorizationSpec{}
				}(in))
			}(in)
		}(in["node_authorization"]),
		CloudLabels: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string]string{}
						for key, in := range in {
							out[key] = string(ExpandString(in))
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["cloud_labels"]),
		Hooks: func(in interface{}) []kopsv1alpha2.HookSpec {
			return func(in interface{}) []kopsv1alpha2.HookSpec {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.HookSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.HookSpec {
						if in == nil {
							return kopsv1alpha2.HookSpec{}
						}
						return ExpandResourceHookSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["hooks"]),
		Assets: func(in interface{}) *kopsv1alpha2.Assets {
			return func(in interface{}) *kopsv1alpha2.Assets {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.Assets) *kopsv1alpha2.Assets {
					return &in
				}(func(in interface{}) kopsv1alpha2.Assets {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceAssets(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.Assets{}
				}(in))
			}(in)
		}(in["assets"]),
		IAM: func(in interface{}) *kopsv1alpha2.IAMSpec {
			return func(in interface{}) *kopsv1alpha2.IAMSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.IAMSpec) *kopsv1alpha2.IAMSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.IAMSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceIAMSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.IAMSpec{}
				}(in))
			}(in)
		}(in["iam"]),
		EncryptionConfig: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
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
		TagSubnets: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
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
		}(in["tag_subnets"]),
		UseHostCertificates: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
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
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["sysctl_parameters"]),
		RollingUpdate: func(in interface{}) *kopsv1alpha2.RollingUpdate {
			return func(in interface{}) *kopsv1alpha2.RollingUpdate {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.RollingUpdate) *kopsv1alpha2.RollingUpdate {
					return &in
				}(func(in interface{}) kopsv1alpha2.RollingUpdate {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceRollingUpdate(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.RollingUpdate{}
				}(in))
			}(in)
		}(in["rolling_update"]),
		ClusterAutoscaler: func(in interface{}) *kopsv1alpha2.ClusterAutoscalerConfig {
			return func(in interface{}) *kopsv1alpha2.ClusterAutoscalerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.ClusterAutoscalerConfig) *kopsv1alpha2.ClusterAutoscalerConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.ClusterAutoscalerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceClusterAutoscalerConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.ClusterAutoscalerConfig{}
				}(in))
			}(in)
		}(in["cluster_autoscaler"]),
		WarmPool: func(in interface{}) *kopsv1alpha2.WarmPoolSpec {
			return func(in interface{}) *kopsv1alpha2.WarmPoolSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.WarmPoolSpec) *kopsv1alpha2.WarmPoolSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.WarmPoolSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceWarmPoolSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.WarmPoolSpec{}
				}(in))
			}(in)
		}(in["warm_pool"]),
		ServiceAccountIssuerDiscovery: func(in interface{}) *kopsv1alpha2.ServiceAccountIssuerDiscoveryConfig {
			return func(in interface{}) *kopsv1alpha2.ServiceAccountIssuerDiscoveryConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.ServiceAccountIssuerDiscoveryConfig) *kopsv1alpha2.ServiceAccountIssuerDiscoveryConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.ServiceAccountIssuerDiscoveryConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceServiceAccountIssuerDiscoveryConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.ServiceAccountIssuerDiscoveryConfig{}
				}(in))
			}(in)
		}(in["service_account_issuer_discovery"]),
		SnapshotController: func(in interface{}) *kopsv1alpha2.SnapshotControllerConfig {
			return func(in interface{}) *kopsv1alpha2.SnapshotControllerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.SnapshotControllerConfig) *kopsv1alpha2.SnapshotControllerConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.SnapshotControllerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceSnapshotControllerConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.SnapshotControllerConfig{}
				}(in))
			}(in)
		}(in["snapshot_controller"]),
		Karpenter: func(in interface{}) *kopsv1alpha2.KarpenterConfig {
			return func(in interface{}) *kopsv1alpha2.KarpenterConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.KarpenterConfig) *kopsv1alpha2.KarpenterConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.KarpenterConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceKarpenterConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.KarpenterConfig{}
				}(in))
			}(in)
		}(in["karpenter"]),
		PodIdentityWebhook: func(in interface{}) *kopsv1alpha2.PodIdentityWebhookSpec {
			return func(in interface{}) *kopsv1alpha2.PodIdentityWebhookSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.PodIdentityWebhookSpec) *kopsv1alpha2.PodIdentityWebhookSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.PodIdentityWebhookSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourcePodIdentityWebhookSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.PodIdentityWebhookSpec{}
				}(in))
			}(in)
		}(in["pod_identity_webhook"]),
	}
}

func FlattenResourceClusterSpecInto(in kopsv1alpha2.ClusterSpec, out map[string]interface{}) {
	out["channel"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Channel)
	out["addons"] = func(in []kopsv1alpha2.AddonSpec) interface{} {
		return func(in []kopsv1alpha2.AddonSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.AddonSpec) interface{} {
					return FlattenResourceAddonSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.Addons)
	out["config_base"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ConfigBase)
	out["legacy_cloud_provider"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.LegacyCloudProvider)
	out["container_runtime"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ContainerRuntime)
	out["kubernetes_version"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.KubernetesVersion)
	out["subnets"] = func(in []kopsv1alpha2.ClusterSubnetSpec) interface{} {
		return func(in []kopsv1alpha2.ClusterSubnetSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.ClusterSubnetSpec) interface{} {
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
	out["topology"] = func(in *kopsv1alpha2.TopologySpec) interface{} {
		return func(in *kopsv1alpha2.TopologySpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.TopologySpec) interface{} {
				return func(in kopsv1alpha2.TopologySpec) []interface{} {
					return []interface{}{FlattenResourceTopologySpec(in)}
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
	out["egress_proxy"] = func(in *kopsv1alpha2.EgressProxySpec) interface{} {
		return func(in *kopsv1alpha2.EgressProxySpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.EgressProxySpec) interface{} {
				return func(in kopsv1alpha2.EgressProxySpec) []interface{} {
					return []interface{}{FlattenResourceEgressProxySpec(in)}
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
				return func(in map[string][]string) []interface{} {
					if in == nil {
						return nil
					}
					keys := make([]string, 0, len(in))
					for key := range in {
						keys = append(keys, key)
					}
					sort.SliceStable(keys, func(i, j int) bool {
						return keys[i] < keys[j]
					})
					var out []interface{}
					for _, key := range keys {
						in := in[key]
						out = append(out, map[string]interface{}{
							"key": key,
							"value": func(in []string) []interface{} {
								var out []interface{}
								for _, in := range in {
									out = append(out, FlattenString(string(in)))
								}
								return out
							}(in),
						})
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
	out["file_assets"] = func(in []kopsv1alpha2.FileAssetSpec) interface{} {
		return func(in []kopsv1alpha2.FileAssetSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.FileAssetSpec) interface{} {
					return FlattenResourceFileAssetSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.FileAssets)
	out["etcd_cluster"] = func(in []kopsv1alpha2.EtcdClusterSpec) interface{} {
		return func(in []kopsv1alpha2.EtcdClusterSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.EtcdClusterSpec) interface{} {
					return FlattenResourceEtcdClusterSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.EtcdClusters)
	out["containerd"] = func(in *kopsv1alpha2.ContainerdConfig) interface{} {
		return func(in *kopsv1alpha2.ContainerdConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.ContainerdConfig) interface{} {
				return func(in kopsv1alpha2.ContainerdConfig) []interface{} {
					return []interface{}{FlattenResourceContainerdConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Containerd)
	out["docker"] = func(in *kopsv1alpha2.DockerConfig) interface{} {
		return func(in *kopsv1alpha2.DockerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.DockerConfig) interface{} {
				return func(in kopsv1alpha2.DockerConfig) []interface{} {
					return []interface{}{FlattenResourceDockerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Docker)
	out["kube_dns"] = func(in *kopsv1alpha2.KubeDNSConfig) interface{} {
		return func(in *kopsv1alpha2.KubeDNSConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.KubeDNSConfig) interface{} {
				return func(in kopsv1alpha2.KubeDNSConfig) []interface{} {
					return []interface{}{FlattenResourceKubeDNSConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeDNS)
	out["kube_api_server"] = func(in *kopsv1alpha2.KubeAPIServerConfig) interface{} {
		return func(in *kopsv1alpha2.KubeAPIServerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.KubeAPIServerConfig) interface{} {
				return func(in kopsv1alpha2.KubeAPIServerConfig) []interface{} {
					return []interface{}{FlattenResourceKubeAPIServerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeAPIServer)
	out["kube_controller_manager"] = func(in *kopsv1alpha2.KubeControllerManagerConfig) interface{} {
		return func(in *kopsv1alpha2.KubeControllerManagerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.KubeControllerManagerConfig) interface{} {
				return func(in kopsv1alpha2.KubeControllerManagerConfig) []interface{} {
					return []interface{}{FlattenResourceKubeControllerManagerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeControllerManager)
	out["external_cloud_controller_manager"] = func(in *kopsv1alpha2.CloudControllerManagerConfig) interface{} {
		return func(in *kopsv1alpha2.CloudControllerManagerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.CloudControllerManagerConfig) interface{} {
				return func(in kopsv1alpha2.CloudControllerManagerConfig) []interface{} {
					return []interface{}{FlattenResourceCloudControllerManagerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ExternalCloudControllerManager)
	out["kube_scheduler"] = func(in *kopsv1alpha2.KubeSchedulerConfig) interface{} {
		return func(in *kopsv1alpha2.KubeSchedulerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.KubeSchedulerConfig) interface{} {
				return func(in kopsv1alpha2.KubeSchedulerConfig) []interface{} {
					return []interface{}{FlattenResourceKubeSchedulerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeScheduler)
	out["kube_proxy"] = func(in *kopsv1alpha2.KubeProxyConfig) interface{} {
		return func(in *kopsv1alpha2.KubeProxyConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.KubeProxyConfig) interface{} {
				return func(in kopsv1alpha2.KubeProxyConfig) []interface{} {
					return []interface{}{FlattenResourceKubeProxyConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeProxy)
	out["kubelet"] = func(in *kopsv1alpha2.KubeletConfigSpec) interface{} {
		return func(in *kopsv1alpha2.KubeletConfigSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.KubeletConfigSpec) interface{} {
				return func(in kopsv1alpha2.KubeletConfigSpec) []interface{} {
					return []interface{}{FlattenResourceKubeletConfigSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Kubelet)
	out["control_plane_kubelet"] = func(in *kopsv1alpha2.KubeletConfigSpec) interface{} {
		return func(in *kopsv1alpha2.KubeletConfigSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.KubeletConfigSpec) interface{} {
				return func(in kopsv1alpha2.KubeletConfigSpec) []interface{} {
					return []interface{}{FlattenResourceKubeletConfigSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ControlPlaneKubelet)
	out["cloud_config"] = func(in *kopsv1alpha2.CloudConfiguration) interface{} {
		return func(in *kopsv1alpha2.CloudConfiguration) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.CloudConfiguration) interface{} {
				return func(in kopsv1alpha2.CloudConfiguration) []interface{} {
					return []interface{}{FlattenResourceCloudConfiguration(in)}
				}(in)
			}(*in)
		}(in)
	}(in.CloudConfig)
	out["external_dns"] = func(in *kopsv1alpha2.ExternalDNSConfig) interface{} {
		return func(in *kopsv1alpha2.ExternalDNSConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.ExternalDNSConfig) interface{} {
				return func(in kopsv1alpha2.ExternalDNSConfig) []interface{} {
					return []interface{}{FlattenResourceExternalDNSConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ExternalDNS)
	out["ntp"] = func(in *kopsv1alpha2.NTPConfig) interface{} {
		return func(in *kopsv1alpha2.NTPConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.NTPConfig) interface{} {
				return func(in kopsv1alpha2.NTPConfig) []interface{} {
					return []interface{}{FlattenResourceNTPConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NTP)
	out["node_termination_handler"] = func(in *kopsv1alpha2.NodeTerminationHandlerSpec) interface{} {
		return func(in *kopsv1alpha2.NodeTerminationHandlerSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.NodeTerminationHandlerSpec) interface{} {
				return func(in kopsv1alpha2.NodeTerminationHandlerSpec) []interface{} {
					return []interface{}{FlattenResourceNodeTerminationHandlerSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NodeTerminationHandler)
	out["node_problem_detector"] = func(in *kopsv1alpha2.NodeProblemDetectorConfig) interface{} {
		return func(in *kopsv1alpha2.NodeProblemDetectorConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.NodeProblemDetectorConfig) interface{} {
				return func(in kopsv1alpha2.NodeProblemDetectorConfig) []interface{} {
					return []interface{}{FlattenResourceNodeProblemDetectorConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NodeProblemDetector)
	out["metrics_server"] = func(in *kopsv1alpha2.MetricsServerConfig) interface{} {
		return func(in *kopsv1alpha2.MetricsServerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.MetricsServerConfig) interface{} {
				return func(in kopsv1alpha2.MetricsServerConfig) []interface{} {
					return []interface{}{FlattenResourceMetricsServerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.MetricsServer)
	out["cert_manager"] = func(in *kopsv1alpha2.CertManagerConfig) interface{} {
		return func(in *kopsv1alpha2.CertManagerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.CertManagerConfig) interface{} {
				return func(in kopsv1alpha2.CertManagerConfig) []interface{} {
					return []interface{}{FlattenResourceCertManagerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.CertManager)
	out["aws_load_balancer_controller"] = func(in *kopsv1alpha2.LoadBalancerControllerSpec) interface{} {
		return func(in *kopsv1alpha2.LoadBalancerControllerSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.LoadBalancerControllerSpec) interface{} {
				return func(in kopsv1alpha2.LoadBalancerControllerSpec) []interface{} {
					return []interface{}{FlattenResourceLoadBalancerControllerSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AWSLoadBalancerController)
	out["legacy_networking"] = func(in *kopsv1alpha2.NetworkingSpec) interface{} {
		return func(in *kopsv1alpha2.NetworkingSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.NetworkingSpec) interface{} {
				return func(in kopsv1alpha2.NetworkingSpec) []interface{} {
					return []interface{}{FlattenResourceNetworkingSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.LegacyNetworking)
	out["networking"] = func(in kopsv1alpha2.NetworkingSpec) interface{} {
		return func(in kopsv1alpha2.NetworkingSpec) []interface{} {
			return []interface{}{FlattenResourceNetworkingSpec(in)}
		}(in)
	}(in.Networking)
	out["legacy_api"] = func(in *kopsv1alpha2.APISpec) interface{} {
		return func(in *kopsv1alpha2.APISpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.APISpec) interface{} {
				return func(in kopsv1alpha2.APISpec) []interface{} {
					return []interface{}{FlattenResourceAPISpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.LegacyAPI)
	out["api"] = func(in kopsv1alpha2.APISpec) interface{} {
		return func(in kopsv1alpha2.APISpec) []interface{} {
			return []interface{}{FlattenResourceAPISpec(in)}
		}(in)
	}(in.API)
	out["authentication"] = func(in *kopsv1alpha2.AuthenticationSpec) interface{} {
		return func(in *kopsv1alpha2.AuthenticationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.AuthenticationSpec) interface{} {
				return func(in kopsv1alpha2.AuthenticationSpec) []interface{} {
					return []interface{}{FlattenResourceAuthenticationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Authentication)
	out["authorization"] = func(in *kopsv1alpha2.AuthorizationSpec) interface{} {
		return func(in *kopsv1alpha2.AuthorizationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.AuthorizationSpec) interface{} {
				return func(in kopsv1alpha2.AuthorizationSpec) []interface{} {
					return []interface{}{FlattenResourceAuthorizationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Authorization)
	out["node_authorization"] = func(in *kopsv1alpha2.NodeAuthorizationSpec) interface{} {
		return func(in *kopsv1alpha2.NodeAuthorizationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.NodeAuthorizationSpec) interface{} {
				return func(in kopsv1alpha2.NodeAuthorizationSpec) []interface{} {
					return []interface{}{FlattenResourceNodeAuthorizationSpec(in)}
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
	out["hooks"] = func(in []kopsv1alpha2.HookSpec) interface{} {
		return func(in []kopsv1alpha2.HookSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.HookSpec) interface{} {
					return FlattenResourceHookSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.Hooks)
	out["assets"] = func(in *kopsv1alpha2.Assets) interface{} {
		return func(in *kopsv1alpha2.Assets) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.Assets) interface{} {
				return func(in kopsv1alpha2.Assets) []interface{} {
					return []interface{}{FlattenResourceAssets(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Assets)
	out["iam"] = func(in *kopsv1alpha2.IAMSpec) interface{} {
		return func(in *kopsv1alpha2.IAMSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.IAMSpec) interface{} {
				return func(in kopsv1alpha2.IAMSpec) []interface{} {
					return []interface{}{FlattenResourceIAMSpec(in)}
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
	out["tag_subnets"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.TagSubnets)
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
	out["rolling_update"] = func(in *kopsv1alpha2.RollingUpdate) interface{} {
		return func(in *kopsv1alpha2.RollingUpdate) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.RollingUpdate) interface{} {
				return func(in kopsv1alpha2.RollingUpdate) []interface{} {
					return []interface{}{FlattenResourceRollingUpdate(in)}
				}(in)
			}(*in)
		}(in)
	}(in.RollingUpdate)
	out["cluster_autoscaler"] = func(in *kopsv1alpha2.ClusterAutoscalerConfig) interface{} {
		return func(in *kopsv1alpha2.ClusterAutoscalerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.ClusterAutoscalerConfig) interface{} {
				return func(in kopsv1alpha2.ClusterAutoscalerConfig) []interface{} {
					return []interface{}{FlattenResourceClusterAutoscalerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ClusterAutoscaler)
	out["warm_pool"] = func(in *kopsv1alpha2.WarmPoolSpec) interface{} {
		return func(in *kopsv1alpha2.WarmPoolSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.WarmPoolSpec) interface{} {
				return func(in kopsv1alpha2.WarmPoolSpec) []interface{} {
					return []interface{}{FlattenResourceWarmPoolSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.WarmPool)
	out["service_account_issuer_discovery"] = func(in *kopsv1alpha2.ServiceAccountIssuerDiscoveryConfig) interface{} {
		return func(in *kopsv1alpha2.ServiceAccountIssuerDiscoveryConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.ServiceAccountIssuerDiscoveryConfig) interface{} {
				return func(in kopsv1alpha2.ServiceAccountIssuerDiscoveryConfig) []interface{} {
					return []interface{}{FlattenResourceServiceAccountIssuerDiscoveryConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ServiceAccountIssuerDiscovery)
	out["snapshot_controller"] = func(in *kopsv1alpha2.SnapshotControllerConfig) interface{} {
		return func(in *kopsv1alpha2.SnapshotControllerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.SnapshotControllerConfig) interface{} {
				return func(in kopsv1alpha2.SnapshotControllerConfig) []interface{} {
					return []interface{}{FlattenResourceSnapshotControllerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.SnapshotController)
	out["karpenter"] = func(in *kopsv1alpha2.KarpenterConfig) interface{} {
		return func(in *kopsv1alpha2.KarpenterConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.KarpenterConfig) interface{} {
				return func(in kopsv1alpha2.KarpenterConfig) []interface{} {
					return []interface{}{FlattenResourceKarpenterConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Karpenter)
	out["pod_identity_webhook"] = func(in *kopsv1alpha2.PodIdentityWebhookSpec) interface{} {
		return func(in *kopsv1alpha2.PodIdentityWebhookSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.PodIdentityWebhookSpec) interface{} {
				return func(in kopsv1alpha2.PodIdentityWebhookSpec) []interface{} {
					return []interface{}{FlattenResourcePodIdentityWebhookSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.PodIdentityWebhook)
}

func FlattenResourceClusterSpec(in kopsv1alpha2.ClusterSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceClusterSpecInto(in, out)
	return out
}
