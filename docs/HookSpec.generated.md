# kops_hook_spec

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `name` | String | :white_check_mark: |  |  |
| `disabled` | Bool |  | :white_check_mark: |  |
| `roles` | List(String) |  | :white_check_mark: |  |
| `requires` | List(String) |  | :white_check_mark: |  |
| `before` | List(String) |  | :white_check_mark: |  |
| `exec_container` | [ExecContainerAction](./ExecContainerAction.generated.md) |  | :white_check_mark: |  |
| `manifest` | String |  | :white_check_mark: |  |
| `use_raw_manifest` | Bool |  | :white_check_mark: |  |
