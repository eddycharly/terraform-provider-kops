package datasources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/utils"
	"k8s.io/kops/pkg/client/simple"
)

// ClusterStatus reports status of a cluster
type ClusterStatus struct {
	// ClusterName defines the target cluster name
	ClusterName string
	// Exists indicates if the cluster exists
	Exists bool
	// IsValid indicates if the cluster is valid
	IsValid bool
	// NeedsUpdate indicates if the cluster needs a rolling update
	NeedsUpdate bool
	// InstanceGroups contains the name of instance groups to be updated
	InstanceGroups []string
}

func (s *ClusterStatus) GetClusterStatus(clientset simple.Clientset) error {
	if exists, err := utils.ClusterExists(clientset, s.ClusterName); err != nil {
		return err
	} else {
		if exists {
			if isValid, err := utils.ClusterIsValid(clientset, s.ClusterName); err != nil {
				return err
			} else {
				s.IsValid = isValid
			}
			if needsUpdate, err := utils.ClusterInstanceGroupsNeedingUpdate(clientset, s.ClusterName); err != nil {
				return err
			} else {
				s.NeedsUpdate = len(needsUpdate) != 0
				s.InstanceGroups = needsUpdate
			}
		}
		return nil
	}
}
