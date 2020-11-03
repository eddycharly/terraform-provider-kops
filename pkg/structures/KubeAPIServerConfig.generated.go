package structures

import (
	"reflect"

	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandKubeAPIServerConfig(in map[string]interface{}) kops.KubeAPIServerConfig {
	if in == nil {
		panic("expand KubeAPIServerConfig failure, in is nil")
	}
	return kops.KubeAPIServerConfig{
		Image: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["image"]),
		DisableBasicAuth: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["disable_basic_auth"]),
		LogLevel: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			return value
		}(in["log_level"]),
		CloudProvider: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cloud_provider"]),
		SecurePort: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			return value
		}(in["secure_port"]),
		InsecurePort: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			return value
		}(in["insecure_port"]),
		Address: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["address"]),
		BindAddress: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["bind_address"]),
		InsecureBindAddress: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["insecure_bind_address"]),
		EnableBootstrapAuthToken: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["enable_bootstrap_auth_token"]),
		EnableAggregatorRouting: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["enable_aggregator_routing"]),
		AdmissionControl: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["admission_control"]),
		AppendAdmissionPlugins: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["append_admission_plugins"]),
		EnableAdmissionPlugins: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["enable_admission_plugins"]),
		DisableAdmissionPlugins: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["disable_admission_plugins"]),
		AdmissionControlConfigFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["admission_control_config_file"]),
		ServiceClusterIPRange: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["service_cluster_ip_range"]),
		ServiceNodePortRange: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["service_node_port_range"]),
		EtcdServers: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["etcd_servers"]),
		EtcdServersOverrides: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["etcd_servers_overrides"]),
		EtcdCAFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["etcd_ca_file"]),
		EtcdCertFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["etcd_cert_file"]),
		EtcdKeyFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["etcd_key_file"]),
		BasicAuthFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["basic_auth_file"]),
		ClientCAFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["client_ca_file"]),
		TLSCertFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["tls_cert_file"]),
		TLSPrivateKeyFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["tls_private_key_file"]),
		TLSCipherSuites: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["tls_cipher_suites"]),
		TLSMinVersion: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["tls_min_version"]),
		TokenAuthFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["token_auth_file"]),
		AllowPrivileged: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["allow_privileged"]),
		APIServerCount: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["api_server_count"]),
		RuntimeConfig: func(in interface{}) map[string]string {
			value := func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
			return value
		}(in["runtime_config"]),
		KubeletClientCertificate: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["kubelet_client_certificate"]),
		KubeletCertificateAuthority: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["kubelet_certificate_authority"]),
		KubeletClientKey: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["kubelet_client_key"]),
		AnonymousAuth: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["anonymous_auth"]),
		KubeletPreferredAddressTypes: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["kubelet_preferred_address_types"]),
		StorageBackend: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["storage_backend"]),
		OIDCUsernameClaim: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["oidc_username_claim"]),
		OIDCUsernamePrefix: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["oidc_username_prefix"]),
		OIDCGroupsClaim: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["oidc_groups_claim"]),
		OIDCGroupsPrefix: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["oidc_groups_prefix"]),
		OIDCIssuerURL: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["oidc_issuer_url"]),
		OIDCClientID: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["oidc_client_id"]),
		OIDCRequiredClaim: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["oidc_required_claim"]),
		OIDCCAFile: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["oidcca_file"]),
		ProxyClientCertFile: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["proxy_client_cert_file"]),
		ProxyClientKeyFile: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["proxy_client_key_file"]),
		AuditLogFormat: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["audit_log_format"]),
		AuditLogPath: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["audit_log_path"]),
		AuditLogMaxAge: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["audit_log_max_age"]),
		AuditLogMaxBackups: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["audit_log_max_backups"]),
		AuditLogMaxSize: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["audit_log_max_size"]),
		AuditPolicyFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["audit_policy_file"]),
		AuditWebhookBatchBufferSize: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["audit_webhook_batch_buffer_size"]),
		AuditWebhookBatchMaxSize: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["audit_webhook_batch_max_size"]),
		AuditWebhookBatchMaxWait: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["audit_webhook_batch_max_wait"]),
		AuditWebhookBatchThrottleBurst: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["audit_webhook_batch_throttle_burst"]),
		AuditWebhookBatchThrottleEnable: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["audit_webhook_batch_throttle_enable"]),
		AuditWebhookBatchThrottleQps: func(in interface{}) *resource.Quantity {
			value := func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandQuantity(in))
			}(in)
			return value
		}(in["audit_webhook_batch_throttle_qps"]),
		AuditWebhookConfigFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["audit_webhook_config_file"]),
		AuditWebhookInitialBackoff: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["audit_webhook_initial_backoff"]),
		AuditWebhookMode: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["audit_webhook_mode"]),
		AuthenticationTokenWebhookConfigFile: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["authentication_token_webhook_config_file"]),
		AuthenticationTokenWebhookCacheTTL: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["authentication_token_webhook_cache_ttl"]),
		AuthorizationMode: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["authorization_mode"]),
		AuthorizationWebhookConfigFile: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["authorization_webhook_config_file"]),
		AuthorizationWebhookCacheAuthorizedTTL: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["authorization_webhook_cache_authorized_ttl"]),
		AuthorizationWebhookCacheUnauthorizedTTL: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["authorization_webhook_cache_unauthorized_ttl"]),
		AuthorizationRBACSuperUser: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["authorization_rbac_super_user"]),
		EncryptionProviderConfig: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["encryption_provider_config"]),
		ExperimentalEncryptionProviderConfig: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["experimental_encryption_provider_config"]),
		RequestheaderUsernameHeaders: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["requestheader_username_headers"]),
		RequestheaderGroupHeaders: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["requestheader_group_headers"]),
		RequestheaderExtraHeaderPrefixes: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["requestheader_extra_header_prefixes"]),
		RequestheaderClientCAFile: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["requestheader_client_ca_file"]),
		RequestheaderAllowedNames: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["requestheader_allowed_names"]),
		FeatureGates: func(in interface{}) map[string]string {
			value := func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
			return value
		}(in["feature_gates"]),
		MaxRequestsInflight: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			return value
		}(in["max_requests_inflight"]),
		MaxMutatingRequestsInflight: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			return value
		}(in["max_mutating_requests_inflight"]),
		HTTP2MaxStreamsPerConnection: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["http_2max_streams_per_connection"]),
		EtcdQuorumRead: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["etcd_quorum_read"]),
		RequestTimeout: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["request_timeout"]),
		MinRequestTimeout: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
			}(in)
			return value
		}(in["min_request_timeout"]),
		TargetRamMb: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			return value
		}(in["target_ram_mb"]),
		ServiceAccountKeyFile: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["service_account_key_file"]),
		ServiceAccountSigningKeyFile: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["service_account_signing_key_file"]),
		ServiceAccountIssuer: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["service_account_issuer"]),
		ServiceAccountJWKSURI: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["service_account_jwksuri"]),
		APIAudiences: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["api_audiences"]),
		CPURequest: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cpu_request"]),
		EventTTL: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["event_ttl"]),
		AuditDynamicConfiguration: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["audit_dynamic_configuration"]),
		EnableProfiling: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["enable_profiling"]),
	}
}

func FlattenKubeAPIServerConfig(in kops.KubeAPIServerConfig) map[string]interface{} {
	return map[string]interface{}{
		"image": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Image),
		"disable_basic_auth": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.DisableBasicAuth),
		"log_level": func(in int32) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.LogLevel),
		"cloud_provider": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.CloudProvider),
		"secure_port": func(in int32) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.SecurePort),
		"insecure_port": func(in int32) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.InsecurePort),
		"address": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Address),
		"bind_address": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.BindAddress),
		"insecure_bind_address": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.InsecureBindAddress),
		"enable_bootstrap_auth_token": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.EnableBootstrapAuthToken),
		"enable_aggregator_routing": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.EnableAggregatorRouting),
		"admission_control": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.AdmissionControl),
		"append_admission_plugins": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.AppendAdmissionPlugins),
		"enable_admission_plugins": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.EnableAdmissionPlugins),
		"disable_admission_plugins": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.DisableAdmissionPlugins),
		"admission_control_config_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.AdmissionControlConfigFile),
		"service_cluster_ip_range": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ServiceClusterIPRange),
		"service_node_port_range": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ServiceNodePortRange),
		"etcd_servers": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.EtcdServers),
		"etcd_servers_overrides": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.EtcdServersOverrides),
		"etcd_ca_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.EtcdCAFile),
		"etcd_cert_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.EtcdCertFile),
		"etcd_key_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.EtcdKeyFile),
		"basic_auth_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.BasicAuthFile),
		"client_ca_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ClientCAFile),
		"tls_cert_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.TLSCertFile),
		"tls_private_key_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.TLSPrivateKeyFile),
		"tls_cipher_suites": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.TLSCipherSuites),
		"tls_min_version": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.TLSMinVersion),
		"token_auth_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.TokenAuthFile),
		"allow_privileged": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.AllowPrivileged),
		"api_server_count": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.APIServerCount),
		"runtime_config": func(in map[string]string) interface{} {
			value := func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
			return value
		}(in.RuntimeConfig),
		"kubelet_client_certificate": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.KubeletClientCertificate),
		"kubelet_certificate_authority": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.KubeletCertificateAuthority),
		"kubelet_client_key": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.KubeletClientKey),
		"anonymous_auth": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.AnonymousAuth),
		"kubelet_preferred_address_types": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.KubeletPreferredAddressTypes),
		"storage_backend": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.StorageBackend),
		"oidc_username_claim": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.OIDCUsernameClaim),
		"oidc_username_prefix": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.OIDCUsernamePrefix),
		"oidc_groups_claim": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.OIDCGroupsClaim),
		"oidc_groups_prefix": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.OIDCGroupsPrefix),
		"oidc_issuer_url": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.OIDCIssuerURL),
		"oidc_client_id": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.OIDCClientID),
		"oidc_required_claim": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.OIDCRequiredClaim),
		"oidcca_file": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.OIDCCAFile),
		"proxy_client_cert_file": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.ProxyClientCertFile),
		"proxy_client_key_file": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.ProxyClientKeyFile),
		"audit_log_format": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.AuditLogFormat),
		"audit_log_path": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.AuditLogPath),
		"audit_log_max_age": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.AuditLogMaxAge),
		"audit_log_max_backups": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.AuditLogMaxBackups),
		"audit_log_max_size": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.AuditLogMaxSize),
		"audit_policy_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.AuditPolicyFile),
		"audit_webhook_batch_buffer_size": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.AuditWebhookBatchBufferSize),
		"audit_webhook_batch_max_size": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.AuditWebhookBatchMaxSize),
		"audit_webhook_batch_max_wait": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.AuditWebhookBatchMaxWait),
		"audit_webhook_batch_throttle_burst": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.AuditWebhookBatchThrottleBurst),
		"audit_webhook_batch_throttle_enable": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.AuditWebhookBatchThrottleEnable),
		"audit_webhook_batch_throttle_qps": func(in *resource.Quantity) interface{} {
			value := func(in *resource.Quantity) interface{} {
				if in == nil {
					return nil
				}
				return func(in resource.Quantity) interface{} {
					return FlattenQuantity(in)
				}(*in)
			}(in)
			return value
		}(in.AuditWebhookBatchThrottleQps),
		"audit_webhook_config_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.AuditWebhookConfigFile),
		"audit_webhook_initial_backoff": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.AuditWebhookInitialBackoff),
		"audit_webhook_mode": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.AuditWebhookMode),
		"authentication_token_webhook_config_file": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.AuthenticationTokenWebhookConfigFile),
		"authentication_token_webhook_cache_ttl": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.AuthenticationTokenWebhookCacheTTL),
		"authorization_mode": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.AuthorizationMode),
		"authorization_webhook_config_file": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.AuthorizationWebhookConfigFile),
		"authorization_webhook_cache_authorized_ttl": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.AuthorizationWebhookCacheAuthorizedTTL),
		"authorization_webhook_cache_unauthorized_ttl": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.AuthorizationWebhookCacheUnauthorizedTTL),
		"authorization_rbac_super_user": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.AuthorizationRBACSuperUser),
		"encryption_provider_config": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.EncryptionProviderConfig),
		"experimental_encryption_provider_config": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.ExperimentalEncryptionProviderConfig),
		"requestheader_username_headers": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.RequestheaderUsernameHeaders),
		"requestheader_group_headers": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.RequestheaderGroupHeaders),
		"requestheader_extra_header_prefixes": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.RequestheaderExtraHeaderPrefixes),
		"requestheader_client_ca_file": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.RequestheaderClientCAFile),
		"requestheader_allowed_names": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.RequestheaderAllowedNames),
		"feature_gates": func(in map[string]string) interface{} {
			value := func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
			return value
		}(in.FeatureGates),
		"max_requests_inflight": func(in int32) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.MaxRequestsInflight),
		"max_mutating_requests_inflight": func(in int32) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.MaxMutatingRequestsInflight),
		"http_2max_streams_per_connection": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.HTTP2MaxStreamsPerConnection),
		"etcd_quorum_read": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.EtcdQuorumRead),
		"request_timeout": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.RequestTimeout),
		"min_request_timeout": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.MinRequestTimeout),
		"target_ram_mb": func(in int32) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.TargetRamMb),
		"service_account_key_file": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.ServiceAccountKeyFile),
		"service_account_signing_key_file": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.ServiceAccountSigningKeyFile),
		"service_account_issuer": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.ServiceAccountIssuer),
		"service_account_jwksuri": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.ServiceAccountJWKSURI),
		"api_audiences": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.APIAudiences),
		"cpu_request": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.CPURequest),
		"event_ttl": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.EventTTL),
		"audit_dynamic_configuration": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.AuditDynamicConfiguration),
		"enable_profiling": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.EnableProfiling),
	}
}
