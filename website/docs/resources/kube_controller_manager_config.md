# kops_kube_controller_manager_config

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `master` | String |  | :white_check_mark: |  |
| `log_level` | Int |  | :white_check_mark: |  |
| `service_account_private_key_file` | String |  | :white_check_mark: |  |
| `image` | String |  | :white_check_mark: |  |
| `cloud_provider` | String |  | :white_check_mark: |  |
| `cluster_name` | String |  | :white_check_mark: |  |
| `cluster_cidr` | String |  | :white_check_mark: |  |
| `allocate_node_cidrs` | Bool |  | :white_check_mark: |  |
| `node_cidr_mask_size` | Int |  | :white_check_mark: |  |
| `configure_cloud_routes` | Bool |  | :white_check_mark: |  |
| `controllers` | List(String) |  | :white_check_mark: |  |
| `cidr_allocator_type` | String |  | :white_check_mark: |  |
| `root_ca_file` | String |  | :white_check_mark: |  |
| `leader_election` | [LeaderElectionConfiguration](./LeaderElectionConfiguration.md) |  | :white_check_mark: |  |
| `attach_detach_reconcile_sync_period` | Duration |  | :white_check_mark: |  |
| `disable_attach_detach_reconcile_sync` | Bool |  | :white_check_mark: |  |
| `terminated_pod_gc_threshold` | Int |  | :white_check_mark: |  |
| `node_monitor_period` | Duration |  | :white_check_mark: |  |
| `node_monitor_grace_period` | Duration |  | :white_check_mark: |  |
| `pod_eviction_timeout` | Duration |  | :white_check_mark: |  |
| `use_service_account_credentials` | Bool |  | :white_check_mark: |  |
| `horizontal_pod_autoscaler_sync_period` | Duration |  | :white_check_mark: |  |
| `horizontal_pod_autoscaler_downscale_delay` | Duration |  | :white_check_mark: |  |
| `horizontal_pod_autoscaler_downscale_stabilization` | Duration |  | :white_check_mark: |  |
| `horizontal_pod_autoscaler_upscale_delay` | Duration |  | :white_check_mark: |  |
| `horizontal_pod_autoscaler_tolerance` | Quantity |  | :white_check_mark: |  |
| `horizontal_pod_autoscaler_use_rest_clients` | Bool |  | :white_check_mark: |  |
| `experimental_cluster_signing_duration` | Duration |  | :white_check_mark: |  |
| `feature_gates` | Map(String) |  | :white_check_mark: |  |
| `tls_cipher_suites` | List(String) |  | :white_check_mark: |  |
| `tls_min_version` | String |  | :white_check_mark: |  |
| `min_resync_period` | String |  | :white_check_mark: |  |
| `kube_api_qps` | Quantity |  | :white_check_mark: |  |
| `kube_api_burst` | Int |  | :white_check_mark: |  |
| `concurrent_deployment_syncs` | Int |  | :white_check_mark: |  |
| `concurrent_endpoint_syncs` | Int |  | :white_check_mark: |  |
| `concurrent_namespace_syncs` | Int |  | :white_check_mark: |  |
| `concurrent_replicaset_syncs` | Int |  | :white_check_mark: |  |
| `concurrent_service_syncs` | Int |  | :white_check_mark: |  |
| `concurrent_resource_quota_syncs` | Int |  | :white_check_mark: |  |
| `concurrent_serviceaccount_token_syncs` | Int |  | :white_check_mark: |  |
| `concurrent_rc_syncs` | Int |  | :white_check_mark: |  |
| `enable_profiling` | Bool |  | :white_check_mark: |  |
