package utils

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/kops/pkg/client/simple"
)

func ClusterExists(clientset simple.Clientset, clusterName string) (bool, error) {
	_, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		if errors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
