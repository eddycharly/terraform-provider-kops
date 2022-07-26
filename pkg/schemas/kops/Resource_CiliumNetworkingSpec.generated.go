package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceCiliumNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"version":                           OptionalString(),
			"memory_request":                    OptionalQuantity(),
			"cpu_request":                       OptionalQuantity(),
			"agent_prometheus_port":             OptionalInt(),
			"metrics":                           OptionalList(String()),
			"chaining_mode":                     OptionalString(),
			"debug":                             OptionalBool(),
			"disable_endpoint_crd":              OptionalBool(),
			"enable_policy":                     OptionalString(),
			"enable_l7_proxy":                   OptionalBool(),
			"enable_bpf_masquerade":             OptionalBool(),
			"enable_endpoint_health_checking":   OptionalBool(),
			"enable_prometheus_metrics":         OptionalBool(),
			"enable_encryption":                 OptionalBool(),
			"encryption_type":                   OptionalString(),
			"identity_allocation_mode":          OptionalString(),
			"identity_change_grace_period":      OptionalString(),
			"masquerade":                        OptionalBool(),
			"agent_pod_annotations":             OptionalMap(String()),
			"tunnel":                            OptionalString(),
			"monitor_aggregation":               OptionalString(),
			"bpfct_global_tcp_max":              OptionalInt(),
			"bpfct_global_any_max":              OptionalInt(),
			"bpflb_algorithm":                   OptionalString(),
			"bpflb_maglev_table_size":           OptionalString(),
			"bpfnat_global_max":                 OptionalInt(),
			"bpf_neigh_global_max":              OptionalInt(),
			"bpf_policy_map_max":                OptionalInt(),
			"bpflb_map_max":                     OptionalInt(),
			"bpflb_sock_host_ns_only":           OptionalBool(),
			"preallocate_bpf_maps":              RequiredBool(),
			"sidecar_istio_proxy_image":         OptionalString(),
			"cluster_name":                      OptionalString(),
			"to_fqdns_dns_reject_response_code": OptionalString(),
			"to_fqdns_enable_poller":            OptionalBool(),
			"ipam":                              OptionalString(),
			"install_iptables_rules":            OptionalBool(),
			"auto_direct_node_routes":           OptionalBool(),
			"enable_host_reachable_services":    OptionalBool(),
			"enable_node_port":                  OptionalBool(),
			"etcd_managed":                      OptionalBool(),
			"enable_remote_node_identity":       RequiredBool(),
			"hubble":                            OptionalStruct(ResourceHubbleSpec()),
			"disable_cnp_status_updates":        OptionalBool(),
			"enable_service_topology":           OptionalBool(),
		},
	}

	return res
}

func ExpandResourceCiliumNetworkingSpec(in map[string]interface{}) kops.CiliumNetworkingSpec {
	if in == nil {
		panic("expand CiliumNetworkingSpec failure, in is nil")
	}
	return kops.CiliumNetworkingSpec{
		Version: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["version"]),
		MemoryRequest: func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
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
		}(in["memory_request"]),
		CPURequest: func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
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
		}(in["cpu_request"]),
		AgentPrometheusPort: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["agent_prometheus_port"]),
		Metrics: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["metrics"]),
		ChainingMode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["chaining_mode"]),
		Debug: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["debug"]),
		DisableEndpointCRD: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disable_endpoint_crd"]),
		EnablePolicy: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["enable_policy"]),
		EnableL7Proxy: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
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
		}(in["enable_l7_proxy"]),
		EnableBPFMasquerade: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
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
		}(in["enable_bpf_masquerade"]),
		EnableEndpointHealthChecking: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
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
		}(in["enable_endpoint_health_checking"]),
		EnablePrometheusMetrics: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_prometheus_metrics"]),
		EnableEncryption: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_encryption"]),
		EncryptionType: func(in interface{}) kops.CiliumEncryptionType {
			return kops.CiliumEncryptionType(ExpandString(in))
		}(in["encryption_type"]),
		IdentityAllocationMode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["identity_allocation_mode"]),
		IdentityChangeGracePeriod: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["identity_change_grace_period"]),
		Masquerade: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
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
		}(in["masquerade"]),
		AgentPodAnnotations: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string]string{}
						for key, in := range in {
							out[key] = string(ExpandString(in))
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["agent_pod_annotations"]),
		Tunnel: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["tunnel"]),
		MonitorAggregation: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["monitor_aggregation"]),
		BPFCTGlobalTCPMax: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["bpfct_global_tcp_max"]),
		BPFCTGlobalAnyMax: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["bpfct_global_any_max"]),
		BPFLBAlgorithm: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["bpflb_algorithm"]),
		BPFLBMaglevTableSize: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["bpflb_maglev_table_size"]),
		BPFNATGlobalMax: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["bpfnat_global_max"]),
		BPFNeighGlobalMax: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["bpf_neigh_global_max"]),
		BPFPolicyMapMax: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["bpf_policy_map_max"]),
		BPFLBMapMax: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["bpflb_map_max"]),
		BPFLBSockHostNSOnly: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["bpflb_sock_host_ns_only"]),
		PreallocateBPFMaps: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["preallocate_bpf_maps"]),
		SidecarIstioProxyImage: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["sidecar_istio_proxy_image"]),
		ClusterName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_name"]),
		ToFQDNsDNSRejectResponseCode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["to_fqdns_dns_reject_response_code"]),
		ToFQDNsEnablePoller: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["to_fqdns_enable_poller"]),
		IPAM: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipam"]),
		InstallIptablesRules: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
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
		}(in["install_iptables_rules"]),
		AutoDirectNodeRoutes: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["auto_direct_node_routes"]),
		EnableHostReachableServices: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_host_reachable_services"]),
		EnableNodePort: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_node_port"]),
		EtcdManaged: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["etcd_managed"]),
		EnableRemoteNodeIdentity: func(in interface{}) *bool {
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
		}(in["enable_remote_node_identity"]),
		Hubble: func(in interface{}) *kops.HubbleSpec {
			return func(in interface{}) *kops.HubbleSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.HubbleSpec) *kops.HubbleSpec {
					return &in
				}(func(in interface{}) kops.HubbleSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.HubbleSpec{}
					}
					return (ExpandResourceHubbleSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["hubble"]),
		DisableCNPStatusUpdates: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
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
		}(in["disable_cnp_status_updates"]),
		EnableServiceTopology: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_service_topology"]),
	}
}

func FlattenResourceCiliumNetworkingSpecInto(in kops.CiliumNetworkingSpec, out map[string]interface{}) {
	out["version"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Version)
	out["memory_request"] = func(in *resource.Quantity) interface{} {
		return func(in *resource.Quantity) interface{} {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) interface{} {
				return FlattenQuantity(in)
			}(*in)
		}(in)
	}(in.MemoryRequest)
	out["cpu_request"] = func(in *resource.Quantity) interface{} {
		return func(in *resource.Quantity) interface{} {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) interface{} {
				return FlattenQuantity(in)
			}(*in)
		}(in)
	}(in.CPURequest)
	out["agent_prometheus_port"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.AgentPrometheusPort)
	out["metrics"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Metrics)
	out["chaining_mode"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ChainingMode)
	out["debug"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Debug)
	out["disable_endpoint_crd"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.DisableEndpointCRD)
	out["enable_policy"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.EnablePolicy)
	out["enable_l7_proxy"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EnableL7Proxy)
	out["enable_bpf_masquerade"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EnableBPFMasquerade)
	out["enable_endpoint_health_checking"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EnableEndpointHealthChecking)
	out["enable_prometheus_metrics"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnablePrometheusMetrics)
	out["enable_encryption"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableEncryption)
	out["encryption_type"] = func(in kops.CiliumEncryptionType) interface{} {
		return FlattenString(string(in))
	}(in.EncryptionType)
	out["identity_allocation_mode"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.IdentityAllocationMode)
	out["identity_change_grace_period"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.IdentityChangeGracePeriod)
	out["masquerade"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Masquerade)
	out["agent_pod_annotations"] = func(in map[string]string) interface{} {
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
	}(in.AgentPodAnnotations)
	out["tunnel"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Tunnel)
	out["monitor_aggregation"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.MonitorAggregation)
	out["bpfct_global_tcp_max"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.BPFCTGlobalTCPMax)
	out["bpfct_global_any_max"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.BPFCTGlobalAnyMax)
	out["bpflb_algorithm"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.BPFLBAlgorithm)
	out["bpflb_maglev_table_size"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.BPFLBMaglevTableSize)
	out["bpfnat_global_max"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.BPFNATGlobalMax)
	out["bpf_neigh_global_max"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.BPFNeighGlobalMax)
	out["bpf_policy_map_max"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.BPFPolicyMapMax)
	out["bpflb_map_max"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.BPFLBMapMax)
	out["bpflb_sock_host_ns_only"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.BPFLBSockHostNSOnly)
	out["preallocate_bpf_maps"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.PreallocateBPFMaps)
	out["sidecar_istio_proxy_image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.SidecarIstioProxyImage)
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
	out["to_fqdns_dns_reject_response_code"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ToFQDNsDNSRejectResponseCode)
	out["to_fqdns_enable_poller"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.ToFQDNsEnablePoller)
	out["ipam"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.IPAM)
	out["install_iptables_rules"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.InstallIptablesRules)
	out["auto_direct_node_routes"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.AutoDirectNodeRoutes)
	out["enable_host_reachable_services"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableHostReachableServices)
	out["enable_node_port"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableNodePort)
	out["etcd_managed"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EtcdManaged)
	out["enable_remote_node_identity"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EnableRemoteNodeIdentity)
	out["hubble"] = func(in *kops.HubbleSpec) interface{} {
		return func(in *kops.HubbleSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.HubbleSpec) interface{} {
				return func(in kops.HubbleSpec) []interface{} {
					return []interface{}{FlattenResourceHubbleSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Hubble)
	out["disable_cnp_status_updates"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.DisableCNPStatusUpdates)
	out["enable_service_topology"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableServiceTopology)
}

func FlattenResourceCiliumNetworkingSpec(in kops.CiliumNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceCiliumNetworkingSpecInto(in, out)
	return out
}
