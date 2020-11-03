package structures

import (
	"reflect"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandKubeletConfigSpec(in map[string]interface{}) kops.KubeletConfigSpec {
	if in == nil {
		panic("expand KubeletConfigSpec failure, in is nil")
	}
	return kops.KubeletConfigSpec{
		APIServers: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["api_servers"]),
		AnonymousAuth: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["anonymous_auth"]),
		AuthorizationMode: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["authorization_mode"]),
		BootstrapKubeconfig: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["bootstrap_kubeconfig"]),
		ClientCAFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["client_ca_file"]),
		TLSCertFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["tls_cert_file"]),
		TLSPrivateKeyFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["tls_private_key_file"]),
		TLSCipherSuites: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["tls_cipher_suites"]),
		TLSMinVersion: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["tls_min_version"]),
		KubeconfigPath: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["kubeconfig_path"]),
		RequireKubeconfig: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["require_kubeconfig"]),
		LogLevel: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["log_level"]),
		PodManifestPath: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["pod_manifest_path"]),
		HostnameOverride: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["hostname_override"]),
		PodInfraContainerImage: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["pod_infra_container_image"]),
		SeccompProfileRoot: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["seccomp_profile_root"]),
		AllowPrivileged: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["allow_privileged"]),
		EnableDebuggingHandlers: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["enable_debugging_handlers"]),
		RegisterNode: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["register_node"]),
		NodeStatusUpdateFrequency: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["node_status_update_frequency"]),
		ClusterDomain: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cluster_domain"]),
		ClusterDNS: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cluster_dns"]),
		NetworkPluginName: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["network_plugin_name"]),
		CloudProvider: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cloud_provider"]),
		KubeletCgroups: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["kubelet_cgroups"]),
		RuntimeCgroups: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["runtime_cgroups"]),
		ReadOnlyPort: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["read_only_port"]),
		SystemCgroups: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["system_cgroups"]),
		CgroupRoot: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cgroup_root"]),
		ConfigureCBR0: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["configure_cbr_00"]),
		HairpinMode: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["hairpin_mode"]),
		BabysitDaemons: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["babysit_daemons"]),
		MaxPods: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["max_pods"]),
		NvidiaGPUs: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			return value
		}(in["nvidia_gp_uss"]),
		PodCIDR: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["pod_cidr"]),
		ResolverConfig: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["resolver_config"]),
		ReconcileCIDR: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["reconcile_cidr"]),
		RegisterSchedulable: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["register_schedulable"]),
		SerializeImagePulls: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["serialize_image_pulls"]),
		NodeLabels: func(in interface{}) map[string]string {
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
		}(in["node_labels"]),
		NonMasqueradeCIDR: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["non_masquerade_cidr"]),
		EnableCustomMetrics: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["enable_custom_metrics"]),
		NetworkPluginMTU: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["network_plugin_mtu"]),
		ImageGCHighThresholdPercent: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["image_gc_high_threshold_percent"]),
		ImageGCLowThresholdPercent: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["image_gc_low_threshold_percent"]),
		ImagePullProgressDeadline: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["image_pull_progress_deadline"]),
		EvictionHard: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["eviction_hard"]),
		EvictionSoft: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["eviction_soft"]),
		EvictionSoftGracePeriod: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["eviction_soft_grace_period"]),
		EvictionPressureTransitionPeriod: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["eviction_pressure_transition_period"]),
		EvictionMaxPodGracePeriod: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			return value
		}(in["eviction_max_pod_grace_period"]),
		EvictionMinimumReclaim: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["eviction_minimum_reclaim"]),
		VolumePluginDirectory: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["volume_plugin_directory"]),
		Taints: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["taints"]),
		FeatureGates: func(in interface{}) map[string]string {
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
		}(in["feature_gates"]),
		KubeReserved: func(in interface{}) map[string]string {
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
		}(in["kube_reserved"]),
		KubeReservedCgroup: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["kube_reserved_cgroup"]),
		SystemReserved: func(in interface{}) map[string]string {
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
		}(in["system_reserved"]),
		SystemReservedCgroup: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["system_reserved_cgroup"]),
		EnforceNodeAllocatable: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["enforce_node_allocatable"]),
		RuntimeRequestTimeout: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["runtime_request_timeout"]),
		VolumeStatsAggPeriod: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["volume_stats_agg_period"]),
		FailSwapOn: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["fail_swap_on"]),
		ExperimentalAllowedUnsafeSysctls: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["experimental_allowed_unsafe_sysctls"]),
		AllowedUnsafeSysctls: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["allowed_unsafe_sysctls"]),
		StreamingConnectionIdleTimeout: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["streaming_connection_idle_timeout"]),
		DockerDisableSharedPID: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["docker_disable_shared_pid"]),
		RootDir: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["root_dir"]),
		AuthenticationTokenWebhook: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["authentication_token_webhook"]),
		AuthenticationTokenWebhookCacheTTL: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["authentication_token_webhook_cache_ttl"]),
		CPUCFSQuota: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["cpucfs_quota"]),
		CPUCFSQuotaPeriod: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["cpucfs_quota_period"]),
		CpuManagerPolicy: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cpu_manager_policy"]),
		RegistryPullQPS: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["registry_pull_qps"]),
		RegistryBurst: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["registry_burst"]),
		TopologyManagerPolicy: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["topology_manager_policy"]),
		RotateCertificates: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["rotate_certificates"]),
		ProtectKernelDefaults: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["protect_kernel_defaults"]),
		CgroupDriver: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cgroup_driver"]),
	}
}

func FlattenKubeletConfigSpec(in kops.KubeletConfigSpec) map[string]interface{} {
	return map[string]interface{}{
		"api_servers": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.APIServers),
		"anonymous_auth": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.AnonymousAuth),
		"authorization_mode": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.AuthorizationMode),
		"bootstrap_kubeconfig": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.BootstrapKubeconfig),
		"client_ca_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ClientCAFile),
		"tls_cert_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.TLSCertFile),
		"tls_private_key_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.TLSPrivateKeyFile),
		"tls_cipher_suites": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.TLSCipherSuites),
		"tls_min_version": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.TLSMinVersion),
		"kubeconfig_path": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.KubeconfigPath),
		"require_kubeconfig": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.RequireKubeconfig),
		"log_level": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.LogLevel),
		"pod_manifest_path": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.PodManifestPath),
		"hostname_override": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.HostnameOverride),
		"pod_infra_container_image": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.PodInfraContainerImage),
		"seccomp_profile_root": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.SeccompProfileRoot),
		"allow_privileged": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.AllowPrivileged),
		"enable_debugging_handlers": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.EnableDebuggingHandlers),
		"register_node": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.RegisterNode),
		"node_status_update_frequency": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.NodeStatusUpdateFrequency),
		"cluster_domain": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ClusterDomain),
		"cluster_dns": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ClusterDNS),
		"network_plugin_name": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.NetworkPluginName),
		"cloud_provider": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.CloudProvider),
		"kubelet_cgroups": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.KubeletCgroups),
		"runtime_cgroups": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.RuntimeCgroups),
		"read_only_port": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.ReadOnlyPort),
		"system_cgroups": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.SystemCgroups),
		"cgroup_root": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.CgroupRoot),
		"configure_cbr_00": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.ConfigureCBR0),
		"hairpin_mode": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.HairpinMode),
		"babysit_daemons": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.BabysitDaemons),
		"max_pods": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.MaxPods),
		"nvidia_gp_uss": func(in int32) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.NvidiaGPUs),
		"pod_cidr": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.PodCIDR),
		"resolver_config": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.ResolverConfig),
		"reconcile_cidr": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.ReconcileCIDR),
		"register_schedulable": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.RegisterSchedulable),
		"serialize_image_pulls": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.SerializeImagePulls),
		"node_labels": func(in map[string]string) interface{} {
			value := func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
			return value
		}(in.NodeLabels),
		"non_masquerade_cidr": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.NonMasqueradeCIDR),
		"enable_custom_metrics": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.EnableCustomMetrics),
		"network_plugin_mtu": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.NetworkPluginMTU),
		"image_gc_high_threshold_percent": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.ImageGCHighThresholdPercent),
		"image_gc_low_threshold_percent": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.ImageGCLowThresholdPercent),
		"image_pull_progress_deadline": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.ImagePullProgressDeadline),
		"eviction_hard": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.EvictionHard),
		"eviction_soft": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.EvictionSoft),
		"eviction_soft_grace_period": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.EvictionSoftGracePeriod),
		"eviction_pressure_transition_period": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.EvictionPressureTransitionPeriod),
		"eviction_max_pod_grace_period": func(in int32) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.EvictionMaxPodGracePeriod),
		"eviction_minimum_reclaim": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.EvictionMinimumReclaim),
		"volume_plugin_directory": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.VolumePluginDirectory),
		"taints": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.Taints),
		"feature_gates": func(in map[string]string) interface{} {
			value := func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
			return value
		}(in.FeatureGates),
		"kube_reserved": func(in map[string]string) interface{} {
			value := func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
			return value
		}(in.KubeReserved),
		"kube_reserved_cgroup": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.KubeReservedCgroup),
		"system_reserved": func(in map[string]string) interface{} {
			value := func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
			return value
		}(in.SystemReserved),
		"system_reserved_cgroup": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.SystemReservedCgroup),
		"enforce_node_allocatable": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.EnforceNodeAllocatable),
		"runtime_request_timeout": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.RuntimeRequestTimeout),
		"volume_stats_agg_period": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.VolumeStatsAggPeriod),
		"fail_swap_on": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.FailSwapOn),
		"experimental_allowed_unsafe_sysctls": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.ExperimentalAllowedUnsafeSysctls),
		"allowed_unsafe_sysctls": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.AllowedUnsafeSysctls),
		"streaming_connection_idle_timeout": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.StreamingConnectionIdleTimeout),
		"docker_disable_shared_pid": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.DockerDisableSharedPID),
		"root_dir": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.RootDir),
		"authentication_token_webhook": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.AuthenticationTokenWebhook),
		"authentication_token_webhook_cache_ttl": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.AuthenticationTokenWebhookCacheTTL),
		"cpucfs_quota": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.CPUCFSQuota),
		"cpucfs_quota_period": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.CPUCFSQuotaPeriod),
		"cpu_manager_policy": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.CpuManagerPolicy),
		"registry_pull_qps": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.RegistryPullQPS),
		"registry_burst": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.RegistryBurst),
		"topology_manager_policy": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.TopologyManagerPolicy),
		"rotate_certificates": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.RotateCertificates),
		"protect_kernel_defaults": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.ProtectKernelDefaults),
		"cgroup_driver": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.CgroupDriver),
	}
}
