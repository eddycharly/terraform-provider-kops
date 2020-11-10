package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/kubeconfig"
)

// Cluster defines the configuration for a cluster
// It includes cluster instance groups.
type Cluster struct {
	// The cluster name
	Name string
	// The cluster admin ssh key
	AdminSshKey string
	// The Channel we are following
	Channel string
	// Additional addons that should be installed on the cluster
	Addons []kops.AddonSpec
	// ConfigBase is the path where we store configuration for the cluster
	// This might be different than the location where the cluster spec itself is stored,
	// both because this must be accessible to the cluster,
	// and because it might be on a different cloud or storage system (etcd vs S3)
	ConfigBase string
	// The CloudProvider to use (aws or gce)
	CloudProvider string
	// Container runtime to use for Kubernetes
	ContainerRuntime string
	// The version of kubernetes to install (optional, and can be a "spec" like stable)
	KubernetesVersion string
	// Configuration of subnets we are targeting
	Subnet []kops.ClusterSubnetSpec
	// Project is the cloud project we should use, required on GCE
	Project string
	// MasterPublicName is the external DNS name for the master nodes
	MasterPublicName string
	// MasterInternalName is the internal DNS name for the master nodes
	MasterInternalName string
	// NetworkCIDR is the CIDR used for the AWS VPC / GCE Network, or otherwise allocated to k8s
	// This is a real CIDR, not the internal k8s network
	// On AWS, it maps to the VPC CIDR.  It is not required on GCE.
	NetworkCIDR string
	// AdditionalNetworkCIDRs is a list of additional CIDR used for the AWS VPC
	// or otherwise allocated to k8s. This is a real CIDR, not the internal k8s network
	// On AWS, it maps to any additional CIDRs added to a VPC.
	AdditionalNetworkCIDRs []string
	// NetworkID is an identifier of a network, if we want to reuse/share an existing network (e.g. an AWS VPC)
	NetworkID string
	// Topology defines the type of network topology to use on the cluster - default public
	// This is heavily weighted towards AWS for the time being, but should also be agnostic enough
	// to port out to GCE later if needed
	Topology *kops.TopologySpec
	// SecretStore is the VFS path to where secrets are stored
	SecretStore string
	// KeyStore is the VFS path to where SSL keys and certificates are stored
	KeyStore string
	// ConfigStore is the VFS path to where the configuration (Cluster, InstanceGroups etc) is stored
	ConfigStore string
	// DNSZone is the DNS zone we should use when configuring DNS
	// This is because some clouds let us define a managed zone foo.bar, and then have
	// kubernetes.dev.foo.bar, without needing to define dev.foo.bar as a hosted zone.
	// DNSZone will probably be a suffix of the MasterPublicName and MasterInternalName
	// Note that DNSZone can either by the host name of the zone (containing dots),
	// or can be an identifier for the zone.
	DNSZone string
	// AdditionalSANs adds additional Subject Alternate Names to apiserver cert that kops generates
	AdditionalSANs []string
	// ClusterDNSDomain is the suffix we use for internal DNS names (normally cluster.local)
	ClusterDNSDomain string
	// ServiceClusterIPRange is the CIDR, from the internal network, where we allocate IPs for services
	ServiceClusterIPRange string
	// PodCIDR is the CIDR from which we allocate IPs for pods
	PodCIDR string
	// NonMasqueradeCIDR is the CIDR for the internal k8s network (on which pods & services live)
	// It cannot overlap ServiceClusterIPRange
	NonMasqueradeCIDR string
	// SSHAccess is a list of the CIDRs that can access SSH.
	SSHAccess []string
	// NodePortAccess is a list of the CIDRs that can access the node ports range (30000-32767).
	NodePortAccess []string
	// HTTPProxy defines connection information to support use of a private cluster behind an forward HTTP Proxy
	EgressProxy *kops.EgressProxySpec
	// SSHKeyName specifies a preexisting SSH key to use
	SSHKeyName *string
	// KubernetesAPIAccess is a list of the CIDRs that can access the Kubernetes API endpoint (master HTTPS)
	KubernetesAPIAccess []string
	// IsolateMasters determines whether we should lock down masters so that they are not on the pod network.
	// true is the kube-up behaviour, but it is very surprising: it means that daemonsets only work on the master
	// if they have hostNetwork=true.
	// false is now the default, and it will:
	//  * give the master a normal PodCIDR
	//  * run kube-proxy on the master
	//  * enable debugging handlers on the master, so kubectl logs works
	IsolateMasters *bool
	// UpdatePolicy determines the policy for applying upgrades automatically.
	// Valid values:
	//   'external' do not apply updates automatically - they are applied manually or by an external system
	//   missing: default policy (currently OS security upgrades that do not require a reboot)
	UpdatePolicy *string
	// ExternalPolicies allows the insertion of pre-existing managed policies on IG Roles
	ExternalPolicies *map[string][]string
	// Additional policies to add for roles
	AdditionalPolicies *map[string]string
	// A collection of files assets for deployed cluster wide
	FileAssets []kops.FileAssetSpec
	// EtcdClusters stores the configuration for each cluster
	EtcdCluster []*kops.EtcdClusterSpec
	// Component configurations
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
	// Networking configuration
	Networking *kops.NetworkingSpec
	// API field controls how the API is exposed outside the cluster
	API *kops.AccessSpec
	// Authentication field controls how the cluster is configured for authentication
	Authentication *kops.AuthenticationSpec
	// Authorization field controls how the cluster is configured for authorization
	Authorization *kops.AuthorizationSpec
	// NodeAuthorization defined the custom node authorization configuration
	NodeAuthorization *kops.NodeAuthorizationSpec
	// Tags for AWS instance groups
	CloudLabels map[string]string
	// Hooks for custom actions e.g. on first installation
	Hooks []kops.HookSpec
	// Assets is alternative locations for files and containers; the API under construction, will remove this comment once this API is fully functional.
	Assets *kops.Assets
	// IAM field adds control over the IAM security policies applied to resources
	IAM *kops.IAMSpec
	// EncryptionConfig controls if encryption is enabled
	EncryptionConfig *bool
	// DisableSubnetTags controls if subnets are tagged in AWS
	DisableSubnetTags bool
	// UseHostCertificates will mount /etc/ssl/certs to inside needed containers.
	// This is needed if some APIs do have self-signed certs
	UseHostCertificates *bool
	// SysctlParameters will configure kernel parameters using sysctl(8). When
	// specified, each parameter must follow the form variable=value, the way
	// it would appear in sysctl.conf.
	SysctlParameters []string
	// RollingUpdate defines the default rolling-update settings for instance groups
	RollingUpdate *kops.RollingUpdate
	// InstanceGroup defines the list of instance groups making the cluster
	InstanceGroup []*InstanceGroup
	// KubeConfig holds the necessary information to connect to the cluster
	KubeConfig *KubeConfig
}

func fromKopsCluster(cluster *kops.Cluster, kubeConfig *kubeconfig.KubeconfigBuilder, instanceGroups ...*kops.InstanceGroup) *Cluster {
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
		KubeConfig:                     fromKubeconfigBuilder(kubeConfig),
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
