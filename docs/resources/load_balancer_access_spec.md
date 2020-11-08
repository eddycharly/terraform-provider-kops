# kops_load_balancer_access_spec

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `type` | String | Yes |  |  |
| `idle_timeout_seconds` | Int |  | Yes |  |
| `security_group_override` | String |  | Yes |  |
| `additional_security_groups` | List(String) |  | Yes |  |
| `use_for_internal_api` | Bool |  | Yes |  |
| `ssl_certificate` | String |  | Yes |  |
| `cross_zone_load_balancing` | Bool |  | Yes |  |
