package resources

import (
	"encoding/json"
	"fmt"

	"k8s.io/kops/upup/pkg/fi"
)

// ClusterSecrets defines cluster secrets
type ClusterSecrets struct {
	// DockerConfig holds a valid docker config.
	// After creating a dockerconfig secret, a /root/.docker/config.json file will be added to newly created nodes.
	// This file will be used by Kubernetes to authenticate to container registries and will also work when using containerd as container runtime.
	DockerConfig string
}

func GetClusterSecrets(secretStore fi.SecretStore) (*ClusterSecrets, error) {
	secret, err := secretStore.FindSecret("dockerconfig")
	if err != nil {
		return nil, err
	}
	if secret == nil {
		return nil, nil
	}
	return &ClusterSecrets{
		DockerConfig: string(secret.Data),
	}, nil
}

func CreateOrUpdateClusterSecrets(secretStore fi.SecretStore, secrets *ClusterSecrets) (*ClusterSecrets, error) {
	if secrets == nil || secrets.DockerConfig == "" {
		secret, err := secretStore.FindSecret("dockerconfig")
		if err != nil {
			return nil, err
		}
		if secret == nil {
			return nil, nil
		}
		return nil, secretStore.DeleteSecret("dockerconfig")
	}
	secret, err := fi.CreateSecret()
	if err != nil {
		return nil, fmt.Errorf("error creating docker config secret: %v", err)
	}
	var parsedData map[string]interface{}
	err = json.Unmarshal([]byte(secrets.DockerConfig), &parsedData)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse JSON: %v", err)
	}
	secret.Data = []byte(secrets.DockerConfig)
	_, err = secretStore.ReplaceSecret("dockerconfig", secret)
	if err != nil {
		return nil, fmt.Errorf("error setting dockerconfig secret: %v", err)
	}
	return &ClusterSecrets{
		DockerConfig: string(secret.Data),
	}, nil
}
