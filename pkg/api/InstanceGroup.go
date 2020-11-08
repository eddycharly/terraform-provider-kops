package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

// InstanceGroup represents a group of instances (either nodes or masters) with the same configuration
type InstanceGroup struct {
	// Instance group name
	Name string
	// Determines the role of instances in this group: masters or nodes
	Role kops.InstanceGroupRole
	// Image is the instance (ami etc) we should use
	Image string
	// MinSize is the minimum size of the pool
	MinSize *int32
	// MaxSize is the maximum size of the pool
	MaxSize *int32
	// MachineType is the instance class
	MachineType string
	// RootVolumeSize is the size of the EBS root volume to use, in GB
	RootVolumeSize *int32
	// RootVolumeType is the type of the EBS root volume to use (e.g. gp2)
	RootVolumeType *string
	// If volume type is io1, then we need to specify the number of Iops.
	RootVolumeIops *int32
	// RootVolumeOptimization enables EBS optimization for an instance
	RootVolumeOptimization *bool
	// RootVolumeDeleteOnTermination configures root volume retention policy upon instance termination.
	// The root volume is deleted by default. Cluster deletion does not remove retained root volumes.
	// NOTE: This setting applies only to the Launch Configuration and does not affect Launch Templates.
	RootVolumeDeleteOnTermination *bool
	// RootVolumeEncryption enables EBS root volume encryption for an instance
	RootVolumeEncryption *bool
	// Volumes is a collection of additional volumes to create for instances within this InstanceGroup
	Volumes []*kops.VolumeSpec
	// VolumeMounts a collection of volume mounts
	VolumeMounts []*kops.VolumeMountSpec
	// Subnets is the names of the Subnets (as specified in the Cluster) where machines in this instance group should be placed
	Subnets []string
	// Zones is the names of the Zones where machines in this instance group should be placed
	// This is needed for regional subnets (e.g. GCE), to restrict placement to particular zones
	Zones []string
	// Hooks is a list of hooks for this instanceGroup, note: these can override the cluster wide ones if required
	Hooks []kops.HookSpec
	// MaxPrice indicates this is a spot-pricing group, with the specified value as our max-price bid
	MaxPrice *string
	// SpotDurationInMinutes reserves a spot block for the period specified
	SpotDurationInMinutes *int64
	// AssociatePublicIP is true if we want instances to have a public IP
	AssociatePublicIP *bool
	// AdditionalSecurityGroups attaches additional security groups (e.g. i-123456)
	AdditionalSecurityGroups []string
	// CloudLabels indicates the labels for instances in this group, at the AWS level
	CloudLabels map[string]string
	// NodeLabels indicates the kubernetes labels for nodes in this group
	NodeLabels map[string]string
	// FileAssets is a collection of file assets for this instance group
	FileAssets []kops.FileAssetSpec
	// Describes the tenancy of the instance group. Can be either default or dedicated. Currently only applies to AWS.
	Tenancy string
	// Kubelet overrides kubelet config from the ClusterSpec
	Kubelet *kops.KubeletConfigSpec
	// Taints indicates the kubernetes taints for nodes in this group
	Taints []string
	// MixedInstancesPolicy defined a optional backing of an AWS ASG by a EC2 Fleet (AWS Only)
	MixedInstancesPolicy *kops.MixedInstancesPolicySpec
	// AdditionalUserData is any additional user-data to be passed to the host
	AdditionalUserData []kops.UserData
	// SuspendProcesses disables the listed Scaling Policies
	SuspendProcesses []string
	// ExternalLoadBalancers define loadbalancers that should be attached to the instancegroup
	ExternalLoadBalancers []kops.LoadBalancer
	// DetailedInstanceMonitoring defines if detailed-monitoring is enabled (AWS only)
	DetailedInstanceMonitoring *bool
	// IAMProfileSpec defines the identity of the cloud group IAM profile (AWS only).
	IAM *kops.IAMProfileSpec
	// SecurityGroupOverride overrides the default security group created by Kops for this IG (AWS only).
	SecurityGroupOverride *string
	// InstanceProtection makes new instances in an autoscaling group protected from scale in
	InstanceProtection *bool
	// SysctlParameters will configure kernel parameters using sysctl(8). When
	// specified, each parameter must follow the form variable=value, the way
	// it would appear in sysctl.conf.
	SysctlParameters []string
	// RollingUpdate defines the rolling-update behavior
	RollingUpdate *kops.RollingUpdate
	// InstanceInterruptionBehavior defines if a spot instance should be terminated, hibernated,
	// or stopped after interruption
	InstanceInterruptionBehavior *string
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
