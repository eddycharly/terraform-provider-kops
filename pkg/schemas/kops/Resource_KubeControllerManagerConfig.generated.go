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
			"allocate_node_cidrs":                       Optional(Nullable(Bool())),
			"node_cidr_mask_size":                       Optional(Nullable(Int())),
			"configure_cloud_routes":                    Optional(Nullable(Bool())),
			"controllers":                               Optional(List(String())),
			"cidr_allocator_type":                       Optional(Nullable(String())),
			"root_ca_file":                              Optional(String()),
			"leader_election":                           Optional(Struct(ResourceLeaderElectionConfiguration())),
			"attach_detach_reconcile_sync_period":       Optional(Nullable(Duration())),
			"disable_attach_detach_reconcile_sync":      Optional(Nullable(Bool())),
			"terminated_pod_gc_threshold":               Optional(Nullable(Int())),
			"node_monitor_period":                       Optional(Nullable(Duration())),
			"node_monitor_grace_period":                 Optional(Nullable(Duration())),
			"pod_eviction_timeout":                      Optional(Nullable(Duration())),
			"use_service_account_credentials":           Optional(Nullable(Bool())),
			"horizontal_pod_autoscaler_sync_period":     Optional(Nullable(Duration())),
			"horizontal_pod_autoscaler_downscale_delay": Optional(Nullable(Duration())),
			"horizontal_pod_autoscaler_downscale_stabilization":   Optional(Nullable(Duration())),
			"horizontal_pod_autoscaler_upscale_delay":             Optional(Nullable(Duration())),
			"horizontal_pod_autoscaler_initial_readiness_delay":   Optional(Nullable(Duration())),
			"horizontal_pod_autoscaler_cpu_initialization_period": Optional(Nullable(Duration())),
			"horizontal_pod_autoscaler_tolerance":                 Optional(Nullable(Quantity())),
			"horizontal_pod_autoscaler_use_rest_clients":          Optional(Nullable(Bool())),
			"experimental_cluster_signing_duration":               Optional(Nullable(Duration())),
			"feature_gates":                                       Optional(Map(String())),
			"tls_cipher_suites":                                   Optional(List(String())),
			"tls_min_version":                                     Optional(String()),
			"min_resync_period":                                   Optional(String()),
			"kube_api_qps":                                        Optional(Nullable(Quantity())),
			"kube_api_burst":                                      Optional(Nullable(Int())),
			"concurrent_deployment_syncs":                         Optional(Nullable(Int())),
			"concurrent_endpoint_syncs":                           Optional(Nullable(Int())),
			"concurrent_namespace_syncs":                          Optional(Nullable(Int())),
			"concurrent_replicaset_syncs":                         Optional(Nullable(Int())),
			"concurrent_service_syncs":                            Optional(Nullable(Int())),
			"concurrent_resource_quota_syncs":                     Optional(Nullable(Int())),
			"concurrent_serviceaccount_token_syncs":               Optional(Nullable(Int())),
			"concurrent_rc_syncs":                                 Optional(Nullable(Int())),
			"authentication_kubeconfig":                           Optional(String()),
			"authorization_kubeconfig":                            Optional(String()),
			"authorization_always_allow_paths":                    Optional(List(String())),
			"external_cloud_volume_plugin":                        Optional(String()),
			"enable_profiling":                                    Optional(Nullable(Bool())),
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["node_cidr_mask_size"]; ok && in != nil {
		out.NodeCIDRMaskSize = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["configure_cloud_routes"]; ok && in != nil {
		out.ConfigureCloudRoutes = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.LeaderElectionConfiguration) *kops.LeaderElectionConfiguration { return &in }(func(in interface{}) kops.LeaderElectionConfiguration {
					return ExpandResourceLeaderElectionConfiguration(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	if in, ok := in["attach_detach_reconcile_sync_period"]; ok && in != nil {
		out.AttachDetachReconcileSyncPeriod = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["disable_attach_detach_reconcile_sync"]; ok && in != nil {
		out.DisableAttachDetachReconcileSync = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["terminated_pod_gc_threshold"]; ok && in != nil {
		out.TerminatedPodGCThreshold = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["node_monitor_period"]; ok && in != nil {
		out.NodeMonitorPeriod = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["node_monitor_grace_period"]; ok && in != nil {
		out.NodeMonitorGracePeriod = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["pod_eviction_timeout"]; ok && in != nil {
		out.PodEvictionTimeout = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["use_service_account_credentials"]; ok && in != nil {
		out.UseServiceAccountCredentials = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["horizontal_pod_autoscaler_sync_period"]; ok && in != nil {
		out.HorizontalPodAutoscalerSyncPeriod = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["horizontal_pod_autoscaler_downscale_delay"]; ok && in != nil {
		out.HorizontalPodAutoscalerDownscaleDelay = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["horizontal_pod_autoscaler_downscale_stabilization"]; ok && in != nil {
		out.HorizontalPodAutoscalerDownscaleStabilization = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["horizontal_pod_autoscaler_upscale_delay"]; ok && in != nil {
		out.HorizontalPodAutoscalerUpscaleDelay = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["horizontal_pod_autoscaler_initial_readiness_delay"]; ok && in != nil {
		out.HorizontalPodAutoscalerInitialReadinessDelay = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["horizontal_pod_autoscaler_cpu_initialization_period"]; ok && in != nil {
		out.HorizontalPodAutoscalerCPUInitializationPeriod = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["horizontal_pod_autoscaler_tolerance"]; ok && in != nil {
		out.HorizontalPodAutoscalerTolerance = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["horizontal_pod_autoscaler_use_rest_clients"]; ok && in != nil {
		out.HorizontalPodAutoscalerUseRestClients = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["experimental_cluster_signing_duration"]; ok && in != nil {
		out.ExperimentalClusterSigningDuration = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["kube_api_burst"]; ok && in != nil {
		out.KubeAPIBurst = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["concurrent_deployment_syncs"]; ok && in != nil {
		out.ConcurrentDeploymentSyncs = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["concurrent_endpoint_syncs"]; ok && in != nil {
		out.ConcurrentEndpointSyncs = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["concurrent_namespace_syncs"]; ok && in != nil {
		out.ConcurrentNamespaceSyncs = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["concurrent_replicaset_syncs"]; ok && in != nil {
		out.ConcurrentReplicasetSyncs = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["concurrent_service_syncs"]; ok && in != nil {
		out.ConcurrentServiceSyncs = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["concurrent_resource_quota_syncs"]; ok && in != nil {
		out.ConcurrentResourceQuotaSyncs = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["concurrent_serviceaccount_token_syncs"]; ok && in != nil {
		out.ConcurrentServiceaccountTokenSyncs = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["concurrent_rc_syncs"]; ok && in != nil {
		out.ConcurrentRcSyncs = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenResourceKubeControllerManagerConfigInto(in kops.KubeControllerManagerConfig, out map[string]interface{}) {
	out["master"] = func(in string) interface{} { return string(in) }(in.Master)
	out["log_level"] = func(in int32) interface{} { return int(in) }(in.LogLevel)
	out["service_account_private_key_file"] = func(in string) interface{} { return string(in) }(in.ServiceAccountPrivateKeyFile)
	out["image"] = func(in string) interface{} { return string(in) }(in.Image)
	out["cloud_provider"] = func(in string) interface{} { return string(in) }(in.CloudProvider)
	out["cluster_name"] = func(in string) interface{} { return string(in) }(in.ClusterName)
	out["cluster_cidr"] = func(in string) interface{} { return string(in) }(in.ClusterCIDR)
	out["allocate_node_cidrs"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.AllocateNodeCIDRs)
	out["node_cidr_mask_size"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.NodeCIDRMaskSize)
	out["configure_cloud_routes"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.ConfigureCloudRoutes)
	out["controllers"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.Controllers)
	out["cidr_allocator_type"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.CIDRAllocatorType)
	out["root_ca_file"] = func(in string) interface{} { return string(in) }(in.RootCAFile)
	out["leader_election"] = func(in *kops.LeaderElectionConfiguration) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.LeaderElectionConfiguration) interface{} {
			return FlattenResourceLeaderElectionConfiguration(in)
		}(*in)
	}(in.LeaderElection)
	out["attach_detach_reconcile_sync_period"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.AttachDetachReconcileSyncPeriod)
	out["disable_attach_detach_reconcile_sync"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.DisableAttachDetachReconcileSync)
	out["terminated_pod_gc_threshold"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.TerminatedPodGCThreshold)
	out["node_monitor_period"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.NodeMonitorPeriod)
	out["node_monitor_grace_period"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.NodeMonitorGracePeriod)
	out["pod_eviction_timeout"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.PodEvictionTimeout)
	out["use_service_account_credentials"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.UseServiceAccountCredentials)
	out["horizontal_pod_autoscaler_sync_period"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.HorizontalPodAutoscalerSyncPeriod)
	out["horizontal_pod_autoscaler_downscale_delay"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.HorizontalPodAutoscalerDownscaleDelay)
	out["horizontal_pod_autoscaler_downscale_stabilization"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.HorizontalPodAutoscalerDownscaleStabilization)
	out["horizontal_pod_autoscaler_upscale_delay"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.HorizontalPodAutoscalerUpscaleDelay)
	out["horizontal_pod_autoscaler_initial_readiness_delay"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.HorizontalPodAutoscalerInitialReadinessDelay)
	out["horizontal_pod_autoscaler_cpu_initialization_period"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.HorizontalPodAutoscalerCPUInitializationPeriod)
	out["horizontal_pod_autoscaler_tolerance"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.HorizontalPodAutoscalerTolerance)
	out["horizontal_pod_autoscaler_use_rest_clients"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.HorizontalPodAutoscalerUseRestClients)
	out["experimental_cluster_signing_duration"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.ExperimentalClusterSigningDuration)
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
	out["tls_cipher_suites"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.TLSCipherSuites)
	out["tls_min_version"] = func(in string) interface{} { return string(in) }(in.TLSMinVersion)
	out["min_resync_period"] = func(in string) interface{} { return string(in) }(in.MinResyncPeriod)
	out["kube_api_qps"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.KubeAPIQPS)
	out["kube_api_burst"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.KubeAPIBurst)
	out["concurrent_deployment_syncs"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.ConcurrentDeploymentSyncs)
	out["concurrent_endpoint_syncs"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.ConcurrentEndpointSyncs)
	out["concurrent_namespace_syncs"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.ConcurrentNamespaceSyncs)
	out["concurrent_replicaset_syncs"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.ConcurrentReplicasetSyncs)
	out["concurrent_service_syncs"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.ConcurrentServiceSyncs)
	out["concurrent_resource_quota_syncs"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.ConcurrentResourceQuotaSyncs)
	out["concurrent_serviceaccount_token_syncs"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.ConcurrentServiceaccountTokenSyncs)
	out["concurrent_rc_syncs"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.ConcurrentRcSyncs)
	out["authentication_kubeconfig"] = func(in string) interface{} { return string(in) }(in.AuthenticationKubeconfig)
	out["authorization_kubeconfig"] = func(in string) interface{} { return string(in) }(in.AuthorizationKubeconfig)
	out["authorization_always_allow_paths"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.AuthorizationAlwaysAllowPaths)
	out["external_cloud_volume_plugin"] = func(in string) interface{} { return string(in) }(in.ExternalCloudVolumePlugin)
	out["enable_profiling"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.EnableProfiling)
}

func FlattenResourceKubeControllerManagerConfig(in kops.KubeControllerManagerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKubeControllerManagerConfigInto(in, out)
	return out
}
