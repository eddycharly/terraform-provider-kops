package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceKubeletConfigSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_servers":                            Computed(String()),
			"anonymous_auth":                         Computed(Nullable(Bool())),
			"authorization_mode":                     Computed(String()),
			"bootstrap_kubeconfig":                   Computed(String()),
			"client_ca_file":                         Computed(String()),
			"tls_cert_file":                          Computed(String()),
			"tls_private_key_file":                   Computed(String()),
			"tls_cipher_suites":                      Computed(List(String())),
			"tls_min_version":                        Computed(String()),
			"kubeconfig_path":                        Computed(String()),
			"require_kubeconfig":                     Computed(Nullable(Bool())),
			"log_level":                              Computed(Nullable(Int())),
			"pod_manifest_path":                      Computed(String()),
			"hostname_override":                      Computed(String()),
			"pod_infra_container_image":              Computed(String()),
			"seccomp_profile_root":                   Computed(Nullable(String())),
			"allow_privileged":                       Computed(Nullable(Bool())),
			"enable_debugging_handlers":              Computed(Nullable(Bool())),
			"register_node":                          Computed(Nullable(Bool())),
			"node_status_update_frequency":           Computed(Nullable(Duration())),
			"cluster_domain":                         Computed(String()),
			"cluster_dns":                            Computed(String()),
			"network_plugin_name":                    Computed(String()),
			"cloud_provider":                         Computed(String()),
			"kubelet_cgroups":                        Computed(String()),
			"runtime_cgroups":                        Computed(String()),
			"read_only_port":                         Computed(Nullable(Int())),
			"system_cgroups":                         Computed(String()),
			"cgroup_root":                            Computed(String()),
			"configure_cbr0":                         Computed(Nullable(Bool())),
			"hairpin_mode":                           Computed(String()),
			"babysit_daemons":                        Computed(Nullable(Bool())),
			"max_pods":                               Computed(Nullable(Int())),
			"nvidia_gp_us":                           Computed(Int()),
			"pod_cidr":                               Computed(String()),
			"resolver_config":                        Computed(Nullable(String())),
			"reconcile_cidr":                         Computed(Nullable(Bool())),
			"register_schedulable":                   Computed(Nullable(Bool())),
			"serialize_image_pulls":                  Computed(Nullable(Bool())),
			"node_labels":                            Computed(Map(String())),
			"non_masquerade_cidr":                    Computed(String()),
			"enable_custom_metrics":                  Computed(Nullable(Bool())),
			"network_plugin_mtu":                     Computed(Nullable(Int())),
			"image_gc_high_threshold_percent":        Computed(Nullable(Int())),
			"image_gc_low_threshold_percent":         Computed(Nullable(Int())),
			"image_pull_progress_deadline":           Computed(Nullable(Duration())),
			"eviction_hard":                          Computed(Nullable(String())),
			"eviction_soft":                          Computed(String()),
			"eviction_soft_grace_period":             Computed(String()),
			"eviction_pressure_transition_period":    Computed(Nullable(Duration())),
			"eviction_max_pod_grace_period":          Computed(Int()),
			"eviction_minimum_reclaim":               Computed(String()),
			"volume_plugin_directory":                Computed(String()),
			"taints":                                 Computed(List(String())),
			"feature_gates":                          Computed(Map(String())),
			"kube_reserved":                          Computed(Map(String())),
			"kube_reserved_cgroup":                   Computed(String()),
			"system_reserved":                        Computed(Map(String())),
			"system_reserved_cgroup":                 Computed(String()),
			"enforce_node_allocatable":               Computed(String()),
			"runtime_request_timeout":                Computed(Nullable(Duration())),
			"volume_stats_agg_period":                Computed(Nullable(Duration())),
			"fail_swap_on":                           Computed(Nullable(Bool())),
			"experimental_allowed_unsafe_sysctls":    Computed(List(String())),
			"allowed_unsafe_sysctls":                 Computed(List(String())),
			"streaming_connection_idle_timeout":      Computed(Nullable(Duration())),
			"docker_disable_shared_pid":              Computed(Nullable(Bool())),
			"root_dir":                               Computed(String()),
			"authentication_token_webhook":           Computed(Nullable(Bool())),
			"authentication_token_webhook_cache_ttl": Computed(Nullable(Duration())),
			"cpucfs_quota":                           Computed(Nullable(Bool())),
			"cpucfs_quota_period":                    Computed(Nullable(Duration())),
			"cpu_manager_policy":                     Computed(String()),
			"registry_pull_qps":                      Computed(Nullable(Int())),
			"registry_burst":                         Computed(Nullable(Int())),
			"topology_manager_policy":                Computed(String()),
			"rotate_certificates":                    Computed(Nullable(Bool())),
			"protect_kernel_defaults":                Computed(Nullable(Bool())),
			"cgroup_driver":                          Computed(String()),
			"housekeeping_interval":                  Computed(Nullable(Duration())),
			"event_qps":                              Computed(Nullable(Int())),
			"event_burst":                            Computed(Nullable(Int())),
			"container_log_max_size":                 Computed(String()),
			"container_log_max_files":                Computed(Nullable(Int())),
			"enable_cadvisor_json_endpoints":         Computed(Nullable(Bool())),
		},
	}
}

func ExpandDataSourceKubeletConfigSpec(in map[string]interface{}) kops.KubeletConfigSpec {
	if in == nil {
		panic("expand KubeletConfigSpec failure, in is nil")
	}
	out := kops.KubeletConfigSpec{}
	if in, ok := in["api_servers"]; ok && in != nil {
		out.APIServers = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["anonymous_auth"]; ok && in != nil {
		out.AnonymousAuth = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["authorization_mode"]; ok && in != nil {
		out.AuthorizationMode = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["bootstrap_kubeconfig"]; ok && in != nil {
		out.BootstrapKubeconfig = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["client_ca_file"]; ok && in != nil {
		out.ClientCAFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["tls_cert_file"]; ok && in != nil {
		out.TLSCertFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["tls_private_key_file"]; ok && in != nil {
		out.TLSPrivateKeyFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["tls_cipher_suites"]; ok && in != nil {
		out.TLSCipherSuites = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["tls_min_version"]; ok && in != nil {
		out.TLSMinVersion = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["kubeconfig_path"]; ok && in != nil {
		out.KubeconfigPath = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["require_kubeconfig"]; ok && in != nil {
		out.RequireKubeconfig = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["log_level"]; ok && in != nil {
		out.LogLevel = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["pod_manifest_path"]; ok && in != nil {
		out.PodManifestPath = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["hostname_override"]; ok && in != nil {
		out.HostnameOverride = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["pod_infra_container_image"]; ok && in != nil {
		out.PodInfraContainerImage = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["seccomp_profile_root"]; ok && in != nil {
		out.SeccompProfileRoot = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["allow_privileged"]; ok && in != nil {
		out.AllowPrivileged = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["enable_debugging_handlers"]; ok && in != nil {
		out.EnableDebuggingHandlers = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["register_node"]; ok && in != nil {
		out.RegisterNode = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["node_status_update_frequency"]; ok && in != nil {
		out.NodeStatusUpdateFrequency = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["cluster_domain"]; ok && in != nil {
		out.ClusterDomain = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cluster_dns"]; ok && in != nil {
		out.ClusterDNS = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["network_plugin_name"]; ok && in != nil {
		out.NetworkPluginName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cloud_provider"]; ok && in != nil {
		out.CloudProvider = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["kubelet_cgroups"]; ok && in != nil {
		out.KubeletCgroups = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["runtime_cgroups"]; ok && in != nil {
		out.RuntimeCgroups = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["read_only_port"]; ok && in != nil {
		out.ReadOnlyPort = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["system_cgroups"]; ok && in != nil {
		out.SystemCgroups = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cgroup_root"]; ok && in != nil {
		out.CgroupRoot = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["configure_cbr0"]; ok && in != nil {
		out.ConfigureCBR0 = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["hairpin_mode"]; ok && in != nil {
		out.HairpinMode = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["babysit_daemons"]; ok && in != nil {
		out.BabysitDaemons = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["max_pods"]; ok && in != nil {
		out.MaxPods = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["nvidia_gp_us"]; ok && in != nil {
		out.NvidiaGPUs = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	if in, ok := in["pod_cidr"]; ok && in != nil {
		out.PodCIDR = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["resolver_config"]; ok && in != nil {
		out.ResolverConfig = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["reconcile_cidr"]; ok && in != nil {
		out.ReconcileCIDR = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["register_schedulable"]; ok && in != nil {
		out.RegisterSchedulable = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["serialize_image_pulls"]; ok && in != nil {
		out.SerializeImagePulls = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["node_labels"]; ok && in != nil {
		out.NodeLabels = func(in interface{}) map[string]string {
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
	if in, ok := in["non_masquerade_cidr"]; ok && in != nil {
		out.NonMasqueradeCIDR = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["enable_custom_metrics"]; ok && in != nil {
		out.EnableCustomMetrics = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["network_plugin_mtu"]; ok && in != nil {
		out.NetworkPluginMTU = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["image_gc_high_threshold_percent"]; ok && in != nil {
		out.ImageGCHighThresholdPercent = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["image_gc_low_threshold_percent"]; ok && in != nil {
		out.ImageGCLowThresholdPercent = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["image_pull_progress_deadline"]; ok && in != nil {
		out.ImagePullProgressDeadline = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["eviction_hard"]; ok && in != nil {
		out.EvictionHard = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["eviction_soft"]; ok && in != nil {
		out.EvictionSoft = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["eviction_soft_grace_period"]; ok && in != nil {
		out.EvictionSoftGracePeriod = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["eviction_pressure_transition_period"]; ok && in != nil {
		out.EvictionPressureTransitionPeriod = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["eviction_max_pod_grace_period"]; ok && in != nil {
		out.EvictionMaxPodGracePeriod = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	if in, ok := in["eviction_minimum_reclaim"]; ok && in != nil {
		out.EvictionMinimumReclaim = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["volume_plugin_directory"]; ok && in != nil {
		out.VolumePluginDirectory = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["taints"]; ok && in != nil {
		out.Taints = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["feature_gates"]; ok && in != nil {
		out.FeatureGates = func(in interface{}) map[string]string {
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
	if in, ok := in["kube_reserved"]; ok && in != nil {
		out.KubeReserved = func(in interface{}) map[string]string {
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
	if in, ok := in["kube_reserved_cgroup"]; ok && in != nil {
		out.KubeReservedCgroup = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["system_reserved"]; ok && in != nil {
		out.SystemReserved = func(in interface{}) map[string]string {
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
	if in, ok := in["system_reserved_cgroup"]; ok && in != nil {
		out.SystemReservedCgroup = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["enforce_node_allocatable"]; ok && in != nil {
		out.EnforceNodeAllocatable = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["runtime_request_timeout"]; ok && in != nil {
		out.RuntimeRequestTimeout = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["volume_stats_agg_period"]; ok && in != nil {
		out.VolumeStatsAggPeriod = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["fail_swap_on"]; ok && in != nil {
		out.FailSwapOn = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["experimental_allowed_unsafe_sysctls"]; ok && in != nil {
		out.ExperimentalAllowedUnsafeSysctls = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["allowed_unsafe_sysctls"]; ok && in != nil {
		out.AllowedUnsafeSysctls = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["streaming_connection_idle_timeout"]; ok && in != nil {
		out.StreamingConnectionIdleTimeout = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["docker_disable_shared_pid"]; ok && in != nil {
		out.DockerDisableSharedPID = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["root_dir"]; ok && in != nil {
		out.RootDir = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["authentication_token_webhook"]; ok && in != nil {
		out.AuthenticationTokenWebhook = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["authentication_token_webhook_cache_ttl"]; ok && in != nil {
		out.AuthenticationTokenWebhookCacheTTL = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["cpucfs_quota"]; ok && in != nil {
		out.CPUCFSQuota = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["cpucfs_quota_period"]; ok && in != nil {
		out.CPUCFSQuotaPeriod = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["cpu_manager_policy"]; ok && in != nil {
		out.CpuManagerPolicy = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["registry_pull_qps"]; ok && in != nil {
		out.RegistryPullQPS = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["registry_burst"]; ok && in != nil {
		out.RegistryBurst = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["topology_manager_policy"]; ok && in != nil {
		out.TopologyManagerPolicy = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["rotate_certificates"]; ok && in != nil {
		out.RotateCertificates = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["protect_kernel_defaults"]; ok && in != nil {
		out.ProtectKernelDefaults = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["cgroup_driver"]; ok && in != nil {
		out.CgroupDriver = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["housekeeping_interval"]; ok && in != nil {
		out.HousekeepingInterval = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["event_qps"]; ok && in != nil {
		out.EventQPS = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["event_burst"]; ok && in != nil {
		out.EventBurst = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["container_log_max_size"]; ok && in != nil {
		out.ContainerLogMaxSize = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["container_log_max_files"]; ok && in != nil {
		out.ContainerLogMaxFiles = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["enable_cadvisor_json_endpoints"]; ok && in != nil {
		out.EnableCadvisorJsonEndpoints = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}

func FlattenDataSourceKubeletConfigSpecInto(in kops.KubeletConfigSpec, out map[string]interface{}) {
	out["api_servers"] = func(in string) interface{} { return string(in) }(in.APIServers)
	out["anonymous_auth"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.AnonymousAuth)
	out["authorization_mode"] = func(in string) interface{} { return string(in) }(in.AuthorizationMode)
	out["bootstrap_kubeconfig"] = func(in string) interface{} { return string(in) }(in.BootstrapKubeconfig)
	out["client_ca_file"] = func(in string) interface{} { return string(in) }(in.ClientCAFile)
	out["tls_cert_file"] = func(in string) interface{} { return string(in) }(in.TLSCertFile)
	out["tls_private_key_file"] = func(in string) interface{} { return string(in) }(in.TLSPrivateKeyFile)
	out["tls_cipher_suites"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.TLSCipherSuites)
	out["tls_min_version"] = func(in string) interface{} { return string(in) }(in.TLSMinVersion)
	out["kubeconfig_path"] = func(in string) interface{} { return string(in) }(in.KubeconfigPath)
	out["require_kubeconfig"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.RequireKubeconfig)
	out["log_level"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.LogLevel)
	out["pod_manifest_path"] = func(in string) interface{} { return string(in) }(in.PodManifestPath)
	out["hostname_override"] = func(in string) interface{} { return string(in) }(in.HostnameOverride)
	out["pod_infra_container_image"] = func(in string) interface{} { return string(in) }(in.PodInfraContainerImage)
	out["seccomp_profile_root"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.SeccompProfileRoot)
	out["allow_privileged"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.AllowPrivileged)
	out["enable_debugging_handlers"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.EnableDebuggingHandlers)
	out["register_node"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.RegisterNode)
	out["node_status_update_frequency"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.NodeStatusUpdateFrequency)
	out["cluster_domain"] = func(in string) interface{} { return string(in) }(in.ClusterDomain)
	out["cluster_dns"] = func(in string) interface{} { return string(in) }(in.ClusterDNS)
	out["network_plugin_name"] = func(in string) interface{} { return string(in) }(in.NetworkPluginName)
	out["cloud_provider"] = func(in string) interface{} { return string(in) }(in.CloudProvider)
	out["kubelet_cgroups"] = func(in string) interface{} { return string(in) }(in.KubeletCgroups)
	out["runtime_cgroups"] = func(in string) interface{} { return string(in) }(in.RuntimeCgroups)
	out["read_only_port"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.ReadOnlyPort)
	out["system_cgroups"] = func(in string) interface{} { return string(in) }(in.SystemCgroups)
	out["cgroup_root"] = func(in string) interface{} { return string(in) }(in.CgroupRoot)
	out["configure_cbr0"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.ConfigureCBR0)
	out["hairpin_mode"] = func(in string) interface{} { return string(in) }(in.HairpinMode)
	out["babysit_daemons"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.BabysitDaemons)
	out["max_pods"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.MaxPods)
	out["nvidia_gp_us"] = func(in int32) interface{} { return int(in) }(in.NvidiaGPUs)
	out["pod_cidr"] = func(in string) interface{} { return string(in) }(in.PodCIDR)
	out["resolver_config"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.ResolverConfig)
	out["reconcile_cidr"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.ReconcileCIDR)
	out["register_schedulable"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.RegisterSchedulable)
	out["serialize_image_pulls"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.SerializeImagePulls)
	out["node_labels"] = func(in map[string]string) interface{} {
		if in == nil {
			return nil
		}
		out := map[string]interface{}{}
		for key, in := range in {
			out[key] = func(in string) interface{} { return string(in) }(in)
		}
		return out
	}(in.NodeLabels)
	out["non_masquerade_cidr"] = func(in string) interface{} { return string(in) }(in.NonMasqueradeCIDR)
	out["enable_custom_metrics"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.EnableCustomMetrics)
	out["network_plugin_mtu"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.NetworkPluginMTU)
	out["image_gc_high_threshold_percent"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.ImageGCHighThresholdPercent)
	out["image_gc_low_threshold_percent"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.ImageGCLowThresholdPercent)
	out["image_pull_progress_deadline"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.ImagePullProgressDeadline)
	out["eviction_hard"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.EvictionHard)
	out["eviction_soft"] = func(in string) interface{} { return string(in) }(in.EvictionSoft)
	out["eviction_soft_grace_period"] = func(in string) interface{} { return string(in) }(in.EvictionSoftGracePeriod)
	out["eviction_pressure_transition_period"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.EvictionPressureTransitionPeriod)
	out["eviction_max_pod_grace_period"] = func(in int32) interface{} { return int(in) }(in.EvictionMaxPodGracePeriod)
	out["eviction_minimum_reclaim"] = func(in string) interface{} { return string(in) }(in.EvictionMinimumReclaim)
	out["volume_plugin_directory"] = func(in string) interface{} { return string(in) }(in.VolumePluginDirectory)
	out["taints"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.Taints)
	out["feature_gates"] = func(in map[string]string) interface{} {
		if in == nil {
			return nil
		}
		out := map[string]interface{}{}
		for key, in := range in {
			out[key] = func(in string) interface{} { return string(in) }(in)
		}
		return out
	}(in.FeatureGates)
	out["kube_reserved"] = func(in map[string]string) interface{} {
		if in == nil {
			return nil
		}
		out := map[string]interface{}{}
		for key, in := range in {
			out[key] = func(in string) interface{} { return string(in) }(in)
		}
		return out
	}(in.KubeReserved)
	out["kube_reserved_cgroup"] = func(in string) interface{} { return string(in) }(in.KubeReservedCgroup)
	out["system_reserved"] = func(in map[string]string) interface{} {
		if in == nil {
			return nil
		}
		out := map[string]interface{}{}
		for key, in := range in {
			out[key] = func(in string) interface{} { return string(in) }(in)
		}
		return out
	}(in.SystemReserved)
	out["system_reserved_cgroup"] = func(in string) interface{} { return string(in) }(in.SystemReservedCgroup)
	out["enforce_node_allocatable"] = func(in string) interface{} { return string(in) }(in.EnforceNodeAllocatable)
	out["runtime_request_timeout"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.RuntimeRequestTimeout)
	out["volume_stats_agg_period"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.VolumeStatsAggPeriod)
	out["fail_swap_on"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.FailSwapOn)
	out["experimental_allowed_unsafe_sysctls"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.ExperimentalAllowedUnsafeSysctls)
	out["allowed_unsafe_sysctls"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.AllowedUnsafeSysctls)
	out["streaming_connection_idle_timeout"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.StreamingConnectionIdleTimeout)
	out["docker_disable_shared_pid"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.DockerDisableSharedPID)
	out["root_dir"] = func(in string) interface{} { return string(in) }(in.RootDir)
	out["authentication_token_webhook"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.AuthenticationTokenWebhook)
	out["authentication_token_webhook_cache_ttl"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.AuthenticationTokenWebhookCacheTTL)
	out["cpucfs_quota"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.CPUCFSQuota)
	out["cpucfs_quota_period"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.CPUCFSQuotaPeriod)
	out["cpu_manager_policy"] = func(in string) interface{} { return string(in) }(in.CpuManagerPolicy)
	out["registry_pull_qps"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.RegistryPullQPS)
	out["registry_burst"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.RegistryBurst)
	out["topology_manager_policy"] = func(in string) interface{} { return string(in) }(in.TopologyManagerPolicy)
	out["rotate_certificates"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.RotateCertificates)
	out["protect_kernel_defaults"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.ProtectKernelDefaults)
	out["cgroup_driver"] = func(in string) interface{} { return string(in) }(in.CgroupDriver)
	out["housekeeping_interval"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.HousekeepingInterval)
	out["event_qps"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.EventQPS)
	out["event_burst"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.EventBurst)
	out["container_log_max_size"] = func(in string) interface{} { return string(in) }(in.ContainerLogMaxSize)
	out["container_log_max_files"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.ContainerLogMaxFiles)
	out["enable_cadvisor_json_endpoints"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.EnableCadvisorJsonEndpoints)
}

func FlattenDataSourceKubeletConfigSpec(in kops.KubeletConfigSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKubeletConfigSpecInto(in, out)
	return out
}
