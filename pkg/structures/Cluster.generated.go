package structures

import (
	"reflect"

	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandCluster(in map[string]interface{}) api.Cluster {
	if in == nil {
		panic("expand Cluster failure, in is nil")
	}
	return api.Cluster{
		Name: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["name"]),
		Channel: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["channel"]),
		Addons: func(in interface{}) []kops.AddonSpec {
			value := func(in interface{}) []kops.AddonSpec {
				var out []kops.AddonSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.AddonSpec {
						if in == nil {
							return kops.AddonSpec{}
						}
						return (ExpandAddonSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
			return value
		}(in["addons"]),
		ConfigBase: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["config_base"]),
		CloudProvider: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cloud_provider"]),
		ContainerRuntime: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["container_runtime"]),
		KubernetesVersion: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["kubernetes_version"]),
		Subnet: func(in interface{}) []kops.ClusterSubnetSpec {
			value := func(in interface{}) []kops.ClusterSubnetSpec {
				var out []kops.ClusterSubnetSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.ClusterSubnetSpec {
						if in == nil {
							return kops.ClusterSubnetSpec{}
						}
						return (ExpandClusterSubnetSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
			return value
		}(in["subnet"]),
		Project: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["project"]),
		MasterPublicName: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["master_public_name"]),
		MasterInternalName: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["master_internal_name"]),
		NetworkCIDR: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["network_cidr"]),
		AdditionalNetworkCIDRs: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["additional_network_cidrs"]),
		NetworkID: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["network_id"]),
		Topology: func(in interface{}) *kops.TopologySpec {
			value := func(in interface{}) *kops.TopologySpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.TopologySpec) *kops.TopologySpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.TopologySpec {
					if in.([]interface{})[0] == nil {
						return kops.TopologySpec{}
					}
					return (ExpandTopologySpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["topology"]),
		SecretStore: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["secret_store"]),
		KeyStore: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["key_store"]),
		ConfigStore: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["config_store"]),
		DNSZone: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["dns_zone"]),
		AdditionalSANs: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["additional_sans"]),
		ClusterDNSDomain: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cluster_dns_domain"]),
		ServiceClusterIPRange: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["service_cluster_ip_range"]),
		PodCIDR: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["pod_cidr"]),
		NonMasqueradeCIDR: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["non_masquerade_cidr"]),
		SSHAccess: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["ssh_access"]),
		NodePortAccess: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["node_port_access"]),
		EgressProxy: func(in interface{}) *kops.EgressProxySpec {
			value := func(in interface{}) *kops.EgressProxySpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.EgressProxySpec) *kops.EgressProxySpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.EgressProxySpec {
					if in.([]interface{})[0] == nil {
						return kops.EgressProxySpec{}
					}
					return (ExpandEgressProxySpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["egress_proxy"]),
		SSHKeyName: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			return value
		}(in["ssh_key_name"]),
		KubernetesAPIAccess: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["kubernetes_api_access"]),
		IsolateMasters: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			return value
		}(in["isolate_masters"]),
		UpdatePolicy: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			return value
		}(in["update_policy"]),
		ExternalPolicies: func(in interface{}) *map[string][]string {
			value := func(in interface{}) *map[string][]string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in map[string][]string) *map[string][]string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
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
				return tmp
			}(in)
			return value
		}(in["external_policies"]),
		AdditionalPolicies: func(in interface{}) *map[string]string {
			value := func(in interface{}) *map[string]string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in map[string]string) *map[string]string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
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
				return tmp
			}(in)
			return value
		}(in["additional_policies"]),
		FileAssets: func(in interface{}) []kops.FileAssetSpec {
			value := func(in interface{}) []kops.FileAssetSpec {
				var out []kops.FileAssetSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.FileAssetSpec {
						if in == nil {
							return kops.FileAssetSpec{}
						}
						return (ExpandFileAssetSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
			return value
		}(in["file_assets"]),
		EtcdCluster: func(in interface{}) []*kops.EtcdClusterSpec {
			value := func(in interface{}) []*kops.EtcdClusterSpec {
				var out []*kops.EtcdClusterSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *kops.EtcdClusterSpec {
						if in == nil {
							return nil
						}
						if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
							return nil
						}
						tmp := func(in kops.EtcdClusterSpec) *kops.EtcdClusterSpec {
							if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
								return nil
							}
							return &in
						}(func(in interface{}) kops.EtcdClusterSpec {
							if in == nil {
								return kops.EtcdClusterSpec{}
							}
							return (ExpandEtcdClusterSpec(in.(map[string]interface{})))
						}(in))
						return tmp
					}(in))
				}
				return out
			}(in)
			return value
		}(in["etcd_cluster"]),
		Containerd: func(in interface{}) *kops.ContainerdConfig {
			value := func(in interface{}) *kops.ContainerdConfig {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.ContainerdConfig) *kops.ContainerdConfig {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.ContainerdConfig {
					if in.([]interface{})[0] == nil {
						return kops.ContainerdConfig{}
					}
					return (ExpandContainerdConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["containerd"]),
		Docker: func(in interface{}) *kops.DockerConfig {
			value := func(in interface{}) *kops.DockerConfig {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.DockerConfig) *kops.DockerConfig {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.DockerConfig {
					if in.([]interface{})[0] == nil {
						return kops.DockerConfig{}
					}
					return (ExpandDockerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["docker"]),
		KubeDNS: func(in interface{}) *kops.KubeDNSConfig {
			value := func(in interface{}) *kops.KubeDNSConfig {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.KubeDNSConfig) *kops.KubeDNSConfig {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.KubeDNSConfig {
					if in.([]interface{})[0] == nil {
						return kops.KubeDNSConfig{}
					}
					return (ExpandKubeDNSConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["kube_dns"]),
		KubeAPIServer: func(in interface{}) *kops.KubeAPIServerConfig {
			value := func(in interface{}) *kops.KubeAPIServerConfig {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.KubeAPIServerConfig) *kops.KubeAPIServerConfig {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.KubeAPIServerConfig {
					if in.([]interface{})[0] == nil {
						return kops.KubeAPIServerConfig{}
					}
					return (ExpandKubeAPIServerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["kube_api_server"]),
		KubeControllerManager: func(in interface{}) *kops.KubeControllerManagerConfig {
			value := func(in interface{}) *kops.KubeControllerManagerConfig {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.KubeControllerManagerConfig) *kops.KubeControllerManagerConfig {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.KubeControllerManagerConfig {
					if in.([]interface{})[0] == nil {
						return kops.KubeControllerManagerConfig{}
					}
					return (ExpandKubeControllerManagerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["kube_controller_manager"]),
		ExternalCloudControllerManager: func(in interface{}) *kops.CloudControllerManagerConfig {
			value := func(in interface{}) *kops.CloudControllerManagerConfig {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.CloudControllerManagerConfig) *kops.CloudControllerManagerConfig {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.CloudControllerManagerConfig {
					if in.([]interface{})[0] == nil {
						return kops.CloudControllerManagerConfig{}
					}
					return (ExpandCloudControllerManagerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["external_cloud_controller_manager"]),
		KubeScheduler: func(in interface{}) *kops.KubeSchedulerConfig {
			value := func(in interface{}) *kops.KubeSchedulerConfig {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.KubeSchedulerConfig) *kops.KubeSchedulerConfig {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.KubeSchedulerConfig {
					if in.([]interface{})[0] == nil {
						return kops.KubeSchedulerConfig{}
					}
					return (ExpandKubeSchedulerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["kube_scheduler"]),
		KubeProxy: func(in interface{}) *kops.KubeProxyConfig {
			value := func(in interface{}) *kops.KubeProxyConfig {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.KubeProxyConfig) *kops.KubeProxyConfig {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.KubeProxyConfig {
					if in.([]interface{})[0] == nil {
						return kops.KubeProxyConfig{}
					}
					return (ExpandKubeProxyConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["kube_proxy"]),
		Kubelet: func(in interface{}) *kops.KubeletConfigSpec {
			value := func(in interface{}) *kops.KubeletConfigSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.KubeletConfigSpec) *kops.KubeletConfigSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.KubeletConfigSpec {
					if in.([]interface{})[0] == nil {
						return kops.KubeletConfigSpec{}
					}
					return (ExpandKubeletConfigSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["kubelet"]),
		MasterKubelet: func(in interface{}) *kops.KubeletConfigSpec {
			value := func(in interface{}) *kops.KubeletConfigSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.KubeletConfigSpec) *kops.KubeletConfigSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.KubeletConfigSpec {
					if in.([]interface{})[0] == nil {
						return kops.KubeletConfigSpec{}
					}
					return (ExpandKubeletConfigSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["master_kubelet"]),
		CloudConfig: func(in interface{}) *kops.CloudConfiguration {
			value := func(in interface{}) *kops.CloudConfiguration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.CloudConfiguration) *kops.CloudConfiguration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.CloudConfiguration {
					if in.([]interface{})[0] == nil {
						return kops.CloudConfiguration{}
					}
					return (ExpandCloudConfiguration(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["cloud_config"]),
		ExternalDNS: func(in interface{}) *kops.ExternalDNSConfig {
			value := func(in interface{}) *kops.ExternalDNSConfig {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.ExternalDNSConfig) *kops.ExternalDNSConfig {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.ExternalDNSConfig {
					if in.([]interface{})[0] == nil {
						return kops.ExternalDNSConfig{}
					}
					return (ExpandExternalDNSConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["external_dns"]),
		Networking: func(in interface{}) *kops.NetworkingSpec {
			value := func(in interface{}) *kops.NetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.NetworkingSpec) *kops.NetworkingSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.NetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.NetworkingSpec{}
					}
					return (ExpandNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["networking"]),
		API: func(in interface{}) *kops.AccessSpec {
			value := func(in interface{}) *kops.AccessSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.AccessSpec) *kops.AccessSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.AccessSpec {
					if in.([]interface{})[0] == nil {
						return kops.AccessSpec{}
					}
					return (ExpandAccessSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["api"]),
		Authentication: func(in interface{}) *kops.AuthenticationSpec {
			value := func(in interface{}) *kops.AuthenticationSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.AuthenticationSpec) *kops.AuthenticationSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.AuthenticationSpec {
					if in.([]interface{})[0] == nil {
						return kops.AuthenticationSpec{}
					}
					return (ExpandAuthenticationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["authentication"]),
		Authorization: func(in interface{}) *kops.AuthorizationSpec {
			value := func(in interface{}) *kops.AuthorizationSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.AuthorizationSpec) *kops.AuthorizationSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.AuthorizationSpec {
					if in.([]interface{})[0] == nil {
						return kops.AuthorizationSpec{}
					}
					return (ExpandAuthorizationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["authorization"]),
		NodeAuthorization: func(in interface{}) *kops.NodeAuthorizationSpec {
			value := func(in interface{}) *kops.NodeAuthorizationSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.NodeAuthorizationSpec) *kops.NodeAuthorizationSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.NodeAuthorizationSpec {
					if in.([]interface{})[0] == nil {
						return kops.NodeAuthorizationSpec{}
					}
					return (ExpandNodeAuthorizationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["node_authorization"]),
		CloudLabels: func(in interface{}) map[string]string {
			value := func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
			return value
		}(in["cloud_labels"]),
		Hooks: func(in interface{}) []kops.HookSpec {
			value := func(in interface{}) []kops.HookSpec {
				var out []kops.HookSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.HookSpec {
						if in == nil {
							return kops.HookSpec{}
						}
						return (ExpandHookSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
			return value
		}(in["hooks"]),
		Assets: func(in interface{}) *kops.Assets {
			value := func(in interface{}) *kops.Assets {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.Assets) *kops.Assets {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.Assets {
					if in.([]interface{})[0] == nil {
						return kops.Assets{}
					}
					return (ExpandAssets(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["assets"]),
		IAM: func(in interface{}) *kops.IAMSpec {
			value := func(in interface{}) *kops.IAMSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.IAMSpec) *kops.IAMSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.IAMSpec {
					if in.([]interface{})[0] == nil {
						return kops.IAMSpec{}
					}
					return (ExpandIAMSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["iam"]),
		EncryptionConfig: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			return value
		}(in["encryption_config"]),
		DisableSubnetTags: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["disable_subnet_tags"]),
		UseHostCertificates: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			return value
		}(in["use_host_certificates"]),
		SysctlParameters: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["sysctl_parameters"]),
		RollingUpdate: func(in interface{}) *kops.RollingUpdate {
			value := func(in interface{}) *kops.RollingUpdate {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.RollingUpdate) *kops.RollingUpdate {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.RollingUpdate {
					if in.([]interface{})[0] == nil {
						return kops.RollingUpdate{}
					}
					return (ExpandRollingUpdate(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["rolling_update"]),
		InstanceGroup: func(in interface{}) []*api.InstanceGroup {
			value := func(in interface{}) []*api.InstanceGroup {
				var out []*api.InstanceGroup
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *api.InstanceGroup {
						if in == nil {
							return nil
						}
						if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
							return nil
						}
						tmp := func(in api.InstanceGroup) *api.InstanceGroup {
							if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
								return nil
							}
							return &in
						}(func(in interface{}) api.InstanceGroup {
							if in == nil {
								return api.InstanceGroup{}
							}
							return (ExpandInstanceGroup(in.(map[string]interface{})))
						}(in))
						return tmp
					}(in))
				}
				return out
			}(in)
			return value
		}(in["instance_group"]),
	}
}

func FlattenCluster(in api.Cluster) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Name),
		"channel": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Channel),
		"addons": func(in []kops.AddonSpec) interface{} {
			value := func(in []kops.AddonSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.AddonSpec) interface{} {
						return FlattenAddonSpec(in)
					}(in))
				}
				return out
			}(in)
			return value
		}(in.Addons),
		"config_base": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ConfigBase),
		"cloud_provider": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.CloudProvider),
		"container_runtime": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ContainerRuntime),
		"kubernetes_version": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.KubernetesVersion),
		"subnet": func(in []kops.ClusterSubnetSpec) interface{} {
			value := func(in []kops.ClusterSubnetSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.ClusterSubnetSpec) interface{} {
						return FlattenClusterSubnetSpec(in)
					}(in))
				}
				return out
			}(in)
			return value
		}(in.Subnet),
		"project": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Project),
		"master_public_name": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.MasterPublicName),
		"master_internal_name": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.MasterInternalName),
		"network_cidr": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.NetworkCIDR),
		"additional_network_cidrs": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.AdditionalNetworkCIDRs),
		"network_id": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.NetworkID),
		"topology": func(in *kops.TopologySpec) interface{} {
			value := func(in *kops.TopologySpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.TopologySpec) interface{} {
					return func(in kops.TopologySpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenTopologySpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Topology),
		"secret_store": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.SecretStore),
		"key_store": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.KeyStore),
		"config_store": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ConfigStore),
		"dns_zone": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.DNSZone),
		"additional_sans": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.AdditionalSANs),
		"cluster_dns_domain": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ClusterDNSDomain),
		"service_cluster_ip_range": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ServiceClusterIPRange),
		"pod_cidr": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.PodCIDR),
		"non_masquerade_cidr": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.NonMasqueradeCIDR),
		"ssh_access": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.SSHAccess),
		"node_port_access": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.NodePortAccess),
		"egress_proxy": func(in *kops.EgressProxySpec) interface{} {
			value := func(in *kops.EgressProxySpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.EgressProxySpec) interface{} {
					return func(in kops.EgressProxySpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenEgressProxySpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.EgressProxy),
		"ssh_key_name": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.SSHKeyName),
		"kubernetes_api_access": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.KubernetesAPIAccess),
		"isolate_masters": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.IsolateMasters),
		"update_policy": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.UpdatePolicy),
		"external_policies": func(in *map[string][]string) interface{} {
			value := func(in *map[string][]string) interface{} {
				if in == nil {
					return nil
				}
				return func(in map[string][]string) interface{} {
					return func(in map[string][]string) map[string]interface{} {
						if in == nil {
							return nil
						}
						// TODO
						return nil
					}(in)
				}(*in)
			}(in)
			return value
		}(in.ExternalPolicies),
		"additional_policies": func(in *map[string]string) interface{} {
			value := func(in *map[string]string) interface{} {
				if in == nil {
					return nil
				}
				return func(in map[string]string) interface{} {
					return func(in map[string]string) map[string]interface{} {
						if in == nil {
							return nil
						}
						// TODO
						return nil
					}(in)
				}(*in)
			}(in)
			return value
		}(in.AdditionalPolicies),
		"file_assets": func(in []kops.FileAssetSpec) interface{} {
			value := func(in []kops.FileAssetSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.FileAssetSpec) interface{} {
						return FlattenFileAssetSpec(in)
					}(in))
				}
				return out
			}(in)
			return value
		}(in.FileAssets),
		"etcd_cluster": func(in []*kops.EtcdClusterSpec) interface{} {
			value := func(in []*kops.EtcdClusterSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in *kops.EtcdClusterSpec) interface{} {
						if in == nil {
							return nil
						}
						return func(in kops.EtcdClusterSpec) interface{} {
							return func(in kops.EtcdClusterSpec) interface{} {
								return FlattenEtcdClusterSpec(in)
							}(in)
						}(*in)
					}(in))
				}
				return out
			}(in)
			return value
		}(in.EtcdCluster),
		"containerd": func(in *kops.ContainerdConfig) interface{} {
			value := func(in *kops.ContainerdConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.ContainerdConfig) interface{} {
					return func(in kops.ContainerdConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenContainerdConfig(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Containerd),
		"docker": func(in *kops.DockerConfig) interface{} {
			value := func(in *kops.DockerConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.DockerConfig) interface{} {
					return func(in kops.DockerConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenDockerConfig(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Docker),
		"kube_dns": func(in *kops.KubeDNSConfig) interface{} {
			value := func(in *kops.KubeDNSConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KubeDNSConfig) interface{} {
					return func(in kops.KubeDNSConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenKubeDNSConfig(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.KubeDNS),
		"kube_api_server": func(in *kops.KubeAPIServerConfig) interface{} {
			value := func(in *kops.KubeAPIServerConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KubeAPIServerConfig) interface{} {
					return func(in kops.KubeAPIServerConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenKubeAPIServerConfig(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.KubeAPIServer),
		"kube_controller_manager": func(in *kops.KubeControllerManagerConfig) interface{} {
			value := func(in *kops.KubeControllerManagerConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KubeControllerManagerConfig) interface{} {
					return func(in kops.KubeControllerManagerConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenKubeControllerManagerConfig(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.KubeControllerManager),
		"external_cloud_controller_manager": func(in *kops.CloudControllerManagerConfig) interface{} {
			value := func(in *kops.CloudControllerManagerConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.CloudControllerManagerConfig) interface{} {
					return func(in kops.CloudControllerManagerConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenCloudControllerManagerConfig(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.ExternalCloudControllerManager),
		"kube_scheduler": func(in *kops.KubeSchedulerConfig) interface{} {
			value := func(in *kops.KubeSchedulerConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KubeSchedulerConfig) interface{} {
					return func(in kops.KubeSchedulerConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenKubeSchedulerConfig(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.KubeScheduler),
		"kube_proxy": func(in *kops.KubeProxyConfig) interface{} {
			value := func(in *kops.KubeProxyConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KubeProxyConfig) interface{} {
					return func(in kops.KubeProxyConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenKubeProxyConfig(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.KubeProxy),
		"kubelet": func(in *kops.KubeletConfigSpec) interface{} {
			value := func(in *kops.KubeletConfigSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KubeletConfigSpec) interface{} {
					return func(in kops.KubeletConfigSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenKubeletConfigSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Kubelet),
		"master_kubelet": func(in *kops.KubeletConfigSpec) interface{} {
			value := func(in *kops.KubeletConfigSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KubeletConfigSpec) interface{} {
					return func(in kops.KubeletConfigSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenKubeletConfigSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.MasterKubelet),
		"cloud_config": func(in *kops.CloudConfiguration) interface{} {
			value := func(in *kops.CloudConfiguration) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.CloudConfiguration) interface{} {
					return func(in kops.CloudConfiguration) []map[string]interface{} {
						return []map[string]interface{}{FlattenCloudConfiguration(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.CloudConfig),
		"external_dns": func(in *kops.ExternalDNSConfig) interface{} {
			value := func(in *kops.ExternalDNSConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.ExternalDNSConfig) interface{} {
					return func(in kops.ExternalDNSConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenExternalDNSConfig(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.ExternalDNS),
		"networking": func(in *kops.NetworkingSpec) interface{} {
			value := func(in *kops.NetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.NetworkingSpec) interface{} {
					return func(in kops.NetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Networking),
		"api": func(in *kops.AccessSpec) interface{} {
			value := func(in *kops.AccessSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.AccessSpec) interface{} {
					return func(in kops.AccessSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenAccessSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.API),
		"authentication": func(in *kops.AuthenticationSpec) interface{} {
			value := func(in *kops.AuthenticationSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.AuthenticationSpec) interface{} {
					return func(in kops.AuthenticationSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenAuthenticationSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Authentication),
		"authorization": func(in *kops.AuthorizationSpec) interface{} {
			value := func(in *kops.AuthorizationSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.AuthorizationSpec) interface{} {
					return func(in kops.AuthorizationSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenAuthorizationSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Authorization),
		"node_authorization": func(in *kops.NodeAuthorizationSpec) interface{} {
			value := func(in *kops.NodeAuthorizationSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.NodeAuthorizationSpec) interface{} {
					return func(in kops.NodeAuthorizationSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenNodeAuthorizationSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.NodeAuthorization),
		"cloud_labels": func(in map[string]string) interface{} {
			value := func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
			return value
		}(in.CloudLabels),
		"hooks": func(in []kops.HookSpec) interface{} {
			value := func(in []kops.HookSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.HookSpec) interface{} {
						return FlattenHookSpec(in)
					}(in))
				}
				return out
			}(in)
			return value
		}(in.Hooks),
		"assets": func(in *kops.Assets) interface{} {
			value := func(in *kops.Assets) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.Assets) interface{} {
					return func(in kops.Assets) []map[string]interface{} {
						return []map[string]interface{}{FlattenAssets(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Assets),
		"iam": func(in *kops.IAMSpec) interface{} {
			value := func(in *kops.IAMSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.IAMSpec) interface{} {
					return func(in kops.IAMSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenIAMSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.IAM),
		"encryption_config": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.EncryptionConfig),
		"disable_subnet_tags": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.DisableSubnetTags),
		"use_host_certificates": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.UseHostCertificates),
		"sysctl_parameters": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.SysctlParameters),
		"rolling_update": func(in *kops.RollingUpdate) interface{} {
			value := func(in *kops.RollingUpdate) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.RollingUpdate) interface{} {
					return func(in kops.RollingUpdate) []map[string]interface{} {
						return []map[string]interface{}{FlattenRollingUpdate(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.RollingUpdate),
		"instance_group": func(in []*api.InstanceGroup) interface{} {
			value := func(in []*api.InstanceGroup) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in *api.InstanceGroup) interface{} {
						if in == nil {
							return nil
						}
						return func(in api.InstanceGroup) interface{} {
							return func(in api.InstanceGroup) interface{} {
								return FlattenInstanceGroup(in)
							}(in)
						}(*in)
					}(in))
				}
				return out
			}(in)
			return value
		}(in.InstanceGroup),
	}
}
