package resources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/utils"
	"k8s.io/kops/pkg/client/simple"
)

// ClusterUpdater takes care of applying changes and/or rolling update the cluster when needed
type ClusterUpdater struct {
	// Revision is incremented every time the resource changes, this is useful for triggering cluster updater
	Revision int
	// ClusterName is the target cluster name
	ClusterName string
	// Keepers contains arbitrary strings used to update the resource when one changes
	Keepers map[string]string
	// RollingUpdate holds cluster rolling update options
	RollingUpdate RollingUpdateOptions
	// Validate holds cluster validation options
	Validate ValidateOptions
}

func (u *ClusterUpdater) UpdateCluster(clientset simple.Clientset) error {
	if err := utils.ClusterApply(clientset, u.ClusterName); err != nil {
		return err
	}
	if !u.RollingUpdate.Skip {
		if err := utils.ClusterRollingUpdate(clientset, u.ClusterName, u.RollingUpdate.RollingUpdateOptions); err != nil {
			return err
		}
	}
	if !u.Validate.Skip {
		if err := utils.ClusterValidate(clientset, u.ClusterName, u.Validate.ValidateOptions); err != nil {
			return err
		}
	}
	return nil
}
