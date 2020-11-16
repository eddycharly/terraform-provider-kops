package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceInstanceGroup() *schema.Resource {
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
			"volumes":                           ComputedList(DataSourceVolumeSpec()),
			"volume_mounts":                     ComputedList(DataSourceVolumeMountSpec()),
			"subnets":                           ComputedList(String()),
			"zones":                             ComputedList(String()),
			"hooks":                             ComputedList(DataSourceHookSpec()),
			"max_price":                         ComputedString(),
			"spot_duration_in_minutes":          ComputedInt(),
			"associate_public_ip":               ComputedBool(),
			"additional_security_groups":        ComputedList(String()),
			"cloud_labels":                      ComputedMap(String()),
			"node_labels":                       ComputedMap(String()),
			"file_assets":                       ComputedList(DataSourceFileAssetSpec()),
			"tenancy":                           ComputedString(),
			"kubelet":                           ComputedStruct(DataSourceKubeletConfigSpec()),
			"taints":                            ComputedList(String()),
			"mixed_instances_policy":            ComputedStruct(DataSourceMixedInstancesPolicySpec()),
			"additional_user_data":              ComputedList(DataSourceUserData()),
			"suspend_processes":                 ComputedList(String()),
			"external_load_balancers":           ComputedList(DataSourceLoadBalancer()),
			"detailed_instance_monitoring":      ComputedBool(),
			"iam":                               ComputedStruct(DataSourceIAMProfileSpec()),
			"security_group_override":           ComputedString(),
			"instance_protection":               ComputedBool(),
			"sysctl_parameters":                 ComputedList(String()),
			"rolling_update":                    ComputedStruct(DataSourceRollingUpdate()),
			"instance_interruption_behavior":    ComputedString(),
		},
	}
}

func ExpandDataSourceInstanceGroup(in map[string]interface{}) datasources.InstanceGroup {
	if in == nil {
		panic("expand InstanceGroup failure, in is nil")
	}
	return datasources.InstanceGroup{
		ClusterName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_name"]),
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		InstanceGroupSpec: func(in interface{}) kops.InstanceGroupSpec {
			return ExpandDataSourceInstanceGroupSpec(in.(map[string]interface{}))
		}(in),
	}
}

func FlattenDataSourceInstanceGroupInto(in datasources.InstanceGroup, out map[string]interface{}) {
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	FlattenDataSourceInstanceGroupSpecInto(in.InstanceGroupSpec, out)
}

func FlattenDataSourceInstanceGroup(in datasources.InstanceGroup) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceInstanceGroupInto(in, out)
	return out
}
