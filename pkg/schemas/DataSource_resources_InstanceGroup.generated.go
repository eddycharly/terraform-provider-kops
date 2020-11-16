package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceResourcesInstanceGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                              RequiredString(),
			"role":                              ComputedString(),
			"image":                             ComputedString(),
			"min_size":                          ComputedInt(),
			"max_size":                          ComputedInt(),
			"machine_type":                      ComputedString(),
			"root_volume_size":                  ComputedInt(),
			"root_volume_type":                  ComputedString(),
			"root_volume_iops":                  ComputedInt(),
			"root_volume_optimization":          ComputedBool(),
			"root_volume_delete_on_termination": ComputedBool(),
			"root_volume_encryption":            ComputedBool(),
			"volumes":                           ComputedList(DataSourceKopsVolumeSpec()),
			"volume_mounts":                     ComputedList(DataSourceKopsVolumeMountSpec()),
			"subnets":                           ComputedList(String()),
			"zones":                             ComputedList(String()),
			"hooks":                             ComputedList(DataSourceKopsHookSpec()),
			"max_price":                         ComputedString(),
			"spot_duration_in_minutes":          ComputedInt(),
			"associate_public_ip":               ComputedBool(),
			"additional_security_groups":        ComputedList(String()),
			"cloud_labels":                      ComputedMap(String()),
			"node_labels":                       ComputedMap(String()),
			"file_assets":                       ComputedList(DataSourceKopsFileAssetSpec()),
			"tenancy":                           ComputedString(),
			"kubelet":                           ComputedStruct(DataSourceKopsKubeletConfigSpec()),
			"taints":                            ComputedList(String()),
			"mixed_instances_policy":            ComputedStruct(DataSourceKopsMixedInstancesPolicySpec()),
			"additional_user_data":              ComputedList(DataSourceKopsUserData()),
			"suspend_processes":                 ComputedList(String()),
			"external_load_balancers":           ComputedList(DataSourceKopsLoadBalancer()),
			"detailed_instance_monitoring":      ComputedBool(),
			"iam":                               ComputedStruct(DataSourceKopsIAMProfileSpec()),
			"security_group_override":           ComputedString(),
			"instance_protection":               ComputedBool(),
			"sysctl_parameters":                 ComputedList(String()),
			"rolling_update":                    ComputedStruct(DataSourceKopsRollingUpdate()),
			"instance_interruption_behavior":    ComputedString(),
		},
	}
}

func ExpandDataSourceResourcesInstanceGroup(in map[string]interface{}) resources.InstanceGroup {
	if in == nil {
		panic("expand InstanceGroup failure, in is nil")
	}
	return resources.InstanceGroup{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		InstanceGroupSpec: func(in interface{}) kops.InstanceGroupSpec {
			return ExpandDataSourceKopsInstanceGroupSpec(in.(map[string]interface{}))
		}(in),
	}
}

func FlattenDataSourceResourcesInstanceGroupInto(in resources.InstanceGroup, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	FlattenDataSourceKopsInstanceGroupSpecInto(in.InstanceGroupSpec, out)
}

func FlattenDataSourceResourcesInstanceGroup(in resources.InstanceGroup) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceResourcesInstanceGroupInto(in, out)
	return out
}
