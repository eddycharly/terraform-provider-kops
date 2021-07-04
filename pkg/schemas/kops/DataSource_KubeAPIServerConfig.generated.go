package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceKubeAPIServerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image":                                        Computed(String()),
			"disable_basic_auth":                           Computed(Nullable(Bool())),
			"log_level":                                    Computed(Int()),
			"cloud_provider":                               Computed(String()),
			"secure_port":                                  Computed(Int()),
			"insecure_port":                                Computed(Int()),
			"address":                                      Computed(String()),
			"bind_address":                                 Computed(String()),
			"insecure_bind_address":                        Computed(String()),
			"enable_bootstrap_auth_token":                  Computed(Nullable(Bool())),
			"enable_aggregator_routing":                    Computed(Nullable(Bool())),
			"admission_control":                            Computed(List(String())),
			"append_admission_plugins":                     Computed(List(String())),
			"enable_admission_plugins":                     Computed(List(String())),
			"disable_admission_plugins":                    Computed(List(String())),
			"admission_control_config_file":                Computed(String()),
			"service_cluster_ip_range":                     Computed(String()),
			"service_node_port_range":                      Computed(String()),
			"etcd_servers":                                 Computed(List(String())),
			"etcd_servers_overrides":                       Computed(List(String())),
			"etcd_ca_file":                                 Computed(String()),
			"etcd_cert_file":                               Computed(String()),
			"etcd_key_file":                                Computed(String()),
			"basic_auth_file":                              Computed(String()),
			"client_ca_file":                               Computed(String()),
			"tls_cert_file":                                Computed(String()),
			"tls_private_key_file":                         Computed(String()),
			"tls_cipher_suites":                            Computed(List(String())),
			"tls_min_version":                              Computed(String()),
			"token_auth_file":                              Computed(String()),
			"allow_privileged":                             Computed(Nullable(Bool())),
			"api_server_count":                             Computed(Nullable(Int())),
			"runtime_config":                               Computed(Map(String())),
			"kubelet_client_certificate":                   Computed(String()),
			"kubelet_certificate_authority":                Computed(String()),
			"kubelet_client_key":                           Computed(String()),
			"anonymous_auth":                               Computed(Nullable(Bool())),
			"kubelet_preferred_address_types":              Computed(List(String())),
			"storage_backend":                              Computed(Nullable(String())),
			"oidc_username_claim":                          Computed(Nullable(String())),
			"oidc_username_prefix":                         Computed(Nullable(String())),
			"oidc_groups_claim":                            Computed(Nullable(String())),
			"oidc_groups_prefix":                           Computed(Nullable(String())),
			"oidc_issuer_url":                              Computed(Nullable(String())),
			"oidc_client_id":                               Computed(Nullable(String())),
			"oidc_required_claim":                          Computed(List(String())),
			"oidcca_file":                                  Computed(Nullable(String())),
			"proxy_client_cert_file":                       Computed(Nullable(String())),
			"proxy_client_key_file":                        Computed(Nullable(String())),
			"audit_log_format":                             Computed(Nullable(String())),
			"audit_log_path":                               Computed(Nullable(String())),
			"audit_log_max_age":                            Computed(Nullable(Int())),
			"audit_log_max_backups":                        Computed(Nullable(Int())),
			"audit_log_max_size":                           Computed(Nullable(Int())),
			"audit_policy_file":                            Computed(String()),
			"audit_webhook_batch_buffer_size":              Computed(Nullable(Int())),
			"audit_webhook_batch_max_size":                 Computed(Nullable(Int())),
			"audit_webhook_batch_max_wait":                 Computed(Nullable(Duration())),
			"audit_webhook_batch_throttle_burst":           Computed(Nullable(Int())),
			"audit_webhook_batch_throttle_enable":          Computed(Nullable(Bool())),
			"audit_webhook_batch_throttle_qps":             Computed(Nullable(Quantity())),
			"audit_webhook_config_file":                    Computed(String()),
			"audit_webhook_initial_backoff":                Computed(Nullable(Duration())),
			"audit_webhook_mode":                           Computed(String()),
			"authentication_token_webhook_config_file":     Computed(Nullable(String())),
			"authentication_token_webhook_cache_ttl":       Computed(Nullable(Duration())),
			"authorization_mode":                           Computed(Nullable(String())),
			"authorization_webhook_config_file":            Computed(Nullable(String())),
			"authorization_webhook_cache_authorized_ttl":   Computed(Nullable(Duration())),
			"authorization_webhook_cache_unauthorized_ttl": Computed(Nullable(Duration())),
			"authorization_rbac_super_user":                Computed(Nullable(String())),
			"encryption_provider_config":                   Computed(Nullable(String())),
			"experimental_encryption_provider_config":      Computed(Nullable(String())),
			"requestheader_username_headers":               Computed(List(String())),
			"requestheader_group_headers":                  Computed(List(String())),
			"requestheader_extra_header_prefixes":          Computed(List(String())),
			"requestheader_client_ca_file":                 Computed(String()),
			"requestheader_allowed_names":                  Computed(List(String())),
			"feature_gates":                                Computed(Map(String())),
			"max_requests_inflight":                        Computed(Int()),
			"max_mutating_requests_inflight":               Computed(Int()),
			"http2_max_streams_per_connection":             Computed(Nullable(Int())),
			"etcd_quorum_read":                             Computed(Nullable(Bool())),
			"request_timeout":                              Computed(Nullable(Duration())),
			"min_request_timeout":                          Computed(Nullable(Int())),
			"target_ram_mb":                                Computed(Int()),
			"service_account_key_file":                     Computed(List(String())),
			"service_account_signing_key_file":             Computed(Nullable(String())),
			"service_account_issuer":                       Computed(Nullable(String())),
			"service_account_jwksuri":                      Computed(Nullable(String())),
			"api_audiences":                                Computed(List(String())),
			"cpu_request":                                  Computed(String()),
			"cpu_limit":                                    Computed(String()),
			"memory_request":                               Computed(String()),
			"memory_limit":                                 Computed(String()),
			"event_ttl":                                    Computed(Nullable(Duration())),
			"audit_dynamic_configuration":                  Computed(Nullable(Bool())),
			"enable_profiling":                             Computed(Nullable(Bool())),
			"cors_allowed_origins":                         Computed(List(String())),
			"default_not_ready_toleration_seconds":         Computed(Nullable(Int())),
			"default_unreachable_toleration_seconds":       Computed(Nullable(Int())),
		},
	}
}

func ExpandDataSourceKubeAPIServerConfig(in map[string]interface{}) kops.KubeAPIServerConfig {
	if in == nil {
		panic("expand KubeAPIServerConfig failure, in is nil")
	}
	out := kops.KubeAPIServerConfig{}
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["disable_basic_auth"]; ok && in != nil {
		out.DisableBasicAuth = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["log_level"]; ok && in != nil {
		out.LogLevel = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	if in, ok := in["cloud_provider"]; ok && in != nil {
		out.CloudProvider = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["secure_port"]; ok && in != nil {
		out.SecurePort = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	if in, ok := in["insecure_port"]; ok && in != nil {
		out.InsecurePort = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	if in, ok := in["address"]; ok && in != nil {
		out.Address = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["bind_address"]; ok && in != nil {
		out.BindAddress = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["insecure_bind_address"]; ok && in != nil {
		out.InsecureBindAddress = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["enable_bootstrap_auth_token"]; ok && in != nil {
		out.EnableBootstrapAuthToken = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["enable_aggregator_routing"]; ok && in != nil {
		out.EnableAggregatorRouting = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["admission_control"]; ok && in != nil {
		out.AdmissionControl = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["append_admission_plugins"]; ok && in != nil {
		out.AppendAdmissionPlugins = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["enable_admission_plugins"]; ok && in != nil {
		out.EnableAdmissionPlugins = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["disable_admission_plugins"]; ok && in != nil {
		out.DisableAdmissionPlugins = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["admission_control_config_file"]; ok && in != nil {
		out.AdmissionControlConfigFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["service_cluster_ip_range"]; ok && in != nil {
		out.ServiceClusterIPRange = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["service_node_port_range"]; ok && in != nil {
		out.ServiceNodePortRange = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["etcd_servers"]; ok && in != nil {
		out.EtcdServers = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["etcd_servers_overrides"]; ok && in != nil {
		out.EtcdServersOverrides = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["etcd_ca_file"]; ok && in != nil {
		out.EtcdCAFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["etcd_cert_file"]; ok && in != nil {
		out.EtcdCertFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["etcd_key_file"]; ok && in != nil {
		out.EtcdKeyFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["basic_auth_file"]; ok && in != nil {
		out.BasicAuthFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["client_ca_file"]; ok && in != nil {
		out.ClientCAFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["tls_cert_file"]; ok && in != nil {
		out.TLSCertFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["tls_private_key_file"]; ok && in != nil {
		out.TLSPrivateKeyFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["tls_cipher_suites"]; ok && in != nil {
		out.TLSCipherSuites = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["tls_min_version"]; ok && in != nil {
		out.TLSMinVersion = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["token_auth_file"]; ok && in != nil {
		out.TokenAuthFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["allow_privileged"]; ok && in != nil {
		out.AllowPrivileged = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["api_server_count"]; ok && in != nil {
		out.APIServerCount = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["runtime_config"]; ok && in != nil {
		out.RuntimeConfig = func(in interface{}) map[string]string {
			if in == nil {
				return nil
			}
			out := map[string]string{}
			for key, in := range in.(map[string]interface{}) {
				out[key] = func(in interface{}) string { return string(in.(string)) }(in)
			}
			return out
		}(in)
	}
	if in, ok := in["kubelet_client_certificate"]; ok && in != nil {
		out.KubeletClientCertificate = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["kubelet_certificate_authority"]; ok && in != nil {
		out.KubeletCertificateAuthority = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["kubelet_client_key"]; ok && in != nil {
		out.KubeletClientKey = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["anonymous_auth"]; ok && in != nil {
		out.AnonymousAuth = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["kubelet_preferred_address_types"]; ok && in != nil {
		out.KubeletPreferredAddressTypes = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["storage_backend"]; ok && in != nil {
		out.StorageBackend = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["oidc_username_claim"]; ok && in != nil {
		out.OIDCUsernameClaim = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["oidc_username_prefix"]; ok && in != nil {
		out.OIDCUsernamePrefix = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["oidc_groups_claim"]; ok && in != nil {
		out.OIDCGroupsClaim = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["oidc_groups_prefix"]; ok && in != nil {
		out.OIDCGroupsPrefix = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["oidc_issuer_url"]; ok && in != nil {
		out.OIDCIssuerURL = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["oidc_client_id"]; ok && in != nil {
		out.OIDCClientID = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["oidc_required_claim"]; ok && in != nil {
		out.OIDCRequiredClaim = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["oidcca_file"]; ok && in != nil {
		out.OIDCCAFile = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["proxy_client_cert_file"]; ok && in != nil {
		out.ProxyClientCertFile = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["proxy_client_key_file"]; ok && in != nil {
		out.ProxyClientKeyFile = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["audit_log_format"]; ok && in != nil {
		out.AuditLogFormat = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["audit_log_path"]; ok && in != nil {
		out.AuditLogPath = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["audit_log_max_age"]; ok && in != nil {
		out.AuditLogMaxAge = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["audit_log_max_backups"]; ok && in != nil {
		out.AuditLogMaxBackups = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["audit_log_max_size"]; ok && in != nil {
		out.AuditLogMaxSize = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["audit_policy_file"]; ok && in != nil {
		out.AuditPolicyFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["audit_webhook_batch_buffer_size"]; ok && in != nil {
		out.AuditWebhookBatchBufferSize = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["audit_webhook_batch_max_size"]; ok && in != nil {
		out.AuditWebhookBatchMaxSize = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["audit_webhook_batch_max_wait"]; ok && in != nil {
		out.AuditWebhookBatchMaxWait = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["audit_webhook_batch_throttle_burst"]; ok && in != nil {
		out.AuditWebhookBatchThrottleBurst = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["audit_webhook_batch_throttle_enable"]; ok && in != nil {
		out.AuditWebhookBatchThrottleEnable = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["audit_webhook_batch_throttle_qps"]; ok && in != nil {
		out.AuditWebhookBatchThrottleQps = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["audit_webhook_config_file"]; ok && in != nil {
		out.AuditWebhookConfigFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["audit_webhook_initial_backoff"]; ok && in != nil {
		out.AuditWebhookInitialBackoff = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["audit_webhook_mode"]; ok && in != nil {
		out.AuditWebhookMode = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["authentication_token_webhook_config_file"]; ok && in != nil {
		out.AuthenticationTokenWebhookConfigFile = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["authentication_token_webhook_cache_ttl"]; ok && in != nil {
		out.AuthenticationTokenWebhookCacheTTL = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["authorization_mode"]; ok && in != nil {
		out.AuthorizationMode = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["authorization_webhook_config_file"]; ok && in != nil {
		out.AuthorizationWebhookConfigFile = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["authorization_webhook_cache_authorized_ttl"]; ok && in != nil {
		out.AuthorizationWebhookCacheAuthorizedTTL = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["authorization_webhook_cache_unauthorized_ttl"]; ok && in != nil {
		out.AuthorizationWebhookCacheUnauthorizedTTL = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["authorization_rbac_super_user"]; ok && in != nil {
		out.AuthorizationRBACSuperUser = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["encryption_provider_config"]; ok && in != nil {
		out.EncryptionProviderConfig = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["experimental_encryption_provider_config"]; ok && in != nil {
		out.ExperimentalEncryptionProviderConfig = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["requestheader_username_headers"]; ok && in != nil {
		out.RequestheaderUsernameHeaders = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["requestheader_group_headers"]; ok && in != nil {
		out.RequestheaderGroupHeaders = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["requestheader_extra_header_prefixes"]; ok && in != nil {
		out.RequestheaderExtraHeaderPrefixes = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["requestheader_client_ca_file"]; ok && in != nil {
		out.RequestheaderClientCAFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["requestheader_allowed_names"]; ok && in != nil {
		out.RequestheaderAllowedNames = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["feature_gates"]; ok && in != nil {
		out.FeatureGates = func(in interface{}) map[string]string {
			if in == nil {
				return nil
			}
			out := map[string]string{}
			for key, in := range in.(map[string]interface{}) {
				out[key] = func(in interface{}) string { return string(in.(string)) }(in)
			}
			return out
		}(in)
	}
	if in, ok := in["max_requests_inflight"]; ok && in != nil {
		out.MaxRequestsInflight = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	if in, ok := in["max_mutating_requests_inflight"]; ok && in != nil {
		out.MaxMutatingRequestsInflight = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	if in, ok := in["http2_max_streams_per_connection"]; ok && in != nil {
		out.HTTP2MaxStreamsPerConnection = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["etcd_quorum_read"]; ok && in != nil {
		out.EtcdQuorumRead = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["request_timeout"]; ok && in != nil {
		out.RequestTimeout = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["min_request_timeout"]; ok && in != nil {
		out.MinRequestTimeout = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["target_ram_mb"]; ok && in != nil {
		out.TargetRamMb = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	if in, ok := in["service_account_key_file"]; ok && in != nil {
		out.ServiceAccountKeyFile = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["service_account_signing_key_file"]; ok && in != nil {
		out.ServiceAccountSigningKeyFile = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["service_account_issuer"]; ok && in != nil {
		out.ServiceAccountIssuer = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["service_account_jwksuri"]; ok && in != nil {
		out.ServiceAccountJWKSURI = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["api_audiences"]; ok && in != nil {
		out.APIAudiences = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["cpu_request"]; ok && in != nil {
		out.CPURequest = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cpu_limit"]; ok && in != nil {
		out.CPULimit = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["memory_request"]; ok && in != nil {
		out.MemoryRequest = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["memory_limit"]; ok && in != nil {
		out.MemoryLimit = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["event_ttl"]; ok && in != nil {
		out.EventTTL = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["audit_dynamic_configuration"]; ok && in != nil {
		out.AuditDynamicConfiguration = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["enable_profiling"]; ok && in != nil {
		out.EnableProfiling = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["cors_allowed_origins"]; ok && in != nil {
		out.CorsAllowedOrigins = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["default_not_ready_toleration_seconds"]; ok && in != nil {
		out.DefaultNotReadyTolerationSeconds = func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			return func(in int64) *int64 { return &in }(func(in interface{}) int64 { return int64(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["default_unreachable_toleration_seconds"]; ok && in != nil {
		out.DefaultUnreachableTolerationSeconds = func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			return func(in int64) *int64 { return &in }(func(in interface{}) int64 { return int64(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}

func FlattenDataSourceKubeAPIServerConfigInto(in kops.KubeAPIServerConfig, out map[string]interface{}) {
	out["image"] = func(in string) interface{} { return string(in) }(in.Image)
	out["disable_basic_auth"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.DisableBasicAuth)
	out["log_level"] = func(in int32) interface{} { return int(in) }(in.LogLevel)
	out["cloud_provider"] = func(in string) interface{} { return string(in) }(in.CloudProvider)
	out["secure_port"] = func(in int32) interface{} { return int(in) }(in.SecurePort)
	out["insecure_port"] = func(in int32) interface{} { return int(in) }(in.InsecurePort)
	out["address"] = func(in string) interface{} { return string(in) }(in.Address)
	out["bind_address"] = func(in string) interface{} { return string(in) }(in.BindAddress)
	out["insecure_bind_address"] = func(in string) interface{} { return string(in) }(in.InsecureBindAddress)
	out["enable_bootstrap_auth_token"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.EnableBootstrapAuthToken)
	out["enable_aggregator_routing"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.EnableAggregatorRouting)
	out["admission_control"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.AdmissionControl)
	out["append_admission_plugins"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.AppendAdmissionPlugins)
	out["enable_admission_plugins"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.EnableAdmissionPlugins)
	out["disable_admission_plugins"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.DisableAdmissionPlugins)
	out["admission_control_config_file"] = func(in string) interface{} { return string(in) }(in.AdmissionControlConfigFile)
	out["service_cluster_ip_range"] = func(in string) interface{} { return string(in) }(in.ServiceClusterIPRange)
	out["service_node_port_range"] = func(in string) interface{} { return string(in) }(in.ServiceNodePortRange)
	out["etcd_servers"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.EtcdServers)
	out["etcd_servers_overrides"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.EtcdServersOverrides)
	out["etcd_ca_file"] = func(in string) interface{} { return string(in) }(in.EtcdCAFile)
	out["etcd_cert_file"] = func(in string) interface{} { return string(in) }(in.EtcdCertFile)
	out["etcd_key_file"] = func(in string) interface{} { return string(in) }(in.EtcdKeyFile)
	out["basic_auth_file"] = func(in string) interface{} { return string(in) }(in.BasicAuthFile)
	out["client_ca_file"] = func(in string) interface{} { return string(in) }(in.ClientCAFile)
	out["tls_cert_file"] = func(in string) interface{} { return string(in) }(in.TLSCertFile)
	out["tls_private_key_file"] = func(in string) interface{} { return string(in) }(in.TLSPrivateKeyFile)
	out["tls_cipher_suites"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.TLSCipherSuites)
	out["tls_min_version"] = func(in string) interface{} { return string(in) }(in.TLSMinVersion)
	out["token_auth_file"] = func(in string) interface{} { return string(in) }(in.TokenAuthFile)
	out["allow_privileged"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.AllowPrivileged)
	out["api_server_count"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.APIServerCount)
	out["runtime_config"] = func(in map[string]string) interface{} {
		if in == nil {
			return nil
		}
		out := map[string]interface{}{}
		for key, in := range in {
			out[key] = func(in string) interface{} { return string(in) }(in)
		}
		return out
	}(in.RuntimeConfig)
	out["kubelet_client_certificate"] = func(in string) interface{} { return string(in) }(in.KubeletClientCertificate)
	out["kubelet_certificate_authority"] = func(in string) interface{} { return string(in) }(in.KubeletCertificateAuthority)
	out["kubelet_client_key"] = func(in string) interface{} { return string(in) }(in.KubeletClientKey)
	out["anonymous_auth"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.AnonymousAuth)
	out["kubelet_preferred_address_types"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.KubeletPreferredAddressTypes)
	out["storage_backend"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.StorageBackend)
	out["oidc_username_claim"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.OIDCUsernameClaim)
	out["oidc_username_prefix"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.OIDCUsernamePrefix)
	out["oidc_groups_claim"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.OIDCGroupsClaim)
	out["oidc_groups_prefix"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.OIDCGroupsPrefix)
	out["oidc_issuer_url"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.OIDCIssuerURL)
	out["oidc_client_id"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.OIDCClientID)
	out["oidc_required_claim"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.OIDCRequiredClaim)
	out["oidcca_file"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.OIDCCAFile)
	out["proxy_client_cert_file"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.ProxyClientCertFile)
	out["proxy_client_key_file"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.ProxyClientKeyFile)
	out["audit_log_format"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.AuditLogFormat)
	out["audit_log_path"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.AuditLogPath)
	out["audit_log_max_age"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.AuditLogMaxAge)
	out["audit_log_max_backups"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.AuditLogMaxBackups)
	out["audit_log_max_size"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.AuditLogMaxSize)
	out["audit_policy_file"] = func(in string) interface{} { return string(in) }(in.AuditPolicyFile)
	out["audit_webhook_batch_buffer_size"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.AuditWebhookBatchBufferSize)
	out["audit_webhook_batch_max_size"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.AuditWebhookBatchMaxSize)
	out["audit_webhook_batch_max_wait"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.AuditWebhookBatchMaxWait)
	out["audit_webhook_batch_throttle_burst"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.AuditWebhookBatchThrottleBurst)
	out["audit_webhook_batch_throttle_enable"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.AuditWebhookBatchThrottleEnable)
	out["audit_webhook_batch_throttle_qps"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.AuditWebhookBatchThrottleQps)
	out["audit_webhook_config_file"] = func(in string) interface{} { return string(in) }(in.AuditWebhookConfigFile)
	out["audit_webhook_initial_backoff"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.AuditWebhookInitialBackoff)
	out["audit_webhook_mode"] = func(in string) interface{} { return string(in) }(in.AuditWebhookMode)
	out["authentication_token_webhook_config_file"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.AuthenticationTokenWebhookConfigFile)
	out["authentication_token_webhook_cache_ttl"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.AuthenticationTokenWebhookCacheTTL)
	out["authorization_mode"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.AuthorizationMode)
	out["authorization_webhook_config_file"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.AuthorizationWebhookConfigFile)
	out["authorization_webhook_cache_authorized_ttl"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.AuthorizationWebhookCacheAuthorizedTTL)
	out["authorization_webhook_cache_unauthorized_ttl"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.AuthorizationWebhookCacheUnauthorizedTTL)
	out["authorization_rbac_super_user"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.AuthorizationRBACSuperUser)
	out["encryption_provider_config"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.EncryptionProviderConfig)
	out["experimental_encryption_provider_config"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.ExperimentalEncryptionProviderConfig)
	out["requestheader_username_headers"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.RequestheaderUsernameHeaders)
	out["requestheader_group_headers"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.RequestheaderGroupHeaders)
	out["requestheader_extra_header_prefixes"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.RequestheaderExtraHeaderPrefixes)
	out["requestheader_client_ca_file"] = func(in string) interface{} { return string(in) }(in.RequestheaderClientCAFile)
	out["requestheader_allowed_names"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.RequestheaderAllowedNames)
	out["feature_gates"] = func(in map[string]string) interface{} {
		if in == nil {
			return nil
		}
		out := map[string]interface{}{}
		for key, in := range in {
			out[key] = func(in string) interface{} { return string(in) }(in)
		}
		return out
	}(in.FeatureGates)
	out["max_requests_inflight"] = func(in int32) interface{} { return int(in) }(in.MaxRequestsInflight)
	out["max_mutating_requests_inflight"] = func(in int32) interface{} { return int(in) }(in.MaxMutatingRequestsInflight)
	out["http2_max_streams_per_connection"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.HTTP2MaxStreamsPerConnection)
	out["etcd_quorum_read"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.EtcdQuorumRead)
	out["request_timeout"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.RequestTimeout)
	out["min_request_timeout"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.MinRequestTimeout)
	out["target_ram_mb"] = func(in int32) interface{} { return int(in) }(in.TargetRamMb)
	out["service_account_key_file"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.ServiceAccountKeyFile)
	out["service_account_signing_key_file"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.ServiceAccountSigningKeyFile)
	out["service_account_issuer"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.ServiceAccountIssuer)
	out["service_account_jwksuri"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.ServiceAccountJWKSURI)
	out["api_audiences"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.APIAudiences)
	out["cpu_request"] = func(in string) interface{} { return string(in) }(in.CPURequest)
	out["cpu_limit"] = func(in string) interface{} { return string(in) }(in.CPULimit)
	out["memory_request"] = func(in string) interface{} { return string(in) }(in.MemoryRequest)
	out["memory_limit"] = func(in string) interface{} { return string(in) }(in.MemoryLimit)
	out["event_ttl"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.EventTTL)
	out["audit_dynamic_configuration"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.AuditDynamicConfiguration)
	out["enable_profiling"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.EnableProfiling)
	out["cors_allowed_origins"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.CorsAllowedOrigins)
	out["default_not_ready_toleration_seconds"] = func(in *int64) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int64) interface{} { return int(in) }(*in)}
	}(in.DefaultNotReadyTolerationSeconds)
	out["default_unreachable_toleration_seconds"] = func(in *int64) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int64) interface{} { return int(in) }(*in)}
	}(in.DefaultUnreachableTolerationSeconds)
}

func FlattenDataSourceKubeAPIServerConfig(in kops.KubeAPIServerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKubeAPIServerConfigInto(in, out)
	return out
}
