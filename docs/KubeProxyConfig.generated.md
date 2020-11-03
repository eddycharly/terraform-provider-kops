# kops_kube_proxy_config

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `image` | String |  | :white_check_mark: |  |
| `cpu_request` | String |  | :white_check_mark: |  |
| `cpu_limit` | String |  | :white_check_mark: |  |
| `memory_request` | String |  | :white_check_mark: |  |
| `memory_limit` | String |  | :white_check_mark: |  |
| `log_level` | Int |  | :white_check_mark: |  |
| `cluster_cidr` | String |  | :white_check_mark: |  |
| `hostname_override` | String |  | :white_check_mark: |  |
| `bind_address` | String |  | :white_check_mark: |  |
| `master` | String |  | :white_check_mark: |  |
| `metrics_bind_address` | String |  | :white_check_mark: |  |
| `enabled` | Bool |  | :white_check_mark: |  |
| `proxy_mode` | String |  | :white_check_mark: |  |
| `ip_vs_exclude_cidr_s` | List(String) |  | :white_check_mark: |  |
| `ip_vs_min_sync_period` | Duration |  | :white_check_mark: |  |
| `ip_vs_scheduler` | String |  | :white_check_mark: |  |
| `ip_vs_sync_period` | Duration |  | :white_check_mark: |  |
| `feature_gates` | Map(String) |  | :white_check_mark: |  |
| `conntrack_max_per_core` | Int |  | :white_check_mark: |  |
| `conntrack_min` | Int |  | :white_check_mark: |  |
