package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kopsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kops"
	resourcesschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

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
			"volumes":                           ComputedList(kopsschemas.DataSourceVolumeSpec()),
			"volume_mounts":                     ComputedList(kopsschemas.DataSourceVolumeMountSpec()),
			"subnets":                           ComputedList(String()),
			"zones":                             ComputedList(String()),
			"hooks":                             ComputedList(kopsschemas.DataSourceHookSpec()),
			"max_price":                         ComputedString(),
			"spot_duration_in_minutes":          ComputedInt(),
			"associate_public_ip":               ComputedBool(),
			"additional_security_groups":        ComputedList(String()),
			"cloud_labels":                      ComputedMap(String()),
			"node_labels":                       ComputedMap(String()),
			"file_assets":                       ComputedList(kopsschemas.DataSourceFileAssetSpec()),
			"tenancy":                           ComputedString(),
			"kubelet":                           ComputedStruct(kopsschemas.DataSourceKubeletConfigSpec()),
			"taints":                            ComputedList(String()),
			"mixed_instances_policy":            ComputedStruct(kopsschemas.DataSourceMixedInstancesPolicySpec()),
			"additional_user_data":              ComputedList(kopsschemas.DataSourceUserData()),
			"suspend_processes":                 ComputedList(String()),
			"external_load_balancers":           ComputedList(kopsschemas.DataSourceLoadBalancer()),
			"detailed_instance_monitoring":      ComputedBool(),
			"iam":                               ComputedStruct(kopsschemas.DataSourceIAMProfileSpec()),
			"security_group_override":           ComputedString(),
			"instance_protection":               ComputedBool(),
			"sysctl_parameters":                 ComputedList(String()),
			"rolling_update":                    ComputedStruct(kopsschemas.DataSourceRollingUpdate()),
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
		InstanceGroup: func(in interface{}) resources.InstanceGroup {
			return resourcesschemas.ExpandDataSourceInstanceGroup(in.(map[string]interface{}))
		}(in),
	}
}

func FlattenDataSourceInstanceGroupInto(in datasources.InstanceGroup, out map[string]interface{}) {
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
	resourcesschemas.FlattenDataSourceInstanceGroupInto(in.InstanceGroup, out)
}

func FlattenDataSourceInstanceGroup(in datasources.InstanceGroup) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceInstanceGroupInto(in, out)
	return out
}
