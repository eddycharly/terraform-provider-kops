package schemas

import (
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	kopsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kops"
	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceCluster(t *testing.T) {
	_default := resources.Cluster{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want resources.Cluster
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"channel":     "",
					"addons":      func() []interface{} { return nil }(),
					"config_base": "",
					"cloud_provider": func() []interface{} {
						return []interface{}{kopsschemas.FlattenResourceCloudProviderSpec(kops.CloudProviderSpec{})}
					}(),
					"container_runtime":                 "",
					"kubernetes_version":                "",
					"subnet":                            func() []interface{} { return nil }(),
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
					"master_kubelet":                    nil,
					"cloud_config":                      nil,
					"external_dns":                      nil,
					"ntp":                               nil,
					"node_termination_handler":          nil,
					"node_problem_detector":             nil,
					"metrics_server":                    nil,
					"cert_manager":                      nil,
					"aws_load_balancer_controller":      nil,
					"networking":                        nil,
					"api":                               nil,
					"authentication":                    nil,
					"authorization":                     nil,
					"node_authorization":                nil,
					"cloud_labels":                      func() map[string]interface{} { return nil }(),
					"hooks":                             func() []interface{} { return nil }(),
					"assets":                            nil,
					"iam":                               nil,
					"encryption_config":                 nil,
					"tag_subnets":                       nil,
					"use_host_certificates":             nil,
					"sysctl_parameters":                 func() []interface{} { return nil }(),
					"rolling_update":                    nil,
					"cluster_autoscaler":                nil,
					"warm_pool":                         nil,
					"service_account_issuer_discovery":  nil,
					"snapshot_controller":               nil,
					"karpenter":                         nil,
					"pod_identity_webhook":              nil,
					"labels":                            func() map[string]interface{} { return nil }(),
					"annotations":                       func() map[string]interface{} { return nil }(),
					"name":                              "",
					"admin_ssh_key":                     "",
					"secrets":                           nil,
					"revision":                          0,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceCluster(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceCluster() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceClusterInto(t *testing.T) {
	_default := map[string]interface{}{
		"channel":     "",
		"addons":      func() []interface{} { return nil }(),
		"config_base": "",
		"cloud_provider": func() []interface{} {
			return []interface{}{kopsschemas.FlattenResourceCloudProviderSpec(kops.CloudProviderSpec{})}
		}(),
		"container_runtime":                 "",
		"kubernetes_version":                "",
		"subnet":                            func() []interface{} { return nil }(),
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
		"master_kubelet":                    nil,
		"cloud_config":                      nil,
		"external_dns":                      nil,
		"ntp":                               nil,
		"node_termination_handler":          nil,
		"node_problem_detector":             nil,
		"metrics_server":                    nil,
		"cert_manager":                      nil,
		"aws_load_balancer_controller":      nil,
		"networking":                        nil,
		"api":                               nil,
		"authentication":                    nil,
		"authorization":                     nil,
		"node_authorization":                nil,
		"cloud_labels":                      func() map[string]interface{} { return nil }(),
		"hooks":                             func() []interface{} { return nil }(),
		"assets":                            nil,
		"iam":                               nil,
		"encryption_config":                 nil,
		"tag_subnets":                       nil,
		"use_host_certificates":             nil,
		"sysctl_parameters":                 func() []interface{} { return nil }(),
		"rolling_update":                    nil,
		"cluster_autoscaler":                nil,
		"warm_pool":                         nil,
		"service_account_issuer_discovery":  nil,
		"snapshot_controller":               nil,
		"karpenter":                         nil,
		"pod_identity_webhook":              nil,
		"labels":                            func() map[string]interface{} { return nil }(),
		"annotations":                       func() map[string]interface{} { return nil }(),
		"name":                              "",
		"admin_ssh_key":                     "",
		"secrets":                           nil,
		"revision":                          0,
	}
	type args struct {
		in resources.Cluster
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: resources.Cluster{},
			},
			want: _default,
		},
		{
			name: "Channel - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Channel = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Addons - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Addons = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigBase - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ConfigBase = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudProvider - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.CloudProvider = kops.CloudProviderSpec{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerRuntime - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ContainerRuntime = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubernetesVersion - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.KubernetesVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Subnet - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Subnets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Project - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Project = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MasterPublicName - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.MasterPublicName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MasterInternalName - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.MasterInternalName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NetworkCidr - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.NetworkCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalNetworkCidrs - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.AdditionalNetworkCIDRs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NetworkID - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.NetworkID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Topology - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Topology = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecretStore - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.SecretStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KeyStore - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.KeyStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigStore - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ConfigStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DnsZone - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.DNSZone = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalSans - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.AdditionalSANs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterDnsDomain - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ClusterDNSDomain = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceClusterIpRange - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ServiceClusterIPRange = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodCidr - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.PodCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NonMasqueradeCidr - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.NonMasqueradeCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SshAccess - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.SSHAccess = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodePortAccess - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.NodePortAccess = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EgressProxy - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.EgressProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SshKeyName - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.SSHKeyName = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubernetesApiAccess - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.KubernetesAPIAccess = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IsolateMasters - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.IsolateMasters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UpdatePolicy - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.UpdatePolicy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalPolicies - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ExternalPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalPolicies - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.AdditionalPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FileAssets - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.FileAssets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdCluster - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.EtcdClusters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Containerd - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Containerd = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Docker - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Docker = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeDns - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.KubeDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeApiServer - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.KubeAPIServer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeControllerManager - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.KubeControllerManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalCloudControllerManager - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ExternalCloudControllerManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeScheduler - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.KubeScheduler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeProxy - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.KubeProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Kubelet - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Kubelet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MasterKubelet - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.MasterKubelet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudConfig - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.CloudConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalDns - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ExternalDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NTP - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.NTP = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeTerminationHandler - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.NodeTerminationHandler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeProblemDetector - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.NodeProblemDetector = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MetricsServer - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.MetricsServer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CertManager - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.CertManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AwsLoadBalancerController - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.AWSLoadBalancerController = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Networking - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Networking = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Api - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.API = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Authentication - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Authentication = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Authorization - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Authorization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeAuthorization - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.NodeAuthorization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudLabels - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.CloudLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Hooks - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Hooks = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Assets - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Assets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IAM - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.IAM = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EncryptionConfig - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.EncryptionConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TagSubnets - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.TagSubnets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseHostCertificates - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.UseHostCertificates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SysctlParameters - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.SysctlParameters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RollingUpdate - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.RollingUpdate = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterAutoscaler - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ClusterAutoscaler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WarmPool - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.WarmPool = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountIssuerDiscovery - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ServiceAccountIssuerDiscovery = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SnapshotController - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.SnapshotController = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Karpenter - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Karpenter = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodIdentityWebhook - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.PodIdentityWebhook = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Labels - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Labels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Annotations - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Annotations = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdminSshKey - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.AdminSshKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secrets - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Secrets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Revision - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Revision = 0
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceClusterInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCluster() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCluster(t *testing.T) {
	_default := map[string]interface{}{
		"channel":     "",
		"addons":      func() []interface{} { return nil }(),
		"config_base": "",
		"cloud_provider": func() []interface{} {
			return []interface{}{kopsschemas.FlattenResourceCloudProviderSpec(kops.CloudProviderSpec{})}
		}(),
		"container_runtime":                 "",
		"kubernetes_version":                "",
		"subnet":                            func() []interface{} { return nil }(),
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
		"master_kubelet":                    nil,
		"cloud_config":                      nil,
		"external_dns":                      nil,
		"ntp":                               nil,
		"node_termination_handler":          nil,
		"node_problem_detector":             nil,
		"metrics_server":                    nil,
		"cert_manager":                      nil,
		"aws_load_balancer_controller":      nil,
		"networking":                        nil,
		"api":                               nil,
		"authentication":                    nil,
		"authorization":                     nil,
		"node_authorization":                nil,
		"cloud_labels":                      func() map[string]interface{} { return nil }(),
		"hooks":                             func() []interface{} { return nil }(),
		"assets":                            nil,
		"iam":                               nil,
		"encryption_config":                 nil,
		"tag_subnets":                       nil,
		"use_host_certificates":             nil,
		"sysctl_parameters":                 func() []interface{} { return nil }(),
		"rolling_update":                    nil,
		"cluster_autoscaler":                nil,
		"warm_pool":                         nil,
		"service_account_issuer_discovery":  nil,
		"snapshot_controller":               nil,
		"karpenter":                         nil,
		"pod_identity_webhook":              nil,
		"labels":                            func() map[string]interface{} { return nil }(),
		"annotations":                       func() map[string]interface{} { return nil }(),
		"name":                              "",
		"admin_ssh_key":                     "",
		"secrets":                           nil,
		"revision":                          0,
	}
	type args struct {
		in resources.Cluster
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: resources.Cluster{},
			},
			want: _default,
		},
		{
			name: "Channel - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Channel = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Addons - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Addons = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigBase - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ConfigBase = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudProvider - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.CloudProvider = kops.CloudProviderSpec{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerRuntime - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ContainerRuntime = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubernetesVersion - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.KubernetesVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Subnet - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Subnets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Project - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Project = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MasterPublicName - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.MasterPublicName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MasterInternalName - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.MasterInternalName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NetworkCidr - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.NetworkCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalNetworkCidrs - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.AdditionalNetworkCIDRs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NetworkID - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.NetworkID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Topology - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Topology = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecretStore - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.SecretStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KeyStore - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.KeyStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigStore - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ConfigStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DnsZone - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.DNSZone = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalSans - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.AdditionalSANs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterDnsDomain - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ClusterDNSDomain = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceClusterIpRange - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ServiceClusterIPRange = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodCidr - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.PodCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NonMasqueradeCidr - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.NonMasqueradeCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SshAccess - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.SSHAccess = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodePortAccess - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.NodePortAccess = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EgressProxy - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.EgressProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SshKeyName - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.SSHKeyName = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubernetesApiAccess - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.KubernetesAPIAccess = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IsolateMasters - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.IsolateMasters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UpdatePolicy - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.UpdatePolicy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalPolicies - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ExternalPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalPolicies - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.AdditionalPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FileAssets - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.FileAssets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdCluster - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.EtcdClusters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Containerd - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Containerd = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Docker - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Docker = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeDns - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.KubeDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeApiServer - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.KubeAPIServer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeControllerManager - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.KubeControllerManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalCloudControllerManager - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ExternalCloudControllerManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeScheduler - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.KubeScheduler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeProxy - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.KubeProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Kubelet - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Kubelet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MasterKubelet - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.MasterKubelet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudConfig - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.CloudConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalDns - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ExternalDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NTP - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.NTP = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeTerminationHandler - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.NodeTerminationHandler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeProblemDetector - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.NodeProblemDetector = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MetricsServer - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.MetricsServer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CertManager - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.CertManager = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AwsLoadBalancerController - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.AWSLoadBalancerController = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Networking - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Networking = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Api - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.API = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Authentication - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Authentication = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Authorization - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Authorization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeAuthorization - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.NodeAuthorization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudLabels - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.CloudLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Hooks - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Hooks = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Assets - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Assets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IAM - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.IAM = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EncryptionConfig - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.EncryptionConfig = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TagSubnets - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.TagSubnets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseHostCertificates - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.UseHostCertificates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SysctlParameters - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.SysctlParameters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RollingUpdate - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.RollingUpdate = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterAutoscaler - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ClusterAutoscaler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WarmPool - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.WarmPool = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountIssuerDiscovery - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.ServiceAccountIssuerDiscovery = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SnapshotController - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.SnapshotController = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Karpenter - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Karpenter = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodIdentityWebhook - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.PodIdentityWebhook = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Labels - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Labels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Annotations - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Annotations = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdminSshKey - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.AdminSshKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secrets - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Secrets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Revision - default",
			args: args{
				in: func() resources.Cluster {
					subject := resources.Cluster{}
					subject.Revision = 0
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceCluster(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCluster() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
