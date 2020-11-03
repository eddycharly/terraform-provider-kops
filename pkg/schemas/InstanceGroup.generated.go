package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                              RequiredString(),
			"role":                              RequiredString(),
			"image":                             OptionalString(),
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
