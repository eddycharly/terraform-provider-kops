package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceDatasourcesInstanceGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_name":                      RequiredString(),
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

func ExpandDataSourceDatasourcesInstanceGroup(in map[string]interface{}) datasources.InstanceGroup {
	if in == nil {
		panic("expand InstanceGroup failure, in is nil")
	}
	return datasources.InstanceGroup{
		ClusterName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_name"]),
		InstanceGroup: func(in interface{}) resources.InstanceGroup {
			return ExpandDataSourceResourcesInstanceGroup(in.(map[string]interface{}))
		}(in),
	}
}

func FlattenDataSourceDatasourcesInstanceGroupInto(in datasources.InstanceGroup, out map[string]interface{}) {
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
	FlattenDataSourceResourcesInstanceGroupInto(in.InstanceGroup, out)
}

func FlattenDataSourceDatasourcesInstanceGroup(in datasources.InstanceGroup) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceDatasourcesInstanceGroupInto(in, out)
	return out
}
