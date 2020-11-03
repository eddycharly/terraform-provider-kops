# kops_load_balancer_access_spec

| attribute | type | optional/required | computed |
| --- | --- | --- | --- |
| `type` | String | Required |  |
| `idle_timeout_seconds` | Int | Optional |  |
| `security_group_override` | String | Optional |  |
| `additional_security_groups` | List(String) | Optional |  |
| `use_for_internal_api` | Bool | Optional |  |
| `ssl_certificate` | String | Optional |  |
| `cross_zone_load_balancing` | Bool | Optional |  |
