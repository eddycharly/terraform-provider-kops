package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func KubeControllerManagerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"master":                                    OptionalString(),
			"log_level":                                 OptionalInt(),
			"service_account_private_key_file":          OptionalString(),
			"image":                                     OptionalString(),
			"cloud_provider":                            OptionalString(),
			"cluster_name":                              OptionalString(),
			"cluster_cidr":                              OptionalString(),
			"allocate_node_cidrs":                       OptionalBool(),
			"node_cidr_mask_size":                       OptionalInt(),
			"configure_cloud_routes":                    OptionalBool(),
			"controllers":                               OptionalList(String()),
			"cidr_allocator_type":                       OptionalString(),
			"root_ca_file":                              OptionalString(),
			"leader_election":                           OptionalStruct(LeaderElectionConfiguration()),
			"attach_detach_reconcile_sync_period":       OptionalDuration(),
			"disable_attach_detach_reconcile_sync":      OptionalBool(),
			"terminated_pod_gc_threshold":               OptionalInt(),
			"node_monitor_period":                       OptionalDuration(),
			"node_monitor_grace_period":                 OptionalDuration(),
			"pod_eviction_timeout":                      OptionalDuration(),
			"use_service_account_credentials":           OptionalBool(),
			"horizontal_pod_autoscaler_sync_period":     OptionalDuration(),
			"horizontal_pod_autoscaler_downscale_delay": OptionalDuration(),
			"horizontal_pod_autoscaler_downscale_stabilization": OptionalDuration(),
			"horizontal_pod_autoscaler_upscale_delay":           OptionalDuration(),
			"horizontal_pod_autoscaler_tolerance":               OptionalQuantity(),
			"horizontal_pod_autoscaler_use_rest_clients":        OptionalBool(),
			"experimental_cluster_signing_duration":             OptionalDuration(),
			"feature_gates":                                     OptionalMap(String()),
			"tls_cipher_suites":                                 OptionalList(String()),
			"tls_min_version":                                   OptionalString(),
			"min_resync_period":                                 OptionalString(),
			"kube_api_qps":                                      OptionalQuantity(),
			"kube_api_burst":                                    OptionalInt(),
			"concurrent_deployment_syncs":                       OptionalInt(),
			"concurrent_endpoint_syncs":                         OptionalInt(),
			"concurrent_namespace_syncs":                        OptionalInt(),
			"concurrent_replicaset_syncs":                       OptionalInt(),
			"concurrent_service_syncs":                          OptionalInt(),
			"concurrent_resource_quota_syncs":                   OptionalInt(),
			"concurrent_serviceaccount_token_syncs":             OptionalInt(),
			"concurrent_rc_syncs":                               OptionalInt(),
			"enable_profiling":                                  OptionalBool(),
		},
	}
}
