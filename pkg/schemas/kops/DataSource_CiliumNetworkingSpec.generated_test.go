package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceCiliumNetworkingSpec(t *testing.T) {
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
					"version":                            "",
					"memory_request":                     nil,
					"cpu_request":                        nil,
					"agent_prometheus_port":              0,
					"chaining_mode":                      "",
					"debug":                              false,
					"disable_endpoint_crd":               false,
					"enable_policy":                      "",
					"enable_l7_proxy":                    nil,
					"enable_bpf_masquerade":              nil,
					"enable_endpoint_health_checking":    nil,
					"enable_prometheus_metrics":          false,
					"enable_encryption":                  false,
					"encryption_type":                    "",
					"identity_allocation_mode":           "",
					"identity_change_grace_period":       "",
					"masquerade":                         nil,
					"agent_pod_annotations":              func() map[string]interface{} { return nil }(),
					"tunnel":                             "",
					"monitor_aggregation":                "",
					"bpfct_global_tcp_max":               0,
					"bpfct_global_any_max":               0,
					"bpflb_algorithm":                    "",
					"bpflb_maglev_table_size":            "",
					"bpfnat_global_max":                  0,
					"bpf_neigh_global_max":               0,
					"bpf_policy_map_max":                 0,
					"bpflb_map_max":                      0,
					"bpflb_sock_host_ns_only":            false,
					"preallocate_bpf_maps":               false,
					"sidecar_istio_proxy_image":          "",
					"cluster_name":                       "",
					"to_fqd_ns_dns_reject_response_code": "",
					"to_fqd_ns_enable_poller":            false,
					"ip_am":                              "",
					"install_iptables_rules":             nil,
					"auto_direct_node_routes":            false,
					"enable_host_reachable_services":     false,
					"enable_node_port":                   false,
					"etcd_managed":                       false,
					"enable_remote_node_identity":        nil,
					"hubble":                             nil,
					"disable_cnp_status_updates":         nil,
					"enable_service_topology":            false,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceCiliumNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceCiliumNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceCiliumNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"version":                            "",
		"memory_request":                     nil,
		"cpu_request":                        nil,
		"agent_prometheus_port":              0,
		"chaining_mode":                      "",
		"debug":                              false,
		"disable_endpoint_crd":               false,
		"enable_policy":                      "",
		"enable_l7_proxy":                    nil,
		"enable_bpf_masquerade":              nil,
		"enable_endpoint_health_checking":    nil,
		"enable_prometheus_metrics":          false,
		"enable_encryption":                  false,
		"encryption_type":                    "",
		"identity_allocation_mode":           "",
		"identity_change_grace_period":       "",
		"masquerade":                         nil,
		"agent_pod_annotations":              func() map[string]interface{} { return nil }(),
		"tunnel":                             "",
		"monitor_aggregation":                "",
		"bpfct_global_tcp_max":               0,
		"bpfct_global_any_max":               0,
		"bpflb_algorithm":                    "",
		"bpflb_maglev_table_size":            "",
		"bpfnat_global_max":                  0,
		"bpf_neigh_global_max":               0,
		"bpf_policy_map_max":                 0,
		"bpflb_map_max":                      0,
		"bpflb_sock_host_ns_only":            false,
		"preallocate_bpf_maps":               false,
		"sidecar_istio_proxy_image":          "",
		"cluster_name":                       "",
		"to_fqd_ns_dns_reject_response_code": "",
		"to_fqd_ns_enable_poller":            false,
		"ip_am":                              "",
		"install_iptables_rules":             nil,
		"auto_direct_node_routes":            false,
		"enable_host_reachable_services":     false,
		"enable_node_port":                   false,
		"etcd_managed":                       false,
		"enable_remote_node_identity":        nil,
		"hubble":                             nil,
		"disable_cnp_status_updates":         nil,
		"enable_service_topology":            false,
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
			name: "CpuRequest - default",
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
			name: "ChainingMode - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ChainingMode = ""
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
			name: "DisableEndpointCRD - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DisableEndpointCRD = false
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
			name: "EnableL7Proxy - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableL7Proxy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableBPFMasquerade - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableBPFMasquerade = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableEndpointHealthChecking - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableEndpointHealthChecking = nil
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
			name: "EncryptionType - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EncryptionType = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdentityAllocationMode - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.IdentityAllocationMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdentityChangeGracePeriod - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.IdentityChangeGracePeriod = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Masquerade - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Masquerade = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AgentPodAnnotations - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AgentPodAnnotations = nil
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
			name: "BPFLBAlgorithm - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFLBAlgorithm = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLBMaglevTableSize - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFLBMaglevTableSize = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFNATGlobalMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFNATGlobalMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFNeighGlobalMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFNeighGlobalMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFPolicyMapMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFPolicyMapMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLBMapMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFLBMapMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLBSockHostNSOnly - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFLBSockHostNSOnly = false
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
			name: "ToFQDNsDnsRejectResponseCode - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ToFQDNsDNSRejectResponseCode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ToFQDNsEnablePoller - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ToFQDNsEnablePoller = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpAM - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.IPAM = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InstallIptablesRules - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.InstallIptablesRules = nil
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
			name: "DisableCNPStatusUpdates - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DisableCNPStatusUpdates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableServiceTopology - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableServiceTopology = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceCiliumNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceCiliumNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceCiliumNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"version":                            "",
		"memory_request":                     nil,
		"cpu_request":                        nil,
		"agent_prometheus_port":              0,
		"chaining_mode":                      "",
		"debug":                              false,
		"disable_endpoint_crd":               false,
		"enable_policy":                      "",
		"enable_l7_proxy":                    nil,
		"enable_bpf_masquerade":              nil,
		"enable_endpoint_health_checking":    nil,
		"enable_prometheus_metrics":          false,
		"enable_encryption":                  false,
		"encryption_type":                    "",
		"identity_allocation_mode":           "",
		"identity_change_grace_period":       "",
		"masquerade":                         nil,
		"agent_pod_annotations":              func() map[string]interface{} { return nil }(),
		"tunnel":                             "",
		"monitor_aggregation":                "",
		"bpfct_global_tcp_max":               0,
		"bpfct_global_any_max":               0,
		"bpflb_algorithm":                    "",
		"bpflb_maglev_table_size":            "",
		"bpfnat_global_max":                  0,
		"bpf_neigh_global_max":               0,
		"bpf_policy_map_max":                 0,
		"bpflb_map_max":                      0,
		"bpflb_sock_host_ns_only":            false,
		"preallocate_bpf_maps":               false,
		"sidecar_istio_proxy_image":          "",
		"cluster_name":                       "",
		"to_fqd_ns_dns_reject_response_code": "",
		"to_fqd_ns_enable_poller":            false,
		"ip_am":                              "",
		"install_iptables_rules":             nil,
		"auto_direct_node_routes":            false,
		"enable_host_reachable_services":     false,
		"enable_node_port":                   false,
		"etcd_managed":                       false,
		"enable_remote_node_identity":        nil,
		"hubble":                             nil,
		"disable_cnp_status_updates":         nil,
		"enable_service_topology":            false,
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
			name: "CpuRequest - default",
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
			name: "ChainingMode - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ChainingMode = ""
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
			name: "DisableEndpointCRD - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DisableEndpointCRD = false
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
			name: "EnableL7Proxy - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableL7Proxy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableBPFMasquerade - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableBPFMasquerade = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableEndpointHealthChecking - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableEndpointHealthChecking = nil
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
			name: "EncryptionType - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EncryptionType = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdentityAllocationMode - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.IdentityAllocationMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdentityChangeGracePeriod - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.IdentityChangeGracePeriod = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Masquerade - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Masquerade = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AgentPodAnnotations - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AgentPodAnnotations = nil
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
			name: "BPFLBAlgorithm - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFLBAlgorithm = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLBMaglevTableSize - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFLBMaglevTableSize = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFNATGlobalMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFNATGlobalMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFNeighGlobalMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFNeighGlobalMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFPolicyMapMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFPolicyMapMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLBMapMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFLBMapMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLBSockHostNSOnly - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFLBSockHostNSOnly = false
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
			name: "ToFQDNsDnsRejectResponseCode - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ToFQDNsDNSRejectResponseCode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ToFQDNsEnablePoller - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ToFQDNsEnablePoller = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IpAM - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.IPAM = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InstallIptablesRules - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.InstallIptablesRules = nil
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
			name: "DisableCNPStatusUpdates - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DisableCNPStatusUpdates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableServiceTopology - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableServiceTopology = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceCiliumNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceCiliumNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
