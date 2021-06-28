package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/upup/pkg/fi"
)

// Cluster defines the configuration for a cluster
type DockerConfig struct {
	// ClusterName defines the cluster name the instance group belongs to
	ClusterName string
	// DockerConfig is the content of the docker config file
	DockerConfig string
}

func GetDockerConfig(clusterName string, clientset simple.Clientset) (*DockerConfig, error) {
	cluster, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return nil, err
	}
	secretStore, err := clientset.SecretStore(cluster)
	if err != nil {
		return nil, err
	}
	secret, err := secretStore.FindSecret("dockerconfig")
	if err != nil {
		return nil, err
	}
	return &DockerConfig{
		ClusterName:  clusterName,
		DockerConfig: string(secret.Data),
	}, nil
}

func CreateOrUpdateDockerConfig(clusterName, dockerConfig string, clientset simple.Clientset) (*DockerConfig, error) {
	secret, err := fi.CreateSecret()
	if err != nil {
		return nil, fmt.Errorf("error creating docker config secret: %v", err)
	}
	cluster, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return nil, err
	}
	secretStore, err := clientset.SecretStore(cluster)
	if err != nil {
		return nil, err
	}
	var parsedData map[string]interface{}
	err = json.Unmarshal([]byte(dockerConfig), &parsedData)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse JSON: %v", err)
	}
	secret.Data = []byte(dockerConfig)
	_, err = secretStore.ReplaceSecret("dockerconfig", secret)
	if err != nil {
		return nil, fmt.Errorf("error setting dockerconfig secret: %v", err)
	}
	return &DockerConfig{
		ClusterName:  clusterName,
		DockerConfig: dockerConfig,
	}, nil
}

func DeleteDockerConfig(clusterName string, clientset simple.Clientset) error {
	cluster, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return err
	}
	secretStore, err := clientset.SecretStore(cluster)
	if err != nil {
		return err
	}
	err = secretStore.DeleteSecret("dockerconfig")
	if err != nil {
		return err
	}
	return nil
}
