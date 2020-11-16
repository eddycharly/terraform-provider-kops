package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func InstanceGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                              RequiredString(),
			"role":                              RequiredString(),
			"image":                             OptionalComputedString(),
			"min_size":                          RequiredInt(),
			"max_size":                          RequiredInt(),
			"machine_type":                      RequiredString(),
			"root_volume_size":                  OptionalInt(),
			"root_volume_type":                  OptionalString(),
			"root_volume_iops":                  OptionalInt(),
			"root_volume_optimization":          OptionalBool(),
			"root_volume_delete_on_termination": OptionalBool(),
			"root_volume_encryption":            OptionalBool(),
			"volumes":                           OptionalList(VolumeSpec()),
			"volume_mounts":                     OptionalList(VolumeMountSpec()),
			"subnets":                           RequiredList(String()),
			"zones":                             OptionalList(String()),
			"hooks":                             OptionalList(HookSpec()),
			"max_price":                         OptionalString(),
			"spot_duration_in_minutes":          OptionalInt(),
			"associate_public_ip":               OptionalBool(),
			"additional_security_groups":        OptionalList(String()),
			"cloud_labels":                      OptionalMap(String()),
			"node_labels":                       OptionalMap(String()),
			"file_assets":                       OptionalList(FileAssetSpec()),
			"tenancy":                           OptionalString(),
			"kubelet":                           OptionalStruct(KubeletConfigSpec()),
			"taints":                            OptionalList(String()),
			"mixed_instances_policy":            OptionalStruct(MixedInstancesPolicySpec()),
			"additional_user_data":              OptionalList(UserData()),
			"suspend_processes":                 OptionalList(String()),
			"external_load_balancers":           OptionalList(LoadBalancer()),
			"detailed_instance_monitoring":      OptionalBool(),
			"iam":                               OptionalStruct(IAMProfileSpec()),
			"security_group_override":           OptionalString(),
			"instance_protection":               OptionalBool(),
			"sysctl_parameters":                 OptionalList(String()),
			"rolling_update":                    OptionalStruct(RollingUpdate()),
			"instance_interruption_behavior":    OptionalString(),
		},
	}
}

func ExpandInstanceGroup(in map[string]interface{}) api.InstanceGroup {
	if in == nil {
		panic("expand InstanceGroup failure, in is nil")
	}
	return api.InstanceGroup{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		InstanceGroupSpec: func(in interface{}) kops.InstanceGroupSpec {
			return func(in interface{}) kops.InstanceGroupSpec {
				if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
					return kops.InstanceGroupSpec{}
				}
				return (ExpandInstanceGroupSpec(in.([]interface{})[0].(map[string]interface{})))
			}(in)
		}(in),
	}
}

func FlattenInstanceGroupInto(in api.InstanceGroup, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	FlattenInstanceGroupSpecInto(in.InstanceGroupSpec, out)
}

func FlattenInstanceGroup(in api.InstanceGroup) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenInstanceGroupInto(in, out)
	return out
}
