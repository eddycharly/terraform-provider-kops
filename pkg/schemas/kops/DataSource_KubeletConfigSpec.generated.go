package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceKubeletConfigSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_servers":                            ComputedString(),
			"anonymous_auth":                         ComputedBool(),
			"authorization_mode":                     ComputedString(),
			"bootstrap_kubeconfig":                   ComputedString(),
			"client_ca_file":                         ComputedString(),
			"tls_cert_file":                          ComputedString(),
			"tls_private_key_file":                   ComputedString(),
			"tls_cipher_suites":                      ComputedList(String()),
			"tls_min_version":                        ComputedString(),
			"kubeconfig_path":                        ComputedString(),
			"require_kubeconfig":                     ComputedBool(),
			"log_level":                              ComputedInt(),
			"pod_manifest_path":                      ComputedString(),
			"hostname_override":                      ComputedString(),
			"pod_infra_container_image":              ComputedString(),
			"seccomp_profile_root":                   ComputedString(),
			"allow_privileged":                       ComputedBool(),
			"enable_debugging_handlers":              ComputedBool(),
			"register_node":                          ComputedBool(),
			"node_status_update_frequency":           ComputedDuration(),
			"cluster_domain":                         ComputedString(),
			"cluster_dns":                            ComputedString(),
			"network_plugin_name":                    ComputedString(),
			"cloud_provider":                         ComputedString(),
			"kubelet_cgroups":                        ComputedString(),
			"runtime_cgroups":                        ComputedString(),
			"read_only_port":                         ComputedInt(),
			"system_cgroups":                         ComputedString(),
			"cgroup_root":                            ComputedString(),
			"configure_cbr0":                         ComputedBool(),
			"hairpin_mode":                           ComputedString(),
			"babysit_daemons":                        ComputedBool(),
			"max_pods":                               ComputedInt(),
			"nvidia_gp_us":                           ComputedInt(),
			"pod_cidr":                               ComputedString(),
			"resolver_config":                        ComputedString(),
			"reconcile_cidr":                         ComputedBool(),
			"register_schedulable":                   ComputedBool(),
			"serialize_image_pulls":                  ComputedBool(),
			"node_labels":                            ComputedMap(String()),
			"non_masquerade_cidr":                    ComputedString(),
			"enable_custom_metrics":                  ComputedBool(),
			"network_plugin_mtu":                     ComputedInt(),
			"image_gc_high_threshold_percent":        ComputedInt(),
			"image_gc_low_threshold_percent":         ComputedInt(),
			"image_pull_progress_deadline":           ComputedDuration(),
			"eviction_hard":                          ComputedString(),
			"eviction_soft":                          ComputedString(),
			"eviction_soft_grace_period":             ComputedString(),
			"eviction_pressure_transition_period":    ComputedDuration(),
			"eviction_max_pod_grace_period":          ComputedInt(),
			"eviction_minimum_reclaim":               ComputedString(),
			"volume_plugin_directory":                ComputedString(),
			"taints":                                 ComputedList(String()),
			"feature_gates":                          ComputedMap(String()),
			"kube_reserved":                          ComputedMap(String()),
			"kube_reserved_cgroup":                   ComputedString(),
			"system_reserved":                        ComputedMap(String()),
			"system_reserved_cgroup":                 ComputedString(),
			"enforce_node_allocatable":               ComputedString(),
			"runtime_request_timeout":                ComputedDuration(),
			"volume_stats_agg_period":                ComputedDuration(),
			"fail_swap_on":                           ComputedBool(),
			"experimental_allowed_unsafe_sysctls":    ComputedList(String()),
			"allowed_unsafe_sysctls":                 ComputedList(String()),
			"streaming_connection_idle_timeout":      ComputedDuration(),
			"docker_disable_shared_pid":              ComputedBool(),
			"root_dir":                               ComputedString(),
			"authentication_token_webhook":           ComputedBool(),
			"authentication_token_webhook_cache_ttl": ComputedDuration(),
			"cpucfs_quota":                           ComputedBool(),
			"cpucfs_quota_period":                    ComputedDuration(),
			"cpu_manager_policy":                     ComputedString(),
			"registry_pull_qps":                      ComputedInt(),
			"registry_burst":                         ComputedInt(),
			"topology_manager_policy":                ComputedString(),
			"rotate_certificates":                    ComputedBool(),
			"protect_kernel_defaults":                ComputedBool(),
			"cgroup_driver":                          ComputedString(),
			"housekeeping_interval":                  ComputedDuration(),
			"event_qps":                              ComputedInt(),
			"event_burst":                            ComputedInt(),
			"container_log_max_size":                 ComputedString(),
			"container_log_max_files":                ComputedInt(),
			"enable_cadvisor_json_endpoints":         ComputedBool(),
		},
	}
}

func ExpandDataSourceKubeletConfigSpec(in map[string]interface{}) kops.KubeletConfigSpec {
	if in == nil {
		panic("expand KubeletConfigSpec failure, in is nil")
	}
	return kops.KubeletConfigSpec{
		APIServers: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["api_servers"]),
		AnonymousAuth: func(in interface{}) *bool {
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
		}(in["anonymous_auth"]),
		AuthorizationMode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["authorization_mode"]),
		BootstrapKubeconfig: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["bootstrap_kubeconfig"]),
		ClientCAFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["client_ca_file"]),
		TLSCertFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["tls_cert_file"]),
		TLSPrivateKeyFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["tls_private_key_file"]),
		TLSCipherSuites: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["tls_cipher_suites"]),
		TLSMinVersion: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["tls_min_version"]),
		KubeconfigPath: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kubeconfig_path"]),
		RequireKubeconfig: func(in interface{}) *bool {
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
		}(in["require_kubeconfig"]),
		LogLevel: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["log_level"]),
		PodManifestPath: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["pod_manifest_path"]),
		HostnameOverride: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["hostname_override"]),
		PodInfraContainerImage: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["pod_infra_container_image"]),
		SeccompProfileRoot: func(in interface{}) *string {
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
		}(in["seccomp_profile_root"]),
		AllowPrivileged: func(in interface{}) *bool {
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
		}(in["allow_privileged"]),
		EnableDebuggingHandlers: func(in interface{}) *bool {
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
		}(in["enable_debugging_handlers"]),
		RegisterNode: func(in interface{}) *bool {
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
		}(in["register_node"]),
		NodeStatusUpdateFrequency: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["node_status_update_frequency"]),
		ClusterDomain: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_domain"]),
		ClusterDNS: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_dns"]),
		NetworkPluginName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["network_plugin_name"]),
		CloudProvider: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cloud_provider"]),
		KubeletCgroups: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kubelet_cgroups"]),
		RuntimeCgroups: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["runtime_cgroups"]),
		ReadOnlyPort: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["read_only_port"]),
		SystemCgroups: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["system_cgroups"]),
		CgroupRoot: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cgroup_root"]),
		ConfigureCBR0: func(in interface{}) *bool {
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
		}(in["configure_cbr0"]),
		HairpinMode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["hairpin_mode"]),
		BabysitDaemons: func(in interface{}) *bool {
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
		}(in["babysit_daemons"]),
		MaxPods: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["max_pods"]),
		NvidiaGPUs: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["nvidia_gp_us"]),
		PodCIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["pod_cidr"]),
		ResolverConfig: func(in interface{}) *string {
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
		}(in["resolver_config"]),
		ReconcileCIDR: func(in interface{}) *bool {
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
		}(in["reconcile_cidr"]),
		RegisterSchedulable: func(in interface{}) *bool {
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
		}(in["register_schedulable"]),
		SerializeImagePulls: func(in interface{}) *bool {
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
		}(in["serialize_image_pulls"]),
		NodeLabels: func(in interface{}) map[string]string {
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
		}(in["node_labels"]),
		NonMasqueradeCIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["non_masquerade_cidr"]),
		EnableCustomMetrics: func(in interface{}) *bool {
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
		}(in["enable_custom_metrics"]),
		NetworkPluginMTU: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["network_plugin_mtu"]),
		ImageGCHighThresholdPercent: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["image_gc_high_threshold_percent"]),
		ImageGCLowThresholdPercent: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["image_gc_low_threshold_percent"]),
		ImagePullProgressDeadline: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["image_pull_progress_deadline"]),
		EvictionHard: func(in interface{}) *string {
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
		}(in["eviction_hard"]),
		EvictionSoft: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["eviction_soft"]),
		EvictionSoftGracePeriod: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["eviction_soft_grace_period"]),
		EvictionPressureTransitionPeriod: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["eviction_pressure_transition_period"]),
		EvictionMaxPodGracePeriod: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["eviction_max_pod_grace_period"]),
		EvictionMinimumReclaim: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["eviction_minimum_reclaim"]),
		VolumePluginDirectory: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["volume_plugin_directory"]),
		Taints: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["taints"]),
		FeatureGates: func(in interface{}) map[string]string {
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
		}(in["feature_gates"]),
		KubeReserved: func(in interface{}) map[string]string {
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
		}(in["kube_reserved"]),
		KubeReservedCgroup: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kube_reserved_cgroup"]),
		SystemReserved: func(in interface{}) map[string]string {
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
		}(in["system_reserved"]),
		SystemReservedCgroup: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["system_reserved_cgroup"]),
		EnforceNodeAllocatable: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["enforce_node_allocatable"]),
		RuntimeRequestTimeout: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["runtime_request_timeout"]),
		VolumeStatsAggPeriod: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["volume_stats_agg_period"]),
		FailSwapOn: func(in interface{}) *bool {
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
		}(in["fail_swap_on"]),
		ExperimentalAllowedUnsafeSysctls: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["experimental_allowed_unsafe_sysctls"]),
		AllowedUnsafeSysctls: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["allowed_unsafe_sysctls"]),
		StreamingConnectionIdleTimeout: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["streaming_connection_idle_timeout"]),
		DockerDisableSharedPID: func(in interface{}) *bool {
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
		}(in["docker_disable_shared_pid"]),
		RootDir: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["root_dir"]),
		AuthenticationTokenWebhook: func(in interface{}) *bool {
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
		}(in["authentication_token_webhook"]),
		AuthenticationTokenWebhookCacheTTL: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["authentication_token_webhook_cache_ttl"]),
		CPUCFSQuota: func(in interface{}) *bool {
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
		}(in["cpucfs_quota"]),
		CPUCFSQuotaPeriod: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["cpucfs_quota_period"]),
		CpuManagerPolicy: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cpu_manager_policy"]),
		RegistryPullQPS: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["registry_pull_qps"]),
		RegistryBurst: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["registry_burst"]),
		TopologyManagerPolicy: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["topology_manager_policy"]),
		RotateCertificates: func(in interface{}) *bool {
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
		}(in["rotate_certificates"]),
		ProtectKernelDefaults: func(in interface{}) *bool {
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
		}(in["protect_kernel_defaults"]),
		CgroupDriver: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cgroup_driver"]),
		HousekeepingInterval: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["housekeeping_interval"]),
		EventQPS: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["event_qps"]),
		EventBurst: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["event_burst"]),
		ContainerLogMaxSize: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["container_log_max_size"]),
		ContainerLogMaxFiles: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["container_log_max_files"]),
		EnableCadvisorJsonEndpoints: func(in interface{}) *bool {
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
		}(in["enable_cadvisor_json_endpoints"]),
	}
}

func FlattenDataSourceKubeletConfigSpecInto(in kops.KubeletConfigSpec, out map[string]interface{}) {
	out["api_servers"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.APIServers)
	out["anonymous_auth"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.AnonymousAuth)
	out["authorization_mode"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AuthorizationMode)
	out["bootstrap_kubeconfig"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.BootstrapKubeconfig)
	out["client_ca_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClientCAFile)
	out["tls_cert_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.TLSCertFile)
	out["tls_private_key_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.TLSPrivateKeyFile)
	out["tls_cipher_suites"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.TLSCipherSuites)
	out["tls_min_version"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.TLSMinVersion)
	out["kubeconfig_path"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.KubeconfigPath)
	out["require_kubeconfig"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.RequireKubeconfig)
	out["log_level"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.LogLevel)
	out["pod_manifest_path"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.PodManifestPath)
	out["hostname_override"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.HostnameOverride)
	out["pod_infra_container_image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.PodInfraContainerImage)
	out["seccomp_profile_root"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.SeccompProfileRoot)
	out["allow_privileged"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.AllowPrivileged)
	out["enable_debugging_handlers"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EnableDebuggingHandlers)
	out["register_node"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.RegisterNode)
	out["node_status_update_frequency"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.NodeStatusUpdateFrequency)
	out["cluster_domain"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterDomain)
	out["cluster_dns"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterDNS)
	out["network_plugin_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.NetworkPluginName)
	out["cloud_provider"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CloudProvider)
	out["kubelet_cgroups"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.KubeletCgroups)
	out["runtime_cgroups"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.RuntimeCgroups)
	out["read_only_port"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.ReadOnlyPort)
	out["system_cgroups"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.SystemCgroups)
	out["cgroup_root"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CgroupRoot)
	out["configure_cbr0"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.ConfigureCBR0)
	out["hairpin_mode"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.HairpinMode)
	out["babysit_daemons"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.BabysitDaemons)
	out["max_pods"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.MaxPods)
	out["nvidia_gp_us"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.NvidiaGPUs)
	out["pod_cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.PodCIDR)
	out["resolver_config"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.ResolverConfig)
	out["reconcile_cidr"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.ReconcileCIDR)
	out["register_schedulable"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.RegisterSchedulable)
	out["serialize_image_pulls"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.SerializeImagePulls)
	out["node_labels"] = func(in map[string]string) interface{} {
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
	}(in.NodeLabels)
	out["non_masquerade_cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.NonMasqueradeCIDR)
	out["enable_custom_metrics"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EnableCustomMetrics)
	out["network_plugin_mtu"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.NetworkPluginMTU)
	out["image_gc_high_threshold_percent"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.ImageGCHighThresholdPercent)
	out["image_gc_low_threshold_percent"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.ImageGCLowThresholdPercent)
	out["image_pull_progress_deadline"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.ImagePullProgressDeadline)
	out["eviction_hard"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.EvictionHard)
	out["eviction_soft"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.EvictionSoft)
	out["eviction_soft_grace_period"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.EvictionSoftGracePeriod)
	out["eviction_pressure_transition_period"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.EvictionPressureTransitionPeriod)
	out["eviction_max_pod_grace_period"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.EvictionMaxPodGracePeriod)
	out["eviction_minimum_reclaim"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.EvictionMinimumReclaim)
	out["volume_plugin_directory"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.VolumePluginDirectory)
	out["taints"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Taints)
	out["feature_gates"] = func(in map[string]string) interface{} {
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
	}(in.FeatureGates)
	out["kube_reserved"] = func(in map[string]string) interface{} {
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
	}(in.KubeReserved)
	out["kube_reserved_cgroup"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.KubeReservedCgroup)
	out["system_reserved"] = func(in map[string]string) interface{} {
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
	}(in.SystemReserved)
	out["system_reserved_cgroup"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.SystemReservedCgroup)
	out["enforce_node_allocatable"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.EnforceNodeAllocatable)
	out["runtime_request_timeout"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.RuntimeRequestTimeout)
	out["volume_stats_agg_period"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.VolumeStatsAggPeriod)
	out["fail_swap_on"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.FailSwapOn)
	out["experimental_allowed_unsafe_sysctls"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.ExperimentalAllowedUnsafeSysctls)
	out["allowed_unsafe_sysctls"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AllowedUnsafeSysctls)
	out["streaming_connection_idle_timeout"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.StreamingConnectionIdleTimeout)
	out["docker_disable_shared_pid"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.DockerDisableSharedPID)
	out["root_dir"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.RootDir)
	out["authentication_token_webhook"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.AuthenticationTokenWebhook)
	out["authentication_token_webhook_cache_ttl"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.AuthenticationTokenWebhookCacheTTL)
	out["cpucfs_quota"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.CPUCFSQuota)
	out["cpucfs_quota_period"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.CPUCFSQuotaPeriod)
	out["cpu_manager_policy"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CpuManagerPolicy)
	out["registry_pull_qps"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.RegistryPullQPS)
	out["registry_burst"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.RegistryBurst)
	out["topology_manager_policy"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.TopologyManagerPolicy)
	out["rotate_certificates"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.RotateCertificates)
	out["protect_kernel_defaults"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.ProtectKernelDefaults)
	out["cgroup_driver"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CgroupDriver)
	out["housekeeping_interval"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.HousekeepingInterval)
	out["event_qps"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.EventQPS)
	out["event_burst"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.EventBurst)
	out["container_log_max_size"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ContainerLogMaxSize)
	out["container_log_max_files"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.ContainerLogMaxFiles)
	out["enable_cadvisor_json_endpoints"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EnableCadvisorJsonEndpoints)
}

func FlattenDataSourceKubeletConfigSpec(in kops.KubeletConfigSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKubeletConfigSpecInto(in, out)
	return out
}
