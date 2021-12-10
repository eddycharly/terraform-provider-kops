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
	// ClusterCaCert string
	// ClusterCaKey  string
}

func GetClusterSecrets(secretStore fi.SecretStore, keyStore fi.CAStore) (*ClusterSecrets, error) {
	d, err := secretStore.FindSecret("dockerconfig")
	if err != nil {
		return nil, err
	}
	dockerConfig := ""
	if d != nil {
		dockerConfig = string(d.Data)
	}
	// TODO
	// clusterCaCert := ""
	// clusterCaKey := ""
	// c, k, err := keyStore.FindPrimaryKeypair(fi.CertificateIDCA)
	// if err != nil {
	// 	return nil, err
	// }
	// if c != nil {
	// 	clusterCaCert, err = c.AsString()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	// if k != nil {
	// 	clusterCaKey, err = k.AsString()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	if dockerConfig == "" /*&& clusterCaCert == "" && clusterCaKey == ""*/ {
		return nil, nil
	}
	return &ClusterSecrets{
		DockerConfig: dockerConfig,
		// ClusterCaCert: clusterCaCert,
		// ClusterCaKey:  clusterCaKey,
	}, nil
}

func createOrUpdateClusterSecret(secretStore fi.SecretStore, name string, s string) error {
	if s == "" {
		secret, err := secretStore.FindSecret(name)
		if err != nil {
			return err
		}
		if secret == nil {
			return nil
		}
		return secretStore.DeleteSecret(name)
	}
	secret, err := fi.CreateSecret()
	if err != nil {
		return fmt.Errorf("error creating secret %s: %v", name, err)
	}
	var parsedData map[string]interface{}
	err = json.Unmarshal([]byte(s), &parsedData)
	if err != nil {
		return fmt.Errorf("Unable to parse JSON: %v", err)
	}
	secret.Data = []byte(s)
	_, err = secretStore.ReplaceSecret(name, secret)
	if err != nil {
		return fmt.Errorf("error setting secret %s: %v", name, err)
	}
	return nil
}

// TODO
// func createOrUpdateClusterKeypair(keyStore fi.CAStore, c string, k string) error {
// 	if c == "" && k == "" {
// 		// TODO: how can we delete the certificate ?
// 		return nil
// 	}
// 	privateKey, err := pki.ParsePEMPrivateKey([]byte(k))
// 	if err != nil {
// 		return fmt.Errorf("error loading private key: %v", err)
// 	}
// 	cert, err := pki.ParsePEMCertificate([]byte(c))
// 	if err != nil {
// 		return fmt.Errorf("error loading certificate: %v", err)
// 	}
// 	err = keyStore.StoreKeypair(fi.CertificateIDCA, cert, privateKey)
// 	if err != nil {
// 		return fmt.Errorf("error storing user provided keys: %v", err)
// 	}
// 	return nil
// }

func CreateOrUpdateClusterSecrets(secretStore fi.SecretStore, keyStore fi.CAStore, secrets *ClusterSecrets) (*ClusterSecrets, error) {
	dockerConfig := ""
	// clusterCaCert := ""
	// clusterCaKey := ""
	if secrets != nil {
		dockerConfig = secrets.DockerConfig
		// clusterCaCert = secrets.ClusterCaCert
		// clusterCaKey = secrets.ClusterCaKey
	}
	if err := createOrUpdateClusterSecret(secretStore, "dockerconfig", dockerConfig); err != nil {
		return nil, err
	}
	// if err := createOrUpdateClusterKeypair(keyStore, clusterCaCert, clusterCaKey); err != nil {
	// 	return nil, err
	// }
	return &ClusterSecrets{
		DockerConfig: dockerConfig,
		// ClusterCaCert: clusterCaCert,
		// ClusterCaKey:  clusterCaKey,
	}, nil
}
