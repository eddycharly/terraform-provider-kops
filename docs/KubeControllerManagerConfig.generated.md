# kops_kube_controller_manager_config

| attribute | type | optional/required | computed |
| --- | --- | --- | --- |
| `master` | String | Optional |  |
| `log_level` | Int | Optional |  |
| `service_account_private_key_file` | String | Optional |  |
| `image` | String | Optional |  |
| `cloud_provider` | String | Optional |  |
| `cluster_name` | String | Optional |  |
| `cluster_cidr` | String | Optional |  |
| `allocate_node_cidrs` | Bool | Optional |  |
| `node_cidr_mask_size` | Int | Optional |  |
| `configure_cloud_routes` | Bool | Optional |  |
| `controllers` | List(String) | Optional |  |
| `cidr_allocator_type` | String | Optional |  |
| `root_ca_file` | String | Optional |  |
| `leader_election` | [LeaderElectionConfiguration](./LeaderElectionConfiguration.generated.md) | Optional |  |
| `attach_detach_reconcile_sync_period` | Duration | Optional |  |
| `disable_attach_detach_reconcile_sync` | Bool | Optional |  |
| `terminated_pod_gc_threshold` | Int | Optional |  |
| `node_monitor_period` | Duration | Optional |  |
| `node_monitor_grace_period` | Duration | Optional |  |
| `pod_eviction_timeout` | Duration | Optional |  |
| `use_service_account_credentials` | Bool | Optional |  |
| `horizontal_pod_autoscaler_sync_period` | Duration | Optional |  |
| `horizontal_pod_autoscaler_downscale_delay` | Duration | Optional |  |
| `horizontal_pod_autoscaler_downscale_stabilization` | Duration | Optional |  |
| `horizontal_pod_autoscaler_upscale_delay` | Duration | Optional |  |
| `horizontal_pod_autoscaler_tolerance` | Quantity | Optional |  |
| `horizontal_pod_autoscaler_use_rest_clients` | Bool | Optional |  |
| `experimental_cluster_signing_duration` | Duration | Optional |  |
| `feature_gates` | Map(String) | Optional |  |
| `tls_cipher_suites` | List(String) | Optional |  |
| `tls_min_version` | String | Optional |  |
| `min_resync_period` | String | Optional |  |
| `kube_api_qps` | Quantity | Optional |  |
| `kube_api_burst` | Int | Optional |  |
| `concurrent_deployment_syncs` | Int | Optional |  |
| `concurrent_endpoint_syncs` | Int | Optional |  |
| `concurrent_namespace_syncs` | Int | Optional |  |
| `concurrent_replicaset_syncs` | Int | Optional |  |
| `concurrent_service_syncs` | Int | Optional |  |
| `concurrent_resource_quota_syncs` | Int | Optional |  |
| `concurrent_serviceaccount_token_syncs` | Int | Optional |  |
| `concurrent_rc_syncs` | Int | Optional |  |
| `enable_profiling` | Bool | Optional |  |
