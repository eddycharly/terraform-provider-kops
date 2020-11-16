package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

// InstanceGroup represents a group of instances (either nodes or masters) with the same configuration
type InstanceGroup struct {
	// Instance group name
	Name string
	kops.InstanceGroupSpec
}

func fromKopsInstanceGroup(instanceGroup *kops.InstanceGroup) *InstanceGroup {
	return &InstanceGroup{
		Name:              instanceGroup.ObjectMeta.Name,
		InstanceGroupSpec: instanceGroup.Spec,
	}
}

func toKopsInstanceGroup(instanceGroup *InstanceGroup) *kops.InstanceGroup {
	return &kops.InstanceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name: instanceGroup.Name,
		},
		Spec: instanceGroup.InstanceGroupSpec,
	}
}
