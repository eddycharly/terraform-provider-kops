package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kopsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kops"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
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
			"autoscale":                         ComputedBool(),
			"machine_type":                      ComputedString(),
			"root_volume_size":                  ComputedInt(),
			"root_volume_type":                  ComputedString(),
			"root_volume_iops":                  ComputedInt(),
			"root_volume_throughput":            ComputedInt(),
			"root_volume_optimization":          ComputedBool(),
			"root_volume_delete_on_termination": ComputedBool(),
			"root_volume_encryption":            ComputedBool(),
			"root_volume_encryption_key":        ComputedString(),
			"volumes":                           ComputedList(kopsschemas.DataSourceVolumeSpec()),
			"volume_mounts":                     ComputedList(kopsschemas.DataSourceVolumeMountSpec()),
			"subnets":                           ComputedList(String()),
			"zones":                             ComputedList(String()),
			"hooks":                             ComputedList(kopsschemas.DataSourceHookSpec()),
			"max_price":                         ComputedString(),
			"spot_duration_in_minutes":          ComputedInt(),
			"cpu_credits":                       ComputedString(),
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
			"compress_user_data":                ComputedBool(),
			"instance_metadata":                 ComputedStruct(kopsschemas.DataSourceInstanceMetadataOptions()),
			"update_policy":                     ComputedString(),
		},
	}
}

func ExpandDataSourceInstanceGroup(in map[string]interface{}) resources.InstanceGroup {
	if in == nil {
		panic("expand InstanceGroup failure, in is nil")
	}
	return resources.InstanceGroup{
		ClusterName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_name"]),
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		InstanceGroupSpec: func(in interface{}) kops.InstanceGroupSpec {
			return kopsschemas.ExpandDataSourceInstanceGroupSpec(in.(map[string]interface{}))
		}(in),
	}
}

func FlattenDataSourceInstanceGroupInto(in resources.InstanceGroup, out map[string]interface{}) {
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	kopsschemas.FlattenDataSourceInstanceGroupSpecInto(in.InstanceGroupSpec, out)
}

func FlattenDataSourceInstanceGroup(in resources.InstanceGroup) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceInstanceGroupInto(in, out)
	return out
}
