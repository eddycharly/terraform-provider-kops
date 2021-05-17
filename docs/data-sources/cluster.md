# kops_cluster

Provides a kOps cluster data source.

## Example usage

```hcl
data "kops_cluster" "cluster" {
  cluster_name = "cluster.example.com"
}
```

## Argument Reference

The following arguments are supported:
- `name` - (Required) - String - Name defines the cluster name.
- `admin_ssh_key` - (Computed) - String - AdminSshKey defines the cluster admin ssh key.
- `channel` - (Computed) - String - The Channel we are following.
- `addons` - (Computed) - List([addon_spec](#addon_spec)) - Additional addons that should be installed on the cluster.
- `config_base` - (Computed) - String - ConfigBase is the path where we store configuration for the cluster<br />This might be different than the location where the cluster spec itself is stored,<br />both because this must be accessible to the cluster,<br />and because it might be on a different cloud or storage system (etcd vs S3).
- `cloud_provider` - (Computed) - String - The CloudProvider to use (aws or gce).
- `container_runtime` - (Computed) - String - Container runtime to use for Kubernetes.
- `kubernetes_version` - (Computed) - String - The version of kubernetes to install (optional, and can be a "spec" like stable).
- `subnet` - (Computed) - List([cluster_subnet_spec](#cluster_subnet_spec)) - Configuration of subnets we are targeting.
- `project` - (Computed) - String - Project is the cloud project we should use, required on GCE.
- `master_public_name` - (Computed) - String - MasterPublicName is the external DNS name for the master nodes.
- `master_internal_name` - (Computed) - String - MasterInternalName is the internal DNS name for the master nodes.
- `network_cidr` - (Computed) - String - NetworkCIDR is the CIDR used for the AWS VPC / GCE Network, or otherwise allocated to k8s<br />This is a real CIDR, not the internal k8s network<br />On AWS, it maps to the VPC CIDR.  It is not required on GCE.
- `additional_network_cidrs` - (Computed) - List(String) - AdditionalNetworkCIDRs is a list of additional CIDR used for the AWS VPC<br />or otherwise allocated to k8s. This is a real CIDR, not the internal k8s network<br />On AWS, it maps to any additional CIDRs added to a VPC.
- `network_id` - (Computed) - String - NetworkID is an identifier of a network, if we want to reuse/share an existing network (e.g. an AWS VPC).
- `topology` - (Computed) - [topology_spec](#topology_spec) - Topology defines the type of network topology to use on the cluster - default public<br />This is heavily weighted towards AWS for the time being, but should also be agnostic enough<br />to port out to GCE later if needed.
- `secret_store` - (Computed) - String - SecretStore is the VFS path to where secrets are stored.
- `key_store` - (Computed) - String - KeyStore is the VFS path to where SSL keys and certificates are stored.
- `config_store` - (Computed) - String - ConfigStore is the VFS path to where the configuration (Cluster, InstanceGroups etc) is stored.
- `dns_zone` - (Computed) - String - DNSZone is the DNS zone we should use when configuring DNS<br />This is because some clouds let us define a managed zone foo.bar, and then have<br />kubernetes.dev.foo.bar, without needing to define dev.foo.bar as a hosted zone.<br />DNSZone will probably be a suffix of the MasterPublicName and MasterInternalName<br />Note that DNSZone can either by the host name of the zone (containing dots),<br />or can be an identifier for the zone.
- `additional_sans` - (Computed) - List(String) - AdditionalSANs adds additional Subject Alternate Names to apiserver cert that kops generates.
- `cluster_dns_domain` - (Computed) - String - ClusterDNSDomain is the suffix we use for internal DNS names (normally cluster.local).
- `service_cluster_ip_range` - (Computed) - String - ServiceClusterIPRange is the CIDR, from the internal network, where we allocate IPs for services.
- `pod_cidr` - (Computed) - String - PodCIDR is the CIDR from which we allocate IPs for pods.
- `non_masquerade_cidr` - (Computed) - String - NonMasqueradeCIDR is the CIDR for the internal k8s network (on which pods & services live)<br />It cannot overlap ServiceClusterIPRange.
- `ssh_access` - (Computed) - List(String) - SSHAccess is a list of the CIDRs that can access SSH.
- `node_port_access` - (Computed) - List(String) - NodePortAccess is a list of the CIDRs that can access the node ports range (30000-32767).
- `egress_proxy` - (Computed) - [egress_proxy_spec](#egress_proxy_spec) - HTTPProxy defines connection information to support use of a private cluster behind an forward HTTP Proxy.
- `ssh_key_name` - (Computed) - String - SSHKeyName specifies a preexisting SSH key to use.
- `kubernetes_api_access` - (Computed) - List(String) - KubernetesAPIAccess is a list of the CIDRs that can access the Kubernetes API endpoint (master HTTPS).
- `isolate_masters` - (Computed) - Bool - IsolateMasters determines whether we should lock down masters so that they are not on the pod network.<br />true is the kube-up behaviour, but it is very surprising: it means that daemonsets only work on the master<br />if they have hostNetwork=true.<br />false is now the default, and it will:<br /> * give the master a normal PodCIDR<br /> * run kube-proxy on the master<br /> * enable debugging handlers on the master, so kubectl logs works.
- `update_policy` - (Computed) - String - UpdatePolicy determines the policy for applying upgrades automatically.<br />Valid values:<br />  'automatic' (default): apply updates automatically (apply OS security upgrades, avoiding rebooting when possible)<br />  'external': do not apply updates automatically; they are applied manually or by an external system.
- `external_policies` - (Computed) - Map(List(String)) - ExternalPolicies allows the insertion of pre-existing managed policies on IG Roles.
- `additional_policies` - (Computed) - Map(String) - Additional policies to add for roles.
- `file_assets` - (Computed) - List([file_asset_spec](#file_asset_spec)) - A collection of files assets for deployed cluster wide.
- `etcd_cluster` - (Computed) - List([etcd_cluster_spec](#etcd_cluster_spec)) - EtcdClusters stores the configuration for each cluster.
- `containerd` - (Computed) - [containerd_config](#containerd_config) - Component configurations.
- `docker` - (Computed) - [docker_config](#docker_config)
- `kube_dns` - (Computed) - [kube_dns_config](#kube_dns_config)
- `kube_api_server` - (Computed) - [kube_api_server_config](#kube_api_server_config)
- `kube_controller_manager` - (Computed) - [kube_controller_manager_config](#kube_controller_manager_config)
- `external_cloud_controller_manager` - (Computed) - [cloud_controller_manager_config](#cloud_controller_manager_config)
- `kube_scheduler` - (Computed) - [kube_scheduler_config](#kube_scheduler_config)
- `kube_proxy` - (Computed) - [kube_proxy_config](#kube_proxy_config)
- `kubelet` - (Computed) - [kubelet_config_spec](#kubelet_config_spec)
- `master_kubelet` - (Computed) - [kubelet_config_spec](#kubelet_config_spec)
- `cloud_config` - (Computed) - [cloud_configuration](#cloud_configuration)
- `external_dns` - (Computed) - [external_dns_config](#external_dns_config)
- `ntp` - (Computed) - [ntp_config](#ntp_config)
- `node_termination_handler` - (Computed) - [node_termination_handler_config](#node_termination_handler_config) - NodeTerminationHandler determines the cluster autoscaler configuration.
- `metrics_server` - (Computed) - [metrics_server_config](#metrics_server_config) - MetricsServer determines the metrics server configuration.
- `cert_manager` - (Computed) - [cert_manager_config](#cert_manager_config) - CertManager determines the metrics server configuration.
- `aws_load_balancer_controller` - (Computed) - [aws_load_balancer_controller_config](#aws_load_balancer_controller_config) - AWSLoadbalancerControllerConfig determines the AWS LB controller configuration.
- `networking` - (Computed) - [networking_spec](#networking_spec) - Networking configuration.
- `api` - (Computed) - [access_spec](#access_spec) - API field controls how the API is exposed outside the cluster.
- `authentication` - (Computed) - [authentication_spec](#authentication_spec) - Authentication field controls how the cluster is configured for authentication.
- `authorization` - (Computed) - [authorization_spec](#authorization_spec) - Authorization field controls how the cluster is configured for authorization.
- `node_authorization` - (Computed) - [node_authorization_spec](#node_authorization_spec) - NodeAuthorization defined the custom node authorization configuration.
- `cloud_labels` - (Computed) - Map(String) - Tags for AWS instance groups.
- `hooks` - (Computed) - List([hook_spec](#hook_spec)) - Hooks for custom actions e.g. on first installation.
- `assets` - (Computed) - [assets](#assets) - Assets is alternative locations for files and containers; the API under construction, will remove this comment once this API is fully functional.
- `iam` - (Computed) - [iam_spec](#iam_spec) - IAM field adds control over the IAM security policies applied to resources.
- `encryption_config` - (Computed) - Bool - EncryptionConfig controls if encryption is enabled.
- `disable_subnet_tags` - (Computed) - Bool - DisableSubnetTags controls if subnets are tagged in AWS.
- `use_host_certificates` - (Computed) - Bool - UseHostCertificates will mount /etc/ssl/certs to inside needed containers.<br />This is needed if some APIs do have self-signed certs.
- `sysctl_parameters` - (Computed) - List(String) - SysctlParameters will configure kernel parameters using sysctl(8). When<br />specified, each parameter must follow the form variable=value, the way<br />it would appear in sysctl.conf.
- `rolling_update` - (Computed) - [rolling_update](#rolling_update) - RollingUpdate defines the default rolling-update settings for instance groups.
- `cluster_autoscaler` - (Computed) - [cluster_autoscaler_config](#cluster_autoscaler_config) - ClusterAutoscaler defines the cluster autoscaler configuration.

## Nested resources

### addon_spec

AddonSpec defines an addon that we want to install in the cluster.

#### Argument Reference

The following arguments are supported:

- `manifest` - (Computed) - String - Manifest is a path to the manifest that defines the addon.

### cluster_subnet_spec

ClusterSubnetSpec defines a subnet.

#### Argument Reference

The following arguments are supported:

- `name` - (Computed) - String - Name is the name of the subnet.
- `cidr` - (Computed) - String - CIDR is the network cidr of the subnet.
- `zone` - (Computed) - String - Zone is the zone the subnet is in, set for subnets that are zonally scoped.
- `region` - (Computed) - String - Region is the region the subnet is in, set for subnets that are regionally scoped.
- `provider_id` - (Computed) - String - ProviderID is the cloud provider id for the objects associated with the zone (the subnet on AWS).
- `egress` - (Computed) - String - Egress defines the method of traffic egress for this subnet.
- `type` - (Computed) - String - Type define which one if the internal types (public, utility, private) the network is.
- `public_ip` - (Computed) - String - PublicIP to attach to NatGateway.

### topology_spec

#### Argument Reference

The following arguments are supported:

- `masters` - (Computed) - String - The environment to launch the Kubernetes masters in public|private.
- `nodes` - (Computed) - String - The environment to launch the Kubernetes nodes in public|private.
- `bastion` - (Computed) - [bastion_spec](#bastion_spec) - Bastion provide an external facing point of entry into a network<br />containing private network instances. This host can provide a single<br />point of fortification or audit and can be started and stopped to enable<br />or disable inbound SSH communication from the Internet, some call bastion<br />as the "jump server".
- `dns` - (Computed) - [dns_spec](#dns_spec) - DNS configures options relating to DNS, in particular whether we use a public or a private hosted zone.

### bastion_spec

#### Argument Reference

The following arguments are supported:

- `bastion_public_name` - (Computed) - String
- `idle_timeout_seconds` - (Computed) - Int - IdleTimeoutSeconds is the bastion's Loadbalancer idle timeout.
- `load_balancer` - (Computed) - [bastion_load_balancer_spec](#bastion_load_balancer_spec)

### bastion_load_balancer_spec

#### Argument Reference

The following arguments are supported:

- `additional_security_groups` - (Computed) - List(String)

### dns_spec

#### Argument Reference

The following arguments are supported:

- `type` - (Computed) - String

### egress_proxy_spec

#### Argument Reference

The following arguments are supported:

- `http_proxy` - (Computed) - [http_proxy](#http_proxy)
- `proxy_excludes` - (Computed) - String

### http_proxy

#### Argument Reference

The following arguments are supported:

- `host` - (Computed) - String
- `port` - (Computed) - Int

### file_asset_spec

FileAssetSpec defines the structure for a file asset.

#### Argument Reference

The following arguments are supported:

- `name` - (Computed) - String - Name is a shortened reference to the asset.
- `path` - (Computed) - String - Path is the location this file should reside.
- `roles` - (Computed) - List(String) - Roles is a list of roles the file asset should be applied, defaults to all.
- `content` - (Computed) - String - Content is the contents of the file.
- `is_base64` - (Computed) - Bool - IsBase64 indicates the contents is base64 encoded.

### etcd_cluster_spec

EtcdClusterSpec is the etcd cluster specification.

#### Argument Reference

The following arguments are supported:

- `name` - (Computed) - String - Name is the name of the etcd cluster (main, events etc).
- `provider` - (Computed) - String - Provider is the provider used to run etcd: Manager, Legacy.<br />Defaults to Manager.
- `member` - (Computed) - List([etcd_member_spec](#etcd_member_spec)) - Members stores the configurations for each member of the cluster (including the data volume).
- `enable_etcd_tls` - (Computed) - Bool - EnableEtcdTLS indicates the etcd service should use TLS between peers and clients.
- `enable_tls_auth` - (Computed) - Bool - EnableTLSAuth indicates client and peer TLS auth should be enforced.
- `version` - (Computed) - String - Version is the version of etcd to run.
- `leader_election_timeout` - (Computed) - Duration - LeaderElectionTimeout is the time (in milliseconds) for an etcd leader election timeout.
- `heartbeat_interval` - (Computed) - Duration - HeartbeatInterval is the time (in milliseconds) for an etcd heartbeat interval.
- `image` - (Computed) - String - Image is the etcd docker image to use. Setting this will ignore the Version specified.
- `backups` - (Computed) - [etcd_backup_spec](#etcd_backup_spec) - Backups describes how we do backups of etcd.
- `manager` - (Computed) - [etcd_manager_spec](#etcd_manager_spec) - Manager describes the manager configuration.
- `memory_request` - (Computed) - Quantity - MemoryRequest specifies the memory requests of each etcd container in the cluster.
- `cpu_request` - (Computed) - Quantity - CPURequest specifies the cpu requests of each etcd container in the cluster.

### etcd_member_spec

EtcdMemberSpec is a specification for a etcd member.

#### Argument Reference

The following arguments are supported:

- `name` - (Computed) - String - Name is the name of the member within the etcd cluster.
- `instance_group` - (Computed) - String - InstanceGroup is the instanceGroup this volume is associated.
- `volume_type` - (Computed) - String - VolumeType is the underlying cloud storage class.
- `volume_iops` - (Computed) - Int - If volume type is io1, then we need to specify the number of Iops.
- `volume_throughput` - (Computed) - Int - Parameter for disks that support provisioned throughput.
- `volume_size` - (Computed) - Int - VolumeSize is the underlying cloud volume size.
- `kms_key_id` - (Computed) - String - KmsKeyId is a AWS KMS ID used to encrypt the volume.
- `encrypted_volume` - (Computed) - Bool - EncryptedVolume indicates you want to encrypt the volume.

### etcd_backup_spec

EtcdBackupSpec describes how we want to do backups of etcd.

#### Argument Reference

The following arguments are supported:

- `backup_store` - (Computed) - String - BackupStore is the VFS path where we will read/write backup data.
- `image` - (Computed) - String - Image is the etcd backup manager image to use.  Setting this will create a sidecar container in the etcd pod with the specified image.

### etcd_manager_spec

EtcdManagerSpec describes how we configure the etcd manager.

#### Argument Reference

The following arguments are supported:

- `image` - (Computed) - String - Image is the etcd manager image to use.
- `env` - (Computed) - List([env_var](#env_var)) - Env allows users to pass in env variables to the etcd-manager container.<br />Variables starting with ETCD_ will be further passed down to the etcd process.<br />This allows etcd setting to be overwriten. No config validation is done.<br />A list of etcd config ENV vars can be found at https://github.com/etcd-io/etcd/blob/master/Documentation/op-guide/configuration.md.
- `discovery_poll_interval` - (Computed) - String - DiscoveryPollInterval which is used for discovering other cluster members. The default is 60 seconds.
- `log_level` - (Computed) - Int - LogLevel allows the klog library verbose log level to be set for etcd-manager. The default is 6.<br />https://github.com/google/glog#verbose-logging.

### env_var

EnvVar represents an environment variable present in a Container.

#### Argument Reference

The following arguments are supported:

- `name` - (Computed) - String - Name of the environment variable. Must be a C_IDENTIFIER.
- `value` - (Computed) - String - Variable references $(VAR_NAME) are expanded<br />using the previous defined environment variables in the container and<br />any service environment variables. If a variable cannot be resolved,<br />the reference in the input string will be unchanged. The $(VAR_NAME)<br />syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped<br />references will never be expanded, regardless of whether the variable<br />exists or not.<br />Defaults to "".<br />+optional.

### containerd_config

ContainerdConfig is the configuration for containerd.

#### Argument Reference

The following arguments are supported:

- `address` - (Computed) - String - Address of containerd's GRPC server (default "/run/containerd/containerd.sock").
- `config_override` - (Computed) - String - ConfigOverride is the complete containerd config file provided by the user.
- `log_level` - (Computed) - String - LogLevel controls the logging details [trace, debug, info, warn, error, fatal, panic] (default "info").
- `packages` - (Computed) - [packages_config](#packages_config) - Packages overrides the URL and hash for the packages.
- `registry_mirrors` - (Computed) - Map(List(String)) - RegistryMirrors is list of image registries.
- `root` - (Computed) - String - Root directory for persistent data (default "/var/lib/containerd").
- `skip_install` - (Computed) - Bool - SkipInstall prevents kOps from installing and modifying containerd in any way (default "false").
- `state` - (Computed) - String - State directory for execution state files (default "/run/containerd").
- `version` - (Computed) - String - Version used to pick the containerd package.

### packages_config

#### Argument Reference

The following arguments are supported:

- `hash_amd64` - (Computed) - String - HashAmd64 overrides the hash for the AMD64 package.
- `hash_arm64` - (Computed) - String - HashArm64 overrides the hash for the ARM64 package.
- `url_amd64` - (Computed) - String - UrlAmd64 overrides the URL for the AMD64 package.
- `url_arm64` - (Computed) - String - UrlArm64 overrides the URL for the ARM64 package.

### docker_config

DockerConfig is the configuration for docker.

#### Argument Reference

The following arguments are supported:

- `authorization_plugins` - (Computed) - List(String) - AuthorizationPlugins is a list of authorization plugins.
- `bridge` - (Computed) - String - Bridge is the network interface containers should bind onto.
- `bridge_ip` - (Computed) - String - BridgeIP is a specific IP address and netmask for the docker0 bridge, using standard CIDR notation.
- `data_root` - (Computed) - String - DataRoot is the root directory of persistent docker state (default "/var/lib/docker").
- `default_ulimit` - (Computed) - List(String) - DefaultUlimit is the ulimits for containers.
- `default_runtime` - (Computed) - String - DefaultRuntime is the default OCI runtime for containers (default "runc").
- `exec_opt` - (Computed) - List(String) - ExecOpt is a series of options passed to the runtime.
- `exec_root` - (Computed) - String - ExecRoot is the root directory for execution state files (default "/var/run/docker").
- `experimental` - (Computed) - Bool - Experimental features permits enabling new features such as dockerd metrics.
- `health_check` - (Computed) - Bool - HealthCheck enables the periodic health-check service.
- `hosts` - (Computed) - List(String) - Hosts enables you to configure the endpoints the docker daemon listens on i.e. tcp://0.0.0.0.2375 or unix:///var/run/docker.sock etc.
- `ip_masq` - (Computed) - Bool - IPMasq enables ip masquerading for containers.
- `ip_tables` - (Computed) - Bool - IPtables enables addition of iptables rules.
- `insecure_registry` - (Computed) - String - InsecureRegistry enable insecure registry communication @question according to dockers this a list??.
- `insecure_registries` - (Computed) - List(String) - InsecureRegistries enables multiple insecure docker registry communications.
- `live_restore` - (Computed) - Bool - LiveRestore enables live restore of docker when containers are still running.
- `log_driver` - (Computed) - String - LogDriver is the default driver for container logs (default "json-file").
- `log_level` - (Computed) - String - LogLevel is the logging level ("debug", "info", "warn", "error", "fatal") (default "info").
- `log_opt` - (Computed) - List(String) - Logopt is a series of options given to the log driver options for containers.
- `metrics_address` - (Computed) - String - Metrics address is the endpoint to serve with Prometheus format metrics.
- `mtu` - (Computed) - Int - MTU is the containers network MTU.
- `packages` - (Computed) - [packages_config](#packages_config) - Packages overrides the URL and hash for the packages.
- `registry_mirrors` - (Computed) - List(String) - RegistryMirrors is a referred list of docker registry mirror.
- `runtimes` - (Computed) - List(String) - Runtimes registers an additional OCI compatible runtime (default []).
- `selinux_enabled` - (Computed) - Bool - SelinuxEnabled enables SELinux support.
- `skip_install` - (Computed) - Bool - SkipInstall when set to true will prevent kops from installing and modifying Docker in any way.
- `storage` - (Computed) - String - Storage is the docker storage driver to use.
- `storage_opts` - (Computed) - List(String) - StorageOpts is a series of options passed to the storage driver.
- `user_namespace_remap` - (Computed) - String - UserNamespaceRemap sets the user namespace remapping option for the docker daemon.
- `version` - (Computed) - String - Version is consumed by the nodeup and used to pick the docker version.

### kube_dns_config

KubeDNSConfig defines the kube dns configuration.

#### Argument Reference

The following arguments are supported:

- `cache_max_size` - (Computed) - Int - CacheMaxSize is the maximum entries to keep in dnsmasq.
- `cache_max_concurrent` - (Computed) - Int - CacheMaxConcurrent is the maximum number of concurrent queries for dnsmasq.
- `core_dns_image` - (Computed) - String - CoreDNSImage is used to override the default image used for CoreDNS.
- `cpa_image` - (Computed) - String - CPAImage is used to override the default image used for Cluster Proportional Autoscaler.
- `domain` - (Computed) - String - Domain is the dns domain.
- `external_core_file` - (Computed) - String - ExternalCoreFile is used to provide a complete CoreDNS CoreFile by the user - ignores other provided flags which modify the CoreFile.
- `image` - (Computed) - String - Image is the name of the docker image to run - @deprecated as this is now in the addon.
- `replicas` - (Computed) - Int - Replicas is the number of pod replicas - @deprecated as this is now in the addon and controlled by autoscaler.
- `provider` - (Computed) - String - Provider indicates whether CoreDNS or kube-dns will be the default service discovery.
- `server_ip` - (Computed) - String - ServerIP is the server ip.
- `stub_domains` - (Computed) - Map(List(String)) - StubDomains redirects a domains to another DNS service.
- `upstream_nameservers` - (Computed) - List(String) - UpstreamNameservers sets the upstream nameservers for queries not on the cluster domain.
- `memory_request` - (Computed) - Quantity - MemoryRequest specifies the memory requests of each dns container in the cluster. Default 70m.
- `cpu_request` - (Computed) - Quantity - CPURequest specifies the cpu requests of each dns container in the cluster. Default 100m.
- `memory_limit` - (Computed) - Quantity - MemoryLimit specifies the memory limit of each dns container in the cluster. Default 170m.
- `node_local_dns` - (Computed) - [node_local_dns_config](#node_local_dns_config) - NodeLocalDNS specifies the configuration for the node-local-dns addon.

### node_local_dns_config

NodeLocalDNSConfig are options of the node-local-dns.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Computed) - Bool - Enabled activates the node-local-dns addon.
- `local_ip` - (Computed) - String - Local listen IP address. It can be any IP in the 169.254.20.0/16 space or any other IP address that can be guaranteed to not collide with any existing IP.
- `forward_to_kube_dns` - (Computed) - Bool - If enabled, nodelocal dns will use kubedns as a default upstream.
- `memory_request` - (Computed) - Quantity - MemoryRequest specifies the memory requests of each node-local-dns container in the daemonset. Default 5Mi.
- `cpu_request` - (Computed) - Quantity - CPURequest specifies the cpu requests of each node-local-dns container in the daemonset. Default 25m.

### kube_api_server_config

KubeAPIServerConfig defines the configuration for the kube api.

#### Argument Reference

The following arguments are supported:

- `image` - (Computed) - String - Image is the docker container used.
- `disable_basic_auth` - (Computed) - Bool - DisableBasicAuth removes the --basic-auth-file flag.
- `log_level` - (Computed) - Int - LogLevel is the logging level of the api.
- `cloud_provider` - (Computed) - String - CloudProvider is the name of the cloudProvider we are using, aws, gce etcd.
- `secure_port` - (Computed) - Int - SecurePort is the port the kube runs on.
- `insecure_port` - (Computed) - Int - InsecurePort is the port the insecure api runs.
- `address` - (Computed) - String - Address is the binding address for the kube api: Deprecated - use insecure-bind-address and bind-address.
- `bind_address` - (Computed) - String - BindAddress is the binding address for the secure kubernetes API.
- `insecure_bind_address` - (Computed) - String - InsecureBindAddress is the binding address for the InsecurePort for the insecure kubernetes API.
- `enable_bootstrap_auth_token` - (Computed) - Bool - EnableBootstrapAuthToken enables 'bootstrap.kubernetes.io/token' in the 'kube-system' namespace to be used for TLS bootstrapping authentication.
- `enable_aggregator_routing` - (Computed) - Bool - EnableAggregatorRouting enables aggregator routing requests to endpoints IP rather than cluster IP.
- `admission_control` - (Computed) - List(String) - AdmissionControl is a list of admission controllers to use: Deprecated - use enable-admission-plugins instead.
- `append_admission_plugins` - (Computed) - List(String) - AppendAdmissionPlugins appends list of enabled admission plugins.
- `enable_admission_plugins` - (Computed) - List(String) - EnableAdmissionPlugins is a list of enabled admission plugins.
- `disable_admission_plugins` - (Computed) - List(String) - DisableAdmissionPlugins is a list of disabled admission plugins.
- `admission_control_config_file` - (Computed) - String - AdmissionControlConfigFile is the location of the admission-control-config-file.
- `service_cluster_ip_range` - (Computed) - String - ServiceClusterIPRange is the service address range.
- `service_node_port_range` - (Computed) - String - Passed as --service-node-port-range to kube-apiserver. Expects 'startPort-endPort' format e.g. 30000-33000.
- `etcd_servers` - (Computed) - List(String) - EtcdServers is a list of the etcd service to connect.
- `etcd_servers_overrides` - (Computed) - List(String) - EtcdServersOverrides is per-resource etcd servers overrides, comma separated. The individual override format: group/resource#servers, where servers are http://ip:port, semicolon separated.
- `etcd_ca_file` - (Computed) - String - EtcdCAFile is the path to a ca certificate.
- `etcd_cert_file` - (Computed) - String - EtcdCertFile is the path to a certificate.
- `etcd_key_file` - (Computed) - String - EtcdKeyFile is the path to a private key.
- `basic_auth_file` - (Computed) - String - TODO: Remove unused BasicAuthFile.
- `client_ca_file` - (Computed) - String - TODO: Remove unused ClientCAFile.
- `tls_cert_file` - (Computed) - String - TODO: Remove unused TLSCertFile.
- `tls_private_key_file` - (Computed) - String - TODO: Remove unused TLSPrivateKeyFile.
- `tls_cipher_suites` - (Computed) - List(String) - TLSCipherSuites indicates the allowed TLS cipher suite.
- `tls_min_version` - (Computed) - String - TLSMinVersion indicates the minimum TLS version allowed.
- `token_auth_file` - (Computed) - String - TODO: Remove unused TokenAuthFile.
- `allow_privileged` - (Computed) - Bool - AllowPrivileged indicates if we can run privileged containers.
- `api_server_count` - (Computed) - Int - APIServerCount is the number of api servers.
- `runtime_config` - (Computed) - Map(String) - RuntimeConfig is a series of keys/values are parsed into the `--runtime-config` parameters.
- `kubelet_client_certificate` - (Computed) - String - KubeletClientCertificate is the path of a certificate for secure communication between api and kubelet.
- `kubelet_certificate_authority` - (Computed) - String - KubeletCertificateAuthority is the path of a certificate authority for secure communication between api and kubelet.
- `kubelet_client_key` - (Computed) - String - KubeletClientKey is the path of a private to secure communication between api and kubelet.
- `anonymous_auth` - (Computed) - Bool - AnonymousAuth indicates if anonymous authentication is permitted.
- `kubelet_preferred_address_types` - (Computed) - List(String) - KubeletPreferredAddressTypes is a list of the preferred NodeAddressTypes to use for kubelet connections.
- `storage_backend` - (Computed) - String - StorageBackend is the backend storage.
- `oidc_username_claim` - (Computed) - String - OIDCUsernameClaim is the OpenID claim to use as the user name.<br />Note that claims other than the default ('sub') is not guaranteed to be<br />unique and immutable.
- `oidc_username_prefix` - (Computed) - String - OIDCUsernamePrefix is the prefix prepended to username claims to prevent<br />clashes with existing names (such as 'system:' users).
- `oidc_groups_claim` - (Computed) - String - OIDCGroupsClaim if provided, the name of a custom OpenID Connect claim for<br />specifying user groups.<br />The claim value is expected to be a string or array of strings.
- `oidc_groups_prefix` - (Computed) - String - OIDCGroupsPrefix is the prefix prepended to group claims to prevent<br />clashes with existing names (such as 'system:' groups).
- `oidc_issuer_url` - (Computed) - String - OIDCIssuerURL is the URL of the OpenID issuer, only HTTPS scheme will<br />be accepted.<br />If set, it will be used to verify the OIDC JSON Web Token (JWT).
- `oidc_client_id` - (Computed) - String - OIDCClientID is the client ID for the OpenID Connect client, must be set<br />if oidc-issuer-url is set.
- `oidc_required_claim` - (Computed) - List(String) - A key=value pair that describes a required claim in the ID Token.<br />If set, the claim is verified to be present in the ID Token with a matching value.<br />Repeat this flag to specify multiple claims.
- `oidcca_file` - (Computed) - String - OIDCCAFile if set, the OpenID server's certificate will be verified by one<br />of the authorities in the oidc-ca-file.
- `proxy_client_cert_file` - (Computed) - String - The apiserver's client certificate used for outbound requests.
- `proxy_client_key_file` - (Computed) - String - The apiserver's client key used for outbound requests.
- `audit_log_format` - (Computed) - String - AuditLogFormat flag specifies the format type for audit log files.
- `audit_log_path` - (Computed) - String - If set, all requests coming to the apiserver will be logged to this file.
- `audit_log_max_age` - (Computed) - Int - The maximum number of days to retain old audit log files based on the timestamp encoded in their filename.
- `audit_log_max_backups` - (Computed) - Int - The maximum number of old audit log files to retain.
- `audit_log_max_size` - (Computed) - Int - The maximum size in megabytes of the audit log file before it gets rotated. Defaults to 100MB.
- `audit_policy_file` - (Computed) - String - AuditPolicyFile is the full path to a advanced audit configuration file e.g. /srv/kubernetes/audit.conf.
- `audit_webhook_batch_buffer_size` - (Computed) - Int - AuditWebhookBatchBufferSize is The size of the buffer to store events before batching and writing. Only used in batch mode. (default 10000).
- `audit_webhook_batch_max_size` - (Computed) - Int - AuditWebhookBatchMaxSize is The maximum size of a batch. Only used in batch mode. (default 400).
- `audit_webhook_batch_max_wait` - (Computed) - Duration - AuditWebhookBatchMaxWait is The amount of time to wait before force writing the batch that hadn't reached the max size. Only used in batch mode. (default 30s).
- `audit_webhook_batch_throttle_burst` - (Computed) - Int - AuditWebhookBatchThrottleBurst is Maximum number of requests sent at the same moment if ThrottleQPS was not utilized before. Only used in batch mode. (default 15).
- `audit_webhook_batch_throttle_enable` - (Computed) - Bool - AuditWebhookBatchThrottleEnable is Whether batching throttling is enabled. Only used in batch mode. (default true).
- `audit_webhook_batch_throttle_qps` - (Computed) - Quantity - AuditWebhookBatchThrottleQps is Maximum average number of batches per second. Only used in batch mode. (default 10).
- `audit_webhook_config_file` - (Computed) - String - AuditWebhookConfigFile is Path to a kubeconfig formatted file that defines the audit webhook configuration. Requires the 'AdvancedAuditing' feature gate.
- `audit_webhook_initial_backoff` - (Computed) - Duration - AuditWebhookInitialBackoff is The amount of time to wait before retrying the first failed request. (default 10s).
- `audit_webhook_mode` - (Computed) - String - AuditWebhookMode is Strategy for sending audit events. Blocking indicates sending events should block server responses. Batch causes the backend to buffer and write events asynchronously. Known modes are batch,blocking. (default "batch").
- `authentication_token_webhook_config_file` - (Computed) - String - File with webhook configuration for token authentication in kubeconfig format. The API server will query the remote service to determine authentication for bearer tokens.
- `authentication_token_webhook_cache_ttl` - (Computed) - Duration - The duration to cache responses from the webhook token authenticator. Default is 2m. (default 2m0s).
- `authorization_mode` - (Computed) - String - AuthorizationMode is the authorization mode the kubeapi is running in.
- `authorization_webhook_config_file` - (Computed) - String - File with webhook configuration for authorization in kubeconfig format. The API server will query the remote service to determine whether to authorize the request.
- `authorization_webhook_cache_authorized_ttl` - (Computed) - Duration - The duration to cache authorized responses from the webhook token authorizer. Default is 5m. (default 5m0s).
- `authorization_webhook_cache_unauthorized_ttl` - (Computed) - Duration - The duration to cache authorized responses from the webhook token authorizer. Default is 30s. (default 30s).
- `authorization_rbac_super_user` - (Computed) - String - AuthorizationRBACSuperUser is the name of the superuser for default rbac.
- `encryption_provider_config` - (Computed) - String - EncryptionProviderConfig enables encryption at rest for secrets.
- `experimental_encryption_provider_config` - (Computed) - String - ExperimentalEncryptionProviderConfig enables encryption at rest for secrets.
- `requestheader_username_headers` - (Computed) - List(String) - List of request headers to inspect for usernames. X-Remote-User is common.
- `requestheader_group_headers` - (Computed) - List(String) - List of request headers to inspect for groups. X-Remote-Group is suggested.
- `requestheader_extra_header_prefixes` - (Computed) - List(String) - List of request header prefixes to inspect. X-Remote-Extra- is suggested.
- `requestheader_client_ca_file` - (Computed) - String - Root certificate bundle to use to verify client certificates on incoming requests before trusting usernames in headers specified by --requestheader-username-headers.
- `requestheader_allowed_names` - (Computed) - List(String) - List of client certificate common names to allow to provide usernames in headers specified by --requestheader-username-headers. If empty, any client certificate validated by the authorities in --requestheader-client-ca-file is allowed.
- `feature_gates` - (Computed) - Map(String) - FeatureGates is set of key=value pairs that describe feature gates for alpha/experimental features.
- `max_requests_inflight` - (Computed) - Int - MaxRequestsInflight The maximum number of non-mutating requests in flight at a given time.
- `max_mutating_requests_inflight` - (Computed) - Int - MaxMutatingRequestsInflight The maximum number of mutating requests in flight at a given time. Defaults to 200.
- `http2_max_streams_per_connection` - (Computed) - Int - HTTP2MaxStreamsPerConnection sets the limit that the server gives to clients for the maximum number of streams in an HTTP/2 connection. Zero means to use golang's default.
- `etcd_quorum_read` - (Computed) - Bool - EtcdQuorumRead configures the etcd-quorum-read flag, which forces consistent reads from etcd.
- `request_timeout` - (Computed) - Duration - RequestTimeout configures the duration a handler must keep a request open before timing it out. (default 1m0s).
- `min_request_timeout` - (Computed) - Int - MinRequestTimeout configures the minimum number of seconds a handler must keep a request open before timing it out.<br />Currently only honored by the watch request handler.
- `target_ram_mb` - (Computed) - Int - Memory limit for apiserver in MB (used to configure sizes of caches, etc.).
- `service_account_key_file` - (Computed) - List(String) - File containing PEM-encoded x509 RSA or ECDSA private or public keys, used to verify ServiceAccount tokens.<br />The specified file can contain multiple keys, and the flag can be specified multiple times with different files.<br />If unspecified, --tls-private-key-file is used.
- `service_account_signing_key_file` - (Computed) - String - Path to the file that contains the current private key of the service account token issuer.<br />The issuer will sign issued ID tokens with this private key. (Requires the 'TokenRequest' feature gate.).
- `service_account_issuer` - (Computed) - String - Identifier of the service account token issuer. The issuer will assert this identifier<br />in "iss" claim of issued tokens. This value is a string or URI.
- `service_account_jwksuri` - (Computed) - String - ServiceAccountJWKSURI overrides the path for the jwks document; this is useful when we are republishing the service account discovery information elsewhere.
- `api_audiences` - (Computed) - List(String) - Identifiers of the API. The service account token authenticator will validate that<br />tokens used against the API are bound to at least one of these audiences. If the<br />--service-account-issuer flag is configured and this flag is not, this field<br />defaults to a single element list containing the issuer URL.
- `cpu_request` - (Computed) - String - CPURequest, cpu request compute resource for api server. Defaults to "150m".
- `cpu_limit` - (Computed) - String - CPULimit, cpu limit compute resource for api server e.g. "500m".
- `memory_request` - (Computed) - String - MemoryRequest, memory request compute resource for api server e.g. "30Mi".
- `memory_limit` - (Computed) - String - MemoryLimit, memory limit compute resource for api server e.g. "30Mi".
- `event_ttl` - (Computed) - Duration - Amount of time to retain Kubernetes events.
- `audit_dynamic_configuration` - (Computed) - Bool - AuditDynamicConfiguration enables dynamic audit configuration via AuditSinks.
- `enable_profiling` - (Computed) - Bool - EnableProfiling enables profiling via web interface host:port/debug/pprof/.
- `cors_allowed_origins` - (Computed) - List(String) - CorsAllowedOrigins is a list of origins for CORS. An allowed origin can be a regular<br />expression to support subdomain matching. If this list is empty CORS will not be enabled.
- `default_not_ready_toleration_seconds` - (Computed) - Int - DefaultNotReadyTolerationSeconds indicates the tolerationSeconds of the toleration for notReady:NoExecute that is added by default to every pod that does not already have such a toleration.
- `default_unreachable_toleration_seconds` - (Computed) - Int - DefaultUnreachableTolerationSeconds indicates the tolerationSeconds of the toleration for unreachable:NoExecute that is added by default to every pod that does not already have such a toleration.

### kube_controller_manager_config

KubeControllerManagerConfig is the configuration for the controller.

#### Argument Reference

The following arguments are supported:

- `master` - (Computed) - String - Master is the url for the kube api master.
- `log_level` - (Computed) - Int - LogLevel is the defined logLevel.
- `service_account_private_key_file` - (Computed) - String - ServiceAccountPrivateKeyFile is the location of the private key for service account token signing.
- `image` - (Computed) - String - Image is the docker image to use.
- `cloud_provider` - (Computed) - String - CloudProvider is the provider for cloud services.
- `cluster_name` - (Computed) - String - ClusterName is the instance prefix for the cluster.
- `cluster_cidr` - (Computed) - String - ClusterCIDR is CIDR Range for Pods in cluster.
- `allocate_node_cidrs` - (Computed) - Bool - AllocateNodeCIDRs enables CIDRs for Pods to be allocated and, if ConfigureCloudRoutes is true, to be set on the cloud provider.
- `node_cidr_mask_size` - (Computed) - Int - NodeCIDRMaskSize set the size for the mask of the nodes.
- `configure_cloud_routes` - (Computed) - Bool - ConfigureCloudRoutes enables CIDRs allocated with to be configured on the cloud provider.
- `controllers` - (Computed) - List(String) - Controllers is a list of controllers to enable on the controller-manager.
- `cidr_allocator_type` - (Computed) - String - CIDRAllocatorType specifies the type of CIDR allocator to use.
- `root_ca_file` - (Computed) - String - rootCAFile is the root certificate authority will be included in service account's token secret. This must be a valid PEM-encoded CA bundle.
- `leader_election` - (Computed) - [leader_election_configuration](#leader_election_configuration) - LeaderElection defines the configuration of leader election client.
- `attach_detach_reconcile_sync_period` - (Computed) - Duration - AttachDetachReconcileSyncPeriod is the amount of time the reconciler sync states loop<br />wait between successive executions. Is set to 1 min by kops by default.
- `disable_attach_detach_reconcile_sync` - (Computed) - Bool - DisableAttachDetachReconcileSync disables the reconcile sync loop in the attach-detach controller.<br />This can cause volumes to become mismatched with pods.
- `terminated_pod_gc_threshold` - (Computed) - Int - TerminatedPodGCThreshold is the number of terminated pods that can exist<br />before the terminated pod garbage collector starts deleting terminated pods.<br />If <= 0, the terminated pod garbage collector is disabled.
- `node_monitor_period` - (Computed) - Duration - NodeMonitorPeriod is the period for syncing NodeStatus in NodeController. (default 5s).
- `node_monitor_grace_period` - (Computed) - Duration - NodeMonitorGracePeriod is the amount of time which we allow running Node to be unresponsive before marking it unhealthy. (default 40s)<br />Must be N-1 times more than kubelet's nodeStatusUpdateFrequency, where N means number of retries allowed for kubelet to post node status.
- `pod_eviction_timeout` - (Computed) - Duration - PodEvictionTimeout is the grace period for deleting pods on failed nodes. (default 5m0s).
- `use_service_account_credentials` - (Computed) - Bool - UseServiceAccountCredentials controls whether we use individual service account credentials for each controller.
- `horizontal_pod_autoscaler_sync_period` - (Computed) - Duration - HorizontalPodAutoscalerSyncPeriod is the amount of time between syncs<br />During each period, the controller manager queries the resource utilization<br />against the metrics specified in each HorizontalPodAutoscaler definition.
- `horizontal_pod_autoscaler_downscale_delay` - (Computed) - Duration - HorizontalPodAutoscalerDownscaleDelay is a duration that specifies<br />how long the autoscaler has to wait before another downscale<br />operation can be performed after the current one has completed.
- `horizontal_pod_autoscaler_downscale_stabilization` - (Computed) - Duration - HorizontalPodAutoscalerDownscaleStabilization is the period for which<br />autoscaler will look backwards and not scale down below any<br />recommendation it made during that period.
- `horizontal_pod_autoscaler_upscale_delay` - (Computed) - Duration - HorizontalPodAutoscalerUpscaleDelay is a duration that specifies how<br />long the autoscaler has to wait before another upscale operation can<br />be performed after the current one has completed.
- `horizontal_pod_autoscaler_initial_readiness_delay` - (Computed) - Duration - HorizontalPodAutoscalerInitialReadinessDelay is the period after pod start<br />during which readiness changes will be treated as initial readiness. (default 30s).
- `horizontal_pod_autoscaler_cpu_initialization_period` - (Computed) - Duration - HorizontalPodAutoscalerCPUInitializationPeriod is the period after pod start<br />when CPU samples might be skipped. (default 5m).
- `horizontal_pod_autoscaler_tolerance` - (Computed) - Quantity - HorizontalPodAutoscalerTolerance is the minimum change (from 1.0) in the<br />desired-to-actual metrics ratio for the horizontal pod autoscaler to<br />consider scaling.
- `horizontal_pod_autoscaler_use_rest_clients` - (Computed) - Bool - HorizontalPodAutoscalerUseRestClients determines if the new-style clients<br />should be used if support for custom metrics is enabled.
- `experimental_cluster_signing_duration` - (Computed) - Duration - ExperimentalClusterSigningDuration is the duration that determines<br />the length of duration that the signed certificates will be given. (default 8760h0m0s).
- `feature_gates` - (Computed) - Map(String) - FeatureGates is set of key=value pairs that describe feature gates for alpha/experimental features.
- `tls_cipher_suites` - (Computed) - List(String) - TLSCipherSuites indicates the allowed TLS cipher suite.
- `tls_min_version` - (Computed) - String - TLSMinVersion indicates the minimum TLS version allowed.
- `min_resync_period` - (Computed) - String - MinResyncPeriod indicates the resync period in reflectors.<br />The resync period will be random between MinResyncPeriod and 2*MinResyncPeriod. (default 12h0m0s).
- `kube_api_qps` - (Computed) - Quantity - KubeAPIQPS QPS to use while talking with kubernetes apiserver. (default 20).
- `kube_api_burst` - (Computed) - Int - KubeAPIBurst Burst to use while talking with kubernetes apiserver. (default 30).
- `concurrent_deployment_syncs` - (Computed) - Int - The number of deployment objects that are allowed to sync concurrently.
- `concurrent_endpoint_syncs` - (Computed) - Int - The number of endpoint objects that are allowed to sync concurrently.
- `concurrent_namespace_syncs` - (Computed) - Int - The number of namespace objects that are allowed to sync concurrently.
- `concurrent_replicaset_syncs` - (Computed) - Int - The number of replicaset objects that are allowed to sync concurrently.
- `concurrent_service_syncs` - (Computed) - Int - The number of service objects that are allowed to sync concurrently.
- `concurrent_resource_quota_syncs` - (Computed) - Int - The number of resourcequota objects that are allowed to sync concurrently.
- `concurrent_serviceaccount_token_syncs` - (Computed) - Int - The number of serviceaccount objects that are allowed to sync concurrently to create tokens.
- `concurrent_rc_syncs` - (Computed) - Int - The number of replicationcontroller objects that are allowed to sync concurrently.
- `authentication_kubeconfig` - (Computed) - String - AuthenticationKubeconfig is the path to an Authentication Kubeconfig.
- `authorization_kubeconfig` - (Computed) - String - AuthorizationKubeconfig is the path to an Authorization Kubeconfig.
- `authorization_always_allow_paths` - (Computed) - List(String) - AuthorizationAlwaysAllowPaths is the list of HTTP paths to skip during authorization.
- `external_cloud_volume_plugin` - (Computed) - String - ExternalCloudVolumePlugin is a fallback mechanism that allows a legacy, in-tree cloudprovider to be used for volume plugins<br />even when an external cloud controller manager is being used.  This can be used instead of installing CSI.  The value should<br />be the same as is used for the --cloud-provider flag, i.e. "aws".
- `enable_profiling` - (Computed) - Bool - EnableProfiling enables profiling via web interface host:port/debug/pprof/.

### leader_election_configuration

LeaderElectionConfiguration defines the configuration of leader election<br />clients for components that can run with leader election enabled.

#### Argument Reference

The following arguments are supported:

- `leader_elect` - (Computed) - Bool - leaderElect enables a leader election client to gain leadership<br />before executing the main loop. Enable this when running replicated<br />components for high availability.
- `leader_elect_lease_duration` - (Computed) - Duration - leaderElectLeaseDuration is the length in time non-leader candidates<br />will wait after observing a leadership renewal until attempting to acquire<br />leadership of a led but unrenewed leader slot. This is effectively the<br />maximum duration that a leader can be stopped before it is replaced by another candidate.
- `leader_elect_renew_deadline_duration` - (Computed) - Duration - LeaderElectRenewDeadlineDuration is the interval between attempts by the acting master to<br />renew a leadership slot before it stops leading. This must be less than or equal to the lease duration.
- `leader_elect_resource_lock` - (Computed) - String - LeaderElectResourceLock is the type of resource object that is used for locking during<br />leader election. Supported options are endpoints (default) and `configmaps`.
- `leader_elect_resource_name` - (Computed) - String - LeaderElectResourceName is the name of resource object that is used for locking during leader election.
- `leader_elect_resource_namespace` - (Computed) - String - LeaderElectResourceNamespace is the namespace of resource object that is used for locking during leader election.
- `leader_elect_retry_period` - (Computed) - Duration - LeaderElectRetryPeriod is The duration the clients should wait between attempting acquisition<br />and renewal of a leadership. This is only applicable if leader election is enabled.

### cloud_controller_manager_config

CloudControllerManagerConfig is the configuration of the cloud controller.

#### Argument Reference

The following arguments are supported:

- `master` - (Computed) - String - Master is the url for the kube api master.
- `log_level` - (Computed) - Int - LogLevel is the verbosity of the logs.
- `image` - (Computed) - String - Image is the OCI image of the cloud controller manager.
- `cloud_provider` - (Computed) - String - CloudProvider is the provider for cloud services.
- `cluster_name` - (Computed) - String - ClusterName is the instance prefix for the cluster.
- `cluster_cidr` - (Computed) - String - ClusterCIDR is CIDR Range for Pods in cluster.
- `allocate_node_cidrs` - (Computed) - Bool - AllocateNodeCIDRs enables CIDRs for Pods to be allocated and, if<br />ConfigureCloudRoutes is true, to be set on the cloud provider.
- `configure_cloud_routes` - (Computed) - Bool - ConfigureCloudRoutes enables CIDRs allocated with to be configured on the cloud provider.
- `cidr_allocator_type` - (Computed) - String - CIDRAllocatorType specifies the type of CIDR allocator to use.
- `leader_election` - (Computed) - [leader_election_configuration](#leader_election_configuration) - LeaderElection defines the configuration of leader election client.
- `use_service_account_credentials` - (Computed) - Bool - UseServiceAccountCredentials controls whether we use individual service account credentials for each controller.

### kube_scheduler_config

KubeSchedulerConfig is the configuration for the kube-scheduler.

#### Argument Reference

The following arguments are supported:

- `master` - (Computed) - String - Master is a url to the kube master.
- `log_level` - (Computed) - Int - LogLevel is the logging level.
- `image` - (Computed) - String - Image is the docker image to use.
- `leader_election` - (Computed) - [leader_election_configuration](#leader_election_configuration) - LeaderElection defines the configuration of leader election client.
- `use_policy_config_map` - (Computed) - Bool - UsePolicyConfigMap enable setting the scheduler policy from a configmap.
- `feature_gates` - (Computed) - Map(String) - FeatureGates is set of key=value pairs that describe feature gates for alpha/experimental features.
- `max_persistent_volumes` - (Computed) - Int - MaxPersistentVolumes changes the maximum number of persistent volumes the scheduler will scheduler onto the same<br />node. Only takes into affect if value is positive. This corresponds to the KUBE_MAX_PD_VOLS environment variable,<br />which has been supported as far back as Kubernetes 1.7. The default depends on the version and the cloud provider<br />as outlined: https://kubernetes.io/docs/concepts/storage/storage-limits/.
- `qps` - (Computed) - Quantity - Qps sets the maximum qps to send to apiserver after the burst quota is exhausted.
- `burst` - (Computed) - Int - Burst sets the maximum qps to send to apiserver after the burst quota is exhausted.
- `authentication_kubeconfig` - (Computed) - String - AuthenticationKubeconfig is the path to an Authentication Kubeconfig.
- `authorization_kubeconfig` - (Computed) - String - AuthorizationKubeconfig is the path to an Authorization Kubeconfig.
- `authorization_always_allow_paths` - (Computed) - List(String) - AuthorizationAlwaysAllowPaths is the list of HTTP paths to skip during authorization.
- `enable_profiling` - (Computed) - Bool - EnableProfiling enables profiling via web interface host:port/debug/pprof/.

### kube_proxy_config

KubeProxyConfig defines the configuration for a proxy.

#### Argument Reference

The following arguments are supported:

- `image` - (Computed) - String
- `cpu_request` - (Computed) - String - TODO: Better type ?<br />CPURequest, cpu request compute resource for kube proxy e.g. "20m".
- `cpu_limit` - (Computed) - String - CPULimit, cpu limit compute resource for kube proxy e.g. "30m".
- `memory_request` - (Computed) - String - MemoryRequest, memory request compute resource for kube proxy e.g. "30Mi".
- `memory_limit` - (Computed) - String - MemoryLimit, memory limit compute resource for kube proxy e.g. "30Mi".
- `log_level` - (Computed) - Int - LogLevel is the logging level of the proxy.
- `cluster_cidr` - (Computed) - String - ClusterCIDR is the CIDR range of the pods in the cluster.
- `hostname_override` - (Computed) - String - HostnameOverride, if non-empty, will be used as the identity instead of the actual hostname.
- `bind_address` - (Computed) - String - BindAddress is IP address for the proxy server to serve on.
- `master` - (Computed) - String - Master is the address of the Kubernetes API server (overrides any value in kubeconfig).
- `metrics_bind_address` - (Computed) - String - MetricsBindAddress is the IP address for the metrics server to serve on.
- `enabled` - (Computed) - Bool - Enabled allows enabling or disabling kube-proxy.
- `proxy_mode` - (Computed) - String - Which proxy mode to use: (userspace, iptables(default), ipvs).
- `ip_vs_exclude_cidr_s` - (Computed) - List(String) - IPVSExcludeCIDRS is comma-separated list of CIDR's which the ipvs proxier should not touch when cleaning up IPVS rules.
- `ip_vs_min_sync_period` - (Computed) - Duration - IPVSMinSyncPeriod is the minimum interval of how often the ipvs rules can be refreshed as endpoints and services change (e.g. '5s', '1m', '2h22m').
- `ip_vs_scheduler` - (Computed) - String - IPVSScheduler is the ipvs scheduler type when proxy mode is ipvs.
- `ip_vs_sync_period` - (Computed) - Duration - IPVSSyncPeriod duration is the maximum interval of how often ipvs rules are refreshed.
- `feature_gates` - (Computed) - Map(String) - FeatureGates is a series of key pairs used to switch on features for the proxy.
- `conntrack_max_per_core` - (Computed) - Int - Maximum number of NAT connections to track per CPU core (default: 131072).
- `conntrack_min` - (Computed) - Int - Minimum number of conntrack entries to allocate, regardless of conntrack-max-per-core.

### kubelet_config_spec

KubeletConfigSpec defines the kubelet configuration.

#### Argument Reference

The following arguments are supported:

- `api_servers` - (Computed) - String - APIServers is not used for clusters version 1.6 and later - flag removed.
- `anonymous_auth` - (Computed) - Bool - AnonymousAuth permits you to control auth to the kubelet api.
- `authorization_mode` - (Computed) - String - AuthorizationMode is the authorization mode the kubelet is running in.
- `bootstrap_kubeconfig` - (Computed) - String - BootstrapKubeconfig is the path to a kubeconfig file that will be used to get client certificate for kubelet.
- `client_ca_file` - (Computed) - String - ClientCAFile is the path to a CA certificate.
- `tls_cert_file` - (Computed) - String - TODO: Remove unused TLSCertFile.
- `tls_private_key_file` - (Computed) - String - TODO: Remove unused TLSPrivateKeyFile.
- `tls_cipher_suites` - (Computed) - List(String) - TLSCipherSuites indicates the allowed TLS cipher suite.
- `tls_min_version` - (Computed) - String - TLSMinVersion indicates the minimum TLS version allowed.
- `kubeconfig_path` - (Computed) - String - KubeconfigPath is the path of kubeconfig for the kubelet.
- `require_kubeconfig` - (Computed) - Bool - RequireKubeconfig indicates a kubeconfig is required.
- `log_level` - (Computed) - Int - LogLevel is the logging level of the kubelet.
- `pod_manifest_path` - (Computed) - String - config is the path to the config file or directory of files.
- `hostname_override` - (Computed) - String - HostnameOverride is the hostname used to identify the kubelet instead of the actual hostname.
- `pod_infra_container_image` - (Computed) - String - PodInfraContainerImage is the image whose network/ipc containers in each pod will use.
- `seccomp_profile_root` - (Computed) - String - SeccompProfileRoot is the directory path for seccomp profiles.
- `allow_privileged` - (Computed) - Bool - AllowPrivileged enables containers to request privileged mode (defaults to false).
- `enable_debugging_handlers` - (Computed) - Bool - EnableDebuggingHandlers enables server endpoints for log collection and local running of containers and commands.
- `register_node` - (Computed) - Bool - RegisterNode enables automatic registration with the apiserver.
- `node_status_update_frequency` - (Computed) - Duration - NodeStatusUpdateFrequency Specifies how often kubelet posts node status to master (default 10s)<br />must work with nodeMonitorGracePeriod in KubeControllerManagerConfig.
- `cluster_domain` - (Computed) - String - ClusterDomain is the DNS domain for this cluster.
- `cluster_dns` - (Computed) - String - ClusterDNS is the IP address for a cluster DNS server.
- `network_plugin_name` - (Computed) - String - NetworkPluginName is the name of the network plugin to be invoked for various events in kubelet/pod lifecycle.
- `cloud_provider` - (Computed) - String - CloudProvider is the provider for cloud services.
- `kubelet_cgroups` - (Computed) - String - KubeletCgroups is the absolute name of cgroups to isolate the kubelet in.
- `runtime_cgroups` - (Computed) - String - Cgroups that container runtime is expected to be isolated in.
- `read_only_port` - (Computed) - Int - ReadOnlyPort is the port used by the kubelet api for read-only access (default 10255).
- `system_cgroups` - (Computed) - String - SystemCgroups is absolute name of cgroups in which to place<br />all non-kernel processes that are not already in a container. Empty<br />for no container. Rolling back the flag requires a reboot.
- `cgroup_root` - (Computed) - String - cgroupRoot is the root cgroup to use for pods. This is handled by the container runtime on a best effort basis.
- `configure_cbr0` - (Computed) - Bool - configureCBR0 enables the kubelet to configure cbr0 based on Node.Spec.PodCIDR.
- `hairpin_mode` - (Computed) - String - How should the kubelet configure the container bridge for hairpin packets.<br />Setting this flag allows endpoints in a Service to loadbalance back to<br />themselves if they should try to access their own Service. Values:<br />  "promiscuous-bridge": make the container bridge promiscuous.<br />  "hairpin-veth":       set the hairpin flag on container veth interfaces.<br />  "none":               do nothing.<br />Setting --configure-cbr0 to false implies that to achieve hairpin NAT<br />one must set --hairpin-mode=veth-flag, because bridge assumes the<br />existence of a container bridge named cbr0.
- `babysit_daemons` - (Computed) - Bool - The node has babysitter process monitoring docker and kubelet. Removed as of 1.7.
- `max_pods` - (Computed) - Int - MaxPods is the number of pods that can run on this Kubelet.
- `nvidia_gp_us` - (Computed) - Int - NvidiaGPUs is the number of NVIDIA GPU devices on this node.
- `pod_cidr` - (Computed) - String - PodCIDR is the CIDR to use for pod IP addresses, only used in standalone mode.<br />In cluster mode, this is obtained from the master.
- `resolver_config` - (Computed) - String - ResolverConfig is the resolver configuration file used as the basis for the container DNS resolution configuration."), [].
- `reconcile_cidr` - (Computed) - Bool - ReconcileCIDR is Reconcile node CIDR with the CIDR specified by the<br />API server. No-op if register-node or configure-cbr0 is false.
- `register_schedulable` - (Computed) - Bool - registerSchedulable tells the kubelet to register the node as schedulable. No-op if register-node is false.
- `serialize_image_pulls` - (Computed) - Bool - // SerializeImagePulls when enabled, tells the Kubelet to pull images one<br />// at a time. We recommend *not* changing the default value on nodes that<br />// run docker daemon with version  < 1.9 or an Aufs storage backend.<br />// Issue #10959 has more details.
- `node_labels` - (Computed) - Map(String) - NodeLabels to add when registering the node in the cluster.
- `non_masquerade_cidr` - (Computed) - String - NonMasqueradeCIDR configures masquerading: traffic to IPs outside this range will use IP masquerade.
- `enable_custom_metrics` - (Computed) - Bool - Enable gathering custom metrics.
- `network_plugin_mtu` - (Computed) - Int - NetworkPluginMTU is the MTU to be passed to the network plugin,<br />and overrides the default MTU for cases where it cannot be automatically<br />computed (such as IPSEC).
- `image_gc_high_threshold_percent` - (Computed) - Int - ImageGCHighThresholdPercent is the percent of disk usage after which<br />image garbage collection is always run.
- `image_gc_low_threshold_percent` - (Computed) - Int - ImageGCLowThresholdPercent is the percent of disk usage before which<br />image garbage collection is never run. Lowest disk usage to garbage<br />collect to.
- `image_pull_progress_deadline` - (Computed) - Duration - ImagePullProgressDeadline is the timeout for image pulls<br />If no pulling progress is made before this deadline, the image pulling will be cancelled. (default 1m0s).
- `eviction_hard` - (Computed) - String - Comma-delimited list of hard eviction expressions.  For example, 'memory.available<300Mi'.
- `eviction_soft` - (Computed) - String - Comma-delimited list of soft eviction expressions.  For example, 'memory.available<300Mi'.
- `eviction_soft_grace_period` - (Computed) - String - Comma-delimited list of grace periods for each soft eviction signal.  For example, 'memory.available=30s'.
- `eviction_pressure_transition_period` - (Computed) - Duration - Duration for which the kubelet has to wait before transitioning out of an eviction pressure condition.
- `eviction_max_pod_grace_period` - (Computed) - Int - Maximum allowed grace period (in seconds) to use when terminating pods in response to a soft eviction threshold being met.
- `eviction_minimum_reclaim` - (Computed) - String - Comma-delimited list of minimum reclaims (e.g. imagefs.available=2Gi) that describes the minimum amount of resource the kubelet will reclaim when performing a pod eviction if that resource is under pressure.
- `volume_plugin_directory` - (Computed) - String - The full path of the directory in which to search for additional third party volume plugins (this path must be writeable, dependent on your choice of OS).
- `taints` - (Computed) - List(String) - Taints to add when registering a node in the cluster.
- `feature_gates` - (Computed) - Map(String) - FeatureGates is set of key=value pairs that describe feature gates for alpha/experimental features.
- `kube_reserved` - (Computed) - Map(String) - Resource reservation for kubernetes system daemons like the kubelet, container runtime, node problem detector, etc.
- `kube_reserved_cgroup` - (Computed) - String - Control group for kube daemons.
- `system_reserved` - (Computed) - Map(String) - Capture resource reservation for OS system daemons like sshd, udev, etc.
- `system_reserved_cgroup` - (Computed) - String - Parent control group for OS system daemons.
- `enforce_node_allocatable` - (Computed) - String - Enforce Allocatable across pods whenever the overall usage across all pods exceeds Allocatable.
- `runtime_request_timeout` - (Computed) - Duration - RuntimeRequestTimeout is timeout for runtime requests on - pull, logs, exec and attach.
- `volume_stats_agg_period` - (Computed) - Duration - VolumeStatsAggPeriod is the interval for kubelet to calculate and cache the volume disk usage for all pods and volumes.
- `fail_swap_on` - (Computed) - Bool - Tells the Kubelet to fail to start if swap is enabled on the node.
- `experimental_allowed_unsafe_sysctls` - (Computed) - List(String) - ExperimentalAllowedUnsafeSysctls are passed to the kubelet config to whitelist allowable sysctls<br />Was promoted to beta and renamed. https://github.com/kubernetes/kubernetes/pull/63717.
- `allowed_unsafe_sysctls` - (Computed) - List(String) - AllowedUnsafeSysctls are passed to the kubelet config to whitelist allowable sysctls.
- `streaming_connection_idle_timeout` - (Computed) - Duration - StreamingConnectionIdleTimeout is the maximum time a streaming connection can be idle before the connection is automatically closed.
- `docker_disable_shared_pid` - (Computed) - Bool - DockerDisableSharedPID uses a shared PID namespace for containers in a pod.
- `root_dir` - (Computed) - String - RootDir is the directory path for managing kubelet files (volume mounts,etc).
- `authentication_token_webhook` - (Computed) - Bool - AuthenticationTokenWebhook uses the TokenReview API to determine authentication for bearer tokens.
- `authentication_token_webhook_cache_ttl` - (Computed) - Duration - AuthenticationTokenWebhook sets the duration to cache responses from the webhook token authenticator. Default is 2m. (default 2m0s).
- `cpucfs_quota` - (Computed) - Bool - CPUCFSQuota enables CPU CFS quota enforcement for containers that specify CPU limits.
- `cpucfs_quota_period` - (Computed) - Duration - CPUCFSQuotaPeriod sets CPU CFS quota period value, cpu.cfs_period_us, defaults to Linux Kernel default.
- `cpu_manager_policy` - (Computed) - String - CpuManagerPolicy allows for changing the default policy of None to static.
- `registry_pull_qps` - (Computed) - Int - RegistryPullQPS if > 0, limit registry pull QPS to this value.  If 0, unlimited. (default 5).
- `registry_burst` - (Computed) - Int - RegistryBurst Maximum size of a bursty pulls, temporarily allows pulls to burst to this number, while still not exceeding registry-qps. Only used if --registry-qps > 0 (default 10).
- `topology_manager_policy` - (Computed) - String - TopologyManagerPolicy determines the allocation policy for the topology manager.
- `rotate_certificates` - (Computed) - Bool - rotateCertificates enables client certificate rotation.
- `protect_kernel_defaults` - (Computed) - Bool - Default kubelet behaviour for kernel tuning. If set, kubelet errors if any of kernel tunables is different than kubelet defaults.<br />(DEPRECATED: This parameter should be set via the config file specified by the Kubelet's --config flag.
- `cgroup_driver` - (Computed) - String - CgroupDriver allows the explicit setting of the kubelet cgroup driver. If omitted, defaults to cgroupfs.
- `housekeeping_interval` - (Computed) - Duration - HousekeepingInterval allows to specify interval between container housekeepings.
- `event_qps` - (Computed) - Int - EventQPS if > 0, limit event creations per second to this value.  If 0, unlimited.
- `event_burst` - (Computed) - Int - EventBurst temporarily allows event records to burst to this number, while still not exceeding EventQPS. Only used if EventQPS > 0.
- `container_log_max_size` - (Computed) - String - ContainerLogMaxSize is the maximum size (e.g. 10Mi) of container log file before it is rotated.
- `container_log_max_files` - (Computed) - Int - ContainerLogMaxFiles is the maximum number of container log files that can be present for a container. The number must be >= 2.
- `enable_cadvisor_json_endpoints` - (Computed) - Bool - EnableCadvisorJsonEndpoints enables cAdvisor json `/spec` and `/stats/*` endpoints. Defaults to False.

### cloud_configuration

CloudConfiguration defines the cloud provider configuration.

#### Argument Reference

The following arguments are supported:

- `manage_storage_classes` - (Computed) - Bool - ManageStorageClasses specifies whether kOps should create and maintain a set of<br />StorageClasses, one of which it nominates as the default class for the cluster.
- `multizone` - (Computed) - Bool - GCE cloud-config options.
- `node_tags` - (Computed) - String
- `node_instance_prefix` - (Computed) - String
- `gce_service_account` - (Computed) - String - GCEServiceAccount specifies the service account with which the GCE VM runs.
- `disable_security_group_ingress` - (Computed) - Bool - AWS cloud-config options.
- `elb_security_group` - (Computed) - String
- `v_sphere_username` - (Computed) - String - VSphereUsername is deprecated and will be removed in a later version.
- `v_sphere_password` - (Computed) - String - VSpherePassword is deprecated and will be removed in a later version.
- `v_sphere_server` - (Computed) - String - VSphereServer is deprecated and will be removed in a later version.
- `v_sphere_datacenter` - (Computed) - String - VShpereDatacenter is deprecated and will be removed in a later version.
- `v_sphere_resource_pool` - (Computed) - String - VSphereResourcePool is deprecated and will be removed in a later version.
- `v_sphere_datastore` - (Computed) - String - VSphereDatastore is deprecated and will be removed in a later version.
- `v_sphere_core_dns_server` - (Computed) - String - VSphereCoreDNSServer is deprecated and will be removed in a later version.
- `spotinst_product` - (Computed) - String - Spotinst cloud-config specs.
- `spotinst_orientation` - (Computed) - String
- `openstack` - (Computed) - [openstack_configuration](#openstack_configuration) - Openstack cloud-config options.
- `azure` - (Computed) - [azure_configuration](#azure_configuration) - Azure cloud-config options.
- `awsebscsi_driver` - (Computed) - [awsebscsi_driver](#awsebscsi_driver) - AWSEBSCSIDriver is the config for the AWS EBS CSI driver.

### openstack_configuration

OpenstackConfiguration defines cloud config elements for the openstack cloud provider.

#### Argument Reference

The following arguments are supported:

- `loadbalancer` - (Computed) - [openstack_loadbalancer_config](#openstack_loadbalancer_config)
- `monitor` - (Computed) - [openstack_monitor](#openstack_monitor)
- `router` - (Computed) - [openstack_router](#openstack_router)
- `block_storage` - (Computed) - [openstack_block_storage_config](#openstack_block_storage_config)
- `insecure_skip_verify` - (Computed) - Bool
- `network` - (Computed) - [openstack_network](#openstack_network)

### openstack_loadbalancer_config

OpenstackLoadbalancerConfig defines the config for a neutron loadbalancer.

#### Argument Reference

The following arguments are supported:

- `method` - (Computed) - String
- `provider` - (Computed) - String
- `use_octavia` - (Computed) - Bool
- `floating_network` - (Computed) - String
- `floating_network_id` - (Computed) - String
- `floating_subnet` - (Computed) - String
- `subnet_id` - (Computed) - String
- `manage_sec_groups` - (Computed) - Bool

### openstack_monitor

OpenstackMonitor defines the config for a health monitor.

#### Argument Reference

The following arguments are supported:

- `delay` - (Computed) - String
- `timeout` - (Computed) - String
- `max_retries` - (Computed) - Int

### openstack_router

OpenstackRouter defines the config for a router.

#### Argument Reference

The following arguments are supported:

- `external_network` - (Computed) - String
- `dns_servers` - (Computed) - String
- `external_subnet` - (Computed) - String
- `availability_zone_hints` - (Computed) - List(String)

### openstack_block_storage_config

#### Argument Reference

The following arguments are supported:

- `version` - (Computed) - String
- `ignore_az` - (Computed) - Bool
- `override_az` - (Computed) - String
- `create_storage_class` - (Computed) - Bool - CreateStorageClass provisions a default class for the Cinder plugin.
- `cs_iplugin_image` - (Computed) - String
- `csi_topology_support` - (Computed) - Bool

### openstack_network

OpenstackNetwork defines the config for a network.

#### Argument Reference

The following arguments are supported:

- `availability_zone_hints` - (Computed) - List(String)

### azure_configuration

AzureConfiguration defines Azure specific cluster configuration.

#### Argument Reference

The following arguments are supported:

- `subscription_id` - (Computed) - String - SubscriptionID specifies the subscription used for the cluster installation.
- `tenant_id` - (Computed) - String - TenantID is the ID of the tenant that the cluster is deployed in.
- `resource_group_name` - (Computed) - String - ResourceGroupName specifies the name of the resource group<br />where the cluster is built.<br />If this is empty, kops will create a new resource group<br />whose name is same as the cluster name. If this is not<br />empty, kops will not create a new resource group, and<br />it will just reuse the existing resource group of the name.<br />This follows the model that kops takes for AWS VPC.
- `route_table_name` - (Computed) - String - RouteTableName is the name of the route table attached to the subnet that the cluster is deployed in.
- `admin_user` - (Computed) - String - AdminUser specifies the admin user of VMs.

### awsebscsi_driver

AWSEBSCSIDriver is the config for the AWS EBS CSI driver.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Computed) - Bool - Enabled enables the AWS EBS CSI driver<br />Default: false.
- `version` - (Computed) - String - Version is the container image tag used.<br />Default: The latest stable release which is compatible with your Kubernetes version.
- `volume_attach_limit` - (Computed) - Int - VolumeAttachLimit is the maximum number of volumes attachable per node.<br />If specified, the limit applies to all nodes.<br />If not specified, the value is approximated from the instance type.<br />Default: -.

### external_dns_config

ExternalDNSConfig are options of the dns-controller.

#### Argument Reference

The following arguments are supported:

- `disable` - (Computed) - Bool - Disable indicates we do not wish to run the dns-controller addon.
- `watch_ingress` - (Computed) - Bool - WatchIngress indicates you want the dns-controller to watch and create dns entries for ingress resources.
- `watch_namespace` - (Computed) - String - WatchNamespace is namespace to watch, defaults to all (use to control whom can creates dns entries).

### ntp_config

NTPConfig is the configuration for NTP.

#### Argument Reference

The following arguments are supported:

- `managed` - (Computed) - Bool - Managed controls if the NTP configuration is managed by kOps.<br />The NTP configuration task is skipped if this is set to false.

### node_termination_handler_config

NodeTerminationHandlerConfig determines the node termination handler configuration.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Computed) - Bool - Enabled enables the node termination handler.<br />Default: true.
- `enable_spot_interruption_draining` - (Computed) - Bool - EnableSpotInterruptionDraining makes node termination handler drain nodes when spot interruption termination notice is received.<br />Default: true.
- `enable_scheduled_event_draining` - (Computed) - Bool - EnableScheduledEventDraining makes node termination handler drain nodes before the maintenance window starts for an EC2 instance scheduled event.<br />Default: false.
- `enable_prometheus_metrics` - (Computed) - Bool - EnablePrometheusMetrics enables the "/metrics" endpoint.

### metrics_server_config

MetricsServerConfig determines the metrics server configuration.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Computed) - Bool - Enabled enables the metrics server.<br />Default: false.
- `image` - (Computed) - String - Image is the docker container used.<br />Default: the latest supported image for the specified kubernetes version.
- `insecure` - (Computed) - Bool - Insecure determines if API server will validate metrics server TLS cert.<br />Default: true.

### cert_manager_config

CertManagerConfig determines the cert manager configuration.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Computed) - Bool - Enabled enables the cert manager.<br />Default: false.
- `image` - (Computed) - String - Image is the docker container used.<br />Default: the latest supported image for the specified kubernetes version.
- `default_issuer` - (Computed) - String - defaultIssuer sets a default clusterIssuer<br />Default: none.

### aws_load_balancer_controller_config

AWSLoadBalancerControllerConfig determines the AWS LB controller configuration.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Computed) - Bool - Enabled enables the loadbalancer controller.<br />Default: false.
- `version` - (Computed) - String - Version is the container image tag used.

### networking_spec

NetworkingSpec allows selection and configuration of a networking plugin.

#### Argument Reference

The following arguments are supported:

- `classic` - (Computed) - [classic_networking_spec](#classic_networking_spec)
- `kubenet` - (Computed) - [kubenet_networking_spec](#kubenet_networking_spec)
- `external` - (Computed) - [external_networking_spec](#external_networking_spec)
- `cni` - (Computed) - [cni_networking_spec](#cni_networking_spec)
- `kopeio` - (Computed) - [kopeio_networking_spec](#kopeio_networking_spec)
- `weave` - (Computed) - [weave_networking_spec](#weave_networking_spec)
- `flannel` - (Computed) - [flannel_networking_spec](#flannel_networking_spec)
- `calico` - (Computed) - [calico_networking_spec](#calico_networking_spec)
- `canal` - (Computed) - [canal_networking_spec](#canal_networking_spec)
- `kuberouter` - (Computed) - [kuberouter_networking_spec](#kuberouter_networking_spec)
- `romana` - (Computed) - [romana_networking_spec](#romana_networking_spec)
- `amazon_vpc` - (Computed) - [amazon_vpc_networking_spec](#amazon_vpc_networking_spec)
- `cilium` - (Computed) - [cilium_networking_spec](#cilium_networking_spec)
- `lyft_vpc` - (Computed) - [lyft_vpc_networking_spec](#lyft_vpc_networking_spec)
- `gce` - (Computed) - [gce_networking_spec](#gce_networking_spec)

### classic_networking_spec

ClassicNetworkingSpec is the specification of classic networking mode, integrated into kubernetes.<br />Support been removed since Kubernetes 1.4.


This resource has no attributes.

### kubenet_networking_spec

KubenetNetworkingSpec is the specification for kubenet networking, largely integrated but intended to replace classic.


This resource has no attributes.

### external_networking_spec

ExternalNetworkingSpec is the specification for networking that is implemented by a user-provided Daemonset that uses the Kubenet kubelet networking plugin.


This resource has no attributes.

### cni_networking_spec

CNINetworkingSpec is the specification for networking that is implemented by a user-provided Daemonset, which uses the CNI kubelet networking plugin.

#### Argument Reference

The following arguments are supported:

- `uses_secondary_ip` - (Computed) - Bool

### kopeio_networking_spec

KopeioNetworkingSpec declares that we want Kopeio networking.


This resource has no attributes.

### weave_networking_spec

WeaveNetworkingSpec declares that we want Weave networking.

#### Argument Reference

The following arguments are supported:

- `mtu` - (Computed) - Int
- `conn_limit` - (Computed) - Int
- `no_masq_local` - (Computed) - Int
- `memory_request` - (Computed) - Quantity - MemoryRequest memory request of weave container. Default 200Mi.
- `cpu_request` - (Computed) - Quantity - CPURequest CPU request of weave container. Default 50m.
- `memory_limit` - (Computed) - Quantity - MemoryLimit memory limit of weave container. Default 200Mi.
- `cpu_limit` - (Computed) - Quantity - CPULimit CPU limit of weave container.
- `net_extra_args` - (Computed) - String - NetExtraArgs are extra arguments that are passed to weave-kube.
- `npc_memory_request` - (Computed) - Quantity - NPCMemoryRequest memory request of weave npc container. Default 200Mi.
- `npccpu_request` - (Computed) - Quantity - NPCCPURequest CPU request of weave npc container. Default 50m.
- `npc_memory_limit` - (Computed) - Quantity - NPCMemoryLimit memory limit of weave npc container. Default 200Mi.
- `npccpu_limit` - (Computed) - Quantity - NPCCPULimit CPU limit of weave npc container.
- `npc_extra_args` - (Computed) - String - NPCExtraArgs are extra arguments that are passed to weave-npc.
- `version` - (Computed) - String - Version specifies the Weave container image tag. The default depends on the kOps version.

### flannel_networking_spec

FlannelNetworkingSpec declares that we want Flannel networking.

#### Argument Reference

The following arguments are supported:

- `backend` - (Computed) - String - Backend is the backend overlay type we want to use (vxlan or udp).
- `disable_tx_checksum_offloading` - (Computed) - Bool - DisableTxChecksumOffloading is deprecated as of kops 1.19 and has no effect.
- `iptables_resync_seconds` - (Computed) - Int - IptablesResyncSeconds sets resync period for iptables rules, in seconds.

### calico_networking_spec

CalicoNetworkingSpec declares that we want Calico networking.

#### Argument Reference

The following arguments are supported:

- `registry` - (Computed) - String - Version overrides the Calico container image registry.
- `version` - (Computed) - String - Version overrides the Calico container image tag.
- `aws_src_dst_check` - (Computed) - String - AWSSrcDstCheck enables/disables ENI source/destination checks (AWS only)<br />Options: DoNothing (default), Enable, or Disable.
- `bpf_enabled` - (Computed) - Bool - BPFEnabled enables the eBPF dataplane mode.
- `bpf_external_service_mode` - (Computed) - String - BPFExternalServiceMode controls how traffic from outside the cluster to NodePorts and ClusterIPs is handled.<br />In Tunnel mode, packet is tunneled from the ingress host to the host with the backing pod and back again.<br />In DSR mode, traffic is tunneled to the host with the backing pod and then returned directly;<br />this requires a network that allows direct return.<br />Default: Tunnel (other options: DSR).
- `bpf_kube_proxy_iptables_cleanup_enabled` - (Computed) - Bool - BPFKubeProxyIptablesCleanupEnabled controls whether Felix will clean up the iptables rules<br />created by the Kubernetes kube-proxy; should only be enabled if kube-proxy is not running.
- `bpf_log_level` - (Computed) - String - BPFLogLevel controls the log level used by the BPF programs. The logs are emitted<br />to the BPF trace pipe, accessible with the command tc exec BPF debug.<br />Default: Off (other options: Info, Debug).
- `chain_insert_mode` - (Computed) - String - ChainInsertMode controls whether Felix inserts rules to the top of iptables chains, or<br />appends to the bottom. Leaving the default option is safest to prevent accidentally<br />breaking connectivity. Default: 'insert' (other options: 'append').
- `cpu_request` - (Computed) - Quantity - CPURequest CPU request of Calico container. Default: 100m.
- `cross_subnet` - (Computed) - Bool - CrossSubnet enables Calico's cross-subnet mode when set to true.
- `encapsulation_mode` - (Computed) - String - EncapsulationMode specifies the network packet encapsulation protocol for Calico to use,<br />employing such encapsulation at the necessary scope per the related CrossSubnet field. In<br />"ipip" mode, Calico will use IP-in-IP encapsulation as needed. In "vxlan" mode, Calico will<br />encapsulate packets as needed using the VXLAN scheme.<br />Options: ipip (default) or vxlan.
- `ip_ip_mode` - (Computed) - String - IPIPMode is the encapsulation mode to use for the default Calico IPv4 pool created at start<br />up, determining when to use IP-in-IP encapsulation, conveyed to the "calico-node" daemon<br />container via the CALICO_IPV4POOL_IPIP environment variable.
- `ipv4_auto_detection_method` - (Computed) - String - IPv4AutoDetectionMethod configures how Calico chooses the IP address used to route<br />between nodes.  This should be set when the host has multiple interfaces<br />and it is important to select the interface used.<br />Options: "first-found" (default), "can-reach=DESTINATION",<br />"interface=INTERFACE-REGEX", or "skip-interface=INTERFACE-REGEX".
- `ipv6_auto_detection_method` - (Computed) - String - IPv6AutoDetectionMethod configures how Calico chooses the IP address used to route<br />between nodes.  This should be set when the host has multiple interfaces<br />and it is important to select the interface used.<br />Options: "first-found" (default), "can-reach=DESTINATION",<br />"interface=INTERFACE-REGEX", or "skip-interface=INTERFACE-REGEX".
- `iptables_backend` - (Computed) - String - IptablesBackend controls which variant of iptables binary Felix uses<br />Default: Auto (other options: Legacy, NFT).
- `log_severity_screen` - (Computed) - String - LogSeverityScreen lets us set the desired log level. (Default: info).
- `mtu` - (Computed) - Int - MTU to be set in the cni-network-config for calico.
- `prometheus_metrics_enabled` - (Computed) - Bool - PrometheusMetricsEnabled can be set to enable the experimental Prometheus<br />metrics server (default: false).
- `prometheus_metrics_port` - (Computed) - Int - PrometheusMetricsPort is the TCP port that the experimental Prometheus<br />metrics server should bind to (default: 9091).
- `prometheus_go_metrics_enabled` - (Computed) - Bool - PrometheusGoMetricsEnabled enables Prometheus Go runtime metrics collection.
- `prometheus_process_metrics_enabled` - (Computed) - Bool - PrometheusProcessMetricsEnabled enables Prometheus process metrics collection.
- `major_version` - (Computed) - String - MajorVersion is deprecated as of kOps 1.20 and has no effect.
- `typha_prometheus_metrics_enabled` - (Computed) - Bool - TyphaPrometheusMetricsEnabled enables Prometheus metrics collection from Typha<br />(default: false).
- `typha_prometheus_metrics_port` - (Computed) - Int - TyphaPrometheusMetricsPort is the TCP port the typha Prometheus metrics server<br />should bind to (default: 9093).
- `typha_replicas` - (Computed) - Int - TyphaReplicas is the number of replicas of Typha to deploy.
- `wireguard_enabled` - (Computed) - Bool - WireguardEnabled enables WireGuard encryption for all on-the-wire pod-to-pod traffic<br />(default: false).

### canal_networking_spec

CanalNetworkingSpec declares that we want Canal networking.

#### Argument Reference

The following arguments are supported:

- `chain_insert_mode` - (Computed) - String - ChainInsertMode controls whether Felix inserts rules to the top of iptables chains, or<br />appends to the bottom. Leaving the default option is safest to prevent accidentally<br />breaking connectivity. Default: 'insert' (other options: 'append').
- `cpu_request` - (Computed) - Quantity - CPURequest CPU request of Canal container. Default: 100m.
- `default_endpoint_to_host_action` - (Computed) - String - DefaultEndpointToHostAction allows users to configure the default behaviour<br />for traffic between pod to host after calico rules have been processed.<br />Default: ACCEPT (other options: DROP, RETURN).
- `disable_flannel_forward_rules` - (Computed) - Bool - DisableFlannelForwardRules configures Flannel to NOT add the<br />default ACCEPT traffic rules to the iptables FORWARD chain.
- `disable_tx_checksum_offloading` - (Computed) - Bool - DisableTxChecksumOffloading is deprecated as of kops 1.19 and has no effect.
- `iptables_backend` - (Computed) - String - IptablesBackend controls which variant of iptables binary Felix uses<br />Default: Auto (other options: Legacy, NFT).
- `log_severity_sys` - (Computed) - String - LogSeveritySys the severity to set for logs which are sent to syslog<br />Default: INFO (other options: DEBUG, WARNING, ERROR, CRITICAL, NONE).
- `mtu` - (Computed) - Int - MTU to be set in the cni-network-config (default: 1500).
- `prometheus_go_metrics_enabled` - (Computed) - Bool - PrometheusGoMetricsEnabled enables Prometheus Go runtime metrics collection.
- `prometheus_metrics_enabled` - (Computed) - Bool - PrometheusMetricsEnabled can be set to enable the experimental Prometheus<br />metrics server (default: false).
- `prometheus_metrics_port` - (Computed) - Int - PrometheusMetricsPort is the TCP port that the experimental Prometheus<br />metrics server should bind to (default: 9091).
- `prometheus_process_metrics_enabled` - (Computed) - Bool - PrometheusProcessMetricsEnabled enables Prometheus process metrics collection.
- `typha_prometheus_metrics_enabled` - (Computed) - Bool - TyphaPrometheusMetricsEnabled enables Prometheus metrics collection from Typha<br />(default: false).
- `typha_prometheus_metrics_port` - (Computed) - Int - TyphaPrometheusMetricsPort is the TCP port the typha Prometheus metrics server<br />should bind to (default: 9093).
- `typha_replicas` - (Computed) - Int - TyphaReplicas is the number of replicas of Typha to deploy.

### kuberouter_networking_spec

KuberouterNetworkingSpec declares that we want Kube-router networking.


This resource has no attributes.

### romana_networking_spec

RomanaNetworkingSpec declares that we want Romana networking<br />Romana is deprecated as of kops 1.18 and removed as of kops 1.19.

#### Argument Reference

The following arguments are supported:

- `daemon_service_ip` - (Computed) - String - DaemonServiceIP is the Kubernetes Service IP for the romana-daemon pod.
- `etcd_service_ip` - (Computed) - String - EtcdServiceIP is the Kubernetes Service IP for the etcd backend used by Romana.

### amazon_vpc_networking_spec

AmazonVPCNetworkingSpec declares that we want Amazon VPC CNI networking.

#### Argument Reference

The following arguments are supported:

- `image_name` - (Computed) - String - The container image name to use.
- `env` - (Computed) - List([env_var](#env_var)) - Env is a list of environment variables to set in the container.

### cilium_networking_spec

CiliumNetworkingSpec declares that we want Cilium networking.

#### Argument Reference

The following arguments are supported:

- `version` - (Computed) - String - Version is the version of the Cilium agent and the Cilium Operator.
- `access_log` - (Computed) - String - AccessLog is not implemented and may be removed in the future.<br />Setting this has no effect.
- `agent_labels` - (Computed) - List(String) - AgentLabels is not implemented and may be removed in the future.<br />Setting this has no effect.
- `agent_prometheus_port` - (Computed) - Int - AgentPrometheusPort is the port to listen to for Prometheus metrics.<br />Defaults to 9090.
- `allow_localhost` - (Computed) - String - AllowLocalhost is not implemented and may be removed in the future.<br />Setting this has no effect.
- `auto_ipv6_node_routes` - (Computed) - Bool - AutoIpv6NodeRoutes is not implemented and may be removed in the future.<br />Setting this has no effect.
- `bpf_root` - (Computed) - String - BPFRoot is not implemented and may be removed in the future.<br />Setting this has no effect.
- `container_runtime` - (Computed) - List(String) - ContainerRuntime is not implemented and may be removed in the future.<br />Setting this has no effect.
- `container_runtime_endpoint` - (Computed) - Map(String) - ContainerRuntimeEndpoint is not implemented and may be removed in the future.<br />Setting this has no effect.
- `debug` - (Computed) - Bool - Debug runs Cilium in debug mode.
- `debug_verbose` - (Computed) - List(String) - DebugVerbose is not implemented and may be removed in the future.<br />Setting this has no effect.
- `device` - (Computed) - String - Device is not implemented and may be removed in the future.<br />Setting this has no effect.
- `disable_conntrack` - (Computed) - Bool - DisableConntrack is not implemented and may be removed in the future.<br />Setting this has no effect.
- `disable_ipv4` - (Computed) - Bool - DisableIpv4 is deprecated: Use EnableIpv4 instead.<br />Setting this flag has no effect.
- `disable_k_8s_services` - (Computed) - Bool - DisableK8sServices is not implemented and may be removed in the future.<br />Setting this has no effect.
- `enable_policy` - (Computed) - String - EnablePolicy specifies the policy enforcement mode.<br />"default": Follows Kubernetes policy enforcement.<br />"always": Cilium restricts all traffic if no policy is in place.<br />"never": Cilium allows all traffic regardless of policies in place.<br />If unspecified, "default" policy mode will be used.
- `enable_tracing` - (Computed) - Bool - EnableTracing is not implemented and may be removed in the future.<br />Setting this has no effect.
- `enable_prometheus_metrics` - (Computed) - Bool - EnablePrometheusMetrics enables the Cilium "/metrics" endpoint for both the agent and the operator.
- `enable_encryption` - (Computed) - Bool - EnableEncryption enables Cilium Encryption.<br />Default: false.
- `envoy_log` - (Computed) - String - EnvoyLog is not implemented and may be removed in the future.<br />Setting this has no effect.
- `ipv4_cluster_cidr_mask_size` - (Computed) - Int - Ipv4ClusterCIDRMaskSize is not implemented and may be removed in the future.<br />Setting this has no effect.
- `ipv4_node` - (Computed) - String - Ipv4Node is not implemented and may be removed in the future.<br />Setting this has no effect.
- `ipv4_range` - (Computed) - String - Ipv4Range is not implemented and may be removed in the future.<br />Setting this has no effect.
- `ipv4_service_range` - (Computed) - String - Ipv4ServiceRange is not implemented and may be removed in the future.<br />Setting this has no effect.
- `ipv6_cluster_alloc_cidr` - (Computed) - String - Ipv6ClusterAllocCidr is not implemented and may be removed in the future.<br />Setting this has no effect.
- `ipv6_node` - (Computed) - String - Ipv6Node is not implemented and may be removed in the future.<br />Setting this has no effect.
- `ipv6_range` - (Computed) - String - Ipv6Range is not implemented and may be removed in the future.<br />Setting this has no effect.
- `ipv6_service_range` - (Computed) - String - Ipv6ServiceRange is not implemented and may be removed in the future.<br />Setting this has no effect.
- `k_8s_api_server` - (Computed) - String - K8sAPIServer is not implemented and may be removed in the future.<br />Setting this has no effect.
- `k_8s_kubeconfig_path` - (Computed) - String - K8sKubeconfigPath is not implemented and may be removed in the future.<br />Setting this has no effect.
- `keep_bpf_templates` - (Computed) - Bool - KeepBPFTemplates is not implemented and may be removed in the future.<br />Setting this has no effect.
- `keep_config` - (Computed) - Bool - KeepConfig is not implemented and may be removed in the future.<br />Setting this has no effect.
- `label_prefix_file` - (Computed) - String - LabelPrefixFile is not implemented and may be removed in the future.<br />Setting this has currently no effect.
- `labels` - (Computed) - List(String) - Labels is not implemented and may be removed in the future.<br />Setting this has no effect.
- `lb` - (Computed) - String - LB is not implemented and may be removed in the future.<br />Setting this has no effect.
- `lib_dir` - (Computed) - String - LibDir is not implemented and may be removed in the future.<br />Setting this has no effect.
- `log_drivers` - (Computed) - List(String) - LogDrivers is not implemented and may be removed in the future.<br />Setting this has no effect.
- `log_opt` - (Computed) - Map(String) - LogOpt is not implemented and may be removed in the future.<br />Setting this has no effect.
- `logstash` - (Computed) - Bool - Logstash is not implemented and may be removed in the future.<br />Setting this has no effect.
- `logstash_agent` - (Computed) - String - LogstashAgent is not implemented and may be removed in the future.<br />Setting this has no effect.
- `logstash_probe_timer` - (Computed) - Int - LogstashProbeTimer is not implemented and may be removed in the future.<br />Setting this has no effect.
- `disable_masquerade` - (Computed) - Bool - DisableMasquerade disables masquerading traffic to external destinations behind the node IP.
- `nat46_range` - (Computed) - String - Nat6Range is not implemented and may be removed in the future.<br />Setting this has no effect.
- `pprof` - (Computed) - Bool - Pprof is not implemented and may be removed in the future.<br />Setting this has no effect.
- `prefilter_device` - (Computed) - String - PrefilterDevice is not implemented and may be removed in the future.<br />Setting this has no effect.
- `prometheus_serve_addr` - (Computed) - String - PrometheusServeAddr is deprecated. Use EnablePrometheusMetrics and AgentPrometheusPort instead.<br />Setting this has no effect.
- `restore` - (Computed) - Bool - Restore is not implemented and may be removed in the future.<br />Setting this has no effect.
- `single_cluster_route` - (Computed) - Bool - SingleClusterRoute is not implemented and may be removed in the future.<br />Setting this has no effect.
- `socket_path` - (Computed) - String - SocketPath is not implemented and may be removed in the future.<br />Setting this has no effect.
- `state_dir` - (Computed) - String - StateDir is not implemented and may be removed in the future.<br />Setting this has no effect.
- `trace_payload_len` - (Computed) - Int - TracePayloadLen is not implemented and may be removed in the future.<br />Setting this has no effect.
- `tunnel` - (Computed) - String - Tunnel specifies the Cilium tunnelling mode. Possible values are "vxlan", "geneve", or "disabled".<br />Default: vxlan.
- `enable_ipv6` - (Computed) - Bool - EnableIpv6 is not implemented and may be removed in the future.<br />Setting this has no effect.
- `enable_ipv4` - (Computed) - Bool - EnableIpv4 is not implemented and may be removed in the future.<br />Setting this has no effect.
- `monitor_aggregation` - (Computed) - String - MonitorAggregation sets the level of packet monitoring. Possible values are "low", "medium", or "maximum".<br />Default: medium.
- `bpfct_global_tcp_max` - (Computed) - Int - BPFCTGlobalTCPMax is the maximum number of entries in the TCP CT table.<br />Default: 524288.
- `bpfct_global_any_max` - (Computed) - Int - BPFCTGlobalAnyMax is the maximum number of entries in the non-TCP CT table.<br />Default: 262144.
- `preallocate_bpf_maps` - (Computed) - Bool - PreallocateBPFMaps reduces the per-packet latency at the expense of up-front memory allocation.<br />Default: true.
- `sidecar_istio_proxy_image` - (Computed) - String - SidecarIstioProxyImage is the regular expression matching compatible Istio sidecar istio-proxy<br />container image names.<br />Default: cilium/istio_proxy.
- `cluster_name` - (Computed) - String - ClusterName is the name of the cluster. It is only relevant when building a mesh of clusters.
- `to_fqdns_dns_reject_response_code` - (Computed) - String - ToFqdnsDNSRejectResponseCode sets the DNS response code for rejecting DNS requests.<br />Possible values are "nameError" or "refused".<br />Default: refused.
- `to_fqdns_enable_poller` - (Computed) - Bool - ToFqdnsEnablePoller replaces the DNS proxy-based implementation of FQDN policies<br />with the less powerful legacy implementation.<br />Default: false.
- `container_runtime_labels` - (Computed) - String - ContainerRuntimeLabels enables fetching of container-runtime labels from the specified container runtime and associating them with endpoints.<br />Supported values are: "none", "containerd", "crio", "docker", "auto"<br />As of Cilium 1.7.0, Cilium no longer fetches information from the<br />container runtime and this field is ignored.<br />Default: none.
- `ipam` - (Computed) - String - Ipam specifies the IP address allocation mode to use.<br />Possible values are "crd" and "eni".<br />"eni" will use AWS native networking for pods. Eni requires masquerade to be set to false.<br />"crd" will use CRDs for controlling IP address management.<br />"hostscope" will use hostscope IPAM mode.<br />"kubernetes" will use addersing based on node pod CIDR.<br />Empty value will use hostscope for cilum <= 1.7 and "kubernetes" otherwise.
- `ip_tables_rules_noinstall` - (Computed) - Bool - IPTablesRulesNoinstall disables installing the base IPTables rules used for masquerading and kube-proxy.<br />Default: false.
- `auto_direct_node_routes` - (Computed) - Bool - AutoDirectNodeRoutes adds automatic L2 routing between nodes.<br />Default: false.
- `enable_host_reachable_services` - (Computed) - Bool - EnableHostReachableServices configures Cilium to enable services to be<br />reached from the host namespace in addition to pod namespaces.<br />https://docs.cilium.io/en/v1.9/gettingstarted/host-services/<br />Default: false.
- `enable_node_port` - (Computed) - Bool - EnableNodePort replaces kube-proxy with Cilium's BPF implementation.<br />Requires spec.kubeProxy.enabled be set to false.<br />Default: false.
- `etcd_managed` - (Computed) - Bool - EtcdManagd installs an additional etcd cluster that is used for Cilium state change.<br />The cluster is operated by cilium-etcd-operator.<br />Default: false.
- `enable_remote_node_identity` - (Computed) - Bool - EnableRemoteNodeIdentity enables the remote-node-identity added in Cilium 1.7.0.<br />Default: true.
- `hubble` - (Computed) - [hubble_spec](#hubble_spec) - Hubble configures the Hubble service on the Cilium agent.
- `remove_cbr_bridge` - (Computed) - Bool - RemoveCbrBridge is not implemented and may be removed in the future.<br />Setting this has no effect.
- `restart_pods` - (Computed) - Bool - RestartPods is not implemented and may be removed in the future.<br />Setting this has no effect.
- `reconfigure_kubelet` - (Computed) - Bool - ReconfigureKubelet is not implemented and may be removed in the future.<br />Setting this has no effect.
- `node_init_bootstrap_file` - (Computed) - String - NodeInitBootstrapFile is not implemented and may be removed in the future.<br />Setting this has no effect.
- `cni_bin_path` - (Computed) - String - CniBinPath is not implemented and may be removed in the future.<br />Setting this has no effect.

### hubble_spec

HubbleSpec configures the Hubble service on the Cilium agent.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Computed) - Bool - Enabled decides if Hubble is enabled on the agent or not.
- `metrics` - (Computed) - List(String) - Metrics is a list of metrics to collect. If empty or null, metrics are disabled.<br />See https://docs.cilium.io/en/stable/configuration/metrics/#hubble-exported-metrics.

### lyft_vpc_networking_spec

LyftVPCNetworkingSpec declares that we want to use the cni-ipvlan-vpc-k8s CNI networking.

#### Argument Reference

The following arguments are supported:

- `subnet_tags` - (Computed) - Map(String)

### gce_networking_spec

GCENetworkingSpec is the specification of GCE's native networking mode, using IP aliases.


This resource has no attributes.

### access_spec

AccessSpec provides configuration details related to kubeapi dns and ELB access.

#### Argument Reference

The following arguments are supported:

- `dns` - (Computed) - [dns_access_spec](#dns_access_spec) - DNS will be used to provide config on kube-apiserver ELB DNS.
- `load_balancer` - (Computed) - [load_balancer_access_spec](#load_balancer_access_spec) - LoadBalancer is the configuration for the kube-apiserver ELB.

### dns_access_spec


This resource has no attributes.

### load_balancer_access_spec

LoadBalancerAccessSpec provides configuration details related to API LoadBalancer and its access.

#### Argument Reference

The following arguments are supported:

- `class` - (Computed) - String - LoadBalancerClass specifies the class of load balancer to create: Classic, Network.
- `type` - (Computed) - String - Type of load balancer to create may Public or Internal.
- `idle_timeout_seconds` - (Computed) - Int - IdleTimeoutSeconds sets the timeout of the api loadbalancer.
- `security_group_override` - (Computed) - String - SecurityGroupOverride overrides the default Kops created SG for the load balancer.
- `additional_security_groups` - (Computed) - List(String) - AdditionalSecurityGroups attaches additional security groups (e.g. sg-123456).
- `use_for_internal_api` - (Computed) - Bool - UseForInternalApi indicates whether the LB should be used by the kubelet.
- `ssl_certificate` - (Computed) - String - SSLCertificate allows you to specify the ACM cert to be used the LB.
- `ssl_policy` - (Computed) - String - SSLPolicy allows you to overwrite the LB listener's Security Policy.
- `cross_zone_load_balancing` - (Computed) - Bool - CrossZoneLoadBalancing allows you to enable the cross zone load balancing.
- `subnets` - (Computed) - List([load_balancer_subnet_spec](#load_balancer_subnet_spec)) - Subnets allows you to specify the subnets that must be used for the load balancer.

### load_balancer_subnet_spec

LoadBalancerSubnetSpec provides configuration for subnets used for a load balancer.

#### Argument Reference

The following arguments are supported:

- `name` - (Computed) - String - Name specifies the name of the cluster subnet.
- `private_ipv4_address` - (Computed) - String - PrivateIPv4Address specifies the private IPv4 address to use for a NLB.
- `allocation_id` - (Computed) - String - AllocationID specifies the Elastic IP Allocation ID for use by a NLB.

### authentication_spec

#### Argument Reference

The following arguments are supported:

- `kopeio` - (Computed) - [kopeio_authentication_spec](#kopeio_authentication_spec)
- `aws` - (Computed) - [aws_authentication_spec](#aws_authentication_spec)

### kopeio_authentication_spec


This resource has no attributes.

### aws_authentication_spec

#### Argument Reference

The following arguments are supported:

- `image` - (Computed) - String - Image is the AWS IAM Authenticator docker image to use.
- `backend_mode` - (Computed) - String - BackendMode is the AWS IAM Authenticator backend to use. Default MountedFile.
- `cluster_id` - (Computed) - String - ClusterID identifies the cluster performing authentication to prevent certain replay attacks. Default master public DNS name.
- `memory_request` - (Computed) - Quantity - MemoryRequest memory request of AWS IAM Authenticator container. Default 20Mi.
- `cpu_request` - (Computed) - Quantity - CPURequest CPU request of AWS IAM Authenticator container. Default 10m.
- `memory_limit` - (Computed) - Quantity - MemoryLimit memory limit of AWS IAM Authenticator container. Default 20Mi.
- `cpu_limit` - (Computed) - Quantity - CPULimit CPU limit of AWS IAM Authenticator container. Default 10m.

### authorization_spec

#### Argument Reference

The following arguments are supported:

- `always_allow` - (Computed) - [always_allow_authorization_spec](#always_allow_authorization_spec)
- `rbac` - (Computed) - [rbac_authorization_spec](#rbac_authorization_spec)

### always_allow_authorization_spec


This resource has no attributes.

### rbac_authorization_spec


This resource has no attributes.

### node_authorization_spec

NodeAuthorizationSpec is used to node authorization.

#### Argument Reference

The following arguments are supported:

- `node_authorizer` - (Computed) - [node_authorizer_spec](#node_authorizer_spec) - NodeAuthorizer defined the configuration for the node authorizer.

### node_authorizer_spec

NodeAuthorizerSpec defines the configuration for a node authorizer.

#### Argument Reference

The following arguments are supported:

- `authorizer` - (Computed) - String - Authorizer is the authorizer to use.
- `features` - (Computed) - List(String) - Features is a series of authorizer features to enable or disable.
- `image` - (Computed) - String - Image is the location of container.
- `node_url` - (Computed) - String - NodeURL is the node authorization service url.
- `port` - (Computed) - Int - Port is the port the service is running on the master.
- `interval` - (Computed) - Duration - Interval the time between retires for authorization request.
- `timeout` - (Computed) - Duration - Timeout the max time for authorization request.
- `token_ttl` - (Computed) - Duration - TokenTTL is the max ttl for an issued token.

### hook_spec

HookSpec is a definition hook.

#### Argument Reference

The following arguments are supported:

- `name` - (Computed) - String - Name is an optional name for the hook, otherwise the name is kops-hook-<index>.
- `disabled` - (Computed) - Bool - Disabled indicates if you want the unit switched off.
- `roles` - (Computed) - List(String) - Roles is an optional list of roles the hook should be rolled out to, defaults to all.
- `requires` - (Computed) - List(String) - Requires is a series of systemd units the action requires.
- `before` - (Computed) - List(String) - Before is a series of systemd units which this hook must run before.
- `exec_container` - (Computed) - [exec_container_action](#exec_container_action) - ExecContainer is the image itself.
- `manifest` - (Computed) - String - Manifest is a raw systemd unit file.
- `use_raw_manifest` - (Computed) - Bool - UseRawManifest indicates that the contents of Manifest should be used as the contents<br />of the systemd unit, unmodified. Before and Requires are ignored when used together<br />with this value (and validation shouldn't allow them to be set).

### exec_container_action

ExecContainerAction defines an hood action.

#### Argument Reference

The following arguments are supported:

- `image` - (Computed) - String - Image is the docker image.
- `command` - (Computed) - List(String) - Command is the command supplied to the above image.
- `environment` - (Computed) - Map(String) - Environment is a map of environment variables added to the hook.

### assets

Assets defines the privately hosted assets.

#### Argument Reference

The following arguments are supported:

- `container_registry` - (Computed) - String - ContainerRegistry is a url for to a docker registry.
- `file_repository` - (Computed) - String - FileRepository is the url for a private file serving repository.
- `container_proxy` - (Computed) - String - ContainerProxy is a url for a pull-through proxy of a docker registry.

### iam_spec

IAMSpec adds control over the IAM security policies applied to resources.

#### Argument Reference

The following arguments are supported:

- `legacy` - (Computed) - Bool - TODO: remove Legacy in next APIVersion.
- `allow_container_registry` - (Computed) - Bool
- `permissions_boundary` - (Computed) - String

### rolling_update

#### Argument Reference

The following arguments are supported:

- `drain_and_terminate` - (Computed) - Bool - DrainAndTerminate enables draining and terminating nodes during rolling updates.<br />Defaults to true.
- `max_unavailable` - (Computed) - IntOrString - MaxUnavailable is the maximum number of nodes that can be unavailable during the update.<br />The value can be an absolute number (for example 5) or a percentage of desired<br />nodes (for example 10%).<br />The absolute number is calculated from a percentage by rounding down.<br />Defaults to 1 if MaxSurge is 0, otherwise defaults to 0.<br />Example: when this is set to 30%, the InstanceGroup can be scaled<br />down to 70% of desired nodes immediately when the rolling update<br />starts. Once new nodes are ready, more old nodes can be drained,<br />ensuring that the total number of nodes available at all times<br />during the update is at least 70% of desired nodes.<br />+optional.
- `max_surge` - (Computed) - IntOrString - MaxSurge is the maximum number of extra nodes that can be created<br />during the update.<br />The value can be an absolute number (for example 5) or a percentage of<br />desired machines (for example 10%).<br />The absolute number is calculated from a percentage by rounding up.<br />Has no effect on instance groups with role "Master".<br />Defaults to 1 on AWS, 0 otherwise.<br />Example: when this is set to 30%, the InstanceGroup can be scaled<br />up immediately when the rolling update starts, such that the total<br />number of old and new nodes do not exceed 130% of desired<br />nodes.<br />+optional.

### cluster_autoscaler_config

ClusterAutoscalerConfig determines the cluster autoscaler configuration.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Computed) - Bool - Enabled enables the cluster autoscaler.<br />Default: false.
- `expander` - (Computed) - String - Expander determines the strategy for which instance group gets expanded.<br />Supported values: least-waste, most-pods, random.<br />Default: least-waste.
- `balance_similar_node_groups` - (Computed) - Bool - BalanceSimilarNodeGroups makes cluster autoscaler treat similar node groups as one.<br />Default: false.
- `scale_down_utilization_threshold` - (Computed) - String - ScaleDownUtilizationThreshold determines the utilization threshold for node scale-down.<br />Default: 0.5.
- `skip_nodes_with_system_pods` - (Computed) - Bool - SkipNodesWithSystemPods makes cluster autoscaler skip scale-down of nodes with non-DaemonSet pods in the kube-system namespace.<br />Default: true.
- `skip_nodes_with_local_storage` - (Computed) - Bool - SkipNodesWithLocalStorage makes cluster autoscaler skip scale-down of nodes with local storage.<br />Default: true.
- `new_pod_scale_up_delay` - (Computed) - String - NewPodScaleUpDelay causes cluster autoscaler to ignore unschedulable pods until they are a certain "age", regardless of the scan-interval<br />Default: 0s.
- `scale_down_delay_after_add` - (Computed) - String - ScaleDownDelayAfterAdd determines the time after scale up that scale down evaluation resumes<br />Default: 10m0s.
- `image` - (Computed) - String - Image is the docker container used.<br />Default: the latest supported image for the specified kubernetes version.
- `memory_request` - (Computed) - Quantity - MemoryRequest of cluster autoscaler container.<br />Default: 300Mi.
- `cpu_request` - (Computed) - Quantity - CPURequest of cluster autoscaler container.<br />Default: 100m.



