# kops_kube_scheduler_config

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `master` | String |  | :white_check_mark: |  |
| `log_level` | Int |  | :white_check_mark: |  |
| `image` | String |  | :white_check_mark: |  |
| `leader_election` | [LeaderElectionConfiguration](./LeaderElectionConfiguration.md) |  | :white_check_mark: |  |
| `use_policy_config_map` | Bool |  | :white_check_mark: |  |
| `feature_gates` | Map(String) |  | :white_check_mark: |  |
| `max_persistent_volumes` | Int |  | :white_check_mark: |  |
| `qps` | Quantity |  | :white_check_mark: |  |
| `burst` | Int |  | :white_check_mark: |  |
| `enable_profiling` | Bool |  | :white_check_mark: |  |
