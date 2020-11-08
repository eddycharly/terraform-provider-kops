# kops_instance_group

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `name` | String | Yes |  |  |
| `role` | String | Yes |  |  |
| `image` | String |  | Yes | Yes |
| `min_size` | Int | Yes |  |  |
| `max_size` | Int | Yes |  |  |
| `machine_type` | String | Yes |  |  |
| `root_volume_size` | Int |  | Yes |  |
| `root_volume_type` | String |  | Yes |  |
| `root_volume_iops` | Int |  | Yes |  |
| `root_volume_optimization` | Bool |  | Yes |  |
| `root_volume_delete_on_termination` | Bool |  | Yes |  |
| `root_volume_encryption` | Bool |  | Yes |  |
| `volumes` | List([VolumeSpec](./VolumeSpec.md)) |  | Yes |  |
| `volume_mounts` | List([VolumeMountSpec](./VolumeMountSpec.md)) |  | Yes |  |
| `subnets` | List(String) | Yes |  |  |
| `zones` | List(String) |  | Yes |  |
| `hooks` | List([HookSpec](./HookSpec.md)) |  | Yes |  |
| `max_price` | String |  | Yes |  |
| `spot_duration_in_minutes` | Int |  | Yes |  |
| `associate_public_ip` | Bool |  | Yes |  |
| `additional_security_groups` | List(String) |  | Yes |  |
| `cloud_labels` | Map(String) |  | Yes |  |
| `node_labels` | Map(String) |  | Yes |  |
| `file_assets` | List([FileAssetSpec](./FileAssetSpec.md)) |  | Yes |  |
| `tenancy` | String |  | Yes |  |
| `kubelet` | [KubeletConfigSpec](./KubeletConfigSpec.md) |  | Yes |  |
| `taints` | List(String) |  | Yes |  |
| `mixed_instances_policy` | [MixedInstancesPolicySpec](./MixedInstancesPolicySpec.md) |  | Yes |  |
| `additional_user_data` | List([UserData](./UserData.md)) |  | Yes |  |
| `suspend_processes` | List(String) |  | Yes |  |
| `external_load_balancers` | List([LoadBalancer](./LoadBalancer.md)) |  | Yes |  |
| `detailed_instance_monitoring` | Bool |  | Yes |  |
| `iam` | [IAMProfileSpec](./IAMProfileSpec.md) |  | Yes |  |
| `security_group_override` | String |  | Yes |  |
| `instance_protection` | Bool |  | Yes |  |
| `sysctl_parameters` | List(String) |  | Yes |  |
| `rolling_update` | [RollingUpdate](./RollingUpdate.md) |  | Yes |  |
| `instance_interruption_behavior` | String |  | Yes |  |
