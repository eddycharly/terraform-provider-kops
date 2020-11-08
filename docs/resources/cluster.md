# kops_cluster

## Argument Reference

The following arguments are supported:
- `name` - (Required) - String
- `admin_ssh_key` - (Required) - String
- `channel` - (Optional) - String
- `addons` - (Optional) - List([addon_spec](#addon_spec))
- `config_base` - (Optional) - (Computed) - String
- `cloud_provider` - (Required) - String
- `container_runtime` - (Optional) - String
- `kubernetes_version` - (Optional) - String
- `subnet` - (Required) - List([cluster_subnet_spec](#cluster_subnet_spec))
- `project` - (Optional) - String
- `master_public_name` - (Optional) - (Computed) - String
- `master_internal_name` - (Optional) - (Computed) - String
- `network_cidr` - (Optional) - (Computed) - String
- `additional_network_cidrs` - (Optional) - List(String)
- `network_id` - (Required) - String
- `topology` - (Required) - [topology_spec](#topology_spec)
- `secret_store` - (Optional) - String
- `key_store` - (Optional) - String
- `config_store` - (Optional) - String
- `dns_zone` - (Optional) - String
- `additional_sans` - (Optional) - List(String)
- `cluster_dns_domain` - (Optional) - String
- `service_cluster_ip_range` - (Optional) - String
- `pod_cidr` - (Optional) - String
- `non_masquerade_cidr` - (Optional) - (Computed) - String
- `ssh_access` - (Optional) - List(String)
- `node_port_access` - (Optional) - List(String)
- `egress_proxy` - (Optional) - [egress_proxy_spec](#egress_proxy_spec)
- `ssh_key_name` - (Optional) - String
- `kubernetes_api_access` - (Optional) - List(String)
- `isolate_masters` - (Optional) - Bool
- `update_policy` - (Optional) - String
- `external_policies` - (Optional) - Map(List(String))
- `additional_policies` - (Optional) - Map(String)
- `file_assets` - (Optional) - List([file_asset_spec](#file_asset_spec))
- `etcd_cluster` - (Required) - List([etcd_cluster_spec](#etcd_cluster_spec))
- `containerd` - (Optional) - [containerd_config](#containerd_config)
- `docker` - (Optional) - [docker_config](#docker_config)
- `kube_dns` - (Optional) - [kube_dns_config](#kube_dns_config)
- `kube_api_server` - (Optional) - [kube_api_server_config](#kube_api_server_config)
- `kube_controller_manager` - (Optional) - [kube_controller_manager_config](#kube_controller_manager_config)
- `external_cloud_controller_manager` - (Optional) - [cloud_controller_manager_config](#cloud_controller_manager_config)
- `kube_scheduler` - (Optional) - [kube_scheduler_config](#kube_scheduler_config)
- `kube_proxy` - (Optional) - [kube_proxy_config](#kube_proxy_config)
- `kubelet` - (Optional) - [kubelet_config_spec](#kubelet_config_spec)
- `master_kubelet` - (Optional) - [kubelet_config_spec](#kubelet_config_spec)
- `cloud_config` - (Optional) - [cloud_configuration](#cloud_configuration)
- `external_dns` - (Optional) - [external_dns_config](#external_dns_config)
- `networking` - (Required) - [networking_spec](#networking_spec)
- `api` - (Optional) - [access_spec](#access_spec)
- `authentication` - (Optional) - [authentication_spec](#authentication_spec)
- `authorization` - (Optional) - [authorization_spec](#authorization_spec)
- `node_authorization` - (Optional) - [node_authorization_spec](#node_authorization_spec)
- `cloud_labels` - (Optional) - Map(String)
- `hooks` - (Optional) - List([hook_spec](#hook_spec))
- `assets` - (Optional) - [assets](#assets)
- `iam` - (Optional) - (Computed) - [iam_spec](#iam_spec)
- `encryption_config` - (Optional) - Bool
- `disable_subnet_tags` - (Optional) - Bool
- `use_host_certificates` - (Optional) - Bool
- `sysctl_parameters` - (Optional) - List(String)
- `rolling_update` - (Optional) - [rolling_update](#rolling_update)
- `instance_group` - (Required) - List([instance_group](#instance_group))
- `kube_config` - (Computed) - [kube_config](#kube_config)
- `rolling_update_options` - (Optional) - [rolling_update_options](#rolling_update_options)
- `validate_options` - (Optional) - [validate_options](#validate_options)

## Nested resources

### cluster

- `name` - (Required) - String
- `admin_ssh_key` - (Required) - String
- `channel` - (Optional) - String
- `addons` - (Optional) - List([addon_spec](#addon_spec))
- `config_base` - (Optional) - (Computed) - String
- `cloud_provider` - (Required) - String
- `container_runtime` - (Optional) - String
- `kubernetes_version` - (Optional) - String
- `subnet` - (Required) - List([cluster_subnet_spec](#cluster_subnet_spec))
- `project` - (Optional) - String
- `master_public_name` - (Optional) - (Computed) - String
- `master_internal_name` - (Optional) - (Computed) - String
- `network_cidr` - (Optional) - (Computed) - String
- `additional_network_cidrs` - (Optional) - List(String)
- `network_id` - (Required) - String
- `topology` - (Required) - [topology_spec](#topology_spec)
- `secret_store` - (Optional) - String
- `key_store` - (Optional) - String
- `config_store` - (Optional) - String
- `dns_zone` - (Optional) - String
- `additional_sans` - (Optional) - List(String)
- `cluster_dns_domain` - (Optional) - String
- `service_cluster_ip_range` - (Optional) - String
- `pod_cidr` - (Optional) - String
- `non_masquerade_cidr` - (Optional) - (Computed) - String
- `ssh_access` - (Optional) - List(String)
- `node_port_access` - (Optional) - List(String)
- `egress_proxy` - (Optional) - [egress_proxy_spec](#egress_proxy_spec)
- `ssh_key_name` - (Optional) - String
- `kubernetes_api_access` - (Optional) - List(String)
- `isolate_masters` - (Optional) - Bool
- `update_policy` - (Optional) - String
- `external_policies` - (Optional) - Map(List(String))
- `additional_policies` - (Optional) - Map(String)
- `file_assets` - (Optional) - List([file_asset_spec](#file_asset_spec))
- `etcd_cluster` - (Required) - List([etcd_cluster_spec](#etcd_cluster_spec))
- `containerd` - (Optional) - [containerd_config](#containerd_config)
- `docker` - (Optional) - [docker_config](#docker_config)
- `kube_dns` - (Optional) - [kube_dns_config](#kube_dns_config)
- `kube_api_server` - (Optional) - [kube_api_server_config](#kube_api_server_config)
- `kube_controller_manager` - (Optional) - [kube_controller_manager_config](#kube_controller_manager_config)
- `external_cloud_controller_manager` - (Optional) - [cloud_controller_manager_config](#cloud_controller_manager_config)
- `kube_scheduler` - (Optional) - [kube_scheduler_config](#kube_scheduler_config)
- `kube_proxy` - (Optional) - [kube_proxy_config](#kube_proxy_config)
- `kubelet` - (Optional) - [kubelet_config_spec](#kubelet_config_spec)
- `master_kubelet` - (Optional) - [kubelet_config_spec](#kubelet_config_spec)
- `cloud_config` - (Optional) - [cloud_configuration](#cloud_configuration)
- `external_dns` - (Optional) - [external_dns_config](#external_dns_config)
- `networking` - (Required) - [networking_spec](#networking_spec)
- `api` - (Optional) - [access_spec](#access_spec)
- `authentication` - (Optional) - [authentication_spec](#authentication_spec)
- `authorization` - (Optional) - [authorization_spec](#authorization_spec)
- `node_authorization` - (Optional) - [node_authorization_spec](#node_authorization_spec)
- `cloud_labels` - (Optional) - Map(String)
- `hooks` - (Optional) - List([hook_spec](#hook_spec))
- `assets` - (Optional) - [assets](#assets)
- `iam` - (Optional) - (Computed) - [iam_spec](#iam_spec)
- `encryption_config` - (Optional) - Bool
- `disable_subnet_tags` - (Optional) - Bool
- `use_host_certificates` - (Optional) - Bool
- `sysctl_parameters` - (Optional) - List(String)
- `rolling_update` - (Optional) - [rolling_update](#rolling_update)
- `instance_group` - (Required) - List([instance_group](#instance_group))
- `kube_config` - (Computed) - [kube_config](#kube_config)
- `rolling_update_options` - (Optional) - [rolling_update_options](#rolling_update_options)
- `validate_options` - (Optional) - [validate_options](#validate_options)

### addon_spec

- `manifest` - (Required) - String

### cluster_subnet_spec

- `name` - (Required) - String
- `cidr` - (Optional) - (Computed) - String
- `zone` - (Required) - String
- `region` - (Optional) - String
- `provider_id` - (Required) - String
- `egress` - (Optional) - String
- `type` - (Required) - String
- `public_ip` - (Optional) - String

### topology_spec

- `masters` - (Required) - String
- `nodes` - (Required) - String
- `bastion` - (Optional) - [bastion_spec](#bastion_spec)
- `dns` - (Required) - [dns_spec](#dns_spec)

### bastion_spec

- `bastion_public_name` - (Required) - String
- `idle_timeout_seconds` - (Optional) - Int
- `load_balancer` - (Optional) - [bastion_load_balancer_spec](#bastion_load_balancer_spec)

### bastion_load_balancer_spec

- `additional_security_groups` - (Required) - List(String)

### dns_spec

- `type` - (Required) - String

### egress_proxy_spec

- `http_proxy` - (Required) - [http_proxy](#http_proxy)
- `proxy_excludes` - (Optional) - String

### http_proxy

- `host` - (Required) - String
- `port` - (Required) - Int

### file_asset_spec

- `name` - (Required) - String
- `path` - (Required) - String
- `roles` - (Optional) - List(String)
- `content` - (Required) - String
- `is_base_64` - (Optional) - Bool

### etcd_cluster_spec

- `name` - (Required) - String
- `provider` - (Optional) - String
- `members` - (Required) - List([etcd_member_spec](#etcd_member_spec))
- `enable_etcd_tls` - (Optional) - Bool
- `enable_tls_auth` - (Optional) - Bool
- `version` - (Optional) - String
- `leader_election_timeout` - (Optional) - Duration
- `heartbeat_interval` - (Optional) - Duration
- `image` - (Optional) - String
- `backups` - (Optional) - [etcd_backup_spec](#etcd_backup_spec)
- `manager` - (Optional) - [etcd_manager_spec](#etcd_manager_spec)
- `memory_request` - (Optional) - Quantity
- `cpu_request` - (Optional) - Quantity

### etcd_member_spec

- `name` - (Required) - String
- `instance_group` - (Required) - String
- `volume_type` - (Optional) - String
- `volume_iops` - (Optional) - Int
- `volume_size` - (Optional) - Int
- `kms_key_id` - (Optional) - String
- `encrypted_volume` - (Optional) - Bool

### etcd_backup_spec

- `backup_store` - (Required) - String
- `image` - (Required) - String

### etcd_manager_spec

- `image` - (Required) - String
- `env` - (Optional) - List([env_var](#env_var))

### env_var

- `name` - (Required) - String
- `value` - (Optional) - String

### containerd_config

- `address` - (Optional) - String
- `config_override` - (Optional) - String
- `log_level` - (Optional) - String
- `root` - (Optional) - String
- `skip_install` - (Optional) - Bool
- `state` - (Optional) - String
- `version` - (Optional) - String

### docker_config

- `authorization_plugins` - (Optional) - List(String)
- `bridge` - (Optional) - String
- `bridge_ip` - (Optional) - String
- `data_root` - (Optional) - String
- `default_ulimit` - (Optional) - List(String)
- `exec_opt` - (Optional) - List(String)
- `exec_root` - (Optional) - String
- `experimental` - (Optional) - Bool
- `health_check` - (Optional) - Bool
- `hosts` - (Optional) - List(String)
- `ip_masq` - (Optional) - Bool
- `ip_tables` - (Optional) - Bool
- `insecure_registry` - (Optional) - String
- `insecure_registries` - (Optional) - List(String)
- `live_restore` - (Optional) - Bool
- `log_driver` - (Optional) - String
- `log_level` - (Optional) - String
- `log_opt` - (Optional) - List(String)
- `metrics_address` - (Optional) - String
- `mtu` - (Optional) - Int
- `registry_mirrors` - (Optional) - List(String)
- `skip_install` - (Optional) - Bool
- `storage` - (Optional) - String
- `storage_opts` - (Optional) - List(String)
- `user_namespace_remap` - (Optional) - String
- `version` - (Optional) - String

### kube_dns_config

- `cache_max_size` - (Optional) - Int
- `cache_max_concurrent` - (Optional) - Int
- `core_dns_image` - (Optional) - String
- `domain` - (Optional) - String
- `external_core_file` - (Optional) - String
- `image` - (Optional) - String
- `replicas` - (Optional) - Int
- `provider` - (Optional) - String
- `server_ip` - (Optional) - String
- `stub_domains` - (Optional) - Map(List(String))
- `upstream_nameservers` - (Optional) - List(String)
- `memory_request` - (Optional) - Quantity
- `cpu_request` - (Optional) - Quantity
- `memory_limit` - (Optional) - Quantity
- `node_local_dns` - (Optional) - [node_local_dns_config](#node_local_dns_config)

### node_local_dns_config

- `enabled` - (Optional) - Bool
- `local_ip` - (Optional) - String
- `memory_request` - (Optional) - Quantity
- `cpu_request` - (Optional) - Quantity

### kube_api_server_config

- `image` - (Optional) - String
- `disable_basic_auth` - (Optional) - Bool
- `log_level` - (Optional) - Int
- `cloud_provider` - (Optional) - String
- `secure_port` - (Optional) - Int
- `insecure_port` - (Optional) - Int
- `address` - (Optional) - String
- `bind_address` - (Optional) - String
- `insecure_bind_address` - (Optional) - String
- `enable_bootstrap_auth_token` - (Optional) - Bool
- `enable_aggregator_routing` - (Optional) - Bool
- `admission_control` - (Optional) - List(String)
- `append_admission_plugins` - (Optional) - List(String)
- `enable_admission_plugins` - (Optional) - List(String)
- `disable_admission_plugins` - (Optional) - List(String)
- `admission_control_config_file` - (Optional) - String
- `service_cluster_ip_range` - (Optional) - String
- `service_node_port_range` - (Optional) - String
- `etcd_servers` - (Optional) - List(String)
- `etcd_servers_overrides` - (Optional) - List(String)
- `etcd_ca_file` - (Optional) - String
- `etcd_cert_file` - (Optional) - String
- `etcd_key_file` - (Optional) - String
- `basic_auth_file` - (Optional) - String
- `client_ca_file` - (Optional) - String
- `tls_cert_file` - (Optional) - String
- `tls_private_key_file` - (Optional) - String
- `tls_cipher_suites` - (Optional) - List(String)
- `tls_min_version` - (Optional) - String
- `token_auth_file` - (Optional) - String
- `allow_privileged` - (Optional) - Bool
- `api_server_count` - (Optional) - Int
- `runtime_config` - (Optional) - Map(String)
- `kubelet_client_certificate` - (Optional) - String
- `kubelet_certificate_authority` - (Optional) - String
- `kubelet_client_key` - (Optional) - String
- `anonymous_auth` - (Optional) - Bool
- `kubelet_preferred_address_types` - (Optional) - List(String)
- `storage_backend` - (Optional) - String
- `oidc_username_claim` - (Optional) - String
- `oidc_username_prefix` - (Optional) - String
- `oidc_groups_claim` - (Optional) - String
- `oidc_groups_prefix` - (Optional) - String
- `oidc_issuer_url` - (Optional) - String
- `oidc_client_id` - (Optional) - String
- `oidc_required_claim` - (Optional) - List(String)
- `oidcca_file` - (Optional) - String
- `proxy_client_cert_file` - (Optional) - String
- `proxy_client_key_file` - (Optional) - String
- `audit_log_format` - (Optional) - String
- `audit_log_path` - (Optional) - String
- `audit_log_max_age` - (Optional) - Int
- `audit_log_max_backups` - (Optional) - Int
- `audit_log_max_size` - (Optional) - Int
- `audit_policy_file` - (Optional) - String
- `audit_webhook_batch_buffer_size` - (Optional) - Int
- `audit_webhook_batch_max_size` - (Optional) - Int
- `audit_webhook_batch_max_wait` - (Optional) - Duration
- `audit_webhook_batch_throttle_burst` - (Optional) - Int
- `audit_webhook_batch_throttle_enable` - (Optional) - Bool
- `audit_webhook_batch_throttle_qps` - (Optional) - Quantity
- `audit_webhook_config_file` - (Optional) - String
- `audit_webhook_initial_backoff` - (Optional) - Duration
- `audit_webhook_mode` - (Optional) - String
- `authentication_token_webhook_config_file` - (Optional) - String
- `authentication_token_webhook_cache_ttl` - (Optional) - Duration
- `authorization_mode` - (Optional) - String
- `authorization_webhook_config_file` - (Optional) - String
- `authorization_webhook_cache_authorized_ttl` - (Optional) - Duration
- `authorization_webhook_cache_unauthorized_ttl` - (Optional) - Duration
- `authorization_rbac_super_user` - (Optional) - String
- `encryption_provider_config` - (Optional) - String
- `experimental_encryption_provider_config` - (Optional) - String
- `requestheader_username_headers` - (Optional) - List(String)
- `requestheader_group_headers` - (Optional) - List(String)
- `requestheader_extra_header_prefixes` - (Optional) - List(String)
- `requestheader_client_ca_file` - (Optional) - String
- `requestheader_allowed_names` - (Optional) - List(String)
- `feature_gates` - (Optional) - Map(String)
- `max_requests_inflight` - (Optional) - Int
- `max_mutating_requests_inflight` - (Optional) - Int
- `http_2max_streams_per_connection` - (Optional) - Int
- `etcd_quorum_read` - (Optional) - Bool
- `request_timeout` - (Optional) - Duration
- `min_request_timeout` - (Optional) - Int
- `target_ram_mb` - (Optional) - Int
- `service_account_key_file` - (Optional) - List(String)
- `service_account_signing_key_file` - (Optional) - String
- `service_account_issuer` - (Optional) - String
- `service_account_jwksuri` - (Optional) - String
- `api_audiences` - (Optional) - List(String)
- `cpu_request` - (Optional) - String
- `event_ttl` - (Optional) - Duration
- `audit_dynamic_configuration` - (Optional) - Bool
- `enable_profiling` - (Optional) - Bool

### kube_controller_manager_config

- `master` - (Optional) - String
- `log_level` - (Optional) - Int
- `service_account_private_key_file` - (Optional) - String
- `image` - (Optional) - String
- `cloud_provider` - (Optional) - String
- `cluster_name` - (Optional) - String
- `cluster_cidr` - (Optional) - String
- `allocate_node_cidrs` - (Optional) - Bool
- `node_cidr_mask_size` - (Optional) - Int
- `configure_cloud_routes` - (Optional) - Bool
- `controllers` - (Optional) - List(String)
- `cidr_allocator_type` - (Optional) - String
- `root_ca_file` - (Optional) - String
- `leader_election` - (Optional) - [leader_election_configuration](#leader_election_configuration)
- `attach_detach_reconcile_sync_period` - (Optional) - Duration
- `disable_attach_detach_reconcile_sync` - (Optional) - Bool
- `terminated_pod_gc_threshold` - (Optional) - Int
- `node_monitor_period` - (Optional) - Duration
- `node_monitor_grace_period` - (Optional) - Duration
- `pod_eviction_timeout` - (Optional) - Duration
- `use_service_account_credentials` - (Optional) - Bool
- `horizontal_pod_autoscaler_sync_period` - (Optional) - Duration
- `horizontal_pod_autoscaler_downscale_delay` - (Optional) - Duration
- `horizontal_pod_autoscaler_downscale_stabilization` - (Optional) - Duration
- `horizontal_pod_autoscaler_upscale_delay` - (Optional) - Duration
- `horizontal_pod_autoscaler_tolerance` - (Optional) - Quantity
- `horizontal_pod_autoscaler_use_rest_clients` - (Optional) - Bool
- `experimental_cluster_signing_duration` - (Optional) - Duration
- `feature_gates` - (Optional) - Map(String)
- `tls_cipher_suites` - (Optional) - List(String)
- `tls_min_version` - (Optional) - String
- `min_resync_period` - (Optional) - String
- `kube_api_qps` - (Optional) - Quantity
- `kube_api_burst` - (Optional) - Int
- `concurrent_deployment_syncs` - (Optional) - Int
- `concurrent_endpoint_syncs` - (Optional) - Int
- `concurrent_namespace_syncs` - (Optional) - Int
- `concurrent_replicaset_syncs` - (Optional) - Int
- `concurrent_service_syncs` - (Optional) - Int
- `concurrent_resource_quota_syncs` - (Optional) - Int
- `concurrent_serviceaccount_token_syncs` - (Optional) - Int
- `concurrent_rc_syncs` - (Optional) - Int
- `enable_profiling` - (Optional) - Bool

### leader_election_configuration

- `leader_elect` - (Optional) - Bool
- `leader_elect_lease_duration` - (Optional) - Duration
- `leader_elect_renew_deadline_duration` - (Optional) - Duration
- `leader_elect_resource_lock` - (Optional) - String
- `leader_elect_resource_name` - (Optional) - String
- `leader_elect_resource_namespace` - (Optional) - String
- `leader_elect_retry_period` - (Optional) - Duration

### cloud_controller_manager_config

- `master` - (Optional) - String
- `log_level` - (Optional) - Int
- `image` - (Optional) - String
- `cloud_provider` - (Optional) - String
- `cluster_name` - (Optional) - String
- `cluster_cidr` - (Optional) - String
- `allocate_node_cidrs` - (Optional) - Bool
- `configure_cloud_routes` - (Optional) - Bool
- `cidr_allocator_type` - (Optional) - String
- `leader_election` - (Optional) - [leader_election_configuration](#leader_election_configuration)
- `use_service_account_credentials` - (Optional) - Bool

### kube_scheduler_config

- `master` - (Optional) - String
- `log_level` - (Optional) - Int
- `image` - (Optional) - String
- `leader_election` - (Optional) - [leader_election_configuration](#leader_election_configuration)
- `use_policy_config_map` - (Optional) - Bool
- `feature_gates` - (Optional) - Map(String)
- `max_persistent_volumes` - (Optional) - Int
- `qps` - (Optional) - Quantity
- `burst` - (Optional) - Int
- `enable_profiling` - (Optional) - Bool

### kube_proxy_config

- `image` - (Optional) - String
- `cpu_request` - (Optional) - String
- `cpu_limit` - (Optional) - String
- `memory_request` - (Optional) - String
- `memory_limit` - (Optional) - String
- `log_level` - (Optional) - Int
- `cluster_cidr` - (Optional) - String
- `hostname_override` - (Optional) - String
- `bind_address` - (Optional) - String
- `master` - (Optional) - String
- `metrics_bind_address` - (Optional) - String
- `enabled` - (Optional) - Bool
- `proxy_mode` - (Optional) - String
- `ip_vs_exclude_cidr_s` - (Optional) - List(String)
- `ip_vs_min_sync_period` - (Optional) - Duration
- `ip_vs_scheduler` - (Optional) - String
- `ip_vs_sync_period` - (Optional) - Duration
- `feature_gates` - (Optional) - Map(String)
- `conntrack_max_per_core` - (Optional) - Int
- `conntrack_min` - (Optional) - Int

### kubelet_config_spec

- `api_servers` - (Optional) - String
- `anonymous_auth` - (Optional) - Bool
- `authorization_mode` - (Optional) - String
- `bootstrap_kubeconfig` - (Optional) - String
- `client_ca_file` - (Optional) - String
- `tls_cert_file` - (Optional) - String
- `tls_private_key_file` - (Optional) - String
- `tls_cipher_suites` - (Optional) - List(String)
- `tls_min_version` - (Optional) - String
- `kubeconfig_path` - (Optional) - String
- `require_kubeconfig` - (Optional) - Bool
- `log_level` - (Optional) - Int
- `pod_manifest_path` - (Optional) - String
- `hostname_override` - (Optional) - String
- `pod_infra_container_image` - (Optional) - String
- `seccomp_profile_root` - (Optional) - String
- `allow_privileged` - (Optional) - Bool
- `enable_debugging_handlers` - (Optional) - Bool
- `register_node` - (Optional) - Bool
- `node_status_update_frequency` - (Optional) - Duration
- `cluster_domain` - (Optional) - String
- `cluster_dns` - (Optional) - String
- `network_plugin_name` - (Optional) - String
- `cloud_provider` - (Optional) - String
- `kubelet_cgroups` - (Optional) - String
- `runtime_cgroups` - (Optional) - String
- `read_only_port` - (Optional) - Int
- `system_cgroups` - (Optional) - String
- `cgroup_root` - (Optional) - String
- `configure_cbr_00` - (Optional) - Bool
- `hairpin_mode` - (Optional) - String
- `babysit_daemons` - (Optional) - Bool
- `max_pods` - (Optional) - Int
- `nvidia_gp_uss` - (Optional) - Int
- `pod_cidr` - (Optional) - String
- `resolver_config` - (Optional) - String
- `reconcile_cidr` - (Optional) - Bool
- `register_schedulable` - (Optional) - Bool
- `serialize_image_pulls` - (Optional) - Bool
- `node_labels` - (Optional) - Map(String)
- `non_masquerade_cidr` - (Optional) - String
- `enable_custom_metrics` - (Optional) - Bool
- `network_plugin_mtu` - (Optional) - Int
- `image_gc_high_threshold_percent` - (Optional) - Int
- `image_gc_low_threshold_percent` - (Optional) - Int
- `image_pull_progress_deadline` - (Optional) - Duration
- `eviction_hard` - (Optional) - String
- `eviction_soft` - (Optional) - String
- `eviction_soft_grace_period` - (Optional) - String
- `eviction_pressure_transition_period` - (Optional) - Duration
- `eviction_max_pod_grace_period` - (Optional) - Int
- `eviction_minimum_reclaim` - (Optional) - String
- `volume_plugin_directory` - (Optional) - String
- `taints` - (Optional) - List(String)
- `feature_gates` - (Optional) - Map(String)
- `kube_reserved` - (Optional) - Map(String)
- `kube_reserved_cgroup` - (Optional) - String
- `system_reserved` - (Optional) - Map(String)
- `system_reserved_cgroup` - (Optional) - String
- `enforce_node_allocatable` - (Optional) - String
- `runtime_request_timeout` - (Optional) - Duration
- `volume_stats_agg_period` - (Optional) - Duration
- `fail_swap_on` - (Optional) - Bool
- `experimental_allowed_unsafe_sysctls` - (Optional) - List(String)
- `allowed_unsafe_sysctls` - (Optional) - List(String)
- `streaming_connection_idle_timeout` - (Optional) - Duration
- `docker_disable_shared_pid` - (Optional) - Bool
- `root_dir` - (Optional) - String
- `authentication_token_webhook` - (Optional) - Bool
- `authentication_token_webhook_cache_ttl` - (Optional) - Duration
- `cpucfs_quota` - (Optional) - Bool
- `cpucfs_quota_period` - (Optional) - Duration
- `cpu_manager_policy` - (Optional) - String
- `registry_pull_qps` - (Optional) - Int
- `registry_burst` - (Optional) - Int
- `topology_manager_policy` - (Optional) - String
- `rotate_certificates` - (Optional) - Bool
- `protect_kernel_defaults` - (Optional) - Bool
- `cgroup_driver` - (Optional) - String

### cloud_configuration

- `multizone` - (Optional) - Bool
- `node_tags` - (Optional) - String
- `node_instance_prefix` - (Optional) - String
- `gce_service_account` - (Optional) - String
- `disable_security_group_ingress` - (Optional) - Bool
- `elb_security_group` - (Optional) - String
- `v_sphere_username` - (Optional) - String
- `v_sphere_password` - (Optional) - String
- `v_sphere_server` - (Optional) - String
- `v_sphere_datacenter` - (Optional) - String
- `v_sphere_resource_pool` - (Optional) - String
- `v_sphere_datastore` - (Optional) - String
- `v_sphere_core_dns_server` - (Optional) - String
- `spotinst_product` - (Optional) - String
- `spotinst_orientation` - (Optional) - String
- `openstack` - (Optional) - [openstack_configuration](#openstack_configuration)

### openstack_configuration

- `loadbalancer` - (Optional) - [openstack_loadbalancer_config](#openstack_loadbalancer_config)
- `monitor` - (Optional) - [openstack_monitor](#openstack_monitor)
- `router` - (Optional) - [openstack_router](#openstack_router)
- `block_storage` - (Optional) - [openstack_block_storage_config](#openstack_block_storage_config)
- `insecure_skip_verify` - (Optional) - Bool

### openstack_loadbalancer_config

- `method` - (Optional) - String
- `provider` - (Optional) - String
- `use_octavia` - (Optional) - Bool
- `floating_network` - (Optional) - String
- `floating_network_id` - (Optional) - String
- `floating_subnet` - (Optional) - String
- `subnet_id` - (Optional) - String
- `manage_sec_groups` - (Optional) - Bool

### openstack_monitor

- `delay` - (Optional) - String
- `timeout` - (Optional) - String
- `max_retries` - (Optional) - Int

### openstack_router

- `external_network` - (Optional) - String
- `dns_servers` - (Optional) - String
- `external_subnet` - (Optional) - String

### openstack_block_storage_config

- `version` - (Optional) - String
- `ignore_az` - (Optional) - Bool
- `override_az` - (Optional) - String

### external_dns_config

- `disable` - (Optional) - Bool
- `watch_ingress` - (Optional) - Bool
- `watch_namespace` - (Optional) - String

### networking_spec

- `classic` - (Optional) - [classic_networking_spec](#classic_networking_spec)
- `kubenet` - (Optional) - [kubenet_networking_spec](#kubenet_networking_spec)
- `external` - (Optional) - [external_networking_spec](#external_networking_spec)
- `cni` - (Optional) - [cni_networking_spec](#cni_networking_spec)
- `kopeio` - (Optional) - [kopeio_networking_spec](#kopeio_networking_spec)
- `weave` - (Optional) - [weave_networking_spec](#weave_networking_spec)
- `flannel` - (Optional) - [flannel_networking_spec](#flannel_networking_spec)
- `calico` - (Optional) - [calico_networking_spec](#calico_networking_spec)
- `canal` - (Optional) - [canal_networking_spec](#canal_networking_spec)
- `kuberouter` - (Optional) - [kuberouter_networking_spec](#kuberouter_networking_spec)
- `romana` - (Optional) - [romana_networking_spec](#romana_networking_spec)
- `amazon_vpc` - (Optional) - [amazon_vpc_networking_spec](#amazon_vpc_networking_spec)
- `cilium` - (Optional) - [cilium_networking_spec](#cilium_networking_spec)
- `lyft_vpc` - (Optional) - [lyft_vpc_networking_spec](#lyft_vpc_networking_spec)
- `gce` - (Optional) - [gce_networking_spec](#gce_networking_spec)

### classic_networking_spec

This resource has no attributes.

### kubenet_networking_spec

This resource has no attributes.

### external_networking_spec

This resource has no attributes.

### cni_networking_spec

- `uses_secondary_ip` - (Optional) - Bool

### kopeio_networking_spec

This resource has no attributes.

### weave_networking_spec

- `mtu` - (Optional) - Int
- `conn_limit` - (Optional) - Int
- `no_masq_local` - (Optional) - Int
- `memory_request` - (Optional) - Quantity
- `cpu_request` - (Optional) - Quantity
- `memory_limit` - (Optional) - Quantity
- `cpu_limit` - (Optional) - Quantity
- `net_extra_args` - (Optional) - String
- `npc_memory_request` - (Optional) - Quantity
- `npccpu_request` - (Optional) - Quantity
- `npc_memory_limit` - (Optional) - Quantity
- `npccpu_limit` - (Optional) - Quantity
- `npc_extra_args` - (Optional) - String

### flannel_networking_spec

- `backend` - (Optional) - String
- `disable_tx_checksum_offloading` - (Optional) - Bool
- `iptables_resync_seconds` - (Optional) - Int

### calico_networking_spec

- `cpu_request` - (Optional) - Quantity
- `cross_subnet` - (Optional) - Bool
- `log_severity_screen` - (Optional) - String
- `mtu` - (Optional) - Int
- `prometheus_metrics_enabled` - (Optional) - Bool
- `prometheus_metrics_port` - (Optional) - Int
- `prometheus_go_metrics_enabled` - (Optional) - Bool
- `prometheus_process_metrics_enabled` - (Optional) - Bool
- `major_version` - (Optional) - String
- `iptables_backend` - (Optional) - String
- `ip_ip_mode` - (Optional) - String
- `typha_prometheus_metrics_enabled` - (Optional) - Bool
- `typha_prometheus_metrics_port` - (Optional) - Int
- `typha_replicas` - (Optional) - Int
- `wireguard_enabled` - (Optional) - Bool

### canal_networking_spec

- `chain_insert_mode` - (Optional) - String
- `cpu_request` - (Optional) - Quantity
- `default_endpoint_to_host_action` - (Optional) - String
- `disable_flannel_forward_rules` - (Optional) - Bool
- `disable_tx_checksum_offloading` - (Optional) - Bool
- `iptables_backend` - (Optional) - String
- `log_severity_sys` - (Optional) - String
- `mtu` - (Optional) - Int
- `prometheus_go_metrics_enabled` - (Optional) - Bool
- `prometheus_metrics_enabled` - (Optional) - Bool
- `prometheus_metrics_port` - (Optional) - Int
- `prometheus_process_metrics_enabled` - (Optional) - Bool
- `typha_prometheus_metrics_enabled` - (Optional) - Bool
- `typha_prometheus_metrics_port` - (Optional) - Int
- `typha_replicas` - (Optional) - Int

### kuberouter_networking_spec

This resource has no attributes.

### romana_networking_spec

- `daemon_service_ip` - (Optional) - String
- `etcd_service_ip` - (Optional) - String

### amazon_vpc_networking_spec

- `image_name` - (Optional) - String
- `env` - (Optional) - List([env_var](#env_var))

### cilium_networking_spec

- `version` - (Optional) - String
- `access_log` - (Optional) - String
- `agent_labels` - (Optional) - List(String)
- `agent_prometheus_port` - (Optional) - Int
- `allow_localhost` - (Optional) - String
- `auto_ipv_6node_routes` - (Optional) - Bool
- `bpf_root` - (Optional) - String
- `container_runtime` - (Optional) - List(String)
- `container_runtime_endpoint` - (Optional) - Map(String)
- `debug` - (Optional) - Bool
- `debug_verbose` - (Optional) - List(String)
- `device` - (Optional) - String
- `disable_conntrack` - (Optional) - Bool
- `disable_ipv_4` - (Optional) - Bool
- `disable_k8s_services` - (Optional) - Bool
- `enable_policy` - (Optional) - String
- `enable_tracing` - (Optional) - Bool
- `enable_prometheus_metrics` - (Optional) - Bool
- `envoy_log` - (Optional) - String
- `ipv_4cluster_cidr_mask_size` - (Optional) - Int
- `ipv_4node` - (Optional) - String
- `ipv_4range` - (Optional) - String
- `ipv_4service_range` - (Optional) - String
- `ipv_6cluster_alloc_cidr` - (Optional) - String
- `ipv_6node` - (Optional) - String
- `ipv_6range` - (Optional) - String
- `ipv_6service_range` - (Optional) - String
- `k8s_api_server` - (Optional) - String
- `k8s_kubeconfig_path` - (Optional) - String
- `keep_bpf_templates` - (Optional) - Bool
- `keep_config` - (Optional) - Bool
- `label_prefix_file` - (Optional) - String
- `labels` - (Optional) - List(String)
- `lb` - (Optional) - String
- `lib_dir` - (Optional) - String
- `log_drivers` - (Optional) - List(String)
- `log_opt` - (Optional) - Map(String)
- `logstash` - (Optional) - Bool
- `logstash_agent` - (Optional) - String
- `logstash_probe_timer` - (Optional) - Int
- `disable_masquerade` - (Optional) - Bool
- `nat_46range` - (Optional) - String
- `pprof` - (Optional) - Bool
- `prefilter_device` - (Optional) - String
- `prometheus_serve_addr` - (Optional) - String
- `restore` - (Optional) - Bool
- `single_cluster_route` - (Optional) - Bool
- `socket_path` - (Optional) - String
- `state_dir` - (Optional) - String
- `trace_payload_len` - (Optional) - Int
- `tunnel` - (Optional) - String
- `enable_ipv_6` - (Optional) - Bool
- `enable_ipv_4` - (Optional) - Bool
- `monitor_aggregation` - (Optional) - String
- `bpfct_global_tcp_max` - (Optional) - Int
- `bpfct_global_any_max` - (Optional) - Int
- `preallocate_bpf_maps` - (Optional) - Bool
- `sidecar_istio_proxy_image` - (Optional) - String
- `cluster_name` - (Optional) - String
- `to_fqdns_dns_reject_response_code` - (Optional) - String
- `to_fqdns_enable_poller` - (Optional) - Bool
- `container_runtime_labels` - (Optional) - String
- `ipam` - (Optional) - String
- `ip_tables_rules_noinstall` - (Optional) - Bool
- `auto_direct_node_routes` - (Optional) - Bool
- `enable_node_port` - (Optional) - Bool
- `etcd_managed` - (Optional) - Bool
- `enable_remote_node_identity` - (Optional) - Bool
- `remove_cbr_bridge` - (Optional) - Bool
- `restart_pods` - (Optional) - Bool
- `reconfigure_kubelet` - (Optional) - Bool
- `node_init_bootstrap_file` - (Optional) - String
- `cni_bin_path` - (Optional) - String

### lyft_vpc_networking_spec

- `subnet_tags` - (Optional) - Map(String)

### gce_networking_spec

This resource has no attributes.

### access_spec

- `dns` - (Optional) - [dns_access_spec](#dns_access_spec)
- `load_balancer` - (Optional) - [load_balancer_access_spec](#load_balancer_access_spec)

### dns_access_spec

This resource has no attributes.

### load_balancer_access_spec

- `type` - (Required) - String
- `idle_timeout_seconds` - (Optional) - Int
- `security_group_override` - (Optional) - String
- `additional_security_groups` - (Optional) - List(String)
- `use_for_internal_api` - (Optional) - Bool
- `ssl_certificate` - (Optional) - String
- `cross_zone_load_balancing` - (Optional) - Bool

### authentication_spec

- `kopeio` - (Optional) - [kopeio_authentication_spec](#kopeio_authentication_spec)
- `aws` - (Optional) - [aws_authentication_spec](#aws_authentication_spec)

### kopeio_authentication_spec

This resource has no attributes.

### aws_authentication_spec

- `image` - (Optional) - String
- `memory_request` - (Optional) - Quantity
- `cpu_request` - (Optional) - Quantity
- `memory_limit` - (Optional) - Quantity
- `cpu_limit` - (Optional) - Quantity

### authorization_spec

- `always_allow` - (Optional) - [always_allow_authorization_spec](#always_allow_authorization_spec)
- `rbac` - (Optional) - [rbac_authorization_spec](#rbac_authorization_spec)

### always_allow_authorization_spec

This resource has no attributes.

### rbac_authorization_spec

This resource has no attributes.

### node_authorization_spec

- `node_authorizer` - (Optional) - [node_authorizer_spec](#node_authorizer_spec)

### node_authorizer_spec

- `authorizer` - (Optional) - String
- `features` - (Optional) - List(String)
- `image` - (Optional) - String
- `node_url` - (Optional) - String
- `port` - (Optional) - Int
- `interval` - (Optional) - Duration
- `timeout` - (Optional) - Duration
- `token_ttl` - (Optional) - Duration

### hook_spec

- `name` - (Required) - String
- `disabled` - (Optional) - Bool
- `roles` - (Optional) - List(String)
- `requires` - (Optional) - List(String)
- `before` - (Optional) - List(String)
- `exec_container` - (Optional) - [exec_container_action](#exec_container_action)
- `manifest` - (Optional) - String
- `use_raw_manifest` - (Optional) - Bool

### exec_container_action

- `image` - (Required) - String
- `command` - (Optional) - List(String)
- `environment` - (Optional) - Map(String)

### assets

- `container_registry` - (Optional) - String
- `file_repository` - (Optional) - String
- `container_proxy` - (Optional) - String

### iam_spec

- `legacy` - (Optional) - Bool
- `allow_container_registry` - (Optional) - Bool

### rolling_update

- `drain_and_terminate` - (Optional) - Bool
- `max_unavailable` - (Optional) - IntOrString
- `max_surge` - (Optional) - IntOrString

### instance_group

- `name` - (Required) - String
- `role` - (Required) - String
- `image` - (Optional) - (Computed) - String
- `min_size` - (Required) - Int
- `max_size` - (Required) - Int
- `machine_type` - (Required) - String
- `root_volume_size` - (Optional) - Int
- `root_volume_type` - (Optional) - String
- `root_volume_iops` - (Optional) - Int
- `root_volume_optimization` - (Optional) - Bool
- `root_volume_delete_on_termination` - (Optional) - Bool
- `root_volume_encryption` - (Optional) - Bool
- `volumes` - (Optional) - List([volume_spec](#volume_spec))
- `volume_mounts` - (Optional) - List([volume_mount_spec](#volume_mount_spec))
- `subnets` - (Required) - List(String)
- `zones` - (Optional) - List(String)
- `hooks` - (Optional) - List([hook_spec](#hook_spec))
- `max_price` - (Optional) - String
- `spot_duration_in_minutes` - (Optional) - Int
- `associate_public_ip` - (Optional) - Bool
- `additional_security_groups` - (Optional) - List(String)
- `cloud_labels` - (Optional) - Map(String)
- `node_labels` - (Optional) - Map(String)
- `file_assets` - (Optional) - List([file_asset_spec](#file_asset_spec))
- `tenancy` - (Optional) - String
- `kubelet` - (Optional) - [kubelet_config_spec](#kubelet_config_spec)
- `taints` - (Optional) - List(String)
- `mixed_instances_policy` - (Optional) - [mixed_instances_policy_spec](#mixed_instances_policy_spec)
- `additional_user_data` - (Optional) - List([user_data](#user_data))
- `suspend_processes` - (Optional) - List(String)
- `external_load_balancers` - (Optional) - List([load_balancer](#load_balancer))
- `detailed_instance_monitoring` - (Optional) - Bool
- `iam` - (Optional) - [iam_profile_spec](#iam_profile_spec)
- `security_group_override` - (Optional) - String
- `instance_protection` - (Optional) - Bool
- `sysctl_parameters` - (Optional) - List(String)
- `rolling_update` - (Optional) - [rolling_update](#rolling_update)
- `instance_interruption_behavior` - (Optional) - String

### volume_spec

- `delete_on_termination` - (Optional) - Bool
- `device` - (Required) - String
- `encrypted` - (Optional) - Bool
- `iops` - (Optional) - Int
- `size` - (Optional) - Int
- `type` - (Optional) - String

### volume_mount_spec

- `device` - (Required) - String
- `filesystem` - (Required) - String
- `format_options` - (Optional) - List(String)
- `mount_options` - (Optional) - List(String)
- `path` - (Required) - String

### mixed_instances_policy_spec

- `instances` - (Optional) - List(String)
- `on_demand_allocation_strategy` - (Optional) - String
- `on_demand_base` - (Optional) - Int
- `on_demand_above_base` - (Optional) - Int
- `spot_allocation_strategy` - (Optional) - String
- `spot_instance_pools` - (Optional) - Int

### user_data

- `name` - (Required) - String
- `type` - (Required) - String
- `content` - (Required) - String

### load_balancer

- `load_balancer_name` - (Optional) - String
- `target_group_arn` - (Optional) - String

### iam_profile_spec

- `profile` - (Required) - String

### kube_config

- `server` - (Computed) - String
- `context` - (Computed) - String
- `namespace` - (Computed) - String
- `kube_bearer_token` - (Computed) - String
- `kube_user` - (Computed) - String
- `kube_password` - (Computed) - String
- `ca_cert` - (Computed) - String
- `client_cert` - (Computed) - String
- `client_key` - (Computed) - String

### rolling_update_options

- `skip` - (Optional) - Bool
- `master_interval` - (Optional) - Duration
- `node_interval` - (Optional) - Duration
- `bastion_interval` - (Optional) - Duration
- `fail_on_drain_error` - (Optional) - Bool
- `fail_on_validate` - (Optional) - Bool
- `post_drain_delay` - (Optional) - Duration
- `validation_timeout` - (Optional) - Duration
- `validate_count` - (Optional) - Int

### validate_options

- `skip` - (Optional) - Bool
- `timeout` - (Optional) - Duration
- `poll_interval` - (Optional) - Duration

