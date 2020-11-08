# kops_kube_scheduler_config

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `master` | String |  | Yes |  |
| `log_level` | Int |  | Yes |  |
| `image` | String |  | Yes |  |
| `leader_election` | [LeaderElectionConfiguration](./LeaderElectionConfiguration.md) |  | Yes |  |
| `use_policy_config_map` | Bool |  | Yes |  |
| `feature_gates` | Map(String) |  | Yes |  |
| `max_persistent_volumes` | Int |  | Yes |  |
| `qps` | Quantity |  | Yes |  |
| `burst` | Int |  | Yes |  |
| `enable_profiling` | Bool |  | Yes |  |
