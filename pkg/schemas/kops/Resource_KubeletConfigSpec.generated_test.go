package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceKubeletConfigSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KubeletConfigSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceKubeletConfigSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceKubeletConfigSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceKubeletConfigSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"api_servers":                            "",
		"anonymous_auth":                         nil,
		"authorization_mode":                     "",
		"bootstrap_kubeconfig":                   "",
		"client_ca_file":                         "",
		"tls_cert_file":                          "",
		"tls_private_key_file":                   "",
		"tls_cipher_suites":                      func() []interface{} { return nil }(),
		"tls_min_version":                        "",
		"kubeconfig_path":                        "",
		"require_kubeconfig":                     nil,
		"log_level":                              nil,
		"pod_manifest_path":                      "",
		"hostname_override":                      "",
		"pod_infra_container_image":              "",
		"seccomp_profile_root":                   nil,
		"allow_privileged":                       nil,
		"enable_debugging_handlers":              nil,
		"register_node":                          nil,
		"node_status_update_frequency":           nil,
		"cluster_domain":                         "",
		"cluster_dns":                            "",
		"network_plugin_name":                    "",
		"cloud_provider":                         "",
		"kubelet_cgroups":                        "",
		"runtime_cgroups":                        "",
		"read_only_port":                         nil,
		"system_cgroups":                         "",
		"cgroup_root":                            "",
		"configure_cbr0":                         nil,
		"hairpin_mode":                           "",
		"babysit_daemons":                        nil,
		"max_pods":                               nil,
		"nvidia_gp_us":                           0,
		"pod_cidr":                               "",
		"resolver_config":                        nil,
		"reconcile_cidr":                         nil,
		"register_schedulable":                   nil,
		"serialize_image_pulls":                  nil,
		"node_labels":                            func() map[string]interface{} { return nil }(),
		"non_masquerade_cidr":                    "",
		"enable_custom_metrics":                  nil,
		"network_plugin_mtu":                     nil,
		"image_gc_high_threshold_percent":        nil,
		"image_gc_low_threshold_percent":         nil,
		"image_pull_progress_deadline":           nil,
		"eviction_hard":                          nil,
		"eviction_soft":                          "",
		"eviction_soft_grace_period":             "",
		"eviction_pressure_transition_period":    nil,
		"eviction_max_pod_grace_period":          0,
		"eviction_minimum_reclaim":               "",
		"volume_plugin_directory":                "",
		"taints":                                 func() []interface{} { return nil }(),
		"feature_gates":                          func() map[string]interface{} { return nil }(),
		"kube_reserved":                          func() map[string]interface{} { return nil }(),
		"kube_reserved_cgroup":                   "",
		"system_reserved":                        func() map[string]interface{} { return nil }(),
		"system_reserved_cgroup":                 "",
		"enforce_node_allocatable":               "",
		"runtime_request_timeout":                nil,
		"volume_stats_agg_period":                nil,
		"fail_swap_on":                           nil,
		"experimental_allowed_unsafe_sysctls":    func() []interface{} { return nil }(),
		"allowed_unsafe_sysctls":                 func() []interface{} { return nil }(),
		"streaming_connection_idle_timeout":      nil,
		"docker_disable_shared_pid":              nil,
		"root_dir":                               "",
		"authentication_token_webhook":           nil,
		"authentication_token_webhook_cache_ttl": nil,
		"cpucfs_quota":                           nil,
		"cpucfs_quota_period":                    nil,
		"cpu_manager_policy":                     "",
		"registry_pull_qps":                      nil,
		"registry_burst":                         nil,
		"topology_manager_policy":                "",
		"rotate_certificates":                    nil,
		"protect_kernel_defaults":                nil,
		"cgroup_driver":                          "",
		"housekeeping_interval":                  nil,
		"event_qps":                              nil,
		"event_burst":                            nil,
		"container_log_max_size":                 "",
		"container_log_max_files":                nil,
		"enable_cadvisor_json_endpoints":         nil,
	}
	type args struct {
		in kops.KubeletConfigSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KubeletConfigSpec{},
			},
			want: _default,
		},
		{
			name: "ApiServers - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.APIServers = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AnonymousAuth - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.AnonymousAuth = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationMode - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.AuthorizationMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BootstrapKubeconfig - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.BootstrapKubeconfig = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClientCAFile - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ClientCAFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSCertFile - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.TLSCertFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSPrivateKeyFile - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.TLSPrivateKeyFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSCipherSuites - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.TLSCipherSuites = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSMinVersion - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.TLSMinVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeconfigPath - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.KubeconfigPath = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RequireKubeconfig - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RequireKubeconfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.LogLevel = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodManifestPath - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.PodManifestPath = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HostnameOverride - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.HostnameOverride = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodInfraContainerImage - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.PodInfraContainerImage = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SeccompProfileRoot - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.SeccompProfileRoot = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowPrivileged - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.AllowPrivileged = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableDebuggingHandlers - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EnableDebuggingHandlers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RegisterNode - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RegisterNode = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeStatusUpdateFrequency - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.NodeStatusUpdateFrequency = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterDomain - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ClusterDomain = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterDns - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ClusterDNS = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NetworkPluginName - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.NetworkPluginName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudProvider - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.CloudProvider = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeletCgroups - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.KubeletCgroups = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RuntimeCgroups - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RuntimeCgroups = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ReadOnlyPort - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ReadOnlyPort = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SystemCgroups - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.SystemCgroups = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CgroupRoot - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.CgroupRoot = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigureCBR0 - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ConfigureCBR0 = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HairpinMode - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.HairpinMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BabysitDaemons - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.BabysitDaemons = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxPods - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.MaxPods = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NvidiaGPUs - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.NvidiaGPUs = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodCidr - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.PodCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ResolverConfig - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ResolverConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ReconcileCidr - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ReconcileCIDR = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RegisterSchedulable - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RegisterSchedulable = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SerializeImagePulls - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.SerializeImagePulls = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeLabels - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.NodeLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NonMasqueradeCidr - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.NonMasqueradeCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableCustomMetrics - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EnableCustomMetrics = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NetworkPluginMTU - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.NetworkPluginMTU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ImageGCHighThresholdPercent - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ImageGCHighThresholdPercent = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ImageGCLowThresholdPercent - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ImageGCLowThresholdPercent = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ImagePullProgressDeadline - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ImagePullProgressDeadline = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EvictionHard - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EvictionHard = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EvictionSoft - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EvictionSoft = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EvictionSoftGracePeriod - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EvictionSoftGracePeriod = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EvictionPressureTransitionPeriod - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EvictionPressureTransitionPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EvictionMaxPodGracePeriod - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EvictionMaxPodGracePeriod = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EvictionMinimumReclaim - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EvictionMinimumReclaim = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VolumePluginDirectory - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.VolumePluginDirectory = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Taints - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.Taints = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FeatureGates - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.FeatureGates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeReserved - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.KubeReserved = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeReservedCgroup - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.KubeReservedCgroup = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SystemReserved - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.SystemReserved = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SystemReservedCgroup - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.SystemReservedCgroup = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnforceNodeAllocatable - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EnforceNodeAllocatable = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RuntimeRequestTimeout - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RuntimeRequestTimeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VolumeStatsAggPeriod - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.VolumeStatsAggPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FailSwapOn - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.FailSwapOn = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExperimentalAllowedUnsafeSysctls - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ExperimentalAllowedUnsafeSysctls = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowedUnsafeSysctls - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.AllowedUnsafeSysctls = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "StreamingConnectionIdleTimeout - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.StreamingConnectionIdleTimeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DockerDisableSharedPID - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.DockerDisableSharedPID = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootDir - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RootDir = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthenticationTokenWebhook - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.AuthenticationTokenWebhook = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthenticationTokenWebhookCacheTTL - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.AuthenticationTokenWebhookCacheTTL = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CPUCFSQuota - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.CPUCFSQuota = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CPUCFSQuotaPeriod - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.CPUCFSQuotaPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuManagerPolicy - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.CpuManagerPolicy = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RegistryPullQPS - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RegistryPullQPS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RegistryBurst - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RegistryBurst = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TopologyManagerPolicy - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.TopologyManagerPolicy = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RotateCertificates - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RotateCertificates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ProtectKernelDefaults - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ProtectKernelDefaults = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CgroupDriver - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.CgroupDriver = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HousekeepingInterval - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.HousekeepingInterval = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EventQPS - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EventQPS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EventBurst - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EventBurst = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerLogMaxSize - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ContainerLogMaxSize = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerLogMaxFiles - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ContainerLogMaxFiles = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableCadvisorJsonEndpoints - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EnableCadvisorJsonEndpoints = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceKubeletConfigSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceKubeletConfigSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceKubeletConfigSpec(t *testing.T) {
	_default := map[string]interface{}{
		"api_servers":                            "",
		"anonymous_auth":                         nil,
		"authorization_mode":                     "",
		"bootstrap_kubeconfig":                   "",
		"client_ca_file":                         "",
		"tls_cert_file":                          "",
		"tls_private_key_file":                   "",
		"tls_cipher_suites":                      func() []interface{} { return nil }(),
		"tls_min_version":                        "",
		"kubeconfig_path":                        "",
		"require_kubeconfig":                     nil,
		"log_level":                              nil,
		"pod_manifest_path":                      "",
		"hostname_override":                      "",
		"pod_infra_container_image":              "",
		"seccomp_profile_root":                   nil,
		"allow_privileged":                       nil,
		"enable_debugging_handlers":              nil,
		"register_node":                          nil,
		"node_status_update_frequency":           nil,
		"cluster_domain":                         "",
		"cluster_dns":                            "",
		"network_plugin_name":                    "",
		"cloud_provider":                         "",
		"kubelet_cgroups":                        "",
		"runtime_cgroups":                        "",
		"read_only_port":                         nil,
		"system_cgroups":                         "",
		"cgroup_root":                            "",
		"configure_cbr0":                         nil,
		"hairpin_mode":                           "",
		"babysit_daemons":                        nil,
		"max_pods":                               nil,
		"nvidia_gp_us":                           0,
		"pod_cidr":                               "",
		"resolver_config":                        nil,
		"reconcile_cidr":                         nil,
		"register_schedulable":                   nil,
		"serialize_image_pulls":                  nil,
		"node_labels":                            func() map[string]interface{} { return nil }(),
		"non_masquerade_cidr":                    "",
		"enable_custom_metrics":                  nil,
		"network_plugin_mtu":                     nil,
		"image_gc_high_threshold_percent":        nil,
		"image_gc_low_threshold_percent":         nil,
		"image_pull_progress_deadline":           nil,
		"eviction_hard":                          nil,
		"eviction_soft":                          "",
		"eviction_soft_grace_period":             "",
		"eviction_pressure_transition_period":    nil,
		"eviction_max_pod_grace_period":          0,
		"eviction_minimum_reclaim":               "",
		"volume_plugin_directory":                "",
		"taints":                                 func() []interface{} { return nil }(),
		"feature_gates":                          func() map[string]interface{} { return nil }(),
		"kube_reserved":                          func() map[string]interface{} { return nil }(),
		"kube_reserved_cgroup":                   "",
		"system_reserved":                        func() map[string]interface{} { return nil }(),
		"system_reserved_cgroup":                 "",
		"enforce_node_allocatable":               "",
		"runtime_request_timeout":                nil,
		"volume_stats_agg_period":                nil,
		"fail_swap_on":                           nil,
		"experimental_allowed_unsafe_sysctls":    func() []interface{} { return nil }(),
		"allowed_unsafe_sysctls":                 func() []interface{} { return nil }(),
		"streaming_connection_idle_timeout":      nil,
		"docker_disable_shared_pid":              nil,
		"root_dir":                               "",
		"authentication_token_webhook":           nil,
		"authentication_token_webhook_cache_ttl": nil,
		"cpucfs_quota":                           nil,
		"cpucfs_quota_period":                    nil,
		"cpu_manager_policy":                     "",
		"registry_pull_qps":                      nil,
		"registry_burst":                         nil,
		"topology_manager_policy":                "",
		"rotate_certificates":                    nil,
		"protect_kernel_defaults":                nil,
		"cgroup_driver":                          "",
		"housekeeping_interval":                  nil,
		"event_qps":                              nil,
		"event_burst":                            nil,
		"container_log_max_size":                 "",
		"container_log_max_files":                nil,
		"enable_cadvisor_json_endpoints":         nil,
	}
	type args struct {
		in kops.KubeletConfigSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KubeletConfigSpec{},
			},
			want: _default,
		},
		{
			name: "ApiServers - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.APIServers = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AnonymousAuth - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.AnonymousAuth = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationMode - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.AuthorizationMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BootstrapKubeconfig - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.BootstrapKubeconfig = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClientCAFile - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ClientCAFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSCertFile - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.TLSCertFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSPrivateKeyFile - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.TLSPrivateKeyFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSCipherSuites - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.TLSCipherSuites = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSMinVersion - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.TLSMinVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeconfigPath - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.KubeconfigPath = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RequireKubeconfig - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RequireKubeconfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.LogLevel = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodManifestPath - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.PodManifestPath = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HostnameOverride - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.HostnameOverride = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodInfraContainerImage - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.PodInfraContainerImage = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SeccompProfileRoot - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.SeccompProfileRoot = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowPrivileged - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.AllowPrivileged = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableDebuggingHandlers - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EnableDebuggingHandlers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RegisterNode - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RegisterNode = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeStatusUpdateFrequency - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.NodeStatusUpdateFrequency = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterDomain - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ClusterDomain = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterDns - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ClusterDNS = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NetworkPluginName - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.NetworkPluginName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudProvider - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.CloudProvider = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeletCgroups - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.KubeletCgroups = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RuntimeCgroups - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RuntimeCgroups = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ReadOnlyPort - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ReadOnlyPort = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SystemCgroups - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.SystemCgroups = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CgroupRoot - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.CgroupRoot = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigureCBR0 - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ConfigureCBR0 = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HairpinMode - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.HairpinMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BabysitDaemons - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.BabysitDaemons = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxPods - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.MaxPods = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NvidiaGPUs - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.NvidiaGPUs = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodCidr - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.PodCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ResolverConfig - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ResolverConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ReconcileCidr - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ReconcileCIDR = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RegisterSchedulable - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RegisterSchedulable = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SerializeImagePulls - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.SerializeImagePulls = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeLabels - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.NodeLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NonMasqueradeCidr - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.NonMasqueradeCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableCustomMetrics - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EnableCustomMetrics = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NetworkPluginMTU - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.NetworkPluginMTU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ImageGCHighThresholdPercent - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ImageGCHighThresholdPercent = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ImageGCLowThresholdPercent - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ImageGCLowThresholdPercent = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ImagePullProgressDeadline - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ImagePullProgressDeadline = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EvictionHard - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EvictionHard = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EvictionSoft - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EvictionSoft = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EvictionSoftGracePeriod - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EvictionSoftGracePeriod = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EvictionPressureTransitionPeriod - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EvictionPressureTransitionPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EvictionMaxPodGracePeriod - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EvictionMaxPodGracePeriod = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EvictionMinimumReclaim - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EvictionMinimumReclaim = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VolumePluginDirectory - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.VolumePluginDirectory = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Taints - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.Taints = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FeatureGates - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.FeatureGates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeReserved - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.KubeReserved = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeReservedCgroup - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.KubeReservedCgroup = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SystemReserved - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.SystemReserved = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SystemReservedCgroup - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.SystemReservedCgroup = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnforceNodeAllocatable - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EnforceNodeAllocatable = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RuntimeRequestTimeout - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RuntimeRequestTimeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VolumeStatsAggPeriod - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.VolumeStatsAggPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FailSwapOn - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.FailSwapOn = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExperimentalAllowedUnsafeSysctls - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ExperimentalAllowedUnsafeSysctls = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowedUnsafeSysctls - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.AllowedUnsafeSysctls = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "StreamingConnectionIdleTimeout - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.StreamingConnectionIdleTimeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DockerDisableSharedPID - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.DockerDisableSharedPID = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootDir - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RootDir = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthenticationTokenWebhook - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.AuthenticationTokenWebhook = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthenticationTokenWebhookCacheTTL - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.AuthenticationTokenWebhookCacheTTL = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CPUCFSQuota - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.CPUCFSQuota = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CPUCFSQuotaPeriod - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.CPUCFSQuotaPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuManagerPolicy - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.CpuManagerPolicy = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RegistryPullQPS - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RegistryPullQPS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RegistryBurst - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RegistryBurst = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TopologyManagerPolicy - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.TopologyManagerPolicy = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RotateCertificates - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.RotateCertificates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ProtectKernelDefaults - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ProtectKernelDefaults = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CgroupDriver - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.CgroupDriver = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HousekeepingInterval - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.HousekeepingInterval = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EventQPS - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EventQPS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EventBurst - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EventBurst = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerLogMaxSize - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ContainerLogMaxSize = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerLogMaxFiles - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.ContainerLogMaxFiles = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableCadvisorJsonEndpoints - default",
			args: args{
				in: func() kops.KubeletConfigSpec {
					subject := kops.KubeletConfigSpec{}
					subject.EnableCadvisorJsonEndpoints = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceKubeletConfigSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceKubeletConfigSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
