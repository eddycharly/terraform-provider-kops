package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceClusterSpec(t *testing.T) {
	_default := kops.ClusterSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ClusterSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"channel":                           "",
					"addons":                            func() []interface{} { return nil }(),
					"config_base":                       "",
					"cloud_provider":                    func() []interface{} { return []interface{}{FlattenResourceCloudProviderSpec(kops.CloudProviderSpec{})} }(),
					"container_runtime":                 "",
					"kubernetes_version":                "",
					"secret_store":                      "",
					"key_store":                         "",
					"config_store":                      "",
					"dns_zone":                          "",
					"cluster_dns_domain":                "",
					"ssh_access":                        func() []interface{} { return nil }(),
					"node_port_access":                  func() []interface{} { return nil }(),
					"ssh_key_name":                      nil,
					"update_policy":                     nil,
					"external_policies":                 nil,
					"additional_policies":               nil,
					"file_assets":                       func() []interface{} { return nil }(),
					"etcd_cluster":                      func() []interface{} { return nil }(),
					"containerd":                        nil,
					"docker":                            nil,
					"kube_dns":                          nil,
					"kube_api_server":                   nil,
					"kube_controller_manager":           nil,
					"external_cloud_controller_manager": nil,
					"kube_scheduler":                    nil,
					"kube_proxy":                        nil,
					"kubelet":                           nil,
					"control_plane_kubelet":             nil,
					"cloud_config":                      nil,
					"external_dns":                      nil,
					"ntp":                               nil,
					"node_problem_detector":             nil,
					"metrics_server":                    nil,
					"cert_manager":                      nil,
					"networking":                        func() []interface{} { return []interface{}{FlattenResourceNetworkingSpec(kops.NetworkingSpec{})} }(),
					"api":                               func() []interface{} { return []interface{}{FlattenResourceAPISpec(kops.APISpec{})} }(),
					"authentication":                    nil,
					"authorization":                     nil,
					"node_authorization":                nil,
					"cloud_labels":                      func() map[string]interface{} { return nil }(),
					"hooks":                             func() []interface{} { return nil }(),
					"assets":                            nil,
					"iam":                               nil,
					"encryption_config":                 nil,
					"use_host_certificates":             nil,
					"sysctl_parameters":                 func() []interface{} { return nil }(),
					"rolling_update":                    nil,
					"cluster_autoscaler":                nil,
					"service_account_issuer_discovery":  nil,
					"snapshot_controller":               nil,
					"karpenter":                         nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceClusterSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceClusterSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceClusterSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"channel":                           "",
		"addons":                            func() []interface{} { return nil }(),
		"config_base":                       "",
		"cloud_provider":                    func() []interface{} { return []interface{}{FlattenResourceCloudProviderSpec(kops.CloudProviderSpec{})} }(),
		"container_runtime":                 "",
		"kubernetes_version":                "",
		"secret_store":                      "",
		"key_store":                         "",
		"config_store":                      "",
		"dns_zone":                          "",
		"cluster_dns_domain":                "",
		"ssh_access":                        func() []interface{} { return nil }(),
		"node_port_access":                  func() []interface{} { return nil }(),
		"ssh_key_name":                      nil,
		"update_policy":                     nil,
		"external_policies":                 nil,
		"additional_policies":               nil,
		"file_assets":                       func() []interface{} { return nil }(),
		"etcd_cluster":                      func() []interface{} { return nil }(),
		"containerd":                        nil,
		"docker":                            nil,
		"kube_dns":                          nil,
		"kube_api_server":                   nil,
		"kube_controller_manager":           nil,
		"external_cloud_controller_manager": nil,
		"kube_scheduler":                    nil,
		"kube_proxy":                        nil,
		"kubelet":                           nil,
		"control_plane_kubelet":             nil,
		"cloud_config":                      nil,
		"external_dns":                      nil,
		"ntp":                               nil,
		"node_problem_detector":             nil,
		"metrics_server":                    nil,
		"cert_manager":                      nil,
		"networking":                        func() []interface{} { return []interface{}{FlattenResourceNetworkingSpec(kops.NetworkingSpec{})} }(),
		"api":                               func() []interface{} { return []interface{}{FlattenResourceAPISpec(kops.APISpec{})} }(),
		"authentication":                    nil,
		"authorization":                     nil,
		"node_authorization":                nil,
		"cloud_labels":                      func() map[string]interface{} { return nil }(),
		"hooks":                             func() []interface{} { return nil }(),
		"assets":                            nil,
		"iam":                               nil,
		"encryption_config":                 nil,
		"use_host_certificates":             nil,
		"sysctl_parameters":                 func() []interface{} { return nil }(),
		"rolling_update":                    nil,
		"cluster_autoscaler":                nil,
		"service_account_issuer_discovery":  nil,
		"snapshot_controller":               nil,
		"karpenter":                         nil,
	}
	type args struct {
		in kops.ClusterSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ClusterSpec{},
			},
			want: _default,
		},
		{
			name: "Channel - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Channel = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Addons - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Addons = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigBase - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ConfigBase = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudProvider - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.CloudProvider = kops.CloudProviderSpec{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerRuntime - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ContainerRuntime = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubernetesVersion - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.KubernetesVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecretStore - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.SecretStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KeyStore - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.KeyStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigStore - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ConfigStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DnsZone - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.DNSZone = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterDnsDomain - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ClusterDNSDomain = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SshAccess - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.SSHAccess = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodePortAccess - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.NodePortAccess = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SshKeyName - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.SSHKeyName = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UpdatePolicy - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.UpdatePolicy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalPolicies - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ExternalPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalPolicies - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.AdditionalPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FileAssets - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.FileAssets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdCluster - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.EtcdClusters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Containerd - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Containerd = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Docker - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Docker = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeDns - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.KubeDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeApiServer - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.KubeAPIServer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeControllerManager - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.KubeControllerManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalCloudControllerManager - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ExternalCloudControllerManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeScheduler - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.KubeScheduler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeProxy - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.KubeProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Kubelet - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Kubelet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ControlPlaneKubelet - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ControlPlaneKubelet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudConfig - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.CloudConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalDns - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ExternalDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NTP - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.NTP = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeProblemDetector - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.NodeProblemDetector = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MetricsServer - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.MetricsServer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CertManager - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.CertManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Networking - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Networking = kops.NetworkingSpec{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Api - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.API = kops.APISpec{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Authentication - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Authentication = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Authorization - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Authorization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeAuthorization - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.NodeAuthorization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudLabels - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.CloudLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Hooks - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Hooks = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Assets - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Assets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IAM - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.IAM = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EncryptionConfig - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.EncryptionConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseHostCertificates - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.UseHostCertificates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SysctlParameters - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.SysctlParameters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RollingUpdate - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.RollingUpdate = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterAutoscaler - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ClusterAutoscaler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountIssuerDiscovery - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ServiceAccountIssuerDiscovery = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SnapshotController - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.SnapshotController = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Karpenter - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Karpenter = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceClusterSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceClusterSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceClusterSpec(t *testing.T) {
	_default := map[string]interface{}{
		"channel":                           "",
		"addons":                            func() []interface{} { return nil }(),
		"config_base":                       "",
		"cloud_provider":                    func() []interface{} { return []interface{}{FlattenResourceCloudProviderSpec(kops.CloudProviderSpec{})} }(),
		"container_runtime":                 "",
		"kubernetes_version":                "",
		"secret_store":                      "",
		"key_store":                         "",
		"config_store":                      "",
		"dns_zone":                          "",
		"cluster_dns_domain":                "",
		"ssh_access":                        func() []interface{} { return nil }(),
		"node_port_access":                  func() []interface{} { return nil }(),
		"ssh_key_name":                      nil,
		"update_policy":                     nil,
		"external_policies":                 nil,
		"additional_policies":               nil,
		"file_assets":                       func() []interface{} { return nil }(),
		"etcd_cluster":                      func() []interface{} { return nil }(),
		"containerd":                        nil,
		"docker":                            nil,
		"kube_dns":                          nil,
		"kube_api_server":                   nil,
		"kube_controller_manager":           nil,
		"external_cloud_controller_manager": nil,
		"kube_scheduler":                    nil,
		"kube_proxy":                        nil,
		"kubelet":                           nil,
		"control_plane_kubelet":             nil,
		"cloud_config":                      nil,
		"external_dns":                      nil,
		"ntp":                               nil,
		"node_problem_detector":             nil,
		"metrics_server":                    nil,
		"cert_manager":                      nil,
		"networking":                        func() []interface{} { return []interface{}{FlattenResourceNetworkingSpec(kops.NetworkingSpec{})} }(),
		"api":                               func() []interface{} { return []interface{}{FlattenResourceAPISpec(kops.APISpec{})} }(),
		"authentication":                    nil,
		"authorization":                     nil,
		"node_authorization":                nil,
		"cloud_labels":                      func() map[string]interface{} { return nil }(),
		"hooks":                             func() []interface{} { return nil }(),
		"assets":                            nil,
		"iam":                               nil,
		"encryption_config":                 nil,
		"use_host_certificates":             nil,
		"sysctl_parameters":                 func() []interface{} { return nil }(),
		"rolling_update":                    nil,
		"cluster_autoscaler":                nil,
		"service_account_issuer_discovery":  nil,
		"snapshot_controller":               nil,
		"karpenter":                         nil,
	}
	type args struct {
		in kops.ClusterSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ClusterSpec{},
			},
			want: _default,
		},
		{
			name: "Channel - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Channel = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Addons - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Addons = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigBase - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ConfigBase = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudProvider - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.CloudProvider = kops.CloudProviderSpec{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerRuntime - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ContainerRuntime = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubernetesVersion - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.KubernetesVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecretStore - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.SecretStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KeyStore - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.KeyStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigStore - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ConfigStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DnsZone - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.DNSZone = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterDnsDomain - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ClusterDNSDomain = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SshAccess - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.SSHAccess = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodePortAccess - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.NodePortAccess = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SshKeyName - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.SSHKeyName = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UpdatePolicy - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.UpdatePolicy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalPolicies - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ExternalPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalPolicies - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.AdditionalPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FileAssets - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.FileAssets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdCluster - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.EtcdClusters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Containerd - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Containerd = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Docker - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Docker = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeDns - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.KubeDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeApiServer - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.KubeAPIServer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeControllerManager - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.KubeControllerManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalCloudControllerManager - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ExternalCloudControllerManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeScheduler - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.KubeScheduler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeProxy - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.KubeProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Kubelet - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Kubelet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ControlPlaneKubelet - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ControlPlaneKubelet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudConfig - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.CloudConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalDns - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ExternalDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NTP - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.NTP = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeProblemDetector - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.NodeProblemDetector = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MetricsServer - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.MetricsServer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CertManager - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.CertManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Networking - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Networking = kops.NetworkingSpec{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Api - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.API = kops.APISpec{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Authentication - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Authentication = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Authorization - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Authorization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeAuthorization - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.NodeAuthorization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudLabels - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.CloudLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Hooks - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Hooks = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Assets - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Assets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IAM - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.IAM = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EncryptionConfig - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.EncryptionConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseHostCertificates - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.UseHostCertificates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SysctlParameters - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.SysctlParameters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RollingUpdate - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.RollingUpdate = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterAutoscaler - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ClusterAutoscaler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountIssuerDiscovery - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.ServiceAccountIssuerDiscovery = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SnapshotController - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.SnapshotController = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Karpenter - default",
			args: args{
				in: func() kops.ClusterSpec {
					subject := kops.ClusterSpec{}
					subject.Karpenter = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceClusterSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceClusterSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
