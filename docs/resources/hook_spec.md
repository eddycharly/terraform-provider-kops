# kops_hook_spec

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `name` | String | Yes |  |  |
| `disabled` | Bool |  | Yes |  |
| `roles` | List(String) |  | Yes |  |
| `requires` | List(String) |  | Yes |  |
| `before` | List(String) |  | Yes |  |
| `exec_container` | [ExecContainerAction](./ExecContainerAction.md) |  | Yes |  |
| `manifest` | String |  | Yes |  |
| `use_raw_manifest` | Bool |  | Yes |  |
