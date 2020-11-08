# kops_kube_proxy_config

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `image` | String |  | Yes |  |
| `cpu_request` | String |  | Yes |  |
| `cpu_limit` | String |  | Yes |  |
| `memory_request` | String |  | Yes |  |
| `memory_limit` | String |  | Yes |  |
| `log_level` | Int |  | Yes |  |
| `cluster_cidr` | String |  | Yes |  |
| `hostname_override` | String |  | Yes |  |
| `bind_address` | String |  | Yes |  |
| `master` | String |  | Yes |  |
| `metrics_bind_address` | String |  | Yes |  |
| `enabled` | Bool |  | Yes |  |
| `proxy_mode` | String |  | Yes |  |
| `ip_vs_exclude_cidr_s` | List(String) |  | Yes |  |
| `ip_vs_min_sync_period` | Duration |  | Yes |  |
| `ip_vs_scheduler` | String |  | Yes |  |
| `ip_vs_sync_period` | Duration |  | Yes |  |
| `feature_gates` | Map(String) |  | Yes |  |
| `conntrack_max_per_core` | Int |  | Yes |  |
| `conntrack_min` | Int |  | Yes |  |
