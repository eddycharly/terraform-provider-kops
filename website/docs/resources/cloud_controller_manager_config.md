# kops_cloud_controller_manager_config

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `master` | String |  | :white_check_mark: |  |
| `log_level` | Int |  | :white_check_mark: |  |
| `image` | String |  | :white_check_mark: |  |
| `cloud_provider` | String |  | :white_check_mark: |  |
| `cluster_name` | String |  | :white_check_mark: |  |
| `cluster_cidr` | String |  | :white_check_mark: |  |
| `allocate_node_cidrs` | Bool |  | :white_check_mark: |  |
| `configure_cloud_routes` | Bool |  | :white_check_mark: |  |
| `cidr_allocator_type` | String |  | :white_check_mark: |  |
| `leader_election` | [LeaderElectionConfiguration](./LeaderElectionConfiguration.md) |  | :white_check_mark: |  |
| `use_service_account_credentials` | Bool |  | :white_check_mark: |  |
