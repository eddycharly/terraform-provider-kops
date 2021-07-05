# kops_instance_group

Provides a kOps cluster instance group.

## Example usage

```hcl
resource "kops_cluster" "cluster" {
  name                 = "cluster.example.com"

  // ....
}

resource "kops_instance_group" "master-0" {
  cluster_name = kops_cluster.cluster.name
  name         = "master-0"
  role         = "Master"
  min_size     = 1
  max_size     = 1
  machine_type = "t3.medium"
  subnets      = ["private-0"]
  depends_on   = [kops_cluster.cluster]
}

resource "kops_instance_group" "master-1" {
  cluster_name = kops_cluster.cluster.name
  name         = "master-1"
  role         = "Master"
  min_size     = 1
  max_size     = 1
  machine_type = "t3.medium"
  subnets      = ["private-1"]
  depends_on   = [kops_cluster.cluster]
}

resource "kops_instance_group" "master-2" {
  cluster_name = kops_cluster.cluster.name
  name         = "master-2"
  role         = "Master"
  min_size     = 1
  max_size     = 1
  machine_type = "t3.medium"
  subnets      = ["private-2"]
  depends_on   = [kops_cluster.cluster]
}
```

## Argument Reference

The following arguments are supported:
- `revision` - (Computed) - Int - Revision is incremented every time the resource changes, this is useful for triggering cluster updater.
- `cluster_name` - (Required) - (Force new) - String - ClusterName defines the cluster name the instance group belongs to.
- `name` - (Required) - (Force new) - String - Name defines the instance group name.
- `role` - (Required) - String - Type determines the role of instances in this instance group: masters or nodes.
- `image` - (Optional) - (Computed) - String - Image is the instance (ami etc) we should use.
- `min_size` - (Required) - Int - MinSize is the minimum size of the pool.
- `max_size` - (Required) - Int - MaxSize is the maximum size of the pool.
- `autoscale` - (Optional) - Bool - Autoscale determines if autoscaling will be enabled for this instance group if cluster autoscaler is enabled.
- `machine_type` - (Required) - String - MachineType is the instance class.
- `root_volume_size` - (Optional) - Int - RootVolumeSize is the size of the EBS root volume to use, in GB.
- `root_volume_type` - (Optional) - String - RootVolumeType is the type of the EBS root volume to use (e.g. gp2).
- `root_volume_iops` - (Optional) - Int - RootVolumeIops is the provisioned IOPS when the volume type is io1, io2 or gp3 (AWS only).
- `root_volume_throughput` - (Optional) - Int - RootVolumeThroughput is the volume throughput in MBps when the volume type is gp3 (AWS only).
- `root_volume_optimization` - (Optional) - Bool - RootVolumeOptimization enables EBS optimization for an instance.
- `root_volume_delete_on_termination` - (Optional) - Bool - RootVolumeDeleteOnTermination configures root volume retention policy upon instance termination.<br />The root volume is deleted by default. Cluster deletion does not remove retained root volumes.<br />NOTE: This setting applies only to the Launch Configuration and does not affect Launch Templates.
- `root_volume_encryption` - (Optional) - Bool - RootVolumeEncryption enables EBS root volume encryption for an instance.
- `root_volume_encryption_key` - (Optional) - String - RootVolumeEncryptionKey provides the key identifier for root volume encryption.
- `volumes` - (Optional) - List([volume_spec](#volume_spec)) - Volumes is a collection of additional volumes to create for instances within this instance group.
- `volume_mounts` - (Optional) - List([volume_mount_spec](#volume_mount_spec)) - VolumeMounts a collection of volume mounts.
- `subnets` - (Required) - List(String) - Subnets is the names of the Subnets (as specified in the Cluster) where machines in this instance group should be placed.
- `zones` - (Optional) - List(String) - Zones is the names of the Zones where machines in this instance group should be placed<br />This is needed for regional subnets (e.g. GCE), to restrict placement to particular zones.
- `hooks` - (Optional) - List([hook_spec](#hook_spec)) - Hooks is a list of hooks for this instance group, note: these can override the cluster wide ones if required.
- `max_price` - (Optional) - String - MaxPrice indicates this is a spot-pricing group, with the specified value as our max-price bid.
- `spot_duration_in_minutes` - (Optional) - Int - SpotDurationInMinutes reserves a spot block for the period specified.
- `cpu_credits` - (Optional) - String - CPUCredits is the credit option for CPU Usage on burstable instance types (AWS only).
- `associate_public_ip` - (Optional) - Bool - AssociatePublicIP is true if we want instances to have a public IP.
- `additional_security_groups` - (Optional) - List(String) - AdditionalSecurityGroups attaches additional security groups (e.g. i-123456).
- `cloud_labels` - (Optional) - Map(String) - CloudLabels indicates the labels for instances in this group, at the AWS level.
- `node_labels` - (Optional) - Map(String) - NodeLabels indicates the kubernetes labels for nodes in this instance group.
- `file_assets` - (Optional) - List([file_asset_spec](#file_asset_spec)) - FileAssets is a collection of file assets for this instance group.
- `tenancy` - (Optional) - String - Describes the tenancy of this instance group. Can be either default or dedicated. Currently only applies to AWS.
- `kubelet` - (Optional) - [kubelet_config_spec](#kubelet_config_spec) - Kubelet overrides kubelet config from the ClusterSpec.
- `taints` - (Optional) - List(String) - Taints indicates the kubernetes taints for nodes in this instance group.
- `mixed_instances_policy` - (Optional) - [mixed_instances_policy_spec](#mixed_instances_policy_spec) - MixedInstancesPolicy defined a optional backing of an AWS ASG by a EC2 Fleet (AWS Only).
- `additional_user_data` - (Optional) - List([user_data](#user_data)) - AdditionalUserData is any additional user-data to be passed to the host.
- `suspend_processes` - (Optional) - List(String) - SuspendProcesses disables the listed Scaling Policies.
- `external_load_balancers` - (Optional) - List([load_balancer](#load_balancer)) - ExternalLoadBalancers define loadbalancers that should be attached to this instance group.
- `detailed_instance_monitoring` - (Optional) - Bool - DetailedInstanceMonitoring defines if detailed-monitoring is enabled (AWS only).
- `iam` - (Optional) - [iam_profile_spec](#iam_profile_spec) - IAMProfileSpec defines the identity of the cloud group IAM profile (AWS only).
- `security_group_override` - (Optional) - String - SecurityGroupOverride overrides the default security group created by Kops for this IG (AWS only).
- `instance_protection` - (Optional) - Bool - InstanceProtection makes new instances in an autoscaling group protected from scale in.
- `sysctl_parameters` - (Optional) - List(String) - SysctlParameters will configure kernel parameters using sysctl(8). When<br />specified, each parameter must follow the form variable=value, the way<br />it would appear in sysctl.conf.
- `rolling_update` - (Optional) - [rolling_update](#rolling_update) - RollingUpdate defines the rolling-update behavior.
- `instance_interruption_behavior` - (Optional) - String - InstanceInterruptionBehavior defines if a spot instance should be terminated, hibernated,<br />or stopped after interruption.
- `compress_user_data` - (Optional) - Bool - CompressUserData compresses parts of the user data to save space.
- `instance_metadata` - (Optional) - [instance_metadata_options](#instance_metadata_options) - InstanceMetadata defines the EC2 instance metadata service options (AWS Only).
- `update_policy` - (Optional) - String - UpdatePolicy determines the policy for applying upgrades automatically.<br />If specified, this value overrides a value specified in the Cluster's "spec.updatePolicy" field.<br />Valid values:<br />  'automatic' (default): apply updates automatically (apply OS security upgrades, avoiding rebooting when possible)<br />  'external': do not apply updates automatically; they are applied manually or by an external system.

## Nested resources

### volume_spec

VolumeSpec defined the spec for an additional volume attached to the instance group.

#### Argument Reference

The following arguments are supported:

- `delete_on_termination` - (Optional) - Bool - DeleteOnTermination configures volume retention policy upon instance termination.<br />The volume is deleted by default. Cluster deletion does not remove retained volumes.<br />NOTE: This setting applies only to the Launch Configuration and does not affect Launch Templates.
- `device` - (Required) - String - Device is an optional device name of the block device.
- `encrypted` - (Optional) - Bool - Encrypted indicates you want to encrypt the volume.
- `iops` - (Optional) - Int - Iops is the provisioned IOPS for the volume when the volume type is io1, io2 or gp3 (AWS only).
- `throughput` - (Optional) - Int - Throughput is the volume throughput in MBps when the volume type is gp3 (AWS only).
- `key` - (Optional) - String - Key is the encryption key identifier for the volume.
- `size` - (Optional) - Int - Size is the size of the volume in GB.
- `type` - (Optional) - String - Type is the type of volume to create and is cloud specific.

### volume_mount_spec

VolumeMountSpec defines the specification for mounting a device.

#### Argument Reference

The following arguments are supported:

- `device` - (Required) - String - Device is the device name to provision and mount.
- `filesystem` - (Required) - String - Filesystem is the filesystem to mount.
- `format_options` - (Optional) - List(String) - FormatOptions is a collection of options passed when formatting the device.
- `mount_options` - (Optional) - List(String) - MountOptions is a collection of mount options - @TODO need to be added.
- `path` - (Required) - String - Path is the location to mount the device.

### hook_spec

HookSpec is a definition hook.

#### Argument Reference

The following arguments are supported:

- `name` - (Required) - String - Name is an optional name for the hook, otherwise the name is kops-hook-<index>.
- `disabled` - (Optional) - Bool - Disabled indicates if you want the unit switched off.
- `roles` - (Optional) - List(String) - Roles is an optional list of roles the hook should be rolled out to, defaults to all.
- `requires` - (Optional) - List(String) - Requires is a series of systemd units the action requires.
- `before` - (Optional) - List(String) - Before is a series of systemd units which this hook must run before.
- `exec_container` - (Optional) - [exec_container_action](#exec_container_action) - ExecContainer is the image itself.
- `manifest` - (Optional) - String - Manifest is a raw systemd unit file.
- `use_raw_manifest` - (Optional) - Bool - UseRawManifest indicates that the contents of Manifest should be used as the contents<br />of the systemd unit, unmodified. Before and Requires are ignored when used together<br />with this value (and validation shouldn't allow them to be set).

### exec_container_action

ExecContainerAction defines an hood action.

#### Argument Reference

The following arguments are supported:

- `image` - (Required) - String - Image is the docker image.
- `command` - (Optional) - List(String) - Command is the command supplied to the above image.
- `environment` - (Optional) - Map(String) - Environment is a map of environment variables added to the hook.

### file_asset_spec

FileAssetSpec defines the structure for a file asset.

#### Argument Reference

The following arguments are supported:

- `name` - (Required) - String - Name is a shortened reference to the asset.
- `path` - (Required) - String - Path is the location this file should reside.
- `roles` - (Optional) - List(String) - Roles is a list of roles the file asset should be applied, defaults to all.
- `content` - (Required) - String - Content is the contents of the file.
- `is_base64` - (Optional) - Bool - IsBase64 indicates the contents is base64 encoded.

### kubelet_config_spec

KubeletConfigSpec defines the kubelet configuration.

#### Argument Reference

The following arguments are supported:

- `api_servers` - (Optional) - String - APIServers is not used for clusters version 1.6 and later - flag removed.
- `anonymous_auth` - (Optional) - Bool - AnonymousAuth permits you to control auth to the kubelet api.
- `authorization_mode` - (Optional) - String - AuthorizationMode is the authorization mode the kubelet is running in.
- `bootstrap_kubeconfig` - (Optional) - String - BootstrapKubeconfig is the path to a kubeconfig file that will be used to get client certificate for kubelet.
- `client_ca_file` - (Optional) - String - ClientCAFile is the path to a CA certificate.
- `tls_cert_file` - (Optional) - String - TODO: Remove unused TLSCertFile.
- `tls_private_key_file` - (Optional) - String - TODO: Remove unused TLSPrivateKeyFile.
- `tls_cipher_suites` - (Optional) - List(String) - TLSCipherSuites indicates the allowed TLS cipher suite.
- `tls_min_version` - (Optional) - String - TLSMinVersion indicates the minimum TLS version allowed.
- `kubeconfig_path` - (Optional) - String - KubeconfigPath is the path of kubeconfig for the kubelet.
- `require_kubeconfig` - (Optional) - Bool - RequireKubeconfig indicates a kubeconfig is required.
- `log_level` - (Optional) - Int - LogLevel is the logging level of the kubelet.
- `pod_manifest_path` - (Optional) - String - config is the path to the config file or directory of files.
- `hostname_override` - (Optional) - String - HostnameOverride is the hostname used to identify the kubelet instead of the actual hostname.
- `pod_infra_container_image` - (Optional) - String - PodInfraContainerImage is the image whose network/ipc containers in each pod will use.
- `seccomp_profile_root` - (Optional) - String - SeccompProfileRoot is the directory path for seccomp profiles.
- `allow_privileged` - (Optional) - Bool - AllowPrivileged enables containers to request privileged mode (defaults to false).
- `enable_debugging_handlers` - (Optional) - Bool - EnableDebuggingHandlers enables server endpoints for log collection and local running of containers and commands.
- `register_node` - (Optional) - Bool - RegisterNode enables automatic registration with the apiserver.
- `node_status_update_frequency` - (Optional) - Duration - NodeStatusUpdateFrequency Specifies how often kubelet posts node status to master (default 10s)<br />must work with nodeMonitorGracePeriod in KubeControllerManagerConfig.
- `cluster_domain` - (Optional) - String - ClusterDomain is the DNS domain for this cluster.
- `cluster_dns` - (Optional) - String - ClusterDNS is the IP address for a cluster DNS server.
- `network_plugin_name` - (Optional) - String - NetworkPluginName is the name of the network plugin to be invoked for various events in kubelet/pod lifecycle.
- `cloud_provider` - (Optional) - String - CloudProvider is the provider for cloud services.
- `kubelet_cgroups` - (Optional) - String - KubeletCgroups is the absolute name of cgroups to isolate the kubelet in.
- `runtime_cgroups` - (Optional) - String - Cgroups that container runtime is expected to be isolated in.
- `read_only_port` - (Optional) - Int - ReadOnlyPort is the port used by the kubelet api for read-only access (default 10255).
- `system_cgroups` - (Optional) - String - SystemCgroups is absolute name of cgroups in which to place<br />all non-kernel processes that are not already in a container. Empty<br />for no container. Rolling back the flag requires a reboot.
- `cgroup_root` - (Optional) - String - cgroupRoot is the root cgroup to use for pods. This is handled by the container runtime on a best effort basis.
- `configure_cbr0` - (Optional) - Bool - configureCBR0 enables the kubelet to configure cbr0 based on Node.Spec.PodCIDR.
- `hairpin_mode` - (Optional) - String - How should the kubelet configure the container bridge for hairpin packets.<br />Setting this flag allows endpoints in a Service to loadbalance back to<br />themselves if they should try to access their own Service. Values:<br />  "promiscuous-bridge": make the container bridge promiscuous.<br />  "hairpin-veth":       set the hairpin flag on container veth interfaces.<br />  "none":               do nothing.<br />Setting --configure-cbr0 to false implies that to achieve hairpin NAT<br />one must set --hairpin-mode=veth-flag, because bridge assumes the<br />existence of a container bridge named cbr0.
- `babysit_daemons` - (Optional) - Bool - The node has babysitter process monitoring docker and kubelet. Removed as of 1.7.
- `max_pods` - (Optional) - Int - MaxPods is the number of pods that can run on this Kubelet.
- `nvidia_gp_us` - (Optional) - Int - NvidiaGPUs is the number of NVIDIA GPU devices on this node.
- `pod_cidr` - (Optional) - String - PodCIDR is the CIDR to use for pod IP addresses, only used in standalone mode.<br />In cluster mode, this is obtained from the master.
- `resolver_config` - (Optional) - String - ResolverConfig is the resolver configuration file used as the basis for the container DNS resolution configuration."), [].
- `reconcile_cidr` - (Optional) - Bool - ReconcileCIDR is Reconcile node CIDR with the CIDR specified by the<br />API server. No-op if register-node or configure-cbr0 is false.
- `register_schedulable` - (Optional) - Bool - registerSchedulable tells the kubelet to register the node as schedulable. No-op if register-node is false.
- `serialize_image_pulls` - (Optional) - Bool - // SerializeImagePulls when enabled, tells the Kubelet to pull images one<br />// at a time. We recommend *not* changing the default value on nodes that<br />// run docker daemon with version  < 1.9 or an Aufs storage backend.<br />// Issue #10959 has more details.
- `node_labels` - (Optional) - Map(String) - NodeLabels to add when registering the node in the cluster.
- `non_masquerade_cidr` - (Optional) - String - NonMasqueradeCIDR configures masquerading: traffic to IPs outside this range will use IP masquerade.
- `enable_custom_metrics` - (Optional) - Bool - Enable gathering custom metrics.
- `network_plugin_mtu` - (Optional) - Int - NetworkPluginMTU is the MTU to be passed to the network plugin,<br />and overrides the default MTU for cases where it cannot be automatically<br />computed (such as IPSEC).
- `image_gc_high_threshold_percent` - (Optional) - Int - ImageGCHighThresholdPercent is the percent of disk usage after which<br />image garbage collection is always run.
- `image_gc_low_threshold_percent` - (Optional) - Int - ImageGCLowThresholdPercent is the percent of disk usage before which<br />image garbage collection is never run. Lowest disk usage to garbage<br />collect to.
- `image_pull_progress_deadline` - (Optional) - Duration - ImagePullProgressDeadline is the timeout for image pulls<br />If no pulling progress is made before this deadline, the image pulling will be cancelled. (default 1m0s).
- `eviction_hard` - (Optional) - String - Comma-delimited list of hard eviction expressions.  For example, 'memory.available<300Mi'.
- `eviction_soft` - (Optional) - String - Comma-delimited list of soft eviction expressions.  For example, 'memory.available<300Mi'.
- `eviction_soft_grace_period` - (Optional) - String - Comma-delimited list of grace periods for each soft eviction signal.  For example, 'memory.available=30s'.
- `eviction_pressure_transition_period` - (Optional) - Duration - Duration for which the kubelet has to wait before transitioning out of an eviction pressure condition.
- `eviction_max_pod_grace_period` - (Optional) - Int - Maximum allowed grace period (in seconds) to use when terminating pods in response to a soft eviction threshold being met.
- `eviction_minimum_reclaim` - (Optional) - String - Comma-delimited list of minimum reclaims (e.g. imagefs.available=2Gi) that describes the minimum amount of resource the kubelet will reclaim when performing a pod eviction if that resource is under pressure.
- `volume_plugin_directory` - (Optional) - String - The full path of the directory in which to search for additional third party volume plugins (this path must be writeable, dependent on your choice of OS).
- `taints` - (Optional) - List(String) - Taints to add when registering a node in the cluster.
- `feature_gates` - (Optional) - Map(String) - FeatureGates is set of key=value pairs that describe feature gates for alpha/experimental features.
- `kube_reserved` - (Optional) - Map(String) - Resource reservation for kubernetes system daemons like the kubelet, container runtime, node problem detector, etc.
- `kube_reserved_cgroup` - (Optional) - String - Control group for kube daemons.
- `system_reserved` - (Optional) - Map(String) - Capture resource reservation for OS system daemons like sshd, udev, etc.
- `system_reserved_cgroup` - (Optional) - String - Parent control group for OS system daemons.
- `enforce_node_allocatable` - (Optional) - String - Enforce Allocatable across pods whenever the overall usage across all pods exceeds Allocatable.
- `runtime_request_timeout` - (Optional) - Duration - RuntimeRequestTimeout is timeout for runtime requests on - pull, logs, exec and attach.
- `volume_stats_agg_period` - (Optional) - Duration - VolumeStatsAggPeriod is the interval for kubelet to calculate and cache the volume disk usage for all pods and volumes.
- `fail_swap_on` - (Optional) - Bool - Tells the Kubelet to fail to start if swap is enabled on the node.
- `experimental_allowed_unsafe_sysctls` - (Optional) - List(String) - ExperimentalAllowedUnsafeSysctls are passed to the kubelet config to whitelist allowable sysctls<br />Was promoted to beta and renamed. https://github.com/kubernetes/kubernetes/pull/63717.
- `allowed_unsafe_sysctls` - (Optional) - List(String) - AllowedUnsafeSysctls are passed to the kubelet config to whitelist allowable sysctls.
- `streaming_connection_idle_timeout` - (Optional) - Duration - StreamingConnectionIdleTimeout is the maximum time a streaming connection can be idle before the connection is automatically closed.
- `docker_disable_shared_pid` - (Optional) - Bool - DockerDisableSharedPID uses a shared PID namespace for containers in a pod.
- `root_dir` - (Optional) - String - RootDir is the directory path for managing kubelet files (volume mounts,etc).
- `authentication_token_webhook` - (Optional) - Bool - AuthenticationTokenWebhook uses the TokenReview API to determine authentication for bearer tokens.
- `authentication_token_webhook_cache_ttl` - (Optional) - Duration - AuthenticationTokenWebhook sets the duration to cache responses from the webhook token authenticator. Default is 2m. (default 2m0s).
- `cpucfs_quota` - (Optional) - Bool - CPUCFSQuota enables CPU CFS quota enforcement for containers that specify CPU limits.
- `cpucfs_quota_period` - (Optional) - Duration - CPUCFSQuotaPeriod sets CPU CFS quota period value, cpu.cfs_period_us, defaults to Linux Kernel default.
- `cpu_manager_policy` - (Optional) - String - CpuManagerPolicy allows for changing the default policy of None to static.
- `registry_pull_qps` - (Optional) - Int - RegistryPullQPS if > 0, limit registry pull QPS to this value.  If 0, unlimited. (default 5).
- `registry_burst` - (Optional) - Int - RegistryBurst Maximum size of a bursty pulls, temporarily allows pulls to burst to this number, while still not exceeding registry-qps. Only used if --registry-qps > 0 (default 10).
- `topology_manager_policy` - (Optional) - String - TopologyManagerPolicy determines the allocation policy for the topology manager.
- `rotate_certificates` - (Optional) - Bool - rotateCertificates enables client certificate rotation.
- `protect_kernel_defaults` - (Optional) - Bool - Default kubelet behaviour for kernel tuning. If set, kubelet errors if any of kernel tunables is different than kubelet defaults.<br />(DEPRECATED: This parameter should be set via the config file specified by the Kubelet's --config flag.
- `cgroup_driver` - (Optional) - String - CgroupDriver allows the explicit setting of the kubelet cgroup driver. If omitted, defaults to cgroupfs.
- `housekeeping_interval` - (Optional) - Duration - HousekeepingInterval allows to specify interval between container housekeepings.
- `event_qps` - (Optional) - Int - EventQPS if > 0, limit event creations per second to this value.  If 0, unlimited.
- `event_burst` - (Optional) - Int - EventBurst temporarily allows event records to burst to this number, while still not exceeding EventQPS. Only used if EventQPS > 0.
- `container_log_max_size` - (Optional) - String - ContainerLogMaxSize is the maximum size (e.g. 10Mi) of container log file before it is rotated.
- `container_log_max_files` - (Optional) - Int - ContainerLogMaxFiles is the maximum number of container log files that can be present for a container. The number must be >= 2.
- `enable_cadvisor_json_endpoints` - (Optional) - Bool - EnableCadvisorJsonEndpoints enables cAdvisor json `/spec` and `/stats/*` endpoints. Defaults to False.

### mixed_instances_policy_spec

MixedInstancesPolicySpec defines the specification for an autoscaling group backed by a ec2 fleet.

#### Argument Reference

The following arguments are supported:

- `instances` - (Optional) - List(String) - Instances is a list of instance types which we are willing to run in the EC2 fleet.
- `on_demand_allocation_strategy` - (Optional) - String - OnDemandAllocationStrategy indicates how to allocate instance types to fulfill On-Demand capacity.
- `on_demand_base` - (Optional) - Int - OnDemandBase is the minimum amount of the Auto Scaling group's capacity that must be<br />fulfilled by On-Demand Instances. This base portion is provisioned first as your group scales.
- `on_demand_above_base` - (Optional) - Int - OnDemandAboveBase controls the percentages of On-Demand Instances and Spot Instances for your<br />additional capacity beyond OnDemandBase. The range is 0–100. The default value is 100. If you<br />leave this parameter set to 100, the percentages are 100% for On-Demand Instances and 0% for<br />Spot Instances.
- `spot_allocation_strategy` - (Optional) - String - SpotAllocationStrategy diversifies your Spot capacity across multiple instance types to<br />find the best pricing. Higher Spot availability may result from a larger number of<br />instance types to choose from.
- `spot_instance_pools` - (Optional) - Int - SpotInstancePools is the number of Spot pools to use to allocate your Spot capacity (defaults to 2)<br />pools are determined from the different instance types in the Overrides array of LaunchTemplate.

### user_data

UserData defines a user-data section.

#### Argument Reference

The following arguments are supported:

- `name` - (Required) - String - Name is the name of the user-data.
- `type` - (Required) - String - Type is the type of user-data.
- `content` - (Required) - String - Content is the user-data content.

### load_balancer

LoadBalancer defines a load balancer.

#### Argument Reference

The following arguments are supported:

- `load_balancer_name` - (Optional) - String - LoadBalancerName to associate with this instance group (AWS ELB).
- `target_group_arn` - (Optional) - String - TargetGroupARN to associate with this instance group (AWS ALB/NLB).

### iam_profile_spec

IAMProfileSpec is the AWS IAM Profile to attach to instances in this instance<br />group. Specify the ARN for the IAM instance profile (AWS only).

#### Argument Reference

The following arguments are supported:

- `profile` - (Required) - String - Profile is the AWS IAM Profile to attach to instances in this instance group.<br />Specify the ARN for the IAM instance profile. (AWS only).

### rolling_update

#### Argument Reference

The following arguments are supported:

- `drain_and_terminate` - (Optional) - Bool - DrainAndTerminate enables draining and terminating nodes during rolling updates.<br />Defaults to true.
- `max_unavailable` - (Optional) - IntOrString - MaxUnavailable is the maximum number of nodes that can be unavailable during the update.<br />The value can be an absolute number (for example 5) or a percentage of desired<br />nodes (for example 10%).<br />The absolute number is calculated from a percentage by rounding down.<br />Defaults to 1 if MaxSurge is 0, otherwise defaults to 0.<br />Example: when this is set to 30%, the InstanceGroup can be scaled<br />down to 70% of desired nodes immediately when the rolling update<br />starts. Once new nodes are ready, more old nodes can be drained,<br />ensuring that the total number of nodes available at all times<br />during the update is at least 70% of desired nodes.<br />+optional.
- `max_surge` - (Optional) - IntOrString - MaxSurge is the maximum number of extra nodes that can be created<br />during the update.<br />The value can be an absolute number (for example 5) or a percentage of<br />desired machines (for example 10%).<br />The absolute number is calculated from a percentage by rounding up.<br />Has no effect on instance groups with role "Master".<br />Defaults to 1 on AWS, 0 otherwise.<br />Example: when this is set to 30%, the InstanceGroup can be scaled<br />up immediately when the rolling update starts, such that the total<br />number of old and new nodes do not exceed 130% of desired<br />nodes.<br />+optional.

### instance_metadata_options

InstanceMetadataOptions defines the EC2 instance metadata service options (AWS Only).

#### Argument Reference

The following arguments are supported:

- `http_put_response_hop_limit` - (Optional) - Int - HTTPPutResponseHopLimit is the desired HTTP PUT response hop limit for instance metadata requests.<br />The larger the number, the further instance metadata requests can travel. The default value is 1.
- `http_tokens` - (Optional) - String - HTTPTokens is the state of token usage for the instance metadata requests.<br />If the parameter is not specified in the request, the default state is "required".


## Import

You can import an existing cluster by creating a `kops_instance_group` configuration
and running the `terraform import` command:

1. Create a terraform configuration:

    ```hcl
    provider "kops" {
      state_store = "s3://cluster.example.com"
    }

    resource "kops_instance_group" "ig-0" {
      cluster_namename  = "ig-0"
      name              = "cluster.example.com"
      
      // ....
    }
    ```

1. Run `terraform import`:

    ```shell
    terraform import kops_instance_group.ig-0 cluster.example.com/ig-0
    ```

~> The id of the instance group to be imported must be given in the 
`cluster name/instance group name` format.

