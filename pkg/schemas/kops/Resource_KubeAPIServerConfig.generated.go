package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceKubeAPIServerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image":                                        OptionalString(),
			"disable_basic_auth":                           OptionalBool(),
			"log_level":                                    OptionalInt(),
			"cloud_provider":                               OptionalString(),
			"secure_port":                                  OptionalInt(),
			"insecure_port":                                OptionalInt(),
			"address":                                      OptionalString(),
			"bind_address":                                 OptionalString(),
			"insecure_bind_address":                        OptionalString(),
			"enable_bootstrap_auth_token":                  OptionalBool(),
			"enable_aggregator_routing":                    OptionalBool(),
			"admission_control":                            OptionalList(String()),
			"append_admission_plugins":                     OptionalList(String()),
			"enable_admission_plugins":                     OptionalList(String()),
			"disable_admission_plugins":                    OptionalList(String()),
			"admission_control_config_file":                OptionalString(),
			"service_cluster_ip_range":                     OptionalString(),
			"service_node_port_range":                      OptionalString(),
			"etcd_servers":                                 OptionalList(String()),
			"etcd_servers_overrides":                       OptionalList(String()),
			"etcd_ca_file":                                 OptionalString(),
			"etcd_cert_file":                               OptionalString(),
			"etcd_key_file":                                OptionalString(),
			"basic_auth_file":                              OptionalString(),
			"client_ca_file":                               OptionalString(),
			"tls_cert_file":                                OptionalString(),
			"tls_private_key_file":                         OptionalString(),
			"tls_cipher_suites":                            OptionalList(String()),
			"tls_min_version":                              OptionalString(),
			"token_auth_file":                              OptionalString(),
			"allow_privileged":                             OptionalBool(),
			"api_server_count":                             OptionalInt(),
			"runtime_config":                               OptionalMap(String()),
			"kubelet_client_certificate":                   OptionalString(),
			"kubelet_certificate_authority":                OptionalString(),
			"kubelet_client_key":                           OptionalString(),
			"anonymous_auth":                               OptionalBool(),
			"kubelet_preferred_address_types":              OptionalList(String()),
			"storage_backend":                              OptionalString(),
			"oidc_username_claim":                          OptionalString(),
			"oidc_username_prefix":                         OptionalString(),
			"oidc_groups_claim":                            OptionalString(),
			"oidc_groups_prefix":                           OptionalString(),
			"oidc_issuer_url":                              OptionalString(),
			"oidc_client_id":                               OptionalString(),
			"oidc_required_claim":                          OptionalList(String()),
			"oidcca_file":                                  OptionalString(),
			"proxy_client_cert_file":                       OptionalString(),
			"proxy_client_key_file":                        OptionalString(),
			"audit_log_format":                             OptionalString(),
			"audit_log_path":                               OptionalString(),
			"audit_log_max_age":                            OptionalInt(),
			"audit_log_max_backups":                        OptionalInt(),
			"audit_log_max_size":                           OptionalInt(),
			"audit_policy_file":                            OptionalString(),
			"audit_webhook_batch_buffer_size":              OptionalInt(),
			"audit_webhook_batch_max_size":                 OptionalInt(),
			"audit_webhook_batch_max_wait":                 OptionalDuration(),
			"audit_webhook_batch_throttle_burst":           OptionalInt(),
			"audit_webhook_batch_throttle_enable":          OptionalBool(),
			"audit_webhook_batch_throttle_qps":             OptionalQuantity(),
			"audit_webhook_config_file":                    OptionalString(),
			"audit_webhook_initial_backoff":                OptionalDuration(),
			"audit_webhook_mode":                           OptionalString(),
			"authentication_token_webhook_config_file":     OptionalString(),
			"authentication_token_webhook_cache_ttl":       OptionalDuration(),
			"authorization_mode":                           OptionalString(),
			"authorization_webhook_config_file":            OptionalString(),
			"authorization_webhook_cache_authorized_ttl":   OptionalDuration(),
			"authorization_webhook_cache_unauthorized_ttl": OptionalDuration(),
			"authorization_rbac_super_user":                OptionalString(),
			"encryption_provider_config":                   OptionalString(),
			"experimental_encryption_provider_config":      OptionalString(),
			"requestheader_username_headers":               OptionalList(String()),
			"requestheader_group_headers":                  OptionalList(String()),
			"requestheader_extra_header_prefixes":          OptionalList(String()),
			"requestheader_client_ca_file":                 OptionalString(),
			"requestheader_allowed_names":                  OptionalList(String()),
			"feature_gates":                                OptionalMap(String()),
			"max_requests_inflight":                        OptionalInt(),
			"max_mutating_requests_inflight":               OptionalInt(),
			"http2_max_streams_per_connection":             OptionalInt(),
			"etcd_quorum_read":                             OptionalBool(),
			"request_timeout":                              OptionalDuration(),
			"min_request_timeout":                          OptionalInt(),
			"target_ram_mb":                                OptionalInt(),
			"service_account_key_file":                     OptionalList(String()),
			"service_account_signing_key_file":             OptionalString(),
			"service_account_issuer":                       OptionalString(),
			"service_account_jwksuri":                      OptionalString(),
			"api_audiences":                                OptionalList(String()),
			"cpu_request":                                  OptionalString(),
			"cpu_limit":                                    OptionalString(),
			"memory_request":                               OptionalString(),
			"memory_limit":                                 OptionalString(),
			"event_ttl":                                    OptionalDuration(),
			"audit_dynamic_configuration":                  OptionalBool(),
			"enable_profiling":                             OptionalBool(),
			"cors_allowed_origins":                         OptionalList(String()),
			"default_not_ready_toleration_seconds":         OptionalInt(),
			"default_unreachable_toleration_seconds":       OptionalInt(),
		},
	}
}

func ExpandResourceKubeAPIServerConfig(in map[string]interface{}) kops.KubeAPIServerConfig {
	if in == nil {
		panic("expand KubeAPIServerConfig failure, in is nil")
	}
	return kops.KubeAPIServerConfig{
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		DisableBasicAuth: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["disable_basic_auth"]),
		LogLevel: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["log_level"]),
		CloudProvider: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cloud_provider"]),
		SecurePort: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["secure_port"]),
		InsecurePort: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["insecure_port"]),
		Address: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["address"]),
		BindAddress: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["bind_address"]),
		InsecureBindAddress: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["insecure_bind_address"]),
		EnableBootstrapAuthToken: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["enable_bootstrap_auth_token"]),
		EnableAggregatorRouting: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["enable_aggregator_routing"]),
		AdmissionControl: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["admission_control"]),
		AppendAdmissionPlugins: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["append_admission_plugins"]),
		EnableAdmissionPlugins: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["enable_admission_plugins"]),
		DisableAdmissionPlugins: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["disable_admission_plugins"]),
		AdmissionControlConfigFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["admission_control_config_file"]),
		ServiceClusterIPRange: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["service_cluster_ip_range"]),
		ServiceNodePortRange: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["service_node_port_range"]),
		EtcdServers: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["etcd_servers"]),
		EtcdServersOverrides: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["etcd_servers_overrides"]),
		EtcdCAFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["etcd_ca_file"]),
		EtcdCertFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["etcd_cert_file"]),
		EtcdKeyFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["etcd_key_file"]),
		BasicAuthFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["basic_auth_file"]),
		ClientCAFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["client_ca_file"]),
		TLSCertFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["tls_cert_file"]),
		TLSPrivateKeyFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["tls_private_key_file"]),
		TLSCipherSuites: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["tls_cipher_suites"]),
		TLSMinVersion: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["tls_min_version"]),
		TokenAuthFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["token_auth_file"]),
		AllowPrivileged: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["allow_privileged"]),
		APIServerCount: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["api_server_count"]),
		RuntimeConfig: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
		}(in["runtime_config"]),
		KubeletClientCertificate: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kubelet_client_certificate"]),
		KubeletCertificateAuthority: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kubelet_certificate_authority"]),
		KubeletClientKey: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kubelet_client_key"]),
		AnonymousAuth: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["anonymous_auth"]),
		KubeletPreferredAddressTypes: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["kubelet_preferred_address_types"]),
		StorageBackend: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["storage_backend"]),
		OIDCUsernameClaim: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["oidc_username_claim"]),
		OIDCUsernamePrefix: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["oidc_username_prefix"]),
		OIDCGroupsClaim: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["oidc_groups_claim"]),
		OIDCGroupsPrefix: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["oidc_groups_prefix"]),
		OIDCIssuerURL: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["oidc_issuer_url"]),
		OIDCClientID: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["oidc_client_id"]),
		OIDCRequiredClaim: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["oidc_required_claim"]),
		OIDCCAFile: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["oidcca_file"]),
		ProxyClientCertFile: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["proxy_client_cert_file"]),
		ProxyClientKeyFile: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["proxy_client_key_file"]),
		AuditLogFormat: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["audit_log_format"]),
		AuditLogPath: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["audit_log_path"]),
		AuditLogMaxAge: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["audit_log_max_age"]),
		AuditLogMaxBackups: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["audit_log_max_backups"]),
		AuditLogMaxSize: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["audit_log_max_size"]),
		AuditPolicyFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["audit_policy_file"]),
		AuditWebhookBatchBufferSize: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["audit_webhook_batch_buffer_size"]),
		AuditWebhookBatchMaxSize: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["audit_webhook_batch_max_size"]),
		AuditWebhookBatchMaxWait: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["audit_webhook_batch_max_wait"]),
		AuditWebhookBatchThrottleBurst: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["audit_webhook_batch_throttle_burst"]),
		AuditWebhookBatchThrottleEnable: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["audit_webhook_batch_throttle_enable"]),
		AuditWebhookBatchThrottleQps: func(in interface{}) *resource.Quantity {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					return &in
				}(ExpandQuantity(in))
			}(in)
		}(in["audit_webhook_batch_throttle_qps"]),
		AuditWebhookConfigFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["audit_webhook_config_file"]),
		AuditWebhookInitialBackoff: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["audit_webhook_initial_backoff"]),
		AuditWebhookMode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["audit_webhook_mode"]),
		AuthenticationTokenWebhookConfigFile: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["authentication_token_webhook_config_file"]),
		AuthenticationTokenWebhookCacheTTL: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["authentication_token_webhook_cache_ttl"]),
		AuthorizationMode: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["authorization_mode"]),
		AuthorizationWebhookConfigFile: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["authorization_webhook_config_file"]),
		AuthorizationWebhookCacheAuthorizedTTL: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["authorization_webhook_cache_authorized_ttl"]),
		AuthorizationWebhookCacheUnauthorizedTTL: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["authorization_webhook_cache_unauthorized_ttl"]),
		AuthorizationRBACSuperUser: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["authorization_rbac_super_user"]),
		EncryptionProviderConfig: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["encryption_provider_config"]),
		ExperimentalEncryptionProviderConfig: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["experimental_encryption_provider_config"]),
		RequestheaderUsernameHeaders: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["requestheader_username_headers"]),
		RequestheaderGroupHeaders: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["requestheader_group_headers"]),
		RequestheaderExtraHeaderPrefixes: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["requestheader_extra_header_prefixes"]),
		RequestheaderClientCAFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["requestheader_client_ca_file"]),
		RequestheaderAllowedNames: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["requestheader_allowed_names"]),
		FeatureGates: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
		}(in["feature_gates"]),
		MaxRequestsInflight: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["max_requests_inflight"]),
		MaxMutatingRequestsInflight: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["max_mutating_requests_inflight"]),
		HTTP2MaxStreamsPerConnection: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["http2_max_streams_per_connection"]),
		EtcdQuorumRead: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["etcd_quorum_read"]),
		RequestTimeout: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["request_timeout"]),
		MinRequestTimeout: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["min_request_timeout"]),
		TargetRamMb: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["target_ram_mb"]),
		ServiceAccountKeyFile: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["service_account_key_file"]),
		ServiceAccountSigningKeyFile: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["service_account_signing_key_file"]),
		ServiceAccountIssuer: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["service_account_issuer"]),
		ServiceAccountJWKSURI: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["service_account_jwksuri"]),
		APIAudiences: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["api_audiences"]),
		CPURequest: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cpu_request"]),
		CPULimit: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cpu_limit"]),
		MemoryRequest: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["memory_request"]),
		MemoryLimit: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["memory_limit"]),
		EventTTL: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["event_ttl"]),
		AuditDynamicConfiguration: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["audit_dynamic_configuration"]),
		EnableProfiling: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["enable_profiling"]),
		CorsAllowedOrigins: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["cors_allowed_origins"]),
		DefaultNotReadyTolerationSeconds: func(in interface{}) *int64 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					return &in
				}(int64(ExpandInt(in)))
			}(in)
		}(in["default_not_ready_toleration_seconds"]),
		DefaultUnreachableTolerationSeconds: func(in interface{}) *int64 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					return &in
				}(int64(ExpandInt(in)))
			}(in)
		}(in["default_unreachable_toleration_seconds"]),
	}
}

func FlattenResourceKubeAPIServerConfigInto(in kops.KubeAPIServerConfig, out map[string]interface{}) {
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
	out["disable_basic_auth"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.DisableBasicAuth)
	out["log_level"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.LogLevel)
	out["cloud_provider"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CloudProvider)
	out["secure_port"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.SecurePort)
	out["insecure_port"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.InsecurePort)
	out["address"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Address)
	out["bind_address"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.BindAddress)
	out["insecure_bind_address"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.InsecureBindAddress)
	out["enable_bootstrap_auth_token"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EnableBootstrapAuthToken)
	out["enable_aggregator_routing"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EnableAggregatorRouting)
	out["admission_control"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AdmissionControl)
	out["append_admission_plugins"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AppendAdmissionPlugins)
	out["enable_admission_plugins"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.EnableAdmissionPlugins)
	out["disable_admission_plugins"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.DisableAdmissionPlugins)
	out["admission_control_config_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AdmissionControlConfigFile)
	out["service_cluster_ip_range"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ServiceClusterIPRange)
	out["service_node_port_range"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ServiceNodePortRange)
	out["etcd_servers"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.EtcdServers)
	out["etcd_servers_overrides"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.EtcdServersOverrides)
	out["etcd_ca_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.EtcdCAFile)
	out["etcd_cert_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.EtcdCertFile)
	out["etcd_key_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.EtcdKeyFile)
	out["basic_auth_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.BasicAuthFile)
	out["client_ca_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClientCAFile)
	out["tls_cert_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.TLSCertFile)
	out["tls_private_key_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.TLSPrivateKeyFile)
	out["tls_cipher_suites"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.TLSCipherSuites)
	out["tls_min_version"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.TLSMinVersion)
	out["token_auth_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.TokenAuthFile)
	out["allow_privileged"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.AllowPrivileged)
	out["api_server_count"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.APIServerCount)
	out["runtime_config"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.RuntimeConfig)
	out["kubelet_client_certificate"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.KubeletClientCertificate)
	out["kubelet_certificate_authority"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.KubeletCertificateAuthority)
	out["kubelet_client_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.KubeletClientKey)
	out["anonymous_auth"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.AnonymousAuth)
	out["kubelet_preferred_address_types"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.KubeletPreferredAddressTypes)
	out["storage_backend"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.StorageBackend)
	out["oidc_username_claim"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.OIDCUsernameClaim)
	out["oidc_username_prefix"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.OIDCUsernamePrefix)
	out["oidc_groups_claim"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.OIDCGroupsClaim)
	out["oidc_groups_prefix"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.OIDCGroupsPrefix)
	out["oidc_issuer_url"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.OIDCIssuerURL)
	out["oidc_client_id"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.OIDCClientID)
	out["oidc_required_claim"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.OIDCRequiredClaim)
	out["oidcca_file"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.OIDCCAFile)
	out["proxy_client_cert_file"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.ProxyClientCertFile)
	out["proxy_client_key_file"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.ProxyClientKeyFile)
	out["audit_log_format"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.AuditLogFormat)
	out["audit_log_path"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.AuditLogPath)
	out["audit_log_max_age"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.AuditLogMaxAge)
	out["audit_log_max_backups"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.AuditLogMaxBackups)
	out["audit_log_max_size"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.AuditLogMaxSize)
	out["audit_policy_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AuditPolicyFile)
	out["audit_webhook_batch_buffer_size"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.AuditWebhookBatchBufferSize)
	out["audit_webhook_batch_max_size"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.AuditWebhookBatchMaxSize)
	out["audit_webhook_batch_max_wait"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.AuditWebhookBatchMaxWait)
	out["audit_webhook_batch_throttle_burst"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.AuditWebhookBatchThrottleBurst)
	out["audit_webhook_batch_throttle_enable"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.AuditWebhookBatchThrottleEnable)
	out["audit_webhook_batch_throttle_qps"] = func(in *resource.Quantity) interface{} {
		return func(in *resource.Quantity) interface{} {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) interface{} {
				return FlattenQuantity(in)
			}(*in)
		}(in)
	}(in.AuditWebhookBatchThrottleQps)
	out["audit_webhook_config_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AuditWebhookConfigFile)
	out["audit_webhook_initial_backoff"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.AuditWebhookInitialBackoff)
	out["audit_webhook_mode"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AuditWebhookMode)
	out["authentication_token_webhook_config_file"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.AuthenticationTokenWebhookConfigFile)
	out["authentication_token_webhook_cache_ttl"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.AuthenticationTokenWebhookCacheTTL)
	out["authorization_mode"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.AuthorizationMode)
	out["authorization_webhook_config_file"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.AuthorizationWebhookConfigFile)
	out["authorization_webhook_cache_authorized_ttl"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.AuthorizationWebhookCacheAuthorizedTTL)
	out["authorization_webhook_cache_unauthorized_ttl"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.AuthorizationWebhookCacheUnauthorizedTTL)
	out["authorization_rbac_super_user"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.AuthorizationRBACSuperUser)
	out["encryption_provider_config"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.EncryptionProviderConfig)
	out["experimental_encryption_provider_config"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.ExperimentalEncryptionProviderConfig)
	out["requestheader_username_headers"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.RequestheaderUsernameHeaders)
	out["requestheader_group_headers"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.RequestheaderGroupHeaders)
	out["requestheader_extra_header_prefixes"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.RequestheaderExtraHeaderPrefixes)
	out["requestheader_client_ca_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.RequestheaderClientCAFile)
	out["requestheader_allowed_names"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.RequestheaderAllowedNames)
	out["feature_gates"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.FeatureGates)
	out["max_requests_inflight"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.MaxRequestsInflight)
	out["max_mutating_requests_inflight"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.MaxMutatingRequestsInflight)
	out["http2_max_streams_per_connection"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.HTTP2MaxStreamsPerConnection)
	out["etcd_quorum_read"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EtcdQuorumRead)
	out["request_timeout"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.RequestTimeout)
	out["min_request_timeout"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.MinRequestTimeout)
	out["target_ram_mb"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.TargetRamMb)
	out["service_account_key_file"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.ServiceAccountKeyFile)
	out["service_account_signing_key_file"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.ServiceAccountSigningKeyFile)
	out["service_account_issuer"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.ServiceAccountIssuer)
	out["service_account_jwksuri"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.ServiceAccountJWKSURI)
	out["api_audiences"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.APIAudiences)
	out["cpu_request"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CPURequest)
	out["cpu_limit"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CPULimit)
	out["memory_request"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.MemoryRequest)
	out["memory_limit"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.MemoryLimit)
	out["event_ttl"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.EventTTL)
	out["audit_dynamic_configuration"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.AuditDynamicConfiguration)
	out["enable_profiling"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EnableProfiling)
	out["cors_allowed_origins"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.CorsAllowedOrigins)
	out["default_not_ready_toleration_seconds"] = func(in *int64) interface{} {
		return func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.DefaultNotReadyTolerationSeconds)
	out["default_unreachable_toleration_seconds"] = func(in *int64) interface{} {
		return func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.DefaultUnreachableTolerationSeconds)
}

func FlattenResourceKubeAPIServerConfig(in kops.KubeAPIServerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKubeAPIServerConfigInto(in, out)
	return out
}
