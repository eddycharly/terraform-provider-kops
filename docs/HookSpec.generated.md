# kops_hook_spec

| attribute | type | optional/required | computed |
| --- | --- | --- | --- |
| `name` | String | Required |  |
| `disabled` | Bool | Optional |  |
| `roles` | List(String) | Optional |  |
| `requires` | List(String) | Optional |  |
| `before` | List(String) | Optional |  |
| `exec_container` | [ExecContainerAction](./ExecContainerAction.generated.md) | Optional |  |
| `manifest` | String | Optional |  |
| `use_raw_manifest` | Bool | Optional |  |
