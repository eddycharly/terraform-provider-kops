package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceKubeControllerManagerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"master":                                    Optional(String()),
			"log_level":                                 Optional(Int()),
			"service_account_private_key_file":          Optional(String()),
			"image":                                     Optional(String()),
			"cloud_provider":                            Optional(String()),
			"cluster_name":                              Optional(String()),
			"cluster_cidr":                              Optional(String()),
			"allocate_node_cidrs":                       Optional(Ptr(Bool())),
			"node_cidr_mask_size":                       Optional(Ptr(Int())),
			"configure_cloud_routes":                    Optional(Ptr(Bool())),
			"controllers":                               Optional(List(String())),
			"cidr_allocator_type":                       Optional(Ptr(String())),
			"root_ca_file":                              Optional(String()),
			"leader_election":                           Optional(Ptr(Struct(ResourceLeaderElectionConfiguration()))),
			"attach_detach_reconcile_sync_period":       Optional(Ptr(Duration())),
			"disable_attach_detach_reconcile_sync":      Optional(Ptr(Bool())),
			"terminated_pod_gc_threshold":               Optional(Ptr(Int())),
			"node_monitor_period":                       Optional(Ptr(Duration())),
			"node_monitor_grace_period":                 Optional(Ptr(Duration())),
			"pod_eviction_timeout":                      Optional(Ptr(Duration())),
			"use_service_account_credentials":           Optional(Ptr(Bool())),
			"horizontal_pod_autoscaler_sync_period":     Optional(Ptr(Duration())),
			"horizontal_pod_autoscaler_downscale_delay": Optional(Ptr(Duration())),
			"horizontal_pod_autoscaler_downscale_stabilization":   Optional(Ptr(Duration())),
			"horizontal_pod_autoscaler_upscale_delay":             Optional(Ptr(Duration())),
			"horizontal_pod_autoscaler_initial_readiness_delay":   Optional(Ptr(Duration())),
			"horizontal_pod_autoscaler_cpu_initialization_period": Optional(Ptr(Duration())),
			"horizontal_pod_autoscaler_tolerance":                 Optional(Ptr(Quantity())),
			"horizontal_pod_autoscaler_use_rest_clients":          Optional(Ptr(Bool())),
			"experimental_cluster_signing_duration":               Optional(Ptr(Duration())),
			"feature_gates":                                       Optional(Map(String())),
			"tls_cipher_suites":                                   Optional(List(String())),
			"tls_min_version":                                     Optional(String()),
			"min_resync_period":                                   Optional(String()),
			"kube_api_qps":                                        Optional(Ptr(Quantity())),
			"kube_api_burst":                                      Optional(Ptr(Int())),
			"concurrent_deployment_syncs":                         Optional(Ptr(Int())),
			"concurrent_endpoint_syncs":                           Optional(Ptr(Int())),
			"concurrent_namespace_syncs":                          Optional(Ptr(Int())),
			"concurrent_replicaset_syncs":                         Optional(Ptr(Int())),
			"concurrent_service_syncs":                            Optional(Ptr(Int())),
			"concurrent_resource_quota_syncs":                     Optional(Ptr(Int())),
			"concurrent_serviceaccount_token_syncs":               Optional(Ptr(Int())),
			"concurrent_rc_syncs":                                 Optional(Ptr(Int())),
			"authentication_kubeconfig":                           Optional(String()),
			"authorization_kubeconfig":                            Optional(String()),
			"authorization_always_allow_paths":                    Optional(List(String())),
			"external_cloud_volume_plugin":                        Optional(String()),
			"enable_profiling":                                    Optional(Ptr(Bool())),
		},
	}
}

func ExpandResourceKubeControllerManagerConfig(in map[string]interface{}) kops.KubeControllerManagerConfig {
	if in == nil {
		panic("expand KubeControllerManagerConfig failure, in is nil")
	}
	out := kops.KubeControllerManagerConfig{}
	if in, ok := in["master"]; ok && in != nil {
		out.Master = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["log_level"]; ok && in != nil {
		out.LogLevel = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	if in, ok := in["service_account_private_key_file"]; ok && in != nil {
		out.ServiceAccountPrivateKeyFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cloud_provider"]; ok && in != nil {
		out.CloudProvider = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cluster_name"]; ok && in != nil {
		out.ClusterName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cluster_cidr"]; ok && in != nil {
		out.ClusterCIDR = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["allocate_node_cidrs"]; ok && in != nil {
		out.AllocateNodeCIDRs = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["node_cidr_mask_size"]; ok && in != nil {
		out.NodeCIDRMaskSize = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["configure_cloud_routes"]; ok && in != nil {
		out.ConfigureCloudRoutes = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["controllers"]; ok && in != nil {
		out.Controllers = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["cidr_allocator_type"]; ok && in != nil {
		out.CIDRAllocatorType = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["root_ca_file"]; ok && in != nil {
		out.RootCAFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["leader_election"]; ok && in != nil {
		out.LeaderElection = func(in interface{}) *kops.LeaderElectionConfiguration {
			if in == nil {
				return nil
			}
			return func(in kops.LeaderElectionConfiguration) *kops.LeaderElectionConfiguration { return &in }(func(in interface{}) kops.LeaderElectionConfiguration {
				if in == nil {
					return kops.LeaderElectionConfiguration{}
				}
				return ExpandResourceLeaderElectionConfiguration(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["attach_detach_reconcile_sync_period"]; ok && in != nil {
		out.AttachDetachReconcileSyncPeriod = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["disable_attach_detach_reconcile_sync"]; ok && in != nil {
		out.DisableAttachDetachReconcileSync = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["terminated_pod_gc_threshold"]; ok && in != nil {
		out.TerminatedPodGCThreshold = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["node_monitor_period"]; ok && in != nil {
		out.NodeMonitorPeriod = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["node_monitor_grace_period"]; ok && in != nil {
		out.NodeMonitorGracePeriod = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["pod_eviction_timeout"]; ok && in != nil {
		out.PodEvictionTimeout = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["use_service_account_credentials"]; ok && in != nil {
		out.UseServiceAccountCredentials = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["horizontal_pod_autoscaler_sync_period"]; ok && in != nil {
		out.HorizontalPodAutoscalerSyncPeriod = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["horizontal_pod_autoscaler_downscale_delay"]; ok && in != nil {
		out.HorizontalPodAutoscalerDownscaleDelay = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["horizontal_pod_autoscaler_downscale_stabilization"]; ok && in != nil {
		out.HorizontalPodAutoscalerDownscaleStabilization = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["horizontal_pod_autoscaler_upscale_delay"]; ok && in != nil {
		out.HorizontalPodAutoscalerUpscaleDelay = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["horizontal_pod_autoscaler_initial_readiness_delay"]; ok && in != nil {
		out.HorizontalPodAutoscalerInitialReadinessDelay = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["horizontal_pod_autoscaler_cpu_initialization_period"]; ok && in != nil {
		out.HorizontalPodAutoscalerCPUInitializationPeriod = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["horizontal_pod_autoscaler_tolerance"]; ok && in != nil {
		out.HorizontalPodAutoscalerTolerance = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["horizontal_pod_autoscaler_use_rest_clients"]; ok && in != nil {
		out.HorizontalPodAutoscalerUseRestClients = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["experimental_cluster_signing_duration"]; ok && in != nil {
		out.ExperimentalClusterSigningDuration = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
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
	if in, ok := in["min_resync_period"]; ok && in != nil {
		out.MinResyncPeriod = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["kube_api_qps"]; ok && in != nil {
		out.KubeAPIQPS = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["kube_api_burst"]; ok && in != nil {
		out.KubeAPIBurst = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["concurrent_deployment_syncs"]; ok && in != nil {
		out.ConcurrentDeploymentSyncs = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["concurrent_endpoint_syncs"]; ok && in != nil {
		out.ConcurrentEndpointSyncs = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["concurrent_namespace_syncs"]; ok && in != nil {
		out.ConcurrentNamespaceSyncs = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["concurrent_replicaset_syncs"]; ok && in != nil {
		out.ConcurrentReplicasetSyncs = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["concurrent_service_syncs"]; ok && in != nil {
		out.ConcurrentServiceSyncs = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["concurrent_resource_quota_syncs"]; ok && in != nil {
		out.ConcurrentResourceQuotaSyncs = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["concurrent_serviceaccount_token_syncs"]; ok && in != nil {
		out.ConcurrentServiceaccountTokenSyncs = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["concurrent_rc_syncs"]; ok && in != nil {
		out.ConcurrentRcSyncs = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["authentication_kubeconfig"]; ok && in != nil {
		out.AuthenticationKubeconfig = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["authorization_kubeconfig"]; ok && in != nil {
		out.AuthorizationKubeconfig = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["authorization_always_allow_paths"]; ok && in != nil {
		out.AuthorizationAlwaysAllowPaths = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["external_cloud_volume_plugin"]; ok && in != nil {
		out.ExternalCloudVolumePlugin = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["enable_profiling"]; ok && in != nil {
		out.EnableProfiling = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
