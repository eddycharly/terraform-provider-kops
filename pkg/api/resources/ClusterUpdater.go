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
	// Apply holds cluster apply options
	Apply ApplyOptions
	// RollingUpdate holds cluster rolling update options
	RollingUpdate RollingUpdateOptions
	// Validate holds cluster validation options
	Validate ValidateOptions
}

func (u *ClusterUpdater) UpdateCluster(clientset simple.Clientset) error {
	if !u.Apply.Skip {
		if lifecycleOverrides, err := utils.ParseLifecycleOverrides(u.Apply.LifecycleOverrides); err != nil {
			return err
		} else {
			if err := utils.ClusterApply(clientset, u.ClusterName, u.Apply.AllowKopsDowngrade, lifecycleOverrides); err != nil {
				return err
			}
		}
	}
	if !u.Validate.Skip {
		if err := utils.ClusterValidate(clientset, u.ClusterName, u.Validate.ValidateOptions); err != nil {
			return err
		}
	}
	if !u.RollingUpdate.Skip {
		if err := utils.ClusterRollingUpdate(clientset, u.ClusterName, u.RollingUpdate.RollingUpdateOptions); err != nil {
			return err
		}
	}
	return nil
}
