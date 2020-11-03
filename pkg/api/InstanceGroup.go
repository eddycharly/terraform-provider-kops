package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

type InstanceGroup struct {
	Name                          string
	Role                          kops.InstanceGroupRole
	Image                         string
	MinSize                       *int32
	MaxSize                       *int32
	MachineType                   string
	RootVolumeSize                *int32
	RootVolumeType                *string
	RootVolumeIops                *int32
	RootVolumeOptimization        *bool
	RootVolumeDeleteOnTermination *bool
	RootVolumeEncryption          *bool
	Volumes                       []*kops.VolumeSpec
	VolumeMounts                  []*kops.VolumeMountSpec
	Subnets                       []string
	Zones                         []string
	Hooks                         []kops.HookSpec
	MaxPrice                      *string
	SpotDurationInMinutes         *int64
	AssociatePublicIP             *bool
	AdditionalSecurityGroups      []string
	CloudLabels                   map[string]string
	NodeLabels                    map[string]string
	FileAssets                    []kops.FileAssetSpec
	Tenancy                       string
	Kubelet                       *kops.KubeletConfigSpec
	Taints                        []string
	MixedInstancesPolicy          *kops.MixedInstancesPolicySpec
	AdditionalUserData            []kops.UserData
	SuspendProcesses              []string
	ExternalLoadBalancers         []kops.LoadBalancer
	DetailedInstanceMonitoring    *bool
	IAM                           *kops.IAMProfileSpec
	SecurityGroupOverride         *string
	InstanceProtection            *bool
	SysctlParameters              []string
	RollingUpdate                 *kops.RollingUpdate
	InstanceInterruptionBehavior  *string
}

func fromKopsInstanceGroup(instanceGroup *kops.InstanceGroup) *InstanceGroup {
	return &InstanceGroup{
		Name:                          instanceGroup.ObjectMeta.Name,
		Role:                          instanceGroup.Spec.Role,
		Image:                         instanceGroup.Spec.Image,
		MinSize:                       instanceGroup.Spec.MinSize,
		MaxSize:                       instanceGroup.Spec.MaxSize,
		MachineType:                   instanceGroup.Spec.MachineType,
		RootVolumeSize:                instanceGroup.Spec.RootVolumeSize,
		RootVolumeType:                instanceGroup.Spec.RootVolumeType,
		RootVolumeIops:                instanceGroup.Spec.RootVolumeIops,
		RootVolumeOptimization:        instanceGroup.Spec.RootVolumeOptimization,
		RootVolumeDeleteOnTermination: instanceGroup.Spec.RootVolumeDeleteOnTermination,
		RootVolumeEncryption:          instanceGroup.Spec.RootVolumeEncryption,
		Volumes:                       instanceGroup.Spec.Volumes,
		VolumeMounts:                  instanceGroup.Spec.VolumeMounts,
		Subnets:                       instanceGroup.Spec.Subnets,
		Zones:                         instanceGroup.Spec.Zones,
		Hooks:                         instanceGroup.Spec.Hooks,
		MaxPrice:                      instanceGroup.Spec.MaxPrice,
		SpotDurationInMinutes:         instanceGroup.Spec.SpotDurationInMinutes,
		AssociatePublicIP:             instanceGroup.Spec.AssociatePublicIP,
		AdditionalSecurityGroups:      instanceGroup.Spec.AdditionalSecurityGroups,
		CloudLabels:                   instanceGroup.Spec.CloudLabels,
		NodeLabels:                    instanceGroup.Spec.NodeLabels,
		FileAssets:                    instanceGroup.Spec.FileAssets,
		Tenancy:                       instanceGroup.Spec.Tenancy,
		Kubelet:                       instanceGroup.Spec.Kubelet,
		Taints:                        instanceGroup.Spec.Taints,
		MixedInstancesPolicy:          instanceGroup.Spec.MixedInstancesPolicy,
		AdditionalUserData:            instanceGroup.Spec.AdditionalUserData,
		SuspendProcesses:              instanceGroup.Spec.SuspendProcesses,
		ExternalLoadBalancers:         instanceGroup.Spec.ExternalLoadBalancers,
		DetailedInstanceMonitoring:    instanceGroup.Spec.DetailedInstanceMonitoring,
		IAM:                           instanceGroup.Spec.IAM,
		SecurityGroupOverride:         instanceGroup.Spec.SecurityGroupOverride,
		InstanceProtection:            instanceGroup.Spec.InstanceProtection,
		SysctlParameters:              instanceGroup.Spec.SysctlParameters,
		RollingUpdate:                 instanceGroup.Spec.RollingUpdate,
		InstanceInterruptionBehavior:  instanceGroup.Spec.InstanceInterruptionBehavior,
	}
}

func toKopsInstanceGroup(instanceGroup *InstanceGroup) *kops.InstanceGroup {
	return &kops.InstanceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name: instanceGroup.Name,
		},
		Spec: kops.InstanceGroupSpec{
			Role:                          instanceGroup.Role,
			Image:                         instanceGroup.Image,
			MinSize:                       instanceGroup.MinSize,
			MaxSize:                       instanceGroup.MaxSize,
			MachineType:                   instanceGroup.MachineType,
			RootVolumeSize:                instanceGroup.RootVolumeSize,
			RootVolumeType:                instanceGroup.RootVolumeType,
			RootVolumeIops:                instanceGroup.RootVolumeIops,
			RootVolumeOptimization:        instanceGroup.RootVolumeOptimization,
			RootVolumeDeleteOnTermination: instanceGroup.RootVolumeDeleteOnTermination,
			RootVolumeEncryption:          instanceGroup.RootVolumeEncryption,
			Volumes:                       instanceGroup.Volumes,
			VolumeMounts:                  instanceGroup.VolumeMounts,
			Subnets:                       instanceGroup.Subnets,
			Zones:                         instanceGroup.Zones,
			Hooks:                         instanceGroup.Hooks,
			MaxPrice:                      instanceGroup.MaxPrice,
			SpotDurationInMinutes:         instanceGroup.SpotDurationInMinutes,
			AssociatePublicIP:             instanceGroup.AssociatePublicIP,
			AdditionalSecurityGroups:      instanceGroup.AdditionalSecurityGroups,
			CloudLabels:                   instanceGroup.CloudLabels,
			NodeLabels:                    instanceGroup.NodeLabels,
			FileAssets:                    instanceGroup.FileAssets,
			Tenancy:                       instanceGroup.Tenancy,
			Kubelet:                       instanceGroup.Kubelet,
			Taints:                        instanceGroup.Taints,
			MixedInstancesPolicy:          instanceGroup.MixedInstancesPolicy,
			AdditionalUserData:            instanceGroup.AdditionalUserData,
			SuspendProcesses:              instanceGroup.SuspendProcesses,
			ExternalLoadBalancers:         instanceGroup.ExternalLoadBalancers,
			DetailedInstanceMonitoring:    instanceGroup.DetailedInstanceMonitoring,
			IAM:                           instanceGroup.IAM,
			SecurityGroupOverride:         instanceGroup.SecurityGroupOverride,
			InstanceProtection:            instanceGroup.InstanceProtection,
			SysctlParameters:              instanceGroup.SysctlParameters,
			RollingUpdate:                 instanceGroup.RollingUpdate,
			InstanceInterruptionBehavior:  instanceGroup.InstanceInterruptionBehavior,
		},
	}
}
