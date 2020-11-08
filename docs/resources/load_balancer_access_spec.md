# kops_load_balancer_access_spec

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `type` | String | :white_check_mark: |  |  |
| `idle_timeout_seconds` | Int |  | :white_check_mark: |  |
| `security_group_override` | String |  | :white_check_mark: |  |
| `additional_security_groups` | List(String) |  | :white_check_mark: |  |
| `use_for_internal_api` | Bool |  | :white_check_mark: |  |
| `ssl_certificate` | String |  | :white_check_mark: |  |
| `cross_zone_load_balancing` | Bool |  | :white_check_mark: |  |
