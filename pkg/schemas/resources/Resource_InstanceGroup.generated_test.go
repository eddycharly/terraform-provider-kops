package schemas

import (
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/google/go-cmp/cmp"
)

func TestExpandResourceInstanceGroup(t *testing.T) {
	_default := resources.InstanceGroup{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want resources.InstanceGroup
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"role":                              "",
					"image":                             "",
					"min_size":                          nil,
					"max_size":                          nil,
					"autoscale":                         nil,
					"machine_type":                      "",
					"root_volume_size":                  nil,
					"root_volume_type":                  nil,
					"root_volume_iops":                  nil,
					"root_volume_throughput":            nil,
					"root_volume_optimization":          nil,
					"root_volume_delete_on_termination": nil,
					"root_volume_encryption":            nil,
					"root_volume_encryption_key":        nil,
					"volumes":                           func() []interface{} { return nil }(),
					"volume_mounts":                     func() []interface{} { return nil }(),
					"subnets":                           func() []interface{} { return nil }(),
					"zones":                             func() []interface{} { return nil }(),
					"hooks":                             func() []interface{} { return nil }(),
					"max_price":                         nil,
					"spot_duration_in_minutes":          nil,
					"cpu_credits":                       nil,
					"associate_public_ip":               nil,
					"additional_security_groups":        func() []interface{} { return nil }(),
					"cloud_labels":                      func() map[string]interface{} { return nil }(),
					"node_labels":                       func() map[string]interface{} { return nil }(),
					"file_assets":                       func() []interface{} { return nil }(),
					"tenancy":                           "",
					"kubelet":                           nil,
					"taints":                            func() []interface{} { return nil }(),
					"mixed_instances_policy":            nil,
					"additional_user_data":              func() []interface{} { return nil }(),
					"suspend_processes":                 func() []interface{} { return nil }(),
					"external_load_balancers":           func() []interface{} { return nil }(),
					"detailed_instance_monitoring":      nil,
					"iam":                               nil,
					"security_group_override":           nil,
					"instance_protection":               nil,
					"sysctl_parameters":                 func() []interface{} { return nil }(),
					"rolling_update":                    nil,
					"instance_interruption_behavior":    nil,
					"compress_user_data":                nil,
					"instance_metadata":                 nil,
					"update_policy":                     nil,
					"warm_pool":                         nil,
					"revision":                          0,
					"cluster_name":                      "",
					"name":                              "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceInstanceGroup(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceInstanceGroup() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceInstanceGroupInto(t *testing.T) {
	_default := map[string]interface{}{
		"role":                              "",
		"image":                             "",
		"min_size":                          nil,
		"max_size":                          nil,
		"autoscale":                         nil,
		"machine_type":                      "",
		"root_volume_size":                  nil,
		"root_volume_type":                  nil,
		"root_volume_iops":                  nil,
		"root_volume_throughput":            nil,
		"root_volume_optimization":          nil,
		"root_volume_delete_on_termination": nil,
		"root_volume_encryption":            nil,
		"root_volume_encryption_key":        nil,
		"volumes":                           func() []interface{} { return nil }(),
		"volume_mounts":                     func() []interface{} { return nil }(),
		"subnets":                           func() []interface{} { return nil }(),
		"zones":                             func() []interface{} { return nil }(),
		"hooks":                             func() []interface{} { return nil }(),
		"max_price":                         nil,
		"spot_duration_in_minutes":          nil,
		"cpu_credits":                       nil,
		"associate_public_ip":               nil,
		"additional_security_groups":        func() []interface{} { return nil }(),
		"cloud_labels":                      func() map[string]interface{} { return nil }(),
		"node_labels":                       func() map[string]interface{} { return nil }(),
		"file_assets":                       func() []interface{} { return nil }(),
		"tenancy":                           "",
		"kubelet":                           nil,
		"taints":                            func() []interface{} { return nil }(),
		"mixed_instances_policy":            nil,
		"additional_user_data":              func() []interface{} { return nil }(),
		"suspend_processes":                 func() []interface{} { return nil }(),
		"external_load_balancers":           func() []interface{} { return nil }(),
		"detailed_instance_monitoring":      nil,
		"iam":                               nil,
		"security_group_override":           nil,
		"instance_protection":               nil,
		"sysctl_parameters":                 func() []interface{} { return nil }(),
		"rolling_update":                    nil,
		"instance_interruption_behavior":    nil,
		"compress_user_data":                nil,
		"instance_metadata":                 nil,
		"update_policy":                     nil,
		"warm_pool":                         nil,
		"revision":                          0,
		"cluster_name":                      "",
		"name":                              "",
	}
	type args struct {
		in resources.InstanceGroup
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: resources.InstanceGroup{},
			},
			want: _default,
		},
		{
			name: "Role - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Role = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MinSize - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.MinSize = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxSize - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.MaxSize = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Autoscale - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Autoscale = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MachineType - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.MachineType = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootVolumeSize - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RootVolumeSize = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootVolumeType - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RootVolumeType = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootVolumeIops - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RootVolumeIops = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootVolumeThroughput - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RootVolumeThroughput = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootVolumeOptimization - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RootVolumeOptimization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootVolumeDeleteOnTermination - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RootVolumeDeleteOnTermination = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootVolumeEncryption - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RootVolumeEncryption = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootVolumeEncryptionKey - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RootVolumeEncryptionKey = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Volumes - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Volumes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VolumeMounts - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.VolumeMounts = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Subnets - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Subnets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Zones - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Zones = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Hooks - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Hooks = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxPrice - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.MaxPrice = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SpotDurationInMinutes - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.SpotDurationInMinutes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CPUCredits - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.CPUCredits = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AssociatePublicIp - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.AssociatePublicIP = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalSecurityGroups - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.AdditionalSecurityGroups = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudLabels - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.CloudLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeLabels - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.NodeLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FileAssets - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.FileAssets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Tenancy - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Tenancy = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Kubelet - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Kubelet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Taints - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Taints = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MixedInstancesPolicy - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.MixedInstancesPolicy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalUserData - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.AdditionalUserData = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SuspendProcesses - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.SuspendProcesses = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalLoadBalancers - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.ExternalLoadBalancers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DetailedInstanceMonitoring - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.DetailedInstanceMonitoring = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IAM - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.IAM = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecurityGroupOverride - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.SecurityGroupOverride = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InstanceProtection - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.InstanceProtection = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SysctlParameters - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.SysctlParameters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RollingUpdate - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RollingUpdate = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InstanceInterruptionBehavior - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.InstanceInterruptionBehavior = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CompressUserData - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.CompressUserData = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InstanceMetadata - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.InstanceMetadata = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UpdatePolicy - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.UpdatePolicy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WarmPool - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.WarmPool = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Revision - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Revision = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceInstanceGroupInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceInstanceGroup() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceInstanceGroup(t *testing.T) {
	_default := map[string]interface{}{
		"role":                              "",
		"image":                             "",
		"min_size":                          nil,
		"max_size":                          nil,
		"autoscale":                         nil,
		"machine_type":                      "",
		"root_volume_size":                  nil,
		"root_volume_type":                  nil,
		"root_volume_iops":                  nil,
		"root_volume_throughput":            nil,
		"root_volume_optimization":          nil,
		"root_volume_delete_on_termination": nil,
		"root_volume_encryption":            nil,
		"root_volume_encryption_key":        nil,
		"volumes":                           func() []interface{} { return nil }(),
		"volume_mounts":                     func() []interface{} { return nil }(),
		"subnets":                           func() []interface{} { return nil }(),
		"zones":                             func() []interface{} { return nil }(),
		"hooks":                             func() []interface{} { return nil }(),
		"max_price":                         nil,
		"spot_duration_in_minutes":          nil,
		"cpu_credits":                       nil,
		"associate_public_ip":               nil,
		"additional_security_groups":        func() []interface{} { return nil }(),
		"cloud_labels":                      func() map[string]interface{} { return nil }(),
		"node_labels":                       func() map[string]interface{} { return nil }(),
		"file_assets":                       func() []interface{} { return nil }(),
		"tenancy":                           "",
		"kubelet":                           nil,
		"taints":                            func() []interface{} { return nil }(),
		"mixed_instances_policy":            nil,
		"additional_user_data":              func() []interface{} { return nil }(),
		"suspend_processes":                 func() []interface{} { return nil }(),
		"external_load_balancers":           func() []interface{} { return nil }(),
		"detailed_instance_monitoring":      nil,
		"iam":                               nil,
		"security_group_override":           nil,
		"instance_protection":               nil,
		"sysctl_parameters":                 func() []interface{} { return nil }(),
		"rolling_update":                    nil,
		"instance_interruption_behavior":    nil,
		"compress_user_data":                nil,
		"instance_metadata":                 nil,
		"update_policy":                     nil,
		"warm_pool":                         nil,
		"revision":                          0,
		"cluster_name":                      "",
		"name":                              "",
	}
	type args struct {
		in resources.InstanceGroup
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: resources.InstanceGroup{},
			},
			want: _default,
		},
		{
			name: "Role - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Role = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MinSize - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.MinSize = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxSize - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.MaxSize = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Autoscale - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Autoscale = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MachineType - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.MachineType = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootVolumeSize - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RootVolumeSize = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootVolumeType - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RootVolumeType = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootVolumeIops - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RootVolumeIops = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootVolumeThroughput - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RootVolumeThroughput = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootVolumeOptimization - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RootVolumeOptimization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootVolumeDeleteOnTermination - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RootVolumeDeleteOnTermination = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootVolumeEncryption - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RootVolumeEncryption = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootVolumeEncryptionKey - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RootVolumeEncryptionKey = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Volumes - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Volumes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VolumeMounts - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.VolumeMounts = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Subnets - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Subnets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Zones - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Zones = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Hooks - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Hooks = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxPrice - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.MaxPrice = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SpotDurationInMinutes - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.SpotDurationInMinutes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CPUCredits - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.CPUCredits = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AssociatePublicIp - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.AssociatePublicIP = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalSecurityGroups - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.AdditionalSecurityGroups = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudLabels - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.CloudLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeLabels - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.NodeLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FileAssets - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.FileAssets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Tenancy - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Tenancy = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Kubelet - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Kubelet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Taints - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Taints = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MixedInstancesPolicy - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.MixedInstancesPolicy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalUserData - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.AdditionalUserData = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SuspendProcesses - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.SuspendProcesses = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalLoadBalancers - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.ExternalLoadBalancers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DetailedInstanceMonitoring - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.DetailedInstanceMonitoring = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IAM - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.IAM = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecurityGroupOverride - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.SecurityGroupOverride = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InstanceProtection - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.InstanceProtection = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SysctlParameters - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.SysctlParameters = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RollingUpdate - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.RollingUpdate = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InstanceInterruptionBehavior - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.InstanceInterruptionBehavior = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CompressUserData - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.CompressUserData = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InstanceMetadata - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.InstanceMetadata = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UpdatePolicy - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.UpdatePolicy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WarmPool - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.WarmPool = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Revision - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Revision = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() resources.InstanceGroup {
					subject := resources.InstanceGroup{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceInstanceGroup(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceInstanceGroup() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
