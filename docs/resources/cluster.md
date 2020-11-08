# kops_cluster

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `name` | String | Yes |  |  |
| `admin_ssh_key` | String | Yes |  |  |
| `channel` | String |  | Yes |  |
| `addons` | List([AddonSpec](./AddonSpec.md)) |  | Yes |  |
| `config_base` | String |  | Yes | Yes |
| `cloud_provider` | String | Yes |  |  |
| `container_runtime` | String |  | Yes |  |
| `kubernetes_version` | String |  | Yes |  |
| `subnet` | List([ClusterSubnetSpec](./ClusterSubnetSpec.md)) | Yes |  |  |
| `project` | String |  | Yes |  |
| `master_public_name` | String |  | Yes | Yes |
| `master_internal_name` | String |  | Yes | Yes |
| `network_cidr` | String |  | Yes | Yes |
| `additional_network_cidrs` | List(String) |  | Yes |  |
| `network_id` | String | Yes |  |  |
| `topology` | [TopologySpec](./TopologySpec.md) | Yes |  |  |
| `secret_store` | String |  | Yes |  |
| `key_store` | String |  | Yes |  |
| `config_store` | String |  | Yes |  |
| `dns_zone` | String |  | Yes |  |
| `additional_sans` | List(String) |  | Yes |  |
| `cluster_dns_domain` | String |  | Yes |  |
| `service_cluster_ip_range` | String |  | Yes |  |
| `pod_cidr` | String |  | Yes |  |
| `non_masquerade_cidr` | String |  | Yes | Yes |
| `ssh_access` | List(String) |  | Yes |  |
| `node_port_access` | List(String) |  | Yes |  |
| `egress_proxy` | [EgressProxySpec](./EgressProxySpec.md) |  | Yes |  |
| `ssh_key_name` | String |  | Yes |  |
| `kubernetes_api_access` | List(String) |  | Yes |  |
| `isolate_masters` | Bool |  | Yes |  |
| `update_policy` | String |  | Yes |  |
| `external_policies` | Map(List(String)) |  | Yes |  |
| `additional_policies` | Map(String) |  | Yes |  |
| `file_assets` | List([FileAssetSpec](./FileAssetSpec.md)) |  | Yes |  |
| `etcd_cluster` | List([EtcdClusterSpec](./EtcdClusterSpec.md)) | Yes |  |  |
| `containerd` | [ContainerdConfig](./ContainerdConfig.md) |  | Yes |  |
| `docker` | [DockerConfig](./DockerConfig.md) |  | Yes |  |
| `kube_dns` | [KubeDNSConfig](./KubeDNSConfig.md) |  | Yes |  |
| `kube_api_server` | [KubeAPIServerConfig](./KubeAPIServerConfig.md) |  | Yes |  |
| `kube_controller_manager` | [KubeControllerManagerConfig](./KubeControllerManagerConfig.md) |  | Yes |  |
| `external_cloud_controller_manager` | [CloudControllerManagerConfig](./CloudControllerManagerConfig.md) |  | Yes |  |
| `kube_scheduler` | [KubeSchedulerConfig](./KubeSchedulerConfig.md) |  | Yes |  |
| `kube_proxy` | [KubeProxyConfig](./KubeProxyConfig.md) |  | Yes |  |
| `kubelet` | [KubeletConfigSpec](./KubeletConfigSpec.md) |  | Yes |  |
| `master_kubelet` | [KubeletConfigSpec](./KubeletConfigSpec.md) |  | Yes |  |
| `cloud_config` | [CloudConfiguration](./CloudConfiguration.md) |  | Yes |  |
| `external_dns` | [ExternalDNSConfig](./ExternalDNSConfig.md) |  | Yes |  |
| `networking` | [NetworkingSpec](./NetworkingSpec.md) | Yes |  |  |
| `api` | [AccessSpec](./AccessSpec.md) |  | Yes |  |
| `authentication` | [AuthenticationSpec](./AuthenticationSpec.md) |  | Yes |  |
| `authorization` | [AuthorizationSpec](./AuthorizationSpec.md) |  | Yes |  |
| `node_authorization` | [NodeAuthorizationSpec](./NodeAuthorizationSpec.md) |  | Yes |  |
| `cloud_labels` | Map(String) |  | Yes |  |
| `hooks` | List([HookSpec](./HookSpec.md)) |  | Yes |  |
| `assets` | [Assets](./Assets.md) |  | Yes |  |
| `iam` | [IAMSpec](./IAMSpec.md) |  | Yes | Yes |
| `encryption_config` | Bool |  | Yes |  |
| `disable_subnet_tags` | Bool |  | Yes |  |
| `use_host_certificates` | Bool |  | Yes |  |
| `sysctl_parameters` | List(String) |  | Yes |  |
| `rolling_update` | [RollingUpdate](./RollingUpdate.md) |  | Yes |  |
| `instance_group` | List([InstanceGroup](./InstanceGroup.md)) | Yes |  |  |
| `kube_config` | [KubeConfig](./KubeConfig.md) |  |  | Yes |
| `rolling_update_options` | [RollingUpdateOptions](./RollingUpdateOptions.md) |  | Yes |  |
| `validate_options` | [ValidateOptions](./ValidateOptions.md) |  | Yes |  |
