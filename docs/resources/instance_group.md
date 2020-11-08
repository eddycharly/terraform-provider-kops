# kops_instance_group

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `name` | String | :white_check_mark: |  |  |
| `role` | String | :white_check_mark: |  |  |
| `image` | String |  | :white_check_mark: | :white_check_mark: |
| `min_size` | Int | :white_check_mark: |  |  |
| `max_size` | Int | :white_check_mark: |  |  |
| `machine_type` | String | :white_check_mark: |  |  |
| `root_volume_size` | Int |  | :white_check_mark: |  |
| `root_volume_type` | String |  | :white_check_mark: |  |
| `root_volume_iops` | Int |  | :white_check_mark: |  |
| `root_volume_optimization` | Bool |  | :white_check_mark: |  |
| `root_volume_delete_on_termination` | Bool |  | :white_check_mark: |  |
| `root_volume_encryption` | Bool |  | :white_check_mark: |  |
| `volumes` | List([VolumeSpec](./VolumeSpec.generated.md)) |  | :white_check_mark: |  |
| `volume_mounts` | List([VolumeMountSpec](./VolumeMountSpec.generated.md)) |  | :white_check_mark: |  |
| `subnets` | List(String) | :white_check_mark: |  |  |
| `zones` | List(String) |  | :white_check_mark: |  |
| `hooks` | List([HookSpec](./HookSpec.generated.md)) |  | :white_check_mark: |  |
| `max_price` | String |  | :white_check_mark: |  |
| `spot_duration_in_minutes` | Int |  | :white_check_mark: |  |
| `associate_public_ip` | Bool |  | :white_check_mark: |  |
| `additional_security_groups` | List(String) |  | :white_check_mark: |  |
| `cloud_labels` | Map(String) |  | :white_check_mark: |  |
| `node_labels` | Map(String) |  | :white_check_mark: |  |
| `file_assets` | List([FileAssetSpec](./FileAssetSpec.generated.md)) |  | :white_check_mark: |  |
| `tenancy` | String |  | :white_check_mark: |  |
| `kubelet` | [KubeletConfigSpec](./KubeletConfigSpec.generated.md) |  | :white_check_mark: |  |
| `taints` | List(String) |  | :white_check_mark: |  |
| `mixed_instances_policy` | [MixedInstancesPolicySpec](./MixedInstancesPolicySpec.generated.md) |  | :white_check_mark: |  |
| `additional_user_data` | List([UserData](./UserData.generated.md)) |  | :white_check_mark: |  |
| `suspend_processes` | List(String) |  | :white_check_mark: |  |
| `external_load_balancers` | List([LoadBalancer](./LoadBalancer.generated.md)) |  | :white_check_mark: |  |
| `detailed_instance_monitoring` | Bool |  | :white_check_mark: |  |
| `iam` | [IAMProfileSpec](./IAMProfileSpec.generated.md) |  | :white_check_mark: |  |
| `security_group_override` | String |  | :white_check_mark: |  |
| `instance_protection` | Bool |  | :white_check_mark: |  |
| `sysctl_parameters` | List(String) |  | :white_check_mark: |  |
| `rolling_update` | [RollingUpdate](./RollingUpdate.generated.md) |  | :white_check_mark: |  |
| `instance_interruption_behavior` | String |  | :white_check_mark: |  |
