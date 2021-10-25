package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceKubeAPIServerConfig(t *testing.T) {
	_default := kops.KubeAPIServerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KubeAPIServerConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"image":                                        "",
					"disable_basic_auth":                           nil,
					"log_format":                                   "",
					"log_level":                                    0,
					"cloud_provider":                               "",
					"secure_port":                                  0,
					"insecure_port":                                0,
					"address":                                      "",
					"bind_address":                                 "",
					"insecure_bind_address":                        "",
					"enable_bootstrap_auth_token":                  nil,
					"enable_aggregator_routing":                    nil,
					"admission_control":                            func() []interface{} { return nil }(),
					"append_admission_plugins":                     func() []interface{} { return nil }(),
					"enable_admission_plugins":                     func() []interface{} { return nil }(),
					"disable_admission_plugins":                    func() []interface{} { return nil }(),
					"admission_control_config_file":                "",
					"service_cluster_ip_range":                     "",
					"service_node_port_range":                      "",
					"etcd_servers":                                 func() []interface{} { return nil }(),
					"etcd_servers_overrides":                       func() []interface{} { return nil }(),
					"etcd_ca_file":                                 "",
					"etcd_cert_file":                               "",
					"etcd_key_file":                                "",
					"basic_auth_file":                              "",
					"client_ca_file":                               "",
					"tls_cert_file":                                "",
					"tls_private_key_file":                         "",
					"tls_cipher_suites":                            func() []interface{} { return nil }(),
					"tls_min_version":                              "",
					"token_auth_file":                              "",
					"allow_privileged":                             nil,
					"api_server_count":                             nil,
					"runtime_config":                               func() map[string]interface{} { return nil }(),
					"kubelet_client_certificate":                   "",
					"kubelet_certificate_authority":                "",
					"kubelet_client_key":                           "",
					"anonymous_auth":                               nil,
					"kubelet_preferred_address_types":              func() []interface{} { return nil }(),
					"storage_backend":                              nil,
					"oidc_username_claim":                          nil,
					"oidc_username_prefix":                         nil,
					"oidc_groups_claim":                            nil,
					"oidc_groups_prefix":                           nil,
					"oidc_issuer_url":                              nil,
					"oidc_client_id":                               nil,
					"oidc_required_claim":                          func() []interface{} { return nil }(),
					"oidcca_file":                                  nil,
					"proxy_client_cert_file":                       nil,
					"proxy_client_key_file":                        nil,
					"audit_log_format":                             nil,
					"audit_log_path":                               nil,
					"audit_log_max_age":                            nil,
					"audit_log_max_backups":                        nil,
					"audit_log_max_size":                           nil,
					"audit_policy_file":                            "",
					"audit_webhook_batch_buffer_size":              nil,
					"audit_webhook_batch_max_size":                 nil,
					"audit_webhook_batch_max_wait":                 nil,
					"audit_webhook_batch_throttle_burst":           nil,
					"audit_webhook_batch_throttle_enable":          nil,
					"audit_webhook_batch_throttle_qps":             nil,
					"audit_webhook_config_file":                    "",
					"audit_webhook_initial_backoff":                nil,
					"audit_webhook_mode":                           "",
					"authentication_token_webhook_config_file":     nil,
					"authentication_token_webhook_cache_ttl":       nil,
					"authorization_mode":                           nil,
					"authorization_webhook_config_file":            nil,
					"authorization_webhook_cache_authorized_ttl":   nil,
					"authorization_webhook_cache_unauthorized_ttl": nil,
					"authorization_rbac_super_user":                nil,
					"encryption_provider_config":                   nil,
					"experimental_encryption_provider_config":      nil,
					"requestheader_username_headers":               func() []interface{} { return nil }(),
					"requestheader_group_headers":                  func() []interface{} { return nil }(),
					"requestheader_extra_header_prefixes":          func() []interface{} { return nil }(),
					"requestheader_client_ca_file":                 "",
					"requestheader_allowed_names":                  func() []interface{} { return nil }(),
					"feature_gates":                                func() map[string]interface{} { return nil }(),
					"max_requests_inflight":                        0,
					"max_mutating_requests_inflight":               0,
					"http2_max_streams_per_connection":             nil,
					"etcd_quorum_read":                             nil,
					"request_timeout":                              nil,
					"min_request_timeout":                          nil,
					"target_ram_mb":                                0,
					"service_account_key_file":                     func() []interface{} { return nil }(),
					"service_account_signing_key_file":             nil,
					"service_account_issuer":                       nil,
					"service_account_jwksuri":                      nil,
					"api_audiences":                                func() []interface{} { return nil }(),
					"cpu_request":                                  "",
					"cpu_limit":                                    "",
					"memory_request":                               "",
					"memory_limit":                                 "",
					"event_ttl":                                    nil,
					"audit_dynamic_configuration":                  nil,
					"enable_profiling":                             nil,
					"cors_allowed_origins":                         func() []interface{} { return nil }(),
					"default_not_ready_toleration_seconds":         nil,
					"default_unreachable_toleration_seconds":       nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceKubeAPIServerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceKubeAPIServerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceKubeAPIServerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"image":                                        "",
		"disable_basic_auth":                           nil,
		"log_format":                                   "",
		"log_level":                                    0,
		"cloud_provider":                               "",
		"secure_port":                                  0,
		"insecure_port":                                0,
		"address":                                      "",
		"bind_address":                                 "",
		"insecure_bind_address":                        "",
		"enable_bootstrap_auth_token":                  nil,
		"enable_aggregator_routing":                    nil,
		"admission_control":                            func() []interface{} { return nil }(),
		"append_admission_plugins":                     func() []interface{} { return nil }(),
		"enable_admission_plugins":                     func() []interface{} { return nil }(),
		"disable_admission_plugins":                    func() []interface{} { return nil }(),
		"admission_control_config_file":                "",
		"service_cluster_ip_range":                     "",
		"service_node_port_range":                      "",
		"etcd_servers":                                 func() []interface{} { return nil }(),
		"etcd_servers_overrides":                       func() []interface{} { return nil }(),
		"etcd_ca_file":                                 "",
		"etcd_cert_file":                               "",
		"etcd_key_file":                                "",
		"basic_auth_file":                              "",
		"client_ca_file":                               "",
		"tls_cert_file":                                "",
		"tls_private_key_file":                         "",
		"tls_cipher_suites":                            func() []interface{} { return nil }(),
		"tls_min_version":                              "",
		"token_auth_file":                              "",
		"allow_privileged":                             nil,
		"api_server_count":                             nil,
		"runtime_config":                               func() map[string]interface{} { return nil }(),
		"kubelet_client_certificate":                   "",
		"kubelet_certificate_authority":                "",
		"kubelet_client_key":                           "",
		"anonymous_auth":                               nil,
		"kubelet_preferred_address_types":              func() []interface{} { return nil }(),
		"storage_backend":                              nil,
		"oidc_username_claim":                          nil,
		"oidc_username_prefix":                         nil,
		"oidc_groups_claim":                            nil,
		"oidc_groups_prefix":                           nil,
		"oidc_issuer_url":                              nil,
		"oidc_client_id":                               nil,
		"oidc_required_claim":                          func() []interface{} { return nil }(),
		"oidcca_file":                                  nil,
		"proxy_client_cert_file":                       nil,
		"proxy_client_key_file":                        nil,
		"audit_log_format":                             nil,
		"audit_log_path":                               nil,
		"audit_log_max_age":                            nil,
		"audit_log_max_backups":                        nil,
		"audit_log_max_size":                           nil,
		"audit_policy_file":                            "",
		"audit_webhook_batch_buffer_size":              nil,
		"audit_webhook_batch_max_size":                 nil,
		"audit_webhook_batch_max_wait":                 nil,
		"audit_webhook_batch_throttle_burst":           nil,
		"audit_webhook_batch_throttle_enable":          nil,
		"audit_webhook_batch_throttle_qps":             nil,
		"audit_webhook_config_file":                    "",
		"audit_webhook_initial_backoff":                nil,
		"audit_webhook_mode":                           "",
		"authentication_token_webhook_config_file":     nil,
		"authentication_token_webhook_cache_ttl":       nil,
		"authorization_mode":                           nil,
		"authorization_webhook_config_file":            nil,
		"authorization_webhook_cache_authorized_ttl":   nil,
		"authorization_webhook_cache_unauthorized_ttl": nil,
		"authorization_rbac_super_user":                nil,
		"encryption_provider_config":                   nil,
		"experimental_encryption_provider_config":      nil,
		"requestheader_username_headers":               func() []interface{} { return nil }(),
		"requestheader_group_headers":                  func() []interface{} { return nil }(),
		"requestheader_extra_header_prefixes":          func() []interface{} { return nil }(),
		"requestheader_client_ca_file":                 "",
		"requestheader_allowed_names":                  func() []interface{} { return nil }(),
		"feature_gates":                                func() map[string]interface{} { return nil }(),
		"max_requests_inflight":                        0,
		"max_mutating_requests_inflight":               0,
		"http2_max_streams_per_connection":             nil,
		"etcd_quorum_read":                             nil,
		"request_timeout":                              nil,
		"min_request_timeout":                          nil,
		"target_ram_mb":                                0,
		"service_account_key_file":                     func() []interface{} { return nil }(),
		"service_account_signing_key_file":             nil,
		"service_account_issuer":                       nil,
		"service_account_jwksuri":                      nil,
		"api_audiences":                                func() []interface{} { return nil }(),
		"cpu_request":                                  "",
		"cpu_limit":                                    "",
		"memory_request":                               "",
		"memory_limit":                                 "",
		"event_ttl":                                    nil,
		"audit_dynamic_configuration":                  nil,
		"enable_profiling":                             nil,
		"cors_allowed_origins":                         func() []interface{} { return nil }(),
		"default_not_ready_toleration_seconds":         nil,
		"default_unreachable_toleration_seconds":       nil,
	}
	type args struct {
		in kops.KubeAPIServerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KubeAPIServerConfig{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableBasicAuth - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.DisableBasicAuth = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogFormat - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.LogFormat = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.LogLevel = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudProvider - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.CloudProvider = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecurePort - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.SecurePort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InsecurePort - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.InsecurePort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Address - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.Address = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BindAddress - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.BindAddress = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InsecureBindAddress - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.InsecureBindAddress = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableBootstrapAuthToken - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EnableBootstrapAuthToken = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableAggregatorRouting - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EnableAggregatorRouting = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdmissionControl - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AdmissionControl = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AppendAdmissionPlugins - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AppendAdmissionPlugins = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableAdmissionPlugins - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EnableAdmissionPlugins = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableAdmissionPlugins - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.DisableAdmissionPlugins = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdmissionControlConfigFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AdmissionControlConfigFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceClusterIpRange - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ServiceClusterIPRange = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceNodePortRange - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ServiceNodePortRange = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdServers - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EtcdServers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdServersOverrides - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EtcdServersOverrides = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdCAFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EtcdCAFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdCertFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EtcdCertFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdKeyFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EtcdKeyFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BasicAuthFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.BasicAuthFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClientCAFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ClientCAFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSCertFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.TLSCertFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSPrivateKeyFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.TLSPrivateKeyFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSCipherSuites - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.TLSCipherSuites = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSMinVersion - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.TLSMinVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TokenAuthFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.TokenAuthFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowPrivileged - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AllowPrivileged = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ApiServerCount - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.APIServerCount = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RuntimeConfig - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.RuntimeConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeletClientCertificate - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.KubeletClientCertificate = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeletCertificateAuthority - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.KubeletCertificateAuthority = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeletClientKey - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.KubeletClientKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AnonymousAuth - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AnonymousAuth = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeletPreferredAddressTypes - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.KubeletPreferredAddressTypes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "StorageBackend - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.StorageBackend = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OIDCUsernameClaim - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.OIDCUsernameClaim = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OIDCUsernamePrefix - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.OIDCUsernamePrefix = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OIDCGroupsClaim - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.OIDCGroupsClaim = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OIDCGroupsPrefix - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.OIDCGroupsPrefix = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OIDCIssuerURL - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.OIDCIssuerURL = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OIDCClientID - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.OIDCClientID = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OIDCRequiredClaim - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.OIDCRequiredClaim = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OIDCCAFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.OIDCCAFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ProxyClientCertFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ProxyClientCertFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ProxyClientKeyFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ProxyClientKeyFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditLogFormat - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditLogFormat = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditLogPath - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditLogPath = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditLogMaxAge - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditLogMaxAge = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditLogMaxBackups - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditLogMaxBackups = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditLogMaxSize - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditLogMaxSize = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditPolicyFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditPolicyFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookBatchBufferSize - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookBatchBufferSize = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookBatchMaxSize - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookBatchMaxSize = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookBatchMaxWait - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookBatchMaxWait = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookBatchThrottleBurst - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookBatchThrottleBurst = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookBatchThrottleEnable - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookBatchThrottleEnable = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookBatchThrottleQps - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookBatchThrottleQps = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookConfigFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookConfigFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookInitialBackoff - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookInitialBackoff = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookMode - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthenticationTokenWebhookConfigFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuthenticationTokenWebhookConfigFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthenticationTokenWebhookCacheTTL - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuthenticationTokenWebhookCacheTTL = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationMode - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuthorizationMode = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationWebhookConfigFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuthorizationWebhookConfigFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationWebhookCacheAuthorizedTTL - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuthorizationWebhookCacheAuthorizedTTL = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationWebhookCacheUnauthorizedTTL - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuthorizationWebhookCacheUnauthorizedTTL = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationRBACSuperUser - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuthorizationRBACSuperUser = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EncryptionProviderConfig - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EncryptionProviderConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExperimentalEncryptionProviderConfig - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ExperimentalEncryptionProviderConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RequestheaderUsernameHeaders - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.RequestheaderUsernameHeaders = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RequestheaderGroupHeaders - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.RequestheaderGroupHeaders = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RequestheaderExtraHeaderPrefixes - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.RequestheaderExtraHeaderPrefixes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RequestheaderClientCAFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.RequestheaderClientCAFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RequestheaderAllowedNames - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.RequestheaderAllowedNames = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FeatureGates - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.FeatureGates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxRequestsInflight - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.MaxRequestsInflight = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxMutatingRequestsInflight - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.MaxMutatingRequestsInflight = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HTTP2MaxStreamsPerConnection - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.HTTP2MaxStreamsPerConnection = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdQuorumRead - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EtcdQuorumRead = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RequestTimeout - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.RequestTimeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MinRequestTimeout - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.MinRequestTimeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TargetRamMb - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.TargetRamMb = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountKeyFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ServiceAccountKeyFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountSigningKeyFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ServiceAccountSigningKeyFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountIssuer - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ServiceAccountIssuer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountJWKSURI - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ServiceAccountJWKSURI = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ApiAudiences - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.APIAudiences = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.CPURequest = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuLimit - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.CPULimit = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.MemoryRequest = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.MemoryLimit = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EventTTL - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EventTTL = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditDynamicConfiguration - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditDynamicConfiguration = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableProfiling - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EnableProfiling = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CorsAllowedOrigins - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.CorsAllowedOrigins = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DefaultNotReadyTolerationSeconds - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.DefaultNotReadyTolerationSeconds = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DefaultUnreachableTolerationSeconds - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.DefaultUnreachableTolerationSeconds = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceKubeAPIServerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceKubeAPIServerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceKubeAPIServerConfig(t *testing.T) {
	_default := map[string]interface{}{
		"image":                                        "",
		"disable_basic_auth":                           nil,
		"log_format":                                   "",
		"log_level":                                    0,
		"cloud_provider":                               "",
		"secure_port":                                  0,
		"insecure_port":                                0,
		"address":                                      "",
		"bind_address":                                 "",
		"insecure_bind_address":                        "",
		"enable_bootstrap_auth_token":                  nil,
		"enable_aggregator_routing":                    nil,
		"admission_control":                            func() []interface{} { return nil }(),
		"append_admission_plugins":                     func() []interface{} { return nil }(),
		"enable_admission_plugins":                     func() []interface{} { return nil }(),
		"disable_admission_plugins":                    func() []interface{} { return nil }(),
		"admission_control_config_file":                "",
		"service_cluster_ip_range":                     "",
		"service_node_port_range":                      "",
		"etcd_servers":                                 func() []interface{} { return nil }(),
		"etcd_servers_overrides":                       func() []interface{} { return nil }(),
		"etcd_ca_file":                                 "",
		"etcd_cert_file":                               "",
		"etcd_key_file":                                "",
		"basic_auth_file":                              "",
		"client_ca_file":                               "",
		"tls_cert_file":                                "",
		"tls_private_key_file":                         "",
		"tls_cipher_suites":                            func() []interface{} { return nil }(),
		"tls_min_version":                              "",
		"token_auth_file":                              "",
		"allow_privileged":                             nil,
		"api_server_count":                             nil,
		"runtime_config":                               func() map[string]interface{} { return nil }(),
		"kubelet_client_certificate":                   "",
		"kubelet_certificate_authority":                "",
		"kubelet_client_key":                           "",
		"anonymous_auth":                               nil,
		"kubelet_preferred_address_types":              func() []interface{} { return nil }(),
		"storage_backend":                              nil,
		"oidc_username_claim":                          nil,
		"oidc_username_prefix":                         nil,
		"oidc_groups_claim":                            nil,
		"oidc_groups_prefix":                           nil,
		"oidc_issuer_url":                              nil,
		"oidc_client_id":                               nil,
		"oidc_required_claim":                          func() []interface{} { return nil }(),
		"oidcca_file":                                  nil,
		"proxy_client_cert_file":                       nil,
		"proxy_client_key_file":                        nil,
		"audit_log_format":                             nil,
		"audit_log_path":                               nil,
		"audit_log_max_age":                            nil,
		"audit_log_max_backups":                        nil,
		"audit_log_max_size":                           nil,
		"audit_policy_file":                            "",
		"audit_webhook_batch_buffer_size":              nil,
		"audit_webhook_batch_max_size":                 nil,
		"audit_webhook_batch_max_wait":                 nil,
		"audit_webhook_batch_throttle_burst":           nil,
		"audit_webhook_batch_throttle_enable":          nil,
		"audit_webhook_batch_throttle_qps":             nil,
		"audit_webhook_config_file":                    "",
		"audit_webhook_initial_backoff":                nil,
		"audit_webhook_mode":                           "",
		"authentication_token_webhook_config_file":     nil,
		"authentication_token_webhook_cache_ttl":       nil,
		"authorization_mode":                           nil,
		"authorization_webhook_config_file":            nil,
		"authorization_webhook_cache_authorized_ttl":   nil,
		"authorization_webhook_cache_unauthorized_ttl": nil,
		"authorization_rbac_super_user":                nil,
		"encryption_provider_config":                   nil,
		"experimental_encryption_provider_config":      nil,
		"requestheader_username_headers":               func() []interface{} { return nil }(),
		"requestheader_group_headers":                  func() []interface{} { return nil }(),
		"requestheader_extra_header_prefixes":          func() []interface{} { return nil }(),
		"requestheader_client_ca_file":                 "",
		"requestheader_allowed_names":                  func() []interface{} { return nil }(),
		"feature_gates":                                func() map[string]interface{} { return nil }(),
		"max_requests_inflight":                        0,
		"max_mutating_requests_inflight":               0,
		"http2_max_streams_per_connection":             nil,
		"etcd_quorum_read":                             nil,
		"request_timeout":                              nil,
		"min_request_timeout":                          nil,
		"target_ram_mb":                                0,
		"service_account_key_file":                     func() []interface{} { return nil }(),
		"service_account_signing_key_file":             nil,
		"service_account_issuer":                       nil,
		"service_account_jwksuri":                      nil,
		"api_audiences":                                func() []interface{} { return nil }(),
		"cpu_request":                                  "",
		"cpu_limit":                                    "",
		"memory_request":                               "",
		"memory_limit":                                 "",
		"event_ttl":                                    nil,
		"audit_dynamic_configuration":                  nil,
		"enable_profiling":                             nil,
		"cors_allowed_origins":                         func() []interface{} { return nil }(),
		"default_not_ready_toleration_seconds":         nil,
		"default_unreachable_toleration_seconds":       nil,
	}
	type args struct {
		in kops.KubeAPIServerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KubeAPIServerConfig{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableBasicAuth - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.DisableBasicAuth = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogFormat - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.LogFormat = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.LogLevel = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudProvider - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.CloudProvider = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecurePort - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.SecurePort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InsecurePort - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.InsecurePort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Address - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.Address = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BindAddress - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.BindAddress = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InsecureBindAddress - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.InsecureBindAddress = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableBootstrapAuthToken - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EnableBootstrapAuthToken = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableAggregatorRouting - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EnableAggregatorRouting = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdmissionControl - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AdmissionControl = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AppendAdmissionPlugins - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AppendAdmissionPlugins = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableAdmissionPlugins - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EnableAdmissionPlugins = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableAdmissionPlugins - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.DisableAdmissionPlugins = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdmissionControlConfigFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AdmissionControlConfigFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceClusterIpRange - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ServiceClusterIPRange = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceNodePortRange - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ServiceNodePortRange = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdServers - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EtcdServers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdServersOverrides - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EtcdServersOverrides = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdCAFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EtcdCAFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdCertFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EtcdCertFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdKeyFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EtcdKeyFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BasicAuthFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.BasicAuthFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClientCAFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ClientCAFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSCertFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.TLSCertFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSPrivateKeyFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.TLSPrivateKeyFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSCipherSuites - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.TLSCipherSuites = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSMinVersion - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.TLSMinVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TokenAuthFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.TokenAuthFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowPrivileged - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AllowPrivileged = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ApiServerCount - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.APIServerCount = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RuntimeConfig - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.RuntimeConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeletClientCertificate - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.KubeletClientCertificate = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeletCertificateAuthority - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.KubeletCertificateAuthority = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeletClientKey - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.KubeletClientKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AnonymousAuth - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AnonymousAuth = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeletPreferredAddressTypes - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.KubeletPreferredAddressTypes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "StorageBackend - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.StorageBackend = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OIDCUsernameClaim - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.OIDCUsernameClaim = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OIDCUsernamePrefix - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.OIDCUsernamePrefix = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OIDCGroupsClaim - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.OIDCGroupsClaim = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OIDCGroupsPrefix - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.OIDCGroupsPrefix = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OIDCIssuerURL - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.OIDCIssuerURL = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OIDCClientID - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.OIDCClientID = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OIDCRequiredClaim - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.OIDCRequiredClaim = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OIDCCAFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.OIDCCAFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ProxyClientCertFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ProxyClientCertFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ProxyClientKeyFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ProxyClientKeyFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditLogFormat - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditLogFormat = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditLogPath - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditLogPath = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditLogMaxAge - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditLogMaxAge = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditLogMaxBackups - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditLogMaxBackups = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditLogMaxSize - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditLogMaxSize = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditPolicyFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditPolicyFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookBatchBufferSize - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookBatchBufferSize = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookBatchMaxSize - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookBatchMaxSize = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookBatchMaxWait - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookBatchMaxWait = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookBatchThrottleBurst - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookBatchThrottleBurst = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookBatchThrottleEnable - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookBatchThrottleEnable = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookBatchThrottleQps - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookBatchThrottleQps = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookConfigFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookConfigFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookInitialBackoff - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookInitialBackoff = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditWebhookMode - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditWebhookMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthenticationTokenWebhookConfigFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuthenticationTokenWebhookConfigFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthenticationTokenWebhookCacheTTL - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuthenticationTokenWebhookCacheTTL = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationMode - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuthorizationMode = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationWebhookConfigFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuthorizationWebhookConfigFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationWebhookCacheAuthorizedTTL - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuthorizationWebhookCacheAuthorizedTTL = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationWebhookCacheUnauthorizedTTL - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuthorizationWebhookCacheUnauthorizedTTL = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationRBACSuperUser - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuthorizationRBACSuperUser = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EncryptionProviderConfig - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EncryptionProviderConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExperimentalEncryptionProviderConfig - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ExperimentalEncryptionProviderConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RequestheaderUsernameHeaders - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.RequestheaderUsernameHeaders = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RequestheaderGroupHeaders - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.RequestheaderGroupHeaders = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RequestheaderExtraHeaderPrefixes - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.RequestheaderExtraHeaderPrefixes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RequestheaderClientCAFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.RequestheaderClientCAFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RequestheaderAllowedNames - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.RequestheaderAllowedNames = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FeatureGates - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.FeatureGates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxRequestsInflight - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.MaxRequestsInflight = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxMutatingRequestsInflight - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.MaxMutatingRequestsInflight = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HTTP2MaxStreamsPerConnection - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.HTTP2MaxStreamsPerConnection = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdQuorumRead - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EtcdQuorumRead = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RequestTimeout - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.RequestTimeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MinRequestTimeout - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.MinRequestTimeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TargetRamMb - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.TargetRamMb = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountKeyFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ServiceAccountKeyFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountSigningKeyFile - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ServiceAccountSigningKeyFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountIssuer - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ServiceAccountIssuer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountJWKSURI - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.ServiceAccountJWKSURI = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ApiAudiences - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.APIAudiences = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.CPURequest = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuLimit - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.CPULimit = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.MemoryRequest = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.MemoryLimit = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EventTTL - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EventTTL = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuditDynamicConfiguration - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.AuditDynamicConfiguration = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableProfiling - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.EnableProfiling = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CorsAllowedOrigins - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.CorsAllowedOrigins = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DefaultNotReadyTolerationSeconds - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.DefaultNotReadyTolerationSeconds = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DefaultUnreachableTolerationSeconds - default",
			args: args{
				in: func() kops.KubeAPIServerConfig {
					subject := kops.KubeAPIServerConfig{}
					subject.DefaultUnreachableTolerationSeconds = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceKubeAPIServerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceKubeAPIServerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
