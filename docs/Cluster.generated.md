# kops_cluster

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `name` | String | :white_check_mark: |  |  |
| `admin_ssh_key` | String | :white_check_mark: |  |  |
| `channel` | String |  | :white_check_mark: |  |
| `addons` | List([AddonSpec](./AddonSpec.generated.md)) |  | :white_check_mark: |  |
| `config_base` | String |  | :white_check_mark: | :white_check_mark: |
| `cloud_provider` | String | :white_check_mark: |  |  |
| `container_runtime` | String |  | :white_check_mark: |  |
| `kubernetes_version` | String |  | :white_check_mark: |  |
| `subnet` | List([ClusterSubnetSpec](./ClusterSubnetSpec.generated.md)) | :white_check_mark: |  |  |
| `project` | String |  | :white_check_mark: |  |
| `master_public_name` | String |  | :white_check_mark: | :white_check_mark: |
| `master_internal_name` | String |  | :white_check_mark: | :white_check_mark: |
| `network_cidr` | String |  | :white_check_mark: | :white_check_mark: |
| `additional_network_cidrs` | List(String) |  | :white_check_mark: |  |
| `network_id` | String | :white_check_mark: |  |  |
| `topology` | [TopologySpec](./TopologySpec.generated.md) | :white_check_mark: |  |  |
| `secret_store` | String |  | :white_check_mark: |  |
| `key_store` | String |  | :white_check_mark: |  |
| `config_store` | String |  | :white_check_mark: |  |
| `dns_zone` | String |  | :white_check_mark: |  |
| `additional_sans` | List(String) |  | :white_check_mark: |  |
| `cluster_dns_domain` | String |  | :white_check_mark: |  |
| `service_cluster_ip_range` | String |  | :white_check_mark: |  |
| `pod_cidr` | String |  | :white_check_mark: |  |
| `non_masquerade_cidr` | String |  | :white_check_mark: | :white_check_mark: |
| `ssh_access` | List(String) |  | :white_check_mark: |  |
| `node_port_access` | List(String) |  | :white_check_mark: |  |
| `egress_proxy` | [EgressProxySpec](./EgressProxySpec.generated.md) |  | :white_check_mark: |  |
| `ssh_key_name` | String |  | :white_check_mark: |  |
| `kubernetes_api_access` | List(String) |  | :white_check_mark: |  |
| `isolate_masters` | Bool |  | :white_check_mark: |  |
| `update_policy` | String |  | :white_check_mark: |  |
| `external_policies` | Map(List(String)) |  | :white_check_mark: |  |
| `additional_policies` | Map(String) |  | :white_check_mark: |  |
| `file_assets` | List([FileAssetSpec](./FileAssetSpec.generated.md)) |  | :white_check_mark: |  |
| `etcd_cluster` | List([EtcdClusterSpec](./EtcdClusterSpec.generated.md)) | :white_check_mark: |  |  |
| `containerd` | [ContainerdConfig](./ContainerdConfig.generated.md) |  | :white_check_mark: |  |
| `docker` | [DockerConfig](./DockerConfig.generated.md) |  | :white_check_mark: |  |
| `kube_dns` | [KubeDNSConfig](./KubeDNSConfig.generated.md) |  | :white_check_mark: |  |
| `kube_api_server` | [KubeAPIServerConfig](./KubeAPIServerConfig.generated.md) |  | :white_check_mark: |  |
| `kube_controller_manager` | [KubeControllerManagerConfig](./KubeControllerManagerConfig.generated.md) |  | :white_check_mark: |  |
| `external_cloud_controller_manager` | [CloudControllerManagerConfig](./CloudControllerManagerConfig.generated.md) |  | :white_check_mark: |  |
| `kube_scheduler` | [KubeSchedulerConfig](./KubeSchedulerConfig.generated.md) |  | :white_check_mark: |  |
| `kube_proxy` | [KubeProxyConfig](./KubeProxyConfig.generated.md) |  | :white_check_mark: |  |
| `kubelet` | [KubeletConfigSpec](./KubeletConfigSpec.generated.md) |  | :white_check_mark: |  |
| `master_kubelet` | [KubeletConfigSpec](./KubeletConfigSpec.generated.md) |  | :white_check_mark: |  |
| `cloud_config` | [CloudConfiguration](./CloudConfiguration.generated.md) |  | :white_check_mark: |  |
| `external_dns` | [ExternalDNSConfig](./ExternalDNSConfig.generated.md) |  | :white_check_mark: |  |
| `networking` | [NetworkingSpec](./NetworkingSpec.generated.md) | :white_check_mark: |  |  |
| `api` | [AccessSpec](./AccessSpec.generated.md) |  | :white_check_mark: |  |
| `authentication` | [AuthenticationSpec](./AuthenticationSpec.generated.md) |  | :white_check_mark: |  |
| `authorization` | [AuthorizationSpec](./AuthorizationSpec.generated.md) |  | :white_check_mark: |  |
| `node_authorization` | [NodeAuthorizationSpec](./NodeAuthorizationSpec.generated.md) |  | :white_check_mark: |  |
| `cloud_labels` | Map(String) |  | :white_check_mark: |  |
| `hooks` | List([HookSpec](./HookSpec.generated.md)) |  | :white_check_mark: |  |
| `assets` | [Assets](./Assets.generated.md) |  | :white_check_mark: |  |
| `iam` | [IAMSpec](./IAMSpec.generated.md) |  | :white_check_mark: | :white_check_mark: |
| `encryption_config` | Bool |  | :white_check_mark: |  |
| `disable_subnet_tags` | Bool |  | :white_check_mark: |  |
| `use_host_certificates` | Bool |  | :white_check_mark: |  |
| `sysctl_parameters` | List(String) |  | :white_check_mark: |  |
| `rolling_update` | [RollingUpdate](./RollingUpdate.generated.md) |  | :white_check_mark: |  |
| `instance_group` | List([InstanceGroup](./InstanceGroup.generated.md)) | :white_check_mark: |  |  |
| `kube_config` | [KubeConfig](./KubeConfig.generated.md) |  |  | :white_check_mark: |
| `rolling_update_options` | [RollingUpdateOptions](./RollingUpdateOptions.generated.md) |  | :white_check_mark: |  |
