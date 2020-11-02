package structures

import (
	"log"
	"reflect"

	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandInstanceGroup(in map[string]interface{}) api.InstanceGroup {
	if in == nil {
		panic("expand InstanceGroup failure, in is nil")
	}
	return api.InstanceGroup{
		Name: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "name", value)
			return value
		}(in["name"]),
		Role: func(in interface{}) kops.InstanceGroupRole {
			value := kops.InstanceGroupRole(ExpandString(in))
			log.Printf("%s - %#v", "role", value)
			return value
		}(in["role"]),
		Image: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "image", value)
			return value
		}(in["image"]),
		MinSize: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "min_size", value)
			return value
		}(in["min_size"]),
		MaxSize: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "max_size", value)
			return value
		}(in["max_size"]),
		MachineType: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "machine_type", value)
			return value
		}(in["machine_type"]),
		RootVolumeSize: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "root_volume_size", value)
			return value
		}(in["root_volume_size"]),
		RootVolumeType: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "root_volume_type", value)
			return value
		}(in["root_volume_type"]),
		RootVolumeIops: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "root_volume_iops", value)
			return value
		}(in["root_volume_iops"]),
		RootVolumeOptimization: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "root_volume_optimization", value)
			return value
		}(in["root_volume_optimization"]),
		RootVolumeDeleteOnTermination: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "root_volume_delete_on_termination", value)
			return value
		}(in["root_volume_delete_on_termination"]),
		RootVolumeEncryption: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "root_volume_encryption", value)
			return value
		}(in["root_volume_encryption"]),
		Volumes: func(in interface{}) []*kops.VolumeSpec {
			value := func(in interface{}) []*kops.VolumeSpec {
				var out []*kops.VolumeSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *kops.VolumeSpec {
						if in == nil {
							return nil
						}
						if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
							return nil
						}
						tmp := func(in kops.VolumeSpec) *kops.VolumeSpec {
							if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
								return nil
							}
							return &in
						}(func(in interface{}) kops.VolumeSpec {
							if in == nil {
								return kops.VolumeSpec{}
							}
							return (ExpandVolumeSpec(in.(map[string]interface{})))
						}(in))
						return tmp
					}(in))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "volumes", value)
			return value
		}(in["volumes"]),
		VolumeMounts: func(in interface{}) []*kops.VolumeMountSpec {
			value := func(in interface{}) []*kops.VolumeMountSpec {
				var out []*kops.VolumeMountSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *kops.VolumeMountSpec {
						if in == nil {
							return nil
						}
						if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
							return nil
						}
						tmp := func(in kops.VolumeMountSpec) *kops.VolumeMountSpec {
							if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
								return nil
							}
							return &in
						}(func(in interface{}) kops.VolumeMountSpec {
							if in == nil {
								return kops.VolumeMountSpec{}
							}
							return (ExpandVolumeMountSpec(in.(map[string]interface{})))
						}(in))
						return tmp
					}(in))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "volume_mounts", value)
			return value
		}(in["volume_mounts"]),
		Subnets: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "subnets", value)
			return value
		}(in["subnets"]),
		Zones: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "zones", value)
			return value
		}(in["zones"]),
		Hooks: func(in interface{}) []kops.HookSpec {
			value := func(in interface{}) []kops.HookSpec {
				var out []kops.HookSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.HookSpec {
						if in == nil {
							return kops.HookSpec{}
						}
						return (ExpandHookSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "hooks", value)
			return value
		}(in["hooks"]),
		MaxPrice: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "max_price", value)
			return value
		}(in["max_price"]),
		SpotDurationInMinutes: func(in interface{}) *int64 {
			value := func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int64) *int64 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int64(ExpandInt(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "spot_duration_in_minutes", value)
			return value
		}(in["spot_duration_in_minutes"]),
		AssociatePublicIP: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "associate_public_ip", value)
			return value
		}(in["associate_public_ip"]),
		AdditionalSecurityGroups: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "additional_security_groups", value)
			return value
		}(in["additional_security_groups"]),
		CloudLabels: func(in interface{}) map[string]string {
			value := func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "cloud_labels", value)
			return value
		}(in["cloud_labels"]),
		NodeLabels: func(in interface{}) map[string]string {
			value := func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "node_labels", value)
			return value
		}(in["node_labels"]),
		FileAssets: func(in interface{}) []kops.FileAssetSpec {
			value := func(in interface{}) []kops.FileAssetSpec {
				var out []kops.FileAssetSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.FileAssetSpec {
						if in == nil {
							return kops.FileAssetSpec{}
						}
						return (ExpandFileAssetSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "file_assets", value)
			return value
		}(in["file_assets"]),
		Tenancy: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "tenancy", value)
			return value
		}(in["tenancy"]),
		Kubelet: func(in interface{}) *kops.KubeletConfigSpec {
			value := func(in interface{}) *kops.KubeletConfigSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.KubeletConfigSpec) *kops.KubeletConfigSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.KubeletConfigSpec {
					if in.([]interface{})[0] == nil {
						return kops.KubeletConfigSpec{}
					}
					return (ExpandKubeletConfigSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			log.Printf("%s - %#v", "kubelet", value)
			return value
		}(in["kubelet"]),
		Taints: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "taints", value)
			return value
		}(in["taints"]),
		MixedInstancesPolicy: func(in interface{}) *kops.MixedInstancesPolicySpec {
			value := func(in interface{}) *kops.MixedInstancesPolicySpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.MixedInstancesPolicySpec) *kops.MixedInstancesPolicySpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.MixedInstancesPolicySpec {
					if in.([]interface{})[0] == nil {
						return kops.MixedInstancesPolicySpec{}
					}
					return (ExpandMixedInstancesPolicySpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			log.Printf("%s - %#v", "mixed_instances_policy", value)
			return value
		}(in["mixed_instances_policy"]),
		AdditionalUserData: func(in interface{}) []kops.UserData {
			value := func(in interface{}) []kops.UserData {
				var out []kops.UserData
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.UserData {
						if in == nil {
							return kops.UserData{}
						}
						return (ExpandUserData(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "additional_user_data", value)
			return value
		}(in["additional_user_data"]),
		SuspendProcesses: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "suspend_processes", value)
			return value
		}(in["suspend_processes"]),
		ExternalLoadBalancers: func(in interface{}) []kops.LoadBalancer {
			value := func(in interface{}) []kops.LoadBalancer {
				var out []kops.LoadBalancer
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.LoadBalancer {
						if in == nil {
							return kops.LoadBalancer{}
						}
						return (ExpandLoadBalancer(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "external_load_balancers", value)
			return value
		}(in["external_load_balancers"]),
		DetailedInstanceMonitoring: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "detailed_instance_monitoring", value)
			return value
		}(in["detailed_instance_monitoring"]),
		IAM: func(in interface{}) *kops.IAMProfileSpec {
			value := func(in interface{}) *kops.IAMProfileSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.IAMProfileSpec) *kops.IAMProfileSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.IAMProfileSpec {
					if in.([]interface{})[0] == nil {
						return kops.IAMProfileSpec{}
					}
					return (ExpandIAMProfileSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			log.Printf("%s - %#v", "iam", value)
			return value
		}(in["iam"]),
		SecurityGroupOverride: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "security_group_override", value)
			return value
		}(in["security_group_override"]),
		InstanceProtection: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "instance_protection", value)
			return value
		}(in["instance_protection"]),
		SysctlParameters: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "sysctl_parameters", value)
			return value
		}(in["sysctl_parameters"]),
		RollingUpdate: func(in interface{}) *kops.RollingUpdate {
			value := func(in interface{}) *kops.RollingUpdate {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.RollingUpdate) *kops.RollingUpdate {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.RollingUpdate {
					if in.([]interface{})[0] == nil {
						return kops.RollingUpdate{}
					}
					return (ExpandRollingUpdate(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			log.Printf("%s - %#v", "rolling_update", value)
			return value
		}(in["rolling_update"]),
		InstanceInterruptionBehavior: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "instance_interruption_behavior", value)
			return value
		}(in["instance_interruption_behavior"]),
	}
}

func FlattenInstanceGroup(in api.InstanceGroup) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "name", value)
			return value
		}(in.Name),
		"role": func(in kops.InstanceGroupRole) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "role", value)
			return value
		}(in.Role),
		"image": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "image", value)
			return value
		}(in.Image),
		"min_size": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			log.Printf("%s - %v", "min_size", value)
			return value
		}(in.MinSize),
		"max_size": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			log.Printf("%s - %v", "max_size", value)
			return value
		}(in.MaxSize),
		"machine_type": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "machine_type", value)
			return value
		}(in.MachineType),
		"root_volume_size": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			log.Printf("%s - %v", "root_volume_size", value)
			return value
		}(in.RootVolumeSize),
		"root_volume_type": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			log.Printf("%s - %v", "root_volume_type", value)
			return value
		}(in.RootVolumeType),
		"root_volume_iops": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			log.Printf("%s - %v", "root_volume_iops", value)
			return value
		}(in.RootVolumeIops),
		"root_volume_optimization": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			log.Printf("%s - %v", "root_volume_optimization", value)
			return value
		}(in.RootVolumeOptimization),
		"root_volume_delete_on_termination": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			log.Printf("%s - %v", "root_volume_delete_on_termination", value)
			return value
		}(in.RootVolumeDeleteOnTermination),
		"root_volume_encryption": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			log.Printf("%s - %v", "root_volume_encryption", value)
			return value
		}(in.RootVolumeEncryption),
		"volumes": func(in []*kops.VolumeSpec) interface{} {
			value := func(in []*kops.VolumeSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in *kops.VolumeSpec) interface{} {
						if in == nil {
							return nil
						}
						return func(in kops.VolumeSpec) interface{} {
							return func(in kops.VolumeSpec) interface{} {
								return FlattenVolumeSpec(in)
							}(in)
						}(*in)
					}(in))
				}
				return out
			}(in)
			log.Printf("%s - %v", "volumes", value)
			return value
		}(in.Volumes),
		"volume_mounts": func(in []*kops.VolumeMountSpec) interface{} {
			value := func(in []*kops.VolumeMountSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in *kops.VolumeMountSpec) interface{} {
						if in == nil {
							return nil
						}
						return func(in kops.VolumeMountSpec) interface{} {
							return func(in kops.VolumeMountSpec) interface{} {
								return FlattenVolumeMountSpec(in)
							}(in)
						}(*in)
					}(in))
				}
				return out
			}(in)
			log.Printf("%s - %v", "volume_mounts", value)
			return value
		}(in.VolumeMounts),
		"subnets": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			log.Printf("%s - %v", "subnets", value)
			return value
		}(in.Subnets),
		"zones": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			log.Printf("%s - %v", "zones", value)
			return value
		}(in.Zones),
		"hooks": func(in []kops.HookSpec) interface{} {
			value := func(in []kops.HookSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.HookSpec) interface{} {
						return FlattenHookSpec(in)
					}(in))
				}
				return out
			}(in)
			log.Printf("%s - %v", "hooks", value)
			return value
		}(in.Hooks),
		"max_price": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			log.Printf("%s - %v", "max_price", value)
			return value
		}(in.MaxPrice),
		"spot_duration_in_minutes": func(in *int64) interface{} {
			value := func(in *int64) interface{} {
				if in == nil {
					return nil
				}
				return func(in int64) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			log.Printf("%s - %v", "spot_duration_in_minutes", value)
			return value
		}(in.SpotDurationInMinutes),
		"associate_public_ip": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			log.Printf("%s - %v", "associate_public_ip", value)
			return value
		}(in.AssociatePublicIP),
		"additional_security_groups": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			log.Printf("%s - %v", "additional_security_groups", value)
			return value
		}(in.AdditionalSecurityGroups),
		"cloud_labels": func(in map[string]string) interface{} {
			value := func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
			log.Printf("%s - %v", "cloud_labels", value)
			return value
		}(in.CloudLabels),
		"node_labels": func(in map[string]string) interface{} {
			value := func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
			log.Printf("%s - %v", "node_labels", value)
			return value
		}(in.NodeLabels),
		"file_assets": func(in []kops.FileAssetSpec) interface{} {
			value := func(in []kops.FileAssetSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.FileAssetSpec) interface{} {
						return FlattenFileAssetSpec(in)
					}(in))
				}
				return out
			}(in)
			log.Printf("%s - %v", "file_assets", value)
			return value
		}(in.FileAssets),
		"tenancy": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "tenancy", value)
			return value
		}(in.Tenancy),
		"kubelet": func(in *kops.KubeletConfigSpec) interface{} {
			value := func(in *kops.KubeletConfigSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KubeletConfigSpec) interface{} {
					return func(in kops.KubeletConfigSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenKubeletConfigSpec(in)}
					}(in)
				}(*in)
			}(in)
			log.Printf("%s - %v", "kubelet", value)
			return value
		}(in.Kubelet),
		"taints": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			log.Printf("%s - %v", "taints", value)
			return value
		}(in.Taints),
		"mixed_instances_policy": func(in *kops.MixedInstancesPolicySpec) interface{} {
			value := func(in *kops.MixedInstancesPolicySpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.MixedInstancesPolicySpec) interface{} {
					return func(in kops.MixedInstancesPolicySpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenMixedInstancesPolicySpec(in)}
					}(in)
				}(*in)
			}(in)
			log.Printf("%s - %v", "mixed_instances_policy", value)
			return value
		}(in.MixedInstancesPolicy),
		"additional_user_data": func(in []kops.UserData) interface{} {
			value := func(in []kops.UserData) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.UserData) interface{} {
						return FlattenUserData(in)
					}(in))
				}
				return out
			}(in)
			log.Printf("%s - %v", "additional_user_data", value)
			return value
		}(in.AdditionalUserData),
		"suspend_processes": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			log.Printf("%s - %v", "suspend_processes", value)
			return value
		}(in.SuspendProcesses),
		"external_load_balancers": func(in []kops.LoadBalancer) interface{} {
			value := func(in []kops.LoadBalancer) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.LoadBalancer) interface{} {
						return FlattenLoadBalancer(in)
					}(in))
				}
				return out
			}(in)
			log.Printf("%s - %v", "external_load_balancers", value)
			return value
		}(in.ExternalLoadBalancers),
		"detailed_instance_monitoring": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			log.Printf("%s - %v", "detailed_instance_monitoring", value)
			return value
		}(in.DetailedInstanceMonitoring),
		"iam": func(in *kops.IAMProfileSpec) interface{} {
			value := func(in *kops.IAMProfileSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.IAMProfileSpec) interface{} {
					return func(in kops.IAMProfileSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenIAMProfileSpec(in)}
					}(in)
				}(*in)
			}(in)
			log.Printf("%s - %v", "iam", value)
			return value
		}(in.IAM),
		"security_group_override": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			log.Printf("%s - %v", "security_group_override", value)
			return value
		}(in.SecurityGroupOverride),
		"instance_protection": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			log.Printf("%s - %v", "instance_protection", value)
			return value
		}(in.InstanceProtection),
		"sysctl_parameters": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			log.Printf("%s - %v", "sysctl_parameters", value)
			return value
		}(in.SysctlParameters),
		"rolling_update": func(in *kops.RollingUpdate) interface{} {
			value := func(in *kops.RollingUpdate) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.RollingUpdate) interface{} {
					return func(in kops.RollingUpdate) []map[string]interface{} {
						return []map[string]interface{}{FlattenRollingUpdate(in)}
					}(in)
				}(*in)
			}(in)
			log.Printf("%s - %v", "rolling_update", value)
			return value
		}(in.RollingUpdate),
		"instance_interruption_behavior": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			log.Printf("%s - %v", "instance_interruption_behavior", value)
			return value
		}(in.InstanceInterruptionBehavior),
	}
}
