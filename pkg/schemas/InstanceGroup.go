package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                              StringRequired(),
			"role":                              StringEnumRequired("Master", "Node", "Bastion"),
			"image":                             StringOptionalComputed(),
			"min_size":                          IntRequired(),
			"max_size":                          IntRequired(),
			"machine_type":                      StringRequired(),
			"root_volume_size":                  IntOptionalComputed(),
			"root_volume_type":                  StringOptionalComputed(),
			"root_volume_iops":                  IntOptionalComputed(),
			"root_volume_optimization":          BoolOptionalComputed(),
			"root_volume_delete_on_termination": BoolOptionalComputed(),
			"root_volume_encryption":            BoolOptionalComputed(),
			"volumes":                           ListOptional(VolumeSpec()),
			"volume_mounts":                     ListOptional(VolumeMountSpec()),
			"subnets":                           ListRequired(String()),
			"zones":                             ListOptional(String()),
			"hooks":                             ListOptional(HookSpec()),
			"max_price":                         StringOptional(),
			"spot_duration_in_minutes":          IntOptional(),
			"associate_public_ip":               BoolOptionalComputed(),
			"additional_security_groups":        ListOptional(String()),
			"cloud_labels":                      MapOptionalComputed(String()),
			"node_labels":                       MapOptionalComputed(String()),
			"file_assets":                       ListOptional(FileAssetSpec()),
			"tenancy":                           StringOptionalComputed(),
			"kubelet":                           StructOptionalComputed(KubeletConfigSpec()),
			"taints":                            ListOptional(String()),
			"mixed_instances_policy":            StructOptionalComputed(MixedInstancesPolicySpec()),
			"additional_user_data":              ListOptional(UserData()),
			"suspend_processes":                 ListOptional(String()),
			"external_load_balancers":           ListOptional(LoadBalancer()),
			"detailed_instance_monitoring":      BoolOptionalComputed(),
			"iam":                               StructOptionalComputed(IAMProfileSpec()),
			"security_group_override":           StringOptionalComputed(),
			"instance_protection":               BoolOptionalComputed(),
			"sysctl_parameters":                 ListOptional(String()),
			"rolling_update":                    StructOptionalComputed(RollingUpdate()),
			"instance_interruption_behavior":    StringOptional(),
		},
	}
}
