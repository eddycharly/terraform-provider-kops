# kops_kube_dns_config

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `cache_max_size` | Int |  | Yes |  |
| `cache_max_concurrent` | Int |  | Yes |  |
| `core_dns_image` | String |  | Yes |  |
| `domain` | String |  | Yes |  |
| `external_core_file` | String |  | Yes |  |
| `image` | String |  | Yes |  |
| `replicas` | Int |  | Yes |  |
| `provider` | String |  | Yes |  |
| `server_ip` | String |  | Yes |  |
| `stub_domains` | Map(List(String)) |  | Yes |  |
| `upstream_nameservers` | List(String) |  | Yes |  |
| `memory_request` | Quantity |  | Yes |  |
| `cpu_request` | Quantity |  | Yes |  |
| `memory_limit` | Quantity |  | Yes |  |
| `node_local_dns` | [NodeLocalDNSConfig](./NodeLocalDNSConfig.md) |  | Yes |  |
