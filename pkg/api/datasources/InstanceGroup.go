package datasources

import (
	"k8s.io/kops/pkg/apis/kops"
)

// InstanceGroup defines a cluster instance group
type InstanceGroup struct {
	// The cluster name
	ClusterName string
	// The instance group name
	Name string
	kops.InstanceGroupSpec
}
