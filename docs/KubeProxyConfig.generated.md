# kops_kube_proxy_config

| attribute | type | optional/required | computed |
| --- | --- | --- | --- |
| `image` | String | Optional |  |
| `cpu_request` | String | Optional |  |
| `cpu_limit` | String | Optional |  |
| `memory_request` | String | Optional |  |
| `memory_limit` | String | Optional |  |
| `log_level` | Int | Optional |  |
| `cluster_cidr` | String | Optional |  |
| `hostname_override` | String | Optional |  |
| `bind_address` | String | Optional |  |
| `master` | String | Optional |  |
| `metrics_bind_address` | String | Optional |  |
| `enabled` | Bool | Optional |  |
| `proxy_mode` | String | Optional |  |
| `ip_vs_exclude_cidr_s` | List(String) | Optional |  |
| `ip_vs_min_sync_period` | Duration | Optional |  |
| `ip_vs_scheduler` | String | Optional |  |
| `ip_vs_sync_period` | Duration | Optional |  |
| `feature_gates` | Map(String) | Optional |  |
| `conntrack_max_per_core` | Int | Optional |  |
| `conntrack_min` | Int | Optional |  |
