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
			"disable_basic_auth":                           Computed(Ptr(Bool())),
			"log_level":                                    Computed(Int()),
			"cloud_provider":                               Computed(String()),
			"secure_port":                                  Computed(Int()),
			"insecure_port":                                Computed(Int()),
			"address":                                      Computed(String()),
			"bind_address":                                 Computed(String()),
			"insecure_bind_address":                        Computed(String()),
			"enable_bootstrap_auth_token":                  Computed(Ptr(Bool())),
			"enable_aggregator_routing":                    Computed(Ptr(Bool())),
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
			"allow_privileged":                             Computed(Ptr(Bool())),
			"api_server_count":                             Computed(Ptr(Int())),
			"runtime_config":                               Computed(Map(String())),
			"kubelet_client_certificate":                   Computed(String()),
			"kubelet_certificate_authority":                Computed(String()),
			"kubelet_client_key":                           Computed(String()),
			"anonymous_auth":                               Computed(Ptr(Bool())),
			"kubelet_preferred_address_types":              Computed(List(String())),
			"storage_backend":                              Computed(Ptr(String())),
			"oidc_username_claim":                          Computed(Ptr(String())),
			"oidc_username_prefix":                         Computed(Ptr(String())),
			"oidc_groups_claim":                            Computed(Ptr(String())),
			"oidc_groups_prefix":                           Computed(Ptr(String())),
			"oidc_issuer_url":                              Computed(Ptr(String())),
			"oidc_client_id":                               Computed(Ptr(String())),
			"oidc_required_claim":                          Computed(List(String())),
			"oidcca_file":                                  Computed(Ptr(String())),
			"proxy_client_cert_file":                       Computed(Ptr(String())),
			"proxy_client_key_file":                        Computed(Ptr(String())),
			"audit_log_format":                             Computed(Ptr(String())),
			"audit_log_path":                               Computed(Ptr(String())),
			"audit_log_max_age":                            Computed(Ptr(Int())),
			"audit_log_max_backups":                        Computed(Ptr(Int())),
			"audit_log_max_size":                           Computed(Ptr(Int())),
			"audit_policy_file":                            Computed(String()),
			"audit_webhook_batch_buffer_size":              Computed(Ptr(Int())),
			"audit_webhook_batch_max_size":                 Computed(Ptr(Int())),
			"audit_webhook_batch_max_wait":                 Computed(Ptr(Duration())),
			"audit_webhook_batch_throttle_burst":           Computed(Ptr(Int())),
			"audit_webhook_batch_throttle_enable":          Computed(Ptr(Bool())),
			"audit_webhook_batch_throttle_qps":             Computed(Ptr(Quantity())),
			"audit_webhook_config_file":                    Computed(String()),
			"audit_webhook_initial_backoff":                Computed(Ptr(Duration())),
			"audit_webhook_mode":                           Computed(String()),
			"authentication_token_webhook_config_file":     Computed(Ptr(String())),
			"authentication_token_webhook_cache_ttl":       Computed(Ptr(Duration())),
			"authorization_mode":                           Computed(Ptr(String())),
			"authorization_webhook_config_file":            Computed(Ptr(String())),
			"authorization_webhook_cache_authorized_ttl":   Computed(Ptr(Duration())),
			"authorization_webhook_cache_unauthorized_ttl": Computed(Ptr(Duration())),
			"authorization_rbac_super_user":                Computed(Ptr(String())),
			"encryption_provider_config":                   Computed(Ptr(String())),
			"experimental_encryption_provider_config":      Computed(Ptr(String())),
			"requestheader_username_headers":               Computed(List(String())),
			"requestheader_group_headers":                  Computed(List(String())),
			"requestheader_extra_header_prefixes":          Computed(List(String())),
			"requestheader_client_ca_file":                 Computed(String()),
			"requestheader_allowed_names":                  Computed(List(String())),
			"feature_gates":                                Computed(Map(String())),
			"max_requests_inflight":                        Computed(Int()),
			"max_mutating_requests_inflight":               Computed(Int()),
			"http2_max_streams_per_connection":             Computed(Ptr(Int())),
			"etcd_quorum_read":                             Computed(Ptr(Bool())),
			"request_timeout":                              Computed(Ptr(Duration())),
			"min_request_timeout":                          Computed(Ptr(Int())),
			"target_ram_mb":                                Computed(Int()),
			"service_account_key_file":                     Computed(List(String())),
			"service_account_signing_key_file":             Computed(Ptr(String())),
			"service_account_issuer":                       Computed(Ptr(String())),
			"service_account_jwksuri":                      Computed(Ptr(String())),
			"api_audiences":                                Computed(List(String())),
			"cpu_request":                                  Computed(String()),
			"cpu_limit":                                    Computed(String()),
			"memory_request":                               Computed(String()),
			"memory_limit":                                 Computed(String()),
			"event_ttl":                                    Computed(Ptr(Duration())),
			"audit_dynamic_configuration":                  Computed(Ptr(Bool())),
			"enable_profiling":                             Computed(Ptr(Bool())),
			"cors_allowed_origins":                         Computed(List(String())),
			"default_not_ready_toleration_seconds":         Computed(Ptr(Int())),
			"default_unreachable_toleration_seconds":       Computed(Ptr(Int())),
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
