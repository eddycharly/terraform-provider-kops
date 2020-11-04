package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/kubeconfig"
)

type Cluster struct {
	Name                           string
	AdminSshKey                    string
	Channel                        string
	Addons                         []kops.AddonSpec
	ConfigBase                     string
	CloudProvider                  string
	ContainerRuntime               string
	KubernetesVersion              string
	Subnet                         []kops.ClusterSubnetSpec
	Project                        string
	MasterPublicName               string
	MasterInternalName             string
	NetworkCIDR                    string
	AdditionalNetworkCIDRs         []string
	NetworkID                      string
	Topology                       *kops.TopologySpec
	SecretStore                    string
	KeyStore                       string
	ConfigStore                    string
	DNSZone                        string
	AdditionalSANs                 []string
	ClusterDNSDomain               string
	ServiceClusterIPRange          string
	PodCIDR                        string
	NonMasqueradeCIDR              string
	SSHAccess                      []string
	NodePortAccess                 []string
	EgressProxy                    *kops.EgressProxySpec
	SSHKeyName                     *string
	KubernetesAPIAccess            []string
	IsolateMasters                 *bool
	UpdatePolicy                   *string
	ExternalPolicies               *map[string][]string
	AdditionalPolicies             *map[string]string
	FileAssets                     []kops.FileAssetSpec
	EtcdCluster                    []*kops.EtcdClusterSpec
	Containerd                     *kops.ContainerdConfig
	Docker                         *kops.DockerConfig
	KubeDNS                        *kops.KubeDNSConfig
	KubeAPIServer                  *kops.KubeAPIServerConfig
	KubeControllerManager          *kops.KubeControllerManagerConfig
	ExternalCloudControllerManager *kops.CloudControllerManagerConfig
	KubeScheduler                  *kops.KubeSchedulerConfig
	KubeProxy                      *kops.KubeProxyConfig
	Kubelet                        *kops.KubeletConfigSpec
	MasterKubelet                  *kops.KubeletConfigSpec
	CloudConfig                    *kops.CloudConfiguration
	ExternalDNS                    *kops.ExternalDNSConfig
	Networking                     *kops.NetworkingSpec
	API                            *kops.AccessSpec
	Authentication                 *kops.AuthenticationSpec
	Authorization                  *kops.AuthorizationSpec
	NodeAuthorization              *kops.NodeAuthorizationSpec
	CloudLabels                    map[string]string
	Hooks                          []kops.HookSpec
	Assets                         *kops.Assets
	IAM                            *kops.IAMSpec
	EncryptionConfig               *bool
	DisableSubnetTags              bool
	UseHostCertificates            *bool
	SysctlParameters               []string
	RollingUpdate                  *kops.RollingUpdate
	KubeServer                     string
	KubeCertificateAuthority       string
	KubeClientCertificate          string
	KubeClientKey                  string
	KubeUsername                   string
	KubePassword                   string
	InstanceGroup                  []*InstanceGroup
	RollingUpdateOptions           *RollingUpdateOptions
}

func fromKopsCluster(cluster *kops.Cluster, config *kubeconfig.KubeconfigBuilder, instanceGroups ...*kops.InstanceGroup) *Cluster {
	return &Cluster{
		Name:                           cluster.ObjectMeta.Name,
		Channel:                        cluster.Spec.Channel,
		ConfigBase:                     cluster.Spec.ConfigBase,
		CloudProvider:                  cluster.Spec.CloudProvider,
		ContainerRuntime:               cluster.Spec.ContainerRuntime,
		KubernetesVersion:              cluster.Spec.KubernetesVersion,
		Subnet:                         cluster.Spec.Subnets,
		Project:                        cluster.Spec.Project,
		MasterPublicName:               cluster.Spec.MasterPublicName,
		MasterInternalName:             cluster.Spec.MasterInternalName,
		NetworkCIDR:                    cluster.Spec.NetworkCIDR,
		AdditionalNetworkCIDRs:         cluster.Spec.AdditionalNetworkCIDRs,
		NetworkID:                      cluster.Spec.NetworkID,
		Topology:                       cluster.Spec.Topology,
		SecretStore:                    cluster.Spec.SecretStore,
		KeyStore:                       cluster.Spec.KeyStore,
		ConfigStore:                    cluster.Spec.ConfigStore,
		DNSZone:                        cluster.Spec.DNSZone,
		AdditionalSANs:                 cluster.Spec.AdditionalSANs,
		ClusterDNSDomain:               cluster.Spec.ClusterDNSDomain,
		ServiceClusterIPRange:          cluster.Spec.ServiceClusterIPRange,
		PodCIDR:                        cluster.Spec.PodCIDR,
		NonMasqueradeCIDR:              cluster.Spec.NonMasqueradeCIDR,
		SSHAccess:                      cluster.Spec.SSHAccess,
		NodePortAccess:                 cluster.Spec.NodePortAccess,
		EgressProxy:                    cluster.Spec.EgressProxy,
		SSHKeyName:                     cluster.Spec.SSHKeyName,
		KubernetesAPIAccess:            cluster.Spec.KubernetesAPIAccess,
		IsolateMasters:                 cluster.Spec.IsolateMasters,
		UpdatePolicy:                   cluster.Spec.UpdatePolicy,
		ExternalPolicies:               cluster.Spec.ExternalPolicies,
		AdditionalPolicies:             cluster.Spec.AdditionalPolicies,
		FileAssets:                     cluster.Spec.FileAssets,
		EtcdCluster:                    cluster.Spec.EtcdClusters,
		Containerd:                     cluster.Spec.Containerd,
		Docker:                         cluster.Spec.Docker,
		KubeDNS:                        cluster.Spec.KubeDNS,
		KubeAPIServer:                  cluster.Spec.KubeAPIServer,
		KubeControllerManager:          cluster.Spec.KubeControllerManager,
		ExternalCloudControllerManager: cluster.Spec.ExternalCloudControllerManager,
		KubeScheduler:                  cluster.Spec.KubeScheduler,
		KubeProxy:                      cluster.Spec.KubeProxy,
		Kubelet:                        cluster.Spec.Kubelet,
		MasterKubelet:                  cluster.Spec.MasterKubelet,
		CloudConfig:                    cluster.Spec.CloudConfig,
		ExternalDNS:                    cluster.Spec.ExternalDNS,
		Networking:                     cluster.Spec.Networking,
		API:                            cluster.Spec.API,
		Authentication:                 cluster.Spec.Authentication,
		Authorization:                  cluster.Spec.Authorization,
		NodeAuthorization:              cluster.Spec.NodeAuthorization,
		CloudLabels:                    cluster.Spec.CloudLabels,
		Hooks:                          cluster.Spec.Hooks,
		Assets:                         cluster.Spec.Assets,
		IAM:                            cluster.Spec.IAM,
		EncryptionConfig:               cluster.Spec.EncryptionConfig,
		DisableSubnetTags:              cluster.Spec.DisableSubnetTags,
		UseHostCertificates:            cluster.Spec.UseHostCertificates,
		SysctlParameters:               cluster.Spec.SysctlParameters,
		RollingUpdate:                  cluster.Spec.RollingUpdate,
		KubeServer:                     config.Server,
		KubeCertificateAuthority:       string(config.CACert),
		KubeClientCertificate:          string(config.ClientCert),
		KubeClientKey:                  string(config.ClientKey),
		KubeUsername:                   config.KubeUser,
		KubePassword:                   config.KubePassword,
		InstanceGroup: func(in ...*kops.InstanceGroup) []*InstanceGroup {
			var out []*InstanceGroup
			for _, in := range in {
				out = append(out, fromKopsInstanceGroup(in))
			}
			return out
		}(instanceGroups...),
	}
}

func toKopsCluster(cluster *Cluster) (*kops.Cluster, []*kops.InstanceGroup) {
	c := kops.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: cluster.Name,
		},
		Spec: kops.ClusterSpec{
			Channel:                        cluster.Channel,
			ConfigBase:                     cluster.ConfigBase,
			CloudProvider:                  cluster.CloudProvider,
			ContainerRuntime:               cluster.ContainerRuntime,
			KubernetesVersion:              cluster.KubernetesVersion,
			Subnets:                        cluster.Subnet,
			Project:                        cluster.Project,
			MasterPublicName:               cluster.MasterPublicName,
			MasterInternalName:             cluster.MasterInternalName,
			NetworkCIDR:                    cluster.NetworkCIDR,
			AdditionalNetworkCIDRs:         cluster.AdditionalNetworkCIDRs,
			NetworkID:                      cluster.NetworkID,
			Topology:                       cluster.Topology,
			SecretStore:                    cluster.SecretStore,
			KeyStore:                       cluster.KeyStore,
			ConfigStore:                    cluster.ConfigStore,
			DNSZone:                        cluster.DNSZone,
			AdditionalSANs:                 cluster.AdditionalSANs,
			ClusterDNSDomain:               cluster.ClusterDNSDomain,
			ServiceClusterIPRange:          cluster.ServiceClusterIPRange,
			PodCIDR:                        cluster.PodCIDR,
			NonMasqueradeCIDR:              cluster.NonMasqueradeCIDR,
			SSHAccess:                      cluster.SSHAccess,
			NodePortAccess:                 cluster.NodePortAccess,
			EgressProxy:                    cluster.EgressProxy,
			SSHKeyName:                     cluster.SSHKeyName,
			KubernetesAPIAccess:            cluster.KubernetesAPIAccess,
			IsolateMasters:                 cluster.IsolateMasters,
			UpdatePolicy:                   cluster.UpdatePolicy,
			ExternalPolicies:               cluster.ExternalPolicies,
			AdditionalPolicies:             cluster.AdditionalPolicies,
			FileAssets:                     cluster.FileAssets,
			EtcdClusters:                   cluster.EtcdCluster,
			Containerd:                     cluster.Containerd,
			Docker:                         cluster.Docker,
			KubeDNS:                        cluster.KubeDNS,
			KubeAPIServer:                  cluster.KubeAPIServer,
			KubeControllerManager:          cluster.KubeControllerManager,
			ExternalCloudControllerManager: cluster.ExternalCloudControllerManager,
			KubeScheduler:                  cluster.KubeScheduler,
			KubeProxy:                      cluster.KubeProxy,
			Kubelet:                        cluster.Kubelet,
			MasterKubelet:                  cluster.MasterKubelet,
			CloudConfig:                    cluster.CloudConfig,
			ExternalDNS:                    cluster.ExternalDNS,
			Networking:                     cluster.Networking,
			API:                            cluster.API,
			Authentication:                 cluster.Authentication,
			Authorization:                  cluster.Authorization,
			NodeAuthorization:              cluster.NodeAuthorization,
			CloudLabels:                    cluster.CloudLabels,
			Hooks:                          cluster.Hooks,
			Assets:                         cluster.Assets,
			IAM:                            cluster.IAM,
			EncryptionConfig:               cluster.EncryptionConfig,
			DisableSubnetTags:              cluster.DisableSubnetTags,
			UseHostCertificates:            cluster.UseHostCertificates,
			SysctlParameters:               cluster.SysctlParameters,
			RollingUpdate:                  cluster.RollingUpdate,
		},
	}
	var ig []*kops.InstanceGroup
	for _, instanceGroup := range cluster.InstanceGroup {
		ig = append(ig, toKopsInstanceGroup(instanceGroup))
	}
	return &c, ig
}
