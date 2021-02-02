package resources

import "github.com/eddycharly/terraform-provider-kops/pkg/api/utils"

type RollingUpdateOptions struct {
	// Skip allows skipping rolling update entirely
	Skip bool
	utils.RollingUpdateOptions
}
