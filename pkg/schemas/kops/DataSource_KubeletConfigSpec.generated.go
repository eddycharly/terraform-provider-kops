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
			"anonymous_auth":                         Computed(Ptr(Bool())),
			"authorization_mode":                     Computed(String()),
			"bootstrap_kubeconfig":                   Computed(String()),
			"client_ca_file":                         Computed(String()),
			"tls_cert_file":                          Computed(String()),
			"tls_private_key_file":                   Computed(String()),
			"tls_cipher_suites":                      Computed(List(String())),
			"tls_min_version":                        Computed(String()),
			"kubeconfig_path":                        Computed(String()),
			"require_kubeconfig":                     Computed(Ptr(Bool())),
			"log_level":                              Computed(Ptr(Int())),
			"pod_manifest_path":                      Computed(String()),
			"hostname_override":                      Computed(String()),
			"pod_infra_container_image":              Computed(String()),
			"seccomp_profile_root":                   Computed(Ptr(String())),
			"allow_privileged":                       Computed(Ptr(Bool())),
			"enable_debugging_handlers":              Computed(Ptr(Bool())),
			"register_node":                          Computed(Ptr(Bool())),
			"node_status_update_frequency":           Computed(Ptr(Duration())),
			"cluster_domain":                         Computed(String()),
			"cluster_dns":                            Computed(String()),
			"network_plugin_name":                    Computed(String()),
			"cloud_provider":                         Computed(String()),
			"kubelet_cgroups":                        Computed(String()),
			"runtime_cgroups":                        Computed(String()),
			"read_only_port":                         Computed(Ptr(Int())),
			"system_cgroups":                         Computed(String()),
			"cgroup_root":                            Computed(String()),
			"configure_cbr0":                         Computed(Ptr(Bool())),
			"hairpin_mode":                           Computed(String()),
			"babysit_daemons":                        Computed(Ptr(Bool())),
			"max_pods":                               Computed(Ptr(Int())),
			"nvidia_gp_us":                           Computed(Int()),
			"pod_cidr":                               Computed(String()),
			"resolver_config":                        Computed(Ptr(String())),
			"reconcile_cidr":                         Computed(Ptr(Bool())),
			"register_schedulable":                   Computed(Ptr(Bool())),
			"serialize_image_pulls":                  Computed(Ptr(Bool())),
			"node_labels":                            Computed(Map(String())),
			"non_masquerade_cidr":                    Computed(String()),
			"enable_custom_metrics":                  Computed(Ptr(Bool())),
			"network_plugin_mtu":                     Computed(Ptr(Int())),
			"image_gc_high_threshold_percent":        Computed(Ptr(Int())),
			"image_gc_low_threshold_percent":         Computed(Ptr(Int())),
			"image_pull_progress_deadline":           Computed(Ptr(Duration())),
			"eviction_hard":                          Computed(Ptr(String())),
			"eviction_soft":                          Computed(String()),
			"eviction_soft_grace_period":             Computed(String()),
			"eviction_pressure_transition_period":    Computed(Ptr(Duration())),
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
			"runtime_request_timeout":                Computed(Ptr(Duration())),
			"volume_stats_agg_period":                Computed(Ptr(Duration())),
			"fail_swap_on":                           Computed(Ptr(Bool())),
			"experimental_allowed_unsafe_sysctls":    Computed(List(String())),
			"allowed_unsafe_sysctls":                 Computed(List(String())),
			"streaming_connection_idle_timeout":      Computed(Ptr(Duration())),
			"docker_disable_shared_pid":              Computed(Ptr(Bool())),
			"root_dir":                               Computed(String()),
			"authentication_token_webhook":           Computed(Ptr(Bool())),
			"authentication_token_webhook_cache_ttl": Computed(Ptr(Duration())),
			"cpucfs_quota":                           Computed(Ptr(Bool())),
			"cpucfs_quota_period":                    Computed(Ptr(Duration())),
			"cpu_manager_policy":                     Computed(String()),
			"registry_pull_qps":                      Computed(Ptr(Int())),
			"registry_burst":                         Computed(Ptr(Int())),
			"topology_manager_policy":                Computed(String()),
			"rotate_certificates":                    Computed(Ptr(Bool())),
			"protect_kernel_defaults":                Computed(Ptr(Bool())),
			"cgroup_driver":                          Computed(String()),
			"housekeeping_interval":                  Computed(Ptr(Duration())),
			"event_qps":                              Computed(Ptr(Int())),
			"event_burst":                            Computed(Ptr(Int())),
			"container_log_max_size":                 Computed(String()),
			"container_log_max_files":                Computed(Ptr(Int())),
			"enable_cadvisor_json_endpoints":         Computed(Ptr(Bool())),
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
