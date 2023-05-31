package resources

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
	kopsutil "k8s.io/kops/pkg/apis/kops/util"
	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/pkg/clusteraddons"
	"k8s.io/kops/pkg/resources"
	"k8s.io/kops/pkg/resources/ops"
	"k8s.io/kops/pkg/wellknownoperators"
	"k8s.io/kops/upup/pkg/fi/cloudup"
)

// Cluster defines the configuration for a cluster
type Cluster struct {
	kops.ClusterSpec
	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and services.
	Labels map[string]string
	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata. They are not
	// queryable and should be preserved when modifying objects.
	Annotations map[string]string
	// Name defines the cluster name
	Name string
	// AdminSshKey defines the cluster admin ssh key
	AdminSshKey string
	// ClusterAddons defines the cluster addons
	ClusterAddons []string
	// Secrets defines the cluster secret
	Secrets *ClusterSecrets
	// Revision is incremented every time the resource changes, this is useful for triggering cluster updater
	Revision int
}

func makeCluster(adminSshKey string, secrets *ClusterSecrets, cluster *kops.Cluster) *Cluster {
	return &Cluster{
		ClusterSpec: cluster.Spec,
		Labels:      cluster.Labels,
		Annotations: cluster.Annotations,
		Name:        cluster.ObjectMeta.Name,
		AdminSshKey: adminSshKey,
		Secrets:     secrets,
	}
}

func makeKopsCluster(name string, labels map[string]string, annotations map[string]string, spec kops.ClusterSpec) *kops.Cluster {
	return &kops.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Labels:      labels,
			Annotations: annotations,
		},
		Spec: spec,
	}
}

func GetCluster(name string, clientset simple.Clientset) (*Cluster, error) {
	kc, err := clientset.GetCluster(context.Background(), name)
	if err != nil {
		return nil, err
	}
	sshCredentialStore, err := clientset.SSHCredentialStore(kc)
	if err != nil {
		return nil, err
	}
	pubKeys, err := sshCredentialStore.FindSSHPublicKeys()
	if err != nil {
		return nil, err
	}
	secretStore, err := clientset.SecretStore(kc)
	if err != nil {
		return nil, err
	}
	keyStore, err := clientset.KeyStore(kc)
	if err != nil {
		return nil, err
	}
	secrets, err := GetClusterSecrets(secretStore, keyStore)
	if err != nil {
		return nil, err
	}
	adminSshKey := ""
	if len(pubKeys) > 0 {
		adminSshKey = pubKeys[0].Spec.PublicKey
	}
	cluster := makeCluster(adminSshKey, secrets, kc)
	return cluster, nil
}

func CreateCluster(name string, labels map[string]string, annotations map[string]string, adminSshKey string, secrets *ClusterSecrets, clusterAddons []string, spec kops.ClusterSpec, clientset simple.Clientset) (*Cluster, error) {
	kc := makeKopsCluster(name, labels, annotations, spec)
	cloud, err := cloudup.BuildCloud(kc)
	if err != nil {
		return nil, err
	}
	if err := cloudup.PerformAssignments(kc, cloud); err != nil {
		return nil, err
	}
	// TODO: deep validate ?
	// TODO: assets builder ?
	channel, err := cloudup.ChannelForCluster(kc)
	if err != nil {
		return nil, err
	}
	kubernetesVersion, err := kopsutil.ParseKubernetesVersion(kc.Spec.KubernetesVersion)
	if err != nil {
		return nil, err
	}
	addons, err := wellknownoperators.CreateAddons(channel, kubernetesVersion, kc)
	if err != nil {
		return nil, err
	}
	for _, addon := range clusterAddons {
		addon, err := clusteraddons.ParseClusterAddon([]byte(addon))
		if err != nil {
			return nil, err
		}
		addons = append(addons, addon.Objects...)
	}
	_, err = clientset.CreateCluster(context.Background(), kc)
	if err != nil {
		return nil, err
	}
	kc, err = clientset.GetCluster(context.Background(), name)
	if err != nil {
		return nil, err
	}
	addonsClient := clientset.AddonsFor(kc)
	if err := addonsClient.Replace(addons); err != nil {
		return nil, err
	}
	if adminSshKey != "" {
		sshCredentialStore, err := clientset.SSHCredentialStore(kc)
		if err != nil {
			return nil, err
		}
		if err = sshCredentialStore.AddSSHPublicKey([]byte(adminSshKey)); err != nil {
			return nil, fmt.Errorf("error adding SSH public key: %v", err)
		}
	}
	secretStore, err := clientset.SecretStore(kc)
	if err != nil {
		return nil, err
	}
	keyStore, err := clientset.KeyStore(kc)
	if err != nil {
		return nil, err
	}
	secrets, err = CreateOrUpdateClusterSecrets(secretStore, keyStore, secrets)
	if err != nil {
		return nil, err
	}
	kc, err = clientset.GetCluster(context.Background(), name)
	if err != nil {
		return nil, err
	}
	return makeCluster(adminSshKey, secrets, kc), nil
}

func UpdateCluster(name string, labels map[string]string, annotations map[string]string, adminSshKey string, secrets *ClusterSecrets, clusterAddons []string, spec kops.ClusterSpec, clientset simple.Clientset) (*Cluster, error) {
	kc := makeKopsCluster(name, labels, annotations, spec)
	cloud, err := cloudup.BuildCloud(kc)
	if err != nil {
		return nil, err
	}
	if err := cloudup.PerformAssignments(kc, cloud); err != nil {
		return nil, err
	}
	// TODO: deep validate ?
	// TODO: assets builder ?
	channel, err := cloudup.ChannelForCluster(kc)
	if err != nil {
		return nil, err
	}
	kubernetesVersion, err := kopsutil.ParseKubernetesVersion(kc.Spec.KubernetesVersion)
	if err != nil {
		return nil, err
	}
	addons, err := wellknownoperators.CreateAddons(channel, kubernetesVersion, kc)
	if err != nil {
		return nil, err
	}
	for _, addon := range clusterAddons {
		addon, err := clusteraddons.ParseClusterAddon([]byte(addon))
		if err != nil {
			return nil, err
		}
		addons = append(addons, addon.Objects...)
	}
	kc, err = clientset.UpdateCluster(context.Background(), kc, nil)
	if err != nil {
		return nil, err
	}
	addonsClient := clientset.AddonsFor(kc)
	if err := addonsClient.Replace(addons); err != nil {
		return nil, err
	}
	sshCredentialStore, err := clientset.SSHCredentialStore(kc)
	if err != nil {
		return nil, err
	}
	if adminSshKey != "" {
		if err = sshCredentialStore.AddSSHPublicKey([]byte(adminSshKey)); err != nil {
			return nil, fmt.Errorf("error adding SSH public key: %v", err)
		}
	} else {
		if pubKeys, err := sshCredentialStore.FindSSHPublicKeys(); err != nil {
			return nil, fmt.Errorf("error listing SSH public keys: %v", err)
		} else if len(pubKeys) > 0 {
			if err = sshCredentialStore.DeleteSSHCredential(); err != nil {
				return nil, fmt.Errorf("error deleting SSH public key: %v", err)
			}
		}
	}
	secretStore, err := clientset.SecretStore(kc)
	if err != nil {
		return nil, err
	}
	keyStore, err := clientset.KeyStore(kc)
	if err != nil {
		return nil, err
	}
	secrets, err = CreateOrUpdateClusterSecrets(secretStore, keyStore, secrets)
	if err != nil {
		return nil, err
	}
	kc, err = clientset.GetCluster(context.Background(), name)
	if err != nil {
		return nil, err
	}
	return makeCluster(adminSshKey, secrets, kc), nil
}

func DeleteCluster(name string, clientset simple.Clientset) error {
	kc, err := clientset.GetCluster(context.Background(), name)
	if err != nil {
		return err
	}
	cloud, err := cloudup.BuildCloud(kc)
	if err != nil {
		return err
	}
	allResources, err := ops.ListResources(cloud, kc, "")
	if err != nil {
		return err
	}
	clusterResources := make(map[string]*resources.Resource)
	for k, resource := range allResources {
		if resource.Shared {
			continue
		}
		clusterResources[k] = resource
	}
	if len(clusterResources) != 0 {
		err = ops.DeleteResources(cloud, clusterResources)
		if err != nil {
			return err
		}
	}
	err = clientset.DeleteCluster(context.Background(), kc)
	if err != nil {
		return err
	}
	return nil
}
