# kops_kube_dns_config

| attribute | type | optional/required | computed |
| --- | --- | --- | --- |
| `cache_max_size` | Int | Optional |  |
| `cache_max_concurrent` | Int | Optional |  |
| `core_dns_image` | String | Optional |  |
| `domain` | String | Optional |  |
| `external_core_file` | String | Optional |  |
| `image` | String | Optional |  |
| `replicas` | Int | Optional |  |
| `provider` | String | Optional |  |
| `server_ip` | String | Optional |  |
| `stub_domains` | Map(List(String)) | Optional |  |
| `upstream_nameservers` | List(String) | Optional |  |
| `memory_request` | Quantity | Optional |  |
| `cpu_request` | Quantity | Optional |  |
| `memory_limit` | Quantity | Optional |  |
| `node_local_dns` | [NodeLocalDNSConfig](./NodeLocalDNSConfig.generated.md) | Optional |  |
