# kops_kube_controller_manager_config

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `master` | String |  | Yes |  |
| `log_level` | Int |  | Yes |  |
| `service_account_private_key_file` | String |  | Yes |  |
| `image` | String |  | Yes |  |
| `cloud_provider` | String |  | Yes |  |
| `cluster_name` | String |  | Yes |  |
| `cluster_cidr` | String |  | Yes |  |
| `allocate_node_cidrs` | Bool |  | Yes |  |
| `node_cidr_mask_size` | Int |  | Yes |  |
| `configure_cloud_routes` | Bool |  | Yes |  |
| `controllers` | List(String) |  | Yes |  |
| `cidr_allocator_type` | String |  | Yes |  |
| `root_ca_file` | String |  | Yes |  |
| `leader_election` | [LeaderElectionConfiguration](./LeaderElectionConfiguration.md) |  | Yes |  |
| `attach_detach_reconcile_sync_period` | Duration |  | Yes |  |
| `disable_attach_detach_reconcile_sync` | Bool |  | Yes |  |
| `terminated_pod_gc_threshold` | Int |  | Yes |  |
| `node_monitor_period` | Duration |  | Yes |  |
| `node_monitor_grace_period` | Duration |  | Yes |  |
| `pod_eviction_timeout` | Duration |  | Yes |  |
| `use_service_account_credentials` | Bool |  | Yes |  |
| `horizontal_pod_autoscaler_sync_period` | Duration |  | Yes |  |
| `horizontal_pod_autoscaler_downscale_delay` | Duration |  | Yes |  |
| `horizontal_pod_autoscaler_downscale_stabilization` | Duration |  | Yes |  |
| `horizontal_pod_autoscaler_upscale_delay` | Duration |  | Yes |  |
| `horizontal_pod_autoscaler_tolerance` | Quantity |  | Yes |  |
| `horizontal_pod_autoscaler_use_rest_clients` | Bool |  | Yes |  |
| `experimental_cluster_signing_duration` | Duration |  | Yes |  |
| `feature_gates` | Map(String) |  | Yes |  |
| `tls_cipher_suites` | List(String) |  | Yes |  |
| `tls_min_version` | String |  | Yes |  |
| `min_resync_period` | String |  | Yes |  |
| `kube_api_qps` | Quantity |  | Yes |  |
| `kube_api_burst` | Int |  | Yes |  |
| `concurrent_deployment_syncs` | Int |  | Yes |  |
| `concurrent_endpoint_syncs` | Int |  | Yes |  |
| `concurrent_namespace_syncs` | Int |  | Yes |  |
| `concurrent_replicaset_syncs` | Int |  | Yes |  |
| `concurrent_service_syncs` | Int |  | Yes |  |
| `concurrent_resource_quota_syncs` | Int |  | Yes |  |
| `concurrent_serviceaccount_token_syncs` | Int |  | Yes |  |
| `concurrent_rc_syncs` | Int |  | Yes |  |
| `enable_profiling` | Bool |  | Yes |  |
