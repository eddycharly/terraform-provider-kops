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
			return string(ExpandString(in))
		}(in["master"]),
		LogLevel: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["log_level"]),
		ServiceAccountPrivateKeyFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["service_account_private_key_file"]),
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		CloudProvider: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cloud_provider"]),
		ClusterName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_name"]),
		ClusterCIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_cidr"]),
		AllocateNodeCIDRs: func(in interface{}) *bool {
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
		}(in["allocate_node_cidrs"]),
		NodeCIDRMaskSize: func(in interface{}) *int32 {
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
		}(in["node_cidr_mask_size"]),
		ConfigureCloudRoutes: func(in interface{}) *bool {
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
		}(in["configure_cloud_routes"]),
		Controllers: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["controllers"]),
		CIDRAllocatorType: func(in interface{}) *string {
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
		}(in["cidr_allocator_type"]),
		RootCAFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["root_ca_file"]),
		LeaderElection: func(in interface{}) *kops.LeaderElectionConfiguration {
			return func(in interface{}) *kops.LeaderElectionConfiguration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.LeaderElectionConfiguration) *kops.LeaderElectionConfiguration {
					return &in
				}(func(in interface{}) kops.LeaderElectionConfiguration {
					if in.([]interface{})[0] == nil {
						return kops.LeaderElectionConfiguration{}
					}
					return (ExpandLeaderElectionConfiguration(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["leader_election"]),
		AttachDetachReconcileSyncPeriod: func(in interface{}) *v1.Duration {
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
		}(in["attach_detach_reconcile_sync_period"]),
		DisableAttachDetachReconcileSync: func(in interface{}) *bool {
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
		}(in["disable_attach_detach_reconcile_sync"]),
		TerminatedPodGCThreshold: func(in interface{}) *int32 {
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
		}(in["terminated_pod_gc_threshold"]),
		NodeMonitorPeriod: func(in interface{}) *v1.Duration {
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
		}(in["node_monitor_period"]),
		NodeMonitorGracePeriod: func(in interface{}) *v1.Duration {
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
		}(in["node_monitor_grace_period"]),
		PodEvictionTimeout: func(in interface{}) *v1.Duration {
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
		}(in["pod_eviction_timeout"]),
		UseServiceAccountCredentials: func(in interface{}) *bool {
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
		}(in["use_service_account_credentials"]),
		HorizontalPodAutoscalerSyncPeriod: func(in interface{}) *v1.Duration {
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
		}(in["horizontal_pod_autoscaler_sync_period"]),
		HorizontalPodAutoscalerDownscaleDelay: func(in interface{}) *v1.Duration {
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
		}(in["horizontal_pod_autoscaler_downscale_delay"]),
		HorizontalPodAutoscalerDownscaleStabilization: func(in interface{}) *v1.Duration {
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
		}(in["horizontal_pod_autoscaler_downscale_stabilization"]),
		HorizontalPodAutoscalerUpscaleDelay: func(in interface{}) *v1.Duration {
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
		}(in["horizontal_pod_autoscaler_upscale_delay"]),
		HorizontalPodAutoscalerTolerance: func(in interface{}) *resource.Quantity {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					return &in
				}(ExpandQuantity(in))
			}(in)
		}(in["horizontal_pod_autoscaler_tolerance"]),
		HorizontalPodAutoscalerUseRestClients: func(in interface{}) *bool {
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
		}(in["horizontal_pod_autoscaler_use_rest_clients"]),
		ExperimentalClusterSigningDuration: func(in interface{}) *v1.Duration {
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
		}(in["experimental_cluster_signing_duration"]),
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
		MinResyncPeriod: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["min_resync_period"]),
		KubeAPIQPS: func(in interface{}) *resource.Quantity {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					return &in
				}(ExpandQuantity(in))
			}(in)
		}(in["kube_api_qps"]),
		KubeAPIBurst: func(in interface{}) *int32 {
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
		}(in["kube_api_burst"]),
		ConcurrentDeploymentSyncs: func(in interface{}) *int32 {
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
		}(in["concurrent_deployment_syncs"]),
		ConcurrentEndpointSyncs: func(in interface{}) *int32 {
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
		}(in["concurrent_endpoint_syncs"]),
		ConcurrentNamespaceSyncs: func(in interface{}) *int32 {
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
		}(in["concurrent_namespace_syncs"]),
		ConcurrentReplicasetSyncs: func(in interface{}) *int32 {
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
		}(in["concurrent_replicaset_syncs"]),
		ConcurrentServiceSyncs: func(in interface{}) *int32 {
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
		}(in["concurrent_service_syncs"]),
		ConcurrentResourceQuotaSyncs: func(in interface{}) *int32 {
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
		}(in["concurrent_resource_quota_syncs"]),
		ConcurrentServiceaccountTokenSyncs: func(in interface{}) *int32 {
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
		}(in["concurrent_serviceaccount_token_syncs"]),
		ConcurrentRcSyncs: func(in interface{}) *int32 {
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
		}(in["concurrent_rc_syncs"]),
		EnableProfiling: func(in interface{}) *bool {
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
		}(in["enable_profiling"]),
	}
}

func FlattenKubeControllerManagerConfig(in kops.KubeControllerManagerConfig) map[string]interface{} {
	return map[string]interface{}{
		"master": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Master),
		"log_level": func(in int32) interface{} {
			return FlattenInt(int(in))
		}(in.LogLevel),
		"service_account_private_key_file": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ServiceAccountPrivateKeyFile),
		"image": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Image),
		"cloud_provider": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.CloudProvider),
		"cluster_name": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ClusterName),
		"cluster_cidr": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ClusterCIDR),
		"allocate_node_cidrs": func(in *bool) interface{} {
			return func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
		}(in.AllocateNodeCIDRs),
		"node_cidr_mask_size": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.NodeCIDRMaskSize),
		"configure_cloud_routes": func(in *bool) interface{} {
			return func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
		}(in.ConfigureCloudRoutes),
		"controllers": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.Controllers),
		"cidr_allocator_type": func(in *string) interface{} {
			return func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
		}(in.CIDRAllocatorType),
		"root_ca_file": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.RootCAFile),
		"leader_election": func(in *kops.LeaderElectionConfiguration) interface{} {
			return func(in *kops.LeaderElectionConfiguration) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.LeaderElectionConfiguration) interface{} {
					return func(in kops.LeaderElectionConfiguration) []map[string]interface{} {
						return []map[string]interface{}{FlattenLeaderElectionConfiguration(in)}
					}(in)
				}(*in)
			}(in)
		}(in.LeaderElection),
		"attach_detach_reconcile_sync_period": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.AttachDetachReconcileSyncPeriod),
		"disable_attach_detach_reconcile_sync": func(in *bool) interface{} {
			return func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
		}(in.DisableAttachDetachReconcileSync),
		"terminated_pod_gc_threshold": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.TerminatedPodGCThreshold),
		"node_monitor_period": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.NodeMonitorPeriod),
		"node_monitor_grace_period": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.NodeMonitorGracePeriod),
		"pod_eviction_timeout": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.PodEvictionTimeout),
		"use_service_account_credentials": func(in *bool) interface{} {
			return func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
		}(in.UseServiceAccountCredentials),
		"horizontal_pod_autoscaler_sync_period": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.HorizontalPodAutoscalerSyncPeriod),
		"horizontal_pod_autoscaler_downscale_delay": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.HorizontalPodAutoscalerDownscaleDelay),
		"horizontal_pod_autoscaler_downscale_stabilization": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.HorizontalPodAutoscalerDownscaleStabilization),
		"horizontal_pod_autoscaler_upscale_delay": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.HorizontalPodAutoscalerUpscaleDelay),
		"horizontal_pod_autoscaler_tolerance": func(in *resource.Quantity) interface{} {
			return func(in *resource.Quantity) interface{} {
				if in == nil {
					return nil
				}
				return func(in resource.Quantity) interface{} {
					return FlattenQuantity(in)
				}(*in)
			}(in)
		}(in.HorizontalPodAutoscalerTolerance),
		"horizontal_pod_autoscaler_use_rest_clients": func(in *bool) interface{} {
			return func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
		}(in.HorizontalPodAutoscalerUseRestClients),
		"experimental_cluster_signing_duration": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.ExperimentalClusterSigningDuration),
		"feature_gates": func(in map[string]string) interface{} {
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
		}(in.FeatureGates),
		"tls_cipher_suites": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.TLSCipherSuites),
		"tls_min_version": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.TLSMinVersion),
		"min_resync_period": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.MinResyncPeriod),
		"kube_api_qps": func(in *resource.Quantity) interface{} {
			return func(in *resource.Quantity) interface{} {
				if in == nil {
					return nil
				}
				return func(in resource.Quantity) interface{} {
					return FlattenQuantity(in)
				}(*in)
			}(in)
		}(in.KubeAPIQPS),
		"kube_api_burst": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.KubeAPIBurst),
		"concurrent_deployment_syncs": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.ConcurrentDeploymentSyncs),
		"concurrent_endpoint_syncs": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.ConcurrentEndpointSyncs),
		"concurrent_namespace_syncs": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.ConcurrentNamespaceSyncs),
		"concurrent_replicaset_syncs": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.ConcurrentReplicasetSyncs),
		"concurrent_service_syncs": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.ConcurrentServiceSyncs),
		"concurrent_resource_quota_syncs": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.ConcurrentResourceQuotaSyncs),
		"concurrent_serviceaccount_token_syncs": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.ConcurrentServiceaccountTokenSyncs),
		"concurrent_rc_syncs": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.ConcurrentRcSyncs),
		"enable_profiling": func(in *bool) interface{} {
			return func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
		}(in.EnableProfiling),
	}
}
