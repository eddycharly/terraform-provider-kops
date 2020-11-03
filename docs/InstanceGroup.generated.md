# kops_instance_group

| attribute | type | optional/required | computed |
| --- | --- | --- | --- |
| `name` | String | Required |  |
| `role` | String | Required |  |
| `image` | String | Optional | Computed |
| `min_size` | Int | Required |  |
| `max_size` | Int | Required |  |
| `machine_type` | String | Required |  |
| `root_volume_size` | Int | Optional |  |
| `root_volume_type` | String | Optional |  |
| `root_volume_iops` | Int | Optional |  |
| `root_volume_optimization` | Bool | Optional |  |
| `root_volume_delete_on_termination` | Bool | Optional |  |
| `root_volume_encryption` | Bool | Optional |  |
| `volumes` | List([VolumeSpec](./VolumeSpec.generated.md)) | Optional |  |
| `volume_mounts` | List([VolumeMountSpec](./VolumeMountSpec.generated.md)) | Optional |  |
| `subnets` | List(String) | Required |  |
| `zones` | List(String) | Optional |  |
| `hooks` | List([HookSpec](./HookSpec.generated.md)) | Optional |  |
| `max_price` | String | Optional |  |
| `spot_duration_in_minutes` | Int | Optional |  |
| `associate_public_ip` | Bool | Optional |  |
| `additional_security_groups` | List(String) | Optional |  |
| `cloud_labels` | Map(String) | Optional |  |
| `node_labels` | Map(String) | Optional |  |
| `file_assets` | List([FileAssetSpec](./FileAssetSpec.generated.md)) | Optional |  |
| `tenancy` | String | Optional |  |
| `kubelet` | [KubeletConfigSpec](./KubeletConfigSpec.generated.md) | Optional |  |
| `taints` | List(String) | Optional |  |
| `mixed_instances_policy` | [MixedInstancesPolicySpec](./MixedInstancesPolicySpec.generated.md) | Optional |  |
| `additional_user_data` | List([UserData](./UserData.generated.md)) | Optional |  |
| `suspend_processes` | List(String) | Optional |  |
| `external_load_balancers` | List([LoadBalancer](./LoadBalancer.generated.md)) | Optional |  |
| `detailed_instance_monitoring` | Bool | Optional |  |
| `iam` | [IAMProfileSpec](./IAMProfileSpec.generated.md) | Optional |  |
| `security_group_override` | String | Optional |  |
| `instance_protection` | Bool | Optional |  |
| `sysctl_parameters` | List(String) | Optional |  |
| `rolling_update` | [RollingUpdate](./RollingUpdate.generated.md) | Optional |  |
| `instance_interruption_behavior` | String | Optional |  |
