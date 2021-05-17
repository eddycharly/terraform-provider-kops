package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kopsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kops"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceInstanceGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"revision":                          ComputedInt(),
			"cluster_name":                      ForceNew(RequiredString()),
			"name":                              ForceNew(RequiredString()),
			"role":                              RequiredString(),
			"image":                             OptionalComputedString(),
			"min_size":                          RequiredInt(),
			"max_size":                          RequiredInt(),
			"autoscale":                         OptionalBool(),
			"machine_type":                      RequiredString(),
			"root_volume_size":                  OptionalInt(),
			"root_volume_type":                  OptionalString(),
			"root_volume_iops":                  OptionalInt(),
			"root_volume_throughput":            OptionalInt(),
			"root_volume_optimization":          OptionalBool(),
			"root_volume_delete_on_termination": OptionalBool(),
			"root_volume_encryption":            OptionalBool(),
			"root_volume_encryption_key":        OptionalString(),
			"volumes":                           OptionalList(kopsschemas.ResourceVolumeSpec()),
			"volume_mounts":                     OptionalList(kopsschemas.ResourceVolumeMountSpec()),
			"subnets":                           RequiredList(String()),
			"zones":                             OptionalList(String()),
			"hooks":                             OptionalList(kopsschemas.ResourceHookSpec()),
			"max_price":                         OptionalString(),
			"spot_duration_in_minutes":          OptionalInt(),
			"cpu_credits":                       OptionalString(),
			"associate_public_ip":               OptionalBool(),
			"additional_security_groups":        OptionalList(String()),
			"cloud_labels":                      OptionalMap(String()),
			"node_labels":                       OptionalMap(String()),
			"file_assets":                       OptionalList(kopsschemas.ResourceFileAssetSpec()),
			"tenancy":                           OptionalString(),
			"kubelet":                           OptionalStruct(kopsschemas.ResourceKubeletConfigSpec()),
			"taints":                            OptionalList(String()),
			"mixed_instances_policy":            OptionalStruct(kopsschemas.ResourceMixedInstancesPolicySpec()),
			"additional_user_data":              OptionalList(kopsschemas.ResourceUserData()),
			"suspend_processes":                 OptionalList(String()),
			"external_load_balancers":           OptionalList(kopsschemas.ResourceLoadBalancer()),
			"detailed_instance_monitoring":      OptionalBool(),
			"iam":                               OptionalStruct(kopsschemas.ResourceIAMProfileSpec()),
			"security_group_override":           OptionalString(),
			"instance_protection":               OptionalBool(),
			"sysctl_parameters":                 OptionalList(String()),
			"rolling_update":                    OptionalStruct(kopsschemas.ResourceRollingUpdate()),
			"instance_interruption_behavior":    OptionalString(),
			"compress_user_data":                OptionalBool(),
			"instance_metadata":                 OptionalStruct(kopsschemas.ResourceInstanceMetadataOptions()),
			"update_policy":                     OptionalString(),
		},
	}
}

func ExpandResourceInstanceGroup(in map[string]interface{}) resources.InstanceGroup {
	if in == nil {
		panic("expand InstanceGroup failure, in is nil")
	}
	return resources.InstanceGroup{
		Revision: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["revision"]),
		ClusterName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_name"]),
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		InstanceGroupSpec: func(in interface{}) kops.InstanceGroupSpec {
			return kopsschemas.ExpandResourceInstanceGroupSpec(in.(map[string]interface{}))
		}(in),
	}
}

func FlattenResourceInstanceGroupInto(in resources.InstanceGroup, out map[string]interface{}) {
	out["revision"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.Revision)
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	kopsschemas.FlattenResourceInstanceGroupSpecInto(in.InstanceGroupSpec, out)
}

func FlattenResourceInstanceGroup(in resources.InstanceGroup) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceInstanceGroupInto(in, out)
	return out
}
