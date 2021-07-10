package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceCiliumNetworkingSpec(t *testing.T) {
	_default := kops.CiliumNetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.CiliumNetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"version":                           "",
					"memory_request":                    nil,
					"cpu_request":                       nil,
					"access_log":                        "",
					"agent_labels":                      func() []interface{} { return nil }(),
					"agent_prometheus_port":             0,
					"allow_localhost":                   "",
					"auto_ipv6_node_routes":             false,
					"bpf_root":                          "",
					"container_runtime":                 func() []interface{} { return nil }(),
					"container_runtime_endpoint":        func() map[string]interface{} { return nil }(),
					"debug":                             false,
					"debug_verbose":                     func() []interface{} { return nil }(),
					"device":                            "",
					"disable_conntrack":                 false,
					"disable_ipv4":                      false,
					"disable_k_8s_services":             false,
					"enable_policy":                     "",
					"enable_tracing":                    false,
					"enable_prometheus_metrics":         false,
					"enable_encryption":                 false,
					"envoy_log":                         "",
					"ipv4_cluster_cidr_mask_size":       0,
					"ipv4_node":                         "",
					"ipv4_range":                        "",
					"ipv4_service_range":                "",
					"ipv6_cluster_alloc_cidr":           "",
					"ipv6_node":                         "",
					"ipv6_range":                        "",
					"ipv6_service_range":                "",
					"k_8s_api_server":                   "",
					"k_8s_kubeconfig_path":              "",
					"keep_bpf_templates":                false,
					"keep_config":                       false,
					"label_prefix_file":                 "",
					"labels":                            func() []interface{} { return nil }(),
					"lb":                                "",
					"lib_dir":                           "",
					"log_drivers":                       func() []interface{} { return nil }(),
					"log_opt":                           func() map[string]interface{} { return nil }(),
					"logstash":                          false,
					"logstash_agent":                    "",
					"logstash_probe_timer":              0,
					"disable_masquerade":                false,
					"nat46_range":                       "",
					"pprof":                             false,
					"prefilter_device":                  "",
					"prometheus_serve_addr":             "",
					"restore":                           false,
					"single_cluster_route":              false,
					"socket_path":                       "",
					"state_dir":                         "",
					"trace_payload_len":                 0,
					"tunnel":                            "",
					"enable_ipv6":                       false,
					"enable_ipv4":                       false,
					"monitor_aggregation":               "",
					"bpfct_global_tcp_max":              0,
					"bpfct_global_any_max":              0,
					"preallocate_bpf_maps":              false,
					"sidecar_istio_proxy_image":         "",
					"cluster_name":                      "",
					"to_fqdns_dns_reject_response_code": "",
					"to_fqdns_enable_poller":            false,
					"container_runtime_labels":          "",
					"ipam":                              "",
					"ip_tables_rules_noinstall":         false,
					"auto_direct_node_routes":           false,
					"enable_host_reachable_services":    false,
					"enable_node_port":                  false,
					"etcd_managed":                      false,
					"enable_remote_node_identity":       nil,
					"hubble":                            nil,
					"remove_cbr_bridge":                 false,
					"restart_pods":                      false,
					"reconfigure_kubelet":               false,
					"node_init_bootstrap_file":          "",
					"cni_bin_path":                      "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceCiliumNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceCiliumNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCiliumNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"version":                           "",
		"memory_request":                    nil,
		"cpu_request":                       nil,
		"access_log":                        "",
		"agent_labels":                      func() []interface{} { return nil }(),
		"agent_prometheus_port":             0,
		"allow_localhost":                   "",
		"auto_ipv6_node_routes":             false,
		"bpf_root":                          "",
		"container_runtime":                 func() []interface{} { return nil }(),
		"container_runtime_endpoint":        func() map[string]interface{} { return nil }(),
		"debug":                             false,
		"debug_verbose":                     func() []interface{} { return nil }(),
		"device":                            "",
		"disable_conntrack":                 false,
		"disable_ipv4":                      false,
		"disable_k_8s_services":             false,
		"enable_policy":                     "",
		"enable_tracing":                    false,
		"enable_prometheus_metrics":         false,
		"enable_encryption":                 false,
		"envoy_log":                         "",
		"ipv4_cluster_cidr_mask_size":       0,
		"ipv4_node":                         "",
		"ipv4_range":                        "",
		"ipv4_service_range":                "",
		"ipv6_cluster_alloc_cidr":           "",
		"ipv6_node":                         "",
		"ipv6_range":                        "",
		"ipv6_service_range":                "",
		"k_8s_api_server":                   "",
		"k_8s_kubeconfig_path":              "",
		"keep_bpf_templates":                false,
		"keep_config":                       false,
		"label_prefix_file":                 "",
		"labels":                            func() []interface{} { return nil }(),
		"lb":                                "",
		"lib_dir":                           "",
		"log_drivers":                       func() []interface{} { return nil }(),
		"log_opt":                           func() map[string]interface{} { return nil }(),
		"logstash":                          false,
		"logstash_agent":                    "",
		"logstash_probe_timer":              0,
		"disable_masquerade":                false,
		"nat46_range":                       "",
		"pprof":                             false,
		"prefilter_device":                  "",
		"prometheus_serve_addr":             "",
		"restore":                           false,
		"single_cluster_route":              false,
		"socket_path":                       "",
		"state_dir":                         "",
		"trace_payload_len":                 0,
		"tunnel":                            "",
		"enable_ipv6":                       false,
		"enable_ipv4":                       false,
		"monitor_aggregation":               "",
		"bpfct_global_tcp_max":              0,
		"bpfct_global_any_max":              0,
		"preallocate_bpf_maps":              false,
		"sidecar_istio_proxy_image":         "",
		"cluster_name":                      "",
		"to_fqdns_dns_reject_response_code": "",
		"to_fqdns_enable_poller":            false,
		"container_runtime_labels":          "",
		"ipam":                              "",
		"ip_tables_rules_noinstall":         false,
		"auto_direct_node_routes":           false,
		"enable_host_reachable_services":    false,
		"enable_node_port":                  false,
		"etcd_managed":                      false,
		"enable_remote_node_identity":       nil,
		"hubble":                            nil,
		"remove_cbr_bridge":                 false,
		"restart_pods":                      false,
		"reconfigure_kubelet":               false,
		"node_init_bootstrap_file":          "",
		"cni_bin_path":                      "",
	}
	type args struct {
		in kops.CiliumNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CiliumNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Version = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CPURequest - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AccessLog - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AccessLog = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AgentLabels - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AgentLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AgentPrometheusPort - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AgentPrometheusPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowLocalhost - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AllowLocalhost = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AutoIpv6NodeRoutes - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AutoIpv6NodeRoutes = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFRoot - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFRoot = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerRuntime - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ContainerRuntime = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerRuntimeEndpoint - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ContainerRuntimeEndpoint = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Debug - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Debug = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DebugVerbose - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DebugVerbose = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Device - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Device = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableConntrack - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DisableConntrack = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableIpv4 - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DisableIpv4 = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableK8sServices - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DisableK8sServices = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnablePolicy - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnablePolicy = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableTracing - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableTracing = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnablePrometheusMetrics - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnablePrometheusMetrics = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableEncryption - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableEncryption = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnvoyLog - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnvoyLog = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv4ClusterCidrMaskSize - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipv4ClusterCIDRMaskSize = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv4Node - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipv4Node = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv4Range - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipv4Range = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv4ServiceRange - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipv4ServiceRange = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6ClusterAllocCidr - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipv6ClusterAllocCidr = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6Node - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipv6Node = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6Range - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipv6Range = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6ServiceRange - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipv6ServiceRange = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "K8sApiServer - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.K8sAPIServer = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "K8sKubeconfigPath - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.K8sKubeconfigPath = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KeepBPFTemplates - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.KeepBPFTemplates = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KeepConfig - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.KeepConfig = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LabelPrefixFile - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.LabelPrefixFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Labels - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Labels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LB - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.LB = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LibDir - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.LibDir = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogDrivers - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.LogDrivers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogOpt - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.LogOpt = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Logstash - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Logstash = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogstashAgent - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.LogstashAgent = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogstashProbeTimer - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.LogstashProbeTimer = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableMasquerade - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DisableMasquerade = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Nat46Range - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Nat46Range = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Pprof - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Pprof = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrefilterDevice - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.PrefilterDevice = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusServeAddr - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.PrometheusServeAddr = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Restore - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Restore = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SingleClusterRoute - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.SingleClusterRoute = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SocketPath - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.SocketPath = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "StateDir - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.StateDir = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TracePayloadLen - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.TracePayloadLen = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Tunnel - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Tunnel = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableIpv6 - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableIpv6 = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableIpv4 - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableIpv4 = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MonitorAggregation - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.MonitorAggregation = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFCTGlobalTCPMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFCTGlobalTCPMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFCTGlobalAnyMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFCTGlobalAnyMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PreallocateBPFMaps - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.PreallocateBPFMaps = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SidecarIstioProxyImage - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.SidecarIstioProxyImage = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ToFqdnsDnsRejectResponseCode - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ToFqdnsDNSRejectResponseCode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ToFqdnsEnablePoller - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ToFqdnsEnablePoller = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerRuntimeLabels - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ContainerRuntimeLabels = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipam - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipam = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpTablesRulesNoinstall - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.IPTablesRulesNoinstall = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AutoDirectNodeRoutes - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AutoDirectNodeRoutes = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableHostReachableServices - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableHostReachableServices = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableNodePort - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableNodePort = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdManaged - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EtcdManaged = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableRemoteNodeIdentity - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableRemoteNodeIdentity = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Hubble - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Hubble = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RemoveCbrBridge - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.RemoveCbrBridge = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RestartPods - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.RestartPods = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ReconfigureKubelet - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ReconfigureKubelet = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeInitBootstrapFile - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.NodeInitBootstrapFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CniBinPath - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.CniBinPath = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceCiliumNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCiliumNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCiliumNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"version":                           "",
		"memory_request":                    nil,
		"cpu_request":                       nil,
		"access_log":                        "",
		"agent_labels":                      func() []interface{} { return nil }(),
		"agent_prometheus_port":             0,
		"allow_localhost":                   "",
		"auto_ipv6_node_routes":             false,
		"bpf_root":                          "",
		"container_runtime":                 func() []interface{} { return nil }(),
		"container_runtime_endpoint":        func() map[string]interface{} { return nil }(),
		"debug":                             false,
		"debug_verbose":                     func() []interface{} { return nil }(),
		"device":                            "",
		"disable_conntrack":                 false,
		"disable_ipv4":                      false,
		"disable_k_8s_services":             false,
		"enable_policy":                     "",
		"enable_tracing":                    false,
		"enable_prometheus_metrics":         false,
		"enable_encryption":                 false,
		"envoy_log":                         "",
		"ipv4_cluster_cidr_mask_size":       0,
		"ipv4_node":                         "",
		"ipv4_range":                        "",
		"ipv4_service_range":                "",
		"ipv6_cluster_alloc_cidr":           "",
		"ipv6_node":                         "",
		"ipv6_range":                        "",
		"ipv6_service_range":                "",
		"k_8s_api_server":                   "",
		"k_8s_kubeconfig_path":              "",
		"keep_bpf_templates":                false,
		"keep_config":                       false,
		"label_prefix_file":                 "",
		"labels":                            func() []interface{} { return nil }(),
		"lb":                                "",
		"lib_dir":                           "",
		"log_drivers":                       func() []interface{} { return nil }(),
		"log_opt":                           func() map[string]interface{} { return nil }(),
		"logstash":                          false,
		"logstash_agent":                    "",
		"logstash_probe_timer":              0,
		"disable_masquerade":                false,
		"nat46_range":                       "",
		"pprof":                             false,
		"prefilter_device":                  "",
		"prometheus_serve_addr":             "",
		"restore":                           false,
		"single_cluster_route":              false,
		"socket_path":                       "",
		"state_dir":                         "",
		"trace_payload_len":                 0,
		"tunnel":                            "",
		"enable_ipv6":                       false,
		"enable_ipv4":                       false,
		"monitor_aggregation":               "",
		"bpfct_global_tcp_max":              0,
		"bpfct_global_any_max":              0,
		"preallocate_bpf_maps":              false,
		"sidecar_istio_proxy_image":         "",
		"cluster_name":                      "",
		"to_fqdns_dns_reject_response_code": "",
		"to_fqdns_enable_poller":            false,
		"container_runtime_labels":          "",
		"ipam":                              "",
		"ip_tables_rules_noinstall":         false,
		"auto_direct_node_routes":           false,
		"enable_host_reachable_services":    false,
		"enable_node_port":                  false,
		"etcd_managed":                      false,
		"enable_remote_node_identity":       nil,
		"hubble":                            nil,
		"remove_cbr_bridge":                 false,
		"restart_pods":                      false,
		"reconfigure_kubelet":               false,
		"node_init_bootstrap_file":          "",
		"cni_bin_path":                      "",
	}
	type args struct {
		in kops.CiliumNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CiliumNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Version = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CPURequest - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AccessLog - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AccessLog = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AgentLabels - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AgentLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AgentPrometheusPort - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AgentPrometheusPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowLocalhost - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AllowLocalhost = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AutoIpv6NodeRoutes - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AutoIpv6NodeRoutes = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFRoot - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFRoot = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerRuntime - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ContainerRuntime = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerRuntimeEndpoint - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ContainerRuntimeEndpoint = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Debug - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Debug = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DebugVerbose - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DebugVerbose = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Device - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Device = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableConntrack - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DisableConntrack = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableIpv4 - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DisableIpv4 = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableK8sServices - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DisableK8sServices = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnablePolicy - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnablePolicy = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableTracing - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableTracing = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnablePrometheusMetrics - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnablePrometheusMetrics = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableEncryption - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableEncryption = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnvoyLog - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnvoyLog = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv4ClusterCidrMaskSize - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipv4ClusterCIDRMaskSize = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv4Node - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipv4Node = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv4Range - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipv4Range = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv4ServiceRange - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipv4ServiceRange = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6ClusterAllocCidr - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipv6ClusterAllocCidr = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6Node - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipv6Node = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6Range - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipv6Range = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6ServiceRange - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipv6ServiceRange = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "K8sApiServer - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.K8sAPIServer = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "K8sKubeconfigPath - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.K8sKubeconfigPath = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KeepBPFTemplates - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.KeepBPFTemplates = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KeepConfig - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.KeepConfig = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LabelPrefixFile - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.LabelPrefixFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Labels - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Labels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LB - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.LB = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LibDir - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.LibDir = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogDrivers - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.LogDrivers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogOpt - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.LogOpt = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Logstash - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Logstash = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogstashAgent - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.LogstashAgent = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogstashProbeTimer - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.LogstashProbeTimer = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableMasquerade - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DisableMasquerade = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Nat46Range - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Nat46Range = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Pprof - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Pprof = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrefilterDevice - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.PrefilterDevice = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PrometheusServeAddr - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.PrometheusServeAddr = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Restore - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Restore = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SingleClusterRoute - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.SingleClusterRoute = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SocketPath - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.SocketPath = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "StateDir - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.StateDir = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TracePayloadLen - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.TracePayloadLen = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Tunnel - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Tunnel = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableIpv6 - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableIpv6 = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableIpv4 - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableIpv4 = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MonitorAggregation - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.MonitorAggregation = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFCTGlobalTCPMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFCTGlobalTCPMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFCTGlobalAnyMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFCTGlobalAnyMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PreallocateBPFMaps - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.PreallocateBPFMaps = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SidecarIstioProxyImage - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.SidecarIstioProxyImage = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ToFqdnsDnsRejectResponseCode - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ToFqdnsDNSRejectResponseCode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ToFqdnsEnablePoller - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ToFqdnsEnablePoller = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerRuntimeLabels - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ContainerRuntimeLabels = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipam - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ipam = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpTablesRulesNoinstall - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.IPTablesRulesNoinstall = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AutoDirectNodeRoutes - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AutoDirectNodeRoutes = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableHostReachableServices - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableHostReachableServices = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableNodePort - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableNodePort = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdManaged - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EtcdManaged = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableRemoteNodeIdentity - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableRemoteNodeIdentity = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Hubble - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Hubble = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RemoveCbrBridge - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.RemoveCbrBridge = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RestartPods - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.RestartPods = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ReconfigureKubelet - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ReconfigureKubelet = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeInitBootstrapFile - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.NodeInitBootstrapFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CniBinPath - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.CniBinPath = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceCiliumNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCiliumNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
