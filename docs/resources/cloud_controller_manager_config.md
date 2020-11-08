# kops_cloud_controller_manager_config

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `master` | String |  | Yes |  |
| `log_level` | Int |  | Yes |  |
| `image` | String |  | Yes |  |
| `cloud_provider` | String |  | Yes |  |
| `cluster_name` | String |  | Yes |  |
| `cluster_cidr` | String |  | Yes |  |
| `allocate_node_cidrs` | Bool |  | Yes |  |
| `configure_cloud_routes` | Bool |  | Yes |  |
| `cidr_allocator_type` | String |  | Yes |  |
| `leader_election` | [LeaderElectionConfiguration](./LeaderElectionConfiguration.md) |  | Yes |  |
| `use_service_account_credentials` | Bool |  | Yes |  |
