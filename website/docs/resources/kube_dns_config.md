# kops_kube_dns_config

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `cache_max_size` | Int |  | :white_check_mark: |  |
| `cache_max_concurrent` | Int |  | :white_check_mark: |  |
| `core_dns_image` | String |  | :white_check_mark: |  |
| `domain` | String |  | :white_check_mark: |  |
| `external_core_file` | String |  | :white_check_mark: |  |
| `image` | String |  | :white_check_mark: |  |
| `replicas` | Int |  | :white_check_mark: |  |
| `provider` | String |  | :white_check_mark: |  |
| `server_ip` | String |  | :white_check_mark: |  |
| `stub_domains` | Map(List(String)) |  | :white_check_mark: |  |
| `upstream_nameservers` | List(String) |  | :white_check_mark: |  |
| `memory_request` | Quantity |  | :white_check_mark: |  |
| `cpu_request` | Quantity |  | :white_check_mark: |  |
| `memory_limit` | Quantity |  | :white_check_mark: |  |
| `node_local_dns` | [NodeLocalDNSConfig](./NodeLocalDNSConfig.md) |  | :white_check_mark: |  |
