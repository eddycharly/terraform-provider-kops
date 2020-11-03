# kops_kube_scheduler_config

| attribute | type | optional/required | computed |
| --- | --- | --- | --- |
| `master` | String | Optional |  |
| `log_level` | Int | Optional |  |
| `image` | String | Optional |  |
| `leader_election` | [LeaderElectionConfiguration](./LeaderElectionConfiguration.generated.md) | Optional |  |
| `use_policy_config_map` | Bool | Optional |  |
| `feature_gates` | Map(String) | Optional |  |
| `max_persistent_volumes` | Int | Optional |  |
| `qps` | Quantity | Optional |  |
| `burst` | Int | Optional |  |
| `enable_profiling` | Bool | Optional |  |
