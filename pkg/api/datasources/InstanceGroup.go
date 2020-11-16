package datasources

import "github.com/eddycharly/terraform-provider-kops/pkg/api/resources"

// InstanceGroup defines a cluster instance group
type InstanceGroup struct {
	// The cluster name
	ClusterName string
	resources.InstanceGroup
}
