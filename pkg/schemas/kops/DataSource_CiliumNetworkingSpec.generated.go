package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceCiliumNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"version":                           Computed(String()),
			"access_log":                        Computed(String()),
			"agent_labels":                      Computed(List(String())),
			"agent_prometheus_port":             Computed(Int()),
			"allow_localhost":                   Computed(String()),
			"auto_ipv6_node_routes":             Computed(Bool()),
			"bpf_root":                          Computed(String()),
			"container_runtime":                 Computed(List(String())),
			"container_runtime_endpoint":        Computed(Map(String())),
			"debug":                             Computed(Bool()),
			"debug_verbose":                     Computed(List(String())),
			"device":                            Computed(String()),
			"disable_conntrack":                 Computed(Bool()),
			"disable_ipv4":                      Computed(Bool()),
			"disable_k_8s_services":             Computed(Bool()),
			"enable_policy":                     Computed(String()),
			"enable_tracing":                    Computed(Bool()),
			"enable_prometheus_metrics":         Computed(Bool()),
			"enable_encryption":                 Computed(Bool()),
			"envoy_log":                         Computed(String()),
			"ipv4_cluster_cidr_mask_size":       Computed(Int()),
			"ipv4_node":                         Computed(String()),
			"ipv4_range":                        Computed(String()),
			"ipv4_service_range":                Computed(String()),
			"ipv6_cluster_alloc_cidr":           Computed(String()),
			"ipv6_node":                         Computed(String()),
			"ipv6_range":                        Computed(String()),
			"ipv6_service_range":                Computed(String()),
			"k_8s_api_server":                   Computed(String()),
			"k_8s_kubeconfig_path":              Computed(String()),
			"keep_bpf_templates":                Computed(Bool()),
			"keep_config":                       Computed(Bool()),
			"label_prefix_file":                 Computed(String()),
			"labels":                            Computed(List(String())),
			"lb":                                Computed(String()),
			"lib_dir":                           Computed(String()),
			"log_drivers":                       Computed(List(String())),
			"log_opt":                           Computed(Map(String())),
			"logstash":                          Computed(Bool()),
			"logstash_agent":                    Computed(String()),
			"logstash_probe_timer":              Computed(Int()),
			"disable_masquerade":                Computed(Bool()),
			"nat46_range":                       Computed(String()),
			"pprof":                             Computed(Bool()),
			"prefilter_device":                  Computed(String()),
			"prometheus_serve_addr":             Computed(String()),
			"restore":                           Computed(Bool()),
			"single_cluster_route":              Computed(Bool()),
			"socket_path":                       Computed(String()),
			"state_dir":                         Computed(String()),
			"trace_payload_len":                 Computed(Int()),
			"tunnel":                            Computed(String()),
			"enable_ipv6":                       Computed(Bool()),
			"enable_ipv4":                       Computed(Bool()),
			"monitor_aggregation":               Computed(String()),
			"bpfct_global_tcp_max":              Computed(Int()),
			"bpfct_global_any_max":              Computed(Int()),
			"preallocate_bpf_maps":              Computed(Bool()),
			"sidecar_istio_proxy_image":         Computed(String()),
			"cluster_name":                      Computed(String()),
			"to_fqdns_dns_reject_response_code": Computed(String()),
			"to_fqdns_enable_poller":            Computed(Bool()),
			"container_runtime_labels":          Computed(String()),
			"ipam":                              Computed(String()),
			"ip_tables_rules_noinstall":         Computed(Bool()),
			"auto_direct_node_routes":           Computed(Bool()),
			"enable_host_reachable_services":    Computed(Bool()),
			"enable_node_port":                  Computed(Bool()),
			"etcd_managed":                      Computed(Bool()),
			"enable_remote_node_identity":       Computed(Ptr(Bool())),
			"hubble":                            Computed(Ptr(Struct(DataSourceHubbleSpec()))),
			"remove_cbr_bridge":                 Computed(Bool()),
			"restart_pods":                      Computed(Bool()),
			"reconfigure_kubelet":               Computed(Bool()),
			"node_init_bootstrap_file":          Computed(String()),
			"cni_bin_path":                      Computed(String()),
		},
	}
}

func ExpandDataSourceCiliumNetworkingSpec(in map[string]interface{}) kops.CiliumNetworkingSpec {
	if in == nil {
		panic("expand CiliumNetworkingSpec failure, in is nil")
	}
	out := kops.CiliumNetworkingSpec{}
	if in, ok := in["version"]; ok && in != nil {
		out.Version = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["access_log"]; ok && in != nil {
		out.AccessLog = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["agent_labels"]; ok && in != nil {
		out.AgentLabels = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["agent_prometheus_port"]; ok && in != nil {
		out.AgentPrometheusPort = func(in interface{}) int { return int(in.(int)) }(in)
	}
	if in, ok := in["allow_localhost"]; ok && in != nil {
		out.AllowLocalhost = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["auto_ipv6_node_routes"]; ok && in != nil {
		out.AutoIpv6NodeRoutes = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["bpf_root"]; ok && in != nil {
		out.BPFRoot = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["container_runtime"]; ok && in != nil {
		out.ContainerRuntime = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["container_runtime_endpoint"]; ok && in != nil {
		out.ContainerRuntimeEndpoint = func(in interface{}) map[string]string {
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
	if in, ok := in["debug"]; ok && in != nil {
		out.Debug = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["debug_verbose"]; ok && in != nil {
		out.DebugVerbose = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["device"]; ok && in != nil {
		out.Device = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["disable_conntrack"]; ok && in != nil {
		out.DisableConntrack = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["disable_ipv4"]; ok && in != nil {
		out.DisableIpv4 = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["disable_k_8s_services"]; ok && in != nil {
		out.DisableK8sServices = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["enable_policy"]; ok && in != nil {
		out.EnablePolicy = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["enable_tracing"]; ok && in != nil {
		out.EnableTracing = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["enable_prometheus_metrics"]; ok && in != nil {
		out.EnablePrometheusMetrics = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["enable_encryption"]; ok && in != nil {
		out.EnableEncryption = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["envoy_log"]; ok && in != nil {
		out.EnvoyLog = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["ipv4_cluster_cidr_mask_size"]; ok && in != nil {
		out.Ipv4ClusterCIDRMaskSize = func(in interface{}) int { return int(in.(int)) }(in)
	}
	if in, ok := in["ipv4_node"]; ok && in != nil {
		out.Ipv4Node = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["ipv4_range"]; ok && in != nil {
		out.Ipv4Range = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["ipv4_service_range"]; ok && in != nil {
		out.Ipv4ServiceRange = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["ipv6_cluster_alloc_cidr"]; ok && in != nil {
		out.Ipv6ClusterAllocCidr = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["ipv6_node"]; ok && in != nil {
		out.Ipv6Node = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["ipv6_range"]; ok && in != nil {
		out.Ipv6Range = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["ipv6_service_range"]; ok && in != nil {
		out.Ipv6ServiceRange = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["k_8s_api_server"]; ok && in != nil {
		out.K8sAPIServer = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["k_8s_kubeconfig_path"]; ok && in != nil {
		out.K8sKubeconfigPath = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["keep_bpf_templates"]; ok && in != nil {
		out.KeepBPFTemplates = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["keep_config"]; ok && in != nil {
		out.KeepConfig = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["label_prefix_file"]; ok && in != nil {
		out.LabelPrefixFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["labels"]; ok && in != nil {
		out.Labels = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["lb"]; ok && in != nil {
		out.LB = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["lib_dir"]; ok && in != nil {
		out.LibDir = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["log_drivers"]; ok && in != nil {
		out.LogDrivers = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["log_opt"]; ok && in != nil {
		out.LogOpt = func(in interface{}) map[string]string {
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
	if in, ok := in["logstash"]; ok && in != nil {
		out.Logstash = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["logstash_agent"]; ok && in != nil {
		out.LogstashAgent = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["logstash_probe_timer"]; ok && in != nil {
		out.LogstashProbeTimer = func(in interface{}) uint32 { return uint32(in.(int)) }(in)
	}
	if in, ok := in["disable_masquerade"]; ok && in != nil {
		out.DisableMasquerade = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["nat46_range"]; ok && in != nil {
		out.Nat46Range = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["pprof"]; ok && in != nil {
		out.Pprof = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["prefilter_device"]; ok && in != nil {
		out.PrefilterDevice = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["prometheus_serve_addr"]; ok && in != nil {
		out.PrometheusServeAddr = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["restore"]; ok && in != nil {
		out.Restore = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["single_cluster_route"]; ok && in != nil {
		out.SingleClusterRoute = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["socket_path"]; ok && in != nil {
		out.SocketPath = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["state_dir"]; ok && in != nil {
		out.StateDir = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["trace_payload_len"]; ok && in != nil {
		out.TracePayloadLen = func(in interface{}) int { return int(in.(int)) }(in)
	}
	if in, ok := in["tunnel"]; ok && in != nil {
		out.Tunnel = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["enable_ipv6"]; ok && in != nil {
		out.EnableIpv6 = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["enable_ipv4"]; ok && in != nil {
		out.EnableIpv4 = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["monitor_aggregation"]; ok && in != nil {
		out.MonitorAggregation = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["bpfct_global_tcp_max"]; ok && in != nil {
		out.BPFCTGlobalTCPMax = func(in interface{}) int { return int(in.(int)) }(in)
	}
	if in, ok := in["bpfct_global_any_max"]; ok && in != nil {
		out.BPFCTGlobalAnyMax = func(in interface{}) int { return int(in.(int)) }(in)
	}
	if in, ok := in["preallocate_bpf_maps"]; ok && in != nil {
		out.PreallocateBPFMaps = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["sidecar_istio_proxy_image"]; ok && in != nil {
		out.SidecarIstioProxyImage = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cluster_name"]; ok && in != nil {
		out.ClusterName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["to_fqdns_dns_reject_response_code"]; ok && in != nil {
		out.ToFqdnsDNSRejectResponseCode = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["to_fqdns_enable_poller"]; ok && in != nil {
		out.ToFqdnsEnablePoller = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["container_runtime_labels"]; ok && in != nil {
		out.ContainerRuntimeLabels = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["ipam"]; ok && in != nil {
		out.Ipam = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["ip_tables_rules_noinstall"]; ok && in != nil {
		out.IPTablesRulesNoinstall = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["auto_direct_node_routes"]; ok && in != nil {
		out.AutoDirectNodeRoutes = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["enable_host_reachable_services"]; ok && in != nil {
		out.EnableHostReachableServices = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["enable_node_port"]; ok && in != nil {
		out.EnableNodePort = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["etcd_managed"]; ok && in != nil {
		out.EtcdManaged = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["enable_remote_node_identity"]; ok && in != nil {
		out.EnableRemoteNodeIdentity = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["hubble"]; ok && in != nil {
		out.Hubble = func(in interface{}) *kops.HubbleSpec {
			if in == nil {
				return nil
			}
			return func(in kops.HubbleSpec) *kops.HubbleSpec { return &in }(func(in interface{}) kops.HubbleSpec {
				if in == nil {
					return kops.HubbleSpec{}
				}
				return ExpandDataSourceHubbleSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["remove_cbr_bridge"]; ok && in != nil {
		out.RemoveCbrBridge = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["restart_pods"]; ok && in != nil {
		out.RestartPods = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["reconfigure_kubelet"]; ok && in != nil {
		out.ReconfigureKubelet = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["node_init_bootstrap_file"]; ok && in != nil {
		out.NodeInitBootstrapFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cni_bin_path"]; ok && in != nil {
		out.CniBinPath = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}
