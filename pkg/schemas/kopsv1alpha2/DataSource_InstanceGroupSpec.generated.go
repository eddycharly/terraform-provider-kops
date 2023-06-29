package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops/v1alpha2"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ExpandDataSourceInstanceGroupSpec(in map[string]interface{}) kopsv1alpha2.InstanceGroupSpec {
	if in == nil {
		panic("expand InstanceGroupSpec failure, in is nil")
	}
	return kopsv1alpha2.InstanceGroupSpec{
		Manager: func(in interface{}) kopsv1alpha2.InstanceManager {
			return v1alpha2.InstanceManager(ExpandString(in))
		}(in["manager"]),
		Role: func(in interface{}) kopsv1alpha2.InstanceGroupRole {
			return v1alpha2.InstanceGroupRole(ExpandString(in))
		}(in["role"]),
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		MinSize: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["min_size"]),
		MaxSize: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["max_size"]),
		Autoscale: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["autoscale"]),
		AutoscalePriority: func(in interface{}) int16 {
			return int16(ExpandInt(in))
		}(in["autoscale_priority"]),
		MachineType: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["machine_type"]),
		RootVolumeSize: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["root_volume_size"]),
		RootVolumeType: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["root_volume_type"]),
		RootVolumeIOPS: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["root_volume_iops"]),
		RootVolumeThroughput: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["root_volume_throughput"]),
		RootVolumeOptimization: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["root_volume_optimization"]),
		RootVolumeDeleteOnTermination: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["root_volume_delete_on_termination"]),
		RootVolumeEncryption: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["root_volume_encryption"]),
		RootVolumeEncryptionKey: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["root_volume_encryption_key"]),
		Volumes: func(in interface{}) []kopsv1alpha2.VolumeSpec {
			return func(in interface{}) []kopsv1alpha2.VolumeSpec {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.VolumeSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.VolumeSpec {
						if in == nil {
							return kopsv1alpha2.VolumeSpec{}
						}
						return ExpandDataSourceVolumeSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["volumes"]),
		VolumeMounts: func(in interface{}) []kopsv1alpha2.VolumeMountSpec {
			return func(in interface{}) []kopsv1alpha2.VolumeMountSpec {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.VolumeMountSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.VolumeMountSpec {
						if in == nil {
							return kopsv1alpha2.VolumeMountSpec{}
						}
						return ExpandDataSourceVolumeMountSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["volume_mounts"]),
		Subnets: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["subnets"]),
		Zones: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["zones"]),
		Hooks: func(in interface{}) []kopsv1alpha2.HookSpec {
			return func(in interface{}) []kopsv1alpha2.HookSpec {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.HookSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.HookSpec {
						if in == nil {
							return kopsv1alpha2.HookSpec{}
						}
						return ExpandDataSourceHookSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["hooks"]),
		MaxPrice: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["max_price"]),
		SpotDurationInMinutes: func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					return &in
				}(int64(ExpandInt(in)))
			}(in)
		}(in["spot_duration_in_minutes"]),
		CPUCredits: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["cpu_credits"]),
		AssociatePublicIP: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["associate_public_ip"]),
		AdditionalSecurityGroups: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["additional_security_groups"]),
		CloudLabels: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string]string{}
						for key, in := range in {
							out[key] = string(ExpandString(in))
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["cloud_labels"]),
		NodeLabels: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string]string{}
						for key, in := range in {
							out[key] = string(ExpandString(in))
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["node_labels"]),
		FileAssets: func(in interface{}) []kopsv1alpha2.FileAssetSpec {
			return func(in interface{}) []kopsv1alpha2.FileAssetSpec {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.FileAssetSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.FileAssetSpec {
						if in == nil {
							return kopsv1alpha2.FileAssetSpec{}
						}
						return ExpandDataSourceFileAssetSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["file_assets"]),
		Tenancy: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["tenancy"]),
		Kubelet: func(in interface{}) *kopsv1alpha2.KubeletConfigSpec {
			return func(in interface{}) *kopsv1alpha2.KubeletConfigSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.KubeletConfigSpec) *kopsv1alpha2.KubeletConfigSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.KubeletConfigSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceKubeletConfigSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.KubeletConfigSpec{}
				}(in))
			}(in)
		}(in["kubelet"]),
		Taints: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["taints"]),
		MixedInstancesPolicy: func(in interface{}) *kopsv1alpha2.MixedInstancesPolicySpec {
			return func(in interface{}) *kopsv1alpha2.MixedInstancesPolicySpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.MixedInstancesPolicySpec) *kopsv1alpha2.MixedInstancesPolicySpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.MixedInstancesPolicySpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceMixedInstancesPolicySpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.MixedInstancesPolicySpec{}
				}(in))
			}(in)
		}(in["mixed_instances_policy"]),
		CapacityRebalance: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["capacity_rebalance"]),
		AdditionalUserData: func(in interface{}) []kopsv1alpha2.UserData {
			return func(in interface{}) []kopsv1alpha2.UserData {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.UserData
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.UserData {
						if in == nil {
							return kopsv1alpha2.UserData{}
						}
						return ExpandDataSourceUserData(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["additional_user_data"]),
		SuspendProcesses: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["suspend_processes"]),
		ExternalLoadBalancers: func(in interface{}) []kopsv1alpha2.LoadBalancerSpec {
			return func(in interface{}) []kopsv1alpha2.LoadBalancerSpec {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.LoadBalancerSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.LoadBalancerSpec {
						if in == nil {
							return kopsv1alpha2.LoadBalancerSpec{}
						}
						return ExpandDataSourceLoadBalancerSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["external_load_balancers"]),
		DetailedInstanceMonitoring: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["detailed_instance_monitoring"]),
		IAM: func(in interface{}) *kopsv1alpha2.IAMProfileSpec {
			return func(in interface{}) *kopsv1alpha2.IAMProfileSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.IAMProfileSpec) *kopsv1alpha2.IAMProfileSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.IAMProfileSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceIAMProfileSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.IAMProfileSpec{}
				}(in))
			}(in)
		}(in["iam"]),
		SecurityGroupOverride: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["security_group_override"]),
		InstanceProtection: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["instance_protection"]),
		SysctlParameters: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["sysctl_parameters"]),
		RollingUpdate: func(in interface{}) *kopsv1alpha2.RollingUpdate {
			return func(in interface{}) *kopsv1alpha2.RollingUpdate {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.RollingUpdate) *kopsv1alpha2.RollingUpdate {
					return &in
				}(func(in interface{}) kopsv1alpha2.RollingUpdate {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceRollingUpdate(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.RollingUpdate{}
				}(in))
			}(in)
		}(in["rolling_update"]),
		InstanceInterruptionBehavior: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["instance_interruption_behavior"]),
		CompressUserData: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["compress_user_data"]),
		InstanceMetadata: func(in interface{}) *kopsv1alpha2.InstanceMetadataOptions {
			return func(in interface{}) *kopsv1alpha2.InstanceMetadataOptions {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.InstanceMetadataOptions) *kopsv1alpha2.InstanceMetadataOptions {
					return &in
				}(func(in interface{}) kopsv1alpha2.InstanceMetadataOptions {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceInstanceMetadataOptions(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.InstanceMetadataOptions{}
				}(in))
			}(in)
		}(in["instance_metadata"]),
		UpdatePolicy: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["update_policy"]),
		WarmPool: func(in interface{}) *kopsv1alpha2.WarmPoolSpec {
			return func(in interface{}) *kopsv1alpha2.WarmPoolSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.WarmPoolSpec) *kopsv1alpha2.WarmPoolSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.WarmPoolSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceWarmPoolSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.WarmPoolSpec{}
				}(in))
			}(in)
		}(in["warm_pool"]),
		Containerd: func(in interface{}) *kopsv1alpha2.ContainerdConfig {
			return func(in interface{}) *kopsv1alpha2.ContainerdConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.ContainerdConfig) *kopsv1alpha2.ContainerdConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.ContainerdConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceContainerdConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.ContainerdConfig{}
				}(in))
			}(in)
		}(in["containerd"]),
		Packages: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["packages"]),
		GuestAccelerators: func(in interface{}) []kopsv1alpha2.AcceleratorConfig {
			return func(in interface{}) []kopsv1alpha2.AcceleratorConfig {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.AcceleratorConfig
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.AcceleratorConfig {
						if in == nil {
							return kopsv1alpha2.AcceleratorConfig{}
						}
						return ExpandDataSourceAcceleratorConfig(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["guest_accelerators"]),
		MaxInstanceLifetime: func(in interface{}) *meta.Duration {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *meta.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in meta.Duration) *meta.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["max_instance_lifetime"]),
		GCPProvisioningModel: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["gcp_provisioning_model"]),
	}
}

func FlattenDataSourceInstanceGroupSpecInto(in kopsv1alpha2.InstanceGroupSpec, out map[string]interface{}) {
	out["manager"] = func(in kopsv1alpha2.InstanceManager) interface{} {
		return FlattenString(string(in))
	}(in.Manager)
	out["role"] = func(in kopsv1alpha2.InstanceGroupRole) interface{} {
		return FlattenString(string(in))
	}(in.Role)
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
	out["min_size"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.MinSize)
	out["max_size"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.MaxSize)
	out["autoscale"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Autoscale)
	out["autoscale_priority"] = func(in int16) interface{} {
		return FlattenInt(int(in))
	}(in.AutoscalePriority)
	out["machine_type"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.MachineType)
	out["root_volume_size"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.RootVolumeSize)
	out["root_volume_type"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.RootVolumeType)
	out["root_volume_iops"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.RootVolumeIOPS)
	out["root_volume_throughput"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.RootVolumeThroughput)
	out["root_volume_optimization"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.RootVolumeOptimization)
	out["root_volume_delete_on_termination"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.RootVolumeDeleteOnTermination)
	out["root_volume_encryption"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.RootVolumeEncryption)
	out["root_volume_encryption_key"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.RootVolumeEncryptionKey)
	out["volumes"] = func(in []kopsv1alpha2.VolumeSpec) interface{} {
		return func(in []kopsv1alpha2.VolumeSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.VolumeSpec) interface{} {
					return FlattenDataSourceVolumeSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.Volumes)
	out["volume_mounts"] = func(in []kopsv1alpha2.VolumeMountSpec) interface{} {
		return func(in []kopsv1alpha2.VolumeMountSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.VolumeMountSpec) interface{} {
					return FlattenDataSourceVolumeMountSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.VolumeMounts)
	out["subnets"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Subnets)
	out["zones"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Zones)
	out["hooks"] = func(in []kopsv1alpha2.HookSpec) interface{} {
		return func(in []kopsv1alpha2.HookSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.HookSpec) interface{} {
					return FlattenDataSourceHookSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.Hooks)
	out["max_price"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.MaxPrice)
	out["spot_duration_in_minutes"] = func(in *int64) interface{} {
		return func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.SpotDurationInMinutes)
	out["cpu_credits"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.CPUCredits)
	out["associate_public_ip"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.AssociatePublicIP)
	out["additional_security_groups"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AdditionalSecurityGroups)
	out["cloud_labels"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.CloudLabels)
	out["node_labels"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.NodeLabels)
	out["file_assets"] = func(in []kopsv1alpha2.FileAssetSpec) interface{} {
		return func(in []kopsv1alpha2.FileAssetSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.FileAssetSpec) interface{} {
					return FlattenDataSourceFileAssetSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.FileAssets)
	out["tenancy"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Tenancy)
	out["kubelet"] = func(in *kopsv1alpha2.KubeletConfigSpec) interface{} {
		return func(in *kopsv1alpha2.KubeletConfigSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.KubeletConfigSpec) interface{} {
				return func(in kopsv1alpha2.KubeletConfigSpec) []interface{} {
					return []interface{}{FlattenDataSourceKubeletConfigSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Kubelet)
	out["taints"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Taints)
	out["mixed_instances_policy"] = func(in *kopsv1alpha2.MixedInstancesPolicySpec) interface{} {
		return func(in *kopsv1alpha2.MixedInstancesPolicySpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.MixedInstancesPolicySpec) interface{} {
				return func(in kopsv1alpha2.MixedInstancesPolicySpec) []interface{} {
					return []interface{}{FlattenDataSourceMixedInstancesPolicySpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.MixedInstancesPolicy)
	out["capacity_rebalance"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.CapacityRebalance)
	out["additional_user_data"] = func(in []kopsv1alpha2.UserData) interface{} {
		return func(in []kopsv1alpha2.UserData) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.UserData) interface{} {
					return FlattenDataSourceUserData(in)
				}(in))
			}
			return out
		}(in)
	}(in.AdditionalUserData)
	out["suspend_processes"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.SuspendProcesses)
	out["external_load_balancers"] = func(in []kopsv1alpha2.LoadBalancerSpec) interface{} {
		return func(in []kopsv1alpha2.LoadBalancerSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.LoadBalancerSpec) interface{} {
					return FlattenDataSourceLoadBalancerSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.ExternalLoadBalancers)
	out["detailed_instance_monitoring"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.DetailedInstanceMonitoring)
	out["iam"] = func(in *kopsv1alpha2.IAMProfileSpec) interface{} {
		return func(in *kopsv1alpha2.IAMProfileSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.IAMProfileSpec) interface{} {
				return func(in kopsv1alpha2.IAMProfileSpec) []interface{} {
					return []interface{}{FlattenDataSourceIAMProfileSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.IAM)
	out["security_group_override"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.SecurityGroupOverride)
	out["instance_protection"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.InstanceProtection)
	out["sysctl_parameters"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.SysctlParameters)
	out["rolling_update"] = func(in *kopsv1alpha2.RollingUpdate) interface{} {
		return func(in *kopsv1alpha2.RollingUpdate) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.RollingUpdate) interface{} {
				return func(in kopsv1alpha2.RollingUpdate) []interface{} {
					return []interface{}{FlattenDataSourceRollingUpdate(in)}
				}(in)
			}(*in)
		}(in)
	}(in.RollingUpdate)
	out["instance_interruption_behavior"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.InstanceInterruptionBehavior)
	out["compress_user_data"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.CompressUserData)
	out["instance_metadata"] = func(in *kopsv1alpha2.InstanceMetadataOptions) interface{} {
		return func(in *kopsv1alpha2.InstanceMetadataOptions) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.InstanceMetadataOptions) interface{} {
				return func(in kopsv1alpha2.InstanceMetadataOptions) []interface{} {
					return []interface{}{FlattenDataSourceInstanceMetadataOptions(in)}
				}(in)
			}(*in)
		}(in)
	}(in.InstanceMetadata)
	out["update_policy"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.UpdatePolicy)
	out["warm_pool"] = func(in *kopsv1alpha2.WarmPoolSpec) interface{} {
		return func(in *kopsv1alpha2.WarmPoolSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.WarmPoolSpec) interface{} {
				return func(in kopsv1alpha2.WarmPoolSpec) []interface{} {
					return []interface{}{FlattenDataSourceWarmPoolSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.WarmPool)
	out["containerd"] = func(in *kopsv1alpha2.ContainerdConfig) interface{} {
		return func(in *kopsv1alpha2.ContainerdConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.ContainerdConfig) interface{} {
				return func(in kopsv1alpha2.ContainerdConfig) []interface{} {
					return []interface{}{FlattenDataSourceContainerdConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Containerd)
	out["packages"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Packages)
	out["guest_accelerators"] = func(in []kopsv1alpha2.AcceleratorConfig) interface{} {
		return func(in []kopsv1alpha2.AcceleratorConfig) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.AcceleratorConfig) interface{} {
					return FlattenDataSourceAcceleratorConfig(in)
				}(in))
			}
			return out
		}(in)
	}(in.GuestAccelerators)
	out["max_instance_lifetime"] = func(in *meta.Duration) interface{} {
		return func(in *meta.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in meta.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.MaxInstanceLifetime)
	out["gcp_provisioning_model"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.GCPProvisioningModel)
}

func FlattenDataSourceInstanceGroupSpec(in kopsv1alpha2.InstanceGroupSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceInstanceGroupSpecInto(in, out)
	return out
}
