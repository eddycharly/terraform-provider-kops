package datasources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/utils"
	"k8s.io/kops/pkg/client/simple"
)

// ClusterStatus reports status of a cluster
type ClusterStatus struct {
	// The target cluster name
	ClusterName string
	// Tells if the cluster exists
	Exists bool
	// Tells if the cluster is valid
	IsValid bool
	// Tells if the cluster needs a rolling update
	NeedsUpdate bool
}

func GetClusterStatus(name string, clientset simple.Clientset) (*ClusterStatus, error) {
	if exists, err := utils.ClusterExists(name, clientset); err != nil {
		return nil, err
	} else {
		isValid := false
		needsUpdate := true
		if exists {
			isValid, _ = utils.IsClusterValid(name, clientset)
			needsUpdate, _ = utils.NeedsUpdate(name, clientset)
		}
		return &ClusterStatus{
				ClusterName: name,
				Exists:      exists,
				IsValid:     isValid,
				NeedsUpdate: needsUpdate,
			},
			nil
	}
}
