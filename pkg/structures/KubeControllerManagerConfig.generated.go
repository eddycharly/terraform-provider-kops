package structures

import (
	"reflect"

	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandKubeControllerManagerConfig(in map[string]interface{}) kops.KubeControllerManagerConfig {
	if in == nil {
		panic("expand KubeControllerManagerConfig failure, in is nil")
	}
	return kops.KubeControllerManagerConfig{
		Master: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["master"]),
		LogLevel: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			return value
		}(in["log_level"]),
		ServiceAccountPrivateKeyFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["service_account_private_key_file"]),
		Image: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["image"]),
		CloudProvider: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cloud_provider"]),
		ClusterName: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cluster_name"]),
		ClusterCIDR: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cluster_cidr"]),
		AllocateNodeCIDRs: func(in interface{}) *bool {
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
		}(in["allocate_node_cidrs"]),
		NodeCIDRMaskSize: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["node_cidr_mask_size"]),
		ConfigureCloudRoutes: func(in interface{}) *bool {
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
		}(in["configure_cloud_routes"]),
		Controllers: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["controllers"]),
		CIDRAllocatorType: func(in interface{}) *string {
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
		}(in["cidr_allocator_type"]),
		RootCAFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["root_ca_file"]),
		LeaderElection: func(in interface{}) *kops.LeaderElectionConfiguration {
			value := func(in interface{}) *kops.LeaderElectionConfiguration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.LeaderElectionConfiguration) *kops.LeaderElectionConfiguration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.LeaderElectionConfiguration {
					if in.([]interface{})[0] == nil {
						return kops.LeaderElectionConfiguration{}
					}
					return (ExpandLeaderElectionConfiguration(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["leader_election"]),
		AttachDetachReconcileSyncPeriod: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["attach_detach_reconcile_sync_period"]),
		DisableAttachDetachReconcileSync: func(in interface{}) *bool {
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
		}(in["disable_attach_detach_reconcile_sync"]),
		TerminatedPodGCThreshold: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["terminated_pod_gc_threshold"]),
		NodeMonitorPeriod: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["node_monitor_period"]),
		NodeMonitorGracePeriod: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["node_monitor_grace_period"]),
		PodEvictionTimeout: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["pod_eviction_timeout"]),
		UseServiceAccountCredentials: func(in interface{}) *bool {
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
		}(in["use_service_account_credentials"]),
		HorizontalPodAutoscalerSyncPeriod: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["horizontal_pod_autoscaler_sync_period"]),
		HorizontalPodAutoscalerDownscaleDelay: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["horizontal_pod_autoscaler_downscale_delay"]),
		HorizontalPodAutoscalerDownscaleStabilization: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["horizontal_pod_autoscaler_downscale_stabilization"]),
		HorizontalPodAutoscalerUpscaleDelay: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["horizontal_pod_autoscaler_upscale_delay"]),
		HorizontalPodAutoscalerTolerance: func(in interface{}) *resource.Quantity {
			value := func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in resource.Quantity) *resource.Quantity {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandQuantity(in))
				return tmp
			}(in)
			return value
		}(in["horizontal_pod_autoscaler_tolerance"]),
		HorizontalPodAutoscalerUseRestClients: func(in interface{}) *bool {
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
		}(in["horizontal_pod_autoscaler_use_rest_clients"]),
		ExperimentalClusterSigningDuration: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["experimental_cluster_signing_duration"]),
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
		MinResyncPeriod: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["min_resync_period"]),
		KubeAPIQPS: func(in interface{}) *resource.Quantity {
			value := func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in resource.Quantity) *resource.Quantity {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandQuantity(in))
				return tmp
			}(in)
			return value
		}(in["kube_api_qps"]),
		KubeAPIBurst: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["kube_api_burst"]),
		ConcurrentDeploymentSyncs: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["concurrent_deployment_syncs"]),
		ConcurrentEndpointSyncs: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["concurrent_endpoint_syncs"]),
		ConcurrentNamespaceSyncs: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["concurrent_namespace_syncs"]),
		ConcurrentReplicasetSyncs: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["concurrent_replicaset_syncs"]),
		ConcurrentServiceSyncs: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["concurrent_service_syncs"]),
		ConcurrentResourceQuotaSyncs: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["concurrent_resource_quota_syncs"]),
		ConcurrentServiceaccountTokenSyncs: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["concurrent_serviceaccount_token_syncs"]),
		ConcurrentRcSyncs: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["concurrent_rc_syncs"]),
		EnableProfiling: func(in interface{}) *bool {
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
		}(in["enable_profiling"]),
	}
}

func FlattenKubeControllerManagerConfig(in kops.KubeControllerManagerConfig) map[string]interface{} {
	return map[string]interface{}{
		"master": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Master),
		"log_level": func(in int32) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.LogLevel),
		"service_account_private_key_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ServiceAccountPrivateKeyFile),
		"image": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Image),
		"cloud_provider": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.CloudProvider),
		"cluster_name": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ClusterName),
		"cluster_cidr": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ClusterCIDR),
		"allocate_node_cidrs": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.AllocateNodeCIDRs),
		"node_cidr_mask_size": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.NodeCIDRMaskSize),
		"configure_cloud_routes": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.ConfigureCloudRoutes),
		"controllers": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.Controllers),
		"cidr_allocator_type": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.CIDRAllocatorType),
		"root_ca_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.RootCAFile),
		"leader_election": func(in *kops.LeaderElectionConfiguration) interface{} {
			value := func(in *kops.LeaderElectionConfiguration) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.LeaderElectionConfiguration) interface{} {
					return func(in kops.LeaderElectionConfiguration) []map[string]interface{} {
						return []map[string]interface{}{FlattenLeaderElectionConfiguration(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.LeaderElection),
		"attach_detach_reconcile_sync_period": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.AttachDetachReconcileSyncPeriod),
		"disable_attach_detach_reconcile_sync": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.DisableAttachDetachReconcileSync),
		"terminated_pod_gc_threshold": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.TerminatedPodGCThreshold),
		"node_monitor_period": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.NodeMonitorPeriod),
		"node_monitor_grace_period": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.NodeMonitorGracePeriod),
		"pod_eviction_timeout": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.PodEvictionTimeout),
		"use_service_account_credentials": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.UseServiceAccountCredentials),
		"horizontal_pod_autoscaler_sync_period": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.HorizontalPodAutoscalerSyncPeriod),
		"horizontal_pod_autoscaler_downscale_delay": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.HorizontalPodAutoscalerDownscaleDelay),
		"horizontal_pod_autoscaler_downscale_stabilization": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.HorizontalPodAutoscalerDownscaleStabilization),
		"horizontal_pod_autoscaler_upscale_delay": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.HorizontalPodAutoscalerUpscaleDelay),
		"horizontal_pod_autoscaler_tolerance": func(in *resource.Quantity) interface{} {
			value := func(in *resource.Quantity) interface{} {
				if in == nil {
					return nil
				}
				return func(in resource.Quantity) interface{} {
					return FlattenQuantity(in)
				}(*in)
			}(in)
			return value
		}(in.HorizontalPodAutoscalerTolerance),
		"horizontal_pod_autoscaler_use_rest_clients": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.HorizontalPodAutoscalerUseRestClients),
		"experimental_cluster_signing_duration": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.ExperimentalClusterSigningDuration),
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
		"min_resync_period": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.MinResyncPeriod),
		"kube_api_qps": func(in *resource.Quantity) interface{} {
			value := func(in *resource.Quantity) interface{} {
				if in == nil {
					return nil
				}
				return func(in resource.Quantity) interface{} {
					return FlattenQuantity(in)
				}(*in)
			}(in)
			return value
		}(in.KubeAPIQPS),
		"kube_api_burst": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.KubeAPIBurst),
		"concurrent_deployment_syncs": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.ConcurrentDeploymentSyncs),
		"concurrent_endpoint_syncs": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.ConcurrentEndpointSyncs),
		"concurrent_namespace_syncs": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.ConcurrentNamespaceSyncs),
		"concurrent_replicaset_syncs": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.ConcurrentReplicasetSyncs),
		"concurrent_service_syncs": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.ConcurrentServiceSyncs),
		"concurrent_resource_quota_syncs": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.ConcurrentResourceQuotaSyncs),
		"concurrent_serviceaccount_token_syncs": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.ConcurrentServiceaccountTokenSyncs),
		"concurrent_rc_syncs": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.ConcurrentRcSyncs),
		"enable_profiling": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.EnableProfiling),
	}
}
