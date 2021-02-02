package resources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/utils"
	"k8s.io/kops/pkg/client/simple"
)

// ClusterUpdater takes care of applying changes and/or rolling update the cluster when needed
type ClusterUpdater struct {
	// ClusterName is the target cluster name
	ClusterName string
	// RollingUpdate holds rolling update options to be used when performing a rolling update
	RollingUpdate RollingUpdateOptions
}

func (u *ClusterUpdater) Apply(clientset simple.Clientset) error {
	if err := utils.ClusterApply(clientset, u.ClusterName); err != nil {
		return err
	}
	if !u.RollingUpdate.Skip {
		if err := utils.ClusterRollingUpdate(clientset, u.ClusterName, u.RollingUpdate.RollingUpdateOptions); err != nil {
			return err
		}
	}
	return nil
}
