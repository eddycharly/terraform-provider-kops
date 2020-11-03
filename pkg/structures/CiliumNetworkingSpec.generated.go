package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandCiliumNetworkingSpec(in map[string]interface{}) kops.CiliumNetworkingSpec {
	if in == nil {
		panic("expand CiliumNetworkingSpec failure, in is nil")
	}
	return kops.CiliumNetworkingSpec{
		Version: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["version"]),
		AccessLog: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["access_log"]),
		AgentLabels: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["agent_labels"]),
		AgentPrometheusPort: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["agent_prometheus_port"]),
		AllowLocalhost: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["allow_localhost"]),
		AutoIpv6NodeRoutes: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["auto_ipv_6node_routes"]),
		BPFRoot: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["bpf_root"]),
		ContainerRuntime: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["container_runtime"]),
		ContainerRuntimeEndpoint: func(in interface{}) map[string]string {
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
		}(in["container_runtime_endpoint"]),
		Debug: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["debug"]),
		DebugVerbose: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["debug_verbose"]),
		Device: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["device"]),
		DisableConntrack: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disable_conntrack"]),
		DisableIpv4: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disable_ipv_4"]),
		DisableK8sServices: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disable_k8s_services"]),
		EnablePolicy: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["enable_policy"]),
		EnableTracing: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_tracing"]),
		EnablePrometheusMetrics: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_prometheus_metrics"]),
		EnvoyLog: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["envoy_log"]),
		Ipv4ClusterCIDRMaskSize: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["ipv_4cluster_cidr_mask_size"]),
		Ipv4Node: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv_4node"]),
		Ipv4Range: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv_4range"]),
		Ipv4ServiceRange: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv_4service_range"]),
		Ipv6ClusterAllocCidr: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv_6cluster_alloc_cidr"]),
		Ipv6Node: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv_6node"]),
		Ipv6Range: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv_6range"]),
		Ipv6ServiceRange: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv_6service_range"]),
		K8sAPIServer: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["k8s_api_server"]),
		K8sKubeconfigPath: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["k8s_kubeconfig_path"]),
		KeepBPFTemplates: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["keep_bpf_templates"]),
		KeepConfig: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["keep_config"]),
		LabelPrefixFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["label_prefix_file"]),
		Labels: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["labels"]),
		LB: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["lb"]),
		LibDir: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["lib_dir"]),
		LogDrivers: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["log_drivers"]),
		LogOpt: func(in interface{}) map[string]string {
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
		}(in["log_opt"]),
		Logstash: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["logstash"]),
		LogstashAgent: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["logstash_agent"]),
		LogstashProbeTimer: func(in interface{}) uint32 {
			return uint32(ExpandInt(in))
		}(in["logstash_probe_timer"]),
		DisableMasquerade: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disable_masquerade"]),
		Nat46Range: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["nat_46range"]),
		Pprof: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["pprof"]),
		PrefilterDevice: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["prefilter_device"]),
		PrometheusServeAddr: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["prometheus_serve_addr"]),
		Restore: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["restore"]),
		SingleClusterRoute: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["single_cluster_route"]),
		SocketPath: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["socket_path"]),
		StateDir: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["state_dir"]),
		TracePayloadLen: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["trace_payload_len"]),
		Tunnel: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["tunnel"]),
		EnableIpv6: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_ipv_6"]),
		EnableIpv4: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_ipv_4"]),
		MonitorAggregation: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["monitor_aggregation"]),
		BPFCTGlobalTCPMax: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["bpfct_global_tcp_max"]),
		BPFCTGlobalAnyMax: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["bpfct_global_any_max"]),
		PreallocateBPFMaps: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["preallocate_bpf_maps"]),
		SidecarIstioProxyImage: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["sidecar_istio_proxy_image"]),
		ClusterName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_name"]),
		ToFqdnsDNSRejectResponseCode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["to_fqdns_dns_reject_response_code"]),
		ToFqdnsEnablePoller: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["to_fqdns_enable_poller"]),
		ContainerRuntimeLabels: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["container_runtime_labels"]),
		Ipam: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipam"]),
		IPTablesRulesNoinstall: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["ip_tables_rules_noinstall"]),
		AutoDirectNodeRoutes: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["auto_direct_node_routes"]),
		EnableNodePort: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_node_port"]),
		EtcdManaged: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["etcd_managed"]),
		EnableRemoteNodeIdentity: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_remote_node_identity"]),
		RemoveCbrBridge: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["remove_cbr_bridge"]),
		RestartPods: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["restart_pods"]),
		ReconfigureKubelet: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["reconfigure_kubelet"]),
		NodeInitBootstrapFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["node_init_bootstrap_file"]),
		CniBinPath: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cni_bin_path"]),
	}
}

func FlattenCiliumNetworkingSpec(in kops.CiliumNetworkingSpec) map[string]interface{} {
	return map[string]interface{}{
		"version": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Version),
		"access_log": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.AccessLog),
		"agent_labels": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.AgentLabels),
		"agent_prometheus_port": func(in int) interface{} {
			return FlattenInt(int(in))
		}(in.AgentPrometheusPort),
		"allow_localhost": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.AllowLocalhost),
		"auto_ipv_6node_routes": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.AutoIpv6NodeRoutes),
		"bpf_root": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.BPFRoot),
		"container_runtime": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.ContainerRuntime),
		"container_runtime_endpoint": func(in map[string]string) interface{} {
			return func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
		}(in.ContainerRuntimeEndpoint),
		"debug": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.Debug),
		"debug_verbose": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.DebugVerbose),
		"device": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Device),
		"disable_conntrack": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.DisableConntrack),
		"disable_ipv_4": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.DisableIpv4),
		"disable_k8s_services": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.DisableK8sServices),
		"enable_policy": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.EnablePolicy),
		"enable_tracing": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.EnableTracing),
		"enable_prometheus_metrics": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.EnablePrometheusMetrics),
		"envoy_log": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.EnvoyLog),
		"ipv_4cluster_cidr_mask_size": func(in int) interface{} {
			return FlattenInt(int(in))
		}(in.Ipv4ClusterCIDRMaskSize),
		"ipv_4node": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Ipv4Node),
		"ipv_4range": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Ipv4Range),
		"ipv_4service_range": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Ipv4ServiceRange),
		"ipv_6cluster_alloc_cidr": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Ipv6ClusterAllocCidr),
		"ipv_6node": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Ipv6Node),
		"ipv_6range": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Ipv6Range),
		"ipv_6service_range": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Ipv6ServiceRange),
		"k8s_api_server": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.K8sAPIServer),
		"k8s_kubeconfig_path": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.K8sKubeconfigPath),
		"keep_bpf_templates": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.KeepBPFTemplates),
		"keep_config": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.KeepConfig),
		"label_prefix_file": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.LabelPrefixFile),
		"labels": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.Labels),
		"lb": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.LB),
		"lib_dir": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.LibDir),
		"log_drivers": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.LogDrivers),
		"log_opt": func(in map[string]string) interface{} {
			return func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
		}(in.LogOpt),
		"logstash": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.Logstash),
		"logstash_agent": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.LogstashAgent),
		"logstash_probe_timer": func(in uint32) interface{} {
			return FlattenInt(int(in))
		}(in.LogstashProbeTimer),
		"disable_masquerade": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.DisableMasquerade),
		"nat_46range": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Nat46Range),
		"pprof": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.Pprof),
		"prefilter_device": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.PrefilterDevice),
		"prometheus_serve_addr": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.PrometheusServeAddr),
		"restore": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.Restore),
		"single_cluster_route": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.SingleClusterRoute),
		"socket_path": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.SocketPath),
		"state_dir": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.StateDir),
		"trace_payload_len": func(in int) interface{} {
			return FlattenInt(int(in))
		}(in.TracePayloadLen),
		"tunnel": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Tunnel),
		"enable_ipv_6": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.EnableIpv6),
		"enable_ipv_4": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.EnableIpv4),
		"monitor_aggregation": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.MonitorAggregation),
		"bpfct_global_tcp_max": func(in int) interface{} {
			return FlattenInt(int(in))
		}(in.BPFCTGlobalTCPMax),
		"bpfct_global_any_max": func(in int) interface{} {
			return FlattenInt(int(in))
		}(in.BPFCTGlobalAnyMax),
		"preallocate_bpf_maps": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.PreallocateBPFMaps),
		"sidecar_istio_proxy_image": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.SidecarIstioProxyImage),
		"cluster_name": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ClusterName),
		"to_fqdns_dns_reject_response_code": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ToFqdnsDNSRejectResponseCode),
		"to_fqdns_enable_poller": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.ToFqdnsEnablePoller),
		"container_runtime_labels": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ContainerRuntimeLabels),
		"ipam": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Ipam),
		"ip_tables_rules_noinstall": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.IPTablesRulesNoinstall),
		"auto_direct_node_routes": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.AutoDirectNodeRoutes),
		"enable_node_port": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.EnableNodePort),
		"etcd_managed": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.EtcdManaged),
		"enable_remote_node_identity": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.EnableRemoteNodeIdentity),
		"remove_cbr_bridge": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.RemoveCbrBridge),
		"restart_pods": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.RestartPods),
		"reconfigure_kubelet": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.ReconfigureKubelet),
		"node_init_bootstrap_file": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.NodeInitBootstrapFile),
		"cni_bin_path": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.CniBinPath),
	}
}
