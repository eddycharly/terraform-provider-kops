package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops/v1alpha2"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceClusterSpec(t *testing.T) {
	_default := kopsv1alpha2.ClusterSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.ClusterSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"channel":                           "",
					"addons":                            func() []interface{} { return nil }(),
					"config_base":                       "",
					"legacy_cloud_provider":             "",
					"container_runtime":                 "",
					"kubernetes_version":                "",
					"subnets":                           func() []interface{} { return nil }(),
					"project":                           "",
					"master_public_name":                "",
					"master_internal_name":              "",
					"network_cidr":                      "",
					"additional_network_cidrs":          func() []interface{} { return nil }(),
					"network_id":                        "",
					"topology":                          nil,
					"secret_store":                      "",
					"key_store":                         "",
					"config_store":                      "",
					"dns_zone":                          "",
					"additional_sans":                   func() []interface{} { return nil }(),
					"cluster_dns_domain":                "",
					"service_cluster_ip_range":          "",
					"pod_cidr":                          "",
					"non_masquerade_cidr":               "",
					"ssh_access":                        func() []interface{} { return nil }(),
					"node_port_access":                  func() []interface{} { return nil }(),
					"egress_proxy":                      nil,
					"ssh_key_name":                      nil,
					"kubernetes_api_access":             func() []interface{} { return nil }(),
					"isolate_masters":                   nil,
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
					"node_termination_handler":          nil,
					"node_problem_detector":             nil,
					"metrics_server":                    nil,
					"cert_manager":                      nil,
					"aws_load_balancer_controller":      nil,
					"legacy_networking":                 nil,
					"networking": func() []interface{} {
						return []interface{}{FlattenResourceNetworkingSpec(kopsv1alpha2.NetworkingSpec{})}
					}(),
					"legacy_api":                       nil,
					"api":                              func() []interface{} { return []interface{}{FlattenResourceAPISpec(kopsv1alpha2.APISpec{})} }(),
					"authentication":                   nil,
					"authorization":                    nil,
					"node_authorization":               nil,
					"cloud_labels":                     func() map[string]interface{} { return nil }(),
					"hooks":                            func() []interface{} { return nil }(),
					"assets":                           nil,
					"iam":                              nil,
					"encryption_config":                nil,
					"tag_subnets":                      nil,
					"use_host_certificates":            nil,
					"sysctl_parameters":                func() []interface{} { return nil }(),
					"rolling_update":                   nil,
					"cluster_autoscaler":               nil,
					"warm_pool":                        nil,
					"service_account_issuer_discovery": nil,
					"snapshot_controller":              nil,
					"karpenter":                        nil,
					"pod_identity_webhook":             nil,
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
		"legacy_cloud_provider":             "",
		"container_runtime":                 "",
		"kubernetes_version":                "",
		"subnets":                           func() []interface{} { return nil }(),
		"project":                           "",
		"master_public_name":                "",
		"master_internal_name":              "",
		"network_cidr":                      "",
		"additional_network_cidrs":          func() []interface{} { return nil }(),
		"network_id":                        "",
		"topology":                          nil,
		"secret_store":                      "",
		"key_store":                         "",
		"config_store":                      "",
		"dns_zone":                          "",
		"additional_sans":                   func() []interface{} { return nil }(),
		"cluster_dns_domain":                "",
		"service_cluster_ip_range":          "",
		"pod_cidr":                          "",
		"non_masquerade_cidr":               "",
		"ssh_access":                        func() []interface{} { return nil }(),
		"node_port_access":                  func() []interface{} { return nil }(),
		"egress_proxy":                      nil,
		"ssh_key_name":                      nil,
		"kubernetes_api_access":             func() []interface{} { return nil }(),
		"isolate_masters":                   nil,
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
		"node_termination_handler":          nil,
		"node_problem_detector":             nil,
		"metrics_server":                    nil,
		"cert_manager":                      nil,
		"aws_load_balancer_controller":      nil,
		"legacy_networking":                 nil,
		"networking": func() []interface{} {
			return []interface{}{FlattenResourceNetworkingSpec(kopsv1alpha2.NetworkingSpec{})}
		}(),
		"legacy_api":                       nil,
		"api":                              func() []interface{} { return []interface{}{FlattenResourceAPISpec(kopsv1alpha2.APISpec{})} }(),
		"authentication":                   nil,
		"authorization":                    nil,
		"node_authorization":               nil,
		"cloud_labels":                     func() map[string]interface{} { return nil }(),
		"hooks":                            func() []interface{} { return nil }(),
		"assets":                           nil,
		"iam":                              nil,
		"encryption_config":                nil,
		"tag_subnets":                      nil,
		"use_host_certificates":            nil,
		"sysctl_parameters":                func() []interface{} { return nil }(),
		"rolling_update":                   nil,
		"cluster_autoscaler":               nil,
		"warm_pool":                        nil,
		"service_account_issuer_discovery": nil,
		"snapshot_controller":              nil,
		"karpenter":                        nil,
		"pod_identity_webhook":             nil,
	}
	type args struct {
		in kopsv1alpha2.ClusterSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.ClusterSpec{},
			},
			want: _default,
		},
		{
			name: "Channel - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Channel = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Addons - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Addons = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigBase - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ConfigBase = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LegacyCloudProvider - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.LegacyCloudProvider = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerRuntime - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ContainerRuntime = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubernetesVersion - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.KubernetesVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Subnets - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Subnets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Project - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Project = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MasterPublicName - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.MasterPublicName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MasterInternalName - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.MasterInternalName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NetworkCidr - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.NetworkCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalNetworkCidrs - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.AdditionalNetworkCIDRs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NetworkId - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.NetworkID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Topology - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Topology = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecretStore - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.SecretStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KeyStore - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.KeyStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigStore - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ConfigStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DnsZone - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.DNSZone = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalSans - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.AdditionalSANs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterDnsDomain - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ClusterDNSDomain = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceClusterIpRange - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ServiceClusterIPRange = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodCidr - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.PodCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NonMasqueradeCidr - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.NonMasqueradeCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SshAccess - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.SSHAccess = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodePortAccess - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.NodePortAccess = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EgressProxy - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.EgressProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SshKeyName - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.SSHKeyName = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubernetesApiAccess - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.KubernetesAPIAccess = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IsolateMasters - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.IsolateMasters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UpdatePolicy - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.UpdatePolicy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalPolicies - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ExternalPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalPolicies - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.AdditionalPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FileAssets - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.FileAssets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdCluster - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.EtcdClusters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Containerd - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Containerd = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Docker - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Docker = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeDns - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.KubeDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeApiServer - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.KubeAPIServer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeControllerManager - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.KubeControllerManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalCloudControllerManager - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ExternalCloudControllerManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeScheduler - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.KubeScheduler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeProxy - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.KubeProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Kubelet - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Kubelet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ControlPlaneKubelet - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ControlPlaneKubelet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudConfig - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.CloudConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalDns - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ExternalDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NTP - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.NTP = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeTerminationHandler - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.NodeTerminationHandler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeProblemDetector - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.NodeProblemDetector = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MetricsServer - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.MetricsServer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CertManager - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.CertManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AwsLoadBalancerController - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.AWSLoadBalancerController = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LegacyNetworking - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.LegacyNetworking = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Networking - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Networking = v1alpha2.NetworkingSpec{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LegacyApi - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.LegacyAPI = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Api - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.API = v1alpha2.APISpec{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Authentication - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Authentication = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Authorization - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Authorization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeAuthorization - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.NodeAuthorization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudLabels - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.CloudLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Hooks - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Hooks = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Assets - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Assets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IAM - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.IAM = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EncryptionConfig - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.EncryptionConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TagSubnets - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.TagSubnets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseHostCertificates - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.UseHostCertificates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SysctlParameters - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.SysctlParameters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RollingUpdate - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.RollingUpdate = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterAutoscaler - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ClusterAutoscaler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WarmPool - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.WarmPool = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountIssuerDiscovery - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ServiceAccountIssuerDiscovery = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SnapshotController - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.SnapshotController = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Karpenter - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Karpenter = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodIdentityWebhook - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.PodIdentityWebhook = nil
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
		"legacy_cloud_provider":             "",
		"container_runtime":                 "",
		"kubernetes_version":                "",
		"subnets":                           func() []interface{} { return nil }(),
		"project":                           "",
		"master_public_name":                "",
		"master_internal_name":              "",
		"network_cidr":                      "",
		"additional_network_cidrs":          func() []interface{} { return nil }(),
		"network_id":                        "",
		"topology":                          nil,
		"secret_store":                      "",
		"key_store":                         "",
		"config_store":                      "",
		"dns_zone":                          "",
		"additional_sans":                   func() []interface{} { return nil }(),
		"cluster_dns_domain":                "",
		"service_cluster_ip_range":          "",
		"pod_cidr":                          "",
		"non_masquerade_cidr":               "",
		"ssh_access":                        func() []interface{} { return nil }(),
		"node_port_access":                  func() []interface{} { return nil }(),
		"egress_proxy":                      nil,
		"ssh_key_name":                      nil,
		"kubernetes_api_access":             func() []interface{} { return nil }(),
		"isolate_masters":                   nil,
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
		"node_termination_handler":          nil,
		"node_problem_detector":             nil,
		"metrics_server":                    nil,
		"cert_manager":                      nil,
		"aws_load_balancer_controller":      nil,
		"legacy_networking":                 nil,
		"networking": func() []interface{} {
			return []interface{}{FlattenResourceNetworkingSpec(kopsv1alpha2.NetworkingSpec{})}
		}(),
		"legacy_api":                       nil,
		"api":                              func() []interface{} { return []interface{}{FlattenResourceAPISpec(kopsv1alpha2.APISpec{})} }(),
		"authentication":                   nil,
		"authorization":                    nil,
		"node_authorization":               nil,
		"cloud_labels":                     func() map[string]interface{} { return nil }(),
		"hooks":                            func() []interface{} { return nil }(),
		"assets":                           nil,
		"iam":                              nil,
		"encryption_config":                nil,
		"tag_subnets":                      nil,
		"use_host_certificates":            nil,
		"sysctl_parameters":                func() []interface{} { return nil }(),
		"rolling_update":                   nil,
		"cluster_autoscaler":               nil,
		"warm_pool":                        nil,
		"service_account_issuer_discovery": nil,
		"snapshot_controller":              nil,
		"karpenter":                        nil,
		"pod_identity_webhook":             nil,
	}
	type args struct {
		in kopsv1alpha2.ClusterSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.ClusterSpec{},
			},
			want: _default,
		},
		{
			name: "Channel - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Channel = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Addons - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Addons = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigBase - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ConfigBase = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LegacyCloudProvider - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.LegacyCloudProvider = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerRuntime - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ContainerRuntime = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubernetesVersion - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.KubernetesVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Subnets - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Subnets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Project - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Project = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MasterPublicName - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.MasterPublicName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MasterInternalName - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.MasterInternalName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NetworkCidr - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.NetworkCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalNetworkCidrs - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.AdditionalNetworkCIDRs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NetworkId - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.NetworkID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Topology - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Topology = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecretStore - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.SecretStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KeyStore - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.KeyStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigStore - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ConfigStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DnsZone - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.DNSZone = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalSans - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.AdditionalSANs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterDnsDomain - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ClusterDNSDomain = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceClusterIpRange - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ServiceClusterIPRange = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodCidr - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.PodCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NonMasqueradeCidr - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.NonMasqueradeCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SshAccess - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.SSHAccess = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodePortAccess - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.NodePortAccess = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EgressProxy - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.EgressProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SshKeyName - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.SSHKeyName = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubernetesApiAccess - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.KubernetesAPIAccess = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IsolateMasters - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.IsolateMasters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UpdatePolicy - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.UpdatePolicy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalPolicies - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ExternalPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalPolicies - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.AdditionalPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FileAssets - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.FileAssets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdCluster - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.EtcdClusters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Containerd - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Containerd = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Docker - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Docker = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeDns - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.KubeDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeApiServer - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.KubeAPIServer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeControllerManager - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.KubeControllerManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalCloudControllerManager - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ExternalCloudControllerManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeScheduler - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.KubeScheduler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeProxy - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.KubeProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Kubelet - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Kubelet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ControlPlaneKubelet - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ControlPlaneKubelet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudConfig - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.CloudConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalDns - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ExternalDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NTP - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.NTP = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeTerminationHandler - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.NodeTerminationHandler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeProblemDetector - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.NodeProblemDetector = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MetricsServer - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.MetricsServer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CertManager - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.CertManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AwsLoadBalancerController - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.AWSLoadBalancerController = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LegacyNetworking - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.LegacyNetworking = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Networking - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Networking = v1alpha2.NetworkingSpec{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LegacyApi - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.LegacyAPI = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Api - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.API = v1alpha2.APISpec{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Authentication - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Authentication = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Authorization - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Authorization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeAuthorization - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.NodeAuthorization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudLabels - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.CloudLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Hooks - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Hooks = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Assets - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Assets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IAM - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.IAM = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EncryptionConfig - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.EncryptionConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TagSubnets - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.TagSubnets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseHostCertificates - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.UseHostCertificates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SysctlParameters - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.SysctlParameters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RollingUpdate - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.RollingUpdate = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterAutoscaler - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ClusterAutoscaler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WarmPool - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.WarmPool = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountIssuerDiscovery - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.ServiceAccountIssuerDiscovery = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SnapshotController - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.SnapshotController = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Karpenter - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.Karpenter = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodIdentityWebhook - default",
			args: args{
				in: func() kopsv1alpha2.ClusterSpec {
					subject := kopsv1alpha2.ClusterSpec{}
					subject.PodIdentityWebhook = nil
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
