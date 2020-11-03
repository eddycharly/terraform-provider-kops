# kops_cloud_controller_manager_config

| attribute | type | optional/required | computed |
| --- | --- | --- | --- |
| `master` | String | Optional |  |
| `log_level` | Int | Optional |  |
| `image` | String | Optional |  |
| `cloud_provider` | String | Optional |  |
| `cluster_name` | String | Optional |  |
| `cluster_cidr` | String | Optional |  |
| `allocate_node_cidrs` | Bool | Optional |  |
| `configure_cloud_routes` | Bool | Optional |  |
| `cidr_allocator_type` | String | Optional |  |
| `leader_election` | [LeaderElectionConfiguration](./LeaderElectionConfiguration.generated.md) | Optional |  |
| `use_service_account_credentials` | Bool | Optional |  |
