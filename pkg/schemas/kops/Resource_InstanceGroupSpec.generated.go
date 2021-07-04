package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ExpandResourceInstanceGroupSpec(in map[string]interface{}) kops.InstanceGroupSpec {
	if in == nil {
		panic("expand InstanceGroupSpec failure, in is nil")
	}
	out := kops.InstanceGroupSpec{}
	if in, ok := in["role"]; ok && in != nil {
		out.Role = func(in interface{}) kops.InstanceGroupRole { return kops.InstanceGroupRole(in.(string)) }(in)
	}
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["min_size"]; ok && in != nil {
		out.MinSize = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["max_size"]; ok && in != nil {
		out.MaxSize = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["autoscale"]; ok && in != nil {
		out.Autoscale = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["machine_type"]; ok && in != nil {
		out.MachineType = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["root_volume_size"]; ok && in != nil {
		out.RootVolumeSize = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["root_volume_type"]; ok && in != nil {
		out.RootVolumeType = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["root_volume_iops"]; ok && in != nil {
		out.RootVolumeIops = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["root_volume_throughput"]; ok && in != nil {
		out.RootVolumeThroughput = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["root_volume_optimization"]; ok && in != nil {
		out.RootVolumeOptimization = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["root_volume_delete_on_termination"]; ok && in != nil {
		out.RootVolumeDeleteOnTermination = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["root_volume_encryption"]; ok && in != nil {
		out.RootVolumeEncryption = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["root_volume_encryption_key"]; ok && in != nil {
		out.RootVolumeEncryptionKey = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["volumes"]; ok && in != nil {
		out.Volumes = func(in interface{}) []kops.VolumeSpec {
			var out []kops.VolumeSpec
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.VolumeSpec {
					if in == nil {
						return kops.VolumeSpec{}
					}
					return ExpandResourceVolumeSpec(in.(map[string]interface{}))
				}(in))
			}
			return out
		}(in)
	}
	if in, ok := in["volume_mounts"]; ok && in != nil {
		out.VolumeMounts = func(in interface{}) []kops.VolumeMountSpec {
			var out []kops.VolumeMountSpec
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.VolumeMountSpec {
					if in == nil {
						return kops.VolumeMountSpec{}
					}
					return ExpandResourceVolumeMountSpec(in.(map[string]interface{}))
				}(in))
			}
			return out
		}(in)
	}
	if in, ok := in["subnets"]; ok && in != nil {
		out.Subnets = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["zones"]; ok && in != nil {
		out.Zones = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["hooks"]; ok && in != nil {
		out.Hooks = func(in interface{}) []kops.HookSpec {
			var out []kops.HookSpec
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.HookSpec {
					if in == nil {
						return kops.HookSpec{}
					}
					return ExpandResourceHookSpec(in.(map[string]interface{}))
				}(in))
			}
			return out
		}(in)
	}
	if in, ok := in["max_price"]; ok && in != nil {
		out.MaxPrice = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["spot_duration_in_minutes"]; ok && in != nil {
		out.SpotDurationInMinutes = func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			return func(in int64) *int64 { return &in }(func(in interface{}) int64 { return int64(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["cpu_credits"]; ok && in != nil {
		out.CPUCredits = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["associate_public_ip"]; ok && in != nil {
		out.AssociatePublicIP = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["additional_security_groups"]; ok && in != nil {
		out.AdditionalSecurityGroups = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["cloud_labels"]; ok && in != nil {
		out.CloudLabels = func(in interface{}) map[string]string {
			if in == nil {
				return nil
			}
			out := map[string]string{}
			for key, in := range in.(map[string]interface{}) {
				out[key] = func(in interface{}) string { return string(in.(string)) }(in)
			}
			return out
		}(in)
	}
	if in, ok := in["node_labels"]; ok && in != nil {
		out.NodeLabels = func(in interface{}) map[string]string {
			if in == nil {
				return nil
			}
			out := map[string]string{}
			for key, in := range in.(map[string]interface{}) {
				out[key] = func(in interface{}) string { return string(in.(string)) }(in)
			}
			return out
		}(in)
	}
	if in, ok := in["file_assets"]; ok && in != nil {
		out.FileAssets = func(in interface{}) []kops.FileAssetSpec {
			var out []kops.FileAssetSpec
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.FileAssetSpec {
					if in == nil {
						return kops.FileAssetSpec{}
					}
					return ExpandResourceFileAssetSpec(in.(map[string]interface{}))
				}(in))
			}
			return out
		}(in)
	}
	if in, ok := in["tenancy"]; ok && in != nil {
		out.Tenancy = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["kubelet"]; ok && in != nil {
		out.Kubelet = func(in interface{}) *kops.KubeletConfigSpec {
			if in == nil {
				return nil
			}
			return func(in kops.KubeletConfigSpec) *kops.KubeletConfigSpec { return &in }(func(in interface{}) kops.KubeletConfigSpec {
				if in == nil {
					return kops.KubeletConfigSpec{}
				}
				return ExpandResourceKubeletConfigSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["taints"]; ok && in != nil {
		out.Taints = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["mixed_instances_policy"]; ok && in != nil {
		out.MixedInstancesPolicy = func(in interface{}) *kops.MixedInstancesPolicySpec {
			if in == nil {
				return nil
			}
			return func(in kops.MixedInstancesPolicySpec) *kops.MixedInstancesPolicySpec { return &in }(func(in interface{}) kops.MixedInstancesPolicySpec {
				if in == nil {
					return kops.MixedInstancesPolicySpec{}
				}
				return ExpandResourceMixedInstancesPolicySpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["additional_user_data"]; ok && in != nil {
		out.AdditionalUserData = func(in interface{}) []kops.UserData {
			var out []kops.UserData
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.UserData {
					if in == nil {
						return kops.UserData{}
					}
					return ExpandResourceUserData(in.(map[string]interface{}))
				}(in))
			}
			return out
		}(in)
	}
	if in, ok := in["suspend_processes"]; ok && in != nil {
		out.SuspendProcesses = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["external_load_balancers"]; ok && in != nil {
		out.ExternalLoadBalancers = func(in interface{}) []kops.LoadBalancer {
			var out []kops.LoadBalancer
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.LoadBalancer {
					if in == nil {
						return kops.LoadBalancer{}
					}
					return ExpandResourceLoadBalancer(in.(map[string]interface{}))
				}(in))
			}
			return out
		}(in)
	}
	if in, ok := in["detailed_instance_monitoring"]; ok && in != nil {
		out.DetailedInstanceMonitoring = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["iam"]; ok && in != nil {
		out.IAM = func(in interface{}) *kops.IAMProfileSpec {
			if in == nil {
				return nil
			}
			return func(in kops.IAMProfileSpec) *kops.IAMProfileSpec { return &in }(func(in interface{}) kops.IAMProfileSpec {
				if in == nil {
					return kops.IAMProfileSpec{}
				}
				return ExpandResourceIAMProfileSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["security_group_override"]; ok && in != nil {
		out.SecurityGroupOverride = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["instance_protection"]; ok && in != nil {
		out.InstanceProtection = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["sysctl_parameters"]; ok && in != nil {
		out.SysctlParameters = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["rolling_update"]; ok && in != nil {
		out.RollingUpdate = func(in interface{}) *kops.RollingUpdate {
			if in == nil {
				return nil
			}
			return func(in kops.RollingUpdate) *kops.RollingUpdate { return &in }(func(in interface{}) kops.RollingUpdate {
				if in == nil {
					return kops.RollingUpdate{}
				}
				return ExpandResourceRollingUpdate(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["instance_interruption_behavior"]; ok && in != nil {
		out.InstanceInterruptionBehavior = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["compress_user_data"]; ok && in != nil {
		out.CompressUserData = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["instance_metadata"]; ok && in != nil {
		out.InstanceMetadata = func(in interface{}) *kops.InstanceMetadataOptions {
			if in == nil {
				return nil
			}
			return func(in kops.InstanceMetadataOptions) *kops.InstanceMetadataOptions { return &in }(func(in interface{}) kops.InstanceMetadataOptions {
				if in == nil {
					return kops.InstanceMetadataOptions{}
				}
				return ExpandResourceInstanceMetadataOptions(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["update_policy"]; ok && in != nil {
		out.UpdatePolicy = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}

func FlattenResourceInstanceGroupSpecInto(in kops.InstanceGroupSpec, out map[string]interface{}) {
	out["role"] = func(in kops.InstanceGroupRole) interface{} { return string(in) }(in.Role)
	out["image"] = func(in string) interface{} { return string(in) }(in.Image)
	out["min_size"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.MinSize)
	out["max_size"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.MaxSize)
	out["autoscale"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.Autoscale)
	out["machine_type"] = func(in string) interface{} { return string(in) }(in.MachineType)
	out["root_volume_size"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.RootVolumeSize)
	out["root_volume_type"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.RootVolumeType)
	out["root_volume_iops"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.RootVolumeIops)
	out["root_volume_throughput"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.RootVolumeThroughput)
	out["root_volume_optimization"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.RootVolumeOptimization)
	out["root_volume_delete_on_termination"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.RootVolumeDeleteOnTermination)
	out["root_volume_encryption"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.RootVolumeEncryption)
	out["root_volume_encryption_key"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.RootVolumeEncryptionKey)
	out["volumes"] = func(in []kops.VolumeSpec) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in kops.VolumeSpec) interface{} { return FlattenResourceVolumeSpec(in) }(in))
		}
		return out
	}(in.Volumes)
	out["volume_mounts"] = func(in []kops.VolumeMountSpec) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in kops.VolumeMountSpec) interface{} { return FlattenResourceVolumeMountSpec(in) }(in))
		}
		return out
	}(in.VolumeMounts)
	out["subnets"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.Subnets)
	out["zones"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.Zones)
	out["hooks"] = func(in []kops.HookSpec) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in kops.HookSpec) interface{} { return FlattenResourceHookSpec(in) }(in))
		}
		return out
	}(in.Hooks)
	out["max_price"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.MaxPrice)
	out["spot_duration_in_minutes"] = func(in *int64) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int64) interface{} { return int(in) }(*in)}
	}(in.SpotDurationInMinutes)
	out["cpu_credits"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.CPUCredits)
	out["associate_public_ip"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.AssociatePublicIP)
	out["additional_security_groups"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.AdditionalSecurityGroups)
	out["cloud_labels"] = func(in map[string]string) interface{} {
		if in == nil {
			return nil
		}
		out := map[string]interface{}{}
		for key, in := range in {
			out[key] = func(in string) interface{} { return string(in) }(in)
		}
		return out
	}(in.CloudLabels)
	out["node_labels"] = func(in map[string]string) interface{} {
		if in == nil {
			return nil
		}
		out := map[string]interface{}{}
		for key, in := range in {
			out[key] = func(in string) interface{} { return string(in) }(in)
		}
		return out
	}(in.NodeLabels)
	out["file_assets"] = func(in []kops.FileAssetSpec) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in kops.FileAssetSpec) interface{} { return FlattenResourceFileAssetSpec(in) }(in))
		}
		return out
	}(in.FileAssets)
	out["tenancy"] = func(in string) interface{} { return string(in) }(in.Tenancy)
	out["kubelet"] = func(in *kops.KubeletConfigSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.KubeletConfigSpec) interface{} { return FlattenResourceKubeletConfigSpec(in) }(*in)
	}(in.Kubelet)
	out["taints"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.Taints)
	out["mixed_instances_policy"] = func(in *kops.MixedInstancesPolicySpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.MixedInstancesPolicySpec) interface{} { return FlattenResourceMixedInstancesPolicySpec(in) }(*in)
	}(in.MixedInstancesPolicy)
	out["additional_user_data"] = func(in []kops.UserData) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in kops.UserData) interface{} { return FlattenResourceUserData(in) }(in))
		}
		return out
	}(in.AdditionalUserData)
	out["suspend_processes"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.SuspendProcesses)
	out["external_load_balancers"] = func(in []kops.LoadBalancer) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in kops.LoadBalancer) interface{} { return FlattenResourceLoadBalancer(in) }(in))
		}
		return out
	}(in.ExternalLoadBalancers)
	out["detailed_instance_monitoring"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.DetailedInstanceMonitoring)
	out["iam"] = func(in *kops.IAMProfileSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.IAMProfileSpec) interface{} { return FlattenResourceIAMProfileSpec(in) }(*in)
	}(in.IAM)
	out["security_group_override"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.SecurityGroupOverride)
	out["instance_protection"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.InstanceProtection)
	out["sysctl_parameters"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.SysctlParameters)
	out["rolling_update"] = func(in *kops.RollingUpdate) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.RollingUpdate) interface{} { return FlattenResourceRollingUpdate(in) }(*in)
	}(in.RollingUpdate)
	out["instance_interruption_behavior"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.InstanceInterruptionBehavior)
	out["compress_user_data"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.CompressUserData)
	out["instance_metadata"] = func(in *kops.InstanceMetadataOptions) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.InstanceMetadataOptions) interface{} { return FlattenResourceInstanceMetadataOptions(in) }(*in)
	}(in.InstanceMetadata)
	out["update_policy"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.UpdatePolicy)
}

func FlattenResourceInstanceGroupSpec(in kops.InstanceGroupSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceInstanceGroupSpecInto(in, out)
	return out
}
